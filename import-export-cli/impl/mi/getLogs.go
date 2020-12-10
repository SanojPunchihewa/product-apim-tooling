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

package mi

// import (
// 	"fmt"
// 	"io"
// 	"text/template"

// 	"github.com/wso2/product-apim-tooling/import-export-cli/credentials"
// 	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
// 	"github.com/wso2/product-apim-tooling/import-export-cli/utils/mi/artifactutils"
// )

// const (
// 	loggerNameHeader      = "NAME"
// 	loggerLevelHeader     = "LOG LEVEL"
// 	loggerComponentHeader = "COMPONENT"

// 	defaultLoggerTableFormat = "table {{.LoggerName}}\t{{.LogLevel}}\t{{.ComponentName}}"
// )

// // GetLoggerInfo returns information about a specific logger
// func GetLoggerInfo(cred credentials.MiCredential, env, loggerName string) (*artifactutils.Logger, error) {

// 	resp, err := getArtifactInfo(utils.MiManagementLoggingResource, "loggerName", loggerName, env, &artifactutils.Logger{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.(*artifactutils.Logger), nil
// }

// // PrintLoggerInfo prints details about a logger
// func PrintLoggerInfo(logger *artifactutils.Logger, format string) {

// 	loggerContext := getContextWithFormat(format, defaultLoggerTableFormat)

// 	renderer := func(w io.Writer, t *template.Template) error {
// 		if err := t.Execute(w, logger); err != nil {
// 			return err
// 		}
// 		_, _ = w.Write([]byte{'\n'})
// 		return nil
// 	}

// 	loggerInfoTableHeaders := map[string]string{
// 		"LoggerName":    loggerNameHeader,
// 		"LogLevel":      loggerLevelHeader,
// 		"ComponentName": loggerComponentHeader,
// 	}

// 	if err := loggerContext.Write(renderer, loggerInfoTableHeaders); err != nil {
// 		fmt.Println("Error executing template:", err.Error())
// 	}
// }
