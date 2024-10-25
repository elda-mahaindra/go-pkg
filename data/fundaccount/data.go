package data

import (
	"database/sql"
	"time"
)

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

type GetAccountForcloseResult struct {
	NamaRekening           sql.NullString  `db:"NAMA_REKENING"`
	KodeProduk             sql.NullString  `db:"KODE_PRODUK"`
	NamaProduk             sql.NullString  `db:"NAMA_PRODUK"`
	Cabang                 sql.NullString  `db:"CABANG"`
	KodeCabang             sql.NullString  `db:"KODE_CABANG"`
	NamaCabang             sql.NullString  `db:"NAMA_CABANG"`
	Saldo                  sql.NullFloat64 `db:"SALDO"`
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING"`
	Valuta                 sql.NullString  `db:"VALUTA"`
	SaldoDitahan           sql.NullFloat64 `db:"SALDO_DITAHAN"`
	IsBlokir               sql.NullString  `db:"IS_BLOKIR"`
	SaldoFloat             sql.NullFloat64 `db:"SALDO_FLOAT"`
	IsSedangDitutup        sql.NullString  `db:"IS_SEDANG_DITUTUP"`
	BiayaPenutupanRekening sql.NullFloat64 `db:"BIAYA_PENUTUPAN_REKENING"`
	IsTutupKartuAtm        sql.NullString  `db:"IS_TUTUP_KARTU_ATM"`
}

type IsAfiliasiJatuhTempoDepositoResult struct {
	NomorRekening     sql.NullString `db:"NOMOR_REKENING"`
	RekeningDisposisi sql.NullString `db:"REKENING_DISPOSISI"`
}

type IsAfiliasiBagiHasilDepositoResult struct {
	NomorRekening          sql.NullString `db:"NOMOR_REKENING"`
	NomorRekeningDisposisi sql.NullString `db:"NOMOR_REKENING_DISPOSISI"`
}

type IsAfiliasiFinancingResult struct {
	NomorRekening   sql.NullString `db:"NOMOR_REKENING"`
	NorekPaymentsrc sql.NullString `db:"NOREK_PAYMENTSRC"`
}

type IsAfiliasiRekeningRencanaResult struct {
	NomorRekening      sql.NullString `db:"NOMOR_REKENING"`
	NomorRekeningInduk sql.NullString `db:"NOMOR_REKENING_INDUK"`
}

type IsAfiliasiPencairanRekeningRencanaResult struct {
	NomorRekening          sql.NullString `db:"NOMOR_REKENING"`
	NomorRekeningPencairan sql.NullString `db:"NOMOR_REKENING_PENCAIRAN"`
}

type IsAfiliasiAtsPerTanggalResult struct {
	Debitaccountno  sql.NullString `db:"DEBITACCOUNTNO"`
	Creditaccountno sql.NullString `db:"CREDITACCOUNTNO"`
}

type IsAfiliasiAtsPerSaldoResult struct {
	NomorRekeningTujuan sql.NullString `db:"NOMOR_REKENING_TUJUAN"`
}

type ResAcctAccessList struct {
	IdIndividu         sql.NullInt32  `db:"ID_INDIVIDU,omitempty"`
	NamaLengkap        sql.NullString `db:"NAMA_LENGKAP,omitempty"`
	JenisIdentitas     sql.NullString `db:"JENIS_IDENTITAS,omitempty"`
	NomorIdentitas     sql.NullString `db:"NOMOR_IDENTITAS,omitempty"`
	JenisIdentitasLain sql.NullString `db:"JENIS_IDENTITAS_LAIN,omitempty"`
	NomorIdentitasLain sql.NullString `db:"NOMOR_IDENTITAS_LAIN,omitempty"`
	TeleponHpNomor     sql.NullString `db:"TELEPON_HP_NOMOR,omitempty"`
	AlamatRumahJalan   sql.NullString `db:"ALAMAT_RUMAH_JALAN,omitempty"`
	TempatLahir        sql.NullString `db:"TEMPAT_LAHIR,omitempty"`
	TanggalLahir       sql.NullString `db:"TANGGAL_LAHIR,omitempty"`
	JenisKelamin       sql.NullString `db:"JENIS_KELAMIN,omitempty"`
	Keterangan         sql.NullString `db:"KETERANGAN,omitempty"`
}

type ListAcctNumber struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING"`
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

type HistblokirAktif struct {
	NomorRekening    sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	StatusBlokir     sql.NullString `db:"STATUS_BLOKIR" json:"status_blokir"`
	SistemPemblokir  sql.NullString `db:"SISTEM_PEMBLOKIR" json:"sistem_pemblokir"`
	UserInput        sql.NullString `db:"USER_INPUT" json:"user_input"`
	TanggalBlokir    sql.NullString `db:"TANGGAL_BLOKIR" json:"tanggal_blokir"`
	KeteranganBlokir sql.NullString `db:"KETERANGAN_BLOKIR" json:"keterangan_blokir"`
}

type DsSaldoRataRata struct {
	NomorRekening sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	Balance       sql.NullFloat64 `db:"BALANCE" json:"balance"`
}

type AgentLiabilitas struct {
	NomorRekening               sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	KodeMarketingCurrent        sql.NullString `db:"KODE_MARKETING_CURRENT" json:"kode_marketing_current"`
	KodeMarketingCurrent_desc   sql.NullString `db:"KODE_MARKETING_CURRENT_DESC" json:"kode_marketing_current_desc"`
	KodeMarketingPertama        sql.NullString `db:"KODE_MARKETING_PERTAMA" json:"kode_marketing_pertama"`
	KodeMarketingPertama_desc   sql.NullString `db:"KODE_MARKETING_PERTAMA_DESC" json:"kode_marketing_pertama_desc"`
	KodeMarketingReferensi      sql.NullString `db:"KODE_MARKETING_REFERENSI" json:"kode_marketing_referensi"`
	KodeMarketingReferensi_desc sql.NullString `db:"KODE_MARKETING_REFERENSI_DESC" json:"kode_marketing_referensi_desc"`
}

type FundingAgent struct {
	Agentcode    sql.NullString `db:"AGENTCODE" json:"agentcode,omitempty"`
	Agentname    sql.NullString `db:"AGENTNAME" json:"agentname,omitempty"`
	Statusactive sql.NullString `db:"STATUSACTIVE" json:"statusactive,omitempty"`
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

type DsPortfolio struct {
	NomorNasabah   sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NomorRekening  sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening   sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeValuta     sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabang     sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	NamaCabang     sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
	Saldo          sql.NullFloat64 `db:"SALDO" json:"saldo"`
	StatusRekening sql.NullString  `db:"STATUS_REKENING" json:"status_rekening"`
	JenisRekening  sql.NullString  `db:"JENIS_REKENING" json:"jenis_rekening"`
	KodeProduk     sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk     sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
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

type ListNomorRekening struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	KodeCabang    sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
	SequenceNo    sql.NullInt64  `db:"SEQUENCE_NO" json:"sequence_no"`
	Status        sql.NullString `db:"STATUS" json:"status"`
}

type Nasabah struct {
	NomorNasabah sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaNasabah  sql.NullString `db:"NAMA_NASABAH" json:"nama_nasabah"`
	JenisNasabah sql.NullString `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	StatusData   sql.NullString `db:"STATUS_DATA" json:"status_data"`
}

type NasabahInfo struct {
	NomorNasabah      sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaNasabah       sql.NullString `db:"NAMA_NASABAH" json:"nama_nasabah"`
	TanggalLahir      sql.NullString `db:"TANGGAL_LAHIR" json:"tanggal_lahir"`
	NamaIbuKandung    sql.NullString `db:"NAMA_IBU_KANDUNG" json:"nama_ibu_kandung"`
	NomorSiup         sql.NullString `db:"NOMOR_SIUP" json:"nomor_siup"`
	Alamat            sql.NullString `db:"ALAMAT" json:"alamat"`
	TeleponRumahNomor sql.NullString `db:"TELEPON_RUMAH_NOMOR" json:"telepon_rumah_nomor"`
	TeleponHpNomor    sql.NullString `db:"TELEPON_HP_NOMOR" json:"telepon_hp_nomor"`
	TelpKantor1Nomor  sql.NullString `db:"TELP_KANTOR1_NOMOR" json:"telp_kantor1_nomor"`
	TelpKantor2Nomor  sql.NullString `db:"TELP_KANTOR2_NOMOR" json:"telp_kantor2_nomor"`
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

type ProdukRencana struct {
	KodeProduk                sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	JangkaWaktuMinimal        sql.NullInt32   `db:"JANGKA_WAKTU_MINIMAL" json:"jangka_waktu_minimal"`
	JangkaWaktuMaksimal       sql.NullInt32   `db:"JANGKA_WAKTU_MAKSIMAL" json:"jangka_waktu_maksimal"`
	JenisRencana              sql.NullString  `db:"JENIS_RENCANA" json:"jenis_rencana"`
	MinimalSetoranRutin       sql.NullFloat64 `db:"MINIMAL_SETORAN_RUTIN" json:"minimal_setoran_rutin"`
	MinimalTargetNominal      sql.NullFloat64 `db:"MINIMAL_TARGET_NOMINAL" json:"minimal_target_nominal"`
	StatusPenarikanSaldo      sql.NullInt32   `db:"STATUS_PENARIKAN_SALDO" json:"status_penarikan_saldo"`
	BiayaPenutupan            sql.NullFloat64 `db:"BIAYA_PENUTUPAN" json:"biaya_penutupan"`
	BiayatutupBawahtarget     sql.NullFloat64 `db:"BIAYATUTUP_BAWAHTARGET" json:"biayatutup_bawahtarget"`
	BiayatutupAtastarget      sql.NullFloat64 `db:"BIAYATUTUP_ATASTARGET" json:"biayatutup_atastarget"`
	MaksimalGagalDebet        sql.NullInt32   `db:"MAKSIMAL_GAGAL_DEBET" json:"maksimal_gagal_debet"`
	AkumulasiProsesGagaldebet sql.NullString  `db:"AKUMULASI_PROSES_GAGALDEBET" json:"akumulasi_proses_gagaldebet"`
	DebetTanggalUlangbulan    sql.NullString  `db:"DEBET_TANGGAL_ULANGBULAN" json:"debet_tanggal_ulangbulan"`
	SetorDenganSi             sql.NullString  `db:"SETOR_DENGAN_SI" json:"setor_dengan_si"`
	TargetNominal             sql.NullFloat64 `db:"TARGET_NOMINAL" json:"target_nominal"`
	BiayaGagalDebet           sql.NullFloat64 `db:"BIAYA_GAGAL_DEBET" json:"biaya_gagal_debet"`
	BiayaTutupOtomatis        sql.NullFloat64 `db:"BIAYA_TUTUP_OTOMATIS" json:"biaya_tutup_otomatis"`
	GraceperiodGagalDebet     sql.NullFloat64 `db:"GRACEPERIOD_GAGAL_DEBET" json:"graceperiod_gagal_debet"`
	IsAutocreateSi            sql.NullString  `db:"IS_AUTOCREATE_SI" json:"is_autocreate_si"`
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

type ResGetRasList struct {
	KodeCabang    sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	NomorRekening sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	Saldo         sql.NullFloat64 `db:"SALDO" json:"saldo"`
	NamaRekening  sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	NamaCabang    sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
}

type GetCifCorpRedundantQuery struct {
	NomorNasabah           string  `excel:"B" db:"NOMOR_NASABAH"`
	KodeCabangInput        string  `excel:"C" db:"KODE_CABANG_INPUT"`
	NamaPerusahaan         string  `excel:"D" db:"NAMA_PERUSAHAAN"`
	NomorNpwp              string  `excel:"E"  db:"NOMOR_NPWP"`
	NomorSiupp             string  `excel:"F" db:"NOMOR_SIUPP"`
	TanggalSiupp           string  `excel:"G" db:"TANGGAL_SIUPP"`
	SkdpNomor              string  `excel:"H" db:"SKDP_NOMOR"`
	AktePendirianNomor     string  `excel:"I" db:"AKTE_PENDIRIAN_NOMOR"`
	TanggalBuka            string  `excel:"J" db:"TANGGAL_BUKA"`
	IsWic                  string  `excel:"K" db:"IS_WIC"`
	Pembiayaan             float64 `excel:"L" db:"PEMBIAYAAN"`
	Tahapan                float64 `excel:"M" db:"TAHAPAN"`
	Deposito               float64 `excel:"N" db:"DEPOSITO"`
	Trib                   float64 `excel:"O" db:"TRIB"`
	Giro                   float64 `excel:"P" db:"GIRO"`
	Groupid                int64   `db:"GROUPID"`
	NamaTrim               string  `db:"NAMA_TRIM"`
	NpwpTrim               string  `db:"NPWP_TRIM"`
	NomorSiuppTrim         string  `db:"NOMOR_SIUPP_TRIM"`
	SkdpNomorTrim          string  `db:"SKDP_NOMOR_TRIM"`
	AktePendirianNomorTrim string  `db:"AKTE_PENDIRIAN_NOMOR_TRIM"`
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

type RekeningCustomer struct {
	NomorRekening         string    `db:"NOMOR_REKENING" json:"nomor_rekening"`
	JenisRekeningCustomer string    `db:"JENIS_REKENING_CUSTOMER" json:"jenis_rekening_customer"`
	NomorNasabah          string    `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	TglAkad               time.Time `db:"TGL_AKAD" json:"tgl_akad"`
}

type RekeningGenerator struct {
	Counter       int64  `db:"COUNTER" json:"counter"`
	KodeGenerator string `db:"KODE_GENERATOR" json:"kode_generator"`
	Grup          string `db:"GRUP" json:"grup"`
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

type JangkaWaktuRencana struct {
	JwDate        sql.NullString `db:"JW_DATE" json:"jw_date"`
	NextMonthDate sql.NullString `db:"NEXT_MONTH_DATE" json:"next_month_date"`
}

type SIRencana struct {
	NomorRekeningKredit sql.NullString `db:"NOMOR_REKENING_KREDIT" json:"nomor_rekening_kredit"`
	Instructionid       sql.NullInt64  `db:"INSTRUCTIONID" json:"instructionid"`
	IdRegister          sql.NullInt64  `db:"ID_REGISTER" json:"id_register"`
	TglProses           sql.NullString `db:"TGL_PROSES" json:"tgl_proses"`
	TglProsesBerikutnya sql.NullString `db:"TGL_PROSES_BERIKUTNYA" json:"tgl_proses_berikutnya"`
	TglKadaluarsa       sql.NullString `db:"TGL_KADALUARSA" json:"tgl_kadaluarsa"`
	JumlahSudahProses   sql.NullInt32  `db:"JUMLAH_SUDAH_PROSES" json:"jumlah_sudah_proses"`
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

type ResGetListTabGiro struct {
	NomorNasabah                sql.NullString `db:"NOMOR_NASABAH"`
	IsWic                       sql.NullString `db:"IS_WIC"`
	NomorRekening               sql.NullString `db:"NOMOR_REKENING"`
	NamaRekening                sql.NullString `db:"NAMA_REKENING"`
	KodeCabang                  sql.NullString `db:"KODE_CABANG"`
	KodeProduk                  sql.NullString `db:"KODE_PRODUK"`
	NamaProduk                  sql.NullString `db:"NAMA_PRODUK"`
	JenisRekeningLiabilitas     sql.NullString `db:"JENIS_REKENING_LIABILITAS"`
	JenisRekeningLiabilitasDesc sql.NullString `db:"JENIS_REKENING_LIABILITAS_DESC"`
	SaldoEfektif                sql.NullString `db:"SALDO_EFEKTIF"`
	StatusRekening              sql.NullString `db:"STATUS_REKENING"`
	StatusRekeningDesc          sql.NullString `db:"STATUS_REKENING_DESC"`
	JenisBlokir                 sql.NullString `db:"JENIS_BLOKIR"`
	IsBlokirDebet               sql.NullString `db:"IS_BLOKIR_DEBET"`
	IsBlokirKredit              sql.NullString `db:"IS_BLOKIR_KREDIT"`
	IsSedangDitutup             sql.NullString `db:"IS_SEDANG_DITUTUP"`
	IsBlokir                    sql.NullString `db:"IS_BLOKIR"`
	Alamat                      sql.NullString `db:"ALAMAT"`
	JenisNasabah                sql.NullString `db:"JENIS_NASABAH"`
	JenisNasabahDesc            sql.NullString `db:"JENIS_NASABAH_DESC"`
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

type DsAlamatAlternatif struct {
	AlamatJalan         string `db:"ALAMAT_JALAN" json:"alamat_jalan"`
	AlamatRtrw          string `db:"ALAMAT_RTRW" json:"alamat_rtrw"`
	AlamatKelurahan     string `db:"ALAMAT_KELURAHAN" json:"alamat_kelurahan"`
	AlamatKecamatan     string `db:"ALAMAT_KECAMATAN" json:"alamat_kecamatan"`
	AlamatKotaKabupaten string `db:"ALAMAT_KOTA_KABUPATEN" json:"alamat_kota_kabupaten"`
	AlamatProvinsi      string `db:"ALAMAT_PROVINSI" json:"alamat_provinsi"`
	AlamatKodePos       string `db:"ALAMAT_KODE_POS" json:"alamat_kode_pos"`
	Telepon             string `db:"TELEPON" json:"telepon"`
}

type JoinAccountDetail struct {
	NomorNasabah   sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaNasabah    sql.NullString `db:"NAMA_NASABAH" json:"nama_nasabah"`
	Alamat         sql.NullString `db:"ALAMAT" json:"alamat"`
	NamaIbuKandung sql.NullString `db:"NAMA_IBU_KANDUNG" json:"nama_ibu_kandung"`
	NomorSiup      sql.NullString `db:"NOMOR_SIUP" json:"nomor_siup"`
	Telepon1       sql.NullString `db:"TELEPON1" json:"telepon1"`
	Telepon2       sql.NullString `db:"TELEPON2" json:"telepon2"`
}

type GetRekeningLayanan struct {
	NomorRekeningLayanan string `db:"NOMOR_REKENING_LAYANAN"`
	NomorRekeningTujuan  string `db:"NOMOR_REKENING_TUJUAN"`
	SaldoMinimal         string `db:"SALDO_MINIMAL"`
}
