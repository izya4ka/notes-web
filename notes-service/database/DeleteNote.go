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

func DeleteNote(base_ctx context.Context, db *gorm.DB, username string, id int) error {
	var note models.Note

	note.ID = uint(id)

	ctx, cancel := context.WithTimeout(base_ctx, 5*time.Second)
	defer cancel()

	res := db.WithContext(ctx).Model(&note).Where("username = ?", username).Delete(&note)
	if err := res.Error; err != nil {
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return noteserrors.ErrTimedOut
		}
		return noteserrors.ErrInternal
	}

	if res.RowsAffected == 0 {
		return noteserrors.ErrNotFound
	}

	return nil
}
