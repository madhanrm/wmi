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

// Win32_MemoryArray struct
type Win32_MemoryArray struct {
	*Win32_SMBIOSMemory

	//
	ErrorGranularity uint16
}

func NewWin32_MemoryArrayEx1(instance *cim.WmiInstance) (newInstance *Win32_MemoryArray, err error) {
	tmp, err := NewWin32_SMBIOSMemoryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_MemoryArray{
		Win32_SMBIOSMemory: tmp,
	}
	return
}

func NewWin32_MemoryArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_MemoryArray, err error) {
	tmp, err := NewWin32_SMBIOSMemoryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_MemoryArray{
		Win32_SMBIOSMemory: tmp,
	}
	return
}

// SetErrorGranularity sets the value of ErrorGranularity for the instance
func (instance *Win32_MemoryArray) SetPropertyErrorGranularity(value uint16) (err error) {
	return instance.SetProperty("ErrorGranularity", value)
}

// GetErrorGranularity gets the value of ErrorGranularity for the instance
func (instance *Win32_MemoryArray) GetPropertyErrorGranularity() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorGranularity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}