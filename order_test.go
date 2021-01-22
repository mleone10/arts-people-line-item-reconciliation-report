package artspeople_test

import (
	"testing"

	"github.com/mleone10/arts-people-reconciler/artspeople"
)

func TestAddLineItem_ItemAppended(t *testing.T) {
	o := artspeople.Order{}
	li := &artspeople.LineItem{}

	o.AddLineItem(li)

	if len(o.LineItems) != 1 {
		t.Errorf("Line item was not added to order")
	}
}

func TestGetCustomer_ReturnsFirstLineItemCustomer(t *testing.T) {
	testCustomer := "John Doe"
	o := artspeople.Order{
		LineItems: []*artspeople.LineItem{
			&artspeople.LineItem{Customer: testCustomer},
		},
	}

	actual := o.GetCustomer()
	if actual != testCustomer {
		t.Errorf("Order returned customer [%v], expected %v", actual, testCustomer)
	}
}

func TestGetItems_ReturnsListOfAllNonPaymentItems(t *testing.T) {
	o := artspeople.Order{
		LineItems: []*artspeople.LineItem{
			&artspeople.LineItem{ItemName: "Payment"},
			&artspeople.LineItem{ItemName: "Ticket"},
			&artspeople.LineItem{ItemName: "Donation"},
			&artspeople.LineItem{ItemName: "Ticket"},
		},
	}

	actual := len(o.GetItems())
	expected := 3

	if actual != expected {
		t.Errorf("Expected %v items returned from GetItems call, only received %v: [%v]", expected, actual, o.GetItems())
	}
}
