package ram

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

// DeleteVirtualMFADevice invokes the ram.DeleteVirtualMFADevice API synchronously
// api document: https://help.aliyun.com/api/ram/deletevirtualmfadevice.html
func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (response *DeleteVirtualMFADeviceResponse, err error) {
	response = CreateDeleteVirtualMFADeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVirtualMFADeviceWithChan invokes the ram.DeleteVirtualMFADevice API asynchronously
// api document: https://help.aliyun.com/api/ram/deletevirtualmfadevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVirtualMFADeviceWithChan(request *DeleteVirtualMFADeviceRequest) (<-chan *DeleteVirtualMFADeviceResponse, <-chan error) {
	responseChan := make(chan *DeleteVirtualMFADeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVirtualMFADevice(request)
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

// DeleteVirtualMFADeviceWithCallback invokes the ram.DeleteVirtualMFADevice API asynchronously
// api document: https://help.aliyun.com/api/ram/deletevirtualmfadevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVirtualMFADeviceWithCallback(request *DeleteVirtualMFADeviceRequest, callback func(response *DeleteVirtualMFADeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVirtualMFADeviceResponse
		var err error
		defer close(result)
		response, err = client.DeleteVirtualMFADevice(request)
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

// DeleteVirtualMFADeviceRequest is the request struct for api DeleteVirtualMFADevice
type DeleteVirtualMFADeviceRequest struct {
	*requests.RpcRequest
	SerialNumber string `position:"Query" name:"SerialNumber"`
}

// DeleteVirtualMFADeviceResponse is the response struct for api DeleteVirtualMFADevice
type DeleteVirtualMFADeviceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVirtualMFADeviceRequest creates a request to invoke DeleteVirtualMFADevice API
func CreateDeleteVirtualMFADeviceRequest() (request *DeleteVirtualMFADeviceRequest) {
	request = &DeleteVirtualMFADeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeleteVirtualMFADevice", "ram", "openAPI")
	return
}

// CreateDeleteVirtualMFADeviceResponse creates a response to parse from DeleteVirtualMFADevice response
func CreateDeleteVirtualMFADeviceResponse() (response *DeleteVirtualMFADeviceResponse) {
	response = &DeleteVirtualMFADeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}