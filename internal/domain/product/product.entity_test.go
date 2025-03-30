package domain_test

import (
	"testing"

	domain "github.com/JeffSilva01/my-order-api/internal/domain/product"
	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	product, err := domain.NewProduct(
		"BACONMAN",
		"Pão brioche , suculenta carne artesanal de 100 gramas, creme cheese e super bacon crispy com nossa maionese da casa.",
		"https://example.com/images/product.jpg",
		10.99,
		nil,
	)

	require.NotNil(t, product)
	require.Nil(t, err)
}
