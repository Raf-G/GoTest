package validation

import (
	"example.com/m/v2/domain"
	"github.com/pkg/errors"
)

func BasketProductValidation(b domain.BasketProduct) error {
	if b.BasketID == 0 {
		return errors.Wrapf(domain.ErrBasketProductNoBasketID, "basket product %+v validation failed", b)
	}

	if b.ProductID == 0 {
		return errors.Wrapf(domain.ErrBasketProductNoProductID, "basket product %+v validation failed", b)
	}

	if b.Count == 0 {
		return errors.Wrapf(domain.ErrBasketProductNoCount, "basket product %+v validation failed", b)
	}

	return nil
}
