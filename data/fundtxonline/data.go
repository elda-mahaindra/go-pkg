package data

import "database/sql"

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

type DsDetailTrxRes struct {
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

type GetRekeningLayanan struct {
	NomorRekeningLayanan string `db:"NOMOR_REKENING_LAYANAN"`
	NomorRekeningTujuan  string `db:"NOMOR_REKENING_TUJUAN"`
	SaldoMinimal         string `db:"SALDO_MINIMAL"`
}

type Account struct {
	NamaRekening            string  `db:"NAMA_REKENING"`
	Saldo                   float64 `db:"SALDO"`
	StatusRekening          int     `db:"STATUS_REKENING"`
	KodeJenis               string  `db:"KODE_JENIS"`
	SaldoEfektif            float64 `db:"SALDO_EFEKTIF"`
	SaldoOverride           float64 `db:"SALDO_OVERRIDE"`
	KodeProduk              string  `db:"KODE_PRODUK"`
	NamaProduk              string  `db:"NAMA_PRODUK"`
	JenisRekeningLiabilitas string  `db:"JENIS_REKENING_LIABILITAS"`
}

type JenisDc struct {
	JenisDc    string `db:"JENIS_DC"`
	Keterangan string `db:"KETERANGAN"`
	GlBeban    string `db:"GL_BEBAN"`
	GlIncome   string `db:"GL_INCOME"`
	GlKl       string `db:"GL_KL"`
	GlKlFee    string `db:"GL_KL_FEE"`
	GlTl       string `db:"GL_TL"`
	GlTlInc    string `db:"GL_TL_INC"`
	NetworkId  string `db:"NETWORK_ID"`
	FeeBin     *int   `db:"FEE_BIN"`
	FeeCwd     *int   `db:"FEE_CWD"`
	FeeTrf     *int   `db:"FEE_TRF"`
	FeeInc     *int   `db:"FEE_INC"`
	FeeTcs     *int   `db:"FEE_TCS"`
}

type GetFeeDcResult struct {
	Feenom float64 `db:"FEENOM"`
	Feeid  int     `db:"FEEID"`
}

type FeeGlCreditResult struct {
	IdRekeningdc  int    `db:"ID_REKENINGDC"`
	NomorRekening string `db:"NOMOR_REKENING"`
	JenisRekening string `db:"JENIS_REKENING"`
	JenisDc       string `db:"JENIS_DC"`
	IsFee         string `db:"IS_FEE"`
	Mnemonic      string `db:"MNEMONIC"`
	TrxCode       string `db:"TRX_CODE"`
}

type GetPairGlResult struct {
	FeeGlCreditResult
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
