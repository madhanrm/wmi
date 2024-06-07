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

// MSNdis_WmiSetHeader struct
type MSNdis_WmiSetHeader struct { 
	*MSNdis

	// 
	Header MSNdis_ObjectHeader

	// 
	NetLuid uint64

	// 
	Padding uint32

	// 
	PortNumber uint32

	// 
	RequestId uint64

	// 
	Timeout uint32
}

	func NewMSNdis_WmiSetHeaderEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiSetHeader, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiSetHeader {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_WmiSetHeaderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_WmiSetHeader, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiSetHeader {
	MSNdis: tmp,
	}
	return
	}
	

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_WmiSetHeader) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) { 
    return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiSetHeader) GetPropertyHeader()(value MSNdis_ObjectHeader, err error) { 
    retValue, err := instance.GetProperty("Header")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_ObjectHeader); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_ObjectHeader(valuetmp)

    return
}

// SetNetLuid sets the value of NetLuid for the instance
func (instance *MSNdis_WmiSetHeader) SetPropertyNetLuid(value uint64) (err error) { 
    return instance.SetProperty("NetLuid", (value))
}

// GetNetLuid gets the value of NetLuid for the instance
func (instance *MSNdis_WmiSetHeader) GetPropertyNetLuid()(value uint64, err error) { 
    retValue, err := instance.GetProperty("NetLuid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetPadding sets the value of Padding for the instance
func (instance *MSNdis_WmiSetHeader) SetPropertyPadding(value uint32) (err error) { 
    return instance.SetProperty("Padding", (value))
}

// GetPadding gets the value of Padding for the instance
func (instance *MSNdis_WmiSetHeader) GetPropertyPadding()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Padding")
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

// SetPortNumber sets the value of PortNumber for the instance
func (instance *MSNdis_WmiSetHeader) SetPropertyPortNumber(value uint32) (err error) { 
    return instance.SetProperty("PortNumber", (value))
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *MSNdis_WmiSetHeader) GetPropertyPortNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PortNumber")
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

// SetRequestId sets the value of RequestId for the instance
func (instance *MSNdis_WmiSetHeader) SetPropertyRequestId(value uint64) (err error) { 
    return instance.SetProperty("RequestId", (value))
}

// GetRequestId gets the value of RequestId for the instance
func (instance *MSNdis_WmiSetHeader) GetPropertyRequestId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("RequestId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetTimeout sets the value of Timeout for the instance
func (instance *MSNdis_WmiSetHeader) SetPropertyTimeout(value uint32) (err error) { 
    return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *MSNdis_WmiSetHeader) GetPropertyTimeout()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Timeout")
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

