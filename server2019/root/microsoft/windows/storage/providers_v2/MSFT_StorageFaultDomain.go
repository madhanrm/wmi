// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_StorageFaultDomain struct
type MSFT_StorageFaultDomain struct { 
	*MSFT_StorageObject

	// A user settable description of the fault domain object.
	Description string

	// A user-friendly string representing the name of the fault domain object.
	FriendlyName string

	// 
	HealthStatus StorageFaultDomain_HealthStatus

	// This field represents the name of the company responsible for the hardware backing the fault domain oject. For physical disk it must match the disk's SCSI inquiry data.
	Manufacturer string

	// This field represents the model number of the hardware. For physical disk it must match the disk's SCSI inquiry data.
	Model string

	// 
	OperationalDetails []string

	// 
	OperationalStatus []StorageFaultDomain_OperationalStatus

	// This field is a free-form string indicating where the hardware is located.
	PhysicalLocation string

	// This field represents the serial number of the hardware. For physical disk it must match the disk's SCSI inquiry data.
	SerialNumber string
}

	func NewMSFT_StorageFaultDomainEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageFaultDomain, err error) {tmp, err := NewMSFT_StorageObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_StorageFaultDomain {
	MSFT_StorageObject: tmp,
	}
	return
	}
	

	func NewMSFT_StorageFaultDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_StorageFaultDomain, err error) {tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_StorageFaultDomain {
	MSFT_StorageObject: tmp,
	}
	return
	}
	

// SetDescription sets the value of Description for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyDescription(value string) (err error) { 
    return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("Description")
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyFriendlyName(value string) (err error) { 
    return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyFriendlyName()(value string, err error) { 
    retValue, err := instance.GetProperty("FriendlyName")
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

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyHealthStatus(value StorageFaultDomain_HealthStatus) (err error) { 
    return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyHealthStatus()(value StorageFaultDomain_HealthStatus, err error) { 
    retValue, err := instance.GetProperty("HealthStatus")
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

    value = StorageFaultDomain_HealthStatus(valuetmp)

    return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyManufacturer(value string) (err error) { 
    return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyManufacturer()(value string, err error) { 
    retValue, err := instance.GetProperty("Manufacturer")
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

// SetModel sets the value of Model for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyModel(value string) (err error) { 
    return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyModel()(value string, err error) { 
    retValue, err := instance.GetProperty("Model")
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

// SetOperationalDetails sets the value of OperationalDetails for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyOperationalDetails(value []string) (err error) { 
    return instance.SetProperty("OperationalDetails", (value))
}

// GetOperationalDetails gets the value of OperationalDetails for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyOperationalDetails()(value []string, err error) { 
    retValue, err := instance.GetProperty("OperationalDetails")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyOperationalStatus(value []StorageFaultDomain_OperationalStatus) (err error) { 
    return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyOperationalStatus()(value []StorageFaultDomain_OperationalStatus, err error) { 
    retValue, err := instance.GetProperty("OperationalStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(int32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, StorageFaultDomain_OperationalStatus(valuetmp))
    }

    return
}

// SetPhysicalLocation sets the value of PhysicalLocation for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyPhysicalLocation(value string) (err error) { 
    return instance.SetProperty("PhysicalLocation", (value))
}

// GetPhysicalLocation gets the value of PhysicalLocation for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyPhysicalLocation()(value string, err error) { 
    retValue, err := instance.GetProperty("PhysicalLocation")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertySerialNumber()(value string, err error) { 
    retValue, err := instance.GetProperty("SerialNumber")
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

