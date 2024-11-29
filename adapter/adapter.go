package main

import (
	"fmt"
	"strings"
)

type Logger interface {
	LogInformation(message string)
}

type CurrentLogger struct{}

func (logger CurrentLogger) LogInformation(message string) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Current Logger: " + message)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

type NewLogger struct{}

func (logger NewLogger) LogInfo(msg string) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("New Logger: " + msg)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

type NewToCurrentAdapter struct {
	NewLogger *NewLogger
}

func (ntc *NewToCurrentAdapter) LogInformation(message string) {
	ntc.NewLogger.LogInfo(message)
}
