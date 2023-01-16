package validation

import (
	"example.com/m/v2/domain"
	"github.com/pkg/errors"
)

func UserValidation(item domain.User) error {
	if item.Login == "" {
		return errors.Wrapf(domain.ErrNoLogin, "[validation] item %+v validation failed", item)
	}

	if item.Surname == "" {
		return errors.Wrapf(domain.ErrNoSurname, "[validation] item %+v validation failed", item)
	}

	return nil
}
