package model

import (
	"gorm.io/gorm"
)

var _ PlayerModel = (*customPlayerModel)(nil)

type (
	// PlayerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlayerModel.
	PlayerModel interface {
		playerModel
		customPlayerLogicModel
	}

	customPlayerModel struct {
		*defaultPlayerModel
	}

	customPlayerLogicModel interface {
	}
)

// NewPlayerModel returns a model for the database table.
func NewPlayerModel(conn *gorm.DB) PlayerModel {
	return &customPlayerModel{
		defaultPlayerModel: newPlayerModel(conn),
	}
}
