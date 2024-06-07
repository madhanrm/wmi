// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS69B56437_FE3D_4523_BE69_32394088573F.User
//////////////////////////////////////////////
package user
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrPowerOptionsV2Setting struct
type RSOP_PolmkrPowerOptionsV2Setting struct { 
	*RSOP_PolmkrPowerSetting
}

	func NewRSOP_PolmkrPowerOptionsV2SettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrPowerOptionsV2Setting, err error) {tmp, err := NewRSOP_PolmkrPowerSettingEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_PolmkrPowerOptionsV2Setting {
	RSOP_PolmkrPowerSetting: tmp,
	}
	return
	}
	

	func NewRSOP_PolmkrPowerOptionsV2SettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_PolmkrPowerOptionsV2Setting, err error) {tmp, err := NewRSOP_PolmkrPowerSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_PolmkrPowerOptionsV2Setting {
	RSOP_PolmkrPowerSetting: tmp,
	}
	return
	}
	

