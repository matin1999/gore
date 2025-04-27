package commands


import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gore/pkg/env"
)

var addCurrencyCmd = &cobra.Command{
	Use:   "AddCurrency",
	Short: "AddCurrency ",
	Long:  `run this command like this AddCurrency`,

	Run: func(cmd *cobra.Command, args []string) {
		envs := env.ReadEnvs()
		db, dbErr := db.Init(envs.PG_DSN, jobLogger)
		if dbErr != nil {
			fmt.Println(fmt.Sprintf("failed err to initilize postgres db %s", dbErr))
			os.Exit(1)
		}
		mongoDatabase, MongoErr := mongo.Init(&envs, jobLogger)
		if MongoErr != nil {
			jobLogger.Log("panic", fmt.Sprintf("[tw-gc-pakban-1.2] cannot connect to monogdb %s", MongoErr))
			panic(MongoErr)
		}
		storage, storageErr := minio.Init(envs.MINIO_CREDS_PATHS, jobLogger)
		if storageErr != nil {
			jobLogger.Log("panic", fmt.Sprintf("cannot create storage clients %s", storageErr))
			panic(storageErr)
		}
		detectionInfo, err := GetDetectionInfo(storage)
		if err != nil {
			os.Exit(1)
		}
		time.Sleep(time.Second * 10)

		kafka := kafka.NewProducer(envs.KAFKA_TOPIC, strings.Split(envs.KAFKA_BROKERS, ","))

		jobsHandler := jobs.DetectionInit(jobLogger, db, mongoDatabase, storage, envs,kafka)
		err = jobsHandler.CreateDetectionJob(&detectionInfo)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to insert to job db %s", err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addDetectionJobCmd)
}

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptGetSelect(pc promptContent, storage minio.StorageInterface) string {
	items := storage.GetListOfStorageNames()
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func GetDetectionInfo(storageClient minio.StorageInterface) (db.DetectionJob, error) {
	fromTimePromptContent := promptContent{
		"provide a job.from unix time",
		"start time of checking can be deleting ts files",
	}
	fromStr := promptGetInput(fromTimePromptContent)

	from, err := strconv.ParseInt(fromStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid from time:", err)
		return db.DetectionJob{}, fmt.Errorf("invalid from time")
	}

	toTimePromptContent := promptContent{
		"provide a job.to unix time",
		"end time of checking can be deleting ts files",
	}
	toStr := promptGetInput(toTimePromptContent)

	to, err := strconv.ParseInt(toStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid to time:", err)
		return db.DetectionJob{}, fmt.Errorf("invalid to time")
	}

	bucketPromptContent := promptContent{
		"provide a bucket name",
		"name of bucket to be clean",
	}
	bucket := promptGetInput(bucketPromptContent)

	if bucket == "" {
		fmt.Println("Invalid bucket name: bucket name cannot be empty")
		return db.DetectionJob{}, fmt.Errorf("invalid bucket name: bucket name cannot be empty")
	}

	storagePromptContent := promptContent{
		"Please provide a storage.",
		fmt.Sprintf("What storage does %s belong to?", bucket),
	}
	storage := promptGetSelect(storagePromptContent, storageClient)
	if storage == "" {
		fmt.Println("Invalid storage: storage cannot be empty")
		return db.DetectionJob{}, fmt.Errorf("invalid storage: storage cannot be empty")
	}
	job := db.DetectionJob{From: from, To: to, Storage: storage, Bucket: bucket}
	confirmationPrompt := promptContent{
		"Please provide (yes/no)",
		fmt.Sprintf("are you sure about this job from: %v to: %v bucket: %s storage: %s (yes/no)", job.From, job.To, job.Bucket, job.Storage),
	}
	confirmation := promptGetInput(confirmationPrompt)
	if confirmation == "yes" {
		return job, nil
	}
	return db.DetectionJob{}, fmt.Errorf("terminated")
}
