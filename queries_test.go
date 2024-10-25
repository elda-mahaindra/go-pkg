package package_query_test

import (
	"testing"

	pquery "github.com/elda-mahaindra/go-pkg"
	"github.com/stretchr/testify/require"
)

// ================================ tests ================================
func TestLoadQueries(t *testing.T) {
	t.Parallel()

	testCases := getTestLoadQueriesTestCases()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			queries, err := pquery.LoadQueries(tc.driver, tc.service)

			switch tc.name {
			case "ok_oracle_fundaccount":
			case "ok_oracle_fundfeat":
			case "ok_oracle_fundparam":
			case "ok_oracle_fundreport":
			case "ok_oracle_fundtimedep":
			case "ok_oracle_fundtx":
			case "ok_oracle_fundtxbatch":
			case "ok_oracle_fundtxonline":
			case "ok_oracle_fundtxremit":
			case "ok_oracle_full":
				{
					require.NoError(t, err)
					require.NotNil(t, queries)

					queryNames, err := pquery.GetQueryNamesForService(tc.service)
					require.NoError(t, err)
					require.NotNil(t, queryNames)

					require.Equal(t, len(queryNames), len(queries))
					for _, name := range queryNames {
						_, ok := queries[name]
						require.True(t, ok)
					}

				}
			case "nok_unrecognized_driver":
				{
					require.Error(t, err)
					require.Nil(t, queries)
					require.ErrorIs(t, err, pquery.ErrUnrecognizedDriver)
				}
			case "nok_unrecognized_service":
				{
					require.Error(t, err)
					require.Nil(t, queries)
					require.ErrorIs(t, err, pquery.ErrUnrecognizedService)
				}
			default:
				t.Fatalf("Unexpected test case: %s", tc.name)
			}
		})
	}
}

// ================================ utilities ================================
func getTestLoadQueriesTestCases() []struct {
	name    string
	driver  pquery.DriverName
	service pquery.ServiceName
} {
	testCases := []struct {
		name    string
		driver  pquery.DriverName
		service pquery.ServiceName
	}{
		{
			name:    "ok_oracle_fundaccount",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundAccount,
		},
		{
			name:    "ok_oracle_fundfeat",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundFeat,
		},
		{
			name:    "ok_oracle_fundparam",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundParam,
		},
		{
			name:    "ok_oracle_fundreport",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundReport,
		},
		{
			name:    "ok_oracle_fundtimedep",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundTimeDep,
		},
		{
			name:    "ok_oracle_fundtx",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundTx,
		},
		{
			name:    "ok_oracle_fundtxbatch",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundTxBatch,
		},
		{
			name:    "ok_oracle_fundtxonline",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundTxOnline,
		},
		{
			name:    "ok_oracle_fundtxremit",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFundTxRemit,
		},
		{
			name:    "ok_oracle_full",
			driver:  pquery.DriverOracle,
			service: pquery.SvcFull,
		},
		{
			name:    "nok_unrecognized_driver",
			driver:  "wrong-driver",
			service: pquery.SvcFundTxOnline,
		},
		{
			name:    "nok_unrecognized_service",
			driver:  pquery.DriverOracle,
			service: "wrong-service",
		},
	}

	return testCases
}
