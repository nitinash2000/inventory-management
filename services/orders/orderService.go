package orders

import (
	"inventory-management/constants"
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository"
	"time"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(req *dtos.Order) error
	UpdateOrder(id string, req *dtos.Order) error
	GetOrder(orderId string) (*dtos.Order, error)
	DeleleOrder(orderId string) error
}

type orderService struct {
	orderRepo     repository.OrderRepo
	orderItemRepo repository.OrderItemRepo
}

func NewOrderService(orderRepo repository.OrderRepo, orderItemRepo repository.OrderItemRepo) OrderService {
	return &orderService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (o *orderService) CreateOrder(req *dtos.Order) error {
	orderModel, itemsModel := OrderDtosToModel(req)

	err := o.orderRepo.Create(orderModel)
	if err != nil {
		return err
	}

	err = o.orderItemRepo.Create(itemsModel...)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) UpdateOrder(id string, req *dtos.Order) error {
	if req.OrderId == "" {
		return constants.ErrorOrderIdEmpty
	}

	orderModel, itemsModel := OrderDtosToModel(req)

	err := o.orderRepo.Create(orderModel)
	if err != nil {
		return err
	}

	err = o.orderItemRepo.Create(itemsModel...)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) GetOrder(orderId string) (*dtos.Order, error) {
	order, err := o.orderRepo.Get(orderId)
	if err != nil {
		return nil, err
	}

	orderItems, err := o.orderItemRepo.GetByOrder(orderId)
	if err != nil {
		return nil, err
	}

	result := OrderModelToDtos(order, orderItems)

	return result, nil
}

func (o *orderService) DeleleOrder(orderId string) error {
	err := o.orderRepo.Delete(orderId)
	if err != nil {
		return err
	}

	return nil
}

func OrderModelToDtos(m *models.Order, i []*models.OrderItem) *dtos.Order {
	o := &dtos.Order{
		OrderId:     m.OrderId,
		CustomerId:  m.CustomerId,
		OrderedAt:   m.OrderedAt,
		TotalAmount: m.TotalAmount,
		NoOfItems:   m.NoOfItems,
		Items:       []*dtos.OrderItems{},
	}

	var items []*dtos.OrderItems
	for _, v := range i {
		items = append(items, &dtos.OrderItems{
			OrderItemId: v.OrderItemId,
			OrderId:     v.OrderId,
			ArticleId:   v.ArticleId,
			Quantity:    v.Quantity,
		})
	}

	o.Items = items

	return o
}

func OrderDtosToModel(m *dtos.Order) (*models.Order, []*models.OrderItem) {
	orderId := m.OrderId
	if orderId == "" {
		orderId = uuid.NewString()
	}

	if m.OrderedAt.IsZero() {
		m.OrderedAt = time.Now().UTC()
	}

	order := &models.Order{
		OrderId:     orderId,
		CustomerId:  m.CustomerId,
		OrderedAt:   m.OrderedAt,
		TotalAmount: m.TotalAmount,
		NoOfItems:   len(m.Items),
	}

	var orderItems []*models.OrderItem
	for _, v := range m.Items {
		if v.OrderItemId == "" {
			v.OrderItemId = uuid.NewString()
		}

		orderItems = append(orderItems, &models.OrderItem{
			OrderItemId: v.OrderItemId,
			OrderId:     orderId,
			ArticleId:   v.ArticleId,
			Quantity:    v.Quantity,
		})
	}

	return order, orderItems
}
