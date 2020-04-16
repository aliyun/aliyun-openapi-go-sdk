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

// AddAsrPipeline invokes the mts.AddAsrPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/addasrpipeline.html
func (client *Client) AddAsrPipeline(request *AddAsrPipelineRequest) (response *AddAsrPipelineResponse, err error) {
	response = CreateAddAsrPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// AddAsrPipelineWithChan invokes the mts.AddAsrPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addasrpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddAsrPipelineWithChan(request *AddAsrPipelineRequest) (<-chan *AddAsrPipelineResponse, <-chan error) {
	responseChan := make(chan *AddAsrPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAsrPipeline(request)
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

// AddAsrPipelineWithCallback invokes the mts.AddAsrPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addasrpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddAsrPipelineWithCallback(request *AddAsrPipelineRequest, callback func(response *AddAsrPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAsrPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddAsrPipeline(request)
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

// AddAsrPipelineRequest is the request struct for api AddAsrPipeline
type AddAsrPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	Name                 string           `position:"Query" name:"Name"`
}

// AddAsrPipelineResponse is the response struct for api AddAsrPipeline
type AddAsrPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateAddAsrPipelineRequest creates a request to invoke AddAsrPipeline API
func CreateAddAsrPipelineRequest() (request *AddAsrPipelineRequest) {
	request = &AddAsrPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddAsrPipeline", "mts", "openAPI")
	return
}

// CreateAddAsrPipelineResponse creates a response to parse from AddAsrPipeline response
func CreateAddAsrPipelineResponse() (response *AddAsrPipelineResponse) {
	response = &AddAsrPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
