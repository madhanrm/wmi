// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetCallToFunctionFailed struct
type MSFT_NetCallToFunctionFailed struct {
	*MSFT_SCMEventLogEvent

	//
	Error uint32

	//
	FunctionName string
}

func NewMSFT_NetCallToFunctionFailedEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetCallToFunctionFailed, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetCallToFunctionFailed{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetCallToFunctionFailedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetCallToFunctionFailed, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetCallToFunctionFailed{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_NetCallToFunctionFailed) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetCallToFunctionFailed) GetPropertyError() (value uint32, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFunctionName sets the value of FunctionName for the instance
func (instance *MSFT_NetCallToFunctionFailed) SetPropertyFunctionName(value string) (err error) {
	return instance.SetProperty("FunctionName", value)
}

// GetFunctionName gets the value of FunctionName for the instance
func (instance *MSFT_NetCallToFunctionFailed) GetPropertyFunctionName() (value string, err error) {
	retValue, err := instance.GetProperty("FunctionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}