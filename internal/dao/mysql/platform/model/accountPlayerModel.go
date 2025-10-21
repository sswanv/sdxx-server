package model

import (
	"context"

	"gorm.io/gorm"
)

var _ AccountPlayerModel = (*customAccountPlayerModel)(nil)

type (
	// AccountPlayerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountPlayerModel.
	AccountPlayerModel interface {
		accountPlayerModel
		customAccountPlayerLogicModel
	}

	customAccountPlayerModel struct {
		*defaultAccountPlayerModel
	}

	customAccountPlayerLogicModel interface {
		FindByAccountId(ctx context.Context, accountId int64) ([]*AccountPlayer, error)
	}
)

// NewAccountPlayerModel returns a model for the database table.
func NewAccountPlayerModel(conn *gorm.DB) AccountPlayerModel {
	return &customAccountPlayerModel{
		defaultAccountPlayerModel: newAccountPlayerModel(conn),
	}
}

func (m *customAccountPlayerModel) FindByAccountId(ctx context.Context, accountId int64) ([]*AccountPlayer, error) {
	var resp []*AccountPlayer
	err := m.conn.WithContext(ctx).Model(&AccountPlayer{}).Where("`account_id` = ?", accountId).Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}
