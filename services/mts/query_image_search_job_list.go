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

// QueryImageSearchJobList invokes the mts.QueryImageSearchJobList API synchronously
// api document: https://help.aliyun.com/api/mts/queryimagesearchjoblist.html
func (client *Client) QueryImageSearchJobList(request *QueryImageSearchJobListRequest) (response *QueryImageSearchJobListResponse, err error) {
	response = CreateQueryImageSearchJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryImageSearchJobListWithChan invokes the mts.QueryImageSearchJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryimagesearchjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryImageSearchJobListWithChan(request *QueryImageSearchJobListRequest) (<-chan *QueryImageSearchJobListResponse, <-chan error) {
	responseChan := make(chan *QueryImageSearchJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryImageSearchJobList(request)
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

// QueryImageSearchJobListWithCallback invokes the mts.QueryImageSearchJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryimagesearchjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryImageSearchJobListWithCallback(request *QueryImageSearchJobListRequest, callback func(response *QueryImageSearchJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryImageSearchJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryImageSearchJobList(request)
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

// QueryImageSearchJobListRequest is the request struct for api QueryImageSearchJobList
type QueryImageSearchJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
}

// QueryImageSearchJobListResponse is the response struct for api QueryImageSearchJobList
type QueryImageSearchJobListResponse struct {
	*responses.BaseResponse
	RequestId          string                               `json:"RequestId" xml:"RequestId"`
	NonExistIds        NonExistIdsInQueryImageSearchJobList `json:"NonExistIds" xml:"NonExistIds"`
	ImageSearchJobList ImageSearchJobList                   `json:"ImageSearchJobList" xml:"ImageSearchJobList"`
}

// CreateQueryImageSearchJobListRequest creates a request to invoke QueryImageSearchJobList API
func CreateQueryImageSearchJobListRequest() (request *QueryImageSearchJobListRequest) {
	request = &QueryImageSearchJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryImageSearchJobList", "mts", "openAPI")
	return
}

// CreateQueryImageSearchJobListResponse creates a response to parse from QueryImageSearchJobList response
func CreateQueryImageSearchJobListResponse() (response *QueryImageSearchJobListResponse) {
	response = &QueryImageSearchJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
