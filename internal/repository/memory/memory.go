package memory

import (
	"context"
	"log/slog"
	"sync"
)

type MemoryRepository struct {
	SubscribedUsers map[int64]struct{}
	sync.RWMutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		SubscribedUsers: make(map[int64]struct{}),
	}
}

func (mr *MemoryRepository) GetAll(ctx context.Context) ([]int64, error) {
	slog.Debug("MemoryRepo GetAll")
	ids := []int64{}
	mr.RLock()
	for key := range mr.SubscribedUsers {
		ids = append(ids, key)
	}
	mr.RUnlock()
	return ids, nil
}

func (mr *MemoryRepository) Add(ctx context.Context, id int64) error {
	slog.Debug("MemoryRepo Add", "id", id)
	mr.Lock()
	mr.SubscribedUsers[id] = struct{}{}
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Del(ctx context.Context, id int64) error {
	slog.Debug("MemoryRepo delete", "id", id)
	mr.Lock()
	delete(mr.SubscribedUsers, id)
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Close() {
}
