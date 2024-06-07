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

// MSNdis_TokenRingCurrentAddress struct
type MSNdis_TokenRingCurrentAddress struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string

	// 
	NdisCurrentAddress MSNdis_NetworkAddress
}

	func NewMSNdis_TokenRingCurrentAddressEx1(instance *cim.WmiInstance) (newInstance *MSNdis_TokenRingCurrentAddress, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_TokenRingCurrentAddress {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_TokenRingCurrentAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_TokenRingCurrentAddress, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_TokenRingCurrentAddress {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_TokenRingCurrentAddress) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_TokenRingCurrentAddress) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_TokenRingCurrentAddress) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_TokenRingCurrentAddress) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// SetNdisCurrentAddress sets the value of NdisCurrentAddress for the instance
func (instance *MSNdis_TokenRingCurrentAddress) SetPropertyNdisCurrentAddress(value MSNdis_NetworkAddress) (err error) { 
    return instance.SetProperty("NdisCurrentAddress", (value))
}

// GetNdisCurrentAddress gets the value of NdisCurrentAddress for the instance
func (instance *MSNdis_TokenRingCurrentAddress) GetPropertyNdisCurrentAddress()(value MSNdis_NetworkAddress, err error) { 
    retValue, err := instance.GetProperty("NdisCurrentAddress")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_NetworkAddress); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_NetworkAddress is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_NetworkAddress(valuetmp)

    return
}

