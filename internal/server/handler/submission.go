package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
	"github.com/smhdhsn/restaurant-order/internal/service/dto"

	log "github.com/smhdhsn/restaurant-order/internal/logger"
	submissionProto "github.com/smhdhsn/restaurant-order/internal/protos/order/submission"
	serviceContract "github.com/smhdhsn/restaurant-order/internal/service/contract"
)

// SubmitHandler contains services that can be used within order submit handler.
type SubmitHandler struct {
	submissionServ serviceContract.SubmissionService
}

// NewSubmitHandler creates a new order submit handler.
func NewSubmitHandler(os serviceContract.SubmissionService) submissionProto.OrderSubmissionServiceServer {
	return &SubmitHandler{
		submissionServ: os,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *SubmitHandler) Submit(ctx context.Context, req *submissionProto.OrderSubmitRequest) (*submissionProto.OrderSubmitResponse, error) {
	oDTO := singleOrderReqToDTO(req)

	err := s.submissionServ.Submit(oDTO)
	if err != nil {
		if errors.Is(err, serviceContract.ErrLackOfComponents) {
			log.Error(err)
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := new(submissionProto.OrderSubmitResponse)

	return resp, nil
}

// singleOrderReqToDTO is responsible for transforming an order request into order dto struct.
func singleOrderReqToDTO(req *submissionProto.OrderSubmitRequest) *dto.Order {
	return &dto.Order{
		FoodID: req.GetFoodId(),
		UserID: req.GetUserId(),
	}
}
