# Arts People Line Item Reconciliation Report Parser

This package provides a parser for the [Arts People](https://www.arts-people.com/) Line Item Reconciliation report.

```
package artspeople // import "github.com/mleone10/arts-people-reconciler"

Package artspeople provides a parser for the Arts People Line Item
Reconciliation Report. To create an instantiated LineItemReconReport struct,
provide the text of the report's CSV representation into the
NewLineItemReconReport() function.

TYPES

type Currency int64
    A Currency is a representation of money within Arts People. Since we're
    parsing these from strings, we have some flexibility around how we represent
    these values.

func NewCurrencyFromString(s string) (Currency, error)
    NewCurrencyFromString converts a string into a cents-only representation of
    a USD monetary amount. This is apparently "Martin Fowler's method".

type LineItem struct {
	OrderID       int
	DateTime      time.Time
	ItemName      string
	Customer      string
	Price         Currency
	Fees          Currency
	PurchaseTotal Currency
	PaymentMethod string
	GCUsed        string // This field does not appear to be used.  As such, we leave it as a string until test data exists.
	Username      string
	// Has unexported fields.
}
    A LineItem represents a single piece of an order, such as a ticket,
    membership, donation, or payment.

func NewLineItem(rawLine []string) (*LineItem, error)
    NewLineItem accepts an array of strings read in from the original CSV and
    returns an initiatlized LineItem.

func (li *LineItem) IsPayment() bool
    IsPayment returns true if the given LineItem pertains to a payment.

type LineItemReconReport struct {
	Orders map[int]*Order
	// Has unexported fields.
}
    A LineItemReconReport is a parsed and type-normalized version of the Line
    Item Reconciliation Report downloaded from Arts People.

func NewLineItemReconReport(reportCsv io.Reader) (*LineItemReconReport, error)
    NewLineItemReconReport accepts an Arts People Line Item Reconciliation
    Report as an io.Reader and returns a fully parsed and type-normalized
    LineItemReconReport.

func (l *LineItemReconReport) GetCustomers() []string
    GetCustomers returns a slice of all unique customers who have an order in
    the report.

func (l *LineItemReconReport) GetItems() []string
    GetItems returns a slice of all item names mentioned in the report.

type Order struct {
	LineItems []*LineItem
}
    An Order represents all details of a single interaction with a customer,
    including all items purchased and the payment method. TODO: Consider
    converting from a struct to a type alias ([]*LineItem)

func NewOrder() *Order
    NewOrder returns an instantiated Order.

func (o *Order) AddLineItem(li *LineItem)
    AddLineItem appends a LineItem to the given Order.

func (o *Order) GetCustomer() string
    GetCustomer returns the customer whose name is attached to the items in this
    order. Sample data suggests that, while each LineItem has its own customer
    field, they are all identical.

func (o *Order) GetItems() []string
    GetItems returns a slice of all items contained within the order, excluding
    payments.
```
