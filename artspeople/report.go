package artspeople

import (
	"encoding/csv"
	"fmt"
	"io"
)

// A LineItemReconReport is a parsed and type-normalized version of the Line Item Reconciliation Report downloaded from Arts People.
type LineItemReconReport struct {
	rawLines [][]string
	orders   map[int]*Order
}

// NewLineItemReconReport accepts an Arts People Line Item Reconciliation Report as an io.Reader and returns a fully parsed and type-normalized LineItemReconReport.
func NewLineItemReconReport(reportCsv io.Reader) (*LineItemReconReport, error) {
	lirReport := LineItemReconReport{}

	err := lirReport.readInput(reportCsv)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse input CSV: %v", err)
	}

	err = lirReport.parseRawLines()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse lines: %v", err)
	}

	return &lirReport, nil
}

// GetOrders returns a read-only map of orders contained within a given LineItemReconReport.
func (l *LineItemReconReport) GetOrders() map[int]Order {
	orders := map[int]Order{}
	for id, o := range l.orders {
		orders[id] = *o
	}
	return orders
}

func (l *LineItemReconReport) readInput(reportCsv io.Reader) error {
	lines, err := csv.NewReader(reportCsv).ReadAll()
	if err != nil {
		return err
	}

	// Throw away the first line, which contains the field headers.
	l.rawLines = lines[1:]
	return nil
}

func (l *LineItemReconReport) parseRawLines() error {
	if l.orders == nil {
		l.orders = map[int]*Order{}
	}

	for _, rl := range l.rawLines {
		// Use the raw line to create a LineItem.
		li, err := NewLineItem(rl)
		if err != nil {
			return fmt.Errorf("Failed to parse line [%s]: %v", rl, err)
		}

		// If this is the first Line for a given order, instantiate a new Order struct.
		if _, ok := l.orders[li.OrderID]; !ok {
			l.orders[li.OrderID] = NewOrder()
		}

		// Add the LineItem to the Order with the LineItem's OrderID.
		l.orders[li.OrderID].AddLineItem(li)
	}

	return nil
}
