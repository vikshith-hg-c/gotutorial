package creational

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

const (
	cash       = 1
	DebitCard  = 2
	CreditCard = 3
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf("payment menthod %d not recogined", m)
	}
}

type CashPM struct{}

func (c *CashPM) Pay(amount float64) string {
	return "Payed Using cash"
}

type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amount float64) string {
	return "Payed Using Debit card"
}

type CreditCardPM struct{}

func (c *CreditCardPM) Pay(amount float64) string {
	return "Payed Using Credit Card"
}
