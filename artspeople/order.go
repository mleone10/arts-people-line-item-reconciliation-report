package artspeople

// An Order represents all details of a single interaction with a customer, including all items purchased and the payment method.
// TODO: Consider converting from a struct to a type alias ([]*LineItem)
type Order struct {
	LineItems []*LineItem
}

// NewOrder returns an instantiated Order.
func NewOrder() *Order {
	return &Order{}
}

// AddLineItem appends a LineItem to the given Order.
func (o *Order) AddLineItem(li *LineItem) {
	o.LineItems = append(o.LineItems, li)
}

// GetCustomer returns the customer whose name is attached to the items in this order.  Sample data suggests that, while each LineItem has its own customer field, they are all identical.
func (o *Order) GetCustomer() string {
	return o.LineItems[0].Customer
}

// GetItems returns a slice of all items contained within the order, excluding payments.
func (o *Order) GetItems() []string {
	is := []string{}
	for _, li := range o.LineItems {
		i := li.ItemName
		if !li.IsPayment() {
			is = append(is, i)
		}
	}
	return is
}
