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

// Win32_ShortcutFile struct
type Win32_ShortcutFile struct {
	*CIM_DataFile

	//
	Target string
}

func NewWin32_ShortcutFileEx1(instance *cim.WmiInstance) (newInstance *Win32_ShortcutFile, err error) {
	tmp, err := NewCIM_DataFileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ShortcutFile{
		CIM_DataFile: tmp,
	}
	return
}

func NewWin32_ShortcutFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ShortcutFile, err error) {
	tmp, err := NewCIM_DataFileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ShortcutFile{
		CIM_DataFile: tmp,
	}
	return
}

// SetTarget sets the value of Target for the instance
func (instance *Win32_ShortcutFile) SetPropertyTarget(value string) (err error) {
	return instance.SetProperty("Target", value)
}

// GetTarget gets the value of Target for the instance
func (instance *Win32_ShortcutFile) GetPropertyTarget() (value string, err error) {
	retValue, err := instance.GetProperty("Target")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}