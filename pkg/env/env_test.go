package env

import (
	"os"
	"testing"
)


func TestReadEnvs(t *testing.T) {

	os.Setenv("SERVICE_PORT", "80")
	os.Setenv("DATABASE_DSN", "localhost user=postgres password=postgres dbname=marginDB port=5433 sslmode=disable TimeZone=Asia/Tehran")

	envs := ReadEnvs()

	if envs.SERVICE_PORT != "80" {
		t.Errorf("Expected service port to be 'true', got %s", envs.SERVICE_PORT)
	}
	if envs.DATABASE_DSN != "localhost user=postgres password=postgres dbname=marginDB port=5433 sslmode=disable TimeZone=Asia/Tehran" {
		t.Errorf("Expected database dsn to be 'test', got %s", envs.DATABASE_DSN)
	}

	os.Unsetenv("SERVICE_PORT")
	os.Unsetenv("DATABASE_DSN")

}
