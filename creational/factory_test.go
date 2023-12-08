package creational

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(1)
	if err != nil {
		t.Fatal("Cash type must exist")
	}
	msg := payment.Pay(1.20)
	if !strings.Contains(msg, "Payed Using cash") {
		t.Error("cash paymenyt method not correct")
	}
	t.Log("Log:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(2)
	if err != nil {
		t.Fatal("Debit Card type must exist")
	}
	msg := payment.Pay(10.20)
	if !strings.Contains(msg, "Payed Using Debit card") {
		t.Error("Debitcard payment method not correct")
	}
	t.Log("Log:", msg)
}
func TestGetPaymentMethodCreditCard(t *testing.T) {
	payment, err := GetPaymentMethod(3)
	if err != nil {
		t.Fatal("Credit card type must exist")
	}
	msg := payment.Pay(10.20)
	if !strings.Contains(msg, "Payed Using Credit Card") {
		t.Error("Creditcard payment method not correct")
	}
	t.Log("Log:", msg)
}

func TestGetPaymentMethodNe(t *testing.T) {
	_, err := GetPaymentMethod(100)
	if err == nil {
		t.Error("Debitcard payment method not correct")
	}
	t.Log("Log:", err)
}
