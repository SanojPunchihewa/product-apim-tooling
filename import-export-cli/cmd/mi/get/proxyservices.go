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

var getProxyServiceCmdEnvironment string
var getProxyServiceCmdFormat string

const artifactProxyServices = "proxy services"
const getProxyServiceCmdLiteral = "proxyservices [proxy-name]"

var getProxyServiceCmd = &cobra.Command{
	Use:     getProxyServiceCmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactProxyServices),
	Long:    generateGetCmdLongDescForArtifact(artifactProxyServices, "proxy-name"),
	Example: generateGetCmdExamplesForArtifact(artifactProxyServices, getTrimmedCmdLiteral(getProxyServiceCmdLiteral), "SampleProxy"),
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetProxyServiceCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getProxyServiceCmd)
	getProxyServiceCmd.Flags().StringVarP(&getProxyServiceCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getProxyServiceCmd.Flags().StringVarP(&getProxyServiceCmdFormat, "format", "", "", generateFormatFlagUsage(artifactProxyServices))
	getProxyServiceCmd.MarkFlagRequired("environment")
}

func handleGetProxyServiceCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get proxyservices called")
	cred, err := credentials.GetMICredentials(getProxyServiceCmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	if len(args) == 1 {
		var proxyServiceName = args[0]
		executeShowProxyService(cred, proxyServiceName)
	} else {
		executeListProxyServices(cred)
	}
}

func executeListProxyServices(cred credentials.MiCredential) {

	// localEntryList, err := impl.GetLocalEntryList(cred, getLocalEntryCmdEnvironment)
	// if err == nil {
	// 	impl.PrintLocalEntryList(localEntryList, getLocalEntryCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting List of localentries", err)
	// }
}

func executeShowProxyService(cred credentials.MiCredential, epName string) {
	// localEntry, err := impl.GetLocalEntry(cred, getLocalEntryCmdEnvironment, epName)
	// if err == nil {
	// 	impl.PrintLocalEntryDetails(localEntry, getLocalEntryCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting Information of the localentry", err)
	// }
}
