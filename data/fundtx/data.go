package data

import (
	"database/sql"
	"time"
)

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

type BalanceToday struct {
	TotalToday sql.NullFloat64 `db:"TOTAL_TODAY" json:"total_today"`
}

type DsChequeBySerialNum struct {
	IdBukuwarkat     sql.NullInt64  `db:"ID_BUKUWARKAT" json:"id_bukuwarkat"`
	StatusWarkat     sql.NullString `db:"STATUS_WARKAT" json:"status_warkat"`
	StatusWarkatDesc sql.NullString `db:"STATUS_WARKAT_DESC" json:"status_warkat_desc"`
	JenisWarkat      sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	JenisWarkatDesc  sql.NullString `db:"JENIS_WARKAT_DESC" json:"jenis_warkat_desc"`
	StatusBuku       sql.NullString `db:"STATUS_BUKU" json:"status_buku"`
	NomorRekening    sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorSeri        sql.NullString `db:"NOMOR_SERI" json:"nomor_seri"`
	NamaRekening     sql.NullString `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeCabang       sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
	KodeProduk       sql.NullString `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk       sql.NullString `db:"NAMA_PRODUK" json:"nama_produk"`
	KodeValuta       sql.NullString `db:"KODE_VALUTA" json:"kode_valuta"`
	NamaCabang       sql.NullString `db:"NAMA_CABANG" json:"nama_cabang"`
}

type PostNotaDebetOutwardDeleted struct {
	JumlahTerhapus int32 `db:"JUMLAH_TERHAPUS" json:"jumlah_terhapus"`
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

type Pod struct {
	Today string `db:"TODAY" json:"today,"`
}

type IdTransaksiData struct {
	IdTransaksi int64 `db:"ID_TRANSAKSI" json:"id_transaksi"`
}

type CustomIdgen struct {
	LastIdGen int32  `db:"LAST_IDGEN" json:"last_idgen"`
	Idname    string `db:"IDNAME" json:"idname"`
	Idcode    string `db:"IDCODE" json:"idcode"`
	LastId    int32  `db:"LAST_ID" json:"last_id"`
	Locked    string `db:"LOCKED" json:"locked"`
}

type GlobalGlInterface struct {
	KodeAccount   string `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeInterface string `db:"KODE_INTERFACE" json:"kode_interface"`
	NamaAccount   string `db:"NAMA_ACCOUNT" json:"nama_account"`
}

type ResGetListGl struct {
	AccountCode sql.NullString `db:"ACCOUNT_CODE" json:"account_code"`
	AccountName sql.NullString `db:"ACCOUNT_NAME" json:"account_name"`
	AccountType sql.NullString `db:"ACCOUNT_TYPE" json:"account_type"`
}

type GLJurnalGetData struct {
	JournalNo               sql.NullString `db:"JOURNAL_NO" json:"journal_no,omitempty"`
	SerialNo                sql.NullString `db:"SERIAL_NO" json:"serial_no,omitempty"`
	IsPosted                sql.NullString `db:"IS_POSTED" json:"is_posted,omitempty"`
	JournalState            sql.NullString `db:"JOURNAL_STATE" json:"journal_state,omitempty"`
	TransactionDate         sql.NullString `db:"TRANSACTION_DATE" json:"transaction_date,omitempty"`
	JournalDate             sql.NullString `db:"JOURNAL_DATE" json:"journal_date,omitempty"`
	JournalDatePosting      sql.NullString `db:"JOURNAL_DATE_POSTING" json:"journal_date_posting,omitempty"`
	JournalDateLastModified sql.NullString `db:"JOURNAL_DATE_LAST_MODIFIED" json:"journal_date_last_modified,omitempty"`
	BranchCode              sql.NullString `db:"BRANCH_CODE" json:"branch_code,omitempty"`
	Description             sql.NullString `db:"DESCRIPTION" json:"description,omitempty"`
	UseridCreate            sql.NullString `db:"USERID_CREATE" json:"userid_create,omitempty"`
	UseridPosting           sql.NullString `db:"USERID_POSTING" json:"userid_posting,omitempty"`
	UseridLastModified      sql.NullString `db:"USERID_LAST_MODIFIED" json:"userid_last_modified,omitempty"`
}

type GLJurnalDetailData struct {
	JournalNo               sql.NullString `db:"JOURNAL_NO" json:"journal_no,omitempty"`
	TransactionDate         sql.NullString `db:"TRANSACTION_DATE" json:"transaction_date,omitempty"`
	JournalDate             sql.NullString `db:"JOURNAL_DATE" json:"journal_date,omitempty"`
	JournalDatePosting      sql.NullString `db:"JOURNAL_DATE_POSTING" json:"journal_date_posting,omitempty"`
	JournalDateLastModified sql.NullString `db:"JOURNAL_DATE_LAST_MODIFIED" json:"journal_date_last_modified,omitempty"`
	BranchCode              sql.NullString `db:"BRANCH_CODE" json:"branch_code,omitempty"`
	Description             sql.NullString `db:"DESCRIPTION" json:"description,omitempty"`
	UseridCreate            sql.NullString `db:"USERID_CREATE" json:"userid_create,omitempty"`
	UseridPosting           sql.NullString `db:"USERID_POSTING" json:"userid_posting,omitempty"`
	IsPosted                sql.NullString `db:"IS_POSTED" json:"is_posted,omitempty"`
	JournalState            sql.NullString `db:"JOURNAL_STATE" json:"journal_state,omitempty"`
}

type GLJurnalItem struct {
	Nomor          *int32         `db:"NOMOR" json:"nomor,omitempty"`
	KodeRekening   sql.NullString `db:"KODE_REKENING" json:"kode_rekening,omitempty"`
	NamaRekening   sql.NullString `db:"NAMA_REKENING" json:"nama_rekening,omitempty"`
	KodeKantor     sql.NullString `db:"KODE_KANTOR" json:"kode_kantor,omitempty"`
	NamaKantor     sql.NullString `db:"NAMA_KANTOR" json:"nama_kantor,omitempty"`
	Valuta         sql.NullString `db:"VALUTA" json:"valuta,omitempty"`
	KodeKurs       sql.NullString `db:"KODE_KURS" json:"kode_kurs,omitempty"`
	NilaiKurs      *float64       `db:"NILAI_KURS" json:"nilai_kurs,omitempty"`
	NamaRc         sql.NullString `db:"NAMA_RC" json:"nama_rc,omitempty"`
	NoBlokJurnal   sql.NullString `db:"NO_BLOK_JURNAL" json:"no_blok_jurnal,omitempty"`
	Subsystemcode  sql.NullString `db:"SUBSYSTEMCODE" json:"subsystemcode,omitempty"`
	Subaccountcode sql.NullString `db:"SUBACCOUNTCODE" json:"subaccountcode,omitempty"`
	Subaccountname sql.NullString `db:"SUBACCOUNTNAME" json:"subaccountname,omitempty"`
	Debet          *float64       `db:"DEBET" json:"debet,omitempty"`
	Kredit         *float64       `db:"KREDIT" json:"kredit,omitempty"`
	UserIdCreate   sql.NullString `db:"USER_ID_CREATE" json:"user_id_create,omitempty"`
	Keterangan     sql.NullString `db:"KETERANGAN" json:"keterangan,omitempty"`
}

type GLJurnalBlock struct {
	KodeSubSistem  sql.NullString `db:"KODE_SUB_SISTEM" json:"kode_sub_sistem,omitempty"`
	ClassId        sql.NullString `db:"CLASS_ID" json:"class_id,omitempty"`
	KeyId          *int32         `db:"KEY_ID" json:"key_id,omitempty"`
	IdJournalblock *int32         `db:"ID_JOURNALBLOCK" json:"id_journalblock,omitempty"`
	Keterangan     sql.NullString `db:"KETERANGAN" json:"keterangan,omitempty"`
}

type InfoPenutupanRekeningView struct {
	Nomor_rekening  sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	Id_transaksi    sql.NullInt64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	Tanggal_proses  sql.NullString  `db:"TANGGAL_PROSES" json:"tanggal_proses"`
	Nilai_transaksi sql.NullFloat64 `db:"NILAI_TRANSAKSI" json:"nilai_transaksi"`
	Keterangan      sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	User_otorisasi  sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	Userotorisasi   sql.NullString  `db:"USEROTORISASI" json:"userotorisasi"`
	Userinput       sql.NullString  `db:"USERINPUT" json:"userinput"`
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

type DenominasitransaksikasDetail struct {
	IdDenominasitransaksikas sql.NullInt64 `db:"ID_DENOMINASITRANSAKSIKAS" json:"id_denominasitransaksikas"`
	IdDenominasiValuta       sql.NullInt64 `db:"ID_DENOMINASI_VALUTA" json:"id_denominasi_valuta"`
	IdDetilTransaksi         sql.NullInt64 `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	JumlahDenominasi         sql.NullInt64 `db:"JUMLAH_DENOMINASI" json:"jumlah_denominasi"`
	JumlahDenominasiTle      sql.NullInt64 `db:"JUMLAH_DENOMINASI_TLE" json:"jumlah_denominasi_tle"`
}

type DenominasiKas struct {
	JumlahDenominasi    int64   `db:"JUMLAH_DENOMINASI"`
	IdDenominasiValuta  int64   `db:"ID_DENOMINASI_VALUTA"`
	NomorRekening       string  `db:"NOMOR_REKENING"`
	IdDenominasikas     int64   `db:"ID_DENOMINASIKAS"`
	JumlahDenominasiTle int64   `db:"JUMLAH_DENOMINASI_TLE"`
	NilaiDenominasi     float64 `db:"NILAI_DENOMINASI"`
	KodeValuta          string  `db:"KODE_VALUTA"`
	JenisDenominasi     string  `db:"JENIS_DENOMINASI"`
}

type DtkliringoutwardDetail struct {
	IdDetilTransaksi   sql.NullInt64  `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	IdNotadebitoutward sql.NullInt64  `db:"ID_NOTADEBITOUTWARD" json:"id_notadebitoutward"`
	IdWic              sql.NullInt64  `db:"ID_WIC" json:"id_wic"`
	NomorRekening      sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
}

type Dtkliringtarikan struct {
	IdDetilTransaksi sql.NullInt64  `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	IdDetilTolakan   sql.NullInt64  `db:"ID_DETIL_TOLAKAN" json:"id_detil_tolakan"`
	JenisWarkat      sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	NomorWarkat      sql.NullString `db:"NOMOR_WARKAT" json:"nomor_warkat"`
	NomorRekening    sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
}

type NotaDebetInternalDetail struct {
	NomorSeri              sql.NullString `db:"NOMOR_SERI" json:"nomor_seri"`
	JenisWarkat            sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	TanggalPakaiWarkat     sql.NullString `db:"TANGGAL_PAKAI_WARKAT" json:"tanggal_pakai_warkat"`
	NoSeriDetilTransaksi   sql.NullString `db:"NO_SERI_DETIL_TRANSAKSI" json:"no_seri_detil_transaksi"`
	IdNotadebetinternal    sql.NullInt64  `db:"ID_NOTADEBETINTERNAL" json:"id_notadebetinternal"`
	StatusWarkat           sql.NullString `db:"STATUS_WARKAT" json:"status_warkat"`
	StatusWarkatDesc       sql.NullString `db:"STATUS_WARKAT_DESC" json:"status_warkat_desc"`
	TanggalPerubahanStatus sql.NullString `db:"TANGGAL_PERUBAHAN_STATUS" json:"tanggal_perubahan_status"`
	UserPengubah           sql.NullString `db:"USER_PENGUBAH" json:"user_pengubah"`
	IdBukuwarkat           sql.NullInt64  `db:"ID_BUKUWARKAT" json:"id_bukuwarkat"`
	MNomrek                sql.NullString `db:"M_NOMREK" json:"m_nomrek"`
	NomorRekening          sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	Used                   sql.NullString `db:"USED" json:"used"`
	UsedCounter            sql.NullString `db:"USED_COUNTER" json:"used_counter"`
}

type DsLimitTrx struct {
	KodeTransaksi       sql.NullString  `db:"KODE_TRANSAKSI"`
	Keterangan          sql.NullString  `db:"KETERANGAN"`
	JenisLimitTransaksi sql.NullString  `db:"JENIS_LIMIT_TRANSAKSI"`
	IdUser              sql.NullString  `db:"ID_USER"`
	NilaiLimit          sql.NullFloat64 `db:"NILAI_LIMIT"`
}

type DsMaxDbHarian struct {
	TotalLimit sql.NullInt64 `db:"TOTAL_LIMIT"`
}

type DsLimitDbHarian struct {
	TotalNilaiTransaksi sql.NullFloat64 `db:"TOTAL_NILAI_TRANSAKSI"`
}

type DsCounterDbHarian struct {
	TotalNilaiTransaksi sql.NullFloat64 `db:"TOTAL_NILAI_TRANSAKSI"`
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

type ResGetRcCodeList struct {
	RcCode sql.NullString `db:"RC_CODE"`
	RcName sql.NullString `db:"RC_NAME"`
}

type RekeningLiabDetail struct {
	KodeCabang                  sql.NullString     `db:"KODE_CABANG" json:"kode_cabang"`
	UserInput                   sql.NullString     `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi               sql.NullString     `db:"USER_OTORISASI" json:"user_otorisasi"`
	AlamatStatement             sql.NullString     `db:"ALAMAT_KIRIM" json:"alamat_statement"`
	NomorNasabah                sql.NullString     `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NomorRekening               sql.NullString     `db:"NOMOR_REKENING" json:"nomor_rekening"`
	TanggalBuka                 sql.NullString     `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	StatusRekening              sql.NullInt32      `db:"STATUS_REKENING" json:"status_rekening"`
	StatusRekeningDesc          sql.NullString     `db:"STATUS_REKENING_DESC" json:"status_rekening_desc"`
	RelasiMultiCif              sql.NullString     `db:"KODE_MULTICIFRELATION" json:"relasi_multi_cif"`
	IsJoinAccount               sql.NullString     `db:"IS_JOIN_ACCOUNT" json:"is_join_account"`
	JenisRekeningLiabilitas     sql.NullString     `db:"JENIS_REKENING_LIABILITAS" json:"jenis_rekening_liabilitas"`
	NamaRekening                sql.NullString     `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeProduk                  sql.NullString     `db:"KODE_PRODUK" json:"kode_produk"`
	KodeProdukDesc              sql.NullString     `db:"KODE_PRODUK_DESC" json:"kode_produk_desc"`
	NisbahSpesial               sql.NullFloat64    `db:"NISBAH_SPESIAL" json:"nisbah_spesial"`
	NisbahDasar                 sql.NullFloat64    `db:"NISBAH_DASAR" json:"nisbah_dasar"`
	TarifPajak                  sql.NullFloat64    `db:"TARIF_PAJAK" json:"tarif_pajak"`
	PersentaseZakatBagiHasil    sql.NullFloat64    `db:"PERSENTASE_ZAKAT_BAGI_HASIL" json:"persentase_zakat_bagi_hasil"`
	KodeSumberDana              sql.NullString     `db:"KODE_SUMBER_DANA" json:"kode_sumber_dana"`
	KodeTujuanRekening          sql.NullString     `db:"KODE_TUJUAN_REKENING" json:"kode_tujuan_rekening"`
	KodeMarketingCurrent        sql.NullString     `db:"KODE_MARKETING_CURRENT" json:"kode_marketing_current"`
	KodeMarketingPertama        sql.NullString     `db:"KODE_MARKETING_PERTAMA" json:"kode_marketing_pertama"`
	JenisRekeningLiabilitasDesc sql.NullString     `db:"JENIS_REKENING_LIABILITAS_DESC" json:"jenis_rekening_liabilitas_desc"`
	IsKenaBiayalayananumum      sql.NullString     `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum"`
	IsStatusPassbook            sql.NullString     `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook"`
	Saldo                       sql.NullFloat64    `db:"SALDO" json:"saldo"`
	SaldoFloat                  sql.NullFloat64    `db:"SALDO_FLOAT" json:"saldo_float"`
	SaldoDeposito               sql.NullFloat64    `db:"SALDO_DEPOSITO" json:"saldo_deposito"`
	SaldoDitahan                sql.NullFloat64    `db:"SALDO_DITAHAN" json:"saldo_ditahan"`
	SaldoMinimum                sql.NullFloat64    `db:"SALDO_MINIMUM" json:"saldo_minimum"`
	SaldoTunggakan              sql.NullFloat64    `db:"SALDO_TUNGGAKAN" json:"saldo_tunggakan"`
	KodeValuta                  sql.NullString     `db:"KODE_VALUTA" json:"kode_valuta"`
	NamaValuta                  sql.NullString     `db:"NAMA_VALUTA" json:"nama_valuta"`
	Keterangan                  sql.NullString     `db:"KETERANGAN" json:"keterangan"`
	IsSedangDitutup             sql.NullString     `db:"IS_SEDANG_DITUTUP" json:"is_sedang_ditutup"`
	KodeProgram                 sql.NullString     `db:"KODE_PROGRAM" json:"kode_program"`
	KodeMarketingReferensi      sql.NullString     `db:"KODE_MARKETING_REFERENSI" json:"kode_marketing_referensi"`
	StatusKelengkapan           sql.NullString     `db:"STATUS_KELENGKAPAN" json:"status_kelengkapan"`
	IsBlokir                    sql.NullString     `db:"IS_BLOKIR" json:"is_blokir"`
	IsBlokirDebet               sql.NullString     `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	IsBolehDebet                sql.NullString     `db:"IS_BOLEH_DEBET" json:"is_boleh_debet"`
	IsBlokirKredit              sql.NullString     `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsBolehKredit               sql.NullString     `db:"IS_BOLEH_KREDIT" json:"is_boleh_kredit"`
	IsCetakNota                 sql.NullString     `db:"IS_CETAK_NOTA" json:"is_cetak_nota"`
	IsDapatBagiHasil            sql.NullString     `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	IsTidakDormant              sql.NullString     `db:"IS_TIDAK_DORMANT" json:"is_tidak_dormant"`
	IsBiayaRekeningDormant      sql.NullString     `db:"IS_BIAYA_REKENING_DORMANT" json:"is_biaya_rekening_dormant"`
	IsBiayaSaldoMinimum         sql.NullString     `db:"IS_BIAYA_SALDO_MINIMUM" json:"is_biaya_saldo_minimum"`
	IsBiayaAtm                  sql.NullString     `db:"IS_BIAYA_ATM" json:"is_biaya_atm"`
	NisbahBagiHasil             sql.NullFloat64    `db:"NISBAH_BAGI_HASIL" json:"nisbah_bagi_hasil"`
	IsTieringNisbahbonus        sql.NullString     `db:"IS_TIERING_NISBAHBONUS" json:"is_tiering_nisbahbonus"`
	SaldoEfektif                sql.NullFloat64    `db:"SALDO_EFEKTIF" json:"saldo_efektif"`
	IsTutupKartuAtm             sql.NullString     `db:"IS_TUTUP_KARTU_ATM" json:"is_tutup_kartu_atm"`
	NomorregisterbukuAktif      sql.NullString     `db:"NOMORREGISTERBUKU_AKTIF" json:"nomorregisterbuku_aktif"`
	AlamatJalan                 sql.NullString     `db:"ALAMAT_JALAN" json:"alamat_jalan"`
	AlamatKodePos               sql.NullString     `db:"ALAMAT_KODE_POS" json:"alamat_kode_pos"`
	KodeJenis                   sql.NullString     `db:"KODE_JENIS" json:"kode_jenis"`
	StatusRestriksi             sql.NullString     `db:"STATUS_RESTRIKSI" json:"status_restriksi"`
	TanggalTutup                sql.NullString     `db:"TANGGAL_TUTUP" json:"tanggal_tutup"`
	HakAksesRekening            []HakAksesRekening `json:"hak_akses_rekening"`
}

type HakAksesRekening struct {
	IdIndividu               sql.NullInt64  `db:"ID_INDIVIDU" json:"id_individu"`
	Nama                     sql.NullString `db:"NAMA_LENGKAP" json:"nama"`
	TempatLahir              sql.NullString `db:"TEMPAT_LAHIR" json:"tempat_lahir"`
	TanggalLahir             sql.NullString `db:"TANGGAL_LAHIR" json:"tanggal_lahir"`
	StatusPerkawinan         sql.NullString `db:"STATUS_PERKAWINAN" json:"status_perkawinan"`
	Kewarganegaraan          sql.NullString `db:"KEWARGANEGARAAN" json:"kewarganegaraan"`
	JenisIdentitas           sql.NullString `db:"JENIS_IDENTITAS" json:"jenis_identitas"`
	NomorIdentitas           sql.NullString `db:"NOMOR_IDENTITAS" json:"nomor_identitas"`
	AlamatRumahJalan         sql.NullString `db:"ALAMAT_RUMAH_JALAN" json:"alamat_rumah_jalan"`
	AlamatRumahRt            sql.NullString `db:"ALAMAT_RUMAH_RT" json:"alamat_rumah_rt"`
	AlamatRumahRw            sql.NullString `db:"ALAMAT_RUMAH_RW" json:"alamat_rumah_rw"`
	AlamatRumahKelurahan     sql.NullString `db:"ALAMAT_RUMAH_KELURAHAN" json:"alamat_rumah_kelurahan"`
	AlamatRumahKecamatan     sql.NullString `db:"ALAMAT_RUMAH_KECAMATAN" json:"alamat_rumah_kecamatan"`
	AlamatRumahKotaKabupaten sql.NullString `db:"ALAMAT_RUMAH_KOTA_KABUPATEN" json:"alamat_rumah_kota_kabupaten"`
	AlamatRumahProvinsi      sql.NullString `db:"ALAMAT_RUMAH_PROVINSI" json:"alamat_rumah_provinsi"`
	AlamatRumahKodePos       sql.NullString `db:"ALAMAT_RUMAH_KODE_POS" json:"alamat_rumah_kode_pos"`
	StatusPep                sql.NullString `db:"STATUS_PEP" json:"status_pep"`
	KodePekerjaan            sql.NullString `db:"KODE_PEKERJAAN" json:"kode_pekerjaan"`
	Keterangan               sql.NullString `db:"KETERANGAN" json:"keterangan"`
	KodeSumberDana           sql.NullString `db:"KODE_SUMBER_DANA" json:"kode_sumber_dana"`
	NomorHandphone           sql.NullString `db:"TELEPON_HP_NOMOR" json:"nomor_handphone"`
	JenisIdentitasLain       sql.NullString `db:"JENIS_IDENTITAS_LAIN" json:"jenis_identitas_lain"`
	NomorIdentitasLain       sql.NullString `db:"NOMOR_IDENTITAS_LAIN" json:"nomor_identitas_lain"`
	IsAktif                  sql.NullString `db:"IS_AKTIF" json:"is_aktif"`
	JenisKelamin             sql.NullString `db:"JENIS_KELAMIN"`
	JenisKelaminDesc         sql.NullString `db:"JENIS_KELAMIN_DESC"`
	KodeNegaraAsal           sql.NullString `db:"KODE_NEGARA_ASAL" json:"kode_negara_asal"`
	KodeNegaraAsalDesc       sql.NullString `db:"KODE_NEGARA_ASAL_DESC"`
}

type RekeningTransaksi struct {
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening           sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening,omitempty"`
	KodeValuta             sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta,omitempty"`
	NamaValuta             sql.NullString  `db:"NAMA_VALUTA" json:"nama_valuta,omitempty"`
	Keterangan             sql.NullString  `db:"KETERANGAN" json:"keterangan,omitempty"`
	JenisRekeningTransaksi sql.NullString  `db:"JENIS_REKENING_TRANSAKSI" json:"jenis_rekening_transaksi,omitempty"`
	StatusRekening         sql.NullInt64   `db:"STATUS_REKENING" json:"status_rekening"`
	StatusRekeningDesc     sql.NullString  `db:"STATUS_REKENING_DESC" json:"status_rekening_desc"`
	Saldo                  sql.NullFloat64 `db:"SALDO" json:"saldo"`
}

type ReverseTransactionView struct {
	IdTransaksi    sql.NullInt64  `db:"ID_TRANSAKSI" json:"id_transaksi"`
	TanggalReverse sql.NullTime   `db:"TANGGAL_REVERSE" json:"tanggal_reverse"`
	UserReverse    sql.NullString `db:"USER_REVERSE" json:"user_reverse"`
	ReverseId      sql.NullInt64  `db:"REVERSE_ID" json:"reverse_id"`
}

type DsRekeningTeller struct {
	IdUser             sql.NullString  `db:"ID_USER" json:"id_user"`
	NomorRekening      sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening       sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeCabang         sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	Saldo              sql.NullFloat64 `db:"SALDO" json:"saldo"`
	KodeValuta         sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	IsOpen             sql.NullString  `db:"IS_OPEN" json:"is_open"`
	IsOpenDesc         sql.NullString  `db:"IS_OPEN_DESC" json:"is_open_desc"`
	StatusRekening     sql.NullInt64   `db:"STATUS_REKENING" json:"status_rekening"`
	StatusRekeningDesc sql.NullString  `db:"STATUS_REKENING_DESC" json:"status_rekening_desc"`
	TanggalUpdate      sql.NullString  `db:"TANGGAL_UPDATE" json:"tanggal_update"`
}

type GetListTrxRes struct {
	IdTransaksi         *int32   `db:"ID_TRANSAKSI" json:"id_transaksi,omitempty"`
	TanggalTransaksi    string   `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi,omitempty"`
	NomorReferensi      string   `db:"NOMOR_REFERENSI" json:"nomor_referensi,omitempty"`
	NilaiTransaksi      *float64 `db:"NILAI_TRANSAKSI" json:"nilai_transaksi,omitempty"`
	Keterangan          string   `db:"KETERANGAN" json:"keterangan,omitempty"`
	KodeCabangTransaksi string   `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi,omitempty"`
}

type Transaksi struct {
	IdTransaksi         sql.NullInt64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	NomorReferensi      sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	Keterangan          sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	JenisAplikasi       sql.NullString  `db:"JENIS_APLIKASI" json:"jenis_aplikasi"`
	KodeValutaTransaksi sql.NullString  `db:"KODE_VALUTA_TRANSAKSI" json:"kode_valuta_transaksi"`
	KursManual          sql.NullFloat64 `db:"KURS_MANUAL" json:"kurs_manual"`
	NilaiTransaksi      sql.NullFloat64 `db:"NILAI_TRANSAKSI" json:"nilai_transaksi"`
	KodeCabangTransaksi sql.NullString  `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi"`
	IsSudahDijurnal     sql.NullString  `db:"IS_SUDAH_DIJURNAL" json:"is_sudah_dijurnal"`
	KodeTransaksi       sql.NullString  `db:"KODE_TRANSAKSI" json:"kode_transaksi"`
	StatusOtorisasi     sql.NullInt32   `db:"STATUS_OTORISASI" json:"status_otorisasi"`
	NomorSeri           sql.NullString  `db:"NOMOR_SERI" json:"nomor_seri"`
	JournalNo           sql.NullString  `db:"JOURNAL_NO" json:"journal_no"`
	NomorCounter        sql.NullInt32   `db:"NOMOR_COUNTER" json:"nomor_counter"`
	UserInput           sql.NullString  `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi       sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	TanggalTransaksi    time.Time       `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	TanggalOtorisasi    time.Time       `db:"TANGGAL_OTORISASI" json:"tanggal_otorisasi"`
	JamInput            time.Time       `db:"JAM_INPUT" json:"jam_input"`
	JamOtorisasi        time.Time       `db:"JAM_OTORISASI" json:"jam_otorisasi"`
	Biaya               sql.NullFloat64 `db:"BIAYA" json:"biaya"`
	TanggalInput        time.Time       `db:"TANGGAL_INPUT" json:"tanggal_input"`
	TerminalInput       sql.NullString  `db:"TERMINAL_INPUT" json:"terminal_input"`
	TerminalOtorisasi   sql.NullString  `db:"TERMINAL_OTORISASI" json:"terminal_otorisasi"`
	Channelname         sql.NullString  `db:"CHANNELNAME" json:"channelname"`
	IsReversed          sql.NullString  `db:"IS_REVERSED" json:"is_reversed"`
	IsReverseAllowed    sql.NullString  `db:"IS_REVERSE_ALLOWED" json:"is_reverse_allowed"`
}

type ListHistTrxByNorek struct {
	IdDetilTransaksi     sql.NullInt64   `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	IdTransaksi          sql.NullInt64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	TanggalTransaksi     sql.NullString  `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	IdParameterTransaksi sql.NullString  `db:"ID_PARAMETER_TRANSAKSI" json:"id_parameter_transaksi"`
	KodeJurnal           sql.NullString  `db:"KODE_JURNAL" json:"kode_jurnal"`
	KodeTransaksi        sql.NullString  `db:"KODE_TRANSAKSI" json:"kode_transaksi"`
	Keterangan           sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	JenisMutasi          sql.NullString  `db:"JENIS_MUTASI" json:"jenis_mutasi"`
	SaldoAwal            sql.NullFloat64 `db:"SALDO_AWAL" json:"saldo_awal"`
	Saldo                sql.NullFloat64 `db:"SALDO" json:"saldo"`
	NilaiMutasi          sql.NullFloat64 `db:"NILAI_MUTASI" json:"nilai_mutasi"`
	KodeValuta           sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	NilaiKursManual      sql.NullInt32   `db:"NILAI_KURS_MANUAL" json:"nilai_kurs_manual"`
	NilaiEkuivalen       sql.NullFloat64 `db:"NILAI_EKUIVALEN" json:"nilai_ekuivalen"`
	NomorReferensi       sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	KodeCabangTransaksi  sql.NullString  `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi"`
	UserInput            sql.NullString  `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi        sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	TanggalInput         sql.NullString  `db:"TANGGAL_INPUT" json:"tanggal_input"`
	JamInput             sql.NullString  `db:"JAM_INPUT" json:"jam_input"`
	KeteranganTambahan   sql.NullString  `db:"KETERANGAN_TAMBAHAN" json:"keterangan_tambahan"`
	JournalNo            sql.NullString  `db:"JOURNAL_NO" json:"journal_no"`
	StatusOtorisasi      sql.NullInt32   `db:"STATUS_OTORISASI" json:"status_otorisasi"`
	StatusOtorisasiDesc  sql.NullString  `db:"STATUS_OTORISASI_DESC" json:"status_otorisasi_desc"`
}

type ListTrxByNorek struct {
	IdDetilTransaksi     sql.NullInt64   `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	IdTransaksi          sql.NullInt64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	TanggalTransaksi     sql.NullString  `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	IdParameterTransaksi sql.NullString  `db:"ID_PARAMETER_TRANSAKSI" json:"id_parameter_transaksi"`
	KodeJurnal           sql.NullString  `db:"KODE_JURNAL" json:"kode_jurnal"`
	KodeTransaksi        sql.NullString  `db:"KODE_TRANSAKSI" json:"kode_transaksi"`
	Keterangan           sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	JenisMutasi          sql.NullString  `db:"JENIS_MUTASI" json:"jenis_mutasi"`
	SaldoAwal            sql.NullFloat64 `db:"SALDO_AWAL" json:"saldo_awal"`
	Saldo                sql.NullFloat64 `db:"SALDO" json:"saldo"`
	NilaiMutasi          sql.NullFloat64 `db:"NILAI_MUTASI" json:"nilai_mutasi"`
	KodeValuta           sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	NilaiKursManual      sql.NullInt32   `db:"NILAI_KURS_MANUAL" json:"nilai_kurs_manual"`
	NilaiEkuivalen       sql.NullFloat64 `db:"NILAI_EKUIVALEN" json:"nilai_ekuivalen"`
	NomorReferensi       sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	KodeCabangTransaksi  sql.NullString  `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi"`
	UserInput            sql.NullString  `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi        sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	TanggalInput         sql.NullString  `db:"TANGGAL_INPUT" json:"tanggal_input"`
	JamInput             sql.NullString  `db:"JAM_INPUT" json:"jam_input"`
	KeteranganTambahan   sql.NullString  `db:"KETERANGAN_TAMBAHAN" json:"keterangan_tambahan"`
	JournalNo            sql.NullString  `db:"JOURNAL_NO" json:"journal_no"`
	StatusOtorisasi      sql.NullInt32   `db:"STATUS_OTORISASI" json:"status_otorisasi"`
	StatusOtorisasiDesc  sql.NullString  `db:"STATUS_OTORISASI_DESC" json:"status_otorisasi_desc"`
}

type InfoTransaksiDetail struct {
	IdTransaksi             sql.NullInt64   `db:"ID_TRANSAKSI"`
	NomorRekeningDebet      sql.NullString  `db:"NOMOR_REKENING_DEBET"`
	NomorRekeningKredit     sql.NullString  `db:"NOMOR_REKENING_KREDIT"`
	NamaRekeningDebet       sql.NullString  `db:"NAMA_REKENING_DEBET"`
	NamaRekeningKredit      sql.NullString  `db:"NAMA_REKENING_KREDIT"`
	IsTunai                 sql.NullString  `db:"IS_TUNAI"`
	KodeBankKliring         sql.NullString  `db:"KODE_BANK_KLIRING"`
	KodeMemberRtgs          sql.NullString  `db:"KODE_MEMBER_RTGS"`
	NominalTransaksi        sql.NullFloat64 `db:"NOMINAL_TRANSAKSI"`
	NominalBiaya            sql.NullFloat64 `db:"NOMINAL_BIAYA"`
	AccountBiaya            sql.NullString  `db:"ACCOUNT_BIAYA"`
	NomorRekeningTransaksi  sql.NullString  `db:"NOMOR_REKENING_TRANSAKSI"`
	NamaRekeningTransaksi   sql.NullString  `db:"NAMA_REKENING_TRANSAKSI"`
	NomorRekeningPair       sql.NullString  `db:"NOMOR_REKENING_PAIR"`
	NamaRekeningPair        sql.NullString  `db:"NAMA_REKENING_PAIR"`
	SumberBiaya             sql.NullString  `db:"SUMBER_BIAYA"`
	JenisMutasiTransaksi    sql.NullString  `db:"JENIS_MUTASI_TRANSAKSI"`
	JenisMutasiPair         sql.NullString  `db:"JENIS_MUTASI_PAIR"`
	CounterCetakValidasi    sql.NullString  `db:"COUNTER_CETAK_VALIDASI"`
	IsConfidential          sql.NullString  `db:"IS_CONFIDENTIAL"`
	NomorNasabahWic         sql.NullString  `db:"NOMOR_NASABAH_WIC"`
	IdNotadebitoutwardgroup sql.NullInt64   `db:"ID_NOTADEBITOUTWARDGROUP"`
	NominalTunaiFisik       sql.NullFloat64 `db:"NOMINAL_TUNAI_FISIK"`
	IsTunaiFisik            sql.NullString  `db:"IS_TUNAI_FISIK"`
	IdSumberDanaTrx         sql.NullInt32   `db:"ID_SUMBER_DANA_TRX"`
	IdTujuanTrx             sql.NullInt32   `db:"ID_TUJUAN_TRX"`
	KdSumberDanaTrx         sql.NullString  `db:"KD_SUMBER_DANA_TRX"`
	KdTujuanTrx             sql.NullString  `db:"KD_TUJUAN_TRX"`
	KetSumberDanaTrx        sql.NullString  `db:"KET_SUMBER_DANA_TRX"`
	KetTujuanTrx            sql.NullString  `db:"KET_TUJUAN_TRX"`
	TipeAksesTransaksi      sql.NullString  `db:"TIPE_AKSES_TRANSAKSI"`
	ListnamaAksesTransaksi  sql.NullString  `db:"LISTNAMA_AKSES_TRANSAKSI"`
}

type CheckBalanceMultiJournalTableTemp struct {
	JenisMutasi  string  `db:"JENIS_MUTASI"`
	TotalNominal float64 `db:"TOTAL_NOMINAL"`
}

type GetRekeningLayanan struct {
	NomorRekeningLayanan string `db:"NOMOR_REKENING_LAYANAN"`
	NomorRekeningTujuan  string `db:"NOMOR_REKENING_TUJUAN"`
	SaldoMinimal         string `db:"SALDO_MINIMAL"`
}

type GetDetailTrxUidRes struct {
	Uidkey              string  `db:"UIDKEY" json:"uidkey"`
	IdTransaksi         string  `db:"ID_TRANSAKSI" json:"id_transaksi"`
	KodeJurnal          string  `db:"KODE_JURNAL" json:"kode_jurnal"`
	KodeCabangTransaksi string  `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi"`
	NamaRekening        string  `db:"NAMA_REKENING" json:"nama_rekening"`
	NomorRekening       string  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	Keterangan          string  `db:"KETERANGAN" json:"keterangan"`
	KodeValuta          string  `db:"KODE_VALUTA" json:"kode_valuta"`
	JenisMutasi         string  `db:"JENIS_MUTASI" json:"jenis_mutasi"`
	JenisMutasiDesc     string  `db:"JENIS_MUTASI_DESC" json:"jenis_mutasi_desc"`
	NilaiMutasi         float64 `db:"NILAI_MUTASI" json:"nilai_mutasi"`
	NilaiKursManual     float64 `db:"NILAI_KURS_MANUAL" json:"nilai_kurs_manual"`
	NilaiEkuivalen      float64 `db:"NILAI_EKUIVALEN" json:"nilai_ekuivalen"`
	KodeRC              string  `db:"KODE_RC" json:"kode_rc"`
}

type AccountingDayAvabilityRes struct {
	Datevalue     string `db:"DATEVALUE"`
	PeriodeStatus string `db:"PERIODE_STATUS"`
}
