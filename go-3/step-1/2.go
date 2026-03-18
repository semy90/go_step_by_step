package main

/*Напишите функцию (logger *OrderLogger) AddOrder(order Order), которая пишет в терминал информацию о добавленом заказе вида Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f

Примечания
Ваше решение должно содержать следующий код:

// Order представляет информацию о заказе.
type Order struct {
    OrderNumber  int
    CustomerName string
    OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct{}

// NewOrderLogger создает новый экземпляр OrderLogger.
func NewOrderLogger() *OrderLogger {
    return &OrderLogger{}
}*/
import (
	"fmt"
)

// Order представляет информацию о заказе.
type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct {
}

func (logger *OrderLogger) AddOrder(order Order) {
	fmt.Println(fmt.Sprintf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f",
		order.OrderNumber, order.CustomerName, order.OrderAmount))
}

// NewOrderLogger создает новый экземпляр OrderLogger.
func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}
