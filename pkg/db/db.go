package db

import (

	"gore/pkg/env"
	"tw-pakban/pkg/mongo"
)

type DataBaseInterface interface {
	CreateCurrency(job *DetectionJob, channels []Channel) error
	
}

type DataBaseWrapper struct {
	PG     *gorm.DB
	logger *logger.Logger
}

func Init(dsn string, log *logger.Logger) (DataBaseInterface, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if db.AutoMigrate(&DetectionJob{}, &ChannelDetection{}, &TsFile{}, &DisallowedEpisode{}, &RemoveBucketJob{}, &RecoverMongoJob{}) != nil {
		return nil, err
	}
	return &DataBaseWrapper{PG: db, logger: log}, nil

}