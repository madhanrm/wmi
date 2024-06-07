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

// MSNdis_TcpOffloadCurrentConfig struct
type MSNdis_TcpOffloadCurrentConfig struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string
}

	func NewMSNdis_TcpOffloadCurrentConfigEx1(instance *cim.WmiInstance) (newInstance *MSNdis_TcpOffloadCurrentConfig, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_TcpOffloadCurrentConfig {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_TcpOffloadCurrentConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_TcpOffloadCurrentConfig, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_TcpOffloadCurrentConfig {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_TcpOffloadCurrentConfig) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_TcpOffloadCurrentConfig) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_TcpOffloadCurrentConfig) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_TcpOffloadCurrentConfig) GetPropertyInstanceName()(value string, err error) { 
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

// 

// <param name="Header" type="MSNdis_WmiMethodHeader "></param>

// <param name="Offload" type="MSNdis_WmiOffload "></param>
func (instance *MSNdis_TcpOffloadCurrentConfig) WmiQueryCurrentOffloadConfig( /* IN */ Header MSNdis_WmiMethodHeader,
 /* OUT */ Offload MSNdis_WmiOffload) (err error) {_, err = instance.InvokeMethod("WmiQueryCurrentOffloadConfig" , Header)
	if err != nil { return }
	return
	
}


