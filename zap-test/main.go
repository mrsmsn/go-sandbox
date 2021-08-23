package main

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	err := errors.New(fmt.Sprintf("Error: %s", "errors.New"))
	logger.Info("Hello Error zap", zap.Error(err))
}
