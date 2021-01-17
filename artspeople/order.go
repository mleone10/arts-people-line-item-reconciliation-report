package artspeople

// An Order represents all details of a single interaction with a customer, including all items purchased and the payment method.
// TODO: Consider converting from a struct to a type alias ([]*LineItem)
type Order struct {
	LineItems []*LineItem
}

// AddLineItem appends a LineItem to the given Order.
func (o *Order) AddLineItem(li *LineItem) {
	o.LineItems = append(o.LineItems, li)
}
