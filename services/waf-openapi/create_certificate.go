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

// CreateCertificate invokes the waf_openapi.CreateCertificate API synchronously
func (client *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
	response = CreateCreateCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCertificateWithChan invokes the waf_openapi.CreateCertificate API asynchronously
func (client *Client) CreateCertificateWithChan(request *CreateCertificateRequest) (<-chan *CreateCertificateResponse, <-chan error) {
	responseChan := make(chan *CreateCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCertificate(request)
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

// CreateCertificateWithCallback invokes the waf_openapi.CreateCertificate API asynchronously
func (client *Client) CreateCertificateWithCallback(request *CreateCertificateRequest, callback func(response *CreateCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCertificateResponse
		var err error
		defer close(result)
		response, err = client.CreateCertificate(request)
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

// CreateCertificateRequest is the request struct for api CreateCertificate
type CreateCertificateRequest struct {
	*requests.RpcRequest
	Certificate     string           `position:"Query" name:"Certificate"`
	PrivateKey      string           `position:"Query" name:"PrivateKey"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	Domain          string           `position:"Query" name:"Domain"`
	CertificateName string           `position:"Query" name:"CertificateName"`
	HttpsCertId     requests.Integer `position:"Query" name:"HttpsCertId"`
}

// CreateCertificateResponse is the response struct for api CreateCertificate
type CreateCertificateResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	CertificateId int64  `json:"CertificateId" xml:"CertificateId"`
}

// CreateCreateCertificateRequest creates a request to invoke CreateCertificate API
func CreateCreateCertificateRequest() (request *CreateCertificateRequest) {
	request = &CreateCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "CreateCertificate", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCertificateResponse creates a response to parse from CreateCertificate response
func CreateCreateCertificateResponse() (response *CreateCertificateResponse) {
	response = &CreateCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
