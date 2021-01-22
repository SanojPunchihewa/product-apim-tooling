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
	"flag"
	"os"
	"testing"

	"github.com/wso2/product-apim-tooling/import-export-cli/integration/base"
	"gopkg.in/yaml.v2"
)

var environment Environment
var miClient MiRESTClient
var config *MiConfig

func TestMain(m *testing.M) {
	flag.Parse()

	readConfigs()

	base.ExtractArchiveFile("../../build/target/")

	miClient = MiRESTClient{}
	miClient.SetupMI(adminUserName, adminPassword, environment.Name, environment.Host, environment.Offset)

	config = &MiConfig{
		Username: adminUserName,
		Password: adminPassword,
		MIClient: miClient,
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func readConfigs() {
	reader, err := os.Open("config.yaml")
	if err != nil {
		base.Fatal(err)
	}
	defer reader.Close()
	environment = Environment{}
	yaml.NewDecoder(reader).Decode(&environment)
	base.Log("env:", environment)
}
