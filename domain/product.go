// domain/product.go

package domain

import (
	"fmt"
)

type Product struct {
	Name        string
	Description string
	Amount      int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Description: %s, Amount: %d", p.Name, p.Description, p.Amount)
}
