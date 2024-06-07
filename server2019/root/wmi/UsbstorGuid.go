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

// UsbstorGuid struct
type UsbstorGuid struct { 
	*EventTrace

	// Enable Flags
	Flags UsbstorGuid_Flags

	// Levels
	Level UsbstorGuid_Level
}

	func NewUsbstorGuidEx1(instance *cim.WmiInstance) (newInstance *UsbstorGuid, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &UsbstorGuid {
	EventTrace: tmp,
	}
	return
	}
	

	func NewUsbstorGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UsbstorGuid, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UsbstorGuid {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *UsbstorGuid) SetPropertyFlags(value UsbstorGuid_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *UsbstorGuid) GetPropertyFlags()(value UsbstorGuid_Flags, err error) { 
    retValue, err := instance.GetProperty("Flags")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = UsbstorGuid_Flags(valuetmp)

    return
}

// SetLevel sets the value of Level for the instance
func (instance *UsbstorGuid) SetPropertyLevel(value UsbstorGuid_Level) (err error) { 
    return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *UsbstorGuid) GetPropertyLevel()(value UsbstorGuid_Level, err error) { 
    retValue, err := instance.GetProperty("Level")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = UsbstorGuid_Level(valuetmp)

    return
}

