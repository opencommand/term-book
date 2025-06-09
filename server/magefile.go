//go:build mage

package main

// Copyright 2025 The OpenCmd Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// Run `mage setup` to setup the project.
//
// Run `mage build` to build the project.
//
// Run `mage clean` to clean the project.

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/caarlos0/log"
	"github.com/magefile/mage/sh"
)

// Clean cleans the project.
func Clean() error {
	log.Info("Start cleaning...")
	goreleaser := exec.Command("goreleaser", "build", "--clean")
	goreleaser.Stdout = os.Stdout
	goreleaser.Stderr = os.Stderr
	err := goreleaser.Run()
	if err != nil {
		log.WithError(err).Fatal("goreleaser build failed")
	}
	dirs := []string{"dist", "tmp"}
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("clean %s failed: %v", dir, err)
		}
	}
	return nil
}

// Build builds the project.
func Build() error {
	log.Info("Start building...")

	goreleaser := exec.Command("goreleaser", "release", "--snapshot", "--clean")
	goreleaser.Stdout = os.Stdout
	goreleaser.Stderr = os.Stderr
	err := goreleaser.Run()
	if err != nil {
		log.WithError(err).Fatal("goreleaser release failed")
	}
	return nil
}

// Lint runs the linter
func Lint() error {
	log.Info("Running linter...")
	cmd := exec.Command("golangci-lint", "run", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.WithError(err).Fatal("lint failed")
	}
	return nil
}

// Format formats the code
func Format() error {
	log.Info("Formatting code...")
	return sh.Run("goimports", "-w", ".")
}

func Test() error {
	log.Info("Running tests...")
	cmd := exec.Command("go", "test", "./...", "-coverprofile=coverage.txt", "-covermode=atomic")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.WithError(err).Fatal("test failed")
	}
	return nil
}
