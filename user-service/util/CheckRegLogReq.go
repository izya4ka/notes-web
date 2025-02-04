package util

import (
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/usererrors"
)

// CheckRegLogReq validates the registration and login request parameters.
// It checks the length of the username and password fields based on predefined rules:
// - Username must be between 5 and 12 characters long.
// - Password must be between 6 and 18 characters long.
// If any of these conditions are not met, it returns an appropriate error.
//
// This function is essential for ensuring that user credentials meet the
// required security standards before being processed further.
func CheckRegLogReq(req *models.LogPassRequest) error {
	if len(req.Username) < 5 {
		return usererrors.ErrStringLen{
			StringName:    "username",
			MustSize:      5,
			LessOrGreater: "greater",
		}
	}

	if len(req.Username) > 12 {
		return usererrors.ErrStringLen{
			StringName:    "username",
			MustSize:      12,
			LessOrGreater: "less",
		}
	}

	if len(req.Password) > 18 {
		return usererrors.ErrStringLen{
			StringName:    "password",
			MustSize:      18,
			LessOrGreater: "less",
		}
	}

	if len(req.Password) < 6 {
		return usererrors.ErrStringLen{
			StringName:    "password",
			MustSize:      6,
			LessOrGreater: "greater",
		}
	}

	return nil
}
