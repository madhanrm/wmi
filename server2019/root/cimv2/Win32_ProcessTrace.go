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

// Win32_ProcessTrace struct
type Win32_ProcessTrace struct {
	*Win32_SystemTrace

	//
	ParentProcessID uint32

	//
	ProcessID uint32

	//
	ProcessName string

	//
	SessionID uint32

	//
	Sid []uint8
}

func NewWin32_ProcessTraceEx1(instance *cim.WmiInstance) (newInstance *Win32_ProcessTrace, err error) {
	tmp, err := NewWin32_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ProcessTrace{
		Win32_SystemTrace: tmp,
	}
	return
}

func NewWin32_ProcessTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ProcessTrace, err error) {
	tmp, err := NewWin32_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ProcessTrace{
		Win32_SystemTrace: tmp,
	}
	return
}

// SetParentProcessID sets the value of ParentProcessID for the instance
func (instance *Win32_ProcessTrace) SetPropertyParentProcessID(value uint32) (err error) {
	return instance.SetProperty("ParentProcessID", value)
}

// GetParentProcessID gets the value of ParentProcessID for the instance
func (instance *Win32_ProcessTrace) GetPropertyParentProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParentProcessID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessID sets the value of ProcessID for the instance
func (instance *Win32_ProcessTrace) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", value)
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Win32_ProcessTrace) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessName sets the value of ProcessName for the instance
func (instance *Win32_ProcessTrace) SetPropertyProcessName(value string) (err error) {
	return instance.SetProperty("ProcessName", value)
}

// GetProcessName gets the value of ProcessName for the instance
func (instance *Win32_ProcessTrace) GetPropertyProcessName() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSessionID sets the value of SessionID for the instance
func (instance *Win32_ProcessTrace) SetPropertySessionID(value uint32) (err error) {
	return instance.SetProperty("SessionID", value)
}

// GetSessionID gets the value of SessionID for the instance
func (instance *Win32_ProcessTrace) GetPropertySessionID() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSid sets the value of Sid for the instance
func (instance *Win32_ProcessTrace) SetPropertySid(value []uint8) (err error) {
	return instance.SetProperty("Sid", value)
}

// GetSid gets the value of Sid for the instance
func (instance *Win32_ProcessTrace) GetPropertySid() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Sid")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}