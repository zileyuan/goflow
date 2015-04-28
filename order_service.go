package goflow

import (
	"time"

	"github.com/satori/go.uuid"
)

type OrderService struct {
}

func (p *OrderService) CreateOrder(process *Process, operator string) *Order {
	//model := process.Model
	order := &Order{
		Id:         uuid.NewV4().String(),
		ProcessId:  process.Id,
		Creator:    operator,
		CreateTime: time.Now(),
	}

	return order
}

func (p *OrderService) SaveOrder(order *Order) {
	historyOrder := new(HistoryOrder)
	historyOrder.DataFromOrder(order)

	historyOrder.OrderState = FS_ACTIVITY
	order.Save()
	historyOrder.Save()
}

func (p *OrderService) CompleteOrder(id string) {
	order := new(Order)
	order.GetOrderById(id)

	historyOrder := new(HistoryOrder)
	historyOrder.GetHistoryOrderById(id)
	historyOrder.OrderState = FS_FINISH

	historyOrder.Update()
	order.Delete()
}

func (p *OrderService) ResumeOrder(id string) {
	historyOrder := new(HistoryOrder)
	historyOrder.GetHistoryOrderById(id)
	historyOrder.OrderState = FS_ACTIVITY
	order := historyOrder.Undo()

	order.Save()
	historyOrder.Save()

}

func (p *OrderService) TerminateOrder(id string, operator string) {

}
