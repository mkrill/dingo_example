// domain/product_api.go
package domain

type ProductAPI interface {
	ProductList() (ProductList, error)
}