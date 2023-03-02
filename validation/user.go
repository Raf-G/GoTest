package validation

import (
	"example.com/m/v2/domain"
	"github.com/pkg/errors"
)

func UserValidation(u domain.User) error {
	if u.Login == "" {
		return errors.Wrapf(domain.ErrNoLogin, "user %+v validation failed", u)
	}

	if u.Surname == "" {
		return errors.Wrapf(domain.ErrNoSurname, "user %+v validation failed", u)
	}

	return nil
}
