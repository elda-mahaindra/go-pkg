package queries

import (
	"embed"
	"fmt"

	"github.com/qustavo/dotsql"
)

//go:embed oracle.sql postgres.sql
var sqlFiles embed.FS

func LoadQueries(driver string) (map[string]string, error) {
	filename := SQL_FILES[driver]
	sqlContent, err := sqlFiles.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file: %w", err)
	}

	dot, err := dotsql.LoadFromString(string(sqlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to load SQL: %w", err)
	}

	queries := map[string]string{}

	for _, name := range QUERY_NAMES {
		query, err := dot.Raw(name)
		if err != nil {
			return nil, fmt.Errorf("failed to get query '%s': %w", name, err)
		}

		queries[name] = query
	}

	return queries, nil
}

func LoadMeIn() {
	fmt.Println("Hello World!")
}

var SQL_FILES = map[string]string{
	"godror":     "oracle.sql",
	"postgresql": "postgres.sql",
}

const (
	QUERY__GET_GL_LIST                   = "get-gl-list"
	QUERY__GET_HIST_TRX_DETAIL_BY_UIDKEY = "get-hist-trx-detail-by-uidkey"
	QUERY__GET_VALIDATION_LIMIT_TRX      = "get-validation-limit-trx"
	QUERY__GET_REKENING_LIAB             = "get-rekening-liab"
	QUERY__GET_MAX_DB_HARIAN             = "get-max-db-harian"
	QUERY__GET_LIMIT_DEBET_HARIAN        = "get-limit-debet-harian"
	QUERY__GET_COUNTER_DEBET_HARIAN      = "get-counter-debet-harian"
	QUERY__GET_RC_CODE_LIST              = "get-rc-code-list"
	QUERY__GET_POD                       = "get-pod"
	QUERY__GET_SALDO_BY_DATE             = "get-saldo-by-date"
	QUERY__GET_HIST_TRX_LIST_BY_NOREK    = "get-hist-trx-list-by-norek"
	QUERY__GET_TRX_LIST_BY_NOREK         = "get-trx-list-by-norek"
	QUERY__GET_TRX_DETAIL                = "get-trx-detail"
	QUERY__GET_TRX_DETAIL_BY_UIDKEY      = "get-trx-detail-by-uidkey"
	QUERY__GET_TRX_LIST                  = "get-trx-list"
	QUERY__DETAIL_GL_JURNAL              = "detail-gl-jurnal"
	QUERY__DETAIL_GL_JURNAL_ITEM         = "detail-gl-jurnal-item"
	QUERY__DETAIL_GL_JURNAL_BLOCK        = "detail-gl-jurnal-block"
	QUERY__GET_GL_JURNAL                 = "get-gl-jurnal"
	QUERY__GET_INFO_UID                  = "get-info-uid"
	QUERY__GET_TRANSAKSI_ID              = "get-transaksi-id"
)

var QUERY_NAMES = []string{
	QUERY__GET_GL_LIST,
	QUERY__GET_HIST_TRX_DETAIL_BY_UIDKEY,
	QUERY__GET_VALIDATION_LIMIT_TRX,
	QUERY__GET_REKENING_LIAB,
	QUERY__GET_MAX_DB_HARIAN,
	QUERY__GET_LIMIT_DEBET_HARIAN,
	QUERY__GET_COUNTER_DEBET_HARIAN,
	QUERY__GET_RC_CODE_LIST,
	QUERY__GET_POD,
	QUERY__GET_SALDO_BY_DATE,
	QUERY__GET_HIST_TRX_LIST_BY_NOREK,
	QUERY__GET_TRX_LIST_BY_NOREK,
	QUERY__GET_TRX_DETAIL,
	QUERY__GET_TRX_DETAIL_BY_UIDKEY,
	QUERY__GET_TRX_LIST,
	QUERY__DETAIL_GL_JURNAL,
	QUERY__DETAIL_GL_JURNAL_ITEM,
	QUERY__DETAIL_GL_JURNAL_BLOCK,
	QUERY__GET_GL_JURNAL,
	QUERY__GET_INFO_UID,
	QUERY__GET_TRANSAKSI_ID,
}
