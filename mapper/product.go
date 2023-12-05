package mapper

import (
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/labd/commercetools-go-sdk/platform"
)

type ProductMapperPort interface {
	MapToProduct(pp platform.ProductProjection) *model.Product
}

type ProductMapper struct{}

func NewProductMapper() *ProductMapper {
	return &ProductMapper{}
}

func (pm *ProductMapper) MapToProduct(productProjection platform.ProductProjection) *model.Product {
	return &model.Product{
		ID:          productProjection.ID,
		Name:        productProjection.Name["de"],
		Description: (*productProjection.Description)["de"],
		Sku:         *productProjection.MasterVariant.Sku,
		Attributes:  mapToAttributes(productProjection.MasterVariant.Attributes),
	}
}

func mapToAttributes(attributes []platform.Attribute) []model.Attribute {
	result := []model.Attribute{}
	for _, attribute := range attributes {
		result = append(result, model.Attribute{Name: attribute.Name, Value: attribute.Value})
	}

	return result
}
