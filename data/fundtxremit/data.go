package data

import "database/sql"

type GetPendingClrdepositByGroupResult struct {
	IdGroup                int32           `db:"ID_GROUP"`
	TanggalInput           sql.NullString  `db:"TANGGAL_INPUT"`
	TanggalEfektif         sql.NullString  `db:"TANGGAL_EFEKTIF"`
	StatusNotadebit        sql.NullString  `db:"STATUS_NOTADEBIT"`
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING"`
	NomorAplikasi          sql.NullString  `db:"NOMOR_APLIKASI"`
	Keterangan             sql.NullString  `db:"KETERANGAN"`
	KodeCabang             sql.NullString  `db:"KODE_CABANG"`
	KodeValuta             sql.NullString  `db:"KODE_VALUTA"`
	Inputer                sql.NullString  `db:"INPUTER"`
	IsOtorisasi            sql.NullString  `db:"IS_OTORISASI"`
	TipeGroup              sql.NullString  `db:"TIPE_GRUP"`
	TanggalEksekusiTitipan sql.NullString  `db:"TANGGAL_EKSEKUSI_TITIPAN"`
	UserEksekusiTitipan    sql.NullString  `db:"USER_EKSEKUSI_TITIPAN"`
	JamEksekusiTitipan     sql.NullString  `db:"JAM_EKSEKUSI_TITIPAN"`
	NamaRekening           sql.NullString  `db:"NAMA_REKENING"`
	Nominal                sql.NullFloat64 `db:"NOMINAL"`
}

type ResPendingCLRDEPList struct {
	IdGroup                int32           `db:"ID_GROUP,omitempty"`
	TanggalInput           sql.NullString  `db:"TANGGAL_INPUT,omitempty"`
	TanggalEfektif         sql.NullString  `db:"TANGGAL_EFEKTIF,omitempty"`
	StatusNotadebit        sql.NullString  `db:"STATUS_NOTADEBIT,omitempty"`
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING,omitempty"`
	NomorAplikasi          sql.NullString  `db:"NOMOR_APLIKASI,omitempty"`
	Keterangan             sql.NullString  `db:"KETERANGAN,omitempty"`
	KodeCabang             sql.NullString  `db:"KODE_CABANG,omitempty"`
	KodeValuta             sql.NullString  `db:"KODE_VALUTA,omitempty"`
	Inputer                sql.NullString  `db:"INPUTER,omitempty"`
	IsOtorisasi            sql.NullString  `db:"IS_OTORISASI,omitempty"`
	TipeGroup              sql.NullString  `db:"TIPE_GRUP,omitempty"`
	TanggalEksekusiTitipan sql.NullString  `db:"TANGGAL_EKSEKUSI_TITIPAN,omitempty"`
	UserEksekusiTitipan    sql.NullString  `db:"USER_EKSEKUSI_TITIPAN,omitempty"`
	JamEksekusiTitipan     sql.NullString  `db:"JAM_EKSEKUSI_TITIPAN,omitempty"`
	NamaRekening           sql.NullString  `db:"NAMA_REKENING,omitempty"`
	TotalNominal           sql.NullFloat64 `db:"TOTAL_NOMINAL,omitempty"`
	TotalWarkat            sql.NullInt32   `db:"TOTAL_WARKAT,omitempty"`
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

type GetInterbankDepositListRes struct {
	IdNotaDebitOutward int32           `db:"ID_NOTADEBITOUTWARD"`
	TanggalInput       sql.NullString  `db:"TANGGAL_INPUT"`
	TanggalEfektif     sql.NullString  `db:"TANGGAL_EFEKTIF"`
	NomorRekening      sql.NullString  `db:"NOMOR_REKENING"`
	KodeBank           sql.NullString  `db:"KODE_BANK"`
	NamaBank           sql.NullString  `db:"NAMA_BANK"`
	NomorReferensi     sql.NullString  `db:"NOMOR_REFERENSI"`
	NomorRekeningNota  sql.NullString  `db:"NOMOR_REKENING_NOTA"`
	NamaPemilikNota    sql.NullString  `db:"NAMA_PEMILIK_NOTA"`
	Nominal            sql.NullFloat64 `db:"NOMINAL"`
	KursTransaksi      sql.NullFloat64 `db:"KURS_TRANSAKSI"`
	NomorAplikasi      sql.NullString  `db:"NOMOR_APLIKASI"`
	Inputer            sql.NullString  `db:"INPUTER"`
	JenisWarkat        sql.NullString  `db:"JENIS_WARKAT"`
	KodeCabang         sql.NullString  `db:"KODE_CABANG"`
	NamaCabang         sql.NullString  `db:"NAMA_CABANG"`
	NamaRekening       sql.NullString  `db:"NAMA_REKENING"`
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

type ParameterBiaya struct {
	AccountCode  sql.NullString  `db:"ACCOUNT_CODE" json:"account_code"`
	NominalBiaya sql.NullFloat64 `db:"NOMINAL_BIAYA" json:"nominal_biaya"`
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

type GlInterface struct {
	KodeInterface string `db:"KODE_INTERFACE" json:"kode_interface"`
	KodeAccount   string `db:"KODE_ACCOUNT" json:"kode_account"`
	KodeProduk    string `db:"KODE_PRODUK" json:"kode_produk"`
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

type NotaDebitOutward struct {
	NomorReferensi         sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	TanggalSetoran         sql.NullString  `db:"TANGGAL_SETORAN" json:"tanggal_setoran"`
	KodeValuta             sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	Nominal                sql.NullFloat64 `db:"NOMINAL" json:"nominal"`
	NamaPemilikNota        sql.NullString  `db:"NAMA_PEMILIK_NOTA" json:"nama_pemilik_nota"`
	Status                 sql.NullString  `db:"STATUS" json:"status"`
	JenisWarkat            sql.NullString  `db:"JENIS_WARKAT" json:"jenis_warkat"`
	TanggalInput           sql.NullString  `db:"TANGGAL_INPUT" json:"tanggal_input"`
	IdNotadebitoutward     sql.NullInt64   `db:"ID_NOTADEBITOUTWARD" json:"id_notadebitoutward"`
	IsOtorisasi            sql.NullString  `db:"IS_OTORISASI" json:"is_otorisasi"`
	KursTransaksi          sql.NullFloat64 `db:"KURS_TRANSAKSI" json:"kurs_transaksi"`
	NominalEkuivalen       sql.NullFloat64 `db:"NOMINAL_EKUIVALEN" json:"nominal_ekuivalen"`
	NomorRekeningNota      sql.NullString  `db:"NOMOR_REKENING_NOTA" json:"nomor_rekening_nota"`
	TanggalEfektif         sql.NullString  `db:"TANGGAL_EFEKTIF" json:"tanggal_efektif"`
	Inputer                sql.NullString  `db:"INPUTER" json:"inputer"`
	KodeCabang             sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	Biaya                  sql.NullFloat64 `db:"BIAYA" json:"biaya"`
	BiayaEkuivalen         sql.NullFloat64 `db:"BIAYA_EKUIVALEN" json:"biaya_ekuivalen"`
	SumberBiaya            sql.NullString  `db:"SUMBER_BIAYA" json:"sumber_biaya"`
	NomorAplikasi          sql.NullString  `db:"NOMOR_APLIKASI" json:"nomor_aplikasi"`
	Keterangan             sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	NomorBatch             sql.NullString  `db:"NOMOR_BATCH" json:"nomor_batch"`
	KodeBank               sql.NullString  `db:"KODE_BANK" json:"kode_bank"`
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	IdGroup                sql.NullInt64   `db:"ID_GROUP" json:"id_group"`
	SandiWilayahKliring    sql.NullString  `db:"SANDI_WILAYAH_KLIRING" json:"sandi_wilayah_kliring"`
	StatusWilayahKliring   sql.NullString  `db:"STATUS_WILAYAH_KLIRING" json:"status_wilayah_kliring"`
	KodeAccountTitipan     sql.NullString  `db:"KODE_ACCOUNT_TITIPAN" json:"kode_account_titipan"`
	TanggalEksekusiEfektif sql.NullString  `db:"TANGGAL_EKSEKUSI_EFEKTIF" json:"tanggal_eksekusi_efektif"`
	WaktuEksekusiEfektif   sql.NullString  `db:"WAKTU_EKSEKUSI_EFEKTIF" json:"waktu_eksekusi_efektif"`
	UserEksekusiEfektif    sql.NullString  `db:"USER_EKSEKUSI_EFEKTIF" json:"user_eksekusi_efektif"`
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

type ListDataOWByIDB struct {
	IdNotaDebitOutWard sql.NullInt32   `db:"ID_NOTADEBITOUTWARD"`
	NomorReferensi     sql.NullString  `db:"NOMOR_REFERENSI"`
	NomorAplikasi      sql.NullString  `db:"NOMOR_APLIKASI"`
	NomorRekening      sql.NullString  `db:"NOMOR_REKENING"`
	Keterangan         sql.NullString  `db:"KETERANGAN"`
	Nominal            sql.NullFloat64 `db:"NOMINAL"`
	TanggalEfektif     sql.NullTime    `db:"TANGGAL_EFEKTIF"`
}

type GetClrRejectReason struct {
	IdParameterTolakan string  `json:"id_parameter_tolakan" db:"ID_PARAMETER_TOLAKAN"`
	Keterangan         string  `json:"keterangan" db:"KETERANGAN"`
	BiayaTolakan       float64 `json:"biaya_tolakan" db:"BIAYA_TOLAKAN"`
}

type GetClrWithdrawList struct {
	TanggalTransaksi  string  `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	JenisWarkat       string  `db:"JENIS_WARKAT" json:"jenis_warkat"`
	NomorReferensi    string  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	NomorRekening     string  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening      string  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeValuta        string  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabang        string  `db:"KODE_CABANG" json:"kode_cabang"`
	NilaiMutasi       float64 `db:"NILAI_MUTASI" json:"nilai_mutasi"`
	IdDetailTransaksi int64   `db:"ID_DETIL_TRANSAKSI" json:"id_detil_mutasi"`
	IdTransaksi       int64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	FullName          int64   `db:"FULL_NAME"`
}

type GetListSetoranKliringResponse struct {
	NomorRekening  sql.NullString  `db:"NOMOR_REKENING"`
	TanggalSetoran sql.NullString  `db:"TANGGAL_SETORAN"`
	TanggalEfektif sql.NullString  `db:"TANGGAL_EFEKTIF"`
	Nominal        sql.NullFloat64 `db:"NOMINAL"`
	KodeCabang     sql.NullString  `db:"KODE_CABANG"`
	NomorReferensi sql.NullString  `db:"NOMOR_REFERENSI"`
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

type GetRekeningLayanan struct {
	NomorRekeningLayanan string `db:"NOMOR_REKENING_LAYANAN"`
	NomorRekeningTujuan  string `db:"NOMOR_REKENING_TUJUAN"`
	SaldoMinimal         string `db:"SALDO_MINIMAL"`
}
