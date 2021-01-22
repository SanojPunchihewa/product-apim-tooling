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

const validDataServiceName = "RESTDataService"
const invalidDataServiceName = "abcDataService"
const dataServiceCmd = "data-services"

func TestGetDataServices(t *testing.T) {
	ValidateDataServicesList(t, config)
}

func TestGetDataServiceByName(t *testing.T) {
	ValidateDataService(t, config, validDataServiceName)
}

func TestGetNonExistingDataServiceByName(t *testing.T) {
	response, _ := getArtifact(t, dataServiceCmd, invalidDataServiceName, config)
	base.Log(response)
	assert.Contains(t, response, "[ERROR]: Getting Information of data services [ "+invalidDataServiceName+" ]  404 Not Found")
}

func TestGetDataServicesWithoutSettingUpEnv(t *testing.T) {
	execGetCommandWithoutSettingEnv(t, dataServiceCmd)
}

func TestGetDataServicesWithoutLogin(t *testing.T) {
	execGetCommandWithoutLogin(t, dataServiceCmd, config)
}

func TestGetDataServicesWithoutEnvFlag(t *testing.T) {
	execGetCommandWithoutEnvFlag(t, dataServiceCmd, config)
}

func TestGetDataServicesWithInvalidArgs(t *testing.T) {
	execGetCommandWithInvalidArgCount(t, config, 1, 2, false, dataServiceCmd, validDataServiceName, invalidDataServiceName)
}

func ValidateDataServicesList(t *testing.T, config *MiConfig) {
	t.Helper()
	output, _ := listArtifacts(t, dataServiceCmd, config)
	artifactList := config.MIClient.GetArtifactList(utils.MiManagementDataServiceResource, &artifactutils.DataServicesList{})
	ValidateDataServiceListEqual(t, output, (artifactList.(*artifactutils.DataServicesList)))
}

func ValidateDataServiceListEqual(t *testing.T, dataServicesListFromCtl string, dataServicesList *artifactutils.DataServicesList) {
	unmatchedCount := dataServicesList.Count
	for _, dataService := range dataServicesList.List {
		assert.Truef(t, strings.Contains(dataServicesListFromCtl, dataService.ServiceName), "dataServicesListFromCtl: "+dataServicesListFromCtl+
			" , does not contain dataService.ServiceName: "+dataService.ServiceName)
		unmatchedCount--
	}
	assert.Equal(t, 0, int(unmatchedCount), "Data Service lists are not equal")
}

func ValidateDataService(t *testing.T, config *MiConfig, dataServiceName string) {
	t.Helper()
	output, _ := getArtifact(t, dataServiceCmd, dataServiceName, config)
	artifactList := config.MIClient.GetArtifact(utils.MiManagementDataServiceResource, "dataServiceName", dataServiceName, &artifactutils.DataServiceInfo{})
	ValidateDataServiceEqual(t, output, (artifactList.(*artifactutils.DataServiceInfo)))
}

func ValidateDataServiceEqual(t *testing.T, dataServicesListFromCtl string, dataService *artifactutils.DataServiceInfo) {
	assert.Contains(t, dataServicesListFromCtl, dataService.ServiceName)
	assert.Contains(t, dataServicesListFromCtl, dataService.ServiceGroupName)
	assert.Contains(t, dataServicesListFromCtl, dataService.Wsdl11)
	assert.Contains(t, dataServicesListFromCtl, dataService.Wsdl20)
	for _, query := range dataService.Queries {
		assert.Contains(t, dataServicesListFromCtl, query.Id)
	}
}
