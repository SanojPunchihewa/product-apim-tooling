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

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/go-resty/resty"
	"github.com/wso2/product-apim-tooling/import-export-cli/credentials"
	"github.com/wso2/product-apim-tooling/import-export-cli/formatter"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

// unmarshalData unmarshal data from the response to the respective struct
// @param url: url of rest api
// @param params: parameters for the HTTP call
// @param env: environment of the micro integrator instance
// @param model: struct object
// @return struct object
// @return error
func unmarshalData(url string, params map[string]string, env string, model interface{}) (interface{}, error) {

	resp, err := invokeGETRequestWithRetry(url, params, env)

	if err != nil {
		utils.HandleErrorAndExit("Unable to connect to "+url, err)
	}

	utils.Logln(utils.LogPrefixInfo+"Response:", resp.Status())

	if resp.StatusCode() == http.StatusOK {
		response := model
		unmarshalError := json.Unmarshal(resp.Body(), &response)

		if unmarshalError != nil {
			utils.HandleErrorAndExit(utils.LogPrefixError+"invalid JSON response", unmarshalError)
		}
		return response, nil
	}
	if resp.StatusCode() == http.StatusUnauthorized {
		fmt.Println("Invalid credentials. Please login to the current Micro Integrator instance")
		utils.HandleErrorAndExit("Execute 'apictl login --help' for more information", nil)
	}
	if len(resp.Body()) == 0 {
		return nil, errors.New(resp.Status())
	}
	data := unmarshalJSONToStringMap(resp.Body())
	return data["Error"], errors.New(resp.Status())
}

func retryHTTPCall(attempts int, env string, f func(string) (*resty.Response, error)) (*resty.Response, error) {
	cred, err := credentials.GetMICredentials(env)
	resp, err := f(cred.AccessToken)
	if resp.StatusCode() == http.StatusUnauthorized {
		if attempts--; attempts > 0 {
			token, err := credentials.GetOAuthAccessTokenForMI(cred.Username, cred.Password, env)
			if err != nil {
				return nil, err
			}
			credentials.UpdateMIAccessToken(env, token)
			return retryHTTPCall(attempts, env, f)
		}
	}
	return resp, err
}

func invokeGETRequestWithRetry(url string, params map[string]string, env string) (*resty.Response, error) {

	return retryHTTPCall(utils.MiHTTPRetryCount, env, func(accessToken string) (*resty.Response, error) {
		headers := make(map[string]string)
		headers[utils.HeaderAuthorization] = utils.HeaderValueAuthBearerPrefix + " " + accessToken
		return utils.InvokeGETRequestWithMultipleQueryParams(params, url, headers)
	})
}

func unmarshalJSONToStringMap(body []byte) map[string]string {
	var data map[string]string
	unmarshalError := json.Unmarshal(body, &data)
	if unmarshalError != nil {
		utils.HandleErrorAndExit(utils.LogPrefixError+"invalid JSON response", unmarshalError)
	}
	return data
}

func getItemRenderer(data interface{}) func(w io.Writer, t *template.Template) error {
	return func(w io.Writer, t *template.Template) error {
		if err := t.Execute(w, data); err != nil {
			return err
		}
		return nil
	}
}

func getArtifactInfo(resource, artifactKey, artifactName, env string, model interface{}) (interface{}, error) {

	params := make(map[string]string)
	params[artifactKey] = artifactName

	return getArtifacts(resource, params, env, model)
}

func getArtifactList(resource, env string, model interface{}) (interface{}, error) {

	return getArtifacts(resource, nil, env, model)
}

func getArtifacts(resource string, params map[string]string, env string, model interface{}) (interface{}, error) {

	url := utils.GetMIManagementEndpointOfResource(resource, env, utils.MainConfigFilePath)

	resp, err := unmarshalData(url, params, env, model)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getContextWithFormat(format, defaultformat string) *formatter.Context {

	if format == "" {
		format = defaultformat
	}
	return formatter.NewContext(os.Stdout, format)
}
