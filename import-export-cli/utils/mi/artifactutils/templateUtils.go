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

package artifactutils

type TemplateList struct {
	SequenceTemplates []Template `json:"sequenceTemplateList"`
	EndpointTemplates []Template `json:"endpointTemplateList"`
}

type TemplateListByType struct {
	Count     int32      `json:"count"`
	Templates []Template `json:"list"`
}

type TemplateSequenceListByName struct {
	Parameters []TemplateSequenceDetail `json:"Parameters"`
	Name       string                   `json:"name"`
}

type TemplateEndpointListByName struct {
	Parameters []string `json:"Parameters"`
	Name       string   `json:"name"`
}

type Template struct {
	Name string `json:"name"`
}

type TemplateSequenceDetail struct {
	Name         string `json:"name"`
	IsMandatory  bool   `json:"mandatory"`
	DefaultValue string `json:"defaultValue"`
}
