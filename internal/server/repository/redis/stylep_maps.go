package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"protos/maps"
)

type StyledMapsRepositoryI struct {
	db *redis.Client
}

func NewStyledMapsRepository(db *redis.Client) *StyledMapsRepositoryI {
	return &StyledMapsRepositoryI{
		db: db,
	}
}

func (r *StyledMapsRepositoryI) Set(styledMap *maps.StyledMap) error {
	data, err := json.Marshal(styledMap)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	if err := r.db.Set(context.TODO(), styledMap.Id, string(data), 0).Err(); err != nil {
		return fmt.Errorf("r.db.Set: %w", err)
	}
	return nil
}
func (r *StyledMapsRepositoryI) Get(id string) (*maps.StyledMap, error) {
	res, err := r.db.Get(context.TODO(), id).Result()
	if err != nil {
		return &maps.StyledMap{}, fmt.Errorf("r.db.Get: %w", err)
	}

	var styledMap *maps.StyledMap
	if err := json.Unmarshal([]byte(res), &styledMap); err != nil {
		return &maps.StyledMap{}, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return styledMap, nil
}
