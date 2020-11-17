package config

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

// ConfigurationListItem is a nested struct in config response
type ConfigurationListItem struct {
	AccountId          int64  `json:"AccountId" xml:"AccountId"`
	AvailabilityZone   string `json:"AvailabilityZone" xml:"AvailabilityZone"`
	CaptureTime        string `json:"CaptureTime" xml:"CaptureTime"`
	ConfigurationDiff  string `json:"ConfigurationDiff" xml:"ConfigurationDiff"`
	Region             string `json:"Region" xml:"Region"`
	Relationship       string `json:"Relationship" xml:"Relationship"`
	RelationshipDiff   string `json:"RelationshipDiff" xml:"RelationshipDiff"`
	ResourceCreateTime string `json:"ResourceCreateTime" xml:"ResourceCreateTime"`
	ResourceId         string `json:"ResourceId" xml:"ResourceId"`
	ResourceName       string `json:"ResourceName" xml:"ResourceName"`
	ResourceType       string `json:"ResourceType" xml:"ResourceType"`
	Tags               string `json:"Tags" xml:"Tags"`
	ResourceEventType  string `json:"ResourceEventType" xml:"ResourceEventType"`
}
