package config_test

import (
	"awesomeProject/book/v2/config"
	"fmt"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := config.LoadConfigfromYaml(fmt.Sprintf("%s/book/v2/application.yaml", os.Getenv("workspaceFolder")))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())

}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("workspaceFolder", "localhost")
	err := config.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}
