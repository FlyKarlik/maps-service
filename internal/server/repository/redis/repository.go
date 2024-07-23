package redis

import (
	"github.com/redis/go-redis/v9"
	"protos/maps"
)

type StyledMapsRepository interface {
	Set(*maps.StyledMap) error
	Get(id string) (*maps.StyledMap, error)
}

type Repository struct {
	StyledMapsRepository
}

func NewRepository(StyledMapsClient *redis.Client) *Repository {
	return &Repository{
		StyledMapsRepository: NewStyledMapsRepository(StyledMapsClient),
	}
}
