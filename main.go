package main

import (
	"fmt"
	"log"

	"gitlab.com/ihsansolusi/posindonesia/package-query/registration"
)

func main() {
	queries, err := LoadQueries(DriverOracle, SvcFundTxBatch)
	if err != nil {
		log.Fatalf("failed to load queries: %s", err.Error())
	}

	// Use the loaded queries
	sqlQuery := fmt.Sprintf(queries[registration.QUERY__GetListTrxPayroll_main], "core")

	log.Println(sqlQuery)
}
