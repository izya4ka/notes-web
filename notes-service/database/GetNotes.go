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

func GetNotes(base_ctx context.Context, db *gorm.DB, username string, amount int) ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(base_ctx, 30*time.Second)
	defer cancel()
	notes := make([]models.Note, amount)
	if err := db.WithContext(ctx).Model(&models.Note{}).Where("username = ?", username).Limit(amount).Find(&notes).Error; err != nil {
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, noteserrors.ErrTimedOut
		}
		return nil, noteserrors.ErrInternal
	}

	return notes, nil
}
