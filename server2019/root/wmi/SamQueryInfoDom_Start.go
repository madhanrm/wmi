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

// SamQueryInfoDom_Start struct
type SamQueryInfoDom_Start struct { 
	*SamQueryInfoDom

	// Client Network Address
	Client string

	// Signature
	Sam string

	// Client SID
	Sid string

	// SamTraceVersion
	Version uint32
}

	func NewSamQueryInfoDom_StartEx1(instance *cim.WmiInstance) (newInstance *SamQueryInfoDom_Start, err error) {tmp, err := NewSamQueryInfoDomEx1(instance)
		
	if err != nil { return }
	newInstance = &SamQueryInfoDom_Start {
	SamQueryInfoDom: tmp,
	}
	return
	}
	

	func NewSamQueryInfoDom_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamQueryInfoDom_Start, err error) {tmp, err := NewSamQueryInfoDomEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamQueryInfoDom_Start {
	SamQueryInfoDom: tmp,
	}
	return
	}
	

// SetClient sets the value of Client for the instance
func (instance *SamQueryInfoDom_Start) SetPropertyClient(value string) (err error) { 
    return instance.SetProperty("Client", (value))
}

// GetClient gets the value of Client for the instance
func (instance *SamQueryInfoDom_Start) GetPropertyClient()(value string, err error) { 
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
func (instance *SamQueryInfoDom_Start) SetPropertySam(value string) (err error) { 
    return instance.SetProperty("Sam", (value))
}

// GetSam gets the value of Sam for the instance
func (instance *SamQueryInfoDom_Start) GetPropertySam()(value string, err error) { 
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
func (instance *SamQueryInfoDom_Start) SetPropertySid(value string) (err error) { 
    return instance.SetProperty("Sid", (value))
}

// GetSid gets the value of Sid for the instance
func (instance *SamQueryInfoDom_Start) GetPropertySid()(value string, err error) { 
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
func (instance *SamQueryInfoDom_Start) SetPropertyVersion(value uint32) (err error) { 
    return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *SamQueryInfoDom_Start) GetPropertyVersion()(value uint32, err error) { 
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

