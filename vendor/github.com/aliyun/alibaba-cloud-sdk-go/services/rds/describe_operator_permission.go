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

func (client *Client) DescribeOperatorPermission(request *DescribeOperatorPermissionRequest) (response *DescribeOperatorPermissionResponse, err error) {
	response = CreateDescribeOperatorPermissionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeOperatorPermissionWithChan(request *DescribeOperatorPermissionRequest) (<-chan *DescribeOperatorPermissionResponse, <-chan error) {
	responseChan := make(chan *DescribeOperatorPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOperatorPermission(request)
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

func (client *Client) DescribeOperatorPermissionWithCallback(request *DescribeOperatorPermissionRequest, callback func(response *DescribeOperatorPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOperatorPermissionResponse
		var err error
		defer close(result)
		response, err = client.DescribeOperatorPermission(request)
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

type DescribeOperatorPermissionRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeOperatorPermissionResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Privileges  string `json:"Privileges" xml:"Privileges"`
	CreatedTime string `json:"CreatedTime" xml:"CreatedTime"`
	ExpiredTime string `json:"ExpiredTime" xml:"ExpiredTime"`
}

func CreateDescribeOperatorPermissionRequest() (request *DescribeOperatorPermissionRequest) {
	request = &DescribeOperatorPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOperatorPermission", "rds", "openAPI")
	return
}

func CreateDescribeOperatorPermissionResponse() (response *DescribeOperatorPermissionResponse) {
	response = &DescribeOperatorPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
