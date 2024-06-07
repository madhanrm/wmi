// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_POTSModemToSerialPort struct
type Win32_POTSModemToSerialPort struct { 
	*CIM_ControlledBy
}

	func NewWin32_POTSModemToSerialPortEx1(instance *cim.WmiInstance) (newInstance *Win32_POTSModemToSerialPort, err error) {tmp, err := NewCIM_ControlledByEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_POTSModemToSerialPort {
	CIM_ControlledBy: tmp,
	}
	return
	}
	

	func NewWin32_POTSModemToSerialPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_POTSModemToSerialPort, err error) {tmp, err := NewCIM_ControlledByEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_POTSModemToSerialPort {
	CIM_ControlledBy: tmp,
	}
	return
	}
	

