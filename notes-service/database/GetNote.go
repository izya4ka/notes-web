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

func GetNote(base_ctx context.Context, db *gorm.DB, username string, id int) (models.Note, error) {
	var note models.Note
	note.ID = uint(id)

	ctx, cancel := context.WithTimeout(base_ctx, 5*time.Second)
	defer cancel()

	if err := db.WithContext(ctx).Model(&note).Where("username = ?", username).Take(&note).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return note, noteserrors.ErrNotFound
		} else {
			log.Println("Error: ", err)
			if errors.Is(err, context.DeadlineExceeded) {
				{
					return note, noteserrors.ErrInternal
				}
			}
			return note, noteserrors.ErrTimedOut
		}
	}
	return note, nil
}
