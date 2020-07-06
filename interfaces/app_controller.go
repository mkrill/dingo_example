package interfaces

import (
	"github.com/mkrill/dingo_example/domain"

	"log"
)

type AppController struct {
	productAPI domain.ProductAPI
}

func (ac *AppController) Inject(productAPI domain.ProductAPI) {
	ac.productAPI = productAPI
}

// New initializes the productAPI struct of the controller and is not need with Dingo
func (ac *AppController) New(productAPI domain.ProductAPI) {
	ac.productAPI = productAPI
}

func (ac *AppController) PrintProducts() {
	productList, err := ac.productAPI.ProductList()
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range productList {
		log.Print(product.String())
	}
}
