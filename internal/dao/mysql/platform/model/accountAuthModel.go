package model

import (
	"context"

	"gorm.io/gorm"
)

var _ AccountAuthModel = (*customAccountAuthModel)(nil)

type (
	// AccountAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountAuthModel.
	AccountAuthModel interface {
		accountAuthModel
		customAccountAuthLogicModel
	}

	customAccountAuthModel struct {
		*defaultAccountAuthModel
	}

	customAccountAuthLogicModel interface {
		UpdateById(ctx context.Context, tx *gorm.DB, id int64, data *AccountAuth) error
	}
)

// NewAccountAuthModel returns a model for the database table.
func NewAccountAuthModel(conn *gorm.DB) AccountAuthModel {
	return &customAccountAuthModel{
		defaultAccountAuthModel: newAccountAuthModel(conn),
	}
}

func (m *customAccountAuthModel) UpdateById(ctx context.Context, tx *gorm.DB, id int64, data *AccountAuth) error {
	db := m.conn
	if tx != nil {
		db = tx
	}

	err := db.WithContext(ctx).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}
