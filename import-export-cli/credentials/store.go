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

package credentials

type Store interface {
	// Has an env in the store
	Has(env string) bool
	// Get an env from store returns intended Credential or an error
	Get(env string) (ApimCredential, error)
	// Get an env from store returns intended Credential or an error
	GetMICredentials(env string) (MiCredential, error)
	// Set credentials for env using given username,password,clientId,clientSecret
	SetMICredentials(env, username, password, accessToken string) error
	// Set credentials for env using given username,password,clientId,clientSecret
	Set(env, username, password, clientID, clientSecret string) error
	// Erase credentials from given env
	Erase(env string) error
	// Erase mi credentials from given env
	EraseMI(env string) error
	// Load store
	Load() error
}
