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
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

var getLocalEntryCmdEnvironment string
var getLocalEntryCmdFormat string

const artifactLocalEntries = "local entries"
const getLocalEntryCmdLiteral = "localentries [localentry-name]"

var getLocalEntryCmd = &cobra.Command{
	Use:     getLocalEntryCmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactLocalEntries),
	Long:    generateGetCmdLongDescForArtifact(artifactLocalEntries, "localentry-name"),
	Example: generateGetCmdExamplesForArtifact(artifactLocalEntries, getTrimmedCmdLiteral(getLocalEntryCmdLiteral), "SampleLocalEntry"),
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetLocalEntryCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getLocalEntryCmd)
	getLocalEntryCmd.Flags().StringVarP(&getLocalEntryCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getLocalEntryCmd.Flags().StringVarP(&getLocalEntryCmdFormat, "format", "", "", generateFormatFlagUsage(artifactLocalEntries))
	getLocalEntryCmd.MarkFlagRequired("environment")
}

func handleGetLocalEntryCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get localentries called")
	cred, err := credentials.GetMICredentials(getLocalEntryCmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	if len(args) == 1 {
		var LocalEntryName = args[0]
		executeShowLocalEntry(cred, LocalEntryName)
	} else {
		executeListLocalEntrys(cred)
	}
}

func executeListLocalEntrys(cred credentials.MiCredential) {

	// localEntryList, err := impl.GetLocalEntryList(cred, getLocalEntryCmdEnvironment)
	// if err == nil {
	// 	impl.PrintLocalEntryList(localEntryList, getLocalEntryCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting List of localentries", err)
	// }
}

func executeShowLocalEntry(cred credentials.MiCredential, epName string) {
	// localEntry, err := impl.GetLocalEntry(cred, getLocalEntryCmdEnvironment, epName)
	// if err == nil {
	// 	impl.PrintLocalEntryDetails(localEntry, getLocalEntryCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting Information of the localentry", err)
	// }
}
