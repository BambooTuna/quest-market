package usecase

import (
	"context"
	"errors"
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/dao"
	"github.com/BambooTuna/quest-market/backend/model/goods"
)

type ProductDetailsUseCase struct {
	ProductDetailsDao dao.ProductDetailsDao
}

func (productDetailsUseCase *ProductDetailsUseCase) GetOpenProducts(ctx context.Context) ([]goods.ProductDetails, error) {
	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByPresenterId(ctx, "")
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) GetMyExistingProducts(ctx context.Context, practitionerId string) ([]goods.ProductDetails, error) {
	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByPresenterId(ctx, practitionerId)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) GetPublicProductDetails(ctx context.Context, productId string) (*goods.ProductDetails, error) {
	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByProductId(ctx, productId, "")
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) GetPrivateProductDetails(ctx context.Context, productId string, practitionerId string) (*goods.ProductDetails, error) {
	if practitionerId == "" {
		return nil, errors.New("practitionerId is empty")
	}

	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByProductId(ctx, productId, practitionerId)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) Exhibition(ctx context.Context, c *command.ExhibitionCommand) (*goods.ProductDetails, error) {
	productDetails, err := c.ToProductDetails()
	if err != nil {
		return nil, err
	}
	if err := productDetailsUseCase.ProductDetailsDao.Insert(ctx, productDetails); err != nil {
		return nil, err
	}
	return productDetails, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) UpdateProductDetails(ctx context.Context, c *command.UpdateProductCommand) (*goods.ProductDetails, error) {
	productDetails, err := productDetailsUseCase.ProductDetailsDao.ResolveByProductId(ctx, c.ProductId, c.PractitionerId)
	if err != nil {
		return nil, err
	}
	productDetails.Title = c.ProductDetail.Title
	productDetails.Detail = c.ProductDetail.Detail
	productDetails.Price = c.ProductDetail.Price
	productDetails.State = c.ProductDetail.State

	r, err := productDetailsUseCase.ProductDetailsDao.Update(ctx, productDetails)
	if err != nil {
		return nil, err
	}
	if r == 0 {
		return nil, errors.New("一件も更新できませんでした")
	}
	return productDetails, nil
}
