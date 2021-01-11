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
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	miUtils "github.com/wso2/product-apim-tooling/import-export-cli/mi/utils"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
	"golang.org/x/crypto/ssh/terminal"
)

var keyStorePropertiesFile = miUtils.GetkeyStorePropertiesFilePath()
var inputPropertiesfile string
var encryptionAlgorithm string
var outputType string

const secretCreateCmdLiteral = "create"
const secretCreateCmdShortDesc = "Encrypt secrets"

const secretCreateCmdLongDesc = "Create secrets based on given arguments"

var secretCreateCmdExamples = "To encrypt secret and get output on console\n" +
	"  " + utils.ProjectName + " " + utils.MiCmdLiteral + " " + secretCmdLiteral + " " + secretCreateCmdLiteral + "\n" +
	"To encrypt secret and get output as a .properties file (stored in the security folder in apictl executable directory)\n" +
	"  " + utils.ProjectName + " " + utils.MiCmdLiteral + " " + secretCmdLiteral + " " + secretCreateCmdLiteral + " -o file\n" +
	"To encrypt secret and get output as a .yaml file (stored in the security folder in apictl executable directory)\n" +
	"  " + utils.ProjectName + " " + utils.MiCmdLiteral + " " + secretCmdLiteral + " " + secretCreateCmdLiteral + " -o k8\n" +
	"To bulk encrypt secrets defined in a properties file\n" +
	"  " + utils.ProjectName + " " + utils.MiCmdLiteral + " " + secretCmdLiteral + " " + secretCreateCmdLiteral + " -f <file_path>\n" +
	"To bulk encrypt secrets defined in a properties file and get a .yaml file (stored in the security folder in apictl executable directory)\n" +
	"  " + utils.ProjectName + " " + utils.MiCmdLiteral + " " + secretCmdLiteral + " " + secretCreateCmdLiteral + " -o k8 -f <file_path>"

var secretCreateCmd = &cobra.Command{
	Use:     secretCreateCmdLiteral,
	Short:   secretCreateCmdShortDesc,
	Long:    secretCreateCmdLongDesc,
	Example: secretCreateCmdExamples,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsFileExist(keyStorePropertiesFile) {
			utils.HandleErrorAndExit("Keystore has not been initialized.\nExecute 'apictl mi secret init --help' for more information", nil)
		}
		initSecretInformation()
	},
}

func init() {
	SecretCmd.AddCommand(secretCreateCmd)
	secretCreateCmd.Flags().StringVarP(&inputPropertiesfile, "from-file", "f", "", "Path to the properties file which contain secrets to be encrypted")
	secretCreateCmd.Flags().StringVarP(&outputType, "output", "o", "console", "Get the output in yaml(k8) or properties(file) format. By default the output is printed to the console")
	secretCreateCmd.Flags().StringVarP(&encryptionAlgorithm, "cipher", "c", "", "Encryption algorithm")
}

func initSecretInformation() {
	encryptionClientPath, err := miUtils.GetEncryptionClientPath()
	if err != nil {
		utils.HandleErrorAndExit(err.Error(), nil)
	}
	var inputs = make(map[string]string)
	inputs["secret.output.type"] = outputType
	if isNonEmptyString(encryptionAlgorithm) {
		os.Setenv("secret.encryption.algorithm", encryptionAlgorithm)
	}
	if isNonEmptyString(inputPropertiesfile) {
		inputs["secret.input.type"] = "file"
		inputs["secret.input.file"] = miUtils.NormalizeFilePath(inputPropertiesfile)
	} else {
		inputs["secret.input.type"] = "console"
		startConsoleForSecretInfo(inputs)
	}
	if miUtils.IsMapWithNonEmptyValues(inputs) {
		secretInfoFilePath := path.Join(miUtils.GetSecurityDirectoryPath(), "secret-info.properties")
		os.Setenv("secret.source", secretInfoFilePath)
		os.Setenv("keystore.source", keyStorePropertiesFile)
		miUtils.WritePropertiesToFile(inputs, secretInfoFilePath)
		execEncryptionClient(encryptionClientPath)
		os.Remove(secretInfoFilePath)
	} else {
		utils.HandleErrorAndExit("Incomplete secret information", nil)
	}
}

func execEncryptionClient(encryptionClientPath string) {
	var stdoutMessage []byte
	var command *exec.Cmd
	commandString := "java -jar " + encryptionClientPath
	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/c", commandString)
	} else {
		command = exec.Command("bash", "-c", commandString)
	}
	stdoutMessage, _ = command.CombinedOutput()
	fmt.Printf("%s", stdoutMessage)
}

func startConsoleForSecretInfo(params map[string]string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter plain alias for secret:")
	updateMap(params, "secret.plaintext.alias", reader)

	fmt.Printf("Enter plain text secret:")
	byteSecret, _ := terminal.ReadPassword(int(syscall.Stdin))
	secret := string(byteSecret)
	fmt.Println()

	fmt.Printf("Repeat plain text secret:")
	byteRepeatSecret, _ := terminal.ReadPassword(int(syscall.Stdin))
	repeatSecret := string(byteRepeatSecret)
	fmt.Println()

	if isMatchingSecrets(secret, repeatSecret) {
		params["secret.plaintext.secret"] = strings.TrimSpace(secret)
	} else {
		fmt.Println("Entered secret values did not match.")
		startConsoleForSecretInfo(params)
	}
}

func isMatchingSecrets(secret, repeatSecret string) bool {
	if secret == repeatSecret {
		return true
	}
	return false
}

func isNonEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}
