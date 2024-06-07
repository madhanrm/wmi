// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_Counters_PacketDirectECUtilization struct
type Win32_PerfFormattedData_Counters_PacketDirectECUtilization struct { 
	*Win32_PerfFormattedData

	// 
	BusyWaitIterationsPersec uint32

	// 
	IterationsPersec uint32

	// 
	PercentBusyWaitingTime uint64

	// 
	PercentBusyWaitIterations uint32

	// 
	PercentIdleTime uint64

	// 
	PercentProcessingTime uint64

	// 
	ProcessorNumber uint32

	// 
	RXQueueCount uint32

	// 
	TotalBusyWaitIterations uint64

	// 
	TotalIterations uint64

	// 
	TXQueueCount uint32
}

	func NewWin32_PerfFormattedData_Counters_PacketDirectECUtilizationEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_PacketDirectECUtilization {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Counters_PacketDirectECUtilizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_PacketDirectECUtilization {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetBusyWaitIterationsPersec sets the value of BusyWaitIterationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyBusyWaitIterationsPersec(value uint32) (err error) { 
    return instance.SetProperty("BusyWaitIterationsPersec", (value))
}

// GetBusyWaitIterationsPersec gets the value of BusyWaitIterationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyBusyWaitIterationsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BusyWaitIterationsPersec")
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

// SetIterationsPersec sets the value of IterationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyIterationsPersec(value uint32) (err error) { 
    return instance.SetProperty("IterationsPersec", (value))
}

// GetIterationsPersec gets the value of IterationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyIterationsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IterationsPersec")
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

// SetPercentBusyWaitingTime sets the value of PercentBusyWaitingTime for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyPercentBusyWaitingTime(value uint64) (err error) { 
    return instance.SetProperty("PercentBusyWaitingTime", (value))
}

// GetPercentBusyWaitingTime gets the value of PercentBusyWaitingTime for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyPercentBusyWaitingTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PercentBusyWaitingTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetPercentBusyWaitIterations sets the value of PercentBusyWaitIterations for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyPercentBusyWaitIterations(value uint32) (err error) { 
    return instance.SetProperty("PercentBusyWaitIterations", (value))
}

// GetPercentBusyWaitIterations gets the value of PercentBusyWaitIterations for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyPercentBusyWaitIterations()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PercentBusyWaitIterations")
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

// SetPercentIdleTime sets the value of PercentIdleTime for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyPercentIdleTime(value uint64) (err error) { 
    return instance.SetProperty("PercentIdleTime", (value))
}

// GetPercentIdleTime gets the value of PercentIdleTime for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyPercentIdleTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PercentIdleTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetPercentProcessingTime sets the value of PercentProcessingTime for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyPercentProcessingTime(value uint64) (err error) { 
    return instance.SetProperty("PercentProcessingTime", (value))
}

// GetPercentProcessingTime gets the value of PercentProcessingTime for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyPercentProcessingTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PercentProcessingTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetProcessorNumber sets the value of ProcessorNumber for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyProcessorNumber(value uint32) (err error) { 
    return instance.SetProperty("ProcessorNumber", (value))
}

// GetProcessorNumber gets the value of ProcessorNumber for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyProcessorNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProcessorNumber")
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

// SetRXQueueCount sets the value of RXQueueCount for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyRXQueueCount(value uint32) (err error) { 
    return instance.SetProperty("RXQueueCount", (value))
}

// GetRXQueueCount gets the value of RXQueueCount for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyRXQueueCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RXQueueCount")
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

// SetTotalBusyWaitIterations sets the value of TotalBusyWaitIterations for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyTotalBusyWaitIterations(value uint64) (err error) { 
    return instance.SetProperty("TotalBusyWaitIterations", (value))
}

// GetTotalBusyWaitIterations gets the value of TotalBusyWaitIterations for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyTotalBusyWaitIterations()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalBusyWaitIterations")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetTotalIterations sets the value of TotalIterations for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyTotalIterations(value uint64) (err error) { 
    return instance.SetProperty("TotalIterations", (value))
}

// GetTotalIterations gets the value of TotalIterations for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyTotalIterations()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalIterations")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetTXQueueCount sets the value of TXQueueCount for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) SetPropertyTXQueueCount(value uint32) (err error) { 
    return instance.SetProperty("TXQueueCount", (value))
}

// GetTXQueueCount gets the value of TXQueueCount for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectECUtilization) GetPropertyTXQueueCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TXQueueCount")
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

