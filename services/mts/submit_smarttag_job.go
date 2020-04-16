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

// SubmitSmarttagJob invokes the mts.SubmitSmarttagJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitsmarttagjob.html
func (client *Client) SubmitSmarttagJob(request *SubmitSmarttagJobRequest) (response *SubmitSmarttagJobResponse, err error) {
	response = CreateSubmitSmarttagJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitSmarttagJobWithChan invokes the mts.SubmitSmarttagJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitsmarttagjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitSmarttagJobWithChan(request *SubmitSmarttagJobRequest) (<-chan *SubmitSmarttagJobResponse, <-chan error) {
	responseChan := make(chan *SubmitSmarttagJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSmarttagJob(request)
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

// SubmitSmarttagJobWithCallback invokes the mts.SubmitSmarttagJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitsmarttagjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitSmarttagJobWithCallback(request *SubmitSmarttagJobRequest, callback func(response *SubmitSmarttagJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSmarttagJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitSmarttagJob(request)
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

// SubmitSmarttagJobRequest is the request struct for api SubmitSmarttagJob
type SubmitSmarttagJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Title                string           `position:"Query" name:"Title"`
	Content              string           `position:"Query" name:"Content"`
	UserData             string           `position:"Query" name:"UserData"`
	NotifyUrl            string           `position:"Query" name:"NotifyUrl"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Params               string           `position:"Query" name:"Params"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	ContentType          string           `position:"Query" name:"ContentType"`
	Input                string           `position:"Query" name:"Input"`
	ContentAddr          string           `position:"Query" name:"ContentAddr"`
}

// SubmitSmarttagJobResponse is the response struct for api SubmitSmarttagJob
type SubmitSmarttagJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitSmarttagJobRequest creates a request to invoke SubmitSmarttagJob API
func CreateSubmitSmarttagJobRequest() (request *SubmitSmarttagJobRequest) {
	request = &SubmitSmarttagJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitSmarttagJob", "mts", "openAPI")
	return
}

// CreateSubmitSmarttagJobResponse creates a response to parse from SubmitSmarttagJob response
func CreateSubmitSmarttagJobResponse() (response *SubmitSmarttagJobResponse) {
	response = &SubmitSmarttagJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
