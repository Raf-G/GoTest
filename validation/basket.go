package validation

import (
	"example.com/m/v2/domain"
	"github.com/pkg/errors"
)

func BasketProductValidation(item domain.BasketProduct) error {
	if item.BasketID == 0 {
		return errors.Wrapf(domain.ErrBasketProductNoBasketID, "[validation] basket product %+v validation failed", item)
	}

	if item.ProductID == 0 {
		return errors.Wrapf(domain.ErrBasketProductNoProductID, "[validation] basket product %+v validation failed", item)
	}

	if item.Count == 0 {
		return errors.Wrapf(domain.ErrBasketProductNoCount, "[validation] basket product %+v validation failed", item)
	}

	return nil
}
