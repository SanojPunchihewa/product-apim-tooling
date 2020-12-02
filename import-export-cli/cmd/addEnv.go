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

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wso2/product-apim-tooling/import-export-cli/impl"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

var envToBeAdded string // Name of the environment to be added

var flagTokenEndpoint string        // token endpoint of the environment to be added
var flagPublisherEndpoint string    // Publisher endpoint of the environment to be added
var flagDevPortalEndpoint string    // DevPortal endpoint of the environment to be added
var flagRegistrationEndpoint string // registration endpoint of the environment to be added
var flagApiManagerEndpoint string   // api manager endpoint of the environment to be added
var flagAdminEndpoint string        // admin endpoint of the environment to be added
var flagMiManagementEndpoint string // mi management endpoint of the environment to be added

// AddEnv command related Info
const AddEnvCmdLiteral = "env [environment]"
const AddEnvCmdLiteralTrimmed = "env"
const addEnvCmdShortDesc = "Add Environment to Config file"
const addEnvCmdLongDesc = "Add new environment and its related endpoints to the config file"
const addEnvCmdExamples = utils.ProjectName + ` ` + AddCmdLiteral + ` ` + AddEnvCmdLiteralTrimmed + ` production \
--apim  https://localhost:9443 

` + utils.ProjectName + ` ` + AddCmdLiteral + ` ` + AddEnvCmdLiteralTrimmed + ` dev \
--mi https://localhost:9164

` + utils.ProjectName + ` ` + AddCmdLiteral + ` ` + AddEnvCmdLiteralTrimmed + ` test \
--registration https://idp.com:9443 \
--publisher https://apim.com:9443 \
--devportal  https://apps.com:9443 \
--admin  https://apim.com:9443 \
--token https://gw.com:8243/token
--mi https://localhost:9164

` + utils.ProjectName + ` ` + AddCmdLiteral + ` ` + AddEnvCmdLiteralTrimmed + ` dev \
--apim https://apim.com:9443 \
--registration https://idp.com:9443 \
--token https://gw.com:8243/token

You can either provide only the flag --apim , or all the other 4 flags (--registration --publisher --devportal --admin) without providing --apim flag.
If you are omitting any of --registration --publisher --devportal --admin flags, you need to specify --apim flag with the API Manager endpoint. In both of the
cases --token flag is optional and use it to specify the gateway token endpoint. This will be used for "apictl get-keys" operation.
To add a micro integrator instance to an environment you can use the --mi flag.`

// addEnvCmd represents the addEnv command
var addEnvCmd = &cobra.Command{
	Use:     AddEnvCmdLiteral,
	Short:   addEnvCmdShortDesc,
	Long:    addEnvCmdLongDesc,
	Example: addEnvCmdExamples,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envToBeAdded = args[0]

		utils.Logln(utils.LogPrefixInfo + AddCmdLiteral + " " + AddEnvCmdLiteralTrimmed + " called")
		executeAddEnvCmd(utils.MainConfigFilePath)
	},
}

func executeAddEnvCmd(mainConfigFilePath string) {
	envEndpoints := new(utils.EnvEndpoints)
	envEndpoints.ApiManagerEndpoint = flagApiManagerEndpoint
	envEndpoints.RegistrationEndpoint = flagRegistrationEndpoint

	envEndpoints.PublisherEndpoint = flagPublisherEndpoint
	envEndpoints.DevPortalEndpoint = flagDevPortalEndpoint
	envEndpoints.AdminEndpoint = flagAdminEndpoint
	envEndpoints.TokenEndpoint = flagTokenEndpoint
	envEndpoints.MiManagementEndpoint = flagMiManagementEndpoint
	err := impl.AddEnv(envToBeAdded, envEndpoints, mainConfigFilePath, AddEnvCmdLiteral)
	if err != nil {
		utils.HandleErrorAndExit("Error adding environment", err)
	}
}

// init using Cobra
func init() {
	AddCmd.AddCommand(addEnvCmd)

	addEnvCmd.Flags().StringVar(&flagApiManagerEndpoint, "apim", "", "API Manager endpoint for the environment")
	addEnvCmd.Flags().StringVar(&flagPublisherEndpoint, "publisher", "", "Publisher endpoint for the environment")
	addEnvCmd.Flags().StringVar(&flagDevPortalEndpoint, "devportal", "", "DevPortal endpoint for the environment")
	addEnvCmd.Flags().StringVar(&flagTokenEndpoint, "token", "", "Token endpoint for the environment")
	addEnvCmd.Flags().StringVar(&flagRegistrationEndpoint, "registration", "",
		"Registration endpoint for the environment")
	addEnvCmd.Flags().StringVar(&flagAdminEndpoint, "admin", "", "Admin endpoint for the environment")
	addEnvCmd.Flags().StringVar(&flagMiManagementEndpoint, "mi", "", "Micro Integrator Management endpoint for the environment")
	_ = addEnvCmd.MarkFlagRequired("environment")
}
