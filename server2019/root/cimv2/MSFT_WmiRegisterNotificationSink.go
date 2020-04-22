// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_WmiRegisterNotificationSink struct
type MSFT_WmiRegisterNotificationSink struct {
	*MSFT_WmiEssEvent

	//
	Namespace string

	//
	Query string

	//
	QueryLanguage string

	//
	Sink uint64
}

func NewMSFT_WmiRegisterNotificationSinkEx1(instance *cim.WmiInstance) (newInstance *MSFT_WmiRegisterNotificationSink, err error) {
	tmp, err := NewMSFT_WmiEssEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiRegisterNotificationSink{
		MSFT_WmiEssEvent: tmp,
	}
	return
}

func NewMSFT_WmiRegisterNotificationSinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WmiRegisterNotificationSink, err error) {
	tmp, err := NewMSFT_WmiEssEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiRegisterNotificationSink{
		MSFT_WmiEssEvent: tmp,
	}
	return
}

// SetNamespace sets the value of Namespace for the instance
func (instance *MSFT_WmiRegisterNotificationSink) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", value)
}

// GetNamespace gets the value of Namespace for the instance
func (instance *MSFT_WmiRegisterNotificationSink) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *MSFT_WmiRegisterNotificationSink) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *MSFT_WmiRegisterNotificationSink) GetPropertyQuery() (value string, err error) {
	retValue, err := instance.GetProperty("Query")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueryLanguage sets the value of QueryLanguage for the instance
func (instance *MSFT_WmiRegisterNotificationSink) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *MSFT_WmiRegisterNotificationSink) GetPropertyQueryLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("QueryLanguage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSink sets the value of Sink for the instance
func (instance *MSFT_WmiRegisterNotificationSink) SetPropertySink(value uint64) (err error) {
	return instance.SetProperty("Sink", value)
}

// GetSink gets the value of Sink for the instance
func (instance *MSFT_WmiRegisterNotificationSink) GetPropertySink() (value uint64, err error) {
	retValue, err := instance.GetProperty("Sink")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}