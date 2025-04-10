package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/codersgyan/expressify/internal/cli_model"
	"github.com/codersgyan/expressify/internal/structure"
)

func main() {
	// Set up logging to a file
	logFile, err := os.OpenFile("expressify.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Include timestamp and file info

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		errMsg := fmt.Errorf("failed to get current working directory: %w", err)
		log.Printf("Error at %s: %v", time.Now().Format(time.RFC3339), errMsg)
		fmt.Printf("Error: %v\n", errMsg)
		os.Exit(1)
	}
	log.Printf("Current working directory retrieved: %s", cwd)

	// Define source and destination paths for the directory copy
	srcPath := cwd + "/.templates/jsbase"
	dstPath := cwd + "/.expressify/auth-service"

	// Copy the directory
	cpErr := structure.CopyDir(srcPath, dstPath)
	if cpErr != nil {
		errMsg := fmt.Errorf("failed to copy directory from %s to %s: %w", srcPath, dstPath, cpErr)
		log.Printf("Error at %s: %v", time.Now().Format(time.RFC3339), errMsg)
		fmt.Printf("Error copying directory: %v\n", errMsg)
		os.Exit(1)
	}
	log.Printf("Directory copied successfully from %s to %s", srcPath, dstPath)
	fmt.Println("Directory copied successfully.")

	// Run the CLI program using bubbletea
	p := tea.NewProgram(cli_model.InitialModel())
	if _, err := p.Run(); err != nil {
		errMsg := fmt.Errorf("failed to run CLI program: %w", err)
		log.Printf("Error at %s: %v", time.Now().Format(time.RFC3339), errMsg)
		fmt.Printf("Error running CLI: %v\n", errMsg)
		os.Exit(1)
	}
	log.Printf("CLI program completed successfully")
	fmt.Println("CLI completed successfully")
}