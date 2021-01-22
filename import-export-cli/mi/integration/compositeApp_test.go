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

const validCAppName = "HealthCareCompositeExporter"
const invalidCAppName = "abcCApp"
const cAppCmd = "composite-apps"

func TestGetCApps(t *testing.T) {
	ValidateCAppsList(t, config)
}

func TestGetCAppByName(t *testing.T) {
	ValidateCApp(t, config, validCAppName)
}

func TestGetNonExistingCAppByName(t *testing.T) {
	response, _ := getArtifact(t, cAppCmd, invalidCAppName, config)
	base.Log(response)
	assert.Contains(t, response, "[ERROR]: Getting Information of composite apps [ "+invalidCAppName+" ]  404 Not Found")
}

func TestGetCAppsWithoutSettingUpEnv(t *testing.T) {
	execGetCommandWithoutSettingEnv(t, cAppCmd)
}

func TestGetCAppsWithoutLogin(t *testing.T) {
	execGetCommandWithoutLogin(t, cAppCmd, config)
}

func TestGetCAppsWithoutEnvFlag(t *testing.T) {
	execGetCommandWithoutEnvFlag(t, cAppCmd, config)
}

func TestGetCAppsWithInvalidArgs(t *testing.T) {
	execGetCommandWithInvalidArgCount(t, config, 1, 2, false, cAppCmd, validCAppName, invalidCAppName)
}

func ValidateCAppsList(t *testing.T, config *MiConfig) {
	t.Helper()
	output, _ := listArtifacts(t, cAppCmd, config)
	artifactList := config.MIClient.GetArtifactList(utils.MiManagementCarbonAppResource, &artifactutils.CompositeAppList{})
	ValidateCAppListEqual(t, output, (artifactList.(*artifactutils.CompositeAppList)))
}

func ValidateCAppListEqual(t *testing.T, cAppsListFromCtl string, cAppsList *artifactutils.CompositeAppList) {
	unmatchedCount := cAppsList.Count
	for _, cApp := range cAppsList.CompositeApps {
		assert.Truef(t, strings.Contains(cAppsListFromCtl, cApp.Name), "cAppsListFromCtl: "+cAppsListFromCtl+
			" , does not contain cApp.Name: "+cApp.Name)
		unmatchedCount--
	}
	assert.Equal(t, 0, int(unmatchedCount), "CApp lists are not equal")
}

func ValidateCApp(t *testing.T, config *MiConfig, cAppName string) {
	t.Helper()
	output, _ := getArtifact(t, cAppCmd, cAppName, config)
	artifactList := config.MIClient.GetArtifact(utils.MiManagementCarbonAppResource, "carbonAppName", cAppName, &artifactutils.CompositeApp{})
	ValidateCAppEqual(t, output, (artifactList.(*artifactutils.CompositeApp)))
}

func ValidateCAppEqual(t *testing.T, CAppsListFromCtl string, cApp *artifactutils.CompositeApp) {
	assert.Contains(t, CAppsListFromCtl, cApp.Name)
	assert.Contains(t, CAppsListFromCtl, cApp.Version)
	for _, artifact := range cApp.Artifacts {
		assert.Contains(t, CAppsListFromCtl, artifact.Name)
		assert.Contains(t, CAppsListFromCtl, artifact.Type)
	}
}
