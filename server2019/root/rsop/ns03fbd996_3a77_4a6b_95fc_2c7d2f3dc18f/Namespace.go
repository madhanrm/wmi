// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS03FBD996_3A77_4A6B_95FC_2C7D2F3DC18F
//////////////////////////////////////////////
package ns03fbd996_3a77_4a6b_95fc_2c7d2f3dc18f
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// __Namespace struct
type __Namespace struct { 
	*__SystemClass

	// 
	Name string
}

	func New__NamespaceEx1(instance *cim.WmiInstance) (newInstance *__Namespace, err error) {tmp, err := New__SystemClassEx1(instance)
		
	if err != nil { return }
	newInstance = &__Namespace {
	__SystemClass: tmp,
	}
	return
	}
	

	func New__NamespaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__Namespace, err error) {tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__Namespace {
	__SystemClass: tmp,
	}
	return
	}
	

// SetName sets the value of Name for the instance
func (instance *__Namespace) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *__Namespace) GetPropertyName()(value string, err error) { 
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

