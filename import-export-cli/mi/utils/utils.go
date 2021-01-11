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

package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

// GetTrimmedCmdLiteral returns the command without the arguments
func GetTrimmedCmdLiteral(cmd string) string {
	cmdParts := strings.Fields(cmd)
	return cmdParts[0]
}

// GetEncryptionClientPath returns the path to the encryption jar client
func GetEncryptionClientPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	programDir := filepath.Dir(exe)
	_ = os.Setenv("wso2.apictl.home", programDir)
	content, err := os.Open(programDir)
	if err != nil {
		log.Fatal(err)
	}
	files, _ := content.Readdir(-1)
	content.Close()
	for _, file := range files {
		if strings.Contains(file.Name(), "encryption-client") {
			return programDir + string(os.PathSeparator) + file.Name(), nil
		}
	}
	return "", errors.New(utils.LogPrefixError + "Encryption client library is missing")
}

// NormalizeFilePath replace the \ file separator with / for windows path
func NormalizeFilePath(path string) string {
	if strings.Contains(path, "\\") {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return strings.TrimSpace(path)
}

// IsMapWithNonEmptyValues iterates over a map and return false if there is an empty value
func IsMapWithNonEmptyValues(inputs map[string]string) bool {
	for key, input := range inputs {
		if len(strings.TrimSpace(input)) == 0 {
			fmt.Println("Invalid input for " + key)
			return false
		}
	}
	return true
}

// WritePropertiesToFile write a map to a .properties file
func WritePropertiesToFile(variables map[string]string, fileName string) {
	props := properties.LoadMap(variables)
	writer, _ := os.Create(fileName)
	props.Write(writer, properties.UTF8)
	writer.Close()
}

// GetSecurityDirectoryPath join mi-security with the config directory path
func GetSecurityDirectoryPath() string {
	return filepath.Join(utils.ConfigDirPath, "mi-security")
}

// GetkeyStorePropertiesFilePath join keystore-info.properties with the mi-security path
func GetkeyStorePropertiesFilePath() string {
	return path.Join(GetSecurityDirectoryPath(), "keystore-info.properties")
}
