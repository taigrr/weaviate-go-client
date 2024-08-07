package test

import (
	"os"
	"os/exec"

	"github.com/weaviate/weaviate-go-client/v4/test/testsuit"
)

// SetupWeaviate run docker compose up
func SetupWeaviate() error {
	app := "docker"
	arguments := []string{
		"compose",
		"up",
		"-d",
	}
	_, _, authEnabled := testsuit.GetPortAndAuthPw()
	if authEnabled {
		arguments = append(arguments, "-f", "docker-compose-wcs.yml")
	}

	return command(app, arguments)
}

// TearDownWeaviate run docker-compose down
func TearDownWeaviate() error {
	app := "docker"
	arguments := []string{
		"compose",
		"down",
		"--remove-orphans",
	}
	return command(app, arguments)
}

func command(app string, arguments []string) error {
	mydir, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := exec.Command(app, arguments...)
	execDir := mydir + "/../"
	cmd.Dir = execDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	return err
}
