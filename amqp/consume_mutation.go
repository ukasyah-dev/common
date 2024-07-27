package amqp

import (
	"context"

	"github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/model"
	"gorm.io/gorm"
)

func ConsumeMutation[T any](ctx context.Context, db *gorm.DB, queue string, consumerGroupID string, m *T) error {
	return Consume(ctx, queue, consumerGroupID, func(ctx context.Context, mutation *model.Mutation[*T]) (*model.Empty, error) {
		log.Debugf("Got a mutation: queue=%s type=%d", queue, mutation.Type)

		if mutation.Type == constant.MutationCreated {
			return createMutationData(ctx, db, mutation.Data)
		} else if mutation.Type == constant.MutationUpdated {
			return updateMutationData(ctx, db, mutation.Data)
		} else if mutation.Type == constant.MutationDeleted {
			return deleteMutationData(ctx, db, mutation.Data)
		}

		return &model.Empty{}, nil
	})
}

func createMutationData[T any](ctx context.Context, db *gorm.DB, req *T) (*model.Empty, error) {
	if err := db.WithContext(ctx).Create(req).Error; err != nil {
		log.Errorf("Failed to create mutation data: %s", err)
		return nil, err
	}

	return &model.Empty{}, nil
}

func updateMutationData[T any](ctx context.Context, db *gorm.DB, req *T) (*model.Empty, error) {
	if err := db.WithContext(ctx).Save(req).Error; err != nil {
		log.Errorf("Failed to update mutation data: %s", err)
		return nil, err
	}

	return &model.Empty{}, nil
}

func deleteMutationData[T any](ctx context.Context, db *gorm.DB, req *T) (*model.Empty, error) {
	if err := db.WithContext(ctx).Delete(req).Error; err != nil {
		log.Errorf("Failed to delete user: %s", err)
		return nil, err
	}

	return &model.Empty{}, nil
}
