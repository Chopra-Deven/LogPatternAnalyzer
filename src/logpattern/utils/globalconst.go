/*
 * Copyright (c) Motadata 2024.  All rights reserved.
 */

package utils

import (
	"os"
)

const (
	SystemLogLevel = "system.log.level"

	NewLineRegexPattern = "\r\n"

	ConfigDirectory = "config"

	Errors = "errors"

	Message = "message"

	Status = "status"

	StatusSucceed = "succeed"

	StatusFail = "fail"

	BlankString = ""

	SpaceSeparator = " "

	NewLineSeparator = "\n"

	TimeFormat = "2006-01-02 15:04:05"

	PathSeparator = string(os.PathSeparator)

	LogLevelTrace = 0

	LogLevelDebug = 1

	LogLevelInfo = 2

	LogDirectory = "logs"

	LogFileDateFormat = "02-January-2006"

	LogFileTimeFormat = "03:04:05.000000 PM"

	LogFile = "@@@-###.log"
)

var (
	CurrentDir, _ = os.Getwd()
)
