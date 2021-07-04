package data

import "github.com/asergio68/imersaoFC3/goapp/entity"

var Products entity.Products

func LoadData() {
	Products.Add(*entity.NewProduct("carrinho"))
	Products.Add(*entity.NewProduct("boneca"))
}