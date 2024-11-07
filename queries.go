package package_query

import (
	"embed"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/elda-mahaindra/go-pkg/registration"
	"github.com/qustavo/dotsql"
)

//go:embed oracle.sql postgres.sql
var sqlFiles embed.FS

// LoadQueries reads the SQL file for the specified driver and loads all defined queries.
// It returns a map of query names to their corresponding SQL strings.
//
// Parameters:
//   - driver: A DriverName specifying the database driver (DriverOracle or DriverPostgres)
//   - service: A ServiceName specifying the service (e.g., SvcFundAccount, SvcFundFeat, etc.)
//
// Returns:
//   - A map[string]string where keys are query names and values are the SQL strings
//   - An error if there was a problem reading the SQL file, parsing the queries,
//     or if an unrecognized driver or service was provided
func LoadQueries(driver DriverName, service ServiceName) (map[string]string, error) {
	filename, ok := SQL_FILES[driver]
	if !ok {
		return nil, ErrUnrecognizedDriver
	}

	sqlContent, err := sqlFiles.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file: %w", err)
	}

	dot, err := dotsql.LoadFromString(string(sqlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to load SQL: %w", err)
	}

	queryNames, err := GetQueryNamesForService(service)
	if err != nil {
		return nil, err
	}

	queries := map[string]string{}
	for _, name := range queryNames {
		query, err := dot.Raw(name)
		if err != nil {
			return nil, fmt.Errorf("failed to get query '%s': %w", name, err)
		}

		queries[name] = cleanQuery(query)
	}

	fmt.Printf("successfully load queries using lib version '%s'", VERSION)

	return queries, nil
}

// GetQueryNamesForService returns the appropriate query names for a given service.
// It returns an error if an unrecognized service is provided.
func GetQueryNamesForService(service ServiceName) ([]string, error) {
	switch service {
	case SvcFundAccount:
		return registration.QUERY_NAMES_FUNDACCOUNT, nil
	case SvcFundFeat:
		return registration.QUERY_NAMES_FUNDFEAT, nil
	case SvcFundParam:
		return registration.QUERY_NAMES_FUNDPARAM, nil
	case SvcFundReport:
		return registration.QUERY_NAMES_FUNDREPORT, nil
	case SvcFundTimeDep:
		return registration.QUERY_NAMES_FUNDTIMEDEP, nil
	case SvcFundTx:
		return registration.QUERY_NAMES_FUNDTX, nil
	case SvcFundTxBatch:
		return registration.QUERY_NAMES_FUNDTXBATCH, nil
	case SvcFundTxOnline:
		return registration.QUERY_NAMES_FUNDTXONLINE, nil
	case SvcFundTxRemit:
		return registration.QUERY_NAMES_FUNDTXREMIT, nil
	case SvcFull:
		return registration.QUERY_NAMES_ALL, nil
	default:
		return nil, ErrUnrecognizedService
	}
}

// GetRandomQueryName returns a random query name from the predefined list of query names.
// This function can be used to select a random query for testing or demonstration purposes.
//
// Returns:
//   - A string representing a randomly selected query name
func GetRandomQueryName() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return registration.QUERY_NAMES_ALL[rand.Intn(len(registration.QUERY_NAMES_ALL))]
}

// cleanQuery sanitizes SQL queries by removing trailing semicolons and normalizing whitespace.
// It helps prevent ORA-00911 errors when using the go-ora driver.
//
// The function performs the following operations:
//   - Removes trailing semicolons
//   - Compresses multiple whitespace characters into a single space
//   - Trims leading and trailing whitespace
//
// Parameters:
//   - query: The SQL query string to be cleaned
//
// Returns:
//   - A cleaned version of the input query
func cleanQuery(query string) string {
	// Remove trailing semicolon
	query = strings.TrimSuffix(query, ";")

	// Optional: compress multiple spaces/newlines
	space := regexp.MustCompile(`\s+`)
	query = space.ReplaceAllString(query, " ")

	return strings.TrimSpace(query)
}
