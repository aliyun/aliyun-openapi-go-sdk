package airec

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// MetaInListFilteringAlgorithms is a nested struct in airec response
type MetaInListFilteringAlgorithms struct {
	MetaType      string    `json:"metaType" xml:"metaType"`
	AlgorithmName string    `json:"algorithmName" xml:"algorithmName"`
	Cron          string    `json:"cron" xml:"cron"`
	CronEnabled   bool      `json:"cronEnabled" xml:"cronEnabled"`
	TaskId        string    `json:"taskId" xml:"taskId"`
	ProjectName   string    `json:"projectName" xml:"projectName"`
	TableName     string    `json:"tableName" xml:"tableName"`
	Type          string    `json:"type" xml:"type"`
	Category      string    `json:"category" xml:"category"`
	ClusterId     string    `json:"clusterId" xml:"clusterId"`
	Description   string    `json:"description" xml:"description"`
	ExtInfo       ExtInfo   `json:"extInfo" xml:"extInfo"`
	Threshold     Threshold `json:"threshold" xml:"threshold"`
}
