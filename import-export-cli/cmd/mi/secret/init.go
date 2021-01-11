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

package secret

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	miUtils "github.com/wso2/product-apim-tooling/import-export-cli/mi/utils"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
	"golang.org/x/crypto/ssh/terminal"
)

const secretInitCmdLiteral = "init"
const secretInitCmdShortDesc = "Initialize keystore"

const secretInitCmdLongDesc = "Initialize the key store information required for secret encryption"

var secretInitCmdExamples = "To initialize keystore information\n" +
	"  " + utils.ProjectName + " " + utils.MiCmdLiteral + " " + secretCmdLiteral + " " + secretInitCmdLiteral

var secretInitCmd = &cobra.Command{
	Use:     secretInitCmdLiteral,
	Short:   secretInitCmdShortDesc,
	Long:    secretInitCmdLongDesc,
	Example: secretInitCmdExamples,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := miUtils.GetEncryptionClientPath()
		if err != nil {
			utils.HandleErrorAndExit(err.Error(), nil)
		}
		startConsoleForKeyStore()
	},
}

func init() {
	SecretCmd.AddCommand(secretInitCmd)
}

func startConsoleForKeyStore() {
	reader := bufio.NewReader(os.Stdin)
	var keyStoreConfigMap = make(map[string]string)

	fmt.Printf("Enter Key Store location: ")
	updateMap(keyStoreConfigMap, "secret.keystore.location", reader)

	fmt.Printf("Enter Key Store type: ")
	updateMap(keyStoreConfigMap, "secret.keystore.type", reader)

	fmt.Printf("Enter Key alias: ")
	updateMap(keyStoreConfigMap, "secret.keystore.alias", reader)

	fmt.Printf("Enter Key Store password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	keyPassword := string(bytePassword)
	fmt.Println()
	keyStoreConfigMap["secret.keystore.password"] = base64.StdEncoding.EncodeToString([]byte(strings.TrimSpace(keyPassword)))

	if miUtils.IsMapWithNonEmptyValues(keyStoreConfigMap) {
		utils.CreateDirIfNotExist(miUtils.GetSecurityDirectoryPath())
		keystorePropertiesPath := miUtils.GetkeyStorePropertiesFilePath()
		miUtils.WritePropertiesToFile(keyStoreConfigMap, keystorePropertiesPath)
		fmt.Println("keystore initialization completed.")
	} else {
		fmt.Println("keystore initialization failed.")
	}
}

func updateMap(params map[string]string, key string, reader *bufio.Reader) {
	param, _ := reader.ReadString('\n')
	params[key] = strings.TrimSpace(param)
}
