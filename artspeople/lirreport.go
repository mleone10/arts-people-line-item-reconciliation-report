package artspeople

import (
	"bufio"
	"io"
)

// A LineItemReconReport is a parsed and type-normalized version of the Line Item Reconciliation Report downloaded from Arts People.
type LineItemReconReport struct {
	rawLines []string
}

// NewLineItemReconReport accepts an Arts People Line Item Reconciliation Report as an io.Reader and returns a fully parsed and type-normalized LineItemReconReport.
func NewLineItemReconReport(reportFile io.Reader) (*LineItemReconReport, error) {
	lirReport := &LineItemReconReport{
		rawLines: []string{},
	}

	scanner := bufio.NewScanner(reportFile)
	for scanner.Scan() {
		lirReport.rawLines = append(lirReport.rawLines, scanner.Text())
	}

	return lirReport, nil
}

// GetRawLines returns the raw list of strings read in when creating the given LineItemReconReport.
func (l *LineItemReconReport) GetRawLines() []string {
	return l.rawLines
}
