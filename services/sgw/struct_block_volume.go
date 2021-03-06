package sgw

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

// BlockVolume is a nested struct in sgw response
type BlockVolume struct {
	Name          string `json:"Name" xml:"Name"`
	DiskId        string `json:"DiskId" xml:"DiskId"`
	DiskType      string `json:"DiskType" xml:"DiskType"`
	Protocol      string `json:"Protocol" xml:"Protocol"`
	Size          int64  `json:"Size" xml:"Size"`
	Enabled       bool   `json:"Enabled" xml:"Enabled"`
	State         string `json:"State" xml:"State"`
	TotalUpload   int64  `json:"TotalUpload" xml:"TotalUpload"`
	TotalDownload int64  `json:"TotalDownload" xml:"TotalDownload"`
	OssBucketName string `json:"OssBucketName" xml:"OssBucketName"`
	OssEndpoint   string `json:"OssEndpoint" xml:"OssEndpoint"`
	OssBucketSsl  bool   `json:"OssBucketSsl" xml:"OssBucketSsl"`
	LocalPath     string `json:"LocalPath" xml:"LocalPath"`
	ChunkSize     int    `json:"ChunkSize" xml:"ChunkSize"`
	CacheMode     string `json:"CacheMode" xml:"CacheMode"`
	Address       string `json:"Address" xml:"Address"`
	SerialNumber  string `json:"SerialNumber" xml:"SerialNumber"`
	IndexId       string `json:"IndexId" xml:"IndexId"`
	Target        string `json:"Target" xml:"Target"`
	Port          int    `json:"Port" xml:"Port"`
	LunId         int    `json:"LunId" xml:"LunId"`
	ChapEnabled   bool   `json:"ChapEnabled" xml:"ChapEnabled"`
	ChapInUser    string `json:"ChapInUser" xml:"ChapInUser"`
	ChapOutUser   string `json:"ChapOutUser" xml:"ChapOutUser"`
	Status        int    `json:"Status" xml:"Status"`
	VolumeState   int    `json:"VolumeState" xml:"VolumeState"`
}
