package main

import (
	"fmt"
	"reflect"
)

func printWithBlankSpace(str string) {
	fmt.Println(str)
	fmt.Println()
}

type State interface {
	cancelOrder()
	verifyPayment()
	shipOrder()
}

type Order struct {
	cancelledOrderState     State
	paymentPendingState     State
	orderBeingPreparedState State
	orderShippedState       State
	currentState            State
}

func (o *Order) NewOrder() {
	o.paymentPendingState = &PaymentPendingState{order: o}
	o.cancelledOrderState = &CancelledOrderState{order: o}
	o.orderShippedState = &ShippedOrderState{order: o}
	o.orderBeingPreparedState = &OrderBeingPreparedState{order: o}
	o.currentState = o.paymentPendingState
}

func (o *Order) setState(state State) {
	o.currentState = state
}

func (o *Order) getState() State {
	return o.currentState
}

func (o *Order) printState() {
	str := fmt.Sprintf("Actual state is %s", reflect.TypeOf(o.currentState).String())
	printWithBlankSpace(str)
}

type PaymentPendingState struct {
	order *Order
}

func (p *PaymentPendingState) cancelOrder() {
	printWithBlankSpace("Cancelling your unpaid order")
	p.order.setState(p.order.cancelledOrderState)
}

func (p *PaymentPendingState) verifyPayment() {
	printWithBlankSpace("Payment verify! Shipping soon.")
	p.order.setState(p.order.orderBeingPreparedState)
}

func (p *PaymentPendingState) shipOrder() {
	printWithBlankSpace("Cannot ship the order when payment is pending!")
}

type OrderBeingPreparedState struct {
	order *Order
}

func (o *OrderBeingPreparedState) cancelOrder() {
	printWithBlankSpace("Cancelling order.")
	o.order.setState(o.order.cancelledOrderState)
}

func (o *OrderBeingPreparedState) verifyPayment() {
	printWithBlankSpace("Payment already verified.")
}
func (o *OrderBeingPreparedState) shipOrder() {
	printWithBlankSpace("Shipping your order now.")
	o.order.setState(o.order.orderShippedState)
}

type CancelledOrderState struct {
	order *Order
}

func (c *CancelledOrderState) cancelOrder() {
	printWithBlankSpace("Your order has already been cancelled.")
}
func (c *CancelledOrderState) verifyPayment() {
	printWithBlankSpace("Order cancelled. Cannot verify payment")
}
func (c *CancelledOrderState) shipOrder() {
	printWithBlankSpace("Cancelled order cannot be shipped.")
}

type ShippedOrderState struct {
	order *Order
}

func (s *ShippedOrderState) cancelOrder() {
	printWithBlankSpace("You cannot cancel order when its already shipped.")
}

func (s *ShippedOrderState) verifyPayment() {
	printWithBlankSpace("You cannot verify and already shipped order.")
}

func (s *ShippedOrderState) shipOrder() {
	printWithBlankSpace("You cannot ship it again. Order already shipped.")
}

func main() {
	order := Order{}
	order.NewOrder()
	order.printState()

	// Should be able to ship order in initial state
	order.getState().shipOrder()

	// Should verify payment and move to OrderBeingPreparedState
	order.getState().verifyPayment()
	order.printState()

	// Should ignore command since payment is already verified
	order.getState().verifyPayment()
	order.printState()

	// Should ship the order and move to ShippedOrderState
	order.getState().shipOrder()
	order.printState()

	anotherOrder := Order{}
	anotherOrder.NewOrder()
	anotherOrder.printState()

	// Should be able to ship order in initial state
	anotherOrder.getState().shipOrder()

	// Should cancel order and move to CancelledOrderState
	anotherOrder.getState().cancelOrder()
	anotherOrder.printState()
}
