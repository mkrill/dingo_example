package main

import (
	"log"

	"flamingo.me/dingo"

	"github.com/mkrill/dingo_example/interfaces"
)

const ProductApi = "https://my-product-api.com"

func main() {
	// Without Dingo
	// *************
	//productAPI := infrastructure.ConcreteProductAPI{ProductApi}
	//
	//appController := interfaces.AppController{}
	//appController.New(&productAPI)

	// With Dingo
	// **********
	injector, err := dingo.NewInjector(new(Module))
	if err != nil {
		panic(err)
	}

	// Now that we've got the injector, we can build objects.
	// We get a new instance, and cast it accordingly:
	instance, err := injector.GetInstance(new(interfaces.AppController))
	if err != nil {
		log.Fatal(err)
	}

	appController := instance.(*interfaces.AppController)

	// Without and with Dingo
	// **********************
	appController.PrintProducts()
}
