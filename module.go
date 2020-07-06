package main

import (
	"flamingo.me/dingo"

	"github.com/mkrill/dingo_example/domain"
	"github.com/mkrill/dingo_example/infrastructure"
)

const ProductApiUrl = "https://my-product-api.com"

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(string)).AnnotatedWith("config:productApiUrl").ToInstance(ProductApiUrl)
	injector.Bind(new(domain.ProductAPI)).To(infrastructure.ConcreteProductAPI{})
}
