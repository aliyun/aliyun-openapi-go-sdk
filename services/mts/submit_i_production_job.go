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

// SubmitIProductionJob invokes the mts.SubmitIProductionJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitiproductionjob.html
func (client *Client) SubmitIProductionJob(request *SubmitIProductionJobRequest) (response *SubmitIProductionJobResponse, err error) {
	response = CreateSubmitIProductionJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitIProductionJobWithChan invokes the mts.SubmitIProductionJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitiproductionjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitIProductionJobWithChan(request *SubmitIProductionJobRequest) (<-chan *SubmitIProductionJobResponse, <-chan error) {
	responseChan := make(chan *SubmitIProductionJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitIProductionJob(request)
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

// SubmitIProductionJobWithCallback invokes the mts.SubmitIProductionJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitiproductionjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitIProductionJobWithCallback(request *SubmitIProductionJobRequest, callback func(response *SubmitIProductionJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitIProductionJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitIProductionJob(request)
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

// SubmitIProductionJobRequest is the request struct for api SubmitIProductionJob
type SubmitIProductionJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobParams            string           `position:"Query" name:"JobParams"`
	UserData             string           `position:"Query" name:"UserData"`
	FunctionName         string           `position:"Query" name:"FunctionName"`
	NotifyUrl            string           `position:"Query" name:"NotifyUrl"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ModelId              string           `position:"Query" name:"ModelId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	ScheduleParams       string           `position:"Query" name:"ScheduleParams"`
}

// SubmitIProductionJobResponse is the response struct for api SubmitIProductionJob
type SubmitIProductionJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitIProductionJobRequest creates a request to invoke SubmitIProductionJob API
func CreateSubmitIProductionJobRequest() (request *SubmitIProductionJobRequest) {
	request = &SubmitIProductionJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitIProductionJob", "mts", "openAPI")
	return
}

// CreateSubmitIProductionJobResponse creates a response to parse from SubmitIProductionJob response
func CreateSubmitIProductionJobResponse() (response *SubmitIProductionJobResponse) {
	response = &SubmitIProductionJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
