// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEConnectionDialUpCredentialsLink struct
type RSOP_IEConnectionDialUpCredentialsLink struct {
	*cim.WmiInstance

	//
	dialUpCredentials RSOP_IEConnectionDialUpCredentials

	//
	policySetting RSOP_IEAKPolicySetting
}

func NewRSOP_IEConnectionDialUpCredentialsLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionDialUpCredentialsLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpCredentialsLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionDialUpCredentialsLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionDialUpCredentialsLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpCredentialsLink{
		WmiInstance: tmp,
	}
	return
}

// SetdialUpCredentials sets the value of dialUpCredentials for the instance
func (instance *RSOP_IEConnectionDialUpCredentialsLink) SetPropertydialUpCredentials(value RSOP_IEConnectionDialUpCredentials) (err error) {
	return instance.SetProperty("dialUpCredentials", value)
}

// GetdialUpCredentials gets the value of dialUpCredentials for the instance
func (instance *RSOP_IEConnectionDialUpCredentialsLink) GetPropertydialUpCredentials() (value RSOP_IEConnectionDialUpCredentials, err error) {
	retValue, err := instance.GetProperty("dialUpCredentials")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEConnectionDialUpCredentials)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEConnectionDialUpCredentialsLink) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", value)
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEConnectionDialUpCredentialsLink) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
	retValue, err := instance.GetProperty("policySetting")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_IEAKPolicySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}