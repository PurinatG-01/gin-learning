package repository

import (
	model "gin-learning/models"
)

type EventRepository interface {
	All() (*[]model.Events, error)
	Create(event *model.Events) (bool, error)
	Get(int) (model.Events, error)
	Update(*model.Events) (bool, error)
	Delete(*model.Events) (bool, error)
	// Load(ctx context.Context, id string) (*model.Event, error)
	// Patch(ctx context.Context, user map[string]interface{}) (int64, error)
}
