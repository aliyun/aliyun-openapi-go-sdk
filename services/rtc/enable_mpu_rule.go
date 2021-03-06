package rtc

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

// EnableMPURule invokes the rtc.EnableMPURule API synchronously
func (client *Client) EnableMPURule(request *EnableMPURuleRequest) (response *EnableMPURuleResponse, err error) {
	response = CreateEnableMPURuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableMPURuleWithChan invokes the rtc.EnableMPURule API asynchronously
func (client *Client) EnableMPURuleWithChan(request *EnableMPURuleRequest) (<-chan *EnableMPURuleResponse, <-chan error) {
	responseChan := make(chan *EnableMPURuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableMPURule(request)
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

// EnableMPURuleWithCallback invokes the rtc.EnableMPURule API asynchronously
func (client *Client) EnableMPURuleWithCallback(request *EnableMPURuleRequest, callback func(response *EnableMPURuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableMPURuleResponse
		var err error
		defer close(result)
		response, err = client.EnableMPURule(request)
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

// EnableMPURuleRequest is the request struct for api EnableMPURule
type EnableMPURuleRequest struct {
	*requests.RpcRequest
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	RuleId  requests.Integer `position:"Query" name:"RuleId"`
}

// EnableMPURuleResponse is the response struct for api EnableMPURule
type EnableMPURuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableMPURuleRequest creates a request to invoke EnableMPURule API
func CreateEnableMPURuleRequest() (request *EnableMPURuleRequest) {
	request = &EnableMPURuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "EnableMPURule", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableMPURuleResponse creates a response to parse from EnableMPURule response
func CreateEnableMPURuleResponse() (response *EnableMPURuleResponse) {
	response = &EnableMPURuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
