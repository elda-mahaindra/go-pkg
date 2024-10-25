package data

import "database/sql"

type ListTrxPayrollRes struct {
	TanggalTransaksi string  `db:"TANGGAL_TRANSAKSI"`
	UserInput        string  `db:"USER_INPUT"`
	Keterangan       string  `db:"KETERANGAN"`
	NomorReferensi   string  `db:"NOMOR_REFERENSI"`
	IdTransaksi      int64   `db:"ID_TRANSAKSI"`
	KodeTransaksi    string  `db:"KODE_TRANSAKSI"`
	TotalNilaiMutasi float64 `db:"TOTAL_NILAI_MUTASI"`
	JumlahRekening   int32   `db:"JUMLAH_REKENING"`
}

type InfoUid struct {
	Uidkey       sql.NullString `db:"UIDKEY" json:"uidkey"`
	TanggalInput sql.NullString `db:"TANGGAL_INPUT" json:"tanggal_input"`
	ActionType   sql.NullString `db:"ACTION_TYPE" json:"action_type"`
	UserInput    sql.NullString `db:"USER_INPUT" json:"user_input"`
	UserOtor     sql.NullString `db:"USER_OTOR" json:"user_otor"`
	Keterangan   sql.NullString `db:"KETERANGAN" json:"keterangan"`
	Info1        sql.NullString `db:"INFO1" json:"info1"`
	Info2        sql.NullString `db:"INFO2" json:"info2"`
	KodeCabang   sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
	Keyint       sql.NullInt64  `db:"KEYINT" json:"keyint"`
	IsExists     sql.NullString `db:"IS_EXISTS" json:"is_exists"`
	IsReversed   sql.NullString `db:"IS_REVERSED" json:"is_reversed"`
}

type CheckBalanceMultiJournalTableTemp struct {
	JenisMutasi  string  `db:"JENIS_MUTASI"`
	TotalNominal float64 `db:"TOTAL_NOMINAL"`
}

type PostTrxMultiJournalInsertToTrxTable struct {
	KodeTipe      string  `json:"kodeTipe" db:"KODE_TIPE"`
	NomorRekening string  `json:"nomorRekening" db:"NOMOR_REKENING"`
	JenisMutasi   string  `json:"jenisMutasi" db:"JENIS_MUTASI"`
	Nominal       float64 `json:"nominal" db:"NOMINAL"`
	KodeCabang    string  `json:"kodeCabang" db:"KODE_CABANG"`
	Keterangan    string  `json:"keterangan" db:"KETERANGAN"`
	KodeAccount   string  `json:"kodeAccount" db:"KODE_ACCOUNT"`
	KodeValuta    string  `json:"kodeValuta" db:"KODE_VALUTA"`
}

type DataTrxMulti struct {
	KodeTipe       string  `db:"KODE_TIPE"`
	NomorRekening  string  `db:"NOMOR_REKENING"`
	NamaRekening   string  `db:"NAMA_REKENING"`
	KodeCabang     string  `db:"KODE_CABANG"`
	JenisMutasi    string  `db:"JENIS_MUTASI"`
	Nominal        float64 `db:"NOMINAL"`
	Keterangan     string  `db:"KETERANGAN"`
	KodeRc         string  `db:"KODE_RC"`
	IsValid        string  `db:"IS_VALID"`
	InvalidMessage string  `db:"INVALID_MESSAGE"`
}

type TotalNominalTmpTrxMulti struct {
	JenisMutasi  string  `db:"JENIS_MUTASI"`
	TotalNominal float64 `db:"TOTAL_NOMINAL"`
}

type ValTrxPayrollR struct {
	NomorRekening  string  `db:"NOMOR_REKENING"`
	NamaRekening   string  `db:"NAMA_REKENING"`
	Nominal        float64 `db:"NOMINAL"`
	Keterangan     string  `db:"KETERANGAN"`
	IsValid        string  `db:"IS_VALID"`
	InvalidMessage string  `db:"INVALID_MESSAGE"`
}

type ValTrxUpFundQuery struct {
	IdSession      int64           `db:"ID_SESSION"`
	NomorRekening  sql.NullString  `db:"NOMOR_REKENING"`
	NamaRekening   sql.NullString  `db:"NAMA_REKENING"`
	Keterangan     sql.NullString  `db:"KETERANGAN"`
	Nominal        sql.NullFloat64 `db:"NOMINAL"`
	NominalSign    sql.NullString  `db:"NOMINAL_SIGN"`
	JenisMutasi    sql.NullString  `db:"JENIS_MUTASI"`
	TransCode      sql.NullString  `db:"TRANS_CODE"`
	IsValid        sql.NullString  `db:"IS_VALID"`
	InvalidMessage sql.NullString  `db:"INVALID_MESSAGE"`
}

type GetUploadGLRes struct {
	IdSession        sql.NullInt32   `db:"ID_SESSION"`
	TanggalInput     sql.NullString  `db:"TANGGAL_INPUT" json:"tanggal_input"`
	TanggalPosting   sql.NullString  `db:"TANGGAL_POSTING" json:"tanggal_posting"`
	KodeCabang       sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	KodeAccount      sql.NullString  `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeValuta       sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	JenisMutasi      sql.NullString  `db:"JENIS_MUTASI" json:"jenis_mutasi"`
	NomorReferensi   sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	Nominal          sql.NullFloat64 `db:"NOMINAL" json:"nominal"`
	NominalEkuivalen sql.NullFloat64 `db:"NOMINAL_EKUIVALEN" json:"nominal_ekuivalen"`
	KodeArea         sql.NullString  `db:"KODE_AREA" json:"kode_area"`
	KodePair         sql.NullString  `db:"KODE_PAIR" json:"kode_pair"`
	Description      sql.NullString  `db:"DESCRIPTION" json:"description"`
	IsValid          sql.NullString  `db:"IS_VALID" json:"is_valid"`
	InvalidMessage   sql.NullString  `db:"INVALID_MESSAGE" json:"invalid_message"`
}

type IdTransaksiData struct {
	IdTransaksi int64 `db:"ID_TRANSAKSI" json:"id_transaksi"`
}

type Pod struct {
	Today string `db:"TODAY" json:"today,"`
}

type CustomIdgen struct {
	LastIdGen int32  `db:"LAST_IDGEN" json:"last_idgen"`
	Idname    string `db:"IDNAME" json:"idname"`
	Idcode    string `db:"IDCODE" json:"idcode"`
	LastId    int32  `db:"LAST_ID" json:"last_id"`
	Locked    string `db:"LOCKED" json:"locked"`
}

type RekKasTrx struct {
	NomorRekening     sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening      sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeValuta        sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabang        sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	Saldo             sql.NullFloat64 `db:"SALDO" json:"saldo"`
	AccountinstanceId sql.NullInt64   `db:"ACCOUNTINSTANCE_ID" json:"accountinstance_id"`
	KodeAccount       sql.NullString  `db:"KODE_ACCOUNT" json:"kode_account"`
	SaldoEfektif      sql.NullFloat64 `db:"SALDO_EFEKTIF" json:"saldo_efektif"`
	StatusRekening    sql.NullInt32   `db:"STATUS_REKENING" json:"status_rekening"`
}

type RekGLTrx struct {
	NomorRekening     sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening      sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeValuta        sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabang        sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	Saldo             sql.NullFloat64 `db:"SALDO" json:"saldo"`
	AccountinstanceId sql.NullInt64   `db:"ACCOUNTINSTANCE_ID" json:"accountinstance_id"`
	KodeAccount       sql.NullString  `db:"KODE_ACCOUNT" json:"kode_account"`
	NormalBalanceType sql.NullString  `db:"NORMAL_BALANCE_TYPE" json:"normal_balance_type"`
	TrxPermitType     sql.NullString  `db:"TRX_PERMIT_TYPE" json:"trx_permit_type"`
	BalanceSign       sql.NullFloat64 `db:"BALANCE_SIGN" json:"balance_sign"`
	Balance           sql.NullFloat64 `db:"BALANCE" json:"balance"`
}

type RekKodeInterfaceTrx struct {
	NomorRekening     sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening      sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeValuta        sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabang        sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	Saldo             sql.NullFloat64 `db:"SALDO" json:"saldo"`
	AccountinstanceId sql.NullInt64   `db:"ACCOUNTINSTANCE_ID" json:"accountinstance_id"`
	KodeAccount       sql.NullString  `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeInterface     sql.NullString  `db:"KODE_INTERFACE" json:"kode_interface"`
}

type RekLiabTrx struct {
	NomorRekening     sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	KodeValuta        sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabang        sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	Saldo             sql.NullFloat64 `db:"SALDO" json:"saldo"`
	AccountinstanceId sql.NullInt64   `db:"ACCOUNTINSTANCE_ID" json:"accountinstance_id"`
	KodeAccount       sql.NullString  `db:"KODE_ACCOUNT" json:"kode_account"`
	SaldoEfektif      sql.NullFloat64 `db:"SALDO_EFEKTIF" json:"saldo_efektif"`
	StatusRekening    sql.NullInt32   `db:"STATUS_REKENING" json:"status_rekening"`
}

type GetGlobalParamResult struct {
	KodeParameter         sql.NullString  `db:"KODE_PARAMETER"`
	TipeParameter         sql.NullString  `db:"TIPE_PARAMETER"`
	NilaiParameter        sql.NullFloat64 `db:"NILAI_PARAMETER"`
	Deskripsi             sql.NullString  `db:"DESKRIPSI"`
	IsParameterSystem     sql.NullString  `db:"IS_PARAMETER_SYSTEM"`
	NilaiParameterTanggal sql.NullString  `db:"NILAI_PARAMETER_TANGGAL"`
	NilaiParameterString  sql.NullString  `db:"NILAI_PARAMETER_STRING"`
	Locked                sql.NullString  `db:"LOCKED"`
}

type GlobalGlInterface struct {
	KodeAccount   string `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeInterface string `db:"KODE_INTERFACE" json:"kode_interface"`
	NamaAccount   string `db:"NAMA_ACCOUNT" json:"nama_account"`
}

type AccountingDayAvabilityRes struct {
	Datevalue     string `db:"DATEVALUE"`
	PeriodeStatus string `db:"PERIODE_STATUS"`
}

type DsValiditasDataJournal struct {
	IdDetilTransaksi sql.NullInt64   `db:"ID_DETIL_TRANSAKSI"`
	IdTransaksi      sql.NullInt64   `db:"ID_TRANSAKSI"`
	NomorRekening    sql.NullString  `db:"NOMOR_REKENING"`
	JenisMutasi      sql.NullString  `db:"JENIS_MUTASI"`
	NilaiMutasi      sql.NullFloat64 `db:"NILAI_MUTASI"`
	KodeTxClass      sql.NullString  `db:"KODE_TX_CLASS"`
	NomorSeri        sql.NullString  `db:"NOMOR_SERI"`
}

type DsCheckAccountInstanceJournal struct {
	IdDetilTransaksi sql.NullInt64  `db:"ID_DETIL_TRANSAKSI"`
	IdTransaksi      sql.NullInt64  `db:"ID_TRANSAKSI"`
	NomorRekening    sql.NullString `db:"NOMOR_REKENING"`
	KodeAccount      sql.NullString `db:"KODE_ACCOUNT"`
	KodeCabang       sql.NullString `db:"KODE_CABANG"`
	KodeValuta       sql.NullString `db:"KODE_VALUTA"`
}

type DsCheckBalancingJournal struct {
	IdTransaksi sql.NullInt64  `db:"ID_TRANSAKSI"`
	KodeValuta  sql.NullString `db:"KODE_VALUTA"`
}
