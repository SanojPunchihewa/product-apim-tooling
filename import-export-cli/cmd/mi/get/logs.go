// /*
// *  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
// *
// *  WSO2 Inc. licenses this file to you under the Apache License,
// *  Version 2.0 (the "License"); you may not use this file except
// *  in compliance with the License.
// *  You may obtain a copy of the License at
// *
// *    http://www.apache.org/licenses/LICENSE-2.0
// *
// * Unless required by applicable law or agreed to in writing,
// * software distributed under the License is distributed on an
// * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// * KIND, either express or implied.  See the License for the
// * specific language governing permissions and limitations
// * under the License.
//  */

package get

// import (
// 	"github.com/spf13/cobra"
// 	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
// )

// var getApplicationCmdEnvironment string
// var getApplicationCmdFormat string
// var appName string

// // Show API command related usage info
// const getApplicationCmdLiteral = "compositeapp [app-name]"
// const getApplicationCmdShortDesc = "Get information about Composite Applications deployed in a Micro Integrator in an environment"

// const getApplicationCmdLongDesc = `Get information about the Composite App specified by command line argument [app-name] If not specified, list all the composite apps
// deployed in a Micro Integrator in the environment specified by the flag --environment, -e`

// var getApplicationCmdExamples = `To list all the composite apps
// ` + utils.ProjectName + ` ` + utils.MiCmdLiteral + ` ` + GetCmdLiteral + ` ` + getTrimmedCmdLiteral(getApplicationCmdLiteral) + ` -e dev
// To get details about a specific composite app
// ` + utils.ProjectName + ` ` + utils.MiCmdLiteral + ` ` + GetCmdLiteral + ` ` + getTrimmedCmdLiteral(getApplicationCmdLiteral) + ` FoodAPI -e dev
// NOTE: The flag (--environment (-e)) is mandatory`

// // apiShowCmd represents the show api command
// var getApplicationCmd = &cobra.Command{
// 	Use:     getApplicationCmdLiteral,
// 	Short:   getApplicationCmdShortDesc,
// 	Long:    getApplicationCmdLongDesc,
// 	Example: getApplicationCmdExamples,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// handleAPICmdArguments(args)
// 	},
// }

// func init() {
// 	GetCmd.AddCommand(getApplicationCmd)
// 	// apiShowCmd.SetHelpTemplate(showAPICmdLongDesc + utils.GetCmdUsage(programName, apiCmdLiteral,
// 	// 	showAPICmdLiteral, "[api-name]") + showAPICmdExamples + utils.GetCmdFlags(apiCmdLiteral))
// 	// apiShowCmd.Flags().StringVarP(&getApisCmdEnvironment, "mi-environment", "",
// 	// 	"", "Environment to be searched")
// 	getApplicationCmd.Flags().StringVarP(&getApplicationCmdEnvironment, "environment", "e",
// 		"", "Environment to be searched")
// 	getApplicationCmd.Flags().StringVarP(&getApplicationCmdFormat, "format", "", "", "Pretty-print apis "+
// 		"using Go Templates. Use \"{{ jsonPretty . }}\" to list all fields")
// 	getApplicationCmd.MarkFlagRequired("environment")
// }
