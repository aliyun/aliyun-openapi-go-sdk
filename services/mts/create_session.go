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

// CreateSession invokes the mts.CreateSession API synchronously
// api document: https://help.aliyun.com/api/mts/createsession.html
func (client *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
	response = CreateCreateSessionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSessionWithChan invokes the mts.CreateSession API asynchronously
// api document: https://help.aliyun.com/api/mts/createsession.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSessionWithChan(request *CreateSessionRequest) (<-chan *CreateSessionResponse, <-chan error) {
	responseChan := make(chan *CreateSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSession(request)
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

// CreateSessionWithCallback invokes the mts.CreateSession API asynchronously
// api document: https://help.aliyun.com/api/mts/createsession.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSessionWithCallback(request *CreateSessionRequest, callback func(response *CreateSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSessionResponse
		var err error
		defer close(result)
		response, err = client.CreateSession(request)
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

// CreateSessionRequest is the request struct for api CreateSession
type CreateSessionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	SessionTime          requests.Integer `position:"Query" name:"SessionTime"`
	EndUserId            string           `position:"Query" name:"EndUserId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
}

// CreateSessionResponse is the response struct for api CreateSession
type CreateSessionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SessionId string `json:"SessionId" xml:"SessionId"`
	Ticket    string `json:"Ticket" xml:"Ticket"`
}

// CreateCreateSessionRequest creates a request to invoke CreateSession API
func CreateCreateSessionRequest() (request *CreateSessionRequest) {
	request = &CreateSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "CreateSession", "mts", "openAPI")
	return
}

// CreateCreateSessionResponse creates a response to parse from CreateSession response
func CreateCreateSessionResponse() (response *CreateSessionResponse) {
	response = &CreateSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
