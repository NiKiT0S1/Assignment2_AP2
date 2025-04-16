package grpc

import (
	"context"
	pb "orderService/internal/delivery/grpc/pb"
	"orderService/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	orderUC domain.OrderUsecase
}

func NewOrderHandler(orderUC domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUC: orderUC}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	o := &domain.Order{
		UserID: int(req.UserId),
		Status: req.Status,
		Items:  toDomainItems(req.Items),
	}
	if err := h.orderUC.Create(o); err != nil {
		return nil, status.Errorf(codes.Internal, "create failed: %v", err)
	}
	return toProto(o), nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.OrderID) (*pb.Order, error) {
	o, err := h.orderUC.GetByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}
	return toProto(o), nil
}

func (h *OrderHandler) UpdateOrderStatus(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	o := &domain.Order{
		ID:     int(req.Id),
		Status: req.Status,
	}
	if err := h.orderUC.UpdateStatus(o.ID, o.Status); err != nil {
		return nil, status.Errorf(codes.Internal, "update failed: %v", err)
	}
	return toProto(o), nil
}

func (h *OrderHandler) ListOrdersByUser(ctx context.Context, req *pb.ListOrdersRequest) (*pb.OrderList, error) {
	orders, err := h.orderUC.ListByUser(int(req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list failed: %v", err)
	}
	var res pb.OrderList
	for _, o := range orders {
		res.Orders = append(res.Orders, toProto(&o))
	}
	return &res, nil
}

func toProto(o *domain.Order) *pb.Order {
	var items []*pb.OrderItem
	for _, item := range o.Items {
		items = append(items, &pb.OrderItem{
			Id:        int32(item.ID),
			OrderId:   int32(item.OrderID),
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}
	return &pb.Order{
		Id:     int32(o.ID),
		UserId: int32(o.UserID),
		Status: o.Status,
		Items:  items,
	}
}

func toDomainItems(items []*pb.OrderItem) []domain.OrderItem {
	var domainItems []domain.OrderItem
	for _, item := range items {
		domainItems = append(domainItems, domain.OrderItem{
			ProductID: int(item.ProductId),
			Quantity:  int(item.Quantity),
		})
	}
	return domainItems
}
