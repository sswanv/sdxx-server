package model

import (
	"sdxx/server/internal/types"

	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type AccountAuthChannel = types.Channel

type AccountAuthStatus byte

const (
	AccountAuthStatusNone   AccountAuthStatus = 0 // None
	AccountAuthStatusNormal AccountAuthStatus = 1 // 正常
	AccountAuthStatusBanned AccountAuthStatus = 1 // 封禁
)

type AccountDeviceType = types.DeviceType

type AccountStatus byte

const (
	AccountStatusNone    AccountStatus = 0 // None
	AccountStatusNormal  AccountStatus = 1 // 正常
	AccountStatusBanned  AccountStatus = 1 // 封禁
	AccountStatusFrozen  AccountStatus = 2 // 冻结
	AccountStatusDeleted AccountStatus = 3 // 注销
)
