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

package integration

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wso2/product-apim-tooling/import-export-cli/integration/base"
	"github.com/wso2/product-apim-tooling/import-export-cli/mi/utils/artifactutils"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

const validAPIName = "HealthcareAPI"
const invalidAPIName = "abcAPI"
const apisCmd = "apis"

func TestGetAPIs(t *testing.T) {
	ValidateAPIsList(t, config)
}

func TestGetAPIByName(t *testing.T) {
	ValidateAPI(t, config, validAPIName)
}

func TestGetNonExistingAPIByName(t *testing.T) {
	response, _ := getArtifact(t, apisCmd, invalidAPIName, config)
	base.Log(response)
	assert.Contains(t, response, "[ERROR]: Getting Information of "+apisCmd+" [ "+invalidAPIName+" ]  404 Not Found")
}

func TestGetAPIsWithoutSettingUpEnv(t *testing.T) {
	execGetCommandWithoutSettingEnv(t, apisCmd)
}

func TestGetAPIsWithoutLogin(t *testing.T) {
	execGetCommandWithoutLogin(t, apisCmd, config)
}

func TestGetAPIsWithoutEnvFlag(t *testing.T) {
	execGetCommandWithoutEnvFlag(t, apisCmd, config)
}

func TestGetAPIsWithInvalidArgs(t *testing.T) {
	execGetCommandWithInvalidArgCount(t, config, 1, 2, false, apisCmd, validAPIName, invalidAPIName)
}

func ValidateAPIsList(t *testing.T, config *MiConfig) {
	t.Helper()
	output, _ := listArtifacts(t, apisCmd, config)
	artifactList := config.MIClient.GetArtifactList(utils.MiManagementAPIResource, &artifactutils.IntegrationAPIList{})
	ValidateAPIListEqual(t, output, (artifactList.(*artifactutils.IntegrationAPIList)))
}

func ValidateAPIListEqual(t *testing.T, apisListFromCtl string, apisList *artifactutils.IntegrationAPIList) {
	unmatchedCount := apisList.Count
	for _, api := range apisList.Apis {
		assert.Truef(t, strings.Contains(apisListFromCtl, api.Name), "apisListFromCtl: "+apisListFromCtl+
			" , does not contain api.Name: "+api.Name)
		unmatchedCount--
	}
	assert.Equal(t, 0, int(unmatchedCount), "API lists are not equal")
}

func ValidateAPI(t *testing.T, config *MiConfig, apiName string) {
	t.Helper()
	output, _ := getArtifact(t, apisCmd, apiName, config)
	artifactList := config.MIClient.GetArtifact(utils.MiManagementAPIResource, "apiName", apiName, &artifactutils.IntegrationAPI{})
	ValidateAPIEqual(t, output, (artifactList.(*artifactutils.IntegrationAPI)))
}

func ValidateAPIEqual(t *testing.T, apisListFromCtl string, api *artifactutils.IntegrationAPI) {
	assert.Contains(t, apisListFromCtl, api.Name)
	assert.Contains(t, apisListFromCtl, api.Stats)
	assert.Contains(t, apisListFromCtl, api.Url)
	assert.Contains(t, apisListFromCtl, api.Version)
}
