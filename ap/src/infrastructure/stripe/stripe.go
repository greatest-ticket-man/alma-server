package stripe

import (
	"alma-server/ap/src/common/error/chk"

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

// CreateProduct 商品を作成する
func (c *AlmaStripe) CreateProduct(productName string) *stripego.Product {
	params := &stripego.ProductParams{
		Name: stripego.String(productName),
	}
	result, err := c.client.Products.New(params)
	chk.SE(err)
	return result
}

// GetProduct 商品を取得
func (c *AlmaStripe) GetProduct(productID string) *stripego.Product {
	result, err := c.client.Products.Get(productID, nil)
	chk.SE(err)
	return result
}

// DeleteProduct 商品を削除する
func (c *AlmaStripe) DeleteProduct(productID string) *stripego.Product {
	result, err := c.client.Products.Del(productID, nil)
	chk.SE(err)
	return result
}

// GetAllProductList 全ての商品を取得する
func (c *AlmaStripe) GetAllProductList() []*stripego.Product {

	params := &stripego.ProductListParams{}

	var productList []*stripego.Product
	i := c.client.Products.List(params)
	for i.Next() {
		productList = append(productList, i.Product())
	}

	return productList
}
