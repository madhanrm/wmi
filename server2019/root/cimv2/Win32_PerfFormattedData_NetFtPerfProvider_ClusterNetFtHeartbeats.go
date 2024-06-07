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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats struct
type Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats struct { 
	*Win32_PerfFormattedData

	// 
	Missingheartbeats uint32

	// 
	Missingheartbeatslimit uint32
}

	func NewWin32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeatsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeatsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetMissingheartbeats sets the value of Missingheartbeats for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) SetPropertyMissingheartbeats(value uint32) (err error) { 
    return instance.SetProperty("Missingheartbeats", (value))
}

// GetMissingheartbeats gets the value of Missingheartbeats for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) GetPropertyMissingheartbeats()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Missingheartbeats")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMissingheartbeatslimit sets the value of Missingheartbeatslimit for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) SetPropertyMissingheartbeatslimit(value uint32) (err error) { 
    return instance.SetProperty("Missingheartbeatslimit", (value))
}

// GetMissingheartbeatslimit gets the value of Missingheartbeatslimit for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) GetPropertyMissingheartbeatslimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Missingheartbeatslimit")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

