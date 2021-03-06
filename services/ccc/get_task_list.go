package ccc

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

// GetTaskList invokes the ccc.GetTaskList API synchronously
func (client *Client) GetTaskList(request *GetTaskListRequest) (response *GetTaskListResponse, err error) {
	response = CreateGetTaskListResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskListWithChan invokes the ccc.GetTaskList API asynchronously
func (client *Client) GetTaskListWithChan(request *GetTaskListRequest) (<-chan *GetTaskListResponse, <-chan error) {
	responseChan := make(chan *GetTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaskList(request)
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

// GetTaskListWithCallback invokes the ccc.GetTaskList API asynchronously
func (client *Client) GetTaskListWithCallback(request *GetTaskListRequest, callback func(response *GetTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskListResponse
		var err error
		defer close(result)
		response, err = client.GetTaskList(request)
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

// GetTaskListRequest is the request struct for api GetTaskList
type GetTaskListRequest struct {
	*requests.RpcRequest
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetTaskListResponse is the response struct for api GetTaskList
type GetTaskListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Tasks          []Task `json:"Tasks" xml:"Tasks"`
}

// CreateGetTaskListRequest creates a request to invoke GetTaskList API
func CreateGetTaskListRequest() (request *GetTaskListRequest) {
	request = &GetTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetTaskList", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTaskListResponse creates a response to parse from GetTaskList response
func CreateGetTaskListResponse() (response *GetTaskListResponse) {
	response = &GetTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
