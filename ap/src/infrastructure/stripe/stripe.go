package stripe

import (
	stripego "github.com/stripe/stripe-go/v71"
	stripeclient "github.com/stripe/stripe-go/v71/client"
)

// AlmaStripe StripeのWrapper
type AlmaStripe struct {
	client *stripeclient.API
}

// GetAllConsumer 全ての顧客をリストする
func (c *AlmaStripe) GetAllConsumer() {

}

// GetAllProductList 全ての商品を取得する
func (c *AlmaStripe) GetAllProductList() []*stripego.Product {

	params := &stripego.ProductListParams{}
	params.Filters.AddFilter("", "", "")

	var productList []*stripego.Product
	i := c.client.Products.List(params)
	for i.Next() {
		productList = append(productList, i.Product())
	}

	return productList
}
