# package-query

`package-query` is a centralized query management system for our core banking database. It provides a unified way to load and manage SQL queries across different services and database drivers.

## Overview

This package serves as a repository for all SQL queries related to our core banking system. It allows for easy management and retrieval of queries based on specific services and database drivers.

## Features

- Centralized storage of SQL queries
- Support for multiple database drivers (Oracle, PostgreSQL)
- Service-specific query loading

## Installation

To use this package in your project, run:

```
go get github.com/elda-mahaindra/go-pkg
```

## Usage

### Loading Queries

To load queries for a specific service and database driver:

```go
import pquery "github.com/elda-mahaindra/go-pkg"
import "github.com/elda-mahaindra/go-pkg/registration"

queries, err := pquery.LoadQueries(pquery.DriverOracle, pquery.SvcFundTxBatch)
if err != nil {
    // Handle error
}

// Use the loaded queries
sqlQuery := fmt.Sprintf(queries[registration.QUERY__GetListTrxPayroll_main], "core")
```

### Supported Services

- fundaccount
- fundfeat
- fundparam
- fundreport
- fundtimedep
- fundtxremit
- fundtx
- fundtxbatch
- fundtxonline
- full

If no specific service is specified, all queries will be loaded.

## Structure

- `oracle.sql`: Contains queries for Oracle database
- `postgres.sql`: Contains queries for PostgreSQL database
- `registration/`: Package containing query name registrations for different services

## Future Plans

This package is designed to be the foundation for separating our datastore package. In the future, it may include separate implementations for different data store operations, allowing services to be built with a two-layer architecture.
