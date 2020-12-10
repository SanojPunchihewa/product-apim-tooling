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

var getInboundEndpointCmdEnvironment string
var getInboundEndpointCmdFormat string

const artifactInboundEndpoints = "inbound endpoints"
const getInboundEndpointCmdLiteral = "inboundendpoints [inbound-name]"

var getInboundEndpointCmd = &cobra.Command{
	Use:     getInboundEndpointCmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactInboundEndpoints),
	Long:    generateGetCmdLongDescForArtifact(artifactInboundEndpoints, "inbound-name"),
	Example: generateGetCmdExamplesForArtifact(artifactInboundEndpoints, getTrimmedCmdLiteral(getInboundEndpointCmdLiteral), "SampleInboundEndpoint"),
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetInboundEndpointCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getInboundEndpointCmd)
	getInboundEndpointCmd.Flags().StringVarP(&getInboundEndpointCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getInboundEndpointCmd.Flags().StringVarP(&getInboundEndpointCmdFormat, "format", "", "", generateFormatFlagUsage(artifactInboundEndpoints))
	getInboundEndpointCmd.MarkFlagRequired("environment")
}

func handleGetInboundEndpointCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get inboundendpoints called")
	_, err := credentials.GetMICredentials(getInboundEndpointCmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	if len(args) == 1 {
		var inboundEndpointName = args[0]
		executeShowInboundEndpoint(inboundEndpointName)
	} else {
		executeListInboundEndpoints()
	}
}

func executeListInboundEndpoints() {

	// epList, err := impl.GetInboundEndpointList(cred, getInboundEndpointCmdEnvironment)
	// if err == nil {
	// 	impl.PrintInboundEndpointList(epList, getInboundEndpointCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting List of inbound endpoints", err)
	// }
}

func executeShowInboundEndpoint(epName string) {
	// inboundEndpoint, err := impl.GetInboundEndpoint(cred, getInboundEndpointCmdEnvironment, epName)
	// if err == nil {
	// 	impl.PrintInboundEndpointDetails(inboundEndpoint, getInboundEndpointCmdFormat)
	// } else {
	// 	utils.Logln(utils.LogPrefixError+"Getting Information of the inbound endpoint", err)
	// }
}
