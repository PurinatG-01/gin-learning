package repository

import (
	model "gin-learning/models"
)

type EventRepository interface {
	All() (*[]model.Event, error)
	Create(event *model.Event) (bool, error)
	// Load(ctx context.Context, id string) (*model.Event, error)
	// Create(ctx context.Context, user *model.Event) (int64, error)
	// Update(ctx context.Context, user *model.Event) (int64, error)
	// Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	// Delete(ctx context.Context, id string) (int64, error)
}
