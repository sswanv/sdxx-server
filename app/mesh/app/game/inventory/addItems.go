package inventory

import (
	"context"
	cfg "sdxx/server/internal/config/gen"
	"sdxx/server/internal/core/msgx"
	"sdxx/server/internal/dao/mysql/game/model"
	"sdxx/server/internal/utils/uuidx"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"

	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/session"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Server) AddItems_addItemByType(ctx context.Context, tx *gorm.DB, playerId uint64, itemCfg *cfg.Item, count uint64) error {
	switch itemCfg.ItemType {
	case cfg.E_ItemType_Equipment:
		// 装备：添加到装备表
		return s.AddItems_addEquipment(ctx, tx, playerId, itemCfg, count)
	case cfg.E_ItemType_Currency, cfg.E_ItemType_Exp, cfg.E_ItemType_Integral:
		// 资产：添加到资产表
		return s.AddItems_addAsset(ctx, tx, playerId, itemCfg, count)

	default:
		// 普通道具：添加到道具表
		return s.AddItems_addItem(ctx, tx, playerId, itemCfg, count)
	}
}

func (s *Server) AddItems_addEquipment(ctx context.Context, tx *gorm.DB, playerId uint64, itemCfg *cfg.Item, count uint64) error {
	return nil
}

func (s *Server) AddItems_addAsset(ctx context.Context, tx *gorm.DB, playerId uint64, itemCfg *cfg.Item, count uint64) error {
	asset := model.PlayerAssets{
		PlayerId: playerId,
		ItemId:   uint64(itemCfg.Id),
		Count:    count,
		Type:     uint64(itemCfg.ItemType),
		SubType:  uint64(itemCfg.ItemSubclass),
	}

	return tx.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "player_id"}, {Name: "item_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"count": gorm.Expr("count + VALUES(count)"),
		}),
	}).Create(&asset).Error
}

func (s *Server) AddItems_addItem(ctx context.Context, tx *gorm.DB, playerId uint64, itemCfg *cfg.Item, count uint64) error {
	if itemCfg.NumberLimit > 1 {
		// 可叠加道具
		item := model.PlayerItem{
			PlayerId:    playerId,
			ItemId:      uint64(itemCfg.Id),
			Count:       count,
			Type:        uint64(itemCfg.ItemType),
			SubType:     uint64(itemCfg.ItemSubclass),
			IsStackable: 1,
			InstanceId:  uuidx.UUID(),
		}

		return tx.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "player_id"}, {Name: "item_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"count": gorm.Expr("count + VALUES(count)"),
			}),
		}).Create(&item).Error
	} else {
		// 不叠加道具
		for range count {
			item := model.PlayerItem{
				PlayerId:    playerId,
				ItemId:      uint64(itemCfg.Id),
				Count:       1,
				Type:        uint64(itemCfg.ItemType),
				SubType:     uint64(itemCfg.ItemSubclass),
				IsStackable: 0,
				InstanceId:  uuidx.UUID(),
			}

			err := s.playerItemModel.Insert(ctx, tx, &item)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (s *Server) AddItems(ctx context.Context, req *inventoryv1.AddItemsReq) (*inventoryv1.AddItemsResp, error) {
	if len(req.Items) == 0 {
		return &inventoryv1.AddItemsResp{}, nil
	}

	tbItem := s.table.TbItem.Load()
	var (
		items        []*commonv1.ItemStack
		assetChanges []*commonv1.ItemStack
	)

	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range req.Items {
			itemCfg := tbItem.Get(int32(item.ItemId))
			if itemCfg == nil {
				continue
			}

			// 如果是资产类型，需要记录变动前后的数量
			if itemCfg.ItemType == cfg.E_ItemType_Currency ||
				itemCfg.ItemType == cfg.E_ItemType_Exp ||
				itemCfg.ItemType == cfg.E_ItemType_Integral {
				// 获取变动前的数量
				beforeCount := uint64(0)
				existingAsset, err := s.playerAssetsModel.FindOneByPlayerIdItemId(ctx, req.PlayerId, item.ItemId)
				if err == nil && existingAsset != nil {
					beforeCount = existingAsset.Count
				}
				// 根据道具类型分发到不同的表
				err = s.AddItems_addItemByType(ctx, tx, req.PlayerId, itemCfg, item.Count)
				if err != nil {
					return err
				}
				// 记录资产变动
				assetChanges = append(assetChanges, &commonv1.ItemStack{
					ItemId: item.ItemId,
					Count:  beforeCount + item.Count,
				})

			} else {
				// 根据道具类型分发到不同的表
				err := s.AddItems_addItemByType(ctx, tx, req.PlayerId, itemCfg, item.Count)
				if err != nil {
					return err
				}
			}
			items = append(items, &commonv1.ItemStack{
				ItemId: item.ItemId,
				Count:  item.Count,
			})
		}
		return nil
	})
	if err != nil {
		log.Errorf("add items err: %v", err)
		return nil, errors.NewError(commonv1.InternalError)
	}

	// 推送道具变动通知
	err = s.proxy.Push(ctx, &cluster.PushArgs{
		Kind:   session.User,
		Target: int64(req.PlayerId),
		Message: msgx.Cluster(
			int32(commonv1.Route_ROUTE_PLAYER_ITEM_CHANGE_NOTIFY),
			&commonv1.RoutePlayerItemChangeNotify{Items: items},
		),
	})
	if err != nil {
		log.Errorf("push items err: %v", err)
	}

	// 推送资产变动通知
	if len(assetChanges) > 0 {
		err = s.proxy.Push(ctx, &cluster.PushArgs{
			Kind:   session.User,
			Target: int64(req.PlayerId),
			Message: msgx.Cluster(
				int32(commonv1.Route_ROUTE_PLAYER_ASSET_CHANGE_NOTIFY),
				&commonv1.RoutePlayerAssetChangeNotify{Assets: assetChanges},
			),
		})
		if err != nil {
			log.Errorf("push asset changes err: %v", err)
		}
	}

	return &inventoryv1.AddItemsResp{}, nil
}
