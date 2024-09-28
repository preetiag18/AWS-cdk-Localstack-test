package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func deploy(t *testing.T) {
	currentDir, err := os.Getwd()
	failOnError(t, err, "Get current directory")
	parentDir := filepath.Dir(currentDir)
	bootstrapCmd := exec.Command("cdklocal", "bootstrap")
	bootstrapCmd.Stdout = os.Stderr
	bootstrapCmd.Stderr = os.Stderr
	bootstrapCmd.Dir = parentDir
	err = bootstrapCmd.Run()
	failOnError(t, err, "CDK bootstrap")
	deployCmd := exec.Command("cdklocal", "deploy", "DataStack", "UserStack", "--require-approval", "never")
	deployCmd.Stdout = os.Stdout
	deployCmd.Stderr = os.Stderr
	deployCmd.Dir = parentDir
	err = deployCmd.Run()
	failOnError(t, err, "CDK deploy")
}
