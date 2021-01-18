package main

import (
	"log"
	"os"

	"github.com/mleone10/arts-people-reconciler/artspeople"
)

func main() {
	r, err := artspeople.NewLineItemReconReport(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read line item reconciliation report: %v", err)
	}

	log.Println(r.GetCustomers())
	log.Println(r.GetItems())
}
