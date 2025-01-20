package util

import (
	"github.com/izya4ka/notes-web/user-service/models"
	"github.com/izya4ka/notes-web/user-service/myerrors"
)

func CheckRegLogReq(req *models.LogPassReq) error {
	if len(req.Username) < 5 {
		return myerrors.ErrStringLen{
			StringName:   "username",
			MustSize:     5,
			LessOrBigger: "bigger",
		}
	}

	if len(req.Username) > 12 {
		return myerrors.ErrStringLen{
			StringName:   "username",
			MustSize:     12,
			LessOrBigger: "less",
		}
	}

	if len(req.Password) > 18 {
		return myerrors.ErrStringLen{
			StringName:   "password",
			MustSize:     18,
			LessOrBigger: "less",
		}
	}

	if len(req.Password) < 6 {
		return myerrors.ErrStringLen{
			StringName:   "password",
			MustSize:     6,
			LessOrBigger: "bigger",
		}
	}

	return nil
}
