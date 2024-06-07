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

// WMI_Tracing_Client_Operations struct
type WMI_Tracing_Client_Operations struct { 
	*EventTrace

	// Enable Flags
	Flags Operations_Flags
}

	func NewWMI_Tracing_Client_OperationsEx1(instance *cim.WmiInstance) (newInstance *WMI_Tracing_Client_Operations, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &WMI_Tracing_Client_Operations {
	EventTrace: tmp,
	}
	return
	}
	

	func NewWMI_Tracing_Client_OperationsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WMI_Tracing_Client_Operations, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WMI_Tracing_Client_Operations {
	EventTrace: tmp,
	}
	return
	}
	

// SetFlags sets the value of Flags for the instance
func (instance *WMI_Tracing_Client_Operations) SetPropertyFlags(value Operations_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *WMI_Tracing_Client_Operations) GetPropertyFlags()(value Operations_Flags, err error) { 
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

    value = Operations_Flags(valuetmp)

    return
}

