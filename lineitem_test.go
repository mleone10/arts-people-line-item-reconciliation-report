package artspeople_test

import (
	"testing"
	"time"

	"github.com/mleone10/arts-people-reconciler/artspeople"
)

func TestNewLineItem_ValidRawLine(t *testing.T) {
	testRawLine := []string{
		"12345678",
		"2020-08-12 10:41 PM",
		"Donation",
		"John Doe",
		"20.00",
		"30.00",
		"40.00",
		"Visa",
		"",
		"Online",
	}

	li, err := artspeople.NewLineItem(testRawLine)
	if err != nil {
		t.Fatalf("Unexpected error when parsing valid raw line: %v", err)
	}

	testOrderID := 12345678
	if li.OrderID != testOrderID {
		t.Errorf("Expected order ID %v, was actually %v", testOrderID, li.OrderID)
	}

	testDatetime := time.Date(2020, time.August, 12, 22, 41, 0, 0, time.UTC)
	if li.DateTime.Year() != 2020 || li.DateTime.Month() != time.August || li.DateTime.Day() != 12 || li.DateTime.Hour() != 22 || li.DateTime.Minute() != 41 {
		t.Errorf("Expected datetime %v, was actually %v", testDatetime, li.DateTime)
	}

	testItemName := "Donation"
	if li.ItemName != testItemName {
		t.Errorf("Expected item name %v, was actually %v", testItemName, li.ItemName)
	}

	testCustomer := "John Doe"
	if li.Customer != testCustomer {
		t.Errorf("Expected customer %v, was actually %v", testCustomer, li.Customer)
	}

	testPrice := artspeople.Currency(2000)
	if li.Price != testPrice {
		t.Errorf("Expected price %v, was actually %v", testPrice, li.Price)
	}

	testFees := artspeople.Currency(3000)
	if li.Fees != testFees {
		t.Errorf("Expected fees %v, was actually %v", testFees, li.Fees)
	}

	testPurchaseTotal := artspeople.Currency(4000)
	if li.PurchaseTotal != testPurchaseTotal {
		t.Errorf("Expected purchase total %v, was actually %v", testPurchaseTotal, li.PurchaseTotal)
	}

	testPaymentMethod := "Visa"
	if li.PaymentMethod != testPaymentMethod {
		t.Errorf("Expected payment method %v, was actually %v", testPaymentMethod, li.PaymentMethod)
	}

	testUsername := "Online"
	if li.Username != testUsername {
		t.Errorf("Expected username %v, was actually %v", testUsername, li.Username)
	}
}

func TestNewLineItem_NotEnoughFields(t *testing.T) {
	testRawLine := []string{
		"foo",
		"bar",
	}

	_, err := artspeople.NewLineItem(testRawLine)
	if err == nil {
		t.Fatalf("Expected an error due to insufficient number of fields")
	}
}

func TestNewLineItem_InvalidOrderID(t *testing.T) {
	testRawLine := []string{
		"invalidOrderID",
		"2020-08-12 10:41 PM",
		"Donation",
		"John Doe",
		"20.00",
		"30.00",
		"40.00",
		"Visa",
		"",
		"Online",
	}

	_, err := artspeople.NewLineItem(testRawLine)
	if err == nil {
		t.Fatalf("Expected an error while parsing a non-int order ID")
	}
}

func TestIsPayment(t *testing.T) {
	paymentItem := artspeople.LineItem{ItemName: "Payment"}
	nonPaymentItem := artspeople.LineItem{ItemName: "NotAPayment"}

	if !paymentItem.IsPayment() {
		t.Errorf("Expected payment item was not determined to be a payment")
	}

	if nonPaymentItem.IsPayment() {
		t.Errorf("Expected nonPayment item was determined to be a payment")
	}
}
