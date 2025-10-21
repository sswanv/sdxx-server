package player

import (
	"context"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
)

func (s *Server) GetAttrs(ctx context.Context, req *playerv1.GetAttrsReq) (*playerv1.GetAttrsResp, error) {
	player, err := s.playerModel.FindOneByPlayerId(ctx, req.PlayerId)
	if err != nil {
		log.Errorf("get attrs: find one by player id failed, err: %v", err)
		return nil, errors.NewError(commonv1.InternalError)
	}

	tbLevel := s.table.TbLevel.Load()
	tbProperty := s.table.TbProperty.Load()

	// 所有属性
	attrs := make(map[uint32]*commonv1.Attr)
	for _, property := range tbProperty.GetDataMap() {
		attrs[uint32(property.PropertyId)] = &commonv1.Attr{
			Id:    uint32(property.PropertyId),
			Value: 0,
		}
	}

	// 基础属性
	baseAttrs := make(map[uint32]*commonv1.Attr)
	for _, property := range tbLevel.Get(int32(player.Level)).Property {
		propertyId := uint32(property.Property)
		propertyValue := uint64(property.Value)

		baseAttrs[propertyId] = &commonv1.Attr{
			Id:    propertyId,
			Value: propertyValue,
		}
		if attr, ok := attrs[propertyId]; ok {
			attr.Value += propertyValue
		}
	}

	// 装备属性
	// Todo yhaha
	var (
		equipmentAttrs = make(map[uint32]*commonv1.Attr)
		equipments     []*commonv1.Equipment
	)

	return &playerv1.GetAttrsResp{
		Attrs:          attrs,
		BaseAttrs:      baseAttrs,
		EquipmentAttrs: equipmentAttrs,
		Equipments:     equipments,
	}, nil
}
