package product

import (
	"context"

	product "github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/status"
)

type Service struct {
	log  *logger.Logger
	repo product.Repository
}

func NewService(log *logger.Logger, repo product.Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) GetProductList(ctx context.Context, req GetRequest) ([]*GetResponse, error) {
	products, err := s.repo.List(ctx, req.CategoryId, req.Limit, req.StartId)
	if err != nil {
		s.log.Usecase("ProductList").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := make([]*GetResponse, 0)
	for _, p := range products {
		resp = append(resp, MakeResponse(p))
	}
	return resp, nil
}

func (s *Service) GetProductById(ctx context.Context, req GetRequest) (*GetResponse, error) {
	res, err := s.repo.GetById(ctx, req.ProductId)
	if err != nil {
		s.log.Usecase("GetProductById").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := MakeResponse(res)
	return resp, nil
}

func (s *Service) GetCategoryList(ctx context.Context) ([]*CategoryResponse, error) {
	categories, err := s.repo.CategoryList(ctx)
	if err != nil {
		s.log.Usecase("CategoryList").Error(err)
		return nil, err
	}
	resp := make([]*CategoryResponse, 0)
	for _, c := range categories {
		resp = append(resp, MakeCategoryResponse(c))
	}
	return resp, nil
}

func (s *Service) GetCartItemList(ctx context.Context, req CartRequest) ([]*CartResponse, error) {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("GetCartItemList").Error(ctxt.ErrGetUserFromContext)
		return nil, status.Error(codes.Internal, "Failed to get cart item list")
	}
	carts, err := s.repo.CartList(ctx, user.Username, req.Limit, req.StartId)
	if err != nil {
		s.log.Usecase("GetCartList").Error(err)
		return nil, status.Error(codes.BadRequest, err.Error())
	}
	resp := make([]*CartResponse, 0)
	for _, c := range carts {
		resp = append(resp, MakeCartResponse(c))
	}
	return resp, nil
}

func (s *Service) AddCartItem(ctx context.Context, req AddCartItemRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("AddCartItem").Error(ctxt.ErrGetUserFromContext)
		return status.Error(codes.Internal, "Failed to add cart item to list")
	}
	err := s.repo.SaveCart(ctx, product.Cart{
		Username:  user.Username,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	})
	if err != nil {
		s.log.Usecase("AddCartItem").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *Service) UpdateCartItemQuantity(ctx context.Context, req UpdateCartItemRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("UpdateCartItemQuantity").Error(ctxt.ErrGetUserFromContext)
		return status.Error(codes.Internal, "Failed to add cart item to list")
	}
	err := s.repo.UpdateCart(ctx, req.Id, product.Cart{
		Username: user.Username,
		Quantity: req.Quantity,
	})
	if err != nil {
		s.log.Usecase("UpdateCartItemQuantity").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *Service) DeleteCartItem(ctx context.Context, req DeleteCartItemRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("DeleteCartItem").Error(ctxt.ErrGetUserFromContext)
		return status.Error(codes.Internal, "Failed to add cart item to list")
	}
	err := s.repo.DeleteCart(ctx, req.Id, product.Cart{
		Username: user.Username,
	})
	if err != nil {
		s.log.Usecase("DeleteCartItem").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *Service) GetDiscountList(ctx context.Context, req DiscountRequest) ([]*DiscountResponse, error) {
	discounts, err := s.repo.DiscountList(ctx, req.Limit, req.StartId)
	if err != nil {
		s.log.Usecase("GetDiscountList").Error(err)
		return nil, status.Error(codes.BadRequest, err.Error())
	}
	resp := make([]*DiscountResponse, 0)
	for _, d := range discounts {
		resp = append(resp, MakeDiscountResponse(d))
	}
	return resp, nil
}
