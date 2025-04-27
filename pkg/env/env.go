package env

import (
	"os"
)

type Envs struct {
	SERVICE_PORT string
	DATABASE_DSN string

}

func ReadEnvs() Envs {
	envs := Envs{}
	envs.SERVICE_PORT = os.Getenv("SERVICE_PORT")
	envs.DATABASE_DSN = os.Getenv("DATABASE_DSN")

	return envs
}
