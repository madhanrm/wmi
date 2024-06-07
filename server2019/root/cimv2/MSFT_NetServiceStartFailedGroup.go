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

// MSFT_NetServiceStartFailedGroup struct
type MSFT_NetServiceStartFailedGroup struct { 
	*MSFT_SCMEventLogEvent

	// 
	Group string

	// 
	Service string
}

	func NewMSFT_NetServiceStartFailedGroupEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetServiceStartFailedGroup, err error) {tmp, err := NewMSFT_SCMEventLogEventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetServiceStartFailedGroup {
	MSFT_SCMEventLogEvent: tmp,
	}
	return
	}
	

	func NewMSFT_NetServiceStartFailedGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetServiceStartFailedGroup, err error) {tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetServiceStartFailedGroup {
	MSFT_SCMEventLogEvent: tmp,
	}
	return
	}
	

// SetGroup sets the value of Group for the instance
func (instance *MSFT_NetServiceStartFailedGroup) SetPropertyGroup(value string) (err error) { 
    return instance.SetProperty("Group", (value))
}

// GetGroup gets the value of Group for the instance
func (instance *MSFT_NetServiceStartFailedGroup) GetPropertyGroup()(value string, err error) { 
    retValue, err := instance.GetProperty("Group")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceStartFailedGroup) SetPropertyService(value string) (err error) { 
    return instance.SetProperty("Service", (value))
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceStartFailedGroup) GetPropertyService()(value string, err error) { 
    retValue, err := instance.GetProperty("Service")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

