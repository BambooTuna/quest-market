package dao

import (
	"context"
	"fmt"
	"github.com/BambooTuna/quest-market/backend/model/goods"
	"gopkg.in/gorp.v1"
)

type ProductDetailsDaoImpl struct {
	DBSession *gorp.DbMap
}

func (a ProductDetailsDaoImpl) ResolveByProductId(ctx context.Context, productId string, practitionerId string) (*goods.ProductDetails, error) {
	var productDetails *goods.ProductDetails

	sql := fmt.Sprintf("select * from product_details where product_id = '%s' and presenter_id = '%s' and state != '%s'", productId, practitionerId, goods.Deleted)
	if practitionerId == "" {
		sql = fmt.Sprintf("select * from product_details where product_id = '%s' and state = '%s'", productId, goods.Open)
	}
	err := a.DBSession.SelectOne(&productDetails, sql)
	return productDetails, err
}

func (a ProductDetailsDaoImpl) ResolveByPurchasedProductId(ctx context.Context, productId string) (*goods.ProductDetails, error) {
	var productDetails *goods.ProductDetails
	sql := fmt.Sprintf("select * from product_details where product_id = '%s' and state == '%s'", productId, goods.Closed)
	err := a.DBSession.SelectOne(&productDetails, sql)
	return productDetails, err
}

func (a ProductDetailsDaoImpl) ResolveByPresenterId(ctx context.Context, practitionerId string) ([]goods.ProductDetails, error) {
	var productDetails []goods.ProductDetails

	sql := fmt.Sprintf("select * from product_details where presenter_id = '%s' and state != '%s'", practitionerId, goods.Deleted)
	if practitionerId == "" {
		sql = fmt.Sprintf("select * from product_details where state = '%s'", goods.Open)
	}
	_, err := a.DBSession.Select(&productDetails, sql)
	return productDetails, err
}

func (a ProductDetailsDaoImpl) Insert(ctx context.Context, record *goods.ProductDetails) error {
	return a.DBSession.Insert(record)
}

func (a ProductDetailsDaoImpl) Update(ctx context.Context, record *goods.ProductDetails) (int64, error) {
	return a.DBSession.Update(record)
}
