package ports

import (
	"context"
	"sophon-core/domain"
)

type AIProvider interface {
	Chat(ctx context.Context, messages []domain.ConvoMessage) (domain.ConvoMessage, error)
	Stream(ctx context.Context, messages []domain.ConvoMessage) (<-chan domain.ConvoMessage, error)
}

type PlanRepository interface {
	Create(ctx context.Context, plan *domain.Plan) error
	Get(ctx context.Context, id string) (*domain.Plan, error)
	List(ctx context.Context, projectId string) ([]*domain.Plan, error)
	Update(ctx context.Context, plan *domain.Plan) error
	Delete(ctx context.Context, id string) error
}

type FileSystem interface {
	ReadFile(path string) (string, error)
	WriteFile(path string, content string) error
	ListFiles(path string) ([]string, error)
}

type VCS interface {
	Commit(message string) error
	Diff() (string, error)
	CurrentBranch() (string, error)
}
