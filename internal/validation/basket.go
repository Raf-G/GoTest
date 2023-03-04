package validation

import (
	domain2 "example.com/m/v2/internal/domain"
	"github.com/pkg/errors"
)

func BasketProductValidation(b domain2.BasketProduct) error {
	if b.BasketID == 0 {
		return errors.Wrapf(domain2.ErrBasketProductNoBasketID, "basket product %+v validation failed", b)
	}

	if b.ProductID == 0 {
		return errors.Wrapf(domain2.ErrBasketProductNoProductID, "basket product %+v validation failed", b)
	}

	if b.Count == 0 {
		return errors.Wrapf(domain2.ErrBasketProductNoCount, "basket product %+v validation failed", b)
	}

	return nil
}
