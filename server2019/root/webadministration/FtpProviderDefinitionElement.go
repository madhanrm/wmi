// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FtpProviderDefinitionElement struct
type FtpProviderDefinitionElement struct { 
	*CollectionElement

	// 
	ClsId string

	// 
	Name string

	// 
	Type string
}

	func NewFtpProviderDefinitionElementEx1(instance *cim.WmiInstance) (newInstance *FtpProviderDefinitionElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpProviderDefinitionElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewFtpProviderDefinitionElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpProviderDefinitionElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpProviderDefinitionElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetClsId sets the value of ClsId for the instance
func (instance *FtpProviderDefinitionElement) SetPropertyClsId(value string) (err error) { 
    return instance.SetProperty("ClsId", (value))
}

// GetClsId gets the value of ClsId for the instance
func (instance *FtpProviderDefinitionElement) GetPropertyClsId()(value string, err error) { 
    retValue, err := instance.GetProperty("ClsId")
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

// SetName sets the value of Name for the instance
func (instance *FtpProviderDefinitionElement) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *FtpProviderDefinitionElement) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetType sets the value of Type for the instance
func (instance *FtpProviderDefinitionElement) SetPropertyType(value string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *FtpProviderDefinitionElement) GetPropertyType()(value string, err error) { 
    retValue, err := instance.GetProperty("Type")
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

