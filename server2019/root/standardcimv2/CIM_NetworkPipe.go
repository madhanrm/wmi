// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_NetworkPipe struct
type CIM_NetworkPipe struct {
	*CIM_EnabledLogicalElement

	//
	AggregationBehavior uint16

	//
	Directionality uint16
}

func NewCIM_NetworkPipeEx1(instance *cim.WmiInstance) (newInstance *CIM_NetworkPipe, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_NetworkPipe{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewCIM_NetworkPipeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_NetworkPipe, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_NetworkPipe{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetAggregationBehavior sets the value of AggregationBehavior for the instance
func (instance *CIM_NetworkPipe) SetPropertyAggregationBehavior(value uint16) (err error) {
	return instance.SetProperty("AggregationBehavior", value)
}

// GetAggregationBehavior gets the value of AggregationBehavior for the instance
func (instance *CIM_NetworkPipe) GetPropertyAggregationBehavior() (value uint16, err error) {
	retValue, err := instance.GetProperty("AggregationBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectionality sets the value of Directionality for the instance
func (instance *CIM_NetworkPipe) SetPropertyDirectionality(value uint16) (err error) {
	return instance.SetProperty("Directionality", value)
}

// GetDirectionality gets the value of Directionality for the instance
func (instance *CIM_NetworkPipe) GetPropertyDirectionality() (value uint16, err error) {
	retValue, err := instance.GetProperty("Directionality")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}