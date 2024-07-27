package paginator

import (
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/model"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, req *model.PaginationRequest) ([]*T, *model.PaginationResponse, error) {
	cursor := paginator.Cursor{}
	keys := []string{"CreatedAt", "UpdatedAt"}
	limit := 10
	order := paginator.ASC

	if req.Sort.Order == "desc" {
		order = paginator.DESC
	}

	if req.Limit > 0 {
		if req.Limit > 100 {
			req.Limit = 100
		} else {
			limit = req.Limit
		}
	}

	if len(req.Keys) > 0 {
		keys = req.Keys
	}

	pgn := paginator.New(
		&paginator.Config{
			Keys:  keys,
			Limit: limit,
			Order: order,
		},
	)

	if req.Cursor.After != "" {
		pgn.SetAfterCursor(req.Cursor.After)
	}

	if req.Cursor.Before != "" {
		pgn.SetBeforeCursor(req.Cursor.Before)
	}

	db = db.Session(&gorm.Session{})

	res := []*T{}

	resDB, cursor, err := pgn.Paginate(db, &res)

	// Handle paginator error
	if err != nil {
		return nil, nil, errors.InvalidArgument(err.Error())
	}

	// Handle db error
	if resDB.Error != nil {
		log.Errorf("Failed to list with pagination: %s", resDB.Error)
		return nil, nil, errors.Internal()
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		log.Errorf("Failed to get total data: %s", err)
		return nil, nil, errors.Internal()
	}

	pagination := &model.PaginationResponse{}

	pagination.Total = total

	if cursor.After != nil {
		pagination.Cursor.After = *cursor.After
	}

	if cursor.Before != nil {
		pagination.Cursor.Before = *cursor.Before
	}

	return res, pagination, nil
}
