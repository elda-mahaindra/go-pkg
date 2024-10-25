package data

import (
	"database/sql"
	"time"
)

type DataLain struct {
	KodeProgram                sql.NullString `db:"KODE_PROGRAM" json:"kode_program"`
	KodeMarketingReferensi     sql.NullString `db:"KODE_MARKETING_REFERENSI" json:"kode_marketing_referensi"`
	KodeMarketingReferensiDesc sql.NullString `db:"KODE_MARKETING_REFERENSI_DESC" json:"kode_marketing_referensi_desc"`
	StatusKelengkapan          sql.NullString `db:"STATUS_KELENGKAPAN" json:"status_kelengkapan"`
}

type AccountingDayByDateResponse struct {
	PeriodStatus sql.NullString `db:"PERIODE_STATUS"`
	IsWorkday    sql.NullString `db:"ISWORKDAY"`
}

type DetailSavPlan struct {
	NomorRekening            sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening             sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	Keterangan               sql.NullString  `db:"KETERANGAN" json:"keterangan" `
	NomorNasabah             sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	JenisNasabah             sql.NullString  `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	TanggalBuka              sql.NullString  `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	TanggalTutup             sql.NullString  `db:"TANGGAL_TUTUP" json:"tanggal_tutup"`
	KodeProduk               sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk               sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
	IsSedangDitutup          sql.NullString  `db:"IS_SEDANG_DITUTUP" json:"is_sedang_ditutup"`
	KodeCabangInput          sql.NullString  `db:"KODE_CABANG" json:"kode_cabang_input"`
	NamaCabang               sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
	IsJoinAccount            sql.NullString  `db:"IS_JOIN_ACCOUNT" json:"is_join_account"`
	KodeValuta               sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	NamaValuta               sql.NullString  `db:"NAMA_VALUTA" json:"nama_valuta"`
	StatusRestriksi          sql.NullString  `db:"STATUS_RESTRIKSI" json:"status_restriksi"`
	KodeStatusRestriksi      sql.NullString  `db:"KODE_STATUS_RETRIKSI" json:"kode_status_retriksi"`
	IsBlokir                 sql.NullString  `db:"IS_BLOKIR" json:"is_blokir"`
	IsBlokirDebet            sql.NullString  `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	IsBlokirKredit           sql.NullString  `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsDapatBagiHasil         sql.NullString  `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	NisbahSpesial            sql.NullFloat64 `db:"NISBAH_SPESIAL" json:"nisbah_spesial"`
	PersentaseZakatBagiHasil sql.NullFloat64 `db:"PERSENTASE_ZAKAT_BAGI_HASIL" json:"persentase_zakat_bagi_hasil"`
	TarifPajak               sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak"`
	NisbahDasar              sql.NullFloat64 `db:"NISBAH_DASAR" json:"nisbah_dasar"`
	IsTieringNisbahbonus     sql.NullString  `db:"IS_TIERING_NISBAHBONUS" json:"is_tiering_nisbahbonus"`
	Saldo                    sql.NullFloat64 `db:"SALDO" json:"saldo"`
	SaldoDitahan             sql.NullFloat64 `db:"SALDO_DITAHAN" json:"saldo_ditahan"`
	SaldoFloat               sql.NullFloat64 `db:"SALDO_FLOAT" json:"saldo_float"`
	SaldoMinimum             sql.NullFloat64 `db:"SALDO_MINIMUM" json:"saldo_minimum"`
	SaldoEfektif             sql.NullFloat64 `db:"SALDO_EFEKTIF" json:"saldo_efektif"`
	IsCetakNota              sql.NullString  `db:"IS_CETAK_NOTA" json:"is_cetak_nota"`
	IsStatusPassbook         sql.NullString  `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook"`
	IsTidakDormant           sql.NullString  `db:"IS_TIDAK_DORMANT" json:"is_tidak_dormant"`
	IsBiayaRekeningDormant   sql.NullString  `db:"IS_BIAYA_REKENING_DORMANT" json:"is_biaya_rekening_dormant"`
	IsKenaBiayalayananumum   sql.NullString  `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum"`
	IsBiayaSaldoMinimum      sql.NullString  `db:"IS_BIAYA_SALDO_MINIMUM" json:"is_biaya_saldo_minimum"`
	IsBiayaAtm               sql.NullString  `db:"IS_BIAYA_ATM" json:"is_biaya_atm"`
	KodeMulticifRelation     sql.NullString  `db:"KODE_MULTICIFRELATION" json:"kode_multicifrelation"`
	UserUpdate               sql.NullString  `db:"USER_UPDATE" json:"user_update"`
	TanggalUpdate            sql.NullString  `db:"TANGGAL_UPDATE" json:"tanggal_update"`
	StatusKelengkapan        sql.NullString  `db:"STATUS_KELENGKAPAN" json:"status_kelengkapan"`
	AlamatKirim              sql.NullString  `db:"ALAMAT_KIRIM" json:"alamat_kirim"`
	TglTransCabangTerakhir   sql.NullString  `db:"TGL_TRANS_CABANG_TERAKHIR" json:"tgl_trans_cabang_terakhir"`
	TglTransEchannelTerakhir sql.NullString  `db:"TGL_TRANS_ECHANNEL_TERAKHIR" json:"tgl_trans_echannel_terakhir"`
	TglTransaksiTerakhir     sql.NullString  `db:"TGL_TRANSAKSI_TERAKHIR" json:"tgl_transaksi_terakhir"`
	KodeProgram              sql.NullString  `db:"KODE_PROGRAM" json:"kode_program"`
	KodeMarketingReferensi   sql.NullString  `db:"KODE_MARKETING_REFERENSI" json:"kode_marketing_referensi"`
	KodeMarketingPertama     sql.NullString  `db:"KODE_MARKETING_PERTAMA" json:"kode_marketing_pertama"`
	KodeMarketingCurrent     sql.NullString  `db:"KODE_MARKETING_CURRENT" json:"kode_marketing_current"`
	NomorRekeningInduk       sql.NullString  `db:"NOMOR_REKENING_INDUK" json:"nomor_rekening_induk"`
	NomorRekeningPencairan   sql.NullString  `db:"NOMOR_REKENING_PENCAIRAN" json:"nomor_rekening_pencairan"`
	TanggalJatuhTempo        sql.NullString  `db:"TANGGAL_JATUH_TEMPO" json:"tanggal_jatuh_tempo"`
	NominalDebet             sql.NullFloat64 `db:"SETORAN_RUTIN" json:"nominal_debet"`
	TanggalSetoranRutin      sql.NullInt32   `db:"TANGGAL_SETORAN_RUTIN" json:"tanggal_setoran_rutin"`
	JangkaWaktu              sql.NullInt32   `db:"JANGKA_WAKTU" json:"jangka_waktu"`
	JumlahKewajibanJt        sql.NullInt32   `db:"JUMLAH_KEWAJIBAN_JT" json:"jumlah_kewajiban_jt"`
	JumlahRealisasi          sql.NullInt32   `db:"JUMLAH_REALISASI" json:"jumlah_realisasi"`
	JumlahTunggakan          sql.NullInt32   `db:"JUMLAH_TUNGGAKAN" json:"jumlah_tunggakan"`
	JenisRekeningLiabilitas  sql.NullString  `db:"JENIS_REKENING_LIABILITAS" json:"jenis_rekening_liabilitas"`
	KodeSumberDana           sql.NullString  `db:"KODE_SUMBER_DANA" json:"kode_sumber_dana"`
	KodeTujuanRekening       sql.NullString  `db:"KODE_TUJUAN_REKENING" json:"kode_tujuan_rekening"`
	NisbahBagiHasil          sql.NullFloat64 `db:"NISBAH_BAGI_HASIL" json:"nisbah_bagi_hasil"`
	StatusRekening           sql.NullInt32   `db:"STATUS_REKENING" json:"status_rekening"`
	StatusRekeningDesc       sql.NullString  `db:"STATUS_REKENING_DESC" json:"status_rekening_desc"`
	TanggalAktifitasTerakhir sql.NullString  `db:"TANGGAL_AKTIFITAS_TERAKHIR" json:"tanggal_aktifitas_terakhir"`
	UserInput                sql.NullString  `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi            sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
}

type ResValidasiRekRT struct {
	NomorRekeningKredit   string `db:"NOMOR_REKENING_KREDIT" json:"nomor_rekening_kredit"`
	StatusRegisterLayanan string `db:"STATUS_REGISTER_LAYANAN" json:"status_register_layanan"`
}

type ListRencanaLayanan struct {
	IdLayanan int64 `db:"ID_LAYANAN" json:"id_layanan"`
}

type ListDataNorekVal struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING"`
	NamaRekening  sql.NullString `db:"NAMA_REKENING"`
	KodeCabang    sql.NullString `db:"KODE_CABANG"`
	NamaCabang    sql.NullString `db:"NAMA_CABANG"`
}

type ResVAL_ACCT_REPOSISI struct {
	IdSeqSession     int64          `db:"ID_SESSION" json:"id_session"`
	NomorRekening    sql.NullString `db:"NOMOR_REKENING"`
	NamaRekening     sql.NullString `db:"NAMA_REKENING"`
	KodeCabang       sql.NullString `db:"KODE_CABANG"`
	NamaCabang       sql.NullString `db:"NAMA_CABANG"`
	KodeCabangTujuan sql.NullString `db:"KODE_CABANG_TUJUAN"`
	NamaCabangTujuan sql.NullString `db:"NAMA_CABANG_TUJUAN"`
	IsValid          string         `db:"IS_VALID"`
	InvalidMessage   string         `db:"INVALID_MESSAGE"`
}

type PostAcctReposisiInsertToInfoReposisi struct {
	NomorRekening    string  `db:"NOMOR_REKENING"`
	KodeCabang       string  `db:"KODE_CABANG"`
	KodeCabangTujuan string  `db:"KODE_CABANG_TUJUAN"`
	SaldoReposisi    float64 `db:"SALDO_REPOSISI"`
	UserInput        string  `db:"USER_INPUT"`
	UserOtorisasi    string  `db:"USER_OTORISASI"`
	KodeCabangInput  string  `db:"KODE_CABANG_INPUT"`
	KodeAccount      string  `db:"KODE_ACCOUNT"`
	KodeProduk       string  `db:"KODE_PRODUK"`
}

type PostAcctChangeProductInsertToInfo struct {
	NomorRekening   string  `db:"NOMOR_REKENING"`
	KodeProduk      string  `db:"KODE_PRODUK"`
	KodeProdukBaru  string  `db:"KODE_PRODUK_BARU"`
	UserInput       string  `db:"USER_INPUT"`
	UserOtorisasi   string  `db:"USER_OTORISASI"`
	KodeCabangInput string  `db:"KODE_CABANG_INPUT"`
	Saldo           float64 `db:"SALDO"`
	KodeCabang      string  `db:"KODE_CABANG"`
	KodeAccount     string  `db:"KODE_ACCOUNT"`
	KodeAccountBaru string  `db:"KODE_ACCOUNT_BARU"`
}

type ResVAL_ACCT_CHANGE_PRODUCT struct {
	IdSeqSession   int64  `db:"ID_SESSION" json:"id_session"`
	NomorRekening  string `db:"NOMOR_REKENING"`
	NamaRekening   string `db:"NAMA_REKENING"`
	KodeProduk     string `db:"KODE_PRODUK"`
	NamaProduk     string `db:"NAMA_PRODUK"`
	JenisProduk    string `db:"JENIS_PRODUK"`
	KodeProdukBaru string `db:"KODE_PRODUK_BARU"`
	NamaProdukBaru string `db:"NAMA_PRODUK_BARU"`
	IsValid        string `db:"IS_VALID"`
	InvalidMessage string `db:"INVALID_MESSAGE"`
}

type AccountingDayAvabilityRes struct {
	Datevalue     string `db:"DATEVALUE"`
	PeriodeStatus string `db:"PERIODE_STATUS"`
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

type BalanceToday struct {
	TotalToday sql.NullFloat64 `db:"TOTAL_TODAY" json:"total_today"`
}

type ResValidasiRegSweep struct {
	NomorRekeningTujuan   string `db:"NOMOR_REKENING_TUJUAN" json:"nomor_rekening_tujuan"`
	StatusRegisterLayanan string `db:"STATUS_REGISTER_LAYANAN" json:"status_register_layanan"`
}

type DsAutotrf struct {
	RekeningDebit             sql.NullString  `db:"REKENING_DEBIT" json:"rekening_debit"`
	NamaRekeningDebit         sql.NullString  `db:"NAMA_REKENING_DEBIT" json:"nama_rekening_debit"`
	RekeningKredit            sql.NullString  `db:"REKENING_KREDIT" json:"rekening_kredit"`
	NamaRekeningKredit        sql.NullString  `db:"NAMA_REKENING_KREDIT" json:"nama_rekening_kredit"`
	TglRegistrasi             sql.NullString  `db:"TGL_REGISTRASI" json:"tgl_registrasi"`
	TglProses                 sql.NullString  `db:"TGL_PROSES" json:"tgl_proses"`
	TglKadaluarsa             sql.NullString  `db:"TGL_KADALUARSA" json:"tgl_kadaluarsa"`
	Keterangan                sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	TglProsesBerikutnya       sql.NullString  `db:"TGL_PROSES_BERIKUTNYA" json:"tgl_proses_berikutnya"`
	KodeTipeInstruksi         sql.NullString  `db:"KODE_TIPE_INSTRUKSI" json:"kode_tipe_instruksi"`
	IdRegister                sql.NullInt64   `db:"ID_REGISTER" json:"id_register"`
	StatusRegisterLayanan     sql.NullString  `db:"STATUS_REGISTER_LAYANAN" json:"status_register_layanan"`
	StatusRegisterLayananDesc sql.NullString  `db:"STATUS_REGISTER_LAYANAN_DESC" json:"status_register_layanan_desc"`
	Nominal                   sql.NullFloat64 `db:"NOMINAL" json:"nominal"`
}

type DsResGetAutosaveList struct {
	NomorRekeningLayanan      sql.NullString  `db:"NOMOR_REKENING_LAYANAN" json:"nomor_rekening_layanan"`
	NamaRekeningLayanan       sql.NullString  `db:"NAMA_REKENING_LAYANAN" json:"nama_rekening_layanan"`
	NomorRekeningTujuan       sql.NullString  `db:"NOMOR_REKENING_TUJUAN" json:"nomor_rekening_tujuan"`
	NamaRekeningTujuan        sql.NullString  `db:"NAMA_REKENING_TUJUAN" json:"nama_rekening_tujuan"`
	TglRegistrasi             sql.NullString  `db:"TGL_REGISTRASI" json:"tgl_registrasi"`
	StatusRegisterLayanan     sql.NullString  `db:"STATUS_REGISTER_LAYANAN" json:"status_register_layanan"`
	StatusRegisterLayananDesc sql.NullString  `db:"STATUS_REGISTER_LAYANAN_DESC" json:"status_register_layanan_desc"`
	JenisRegisterLayanan      sql.NullString  `db:"JENIS_REGISTER_LAYANAN" json:"jenis_register_layanan"`
	IdRegister                sql.NullInt64   `db:"ID_REGISTER" json:"id_register"`
	Description               sql.NullString  `db:"DESCRIPTION" json:"description"`
	SaldoMinimal              sql.NullFloat64 `db:"SALDO_MINIMAL" json:"saldo_minimal"`
	SaldoMaksimal             sql.NullFloat64 `db:"SALDO_MAKSIMAL" json:"saldo_maksimal"`
	BiayaLayanan              sql.NullFloat64 `db:"BIAYA_LAYANAN" json:"biaya_layanan"`
}

type GetBlobDataLengthResult struct {
	BlobDataLength sql.NullInt32 `db:"BLOB_DATA_LENGTH"`
}

type GetBlobDataResult struct {
	BlobData string `db:"BLOB_DATA"`
}

type BlokirDana struct {
	IdHoldDana    *int32   `db:"ID_HOLD_DANA"`
	NominalHold   *float64 `db:"NOMINAL_HOLD"`
	AlasanHold    string   `db:"ALASAN_HOLD"`
	NomorRekening string   `db:"NOMOR_REKENING"`
	Status        string   `db:"STATUS"`
}

type BukuWarkat struct {
	NomorSeriAwal            sql.NullString `db:"NOMOR_SERI_AWAL" json:"nomor_seri_awal"`
	NomorSeriAkhir           sql.NullString `db:"NOMOR_SERI_AKHIR" json:"nomor_seri_akhir"`
	LembarTerpakai           sql.NullString `db:"LEMBAR_TERPAKAI" json:"lembar_terpakai"`
	LembarTersedia           sql.NullString `db:"LEMBAR_TERSEDIA" json:"lembar_tersedia"`
	IdBukuwarkat             sql.NullString `db:"ID_BUKUWARKAT" json:"id_bukuwarkat"`
	JenisWarkat              sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	IsOtorisasi              sql.NullString `db:"IS_OTORISASI" json:"is_otorisasi"`
	TanggalOtorisasi         sql.NullString `db:"TANGGAL_OTORISASI" json:"tanggal_otorisasi"`
	UserOtorisasi            sql.NullString `db:"USER_OTORISASI" json:"user_otorisasi"`
	UserInput                sql.NullString `db:"USER_INPUT" json:"user_input"`
	NomorRekening            sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	StatusBuku               sql.NullString `db:"STATUS_BUKU" json:"status_buku"`
	TanggalInput             sql.NullString `db:"TANGGAL_INPUT" json:"tanggal_input"`
	TanggalAktivasi          sql.NullString `db:"TANGGAL_AKTIVASI" json:"tanggal_aktivasi"`
	UserOtorisasiAktivasi    sql.NullString `db:"USER_OTORISASI_AKTIVASI" json:"user_otorisasi_aktivasi"`
	TanggalOtorisasiAktivasi sql.NullString `db:"TANGGAL_OTORISASI_AKTIVASI" json:"tanggal_otorisasi_aktivasi"`
	UserAktivasi             sql.NullString `db:"USER_AKTIVASI" json:"user_aktivasi"`
}

type Card struct {
	Cardid                  string    `db:"CARDID" json:"cardid"`
	Cardno                  string    `db:"CARDNO" json:"cardno"`
	Sequenceno              string    `db:"SEQUENCENO" json:"sequenceno"`
	Branchcode              string    `db:"BRANCHCODE" json:"branchcode"`
	Customerno              string    `db:"CUSTOMERNO" json:"customerno"`
	Customername            string    `db:"CUSTOMERNAME" json:"customername"`
	Cardstate               int       `db:"CARDSTATE" json:"cardstate"`
	Isprimarycard           string    `db:"ISPRIMARYCARD" json:"isprimarycard"`
	Savingaccountno         string    `db:"SAVINGACCOUNTNO" json:"savingaccountno"`
	Currentaccountno        string    `db:"CURRENTACCOUNTNO" json:"currentaccountno"`
	Registrationdate        time.Time `db:"REGISTRATIONDATE" json:"registrationdate"`
	Printeddate             time.Time `db:"PRINTEDDATE" json:"printeddate"`
	Expirationdate          time.Time `db:"EXPIRATIONDATE" json:"expirationdate"`
	Pinstate                int       `db:"PINSTATE" json:"pinstate"`
	Lastpinmodifieddate     time.Time `db:"LASTPINMODIFIEDDATE" json:"lastpinmodifieddate"`
	Lastpinmodifiedtime     time.Time `db:"LASTPINMODIFIEDTIME" json:"lastpinmodifiedtime"`
	Pinprinteddate          time.Time `db:"PINPRINTEDDATE" json:"pinprinteddate"`
	Delivered               string    `db:"DELIVERED" json:"delivered"`
	Islogon                 string    `db:"ISLOGON" json:"islogon"`
	Cardkind                string    `db:"CARDKIND" json:"cardkind"`
	Isauthorized            string    `db:"ISAUTHORIZED" json:"isauthorized"`
	Wrongpincounter         int       `db:"WRONGPINCOUNTER" json:"wrongpincounter"`
	Insufficientfundcounter int       `db:"INSUFFICIENTFUNDCOUNTER" json:"insufficientfundcounter"`
	Generateddate           time.Time `db:"GENERATEDDATE" json:"generateddate"`
	Changestatusdesc        string    `db:"CHANGESTATUSDESC" json:"changestatusdesc"`
	UserChangepin           string    `db:"USER_CHANGEPIN" json:"user_changepin"`
	InputerRegistrasi       string    `db:"INPUTER_REGISTRASI" json:"inputer_registrasi"`
	UserChangetin           string    `db:"USER_CHANGETIN" json:"user_changetin"`
	Datetinprinted          time.Time `db:"DATETINPRINTED" json:"datetinprinted"`
	Checksum                string    `db:"CHECKSUM" json:"checksum"`
	ProductCode             string    `db:"PRODUCT_CODE" json:"product_code"`
	UpdateDate              time.Time `db:"UPDATE_DATE" json:"update_date"`
	UserUpdate              string    `db:"USER_UPDATE" json:"user_update"`
	OtorUpdate              string    `db:"OTOR_UPDATE" json:"otor_update"`
	Inhtype                 string    `db:"INHTYPE" json:"inhtype"`
	Track2discdata          string    `db:"TRACK2DISCDATA" json:"track2discdata"`
	Effectivedate           time.Time `db:"EFFECTIVEDATE" json:"effectivedate"`
	Basecardno              string    `db:"BASECARDNO" json:"basecardno"`
}

type Cardpin struct {
	Cardid    string `db:"CARDID" json:"cardid"`
	Pinoffset string `db:"PINOFFSET" json:"pinoffset"`
	Pinenc    string `db:"PINENC" json:"pinenc"`
}

type DSCard struct {
	Cardid           sql.NullString `db:"CARDID" json:"cardid"`
	Cardno           sql.NullString `db:"CARDNO" json:"cardno"`
	Cardstate        sql.NullInt64  `db:"CARDSTATE" json:"cardstate"`
	Branchcode       sql.NullString `db:"BRANCHCODE" json:"branchcode"`
	Customerno       sql.NullString `db:"CUSTOMERNO" json:"customerno"`
	Customername     sql.NullString `db:"CUSTOMERNAME" json:"customername"`
	Savingaccountno  sql.NullString `db:"SAVINGACCOUNTNO" json:"savingaccountno"`
	Currentaccountno sql.NullString `db:"CURRENTACCOUNTNO" json:"currentaccountno"`
}

type PostChqValidate struct {
	JenisWarkat   sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	NomorSeriAwal sql.NullString `db:"NOMOR_SERI_AWAL"`
	NomorRekening sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	StatusBuku    sql.NullString `db:"STATUS_BUKU" json:"status_buku"`
}

type DsResGetListCheque struct {
	NomorRekening             sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorSeriAwal             sql.NullString `db:"NOMOR_SERI_AWAL" json:"nomor_seri_awal"`
	NomorSeriAkhir            sql.NullString `db:"NOMOR_SERI_AKHIR" json:"nomor_seri_akhir"`
	LembarTerpakai            sql.NullInt64  `db:"LEMBAR_TERPAKAI" json:"lembar_terpakai"`
	LembarTersedia            sql.NullInt64  `db:"LEMBAR_TERSEDIA" json:"lembar_tersedia"`
	IdBukuwarkat              sql.NullInt64  `db:"ID_BUKUWARKAT" json:"id_bukuwarkat"`
	JenisWarkat               sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	IsOtorisasi               sql.NullString `db:"IS_OTORISASI" json:"is_otorisasi"`
	TanggalOtorisasi          sql.NullString `db:"TANGGAL_OTORISASI" json:"tanggal_otorisasi"`
	TanggalInput              sql.NullString `db:"TANGGAL_INPUT" json:"tanggal_input"`
	TanggalAktivasi           sql.NullString `db:"TANGGAL_AKTIVASI" json:"tanggal_aktivasi"`
	TanggalOtorisasi_aktivasi sql.NullString `db:"TANGGAL_OTORISASI_AKTIVASI" json:"tanggal_otorisasi_aktivasi"`
	UserOtorisasi             sql.NullString `db:"USER_OTORISASI" json:"user_otorisasi"`
	UserInput                 sql.NullString `db:"USER_INPUT" json:"user_input"`
	StatusBuku                sql.NullString `db:"STATUS_BUKU" json:"status_buku"`
}

type DsResGetDetailCheque struct {
	NomorRekening          sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorSeri              sql.NullString `db:"NOMOR_SERI" json:"nomor_seri"`
	TanggalPakaiWarkat     sql.NullString `db:"TANGGAL_PAKAI_WARKAT" json:"tanggal_pakai_warkat"`
	TanggalAktivasi        sql.NullString `db:"TANGGAL_AKTIVASI" json:"tanggal_aktivasi"`
	IdBukuwarkat           sql.NullInt64  `db:"ID_BUKUWARKAT" json:"id_bukuwarkat"`
	JenisWarkat            sql.NullString `db:"JENIS_WARKAT" json:"jenis_warkat"`
	JenisWarkatDesc        sql.NullString `db:"JENIS_WARKAT_DESC" json:"jenis_warkat_desc"`
	UserPengubah           sql.NullString `db:"USER_PENGUBAH" json:"user_pengubah"`
	UserInput              sql.NullString `db:"USER_INPUT" json:"user_input"`
	StatusWarkat           sql.NullString `db:"STATUS_WARKAT" json:"status_warkat"`
	StatusWarkatDesc       sql.NullString `db:"STATUS_WARKAT_DESC" json:"status_warkat_desc"`
	UsedCounter            sql.NullInt64  `db:"USED_COUNTER" json:"used_counter"`
	TanggalPerubahanStatus sql.NullString `db:"TANGGAL_PERUBAHAN_STATUS" json:"tanggal_perubahan_status"`
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

type GetLastID struct {
	LastID *int `db:"LAST_ID" json:"last_id"`
}

type GetBranchCode struct {
	BranchCode string `db:"BRANCH_CODE" json:"branch_code"`
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

type ParameterBiaya struct {
	AccountCode  sql.NullString  `db:"ACCOUNT_CODE" json:"account_code"`
	NominalBiaya sql.NullFloat64 `db:"NOMINAL_BIAYA" json:"nominal_biaya"`
}

type CustomIdgen struct {
	LastIdGen int32  `db:"LAST_IDGEN" json:"last_idgen"`
	Idname    string `db:"IDNAME" json:"idname"`
	Idcode    string `db:"IDCODE" json:"idcode"`
	LastId    int32  `db:"LAST_ID" json:"last_id"`
	Locked    string `db:"LOCKED" json:"locked"`
}

type GlInterface struct {
	KodeInterface string `db:"KODE_INTERFACE" json:"kode_interface"`
	KodeAccount   string `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeProduk    string `db:"KODE_PRODUK" json:"kode_produk"`
}

type Glinterface struct {
	KodeInterface sql.NullString `db:"KODE_INTERFACE" json:"kode_interface"`
	KodeAccount   sql.NullString `db:"KODE_ACCOUNT" json:"kode_account"`
	Deskripsi     sql.NullString `db:"DESKRIPSI" json:"deskripsi"`
	IdGlinterface sql.NullInt64  `db:"ID_GLINTERFACE" json:"id_glinterface"`
	KodeProduk    sql.NullString `db:"KODE_PRODUK" json:"kode_produk"`
}

type GlAccount struct {
	AccountinstanceId int    `db:"ACCOUNTINSTANCE_ID" json:"accountinstance_id"`
	NomorRekening     string `db:"NOMOR_REKENING" json:"nomor_rekening"`
}

type GlobalGlInterface struct {
	KodeAccount   string `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeInterface string `db:"KODE_INTERFACE" json:"kode_interface"`
	NamaAccount   string `db:"NAMA_ACCOUNT" json:"nama_account"`
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

type NotaDebitOutwardGroup struct {
	IdGroup                sql.NullInt64  `db:"ID_GROUP" json:"id_group"`
	TanggalInput           sql.NullString `db:"TANGGAL_INPUT" json:"tanggal_input"`
	TanggalEfektif         sql.NullString `db:"TANGGAL_EFEKTIF" json:"tanggal_efektif"`
	StatusNotadebit        sql.NullString `db:"STATUS_NOTADEBIT" json:"status_notadebit"`
	NomorRekening          sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorAplikasi          sql.NullString `db:"NOMOR_APLIKASI" json:"nomor_aplikasi"`
	Keterangan             sql.NullString `db:"KETERANGAN" json:"keterangan"`
	KodeCabang             sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
	KodeValuta             sql.NullString `db:"KODE_VALUTA" json:"kode_valuta"`
	Inputer                sql.NullString `db:"INPUTER" json:"inputer"`
	IsOtorisasi            sql.NullString `db:"IS_OTORISASI" json:"is_otorisasi"`
	TipeGrup               sql.NullString `db:"TIPE_GRUP" json:"tipe_grup"`
	TanggalEksekusiTitipan sql.NullString `db:"TANGGAL_EKSEKUSI_TITIPAN" json:"tanggal_eksekusi_titipan"`
	UserEksekusiTitipan    sql.NullString `db:"USER_EKSEKUSI_TITIPAN" json:"user_eksekusi_titipan"`
	JamEksekusiTitipan     sql.NullString `db:"JAM_EKSEKUSI_TITIPAN" json:"jam_eksekusi_titipan"`
}

type Layanan struct {
	NamaLayanan             sql.NullString  `db:"NAMA_LAYANAN" json:"nama_layanan"`
	Deskripsi               sql.NullString  `db:"DESKRIPSI" json:"deskripsi"`
	JenisLayanan            sql.NullString  `db:"JENIS_LAYANAN" json:"jenis_layanan"`
	IdLayanan               sql.NullInt64   `db:"ID_LAYANAN" json:"id_layanan"`
	PeriodeProsesLayanan    sql.NullString  `db:"PERIODE_PROSES_LAYANAN" json:"periode_proses_layanan"`
	TanggalProsesBerikutnya sql.NullString  `db:"TANGGAL_PROSES_BERIKUTNYA" json:"tanggal_proses_berikutnya"`
	TanggalProsesTerakhir   sql.NullString  `db:"TANGGAL_PROSES_TERAKHIR" json:"tanggal_proses_terakhir"`
	JumlahMaksimalRegister  sql.NullInt32   `db:"JUMLAH_MAKSIMAL_REGISTER" json:"jumlah_maksimal_register"`
	IsAdaProsesPeriodik     sql.NullString  `db:"IS_ADA_PROSES_PERIODIK" json:"is_ada_proses_periodik"`
	TanggalProsesPertama    sql.NullString  `db:"TANGGAL_PROSES_PERTAMA" json:"tanggal_proses_pertama"`
	IsLayananSystem         sql.NullString  `db:"IS_LAYANAN_SYSTEM" json:"is_layanan_system"`
	IsOverrideTanggalProses sql.NullString  `db:"IS_OVERRIDE_TANGGAL_PROSES" json:"is_override_tanggal_proses"`
	NomorRekeningBiaya      sql.NullString  `db:"NOMOR_REKENING_BIAYA" json:"nomor_rekening_biaya"`
	JarakPeriodeProses      sql.NullInt32   `db:"JARAK_PERIODE_PROSES" json:"jarak_periode_proses"`
	StatusBiayaProses       sql.NullInt32   `db:"STATUS_BIAYA_PROSES" json:"status_biaya_proses"`
	StatusData              sql.NullString  `db:"STATUS_DATA" json:"status_data"`
	TagLayanan              sql.NullString  `db:"TAG_LAYANAN" json:"tag_layanan"`
	BiayaLayanan            sql.NullFloat64 `db:"BIAYA_LAYANAN" json:"biaya_layanan"`
	IdEvent                 sql.NullInt32   `db:"ID_EVENT" json:"id_event"`
	UserPembuat             sql.NullString  `db:"USER_PEMBUAT" json:"user_pembuat"`
	UserPengubah            sql.NullString  `db:"USER_PENGUBAH" json:"user_pengubah"`
}

type Nasabah struct {
	NomorNasabah sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaNasabah  sql.NullString `db:"NAMA_NASABAH" json:"nama_nasabah"`
	JenisNasabah sql.NullString `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	StatusData   sql.NullString `db:"STATUS_DATA" json:"status_data"`
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

type ResRegPassbookLast struct {
	IdCetakPassbook   sql.NullInt64   `db:"ID_CETAK_PASSBOOK"`
	Baris             sql.NullInt64   `db:"BARIS"`
	NilaiMutasi       sql.NullFloat64 `db:"NILAI_MUTASI"`
	JenisMutasi       sql.NullString  `db:"JENIS_MUTASI"`
	Halaman           sql.NullInt64   `db:"HALAMAN"`
	TanggalTransaksi  sql.NullString  `db:"TANGGAL_TRANSAKSI"`
	TanggalCetak      sql.NullString  `db:"TANGGAL_CETAK"`
	IsSudahCetak      sql.NullString  `db:"IS_SUDAH_CETAK"`
	IsGtu             sql.NullString  `db:"IS_GTU"`
	NomorRekening     sql.NullString  `db:"NOMOR_REKENING"`
	KodeTransaksi     sql.NullString  `db:"KODE_TRANSAKSI"`
	Keterangan        sql.NullString  `db:"KETERANGAN"`
	UserCetak         sql.NullString  `db:"USER_CETAK"`
	NoBuku            sql.NullInt64   `db:"NO_BUKU"`
	NomorRegisterBuku sql.NullString  `db:"NOMOR_REGISTER_BUKU"`
}

type DataInfoCetakPassbook struct {
	Halaman sql.NullFloat64 `db:"HALAMAN"`
	Baris   sql.NullFloat64 `db:"BARIS"`
	NoBuku  sql.NullFloat64 `db:"NO_BUKU"`
}

type DsRegisterLayanan struct {
	IdRegister            sql.NullInt64   `db:"ID_REGISTER" json:"id_register"`
	IdLayanan             sql.NullInt64   `db:"ID_LAYANAN" json:"id_layanan"`
	BiayaLayanan          sql.NullFloat64 `db:"BIAYA_LAYANAN" json:"biaya_layanan"`
	NomorRekeningLayanan  sql.NullString  `db:"NOMOR_REKENING_LAYANAN" json:"nomor_rekening_layanan"`
	StatusRegisterLayanan sql.NullString  `db:"STATUS_REGISTER_LAYANAN" json:"status_layanan"`
	TanggalRegistrasi     sql.NullString  `db:"TANGGAL_REGISTRASI" json:"tanggal_registrasi"`
	UserPembuat           sql.NullString  `db:"USER_PEMBUAT" json:"user_pembuat"`
	UserOtorisasi         sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	TanggalOtorisasi      sql.NullString  `db:"TANGGAL_OTORISASI" json:"tanggal_otorisasi"`
	KodeCabangInput       sql.NullString  `db:"KODE_CABANG_INPUT" json:"kode_cabang_input"`
	JenisRegisterLayanan  sql.NullString  `db:"JENIS_REGISTER_LAYANAN" json:"jenis_register_layanan"`
	TanggalNonaktif       sql.NullString  `db:"TANGGAL_NONAKTIF" json:"tanggal_nonaktif"`
	UserInputNonaktif     sql.NullString  `db:"USER_INPUT_NONAKTIF" json:"user_input_nonaktif"`
	UserOtorisasiNonaktif sql.NullString  `db:"USER_OTORISASI_NONAKTIF" json:"user_otorisasi_nonaktif"`
	KodeCabangNonaktif    sql.NullString  `db:"KODE_CABANG_NONAKTIF" json:"kode_cabang_nonaktif"`
}

type Registerrt struct {
	Recid               int64  `db:"RECID" json:"recid"`
	IdRegister          int64  `db:"ID_REGISTER" json:"id_register"`
	Description         string `db:"DESCRIPTION" json:"description"`
	NomorRekeningKredit string `db:"NOMOR_REKENING_KREDIT" json:"nomor_rekening_kredit"`
}

type Produk struct {
	KodeProduk                sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk                sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
	IsOnBalanceSheetMode      sql.NullString  `db:"IS_ONBALANCESHEETMODE" json:"is_on_balance_sheet_mode"`
	SaldoMinimumTidakAktif    sql.NullFloat64 `db:"SALDO_MINIMUM_TIDAK_AKTIF" json:"saldo_minimum_tidak_aktif"`
	JumlahHariJadiTidakAktif  sql.NullInt32   `db:"JUMLAH_HARI_JADI_TIDAK_AKTIF" json:"jumlah_hari_jadi_tidak_aktif"`
	SaldoMinimum              sql.NullFloat64 `db:"SALDO_MINIMUM" json:"saldo_minimum"`
	SaldoMaksimum             sql.NullFloat64 `db:"SALDO_MAKSIMUM" json:"saldo_maksimum"`
	IsBackdated               sql.NullString  `db:"IS_BACKDATED" json:"is_backdated"`
	JumlahBulanMaksBackdated  sql.NullInt32   `db:"JUMLAH_BULAN_MAKS_BACKDATED" json:"jumlah_bulan_maks_backdated"`
	IsKenaZakatBagiHasil      sql.NullString  `db:"IS_KENA_ZAKAT_BAGI_HASIL" json:"is_kena_zakat_bagi_hasil"`
	PresentaseZakatBagiHasil  sql.NullFloat64 `db:"PRESENTASE_ZAKAT_BAGI_HASIL" json:"presentase_zakat_bagi_hasil"`
	IsKenaPajak               sql.NullString  `db:"IS_KENA_PAJAK" json:"is_kena_pajak"`
	TarifPajak                sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak"`
	DisposisiBagiHasil        sql.NullString  `db:"DISPOSISI_BAGI_HASIL" json:"disposisi_bagi_hasil"`
	PeriodeBagiHasil          sql.NullString  `db:"PERIODE_BAGI_HASIL" json:"periode_bagi_hasil"`
	JenisProduk               sql.NullString  `db:"JENIS_PRODUK" json:"jenis_produk"`
	KodeValuta                sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	NamaValuta                sql.NullString  `db:"NAMA_VALUTA" json:"nama_valuta"`
	Status                    sql.NullString  `db:"STATUS" json:"status"`
	IsProdukBerasuransi       sql.NullString  `db:"IS_PRODUK_BERASURANSI" json:"is_produk_berasuransi"`
	IsProdukPartnership       sql.NullString  `db:"IS_PRODUK_PARTNERSHIP" json:"is_produk_partnership"`
	IsOverrideBagiHasilKhusus sql.NullString  `db:"IS_OVERRIDE_BAGIHASILKHUSUS" json:"is_override_bagihasil_khusus"`
	IsOverrideTarifPajak      sql.NullString  `db:"IS_OVERRIDE_TARIFPAJAK" json:"is_override_tarif_pajak"`
	IsOverrideZakatBagiHasil  sql.NullString  `db:"IS_OVERRIDE_ZAKATBAGIHASIL" json:"is_override_zakat_bagihasil"`
	IsOverrideDisposisiBGH    sql.NullString  `db:"IS_OVERRIDE_DISPOSISI_BGH" json:"is_override_disposisi_bgh"`
	JarakPeriodeBagiHasil     sql.NullInt32   `db:"JARAK_PERIODE_BAGIHASIL" json:"jarak_periode_bagihasil"`
	IsOverrideBackdated       sql.NullString  `db:"IS_OVERRIDE_BACKDATED" json:"is_override_backdated"`
	IsKenaBiayalayananumum    sql.NullString  `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biaya_layanan_umum"`
	IsProdukKartu             sql.NullString  `db:"IS_PRODUKKARTU" json:"is_produk_kartu"`
	NamaAsuransi              sql.NullString  `db:"NAMA_ASURANSI" json:"nama_asuransi"`
	IDPartner                 sql.NullString  `db:"ID_PARTNER" json:"id_partner"`
	NamaGrup                  sql.NullString  `db:"NAMA_GRUP" json:"nama_grup"`
	InsplanID                 sql.NullInt64   `db:"INSPLANID" json:"insplan_id"`
	Psplanid                  sql.NullInt64   `db:"PSPLANID" json:"psplanid"`
	BiayaAdminBulanan         sql.NullFloat64 `db:"BIAYA_ADM_BULANAN" json:"biaya_admin_bulanan"`
	IsTieringBiayaadm         sql.NullString  `db:"IS_TIERING_BIAYAADM" json:"is_tiering_biayaadm"`
	IsTieringNisbahbonus      sql.NullString  `db:"IS_TIERING_NISBAHBONUS" json:"is_tiering_nisbahbonus"`
	JenisAkad                 sql.NullString  `db:"JENIS_AKAD" json:"jenis_akad"`
	KodeRencanaAsuransi       sql.NullString  `db:"KODE_RENCANA_ASURANSI" json:"kode_rencana_asuransi"`
	NisbahBonusDasar          sql.NullFloat64 `db:"NISBAH_BONUS_DASAR" json:"nisbah_bonus_dasar"`
	SetoranAwalMinimum        sql.NullFloat64 `db:"SETORAN_AWAL_MINIMUM" json:"setoran_awal_minimum"`
	IsParamSaldoTidakAktif    sql.NullString  `db:"IS_PARAM_SALDO_TIDAK_AKTIF" json:"is_param_saldo_tidak_aktif"`
	IsOverrideBiayaPenutupan  sql.NullString  `db:"IS_OVERRIDE_BIAYA_PENUTUPAN" json:"is_override_biaya_penutupan"`
	BiayaPenutupanRekening    sql.NullFloat64 `db:"BIAYA_PENUTUPAN_REKENING" json:"biaya_penutupan_rekening"`
	SumberBiayaADM            sql.NullInt32   `db:"SUMBER_BIAYA_ADM" json:"sumber_biaya_adm"`
	JumlahHariPertahun        sql.NullInt32   `db:"JUMLAH_HARI_PERTAHUN" json:"jumlah_hari_pertahun"`
	KodeDigit                 sql.NullString  `db:"KODE_DIGIT" json:"kode_digit"`
	IsBlokirDebetInternal     sql.NullString  `db:"IS_BLOKIR_DEBET_INTERNAL" json:"is_blokir_debet_internal"`
	IsBlokirDebetEksternal    sql.NullString  `db:"IS_BLOKIR_DEBET_EKSTERNAL" json:"is_blokir_debet_eksternal"`
	IsBlokirKreditInternal    sql.NullString  `db:"IS_BLOKIR_KREDIT_INTERNAL" json:"is_blokir_kredit_internal"`
	IsBlokirKreditEksternal   sql.NullString  `db:"IS_BLOKIR_KREDIT_EKSTERNAL" json:"is_blokir_kredit_eksternal"`
	IsBiayaSaldoMinimum       sql.NullString  `db:"IS_BIAYA_SALDO_MINIMUM" json:"is_biaya_saldo_minimum"`
	IsBiayaRekeningDormant    sql.NullString  `db:"IS_BIAYA_REKENING_DORMANT" json:"is_biaya_rekening_dormant"`
	IsBiayaAtm                sql.NullString  `db:"IS_BIAYA_ATM" json:"is_biaya_atm"`
	IsCetakNota               sql.NullString  `db:"IS_CETAK_NOTA" json:"is_cetak_nota"`
	IsStatusPassbook          sql.NullString  `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook"`
	Kolektibilitas            sql.NullString  `db:"KOLEKTIBILITAS" json:"kolektibilitas"`
	IsDapatBagiHasil          sql.NullString  `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	IsTidakDormant            sql.NullString  `db:"IS_TIDAK_DORMANT" json:"is_tidak_dormant"`
	TanggalAcuanDormant       sql.NullString  `db:"TANGGAL_ACUAN_DORMANT" json:"tanggal_acuan_dormant"`
	BiayaRekeningDormant      sql.NullFloat64 `db:"BIAYA_REKENING_DORMANT" json:"biaya_rekening_dormant"`
	BiayaSaldoMinimum         sql.NullFloat64 `db:"BIAYA_SALDO_MINIMUM" json:"biaya_saldo_minimum"`
	JumlahHariTutupOtomatis   sql.NullInt64   `db:"JUMLAH_HARI_TUTUP_OTOMATIS" json:"jumlah_hari_tutup_otomatis"`
	IsBlokirDebet             sql.NullString  `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	IsBlokirKredit            sql.NullString  `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsTutupOtomatisDormant    sql.NullString  `db:"IS_TUTUP_OTOMATIS_DORMANT" json:"is_tutup_otomatis_dormant"`
	BiayaADMATM               sql.NullFloat64 `db:"BIAYA_ADM_ATM" json:"biaya_adm_atm"`
	BiayaKartuATM             sql.NullFloat64 `db:"BIAYA_KARTU_ATM" json:"biaya_kartu_atm"`
	EkuivalenRate             sql.NullFloat64 `db:"EKUIVALEN_RATE" json:"ekuivalen_rate"`
	SaldoMinimumBagiHasil     sql.NullFloat64 `db:"SALDO_MINIMUM_BAGIHASIL" json:"saldo_minimum_bagihasil"`
	IsProsesCadangan          sql.NullString  `db:"IS_PROSES_CADANGAN" json:"is_proses_cadangan"`
	PersentaseCadangan        sql.NullFloat64 `db:"PERSENTASE_CADANGAN" json:"persentase_cadangan"`
	OrderMB                   sql.NullString  `db:"ORDER_MB" json:"order_mb"`
	IsBiayaMaterai            sql.NullString  `db:"IS_BIAYA_MATERAI" json:"is_biaya_materai"`
	IsOverrideBiayaatm        sql.NullString  `db:"IS_OVERRIDE_BIAYAATM" json:"is_override_biayaatm"`
	IDRencanaBagiHasilDefault sql.NullInt64   `db:"ID_RENCANABAGIHASILDEFAULT" json:"id_rencana_bagihasil_default"`
	IDBiayaLayananUmum        sql.NullInt64   `db:"ID_BIAYALAYANANUMUM" json:"id_biaya_layanan_umum"`
	IDTieringBiayaADM         sql.NullInt64   `db:"ID_TIERING_BIAYAADM" json:"id_tiering_biaya_adm"`
	IDTieringNisbahBonus      sql.NullInt64   `db:"ID_TIERING_NISBAHBONUS" json:"id_tiering_nisbah_bonus"`
	IDAsuransi                sql.NullFloat64 `db:"ID_ASURANSI" json:"id_asuransi"`
}
type ResGetListSavPlan struct {
	Produk
	KodeCabang sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
}

type ResGetListDep struct {
	KodeProduk               sql.NullString `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk               sql.NullString `db:"NAMA_PRODUK" json:"nama_produk"`
	NisbahBonusDasar         sql.NullString `db:"NISBAH_BONUS_DASAR" json:"nisbah_bonus_dasar"`
	NamaValuta               sql.NullString `db:"NAMA_VALUTA" json:"nama_valuta"`
	IsDapatBagiHasil         sql.NullString `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	TarifPajak               sql.NullString `db:"TARIF_PAJAK" json:"tarif_pajak"`
	PresentaseZakatBagiHasil sql.NullString `db:"PRESENTASE_ZAKAT_BAGI_HASIL" json:"presentase_zakat_bagi_hasil"`
	KodeValuta               sql.NullString `db:"KODE_VALUTA" json:"kode_valuta"`
	IsKenaBiayaLayananUmum   sql.NullString `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum"`
	IsBiayaAtm               sql.NullString `db:"IS_BIAYA_ATM" json:"is_biaya_atm"`
	JenisProduk              sql.NullString `db:"JENIS_PRODUK" json:"jenis_produk"`
	KodeCabang               sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
}

type AccountLiabilitas struct {
	NomorNasabah            string `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NomorRekening           string `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening            string `db:"NAMA_REKENING" json:"nama_rekening"`
	StatusRekening          int32  `db:"STATUS_REKENING" json:"status_rekening"`
	KodeCabang              string `db:"KODE_CABANG" json:"kode_cabang"`
	IsSedangDitutup         string `db:"IS_SEDANG_DITUTUP" json:"is_sedang_ditutup"`
	IsBlokir                string `db:"IS_BLOKIR" json:"is_blokir"`
	IsBlokirKredit          string `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsBlokirDebet           string `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	KodeProduk              string `db:"KODE_PRODUK" json:"kode_produk"`
	JenisRekeningLiabilitas string `db:"JENIS_REKENING_LIABILITAS" json:"jenis_rekening_liabilitas"`
}

type RekeningNasabah struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorNasabah  sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaNasabah   sql.NullString `db:"NAMA_NASABAH" json:"nama_nasabah"`
}

type SaldoDitahan struct {
	SaldoDitahan *float64 `db:"SALDO_DITAHAN"`
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

type GetDataHeaderSvsResult struct {
	NomorRekening    sql.NullString `db:"NOMOR_REKENING"`
	NamaRekening     sql.NullString `db:"NAMA_REKENING"`
	NomorNasabah     sql.NullString `db:"NOMOR_NASABAH"`
	Remark           sql.NullString `db:"REMARK"`
	Additionalremark sql.NullString `db:"ADDITIONALREMARK"`
}

type GetSpecimentAcctListOfImage struct {
	Imageid  int64  `db:"IMAGEID"`
	Imagetag string `db:"IMAGETAG"`
	Remark   string `db:"REMARK"`
	IdData   int64  `db:"ID_DATA"`
	Bucket   string
	FilePath string
	FileUrl  string
}

type CheckSpecimentAcctListResponse struct {
	AccountNo string `db:"ACCOUNTNO"`
}

type GetTagSequenceResponse struct {
	ImageTag    sql.NullString `db:"IMAGETAG"`
	MaxSequence sql.NullInt32  `db:"MAXSEQUENCE"`
}

type RekeningTabgirDetail struct {
	KodeCabangInput             sql.NullString  `db:"KODE_CABANG" json:"kode_cabang_input"`
	NamaCabang                  sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
	UserInput                   sql.NullString  `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi               sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	AlamatStatement             sql.NullString  `db:"ALAMAT_KIRIM" json:"alamat_statement"`
	NomorNasabah                sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	JenisNasabah                sql.NullString  `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	NomorRekening               sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	TanggalBuka                 sql.NullString  `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	TanggalTutup                sql.NullString  `db:"TANGGAL_TUTUP" json:"tanggal_tutup"`
	StatusRekening              sql.NullInt32   `db:"STATUS_REKENING" json:"status_rekening"`
	RelasiMultiCif              sql.NullString  `db:"RELASI_MULTI_CIF" json:"relasi_multi_cif"`
	KodeMulticifrelation        sql.NullString  `db:"KODE_MULTICIFRELATION" json:"kode_multicifrelation"`
	JenisRekeningLiabilitas     sql.NullString  `db:"JENIS_REKENING_LIABILITAS" json:"jenis_rekening_liabilitas"`
	NamaRekening                sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeProduk                  sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk                  sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
	NisbahSpesial               sql.NullFloat64 `db:"NISBAH_SPESIAL" json:"nisbah_spesial"`
	NisbahDasar                 sql.NullFloat64 `db:"NISBAH_DASAR" json:"nisbah_dasar"`
	TarifPajak                  sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak"`
	PersentaseZakatBagiHasil    sql.NullFloat64 `db:"PERSENTASE_ZAKAT_BAGI_HASIL" json:"persentase_zakat_bagi_hasil"`
	KodeSumberDana              sql.NullString  `db:"KODE_SUMBER_DANA" json:"kode_sumber_dana"`
	KodeTujuanRekening          sql.NullString  `db:"KODE_TUJUAN_REKENING" json:"kode_tujuan_rekening"`
	KodeMarketingCurrent        sql.NullString  `db:"KODE_MARKETING_CURRENT" json:"kode_marketing_current"`
	KodeMarketingPertama        sql.NullString  `db:"KODE_MARKETING_PERTAMA" json:"kode_marketing_pertama"`
	JenisRekeningLiabilitasDesc sql.NullString  `db:"JENIS_REKENING_LIABILITAS_DESC" json:"jenis_rekening_liabilitas_desc"`
	IsKenaBiayalayananumum      sql.NullString  `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum"`
	IsStatusPassbook            sql.NullString  `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook"`
	Saldo                       sql.NullFloat64 `db:"SALDO" json:"saldo"`
	SaldoFloat                  sql.NullFloat64 `db:"SALDO_FLOAT" json:"saldo_float"`
	SaldoDeposito               sql.NullFloat64 `db:"SALDO_DEPOSITO" json:"saldo_deposito"`
	SaldoDitahan                sql.NullFloat64 `db:"SALDO_DITAHAN" json:"saldo_ditahan"`
	SaldoMinimum                sql.NullFloat64 `db:"SALDO_MINIMUM" json:"saldo_minimum"`
	SaldoEfektif                sql.NullFloat64 `db:"SALDO_EFEKTIF" json:"saldo_efektif"`
	SaldoAccrualBagihasil       sql.NullFloat64 `db:"SALDO_ACCRUAL_BAGIHASIL" json:"saldo_accrual_bagihasil"`
	SaldoTunggakan              sql.NullFloat64 `db:"SALDO_TUNGGAKAN" json:"saldo_tunggakan"`
	IsJoinAccount               sql.NullString  `db:"IS_JOIN_ACCOUNT" json:"is_join_account"`
	KodeValuta                  sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	NamaValuta                  sql.NullString  `db:"NAMA_VALUTA" json:"nama_valuta"`
	Keterangan                  sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	IsSedangDitutup             sql.NullString  `db:"IS_SEDANG_DITUTUP" json:"is_sedang_ditutup"`
	KodeProgram                 sql.NullString  `db:"KODE_PROGRAM" json:"kode_program"`
	KodeMarketingReferensi      sql.NullString  `db:"KODE_MARKETING_REFERENSI" json:"kode_marketing_referensi"`
	StatusKelengkapan           sql.NullString  `db:"STATUS_KELENGKAPAN" json:"status_kelengkapan"`
	IsBlokir                    sql.NullString  `db:"IS_BLOKIR" json:"is_blokir"`
	IsBlokirDebet               sql.NullString  `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	IsBlokirKredit              sql.NullString  `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsCetakNota                 sql.NullString  `db:"IS_CETAK_NOTA" json:"is_cetak_nota"`
	IsDapatBagiHasil            sql.NullString  `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	IsTidakDormant              sql.NullString  `db:"IS_TIDAK_DORMANT" json:"is_tidak_dormant"`
	IsBiayaRekeningDormant      sql.NullString  `db:"IS_BIAYA_REKENING_DORMANT" json:"is_biaya_rekening_dormant"`
	IsBiayaSaldoMinimum         sql.NullString  `db:"IS_BIAYA_SALDO_MINIMUM" json:"is_biaya_saldo_minimum"`
	IsBiayaAtm                  sql.NullString  `db:"IS_BIAYA_ATM" json:"is_biaya_atm"`
	NisbahBagiHasil             sql.NullFloat64 `db:"NISBAH_BAGI_HASIL" json:"nisbah_bagi_hasil"`
	IsTieringNisbahbonus        sql.NullString  `db:"IS_TIERING_NISBAHBONUS" json:"is_tiering_nisbahbonus"`
	TanggalAktifitasTerakhir    sql.NullString  `db:"TANGGAL_AKTIFITAS_TERAKHIR" json:"tanggal_aktifitas_terakhir"`
	TglTransCabangTerakhir      sql.NullString  `db:"TGL_TRANS_CABANG_TERAKHIR" json:"tgl_trans_cabang_terakhir"`
	TglTransEchannelTerakhir    sql.NullString  `db:"TGL_TRANS_ECHANNEL_TERAKHIR" json:"tgl_trans_echannel_terakhir"`
	TglTransaksiTerakhir        sql.NullString  `db:"TGL_TRANSAKSI_TERAKHIR" json:"tgl_transaksi_terakhir"`
	UserUpdate                  sql.NullString  `db:"USER_UPDATE" json:"user_update"`
	TanggalUpdate               sql.NullString  `db:"TANGGAL_UPDATE" json:"tanggal_update"`
	JenisRegisterLayanan        sql.NullString  `db:"JENIS_REGISTER_LAYANAN" json:"jenis_register_layanan"`
	StatusRestriksi             sql.NullString  `db:"STATUS_RESTRIKSI" json:"status_restriksi"`
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

type GetDetailTrxRes struct {
	IdDetilTransaksi     sql.NullInt64   `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	IdTransaksi          sql.NullInt64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	JenisMutasi          sql.NullString  `db:"JENIS_MUTASI" json:"jenis_mutasi"`
	KodeValuta           sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeValutaRak        sql.NullString  `db:"KODE_VALUTA_RAK" json:"kode_valuta_rak"`
	TanggalTransaksi     sql.NullString  `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	NilaiMutasi          sql.NullFloat64 `db:"NILAI_MUTASI" json:"nilai_mutasi"`
	Keterangan           sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	KodeCabang           sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	JenisDetilTransaksi  sql.NullString  `db:"JENIS_DETIL_TRANSAKSI" json:"jenis_detil_transaksi"`
	NilaiKursManual      sql.NullFloat64 `db:"NILAI_KURS_MANUAL" json:"nilai_kurs_manual"`
	SaldoAwal            sql.NullFloat64 `db:"SALDO_AWAL" json:"saldo_awal"`
	NilaiEkuivalen       sql.NullFloat64 `db:"NILAI_EKUIVALEN" json:"nilai_ekuivalen"`
	NomorReferensi       sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	KodeAccount          sql.NullString  `db:"KODE_ACCOUNT" json:"kode_account"`
	NomorRekening        sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	KodeJurnal           sql.NullString  `db:"KODE_JURNAL" json:"kode_jurnal"`
	IdParameterTransaksi sql.NullString  `db:"ID_PARAMETER_TRANSAKSI" json:"id_parameter_transaksi"`
	FlagProses           sql.NullString  `db:"FLAG_PROSES" json:"flag_proses"`
	Sequence             sql.NullString  `db:"SEQUENCE" json:"sequence"`
	KodeTxClass          sql.NullString  `db:"KODE_TX_CLASS" json:"kode_tx_class"`
	NomorSeriWarkat      sql.NullString  `db:"NOMOR_SERI_WARKAT" json:"nomor_seri_warkat"`
	FlagCode             sql.NullString  `db:"FLAG_CODE" json:"flag_code"`
}

type DetailTrxTDN struct {
	IdDetilTransaksi   sql.NullInt64  `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	IdTransaksi        sql.NullInt64  `db:"ID_TRANSAKSI" json:"id_transaksi"`
	NomorRekening      sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	IdNotadebitoutward sql.NullInt64  `db:"ID_NOTADEBITOUTWARD" json:"id_notadebitoutward"`
}

type GetVaResult struct {
	Accountname    string `db:"ACCOUNTNAME"`
	Admcosttype    string `db:"ADMCOSTTYPE"`
	Admcost        int    `db:"ADMCOST"`
	Paymenttype    string `db:"PAYMENTTYPE"`
	NomorRekening  string `db:"NOMOR_REKENING"`
	KodeCabang     string `db:"KODE_CABANG"`
	Acctype        string
	KodeJenis      string
	StatusRekening string
	NamaRekening   string
	IsBlokir       string
	IsBolehKredit  string
	Isfee          int
	Fee            int
}

type GetVaBillingResult struct {
	Billamount    int64  `db:"BILLAMOUNT"`
	Paymentstatus int    `db:"PAYMENTSTATUS"`
	Billingid     int    `db:"BILLINGID"`
	Status        string `db:"STATUS"`
}

type GetVaPaidResult struct {
	Totalpayment int `db:"TOTALPAYMENT"`
}

type GetVaHistPaidResult struct {
	Totalpayment int `db:"TOTALPAYMENT"`
}
