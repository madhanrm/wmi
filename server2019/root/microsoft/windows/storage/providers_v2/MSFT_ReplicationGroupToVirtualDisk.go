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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_ReplicationGroupToVirtualDisk struct
type MSFT_ReplicationGroupToVirtualDisk struct { 
	*cim.WmiInstance

	// 
	ReplicationGroup MSFT_ReplicationGroup

	// 
	VirtualDisk MSFT_VirtualDisk
}

	func NewMSFT_ReplicationGroupToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_ReplicationGroupToVirtualDisk, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_ReplicationGroupToVirtualDisk {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_ReplicationGroupToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_ReplicationGroupToVirtualDisk, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_ReplicationGroupToVirtualDisk {
	WmiInstance: tmp,
	}
	return
	}
	

// SetReplicationGroup sets the value of ReplicationGroup for the instance
func (instance *MSFT_ReplicationGroupToVirtualDisk) SetPropertyReplicationGroup(value MSFT_ReplicationGroup) (err error) { 
    return instance.SetProperty("ReplicationGroup", (value))
}

// GetReplicationGroup gets the value of ReplicationGroup for the instance
func (instance *MSFT_ReplicationGroupToVirtualDisk) GetPropertyReplicationGroup()(value MSFT_ReplicationGroup, err error) { 
    retValue, err := instance.GetProperty("ReplicationGroup")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSFT_ReplicationGroup); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSFT_ReplicationGroup is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSFT_ReplicationGroup(valuetmp)

    return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_ReplicationGroupToVirtualDisk) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) { 
    return instance.SetProperty("VirtualDisk", (value))
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_ReplicationGroupToVirtualDisk) GetPropertyVirtualDisk()(value MSFT_VirtualDisk, err error) { 
    retValue, err := instance.GetProperty("VirtualDisk")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSFT_VirtualDisk); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSFT_VirtualDisk is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSFT_VirtualDisk(valuetmp)

    return
}

