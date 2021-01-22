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

const validEndpointName = "GrandOakEndpoint"
const invalidEndpointName = "abcEndpoint"
const endpointCmd = "endpoints"

func TestGetEndpoints(t *testing.T) {
	ValidateEndpointList(t, config)
}

func TestGetEndpointByName(t *testing.T) {
	ValidateEndpoint(t, config, validEndpointName)
}

func TestGetNonExistingEndpointByName(t *testing.T) {
	response, _ := getArtifact(t, endpointCmd, invalidEndpointName, config)
	base.Log(response)
	assert.Contains(t, response, "[ERROR]: Getting Information of endpoints [ "+invalidEndpointName+" ]  404 Not Found")
}

func TestGetEndpointsWithoutSettingUpEnv(t *testing.T) {
	execGetCommandWithoutSettingEnv(t, endpointCmd)
}

func TestGetEndpointsWithoutLogin(t *testing.T) {
	execGetCommandWithoutLogin(t, endpointCmd, config)
}

func TestGetEndpointsWithoutEnvFlag(t *testing.T) {
	execGetCommandWithoutEnvFlag(t, endpointCmd, config)
}

func TestGetEndpointsWithInvalidArgs(t *testing.T) {
	execGetCommandWithInvalidArgCount(t, config, 1, 2, false, endpointCmd, validEndpointName, invalidEndpointName)
}

func ValidateEndpointList(t *testing.T, config *MiConfig) {
	t.Helper()
	output, _ := listArtifacts(t, endpointCmd, config)
	artifactList := config.MIClient.GetArtifactList(utils.MiManagementEndpointResource, &artifactutils.EndpointList{})
	ValidateEndpointListEqual(t, output, (artifactList.(*artifactutils.EndpointList)))
}

func ValidateEndpointListEqual(t *testing.T, endpointsListFromCtl string, endpointList *artifactutils.EndpointList) {
	unmatchedCount := endpointList.Count
	for _, endpoint := range endpointList.Endpoints {
		assert.Truef(t, strings.Contains(endpointsListFromCtl, endpoint.Name), "endpointsListFromCtl: "+endpointsListFromCtl+
			" , does not contain endpoint.Name: "+endpoint.Name)
		unmatchedCount--
	}
	assert.Equal(t, 0, int(unmatchedCount), "Endpoint lists are not equal")
}

func ValidateEndpoint(t *testing.T, config *MiConfig, endpointName string) {
	t.Helper()
	output, _ := getArtifact(t, endpointCmd, endpointName, config)
	artifactList := config.MIClient.GetArtifact(utils.MiManagementEndpointResource, "endpointName", endpointName, &artifactutils.Endpoint{})
	ValidateEndpointEqual(t, output, (artifactList.(*artifactutils.Endpoint)))
}

func ValidateEndpointEqual(t *testing.T, endpointFromCtl string, endpoint *artifactutils.Endpoint) {
	assert.Contains(t, endpointFromCtl, endpoint.Name)
	assert.Contains(t, endpointFromCtl, endpoint.Type)
	assert.Contains(t, endpointFromCtl, endpoint.Method)
	assert.Contains(t, endpointFromCtl, endpoint.URITemplate)
}
