package model

import (
	"context"

	"gorm.io/gorm"
)

var _ PlayerAssetsModel = (*customPlayerAssetsModel)(nil)

type (
	// PlayerAssetsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlayerAssetsModel.
	PlayerAssetsModel interface {
		playerAssetsModel
		customPlayerAssetsLogicModel
	}

	customPlayerAssetsModel struct {
		*defaultPlayerAssetsModel
	}

	customPlayerAssetsLogicModel interface {
		FindByPlayerId(ctx context.Context, playerId uint64) ([]*PlayerAssets, error)
		FindByPlayerIdAndItemId(ctx context.Context, playerId uint64, itemId []uint64) ([]*PlayerAssets, error)
		BatchUpsert(ctx context.Context, tx *gorm.DB, data []PlayerAssets) error
	}
)

// NewPlayerAssetsModel returns a model for the database table.
func NewPlayerAssetsModel(conn *gorm.DB) PlayerAssetsModel {
	return &customPlayerAssetsModel{
		defaultPlayerAssetsModel: newPlayerAssetsModel(conn),
	}
}

func (m *customPlayerAssetsModel) FindByPlayerId(ctx context.Context, playerId uint64) ([]*PlayerAssets, error) {
	var playerAssets []*PlayerAssets
	err := m.conn.WithContext(ctx).
		Model(&PlayerAssets{}).
		Where("player_id = ?", playerId).Find(&playerAssets).Error
	if err != nil {
		return nil, err
	}
	return playerAssets, nil
}

func (m *customPlayerAssetsModel) FindByPlayerIdAndItemId(ctx context.Context, playerId uint64, itemId []uint64) ([]*PlayerAssets, error) {
	var playerAssets []*PlayerAssets
	err := m.conn.WithContext(ctx).
		Model(&PlayerAssets{}).
		Where("player_id = ? and item_id in (?)", playerId, itemId).
		Find(&playerAssets).Error
	if err != nil {
		return nil, err
	}
	return playerAssets, nil
}

func (m *customPlayerAssetsModel) BatchUpsert(ctx context.Context, tx *gorm.DB, data []PlayerAssets) error {
	db := m.conn
	if tx != nil {
		db = tx
	}

	// 使用批量更新操作，因为我们已经查询了现有数据并进行了累加
	// 这样可以确保数据的一致性
	for _, asset := range data {
		err := db.WithContext(ctx).Model(&PlayerAssets{}).
			Where("player_id = ? AND item_id = ?", asset.PlayerId, asset.ItemId).
			Updates(map[string]interface{}{
				"count":      asset.Count,
				"updated_at": asset.UpdatedAt,
			}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
