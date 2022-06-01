package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-order/internal/model"

	ospb "github.com/smhdhsn/restaurant-order/internal/protos/order/submission"
	serviceContract "github.com/smhdhsn/restaurant-order/internal/service/contract"
)

// OrderSubmitHandler contains services that can be used within order submit handler.
type OrderSubmitHandler struct {
	osServ serviceContract.OrderSubmitService
}

// NewOrderSubmitHandler creates a new order submit handler.
func NewOrderSubmitHandler(osServ serviceContract.OrderSubmitService) ospb.OrderSubmissionServiceServer {
	return &OrderSubmitHandler{
		osServ: osServ,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *OrderSubmitHandler) Submit(ctx context.Context, req *ospb.OrderSubmitRequest) (*ospb.OrderSubmitResponse, error) {
	oReq := model.OrderDTO{
		FoodID: req.GetFoodId(),
		UserID: req.GetUserId(),
	}

	err := s.osServ.Submit(&oReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := ospb.OrderSubmitResponse{
		Status: true,
	}

	return &resp, nil
}
