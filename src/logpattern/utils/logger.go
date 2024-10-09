/*
 * Copyright (c) Motadata 2024.  All rights reserved.
 */

package utils

import (
	"os"
)

type Logger struct {
	component MotadataString

	module MotadataString
}

var logLevel = LogLevelInfo

func TraceEnabled() bool {

	return LogLevelTrace >= logLevel
}

func DebugEnabled() bool {

	return LogLevelDebug >= logLevel
}

func SetLogLevel(value int) {

	logLevel = value
}

func NewLogger(component MotadataString, module MotadataString) Logger {

	return Logger{component: component, module: module}
}

func (logger *Logger) Trace(message MotadataString) {

	if TraceEnabled() {

		currentDate := MotadataTimeString(LogFileDateFormat).Format()

		currentTime := MotadataTimeString(LogFileTimeFormat).Format()

		message := currentDate + SpaceSeparator + currentTime + SpaceSeparator + "TRACE [" + logger.component + "]:" +
			message

		logger.write(message.ToString())

	}
}

func (logger *Logger) Debug(message MotadataString) {

	if DebugEnabled() {

		currentDate := MotadataTimeString(LogFileDateFormat).Format()

		currentTime := MotadataTimeString(LogFileTimeFormat).Format()

		message := currentDate + SpaceSeparator + currentTime + SpaceSeparator + "DEBUG [" + logger.component + "]:" +
			message

		logger.write(message.ToString())

	}
}

func (logger *Logger) Info(message MotadataString) {

	currentDate := MotadataTimeString(LogFileDateFormat).Format()

	currentTime := MotadataTimeString(LogFileTimeFormat).Format()

	message = currentDate + SpaceSeparator + currentTime + SpaceSeparator + "INFO [" + logger.component + "]:" +
		message

	logger.write(message.ToString())

}

func (logger *Logger) Warn(message MotadataString) {

	currentDate := MotadataTimeString(LogFileDateFormat).Format()

	currentTime := MotadataTimeString(LogFileTimeFormat).Format()

	message = currentDate + SpaceSeparator + currentTime + SpaceSeparator + "WARN [" + logger.component + "]:" + message

	logger.write(message.ToString())

}

func (logger *Logger) Fatal(message MotadataString) {

	currentDate := MotadataTimeString(LogFileDateFormat).Format()

	currentTime := MotadataTimeString(LogFileTimeFormat).Format()

	message = currentDate + SpaceSeparator + currentTime + SpaceSeparator + "FATAL [" + logger.component + "]:" +
		message

	logger.write(message.ToString())

}

func (logger *Logger) write(message string) {

	logDir := CurrentDir + PathSeparator + LogDirectory + PathSeparator

	_, err := os.Stat(logDir)

	if os.IsNotExist(err) {

		err := os.MkdirAll(logDir, 0755)

		if err != nil {

			panic(err)
		}

	}

	logFile := logDir + PathSeparator

	currentDate := MotadataTimeString(LogFileDateFormat).Format()

	if logger.module.IsNotEmpty() {

		logFile = logFile + MotadataString(LogFile).ReplaceAll("@@@", currentDate).ReplaceAll("###", logger.module).ToString()

	} else {

		logFile = logFile + MotadataString(LogFile).ReplaceAll("@@@", currentDate).ReplaceAll("###", "Plugin Engine").ToString()
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {

		panic(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if _, err = file.WriteString(message + NewLineSeparator); err != nil {

		panic(err)
	}
}
