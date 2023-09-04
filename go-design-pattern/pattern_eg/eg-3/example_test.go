package example

import "testing"

func TestNewLogManager(t *testing.T) {
	fileLogging := &FileLogging{}
	logManager := NewLogManager(fileLogging)

	logManager.Info()
	logManager.Error()

	dbLogging := &DBLogging{}
	logManager.Logging = dbLogging
	logManager.Info()

	logManager.Error()
}
