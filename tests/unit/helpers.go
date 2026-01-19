package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func getTerraformWorkingDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func readJsonFile(filePath string, result interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}
	return nil
}

func writeJsonFile(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getTerraformBinaryPath() (string, error) {
	terraformBinaryPath := os.Getenv("TERRAFORM_BINARY_PATH")
	if terraformBinaryPath == "" {
		terraformBinaryPath = "terraform"
	}
	return terraformBinaryPath, nil
}

func newTerraformExecutor(workingDir string) (*tfexec.Terraform, error) {
	tfBinaryPath, err := getTerraformBinaryPath()
	if err != nil {
		return nil, err
	}
	tf, err := tfexec.NewTerraform(workingDir, tfBinaryPath)
	if err != nil {
		return nil, err
	}
	return tf, nil
}

func initTerraform(workingDir string) (*tfexec.Terraform, error) {
	tf, err := newTerraformExecutor(workingDir)
	if err != nil {
		return nil, err
	}
	err = tf.Init()
	if err != nil {
		return nil, err
	}
	return tf, nil
}

func applyTerraform(workingDir string) (*tfexec.Terraform, error) {
	tf, err := initTerraform(workingDir)
	if err != nil {
		return nil, err
	}
	err = tf.Apply()
	if err != nil {
		return nil, err
	}
	return tf, nil
}

func destroyTerraform(workingDir string) (*tfexec.Terraform, error) {
	tf, err := initTerraform(workingDir)
	if err != nil {
		return nil, err
	}
	err = tf.Destroy()
	if err != nil {
		return nil, err
	}
	return tf, nil
}

func getProjectRoot() (string, error) {
	wd, err := getTerraformWorkingDir()
	if err != nil {
		return "", err
	}
	root := wd
	for {
		_, err := os.Stat(filepath.Join(root, "main.tf"))
		if err == nil {
			return root, nil
		}
		parentDir := filepath.Dir(root)
		if parentDir == root {
			return "", errors.New("could not find project root")
		}
		root = parentDir
	}
}