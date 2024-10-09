/*
 * Copyright (c) Motadata 2024.  All rights reserved.
 */

package utils

import (
	"os"
)

const (
	SystemLogLevel = "system.log.level"

	NotAvailable = -1

	NewLineRegexPattern = "\r\n"

	Yes = "yes"

	No = "no"

	Password = "password"

	ConfigDirectory = "config"

	Errors = "errors"

	Result = "result"

	Message = "message"

	ObjectContext = "object.context"

	Status = "status"

	StatusSucceed = "succeed"

	StatusFail = "fail"

	BlankString = ""

	ColonSeparator = ":"

	EqualSeparator = "="

	DotSeparator = "."

	SpaceSeparator = " "

	CommaSeparator = ","

	ForwardSlashSeparator = "/"

	NewLineSeparator = "\n"

	PipeSeparator = '|'

	TabSeparator = "\t"

	InstanceSeparator = "_"

	TimeFormat = "2006-01-02 15:04:05"

	DiscoveryCredentialProfiles = "discovery.credential.profiles"

	Id = "id"

	// debug/info/error message related const

	InfoMessageConnectionDestroyed = "disconnecting from %s : %s"

	ErrorMessageConnectionTimeout = "Timeout Error: The %s request for %s:%d has timed out."

	ErrorMessageInvalidPort = "Connection Error: Invalid Port %d"

	ErrorMessageConnectionFailed = "Connection Error: Failed to Establish %s Connection on %s:%d"

	ErrorMessageInvalidCredentials = "Authentication Error: Invalid Credentials for %s:%d"

	DebugMessageRequest = "Request of : %s"

	DebugMessageResult = "Result %s : %s"

	//Log related Constants

	PathSeparator = string(os.PathSeparator)

	LogLevelTrace = 0

	LogLevelDebug = 1

	LogLevelInfo = 2

	LogDirectory = "logs"

	LogFileDateFormat = "02-January-2006"

	LogFileTimeFormat = "03:04:05.000000 PM"

	LogFile = "@@@-###.log"

	//Error Code

	ErrorCodeInternalError = "MD031"

	ErrorCodeTimeout = "MD004"

	ErrorCodeInvalidPort = "MD002"

	ErrorCodeInvalidCredentials = "MD003"

	ErrorCodeConnectionFailed = "MD047"

	ErrorCodeInvalidPublicOrPrivateSSHKey = "MD103"

	ErrorCodeConnectionReset = "MD058"

	ErrorCodeCommandExecutionFailed = "MD059"

	// -----> Error/Info messages <-------

	ErrorMessageFailReadingOutput                         = "Got Error while reading output, error : %v for host %s"
	ErrorMessageTimeOutWhileReadingOutput                 = "Timed out while reading output for command : %s , pattern : %s for host %s"
	ErrorMessageReadingOutputTimeOut                      = "Timed out while reading output for command : %s , pattern : %s"
	InfoMessageExecutingCommand                           = "Executing command : %s for host %s"
	InfoMessageExecutedCommand                            = "Command : %s executed for host %s with execution time is : %d"
	InfoMessageTryingTemplate                             = "Trying template : %s for the host : %s"
	InfoMessageQualifiedTemplate                          = "Qualified template is : %s for the host : %s"
	InfoMessageTryingCredentialProfile                    = "Trying credential profile : %s for the host : %s"
	InfoMessageQualifiedCredentialProfile                 = "Qualified credential profile is : %s for the host : %s"
	InfoMessageBackupExecutionResult                      = "%s result for the template : %s with host : %s is %v"
	InfoMessageExecutedCommandMessage                     = "Successfully executed command : %s for the host : %s"
	InfoMessageBackupEndPrompt                            = "Upgrade end prompt for the host : %s is : %s, prompt index : %d"
	InfoCommandNotFound                                   = "command not found for the %s , template : %s,  host : %s"
	ErrorMessageSessionCreationIssue                      = "Error %v occurred for the host : %s while creating ssh session"
	ErrorMessageWhileInitSSHSession                       = "Error while initialising session for the host : %s:%d"
	ErrorMessageWhileRequestingPTY                        = "Error %v occurred while requesting Pty for the host : %s"
	ErrorMessageWhileRequestingPTYWithHost                = "Error while requesting Pty for the host : %s:%d"
	ErrorMessageWhileInvokingSSHShell                     = "Error %v occurred while invoking ssh Shell for the host : %s"
	ErrorMessageWhileInvokingSSHShellWithHost             = "Error while invoking ssh shell for the host : %s:%d"
	ErrorMessageWhileSendingEnableCommand                 = "Error %v occurred while sending enable command for the host : %s"
	ErrorMessageWhileSendingEnableCommandWithHost         = "Error while sending enable command for the host : %s:%d"
	ErrorMessageWhileSendingEnableUserCommand             = "Error %v occurred while sending enable user command for the host : %s"
	ErrorMessageWhileSendingEnableUserCommandWithHost     = "Error while sending enable user command for the host : %s:%d"
	ErrorMessageWhileSendingEnablePasswordCommand         = "Error %v occurred while sending enable password command for the host : %s"
	ErrorMessageWhileSendingEnablePasswordCommandWithHost = "Error while sending enable password command for the host : %s:%d"
	ErrorMessageFailToEnterEnableMode                     = "Failed to enter enable mode for the host : %s"
	ErrorMessageFailToEnterEnableModeWithHost             = "Failed to enter enable mode for the host : %s:%d"
	ErrorMessageEnablePasswordNotProvided                 = "Enable Password not provided for the host : %s"
	ErrorMessageCommandExecutionFail                      = "Execution failed for the request : %s for the host : %s"
	ErrorMessagePromptCommandExecutionFail                = "Prompt command execution failed for the request : %s for the host : %s"
	ErrorMessageConfigOutputError                         = "Error found in output of the command : %s for host %s, output : %s"
	ErrorMessageConfigOutputPromptCommandError            = "Error found in output of the prompt command : %s for host %s, output : %s"
	ErrorMessageConfigOutputErrorCheck                    = "Error found in output of the command"
	ErrorMessageConfigTemplateNotAvailable                = "No config template provided for the host : %s"
	ErrorMessageConfigOperationCommandsNotAvailable       = "No config operation commands provided for the host : %s, operation : %s"
	ErrorMessageOperationCommandNotAvailable              = "No operation group command added for the request : %s, template : %s, protocol : %s, host : %s"
	InfoMessageExtractingBackupResult                     = "Extracting backup result for host : %s"
	InfoMessageBackupFileName                             = "Upgrade file name %s"
	ErrorMessageProcessTimeout                            = "Process was killed due to time out"

	ErrorCodeFailExecutingCommand                      = "MD128"
	ErrorCodeWhileInitSSHSession                       = "MD129"
	ErrorCodeWhileSendingEnableCommandWithHost         = "MD130"
	ErrorCodeWhileSendingEnableUserCommandWithHost     = "MD131"
	ErrorCodeWhileSendingEnablePasswordCommandWithHost = "MD132"
	ErrorCodeFailToEnterEnableModeWithHost             = "MD133"
	ErrorCodeEnablePasswordNotProvided                 = "MD134"
	ErrorCodeConfigTemplateNotAvailable                = "MD135"
	ErrorCodeConfigOperationCommandNotAvailable        = "MD136"
	ErrorCodeWhileRequestingPTYWithHost                = "MD117"
	ErrorCodeWhileInvokingSSHShellWithHost             = "MD118"
)

var (
	CurrentDir, _ = os.Getwd()

	SensitiveFields = []string{"$$$password$$$", "$$$snmp.private.password$$$", "$$$snmp.community$$$",
		"$$$snmp.authentication.password$$$", "$$$snmp.community$$$",
		"$$$snmp.authentication.password$$$", "$$$snmp.private.password$$$",
		"$$$user.password$$$", "$$$secret.key$$$", "$$$access.id$$$",
		"$$$client.id$$$", "$$$ssh.key$$$", "$$$passphrase$$$", Password,
		"snmp.community", "snmp.authentication.password", "snmp.private.password",
		ObjectContext, DiscoveryCredentialProfiles,
		"cloud.secret.key", "cloud.client.id", "cloud.access.id", "cloud.tenant.id",
		"ssh.key", "passphrase", "SNMP Community", "Authentication Password",
		"Private Password", "metric.plugin.variables", "mail.server.password",
		"ldap.server.password", "discovery.context", "credential.profile.context",
		"user.password", "enable.password", "storage.profile.context"}
)
