package waf_openapi

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeProtectionModuleMode invokes the waf_openapi.DescribeProtectionModuleMode API synchronously
func (client *Client) DescribeProtectionModuleMode(request *DescribeProtectionModuleModeRequest) (response *DescribeProtectionModuleModeResponse, err error) {
	response = CreateDescribeProtectionModuleModeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProtectionModuleModeWithChan invokes the waf_openapi.DescribeProtectionModuleMode API asynchronously
func (client *Client) DescribeProtectionModuleModeWithChan(request *DescribeProtectionModuleModeRequest) (<-chan *DescribeProtectionModuleModeResponse, <-chan error) {
	responseChan := make(chan *DescribeProtectionModuleModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProtectionModuleMode(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeProtectionModuleModeWithCallback invokes the waf_openapi.DescribeProtectionModuleMode API asynchronously
func (client *Client) DescribeProtectionModuleModeWithCallback(request *DescribeProtectionModuleModeRequest, callback func(response *DescribeProtectionModuleModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProtectionModuleModeResponse
		var err error
		defer close(result)
		response, err = client.DescribeProtectionModuleMode(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeProtectionModuleModeRequest is the request struct for api DescribeProtectionModuleMode
type DescribeProtectionModuleModeRequest struct {
	*requests.RpcRequest
	DefenseType     string `position:"Query" name:"DefenseType"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeProtectionModuleModeResponse is the response struct for api DescribeProtectionModuleMode
type DescribeProtectionModuleModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Mode      int    `json:"Mode" xml:"Mode"`
}

// CreateDescribeProtectionModuleModeRequest creates a request to invoke DescribeProtectionModuleMode API
func CreateDescribeProtectionModuleModeRequest() (request *DescribeProtectionModuleModeRequest) {
	request = &DescribeProtectionModuleModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeProtectionModuleMode", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProtectionModuleModeResponse creates a response to parse from DescribeProtectionModuleMode response
func CreateDescribeProtectionModuleModeResponse() (response *DescribeProtectionModuleModeResponse) {
	response = &DescribeProtectionModuleModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
