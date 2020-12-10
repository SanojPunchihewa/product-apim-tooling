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

var getIntegrationAPICmdEnvironment string
var getIntegrationAPICmdFormat string

const artifactAPIs = "integration apis"
const getIntegrationAPICmdLiteral = "integrationapis [api-name]"

var getIntegrationAPICmd = &cobra.Command{
	Use:     getIntegrationAPICmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactAPIs),
	Long:    generateGetCmdLongDescForArtifact(artifactAPIs, "api-name"),
	Example: generateGetCmdExamplesForArtifact(artifactAPIs, getTrimmedCmdLiteral(getIntegrationAPICmdLiteral), "SampleIntegrationAPI"),
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetIntegrationAPICmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getIntegrationAPICmd)
	getIntegrationAPICmd.Flags().StringVarP(&getIntegrationAPICmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getIntegrationAPICmd.Flags().StringVarP(&getIntegrationAPICmdFormat, "format", "", "", generateFormatFlagUsage(artifactAPIs))
	getIntegrationAPICmd.MarkFlagRequired("environment")
}

func handleGetIntegrationAPICmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get integrationapis called")
	_, err := credentials.GetMICredentials(getIntegrationAPICmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	if len(args) == 1 {
		var IntegrationAPIName = args[0]
		executeShowIntegrationAPI(IntegrationAPIName)
	} else {
		executeListIntegrationAPIs()
	}
}

func executeListIntegrationAPIs() {

	apiList, err := impl.GetIntegrationAPIList(getIntegrationAPICmdEnvironment)
	if err == nil {
		impl.PrintIntegrationAPIList(apiList, getIntegrationAPICmdFormat)
	} else {
		utils.Logln(utils.LogPrefixError+"Getting List of integration apis", err)
	}
}

func executeShowIntegrationAPI(epName string) {
	integrationAPI, err := impl.GetIntegrationAPI(getIntegrationAPICmdEnvironment, epName)
	if err == nil {
		impl.PrintIntegrationAPIDetails(integrationAPI, getIntegrationAPICmdFormat)
	} else {
		utils.Logln(utils.LogPrefixError+"Getting Information of the integration api", err)
	}
}
