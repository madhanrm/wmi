// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS444D93E5_3829_4E25_8143_5908FDABA7CF
//////////////////////////////////////////////
package ns444d93e5_3829_4e25_8143_5908fdaba7cf
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ClassDeletion struct
type CIM_ClassDeletion struct { 
	*CIM_ClassIndication
}

	func NewCIM_ClassDeletionEx1(instance *cim.WmiInstance) (newInstance *CIM_ClassDeletion, err error) {tmp, err := NewCIM_ClassIndicationEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_ClassDeletion {
	CIM_ClassIndication: tmp,
	}
	return
	}
	

	func NewCIM_ClassDeletionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_ClassDeletion, err error) {tmp, err := NewCIM_ClassIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_ClassDeletion {
	CIM_ClassIndication: tmp,
	}
	return
	}
	

