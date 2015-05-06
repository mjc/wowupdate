package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GitUpdate(addon Addon) (err error) {
	log.Printf("Using git to update %v", addon.title)
	originalWorkingDirectory, _ := os.Getwd()

	repo := filepath.Join(getWowPath(), addon.path)
	os.Chdir(repo)

	cmd := exec.Command("git", "pull")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()

	os.Chdir(originalWorkingDirectory)

	return err
}
