package usecase

import (
	"context"
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/dao"
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/model/goods"
)

type ProductDetailsUseCase struct {
	ProductDetailsDao dao.ProductDetailsDao
}

func (productDetailsUseCase *ProductDetailsUseCase) GetOpenProducts(ctx context.Context) ([]goods.ProductDetails, error) {
	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByPresenterId(ctx, "")
	if err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) GetMyExistingProducts(ctx context.Context, practitionerId string) ([]goods.ProductDetails, error) {
	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByPresenterId(ctx, practitionerId)
	if err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) GetPublicProductDetails(ctx context.Context, productId string) (*goods.ProductDetails, error) {
	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByProductId(ctx, productId, "")
	if err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) GetPrivateProductDetails(ctx context.Context, productId string, practitionerId string) (*goods.ProductDetails, error) {
	if practitionerId == "" {
		return nil, error2.Error(error2.RequestFieldEmpty)
	}

	r, err := productDetailsUseCase.ProductDetailsDao.ResolveByProductId(ctx, productId, practitionerId)
	if err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	}
	return r, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) Exhibition(ctx context.Context, c *command.ExhibitionCommand) (*goods.ProductDetails, error) {
	productDetails, err := c.ToProductDetails()
	if err != nil {
		return nil, err
	}
	if err := productDetailsUseCase.ProductDetailsDao.Insert(ctx, productDetails); err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	}
	return productDetails, nil
}

func (productDetailsUseCase *ProductDetailsUseCase) UpdateProductDetails(ctx context.Context, c *command.UpdateProductCommand) (*goods.ProductDetails, error) {
	productDetails, err := productDetailsUseCase.ProductDetailsDao.ResolveByProductId(ctx, c.ProductId, c.PractitionerId)
	if err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	}
	productDetails.Title = c.ProductDetail.Title
	productDetails.Detail = c.ProductDetail.Detail
	productDetails.Price = c.ProductDetail.Price
	productDetails.State = c.ProductDetail.State
	if _, err := productDetails.Validate(); err != nil {
		return nil, err
	}

	if r, err := productDetailsUseCase.ProductDetailsDao.Update(ctx, productDetails); err != nil {
		return nil, error2.Error(error2.SqlRequestFailed)
	} else if r == 0 {
		return nil, error2.Error(error2.NoUpdatesWereFound)
	}
	return productDetails, nil
}
