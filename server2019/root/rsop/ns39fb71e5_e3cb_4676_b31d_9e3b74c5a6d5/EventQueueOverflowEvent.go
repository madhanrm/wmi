// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS39FB71E5_E3CB_4676_B31D_9E3B74C5A6D5
//////////////////////////////////////////////
package ns39fb71e5_e3cb_4676_b31d_9e3b74c5a6d5
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// __EventQueueOverflowEvent struct
type __EventQueueOverflowEvent struct { 
	*__EventDroppedEvent

	// 
	CurrentQueueSize uint32
}

	func New__EventQueueOverflowEventEx1(instance *cim.WmiInstance) (newInstance *__EventQueueOverflowEvent, err error) {tmp, err := New__EventDroppedEventEx1(instance)
		
	if err != nil { return }
	newInstance = &__EventQueueOverflowEvent {
	__EventDroppedEvent: tmp,
	}
	return
	}
	

	func New__EventQueueOverflowEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__EventQueueOverflowEvent, err error) {tmp, err := New__EventDroppedEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__EventQueueOverflowEvent {
	__EventDroppedEvent: tmp,
	}
	return
	}
	

// SetCurrentQueueSize sets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) SetPropertyCurrentQueueSize(value uint32) (err error) { 
    return instance.SetProperty("CurrentQueueSize", (value))
}

// GetCurrentQueueSize gets the value of CurrentQueueSize for the instance
func (instance *__EventQueueOverflowEvent) GetPropertyCurrentQueueSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentQueueSize")
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

