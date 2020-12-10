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

var getDataServiceCmdEnvironment string
var getDataServiceCmdFormat string

const artifactDataServices = "data services"
const getDataServiceCmdLiteral = "dataservices [dataservice-name]"

var getDataServiceCmd = &cobra.Command{
	Use:     getDataServiceCmdLiteral,
	Short:   generateGetCmdShortDescForArtifact(artifactDataServices),
	Long:    generateGetCmdLongDescForArtifact(artifactDataServices, "dataservice-name"),
	Example: generateGetCmdExamplesForArtifact(artifactDataServices, getTrimmedCmdLiteral(getDataServiceCmdLiteral), "SampleDataService"),
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleGetDataServiceCmdArguments(args)
	},
}

func init() {
	GetCmd.AddCommand(getDataServiceCmd)
	getDataServiceCmd.Flags().StringVarP(&getDataServiceCmdEnvironment, "environment", "e",
		"", "Environment to be searched")
	getDataServiceCmd.Flags().StringVarP(&getDataServiceCmdFormat, "format", "", "", generateFormatFlagUsage(artifactDataServices))
	getDataServiceCmd.MarkFlagRequired("environment")
}

func handleGetDataServiceCmdArguments(args []string) {
	utils.Logln(utils.LogPrefixInfo + "get dataservices called")
	_, err := credentials.GetMICredentials(getDataServiceCmdEnvironment)
	if err != nil {
		utils.HandleErrorAndExit("Error getting credentials", err)
	}
	if len(args) == 1 {
		var dataServiceName = args[0]
		executeShowDataService(dataServiceName)
	} else {
		executeListDataServices()
	}
}

func executeListDataServices() {

	dataServiceList, err := impl.GetDataServiceList(getDataServiceCmdEnvironment)
	if err == nil {
		impl.PrintDataServiceList(dataServiceList, getDataServiceCmdFormat)
	} else {
		utils.Logln(utils.LogPrefixError+"Getting List of DataServices", err)
	}
}

func executeShowDataService(dataserviceName string) {
	dataservice, err := impl.GetDataService(getDataServiceCmdEnvironment, dataserviceName)
	if err == nil {
		impl.PrintDataServiceDetails(dataservice, getDataServiceCmdFormat)
	} else {
		utils.Logln(utils.LogPrefixError+"Getting Information of the DataService", err)
	}
}
