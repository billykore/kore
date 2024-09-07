package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
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
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		uc.log.Usecase("GetCartItemList").Error(ctxt.ErrGetUserFromContext)
		return nil, status.Error(codes.Internal, "Failed to get cart item list")
	}
	carts, err := uc.cartRepo.List(ctx, user.Username, req.Limit, req.StartId)
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
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		uc.log.Usecase("AddCartItem").Error(ctxt.ErrGetUserFromContext)
		return status.Error(codes.Internal, "Failed to add cart item to list")
	}
	err := uc.cartRepo.Save(ctx, model.Cart{
		Username:  user.Username,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	})
	if err != nil {
		uc.log.Usecase("AddCartItem").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (uc *CartUsecase) UpdateCartItemQuantity(ctx context.Context, req entity.UpdateCartItemRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		uc.log.Usecase("UpdateCartItemQuantity").Error(ctxt.ErrGetUserFromContext)
		return status.Error(codes.Internal, "Failed to add cart item to list")
	}
	err := uc.cartRepo.Update(ctx, req.Id, model.Cart{
		Username: user.Username,
		Quantity: req.Quantity,
	})
	if err != nil {
		uc.log.Usecase("UpdateCartItemQuantity").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (uc *CartUsecase) DeleteCartItem(ctx context.Context, req entity.DeleteCartItemRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		uc.log.Usecase("DeleteCartItem").Error(ctxt.ErrGetUserFromContext)
		return status.Error(codes.Internal, "Failed to add cart item to list")
	}
	err := uc.cartRepo.Delete(ctx, req.Id, model.Cart{
		Username: user.Username,
	})
	if err != nil {
		uc.log.Usecase("DeleteCartItem").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
