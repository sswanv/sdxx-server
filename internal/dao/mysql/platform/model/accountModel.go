package model

import (
	"context"

	"gorm.io/gorm"
)

var _ AccountModel = (*customAccountModel)(nil)

type (
	// AccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountModel.
	AccountModel interface {
		accountModel
		customAccountLogicModel
	}

	customAccountModel struct {
		*defaultAccountModel
	}

	customAccountLogicModel interface {
		UpdateByAccountId(ctx context.Context, tx *gorm.DB, accountId int64, data *Account) error
		UpdateById(ctx context.Context, tx *gorm.DB, id int64, data *Account) error
	}
)

// NewAccountModel returns a model for the database table.
func NewAccountModel(conn *gorm.DB) AccountModel {
	return &customAccountModel{
		defaultAccountModel: newAccountModel(conn),
	}
}

func (m *customAccountModel) UpdateByAccountId(ctx context.Context, tx *gorm.DB, accountId int64, data *Account) error {
	db := m.conn
	if tx != nil {
		db = tx
	}

	err := db.WithContext(ctx).Where("account_id = ?", accountId).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *customAccountModel) UpdateById(ctx context.Context, tx *gorm.DB, id int64, data *Account) error {
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
