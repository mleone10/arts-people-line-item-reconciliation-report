package artspeople_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mleone10/arts-people-reconciler/artspeople"
)

func TestNewLineItemReconReport(t *testing.T) {
	testCsv := []string{
		`"Order ID","Date/time","Item name","Customer","Price","Fees","Purchase total","Payment method","GC used","Username"`,
		`"12345678","2020-08-12 10:41 PM","Donation","John Doe","20.00","","","Visa","","Online"`,
		`"12345678","2020-02-12 09:27 PM","A Daughter's A Daughter","John Doe","0.00","","","Pass","","Online"`,
		`"12345678","2020-02-12 09:27 PM","A Daughter's A Daughter","John Doe","23.00","","","Visa","","Online"`,
		`"12345678","2020-02-12 09:27 PM","Item Fee - FEE per Ticket","John Doe","","2.00","","Visa","","Online"`,
		`"12345678","2020-02-12 09:27 PM","Payment","John Doe","","","45.00","Visa","","Online"`,
	}

	lirReport, _ := artspeople.NewLineItemReconReport(bytes.NewBufferString(strings.Join(testCsv, "\n")))

	if len(lirReport.GetRawLines()) != len(testCsv) {
		t.Fatalf("Report contained %d raw lines, expected %d", len(lirReport.GetRawLines()), len(testCsv))
	}
}
