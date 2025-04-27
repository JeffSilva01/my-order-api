package domain

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID              string   `json:"product_id" valid:"uuid,required" gorm:"type:uuid;primary_key"`
	Name            string   `json:"name" valid:"required,stringlength(3|500)"`
	Description     string   `json:"description" valid:"required,stringlength(3|500)"`
	UrlImage        string   `json:"url_image" valid:"url,required"`
	Price           float64  `json:"price" valid:"required"`
	DiscountPercent *float64 `json:"discount_percent" valid:"optional"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduct(name string, description string, urlImage string, price float64, discountPercent *float64) (*Product, error) {
	product := Product{
		ID:              uuid.NewV4().String(),
		Name:            name,
		Description:     description,
		UrlImage:        urlImage,
		Price:           price,
		DiscountPercent: discountPercent,
	}

	err := product.validate()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &product, nil
}

func (product *Product) validate() error {
	var errList []string

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		errList = append(errList, err.Error())
	}

	isPriceRange := govalidator.InRangeFloat64(product.Price, 0.01, 10000000)
	if !isPriceRange {
		errList = append(errList, "product price must be between $0.01 and $10,000,000.00")
	}

	if len(errList) > 0 {
		return errors.New(strings.Join(errList, "; "))
	}

	return nil
}

func (p *Product) GetFinalPrice() float64 {
	if p.DiscountPercent == nil || *p.DiscountPercent <= 0 {
		return p.Price
	}

	discountAmount := p.Price * (*p.DiscountPercent / 100)
	return p.Price - discountAmount
}

func (p *Product) HasDiscount() bool {
	return p.DiscountPercent != nil && *p.DiscountPercent > 0
}

func (p *Product) GetDiscountAmount() float64 {
	if !p.HasDiscount() {
		return 0
	}
	return p.Price * (*p.DiscountPercent / 100)
}
