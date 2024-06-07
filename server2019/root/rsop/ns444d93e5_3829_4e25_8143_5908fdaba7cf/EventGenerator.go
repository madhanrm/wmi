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

// __EventGenerator struct
type __EventGenerator struct { 
	*__IndicationRelated
}

	func New__EventGeneratorEx1(instance *cim.WmiInstance) (newInstance *__EventGenerator, err error) {tmp, err := New__IndicationRelatedEx1(instance)
		
	if err != nil { return }
	newInstance = &__EventGenerator {
	__IndicationRelated: tmp,
	}
	return
	}
	

	func New__EventGeneratorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__EventGenerator, err error) {tmp, err := New__IndicationRelatedEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__EventGenerator {
	__IndicationRelated: tmp,
	}
	return
	}
	

