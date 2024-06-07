

// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
//
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// DS_LDAP_Root_Class struct
type DS_LDAP_Root_Class struct { 
	*cim.WmiInstance
}

	func NewDS_LDAP_Root_ClassEx1(instance *cim.WmiInstance) (newInstance *DS_LDAP_Root_Class, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DS_LDAP_Root_Class {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDS_LDAP_Root_ClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DS_LDAP_Root_Class, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DS_LDAP_Root_Class {
	WmiInstance: tmp,
	}
	return
	}
	

