package rds

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

func (client *Client) DescribeInstanceAutoRenewalAttribute(request *DescribeInstanceAutoRenewalAttributeRequest) (response *DescribeInstanceAutoRenewalAttributeResponse, err error) {
	response = CreateDescribeInstanceAutoRenewalAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceAutoRenewalAttributeWithChan(request *DescribeInstanceAutoRenewalAttributeRequest) (<-chan *DescribeInstanceAutoRenewalAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAutoRenewalAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAutoRenewalAttribute(request)
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

func (client *Client) DescribeInstanceAutoRenewalAttributeWithCallback(request *DescribeInstanceAutoRenewalAttributeRequest, callback func(response *DescribeInstanceAutoRenewalAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAutoRenewalAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAutoRenewalAttribute(request)
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

type DescribeInstanceAutoRenewalAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeInstanceAutoRenewalAttributeResponse struct {
	*responses.BaseResponse
	RequestId        string                                      `json:"RequestId" xml:"RequestId"`
	PageNumber       int                                         `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int                                         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                                         `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeInstanceAutoRenewalAttribute `json:"Items" xml:"Items"`
}

func CreateDescribeInstanceAutoRenewalAttributeRequest() (request *DescribeInstanceAutoRenewalAttributeRequest) {
	request = &DescribeInstanceAutoRenewalAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceAutoRenewalAttribute", "rds", "openAPI")
	return
}

func CreateDescribeInstanceAutoRenewalAttributeResponse() (response *DescribeInstanceAutoRenewalAttributeResponse) {
	response = &DescribeInstanceAutoRenewalAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
