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

// Win32_FontInfoAction struct
type Win32_FontInfoAction struct {
	*CIM_Action

	//
	File string

	//
	FontTitle string
}

func NewWin32_FontInfoActionEx1(instance *cim.WmiInstance) (newInstance *Win32_FontInfoAction, err error) {
	tmp, err := NewCIM_ActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_FontInfoAction{
		CIM_Action: tmp,
	}
	return
}

func NewWin32_FontInfoActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_FontInfoAction, err error) {
	tmp, err := NewCIM_ActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_FontInfoAction{
		CIM_Action: tmp,
	}
	return
}

// SetFile sets the value of File for the instance
func (instance *Win32_FontInfoAction) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *Win32_FontInfoAction) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFontTitle sets the value of FontTitle for the instance
func (instance *Win32_FontInfoAction) SetPropertyFontTitle(value string) (err error) {
	return instance.SetProperty("FontTitle", value)
}

// GetFontTitle gets the value of FontTitle for the instance
func (instance *Win32_FontInfoAction) GetPropertyFontTitle() (value string, err error) {
	retValue, err := instance.GetProperty("FontTitle")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}