package package_query

import (
	"fmt"
	"log"

	"github.com/elda-mahaindra/go-pkg/registration"
)

func Sample() {
	queries, err := LoadQueries(DriverOracle, SvcFundTxBatch)
	if err != nil {
		log.Fatalf("failed to load queries: %s", err.Error())
	}

	// Use the loaded queries
	sqlQuery := fmt.Sprintf(queries[registration.QUERY__GetListTrxPayroll_main], "core")

	log.Println(sqlQuery)
}
