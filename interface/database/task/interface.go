package task

import (
	"context"

	"github.com/Hajime3778/go-clean-arch/domain"
)

// TaskRepository
type TaskRepository interface {
	FindByUserID(ctx context.Context, userID int64, limit int64, offset int64) ([]domain.Task, error)
	GetByID(ctx context.Context, id int64) (domain.Task, error)
	Create(ctx context.Context, task domain.Task) (int64, error)
	Update(ctx context.Context, task domain.Task) error
	Delete(ctx context.Context, id int64) error
}
