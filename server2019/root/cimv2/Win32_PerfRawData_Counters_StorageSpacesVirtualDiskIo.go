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

// Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo struct
type Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo struct { 
	*Win32_PerfRawData

	// 
	VirtualDiskFlushLatencyms uint32

	// 
	VirtualDiskFlushLatencyms_Base uint32

	// 
	VirtualDiskReadBytesPersec uint64

	// 
	VirtualDiskReadLatencyms uint32

	// 
	VirtualDiskReadLatencyms_Base uint32

	// 
	VirtualDiskWriteBytesPersec uint64

	// 
	VirtualDiskWriteLatencyms uint32

	// 
	VirtualDiskWriteLatencyms_Base uint32
}

	func NewWin32_PerfRawData_Counters_StorageSpacesVirtualDiskIoEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_Counters_StorageSpacesVirtualDiskIoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetVirtualDiskFlushLatencyms sets the value of VirtualDiskFlushLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskFlushLatencyms(value uint32) (err error) { 
    return instance.SetProperty("VirtualDiskFlushLatencyms", (value))
}

// GetVirtualDiskFlushLatencyms gets the value of VirtualDiskFlushLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskFlushLatencyms()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskFlushLatencyms")
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

// SetVirtualDiskFlushLatencyms_Base sets the value of VirtualDiskFlushLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskFlushLatencyms_Base(value uint32) (err error) { 
    return instance.SetProperty("VirtualDiskFlushLatencyms_Base", (value))
}

// GetVirtualDiskFlushLatencyms_Base gets the value of VirtualDiskFlushLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskFlushLatencyms_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskFlushLatencyms_Base")
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

// SetVirtualDiskReadBytesPersec sets the value of VirtualDiskReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskReadBytesPersec(value uint64) (err error) { 
    return instance.SetProperty("VirtualDiskReadBytesPersec", (value))
}

// GetVirtualDiskReadBytesPersec gets the value of VirtualDiskReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskReadBytesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskReadBytesPersec")
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

// SetVirtualDiskReadLatencyms sets the value of VirtualDiskReadLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskReadLatencyms(value uint32) (err error) { 
    return instance.SetProperty("VirtualDiskReadLatencyms", (value))
}

// GetVirtualDiskReadLatencyms gets the value of VirtualDiskReadLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskReadLatencyms()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskReadLatencyms")
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

// SetVirtualDiskReadLatencyms_Base sets the value of VirtualDiskReadLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskReadLatencyms_Base(value uint32) (err error) { 
    return instance.SetProperty("VirtualDiskReadLatencyms_Base", (value))
}

// GetVirtualDiskReadLatencyms_Base gets the value of VirtualDiskReadLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskReadLatencyms_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskReadLatencyms_Base")
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

// SetVirtualDiskWriteBytesPersec sets the value of VirtualDiskWriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskWriteBytesPersec(value uint64) (err error) { 
    return instance.SetProperty("VirtualDiskWriteBytesPersec", (value))
}

// GetVirtualDiskWriteBytesPersec gets the value of VirtualDiskWriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskWriteBytesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskWriteBytesPersec")
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

// SetVirtualDiskWriteLatencyms sets the value of VirtualDiskWriteLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskWriteLatencyms(value uint32) (err error) { 
    return instance.SetProperty("VirtualDiskWriteLatencyms", (value))
}

// GetVirtualDiskWriteLatencyms gets the value of VirtualDiskWriteLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskWriteLatencyms()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskWriteLatencyms")
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

// SetVirtualDiskWriteLatencyms_Base sets the value of VirtualDiskWriteLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskWriteLatencyms_Base(value uint32) (err error) { 
    return instance.SetProperty("VirtualDiskWriteLatencyms_Base", (value))
}

// GetVirtualDiskWriteLatencyms_Base gets the value of VirtualDiskWriteLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskWriteLatencyms_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VirtualDiskWriteLatencyms_Base")
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

