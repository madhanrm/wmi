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

// SamSetMemAttrsOfGrp_Start struct
type SamSetMemAttrsOfGrp_Start struct { 
	*SamSetMemAttrsOfGrp

	// Client Network Address
	Client string

	// Signature
	Sam string

	// Client SID
	Sid string

	// SamTraceVersion
	Version uint32
}

	func NewSamSetMemAttrsOfGrp_StartEx1(instance *cim.WmiInstance) (newInstance *SamSetMemAttrsOfGrp_Start, err error) {tmp, err := NewSamSetMemAttrsOfGrpEx1(instance)
		
	if err != nil { return }
	newInstance = &SamSetMemAttrsOfGrp_Start {
	SamSetMemAttrsOfGrp: tmp,
	}
	return
	}
	

	func NewSamSetMemAttrsOfGrp_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamSetMemAttrsOfGrp_Start, err error) {tmp, err := NewSamSetMemAttrsOfGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamSetMemAttrsOfGrp_Start {
	SamSetMemAttrsOfGrp: tmp,
	}
	return
	}
	

// SetClient sets the value of Client for the instance
func (instance *SamSetMemAttrsOfGrp_Start) SetPropertyClient(value string) (err error) { 
    return instance.SetProperty("Client", (value))
}

// GetClient gets the value of Client for the instance
func (instance *SamSetMemAttrsOfGrp_Start) GetPropertyClient()(value string, err error) { 
    retValue, err := instance.GetProperty("Client")
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

// SetSam sets the value of Sam for the instance
func (instance *SamSetMemAttrsOfGrp_Start) SetPropertySam(value string) (err error) { 
    return instance.SetProperty("Sam", (value))
}

// GetSam gets the value of Sam for the instance
func (instance *SamSetMemAttrsOfGrp_Start) GetPropertySam()(value string, err error) { 
    retValue, err := instance.GetProperty("Sam")
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

// SetSid sets the value of Sid for the instance
func (instance *SamSetMemAttrsOfGrp_Start) SetPropertySid(value string) (err error) { 
    return instance.SetProperty("Sid", (value))
}

// GetSid gets the value of Sid for the instance
func (instance *SamSetMemAttrsOfGrp_Start) GetPropertySid()(value string, err error) { 
    retValue, err := instance.GetProperty("Sid")
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

// SetVersion sets the value of Version for the instance
func (instance *SamSetMemAttrsOfGrp_Start) SetPropertyVersion(value uint32) (err error) { 
    return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *SamSetMemAttrsOfGrp_Start) GetPropertyVersion()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Version")
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

