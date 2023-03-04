package validation

import (
	domain2 "example.com/m/v2/internal/domain"
	"github.com/pkg/errors"
)

func UserValidation(u domain2.User) error {
	if u.Login == "" {
		return errors.Wrapf(domain2.ErrNoLogin, "user %+v validation failed", u)
	}

	if u.Surname == "" {
		return errors.Wrapf(domain2.ErrNoSurname, "user %+v validation failed", u)
	}

	return nil
}
