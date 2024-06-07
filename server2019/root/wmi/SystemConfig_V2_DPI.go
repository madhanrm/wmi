// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SystemConfig_V2_DPI struct
type SystemConfig_V2_DPI struct { 
	*SystemConfig_V2

	// 
	MachineDPI uint32

	// 
	UserDPI uint32
}

	func NewSystemConfig_V2_DPIEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_DPI, err error) {tmp, err := NewSystemConfig_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SystemConfig_V2_DPI {
	SystemConfig_V2: tmp,
	}
	return
	}
	

	func NewSystemConfig_V2_DPIEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemConfig_V2_DPI, err error) {tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemConfig_V2_DPI {
	SystemConfig_V2: tmp,
	}
	return
	}
	

// SetMachineDPI sets the value of MachineDPI for the instance
func (instance *SystemConfig_V2_DPI) SetPropertyMachineDPI(value uint32) (err error) { 
    return instance.SetProperty("MachineDPI", (value))
}

// GetMachineDPI gets the value of MachineDPI for the instance
func (instance *SystemConfig_V2_DPI) GetPropertyMachineDPI()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MachineDPI")
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

// SetUserDPI sets the value of UserDPI for the instance
func (instance *SystemConfig_V2_DPI) SetPropertyUserDPI(value uint32) (err error) { 
    return instance.SetProperty("UserDPI", (value))
}

// GetUserDPI gets the value of UserDPI for the instance
func (instance *SystemConfig_V2_DPI) GetPropertyUserDPI()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UserDPI")
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

