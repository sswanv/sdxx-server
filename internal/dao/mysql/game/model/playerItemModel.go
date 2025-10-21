package model

import (
	"context"

	"gorm.io/gorm"
)

var _ PlayerItemModel = (*customPlayerItemModel)(nil)

type (
	// PlayerItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlayerItemModel.
	PlayerItemModel interface {
		playerItemModel
		customPlayerItemLogicModel
	}

	customPlayerItemModel struct {
		*defaultPlayerItemModel
	}

	customPlayerItemLogicModel interface {
		FindByPlayerId(ctx context.Context, playerId uint64) ([]*PlayerItem, error)
	}
)

// NewPlayerItemModel returns a model for the database table.
func NewPlayerItemModel(conn *gorm.DB) PlayerItemModel {
	return &customPlayerItemModel{
		defaultPlayerItemModel: newPlayerItemModel(conn),
	}
}

func (m *customPlayerItemModel) FindByPlayerId(ctx context.Context, playerId uint64) ([]*PlayerItem, error) {
	var playerItems []*PlayerItem
	err := m.conn.WithContext(ctx).
		Where("player_id = ?", playerId).
		Find(&playerItems).Error
	if err != nil {
		return nil, err
	}
	return playerItems, nil
}
