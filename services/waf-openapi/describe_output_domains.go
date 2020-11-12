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

// DescribeOutputDomains invokes the waf_openapi.DescribeOutputDomains API synchronously
func (client *Client) DescribeOutputDomains(request *DescribeOutputDomainsRequest) (response *DescribeOutputDomainsResponse, err error) {
	response = CreateDescribeOutputDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOutputDomainsWithChan invokes the waf_openapi.DescribeOutputDomains API asynchronously
func (client *Client) DescribeOutputDomainsWithChan(request *DescribeOutputDomainsRequest) (<-chan *DescribeOutputDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeOutputDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOutputDomains(request)
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

// DescribeOutputDomainsWithCallback invokes the waf_openapi.DescribeOutputDomains API asynchronously
func (client *Client) DescribeOutputDomainsWithCallback(request *DescribeOutputDomainsRequest, callback func(response *DescribeOutputDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOutputDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOutputDomains(request)
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

// DescribeOutputDomainsRequest is the request struct for api DescribeOutputDomains
type DescribeOutputDomainsRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	Domain          string           `position:"Query" name:"Domain"`
	Region          string           `position:"Query" name:"Region"`
}

// DescribeOutputDomainsResponse is the response struct for api DescribeOutputDomains
type DescribeOutputDomainsResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	WafTaskId  string           `json:"WafTaskId" xml:"WafTaskId"`
	TaskStatus int              `json:"TaskStatus" xml:"TaskStatus"`
	Total      int              `json:"Total" xml:"Total"`
	DomainList []DomainListItem `json:"DomainList" xml:"DomainList"`
}

// CreateDescribeOutputDomainsRequest creates a request to invoke DescribeOutputDomains API
func CreateDescribeOutputDomainsRequest() (request *DescribeOutputDomainsRequest) {
	request = &DescribeOutputDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeOutputDomains", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOutputDomainsResponse creates a response to parse from DescribeOutputDomains response
func CreateDescribeOutputDomainsResponse() (response *DescribeOutputDomainsResponse) {
	response = &DescribeOutputDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
