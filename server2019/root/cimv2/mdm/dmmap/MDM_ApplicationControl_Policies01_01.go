// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_ApplicationControl_Policies01_01 struct
type MDM_ApplicationControl_Policies01_01 struct { 
	*cim.WmiInstance

	// 
	InstanceID string

	// 
	ParentID string

	// 
	Policy string
}

	func NewMDM_ApplicationControl_Policies01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_ApplicationControl_Policies01_01, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_ApplicationControl_Policies01_01 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_ApplicationControl_Policies01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_ApplicationControl_Policies01_01, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_ApplicationControl_Policies01_01 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl_Policies01_01) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl_Policies01_01) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_Policies01_01) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_Policies01_01) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
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

// SetPolicy sets the value of Policy for the instance
func (instance *MDM_ApplicationControl_Policies01_01) SetPropertyPolicy(value string) (err error) { 
    return instance.SetProperty("Policy", (value))
}

// GetPolicy gets the value of Policy for the instance
func (instance *MDM_ApplicationControl_Policies01_01) GetPropertyPolicy()(value string, err error) { 
    retValue, err := instance.GetProperty("Policy")
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

