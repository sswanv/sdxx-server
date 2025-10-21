package types

import (
	"fmt"
	"strings"
)

type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "unknown" // 未知设备
	DeviceTypeIOS     DeviceType = "ios"     // iOS设备
	DeviceTypeAndroid DeviceType = "android" // Android设备
	DeviceTypeWeb     DeviceType = "web"     // Web设备
)

func (d DeviceType) String() string {
	return string(d)
}

func (d DeviceType) IsValid() bool {
	switch d {
	case DeviceTypeIOS, DeviceTypeAndroid, DeviceTypeWeb:
		return true
	default:
		return false
	}
}

func ParseDeviceType(s string) (DeviceType, error) {
	if s == "" {
		return DeviceTypeUnknown, nil
	}

	dt := DeviceType(strings.ToLower(s))
	if !dt.IsValid() {
		return DeviceTypeUnknown, fmt.Errorf("invalid device type: %s", s)
	}
	return dt, nil
}

func MustParseDeviceType(s string) DeviceType {
	dt, err := ParseDeviceType(s)
	if err != nil {
		return DeviceTypeUnknown
	}
	return dt
}

func AllDeviceTypes() []DeviceType {
	return []DeviceType{
		DeviceTypeIOS,
		DeviceTypeAndroid,
		DeviceTypeWeb,
	}
}
