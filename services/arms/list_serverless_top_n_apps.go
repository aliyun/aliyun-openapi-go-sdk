package arms

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

// ListServerlessTopNApps invokes the arms.ListServerlessTopNApps API synchronously
func (client *Client) ListServerlessTopNApps(request *ListServerlessTopNAppsRequest) (response *ListServerlessTopNAppsResponse, err error) {
	response = CreateListServerlessTopNAppsResponse()
	err = client.DoAction(request, response)
	return
}

// ListServerlessTopNAppsWithChan invokes the arms.ListServerlessTopNApps API asynchronously
func (client *Client) ListServerlessTopNAppsWithChan(request *ListServerlessTopNAppsRequest) (<-chan *ListServerlessTopNAppsResponse, <-chan error) {
	responseChan := make(chan *ListServerlessTopNAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListServerlessTopNApps(request)
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

// ListServerlessTopNAppsWithCallback invokes the arms.ListServerlessTopNApps API asynchronously
func (client *Client) ListServerlessTopNAppsWithCallback(request *ListServerlessTopNAppsRequest, callback func(response *ListServerlessTopNAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListServerlessTopNAppsResponse
		var err error
		defer close(result)
		response, err = client.ListServerlessTopNApps(request)
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

// ListServerlessTopNAppsRequest is the request struct for api ListServerlessTopNApps
type ListServerlessTopNAppsRequest struct {
	*requests.RpcRequest
	Limit     requests.Integer `position:"Query" name:"Limit"`
	OrderBy   string           `position:"Query" name:"OrderBy"`
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
}

// ListServerlessTopNAppsResponse is the response struct for api ListServerlessTopNApps
type ListServerlessTopNAppsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListServerlessTopNAppsRequest creates a request to invoke ListServerlessTopNApps API
func CreateListServerlessTopNAppsRequest() (request *ListServerlessTopNAppsRequest) {
	request = &ListServerlessTopNAppsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListServerlessTopNApps", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListServerlessTopNAppsResponse creates a response to parse from ListServerlessTopNApps response
func CreateListServerlessTopNAppsResponse() (response *ListServerlessTopNAppsResponse) {
	response = &ListServerlessTopNAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
