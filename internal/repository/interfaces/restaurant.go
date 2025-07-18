package interfaces

import (
	"context"
	"github.com/rezalaal/coral/internal/models"
)

type RestaurantRepository interface {
	Create(ctx context.Context, restaurant *models.Restaurant) error
	GetByID(ctx context.Context, id int64) (*models.Restaurant, error)
	List(ctx context.Context) ([]*models.Restaurant, error)
}
