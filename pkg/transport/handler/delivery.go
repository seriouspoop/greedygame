package handler

import (
	"context"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/protos/go/deliverypb"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type DeliveryHandler struct {
	s      deliveryServicer
	logger *logging.Logger
	deliverypb.UnimplementedDeliveryServer
}

// New Grpc server handler
func NewDeliveryHandler(grpc *grpc.Server, s deliveryServicer, logger *logging.Logger) {
	handler := &DeliveryHandler{s: s, logger: logger}
	deliverypb.RegisterDeliveryServer(grpc, handler)
}

// New Grpc gateway handler
func NewDeliveryGatewayHandler(ctx context.Context, runtime *runtime.ServeMux, s deliveryServicer, logger *logging.Logger) {
	handler := &DeliveryHandler{s: s, logger: logger}
	deliverypb.RegisterDeliveryHandlerServer(ctx, runtime, handler)
}

func (d *DeliveryHandler) GetDelivery(ctx context.Context, req *deliverypb.DeliveryRequest) (*deliverypb.DeliveryResponse, error) {
	logger := d.logger.Ctx(ctx)
	app := req.GetApp()
	os := req.GetOs()
	country := req.GetCountry()

	logger.Info("helloooo")
	time.Sleep(5 * time.Second)
	campaigns, err := d.s.GetActiveCampaignForDelivery(ctx, app, os, country)

	if err != nil {
		logger.Error("error while getting campaigns for delivery", zap.Error(err))
		return nil, grpcError(err)
	}

	deliveryItems := make([]*deliverypb.DeliveryResponseItem, 0, len(campaigns))
	for _, campaign := range campaigns {
		delivery := &deliverypb.DeliveryResponseItem{
			Cid:   campaign.ID.String(),
			Image: campaign.Image.String(),
			Cta:   campaign.CTA,
		}
		deliveryItems = append(deliveryItems, delivery)
	}

	res := &deliverypb.DeliveryResponse{
		Items: deliveryItems,
	}
	return res, nil
}
