package repository

import (
	model "gin-learning/models"
)

type EventRepository interface {
	All() (*[]model.Event, error)
	Create(event *model.Event) (bool, error)
	Get(int) (model.Event, error)
	Update(*model.Event) (bool, error)
	Delete(*model.Event) (bool, error)
	// Load(ctx context.Context, id string) (*model.Event, error)
	// Patch(ctx context.Context, user map[string]interface{}) (int64, error)
}
