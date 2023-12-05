package mapper

import (
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/labd/commercetools-go-sdk/platform"
)

type IProductMapper interface {
	MapToProduct(pp platform.ProductProjection) *model.Product
}

type ProductMapper struct {}

func NewProductMapper() *ProductMapper {
	return &ProductMapper{}
}

func (pm *ProductMapper) MapToProduct(pp platform.ProductProjection) *model.Product {
	return &model.Product{
		ID:          pp.ID,
		Name:        pp.Name["de"],
		Description: (*pp.Description)["de"],
		Sku:         *pp.MasterVariant.Sku,
		Attributes: mapToAttributes(pp.MasterVariant.Attributes),
	}
}

func mapToAttributes(attributes []platform.Attribute) []model.Attribute {
	result := []model.Attribute{}
	for _, attribute := range attributes {
		result = append(result, model.Attribute{Name: attribute.Name, Value: attribute.Value})
	}
	return result
}
