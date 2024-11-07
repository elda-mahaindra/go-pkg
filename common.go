package package_query

import "fmt"

const VERSION = "1.0.7"

type DriverName string
type ServiceName string

const (
	DriverOracle   DriverName = "godror"
	DriverPostgres DriverName = "postgresql"
)

const (
	SvcFundAccount  ServiceName = "fundaccount"
	SvcFundFeat     ServiceName = "fundfeat"
	SvcFundParam    ServiceName = "fundparam"
	SvcFundReport   ServiceName = "fundreport"
	SvcFundTimeDep  ServiceName = "fundtimedep"
	SvcFundTx       ServiceName = "fundtx"
	SvcFundTxBatch  ServiceName = "fundtxbatch"
	SvcFundTxOnline ServiceName = "fundtxonline"
	SvcFundTxRemit  ServiceName = "fundtxremit"
	SvcFull         ServiceName = "full"
)

// SQL_FILES maps database driver names to their corresponding SQL file names.
var SQL_FILES = map[DriverName]string{
	DriverOracle:   "oracle.sql",
	DriverPostgres: "postgres.sql",
}

// ErrUnrecognizedDriver is returned when an unrecognized driver is provided.
var ErrUnrecognizedDriver = fmt.Errorf("unrecognized driver")

// ErrUnrecognizedService is returned when an unrecognized service is provided.
var ErrUnrecognizedService = fmt.Errorf("unrecognized service")
