package handlers

import (
	"fmt"
	"service-order/constant"
	orders "service-order/module/order"
)

func NewOrderControllerKafka(orderService orders.OrderServiceInterface) {
	go ListenTopicProduct(orderService)
}

func ListenTopicProduct(orderService orders.OrderServiceInterface) {

	for {
		select {
		case data := <-constant.ChanTopicProduct:
			fmt.Println("bebek ", data)
		}
	}

}
