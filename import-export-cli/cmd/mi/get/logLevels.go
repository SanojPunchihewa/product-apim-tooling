/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package get

import (
	"github.com/spf13/cobra"
	"github.com/wso2/product-apim-tooling/import-export-cli/credentials"
	impl "github.com/wso2/product-apim-tooling/import-export-cli/impl/mi"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

var getLogLevelCmdEnvironment string
var getLogLevelCmdFormat string

const getLogLevelCmdLiteral = "log-levels [logger-name]"
const getLogLevelCmdShortDesc = "Get information about a Logger configured in a Micro Integrator"

const getLogLevelCmdLongDesc = "Get information about the Logger specified by command line argument [logger-name]\nconfigured in a Micro Integrator in the environment specified by the flag --environment, -e"

var getLogLevelCmdExamples = "To get details about a specific logger\n" +
	utils.ProjectName + " " + utils.MiCmdLiteral + " " + GetCmdLiteral + " " + getTrimmedCmdLiteral(getLogLevelCmdLiteral) + " org-apache-coyote -e dev\n" +
	"NOTE: The flag (--environment (-e)) is mandatory"

var getLogLevelCmd = &cobra.Command{
	Use:     getLogLevelCmdLiteral,
	Short:   getLogLevelCmdShortDesc,
	Long:    getLogLevelCmdLongDesc,
	Example: getLogLevelCmdExamples,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetLogLevelCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getLogLevelCmd)
	getLogLevelCmd.Flags().StringVarP(&getLogLevelCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getLogLevelCmd.Flags().StringVarP(&getLogLevelCmdFormat, "format", "", "", generateFormatFlagUsage(getTrimmedCmdLiteral(getLogLevelCmdLiteral)))
	getLogLevelCmd.MarkFlagRequired("environment")
}

func handleGetLogLevelCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get LogLevels called")
	_, err := credentials.GetMICredentials(getLogLevelCmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	var loggerName = args[0]
	executeShowLogLevel(loggerName)
}

func executeShowLogLevel(loggerName string) {

	LogLevelList, err := impl.GetLoggerInfo(getLogLevelCmdEnvironment, loggerName)
	if err == nil {
		impl.PrintLoggerInfo(LogLevelList, getLogLevelCmdFormat)
	} else {
		printErrorForArtifact("logger", loggerName, err)
	}
}
