package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/billykore/kore/backend/pkg/status"
)

type CartUsecase struct {
	log      *log.Logger
	cartRepo repo.CartRepository
}

func NewCartUsecase(log *log.Logger, cartRepo repo.CartRepository) *CartUsecase {
	return &CartUsecase{
		log:      log,
		cartRepo: cartRepo,
	}
}

func (uc *CartUsecase) GetCartItemList(ctx context.Context, req entity.CartRequest) ([]*entity.CartResponse, error) {
	carts, err := uc.cartRepo.List(ctx, req.UserId, req.Limit, req.StartId)
	if err != nil {
		uc.log.Usecase("GetCartList").Error(err)
		return nil, status.Error(codes.BadRequest, err.Error())
	}
	resp := make([]*entity.CartResponse, 0)
	for _, c := range carts {
		resp = append(resp, entity.MakeCartResponse(c))
	}
	return resp, nil
}

func (uc *CartUsecase) AddCartItem(ctx context.Context, req entity.AddCartItemRequest) error {
	err := uc.cartRepo.Save(ctx, model.Cart{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	})
	if err != nil {
		uc.log.Usecase("AddCart").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (uc *CartUsecase) UpdateCartItemQuantity(ctx context.Context, req entity.UpdateCartItemRequest) error {
	err := uc.cartRepo.Update(ctx, req.Id, req.Quantity)
	if err != nil {
		uc.log.Usecase("UpdateCartItemQuantity").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (uc *CartUsecase) DeleteCartItem(ctx context.Context, req entity.DeleteCartItemRequest) error {
	err := uc.cartRepo.Delete(ctx, req.Id)
	if err != nil {
		uc.log.Usecase("DeleteCartItem").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
