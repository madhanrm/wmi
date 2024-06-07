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

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages struct { 
	*Win32_PerfFormattedData

	// 
	AverageDatabaseMessagesExecutionTime uint32

	// 
	AverageMessagesExecutionTime uint32

	// 
	AverageWaitingTimeToExecuteDatabaseMessages uint32

	// 
	AverageWaitingTimeToExecuteMessages uint32

	// 
	DatabaseMessagesQueueLength uint64

	// 
	DatabaseUpdateMessages uint64

	// 
	DatabaseUpdateMessagesPersec uint64

	// 
	MessagesExecutionQueueLength uint64

	// 
	MessagesQueueLength uint64

	// 
	UpdateMessages uint64

	// 
	UpdateMessagesPersec uint64
}

	func NewWin32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessagesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessagesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetAverageDatabaseMessagesExecutionTime sets the value of AverageDatabaseMessagesExecutionTime for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyAverageDatabaseMessagesExecutionTime(value uint32) (err error) { 
    return instance.SetProperty("AverageDatabaseMessagesExecutionTime", (value))
}

// GetAverageDatabaseMessagesExecutionTime gets the value of AverageDatabaseMessagesExecutionTime for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyAverageDatabaseMessagesExecutionTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AverageDatabaseMessagesExecutionTime")
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

// SetAverageMessagesExecutionTime sets the value of AverageMessagesExecutionTime for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyAverageMessagesExecutionTime(value uint32) (err error) { 
    return instance.SetProperty("AverageMessagesExecutionTime", (value))
}

// GetAverageMessagesExecutionTime gets the value of AverageMessagesExecutionTime for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyAverageMessagesExecutionTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AverageMessagesExecutionTime")
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

// SetAverageWaitingTimeToExecuteDatabaseMessages sets the value of AverageWaitingTimeToExecuteDatabaseMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyAverageWaitingTimeToExecuteDatabaseMessages(value uint32) (err error) { 
    return instance.SetProperty("AverageWaitingTimeToExecuteDatabaseMessages", (value))
}

// GetAverageWaitingTimeToExecuteDatabaseMessages gets the value of AverageWaitingTimeToExecuteDatabaseMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyAverageWaitingTimeToExecuteDatabaseMessages()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AverageWaitingTimeToExecuteDatabaseMessages")
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

// SetAverageWaitingTimeToExecuteMessages sets the value of AverageWaitingTimeToExecuteMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyAverageWaitingTimeToExecuteMessages(value uint32) (err error) { 
    return instance.SetProperty("AverageWaitingTimeToExecuteMessages", (value))
}

// GetAverageWaitingTimeToExecuteMessages gets the value of AverageWaitingTimeToExecuteMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyAverageWaitingTimeToExecuteMessages()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AverageWaitingTimeToExecuteMessages")
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

// SetDatabaseMessagesQueueLength sets the value of DatabaseMessagesQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyDatabaseMessagesQueueLength(value uint64) (err error) { 
    return instance.SetProperty("DatabaseMessagesQueueLength", (value))
}

// GetDatabaseMessagesQueueLength gets the value of DatabaseMessagesQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyDatabaseMessagesQueueLength()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DatabaseMessagesQueueLength")
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

// SetDatabaseUpdateMessages sets the value of DatabaseUpdateMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyDatabaseUpdateMessages(value uint64) (err error) { 
    return instance.SetProperty("DatabaseUpdateMessages", (value))
}

// GetDatabaseUpdateMessages gets the value of DatabaseUpdateMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyDatabaseUpdateMessages()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DatabaseUpdateMessages")
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

// SetDatabaseUpdateMessagesPersec sets the value of DatabaseUpdateMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyDatabaseUpdateMessagesPersec(value uint64) (err error) { 
    return instance.SetProperty("DatabaseUpdateMessagesPersec", (value))
}

// GetDatabaseUpdateMessagesPersec gets the value of DatabaseUpdateMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyDatabaseUpdateMessagesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DatabaseUpdateMessagesPersec")
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

// SetMessagesExecutionQueueLength sets the value of MessagesExecutionQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyMessagesExecutionQueueLength(value uint64) (err error) { 
    return instance.SetProperty("MessagesExecutionQueueLength", (value))
}

// GetMessagesExecutionQueueLength gets the value of MessagesExecutionQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyMessagesExecutionQueueLength()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MessagesExecutionQueueLength")
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

// SetMessagesQueueLength sets the value of MessagesQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyMessagesQueueLength(value uint64) (err error) { 
    return instance.SetProperty("MessagesQueueLength", (value))
}

// GetMessagesQueueLength gets the value of MessagesQueueLength for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyMessagesQueueLength()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MessagesQueueLength")
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

// SetUpdateMessages sets the value of UpdateMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyUpdateMessages(value uint64) (err error) { 
    return instance.SetProperty("UpdateMessages", (value))
}

// GetUpdateMessages gets the value of UpdateMessages for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyUpdateMessages()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UpdateMessages")
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

// SetUpdateMessagesPersec sets the value of UpdateMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) SetPropertyUpdateMessagesPersec(value uint64) (err error) { 
    return instance.SetProperty("UpdateMessagesPersec", (value))
}

// GetUpdateMessagesPersec gets the value of UpdateMessagesPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGlobalUpdateManagerMessages) GetPropertyUpdateMessagesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UpdateMessagesPersec")
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

