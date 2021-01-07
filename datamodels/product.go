package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" homework:"id"`
	ProductName  string `json:"ProductName" sql:"productName" homework:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" homework:"ProductNum"`
	ProductPrice   string `json:"ProductPrice" sql:"productPrice" homework:"ProductPrice"`
	ProductOriginPrice   string `json:"ProductOriginPrice" sql:"productOriginPrice" homework:"ProductOriginPrice"`
	ProductImage string `json:"ProductImage" sql:"productImage" homework:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" homework:"ProductUrl"`
}
