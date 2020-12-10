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
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

var getTemplateCmdEnvironment string
var getTemplateCmdFormat string

const artifactTemplates = "templates"
const getTemplateCmdLiteral = "templates [template-type] [template-name]"

var getTemplateCmd = &cobra.Command{
	Use:     getTemplateCmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactTemplates),
	Long:    generateGetCmdLongDescForArtifact(artifactTemplates, ""),                                              //
	Example: generateGetCmdExamplesForArtifact(artifactTemplates, getTrimmedCmdLiteral(getTemplateCmdLiteral), ""), //
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetTemplateCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getTemplateCmd)
	getTemplateCmd.Flags().StringVarP(&getTemplateCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getTemplateCmd.Flags().StringVarP(&getTemplateCmdFormat, "format", "", "", generateFormatFlagUsage(artifactTemplates))
	getTemplateCmd.MarkFlagRequired("environment")
}

func handleGetTemplateCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get templates called")
	// cred, err := credentials.GetMICredentials(getTemplateCmdEnvironment)
	// if err != nil {
	// 	utils.HandleErrorAndExit("Error getting credentials", err)
	// }
	//
}

// func executeListLocalEntrys(cred credentials.MiCredential) {

// 	// localEntryList, err := impl.GetLocalEntryList(cred, getLocalEntryCmdEnvironment)
// 	// if err == nil {
// 	// 	impl.PrintLocalEntryList(localEntryList, getLocalEntryCmdFormat)
// 	// } else {
// 	// 	utils.Logln(utils.LogPrefixError+"Getting List of localentries", err)
// 	// }
// }

// func executeShowLocalEntry(cred credentials.MiCredential, epName string) {
// 	// localEntry, err := impl.GetLocalEntry(cred, getLocalEntryCmdEnvironment, epName)
// 	// if err == nil {
// 	// 	impl.PrintLocalEntryDetails(localEntry, getLocalEntryCmdFormat)
// 	// } else {
// 	// 	utils.Logln(utils.LogPrefixError+"Getting Information of the localentry", err)
// 	// }
// }
