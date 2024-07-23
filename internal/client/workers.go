package client

import (
	"comet/utils"
	"context"
	"sync"
	"time"
)

func (m *MapsConsumerGroup) layerWorker(
	ctx context.Context,
	cancel context.CancelFunc,
	wg *sync.WaitGroup,
	workerID int,
) {
	defer wg.Done()
	defer cancel()

	for {
		message, err := m.Client.C.ReadMessage(5 * time.Millisecond)
		if err != nil {
			continue
		}

		m.log.Info("[Consumer workers]",
			"worker", workerID,
			"topic", *message.TopicPartition.Topic,
			"partition", message.TopicPartition.Partition,
			"offset", message.TopicPartition.Offset.String(),
			"key", message.Key,
		)

		switch message.Key[0] {
		/************* Layer *************/
		case utils.AddLayerRequest:
			m.HandleAddLayerRequest(ctx, message)
			continue
		case utils.LayerRequest:
			m.HandleLayerRequest(ctx, message)
			continue
		case utils.EditLayerRequest:
			m.HandleEditLayerRequest(ctx, message)
			continue
		case utils.DeleteLayerRequest:
			m.HandleDeleteLayerRequest(ctx, message)
			continue
		case utils.LayersRequest:
			m.HandleLayersRequest(ctx, message)
			continue

		/************* Group *************/
		case utils.AddGroupRequest:
			m.HandleAddGroupRequest(ctx, message)
			continue
		case utils.GroupRequest:
			m.HandleGroupRequest(ctx, message)
			continue
		case utils.EditGroupRequest:
			m.HandleEditGroupRequest(ctx, message)
			continue
		case utils.DeleteGroupRequest:
			m.HandleDeleteGroupRequest(ctx, message)
			continue
		case utils.GroupsRequest:
			m.HandleGroupsRequest(ctx, message)
			continue

		/************* Style *************/
		case utils.AddStyleRequest:
			m.HandleAddStyleRequest(ctx, message)
			continue
		case utils.StyleRequest:
			m.HandleStyleRequest(ctx, message)
			continue
		case utils.EditStyleRequest:
			m.HandleEditStyleRequest(ctx, message)
			continue
		case utils.DeleteStyleRequest:
			m.HandleDeleteStyleRequest(ctx, message)
			continue
		case utils.StylesRequest:
			m.HandleStylesRequest(ctx, message)
			continue
		case utils.StylesPaginationRequest:
			m.HandleStylesPaginationRequest(ctx, message)
			continue

		/************* GroupLayerRelation *************/
		case utils.AddGroupLayerRelationRequest:
			m.HandleAddGroupLayerRelationRequest(ctx, message)
			continue
		case utils.DeleteGroupLayerRelationRequest:
			m.HandleDeleteGroupLayerRelationRequest(ctx, message)
			continue
		case utils.GroupLayerRelationsRequest:
			m.HandleGroupLayerRelationsRequest(ctx, message)
			continue
		case utils.LayerRelationGroupsRequest:
			m.HandleLayerRelationGroupsRequest(ctx, message)
			continue
		case utils.GroupRelationLayersRequest:
			m.HandleGroupRelationLayersRequest(ctx, message)
			continue
		case utils.GroupLayerOrderUpRequest:
			m.HandleGroupLayerOrderUpRequest(ctx, message)
			continue
		case utils.GroupLayerOrderDownRequest:
			m.HandleGroupLayerOrderDownRequest(ctx, message)
			continue

		/************* MapGroupRelation *************/
		case utils.AddMapGroupRelationRequest:
			m.HandleAddMapGroupRelationRequest(ctx, message)
			continue
		case utils.DeleteMapGroupRelationRequest:
			m.HandleDeleteMapGroupRelationRequest(ctx, message)
			continue
		case utils.MapGroupRelationsRequest:
			m.HandleMapGroupRelationsRequest(ctx, message)
			continue
		case utils.MapRelationGroupsRequest:
			m.HandleMapRelationGroupsRequest(ctx, message)
			continue
		case utils.GroupRelationMapsRequest:
			m.HandleGroupRelationMapsRequest(ctx, message)
			continue
		case utils.MapGroupOrderUpRequest:
			m.HandleMapGroupOrderUpRequest(ctx, message)
			continue
		case utils.MapGroupOrderDownResponse:
			m.HandleMapGroupOrderDownRequest(ctx, message)
			continue

		}

		/************* LayerStyleRelation *************/
		if next := m.LayerStyleRelationSwitcher(ctx, message); next {
			continue
		}
		/************* StyledMap *************/
		if next := m.StyledMapSwitcher(ctx, message); next {
			continue
		}
		/************* Patterns *************/
		if next := m.PatternSwitcher(ctx, message); next {
			continue
		}
		/************* Map *************/
		if next := m.MapSwitcher(ctx, message); next {
			continue
		}
		/************* Tables *************/
		if next := m.TableSwitcher(ctx, message); next {
			continue
		}

	}
}
