// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NSCBC80CE2_C2E8_4CBC_8C8C_1C374E56CEFC
//////////////////////////////////////////////
package nscbc80ce2_c2e8_4cbc_8c8c_1c374e56cefc
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceCreationEvent struct
type __InstanceCreationEvent struct { 
	*__InstanceOperationEvent
}

	func New__InstanceCreationEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceCreationEvent, err error) {tmp, err := New__InstanceOperationEventEx1(instance)
		
	if err != nil { return }
	newInstance = &__InstanceCreationEvent {
	__InstanceOperationEvent: tmp,
	}
	return
	}
	

	func New__InstanceCreationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__InstanceCreationEvent, err error) {tmp, err := New__InstanceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__InstanceCreationEvent {
	__InstanceOperationEvent: tmp,
	}
	return
	}
	

