package mts

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

// ActivateMediaWorkflow invokes the mts.ActivateMediaWorkflow API synchronously
// api document: https://help.aliyun.com/api/mts/activatemediaworkflow.html
func (client *Client) ActivateMediaWorkflow(request *ActivateMediaWorkflowRequest) (response *ActivateMediaWorkflowResponse, err error) {
	response = CreateActivateMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// ActivateMediaWorkflowWithChan invokes the mts.ActivateMediaWorkflow API asynchronously
// api document: https://help.aliyun.com/api/mts/activatemediaworkflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActivateMediaWorkflowWithChan(request *ActivateMediaWorkflowRequest) (<-chan *ActivateMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *ActivateMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActivateMediaWorkflow(request)
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

// ActivateMediaWorkflowWithCallback invokes the mts.ActivateMediaWorkflow API asynchronously
// api document: https://help.aliyun.com/api/mts/activatemediaworkflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActivateMediaWorkflowWithCallback(request *ActivateMediaWorkflowRequest, callback func(response *ActivateMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActivateMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.ActivateMediaWorkflow(request)
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

// ActivateMediaWorkflowRequest is the request struct for api ActivateMediaWorkflow
type ActivateMediaWorkflowRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaWorkflowId      string           `position:"Query" name:"MediaWorkflowId"`
}

// ActivateMediaWorkflowResponse is the response struct for api ActivateMediaWorkflow
type ActivateMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	MediaWorkflow MediaWorkflow `json:"MediaWorkflow" xml:"MediaWorkflow"`
}

// CreateActivateMediaWorkflowRequest creates a request to invoke ActivateMediaWorkflow API
func CreateActivateMediaWorkflowRequest() (request *ActivateMediaWorkflowRequest) {
	request = &ActivateMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ActivateMediaWorkflow", "mts", "openAPI")
	return
}

// CreateActivateMediaWorkflowResponse creates a response to parse from ActivateMediaWorkflow response
func CreateActivateMediaWorkflowResponse() (response *ActivateMediaWorkflowResponse) {
	response = &ActivateMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
