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

func AddNote(base_ctx context.Context, db *gorm.DB, req *models.BaseNoteRequest, username string) error {

	ctx, cancel := context.WithTimeout(base_ctx, 5*time.Second)
	defer cancel()
	note := models.Note{}

	note.Username = username
	note.Title = req.Title
	note.Description = req.Description

	if err := db.WithContext(ctx).Create(&note).Error; err != nil {
		log.Println("Error: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return noteserrors.ErrTimedOut
		}
		return noteserrors.ErrInternal
	}

	return nil
}
