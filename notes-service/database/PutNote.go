package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/izya4ka/notes-web/notes-service/models"
	"github.com/izya4ka/notes-web/notes-service/noteserrors"
	"gorm.io/gorm"
)

func PutNote(base_ctx context.Context, db *gorm.DB, req *models.BaseNoteRequest, username string, id int) (models.Note, error) {
	var note models.Note

	note.ID = uint(id)
	note.Title = req.Title
	note.Description = req.Description

	ctx, cancel := context.WithTimeout(base_ctx, 5*time.Second)
	defer cancel()

	result := db.WithContext(ctx).Model(&note).Where("username = ?", username).Updates(&note)
	if err := result.Error; err != nil {
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return note, noteserrors.ErrTimedOut
		}
		return note, noteserrors.ErrInternal
	}

	if result.RowsAffected == 0 {
		return note, noteserrors.ErrNotFound
	}

	return note, nil
}
