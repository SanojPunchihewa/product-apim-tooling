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

var getMessageStoreCmdEnvironment string
var getMessageStoreCmdFormat string

const artifactMessageStores = "message stores"
const getMessageStoreCmdLiteral = "messagestores [messagestore-name]"

var getMessageStoreCmd = &cobra.Command{
	Use:     getMessageStoreCmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactMessageStores),
	Long:    generateGetCmdLongDescForArtifact(artifactMessageStores, "messagestore-name"),
	Example: generateGetCmdExamplesForArtifact(artifactMessageStores, getTrimmedCmdLiteral(getMessageStoreCmdLiteral), "TestMessageStore"),
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetMessageStoreCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getMessageStoreCmd)
	getMessageStoreCmd.Flags().StringVarP(&getMessageStoreCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getMessageStoreCmd.Flags().StringVarP(&getMessageStoreCmdFormat, "format", "", "", generateFormatFlagUsage(artifactMessageStores))
	getMessageStoreCmd.MarkFlagRequired("environment")
}

func handleGetMessageStoreCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get messagestores called")
	cred, err := credentials.GetMICredentials(getMessageStoreCmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	if len(args) == 1 {
		var messageStoreName = args[0]
		executeShowMessageStore(cred, messageStoreName)
	} else {
		executeListMessageStores(cred)
	}
}

func executeListMessageStores(cred credentials.MiCredential) {

	// localEntryList, err := impl.GetLocalEntryList(cred, getLocalEntryCmdEnvironment)
	// if err == nil {
	// 	impl.PrintLocalEntryList(localEntryList, getLocalEntryCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting List of localentries", err)
	// }
}

func executeShowMessageStore(cred credentials.MiCredential, epName string) {
	// localEntry, err := impl.GetLocalEntry(cred, getLocalEntryCmdEnvironment, epName)
	// if err == nil {
	// 	impl.PrintLocalEntryDetails(localEntry, getLocalEntryCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting Information of the localentry", err)
	// }
}
