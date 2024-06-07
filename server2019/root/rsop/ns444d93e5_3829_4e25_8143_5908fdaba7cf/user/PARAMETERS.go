// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS444D93E5_3829_4E25_8143_5908FDABA7CF.User
//////////////////////////////////////////////
package user
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// __PARAMETERS struct
type __PARAMETERS struct { 
	*cim.WmiInstance
}

	func New__PARAMETERSEx1(instance *cim.WmiInstance) (newInstance *__PARAMETERS, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &__PARAMETERS {
	WmiInstance: tmp,
	}
	return
	}
	

	func New__PARAMETERSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__PARAMETERS, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__PARAMETERS {
	WmiInstance: tmp,
	}
	return
	}
	

