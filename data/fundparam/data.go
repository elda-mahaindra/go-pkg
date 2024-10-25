package data

import (
	"database/sql"
)

type ResKodeBank struct {
	KodeMember        sql.NullString `db:"KODE_MEMBER" json:"kode_member"`
	NamaMember        sql.NullString `db:"NAMA_MEMBER" json:"nama_member"`
	NamaSingkat       sql.NullString `db:"NAMA_SINGKAT" json:"nama_singkat"`
	StatusPesertaRtgs sql.NullString `db:"STATUS_PESERTA_RTGS" json:"status_peserta_rtgs"`
	StatusPesertaSkn  sql.NullString `db:"STATUS_PESERTA_SKN" json:"status_peserta_skn"`
	AccountNo         sql.NullString `db:"ACCOUNT_NO" json:"account_no"`
	AccountName       sql.NullString `db:"ACCOUNT_NAME" json:"account_name"`
}

type ResKliringKota struct {
	SandiKota     sql.NullString `db:"SANDI_KOTA" json:"sandi_kota"`
	SandiPropinsi sql.NullString `db:"SANDI_PROPINSI" json:"sandi_propinsi"`
	NamaKota      sql.NullString `db:"NAMA_KOTA" json:"nama_kota"`
	Keterangan    sql.NullString `db:"KETERANGAN" json:"keterangan"`
}

type ResTipeMessage struct {
	TipeMessage sql.NullString `db:"TIPE_MESSAGE" json:"tipe_message"`
	NamaMessage sql.NullString `db:"NAMA_MESSAGE" json:"nama_message"`
}

type ResTipeTrx struct {
	KodeTxRtgs sql.NullString `db:"KODE_TX_RTGS" json:"kode_tx_rtgs"`
	NamaTxRtgs sql.NullString `db:"NAMA_TX_RTGS" json:"nama_tx_rtgs"`
}

type ResTipePrior struct {
	KodePrioritas sql.NullString `db:"KODE_PRIORITAS" json:"kode_prioritas"`
	NamaPrioritas sql.NullString `db:"NAMA_PRIORITAS" json:"nama_prioritas"`
}

type ResGetCabangDetail struct {
	KodeCabang          sql.NullString `db:"KODE_CABANG" json:"kode_cabang,omitempty"`
	NamaCabang          sql.NullString `db:"NAMA_CABANG" json:"nama_cabang,omitempty"`
	TipeCabang          sql.NullString `db:"TIPE_CABANG" json:"tipe_cabang,omitempty"`
	StatusAkses         sql.NullString `db:"STATUS_AKSES" json:"status_akses,omitempty"`
	StatusAksesDesc     sql.NullString `db:"STATUS_AKSES_DESC" json:"status_akses_desc,omitempty"`
	KodeWilayahKliring  sql.NullString `db:"KODE_WILAYAH_KLIRING" json:"kode_wilayah_kliring,omitempty"`
	KodeWilayah         sql.NullString `db:"KODE_WILAYAH" json:"kode_wilayah,omitempty"`
	IsKordinatorKliring sql.NullString `db:"IS_KORDINATOR_KLIRING" json:"is_kordinator_kliring,omitempty"`
	StatusCabang        sql.NullString `db:"STATUS_CABANG" json:"status_cabang,omitempty"`
	StatusAktif         sql.NullString `db:"STATUS_AKTIF" json:"status_aktif,omitempty"`
}

type GetLbvQuery struct {
	GlNo              sql.NullString  `excel:"B" db:"ACCOUNT_CODE"`
	GlName            sql.NullString  `excel:"C" db:"ACCOUNT_NAME"`
	Cabang            sql.NullString  `excel:"D" db:"BRANCH_CODE"`
	Valuta            sql.NullString  `excel:"E" db:"CURRENCY_CODE"`
	NominativeBalance sql.NullFloat64 `excel:"F" db:"BALANCE_NOMINATIVE"`
	GlBalance         sql.NullFloat64 `excel:"G" db:"BALANCECUMULATIVE"`
	Selisih           sql.NullFloat64 `excel:"H" db:"SELISIH"`

	AccountInstanceId int64 `db:"ACCOUNTINSTANCE_ID"`
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

type ParameterBiaya struct {
	AccountCode  sql.NullString  `db:"ACCOUNT_CODE" json:"account_code"`
	NominalBiaya sql.NullFloat64 `db:"NOMINAL_BIAYA" json:"nominal_biaya"`
}

type GLInterfaceTabungan struct {
	IdGlinterface int64  `db:"ID_GLINTERFACE" json:"id_glinterface,omitempty"`
	KodeInterface string `db:"KODE_INTERFACE" json:"kode_interface,omitempty"`
	Deskripsi     string `db:"DESKRIPSI" json:"deskripsi,omitempty"`
	KodeAccount   string `db:"KODE_ACCOUNT" json:"kode_account,omitempty"`
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

type ResGetPassBookLastType struct {
	KodePassbook string `db:"KODE_PASSBOOK" json:"kode_passbook"`
	NamaPassbook string `db:"NAMA_PASSBOOK" json:"nama_passbook"`
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
type ProdukTabungan struct {
	KodeProduk                sql.NullString  `db:"KODE_PRODUK" json:"kode_produk,omitempty"`
	NamaProduk                sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk,omitempty"`
	IsOnbalancesheetmode      sql.NullString  `db:"IS_ONBALANCESHEETMODE" json:"is_onbalancesheetmode,omitempty"`
	SaldoMinimumTidakAktif    sql.NullFloat64 `db:"SALDO_MINIMUM_TIDAK_AKTIF" json:"saldo_minimum_tidak_aktif,omitempty"`
	JumlahHariJadiTidakAktif  sql.NullFloat64 `db:"JUMLAH_HARI_JADI_TIDAK_AKTIF" json:"jumlah_hari_jadi_tidak_aktif,omitempty"`
	SaldoMinimum              sql.NullFloat64 `db:"SALDO_MINIMUM" json:"saldo_minimum,omitempty"`
	SaldoMaksimum             sql.NullFloat64 `db:"SALDO_MAKSIMUM" json:"saldo_maksimum,omitempty"`
	IsBackdated               sql.NullString  `db:"IS_BACKDATED" json:"is_backdated,omitempty"`
	JumlahBulanMaksBackdated  sql.NullFloat64 `db:"JUMLAH_BULAN_MAKS_BACKDATED" json:"jumlah_bulan_maks_backdated,omitempty"`
	IsKenaZakatBagiHasil      sql.NullString  `db:"IS_KENA_ZAKAT_BAGI_HASIL" json:"is_kena_zakat_bagi_hasil,omitempty"`
	PresentaseZakatBagiHasil  sql.NullFloat64 `db:"PRESENTASE_ZAKAT_BAGI_HASIL" json:"presentase_zakat_bagi_hasil,omitempty"`
	IsKenaPajak               sql.NullString  `db:"IS_KENA_PAJAK" json:"is_kena_pajak,omitempty"`
	TarifPajak                sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak,omitempty"`
	DisposisiBagiHasil        sql.NullString  `db:"DISPOSISI_BAGI_HASIL" json:"disposisi_bagi_hasil,omitempty"`
	PeriodeBagiHasil          sql.NullString  `db:"PERIODE_BAGI_HASIL" json:"periode_bagi_hasil,omitempty"`
	JenisProduk               sql.NullString  `db:"JENIS_PRODUK" json:"jenis_produk,omitempty"`
	KodeValuta                sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta,omitempty"`
	Status                    sql.NullString  `db:"STATUS" json:"status,omitempty"`
	IsProdukBerasuransi       sql.NullString  `db:"IS_PRODUK_BERASURANSI" json:"is_produk_berasuransi,omitempty"`
	IsProdukPartnership       sql.NullString  `db:"IS_PRODUK_PARTNERSHIP" json:"is_produk_partnership,omitempty"`
	IsOverrideBagihasilkhusus sql.NullString  `db:"IS_OVERRIDE_BAGIHASILKHUSUS" json:"is_override_bagihasilkhusus,omitempty"`
	IsOverrideTarifpajak      sql.NullString  `db:"IS_OVERRIDE_TARIFPAJAK" json:"is_override_tarifpajak,omitempty"`
	IsOverrideZakatbagihasil  sql.NullString  `db:"IS_OVERRIDE_ZAKATBAGIHASIL" json:"is_override_zakatbagihasil,omitempty"`
	IsOverrideDisposisiBgh    sql.NullString  `db:"IS_OVERRIDE_DISPOSISI_BGH" json:"is_override_disposisi_bgh,omitempty"`
	JarakPeriodeBagihasil     sql.NullFloat64 `db:"JARAK_PERIODE_BAGIHASIL" json:"jarak_periode_bagihasil,omitempty"`
	IsOverrideBackdated       sql.NullString  `db:"IS_OVERRIDE_BACKDATED" json:"is_override_backdated,omitempty"`
	IdRencanabagihasildefault sql.NullFloat64 `db:"ID_RENCANABAGIHASILDEFAULT" json:"id_rencanabagihasildefault,omitempty"`
	IsKenaBiayalayananumum    sql.NullString  `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum,omitempty"`
	IdBiayalayananumum        sql.NullFloat64 `db:"ID_BIAYALAYANANUMUM" json:"id_biayalayananumum,omitempty"`
	NamaValuta                sql.NullString  `db:"NAMA_VALUTA" json:"nama_valuta,omitempty"`
	IsProdukkartu             sql.NullString  `db:"IS_PRODUKKARTU" json:"is_produkkartu,omitempty"`
	NamaAsuransi              sql.NullString  `db:"NAMA_ASURANSI" json:"nama_asuransi,omitempty"`
	IdPartner                 sql.NullString  `db:"ID_PARTNER" json:"id_partner,omitempty"`
	NamaGrup                  sql.NullString  `db:"NAMA_GRUP" json:"nama_grup,omitempty"`
	Insplanid                 sql.NullFloat64 `db:"INSPLANID" json:"insplanid,omitempty"`
	Psplanid                  sql.NullFloat64 `db:"PSPLANID" json:"psplanid,omitempty"`
	BiayaAdmBulanan           sql.NullFloat64 `db:"BIAYA_ADM_BULANAN" json:"biaya_adm_bulanan,omitempty"`
	IdTieringBiayaadm         sql.NullInt32   `db:"ID_TIERING_BIAYAADM" json:"id_tiering_biayaadm,omitempty"`
	IdTieringNisbahbonus      sql.NullInt32   `db:"ID_TIERING_NISBAHBONUS" json:"id_tiering_nisbahbonus,omitempty"`
	IsTieringBiayaadm         sql.NullString  `db:"IS_TIERING_BIAYAADM" json:"is_tiering_biayaadm,omitempty"`
	IsTieringNisbahbonus      sql.NullString  `db:"IS_TIERING_NISBAHBONUS" json:"is_tiering_nisbahbonus,omitempty"`
	JenisAkad                 sql.NullString  `db:"JENIS_AKAD" json:"jenis_akad,omitempty"`
	KodeRencanaAsuransi       sql.NullString  `db:"KODE_RENCANA_ASURANSI" json:"kode_rencana_asuransi,omitempty"`
	NisbahBonusDasar          sql.NullFloat64 `db:"NISBAH_BONUS_DASAR" json:"nisbah_bonus_dasar,omitempty"`
	SetoranAwalMinimum        sql.NullFloat64 `db:"SETORAN_AWAL_MINIMUM" json:"setoran_awal_minimum,omitempty"`
	IdAsuransi                sql.NullFloat64 `db:"ID_ASURANSI" json:"id_asuransi,omitempty"`
	IsParamSaldoTidakAktif    sql.NullString  `db:"IS_PARAM_SALDO_TIDAK_AKTIF" json:"is_param_saldo_tidak_aktif,omitempty"`
	IsOverrideBiayaPenutupan  sql.NullString  `db:"IS_OVERRIDE_BIAYA_PENUTUPAN" json:"is_override_biaya_penutupan,omitempty"`
	BiayaPenutupanRekening    sql.NullFloat64 `db:"BIAYA_PENUTUPAN_REKENING" json:"biaya_penutupan_rekening,omitempty"`
	SumberBiayaAdm            sql.NullFloat64 `db:"SUMBER_BIAYA_ADM" json:"sumber_biaya_adm,omitempty"`
	JumlahHariPertahun        sql.NullFloat64 `db:"JUMLAH_HARI_PERTAHUN" json:"jumlah_hari_pertahun,omitempty"`
	KodeDigit                 sql.NullString  `db:"KODE_DIGIT" json:"kode_digit,omitempty"`
	IsBiayaSaldoMinimum       sql.NullString  `db:"IS_BIAYA_SALDO_MINIMUM" json:"is_biaya_saldo_minimum,omitempty"`
	IsBiayaRekeningDormant    sql.NullString  `db:"IS_BIAYA_REKENING_DORMANT" json:"is_biaya_rekening_dormant,omitempty"`
	IsBiayaAtm                sql.NullString  `db:"IS_BIAYA_ATM" json:"is_biaya_atm,omitempty"`
	IsCetakNota               sql.NullString  `db:"IS_CETAK_NOTA" json:"is_cetak_nota,omitempty"`
	IsStatusPassbook          sql.NullString  `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook,omitempty"`
	Kolektibilitas            sql.NullString  `db:"KOLEKTIBILITAS" json:"kolektibilitas,omitempty"`
	IsDapatBagiHasil          sql.NullString  `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil,omitempty"`
	IsTidakDormant            sql.NullString  `db:"IS_TIDAK_DORMANT" json:"is_tidak_dormant,omitempty"`
	TanggalAcuanDormant       sql.NullString  `db:"TANGGAL_ACUAN_DORMANT" json:"tanggal_acuan_dormant,omitempty"`
	BiayaRekeningDormant      sql.NullFloat64 `db:"BIAYA_REKENING_DORMANT" json:"biaya_rekening_dormant,omitempty"`
	BiayaSaldoMinimum         sql.NullFloat64 `db:"BIAYA_SALDO_MINIMUM" json:"biaya_saldo_minimum,omitempty"`
	JumlahHariTutupOtomatis   sql.NullFloat64 `db:"JUMLAH_HARI_TUTUP_OTOMATIS" json:"jumlah_hari_tutup_otomatis,omitempty"`
	IsBlokirDebet             sql.NullString  `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet,omitempty"`
	IsBlokirKredit            sql.NullString  `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit,omitempty"`
	IsTutupOtomatisDormant    sql.NullString  `db:"IS_TUTUP_OTOMATIS_DORMANT" json:"is_tutup_otomatis_dormant,omitempty"`
	BiayaAdmAtm               sql.NullFloat64 `db:"BIAYA_ADM_ATM" json:"biaya_adm_atm,omitempty"`
	BiayaKartuAtm             sql.NullFloat64 `db:"BIAYA_KARTU_ATM" json:"biaya_kartu_atm,omitempty"`
	IsBiayaMaterai            sql.NullString  `db:"IS_BIAYA_MATERAI" json:"is_biaya_materai,omitempty"`
	IsOverrideBiayaatm        sql.NullString  `db:"IS_OVERRIDE_BIAYAATM" json:"is_override_biayaatm,omitempty"`
	KodePofDefault            sql.NullString  `db:"KODE_POF_DEFAULT" json:"kode_pof_default,omitempty"`
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

type ResGetListProdSav struct {
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

type ResGetProdDepList struct {
	Produk
	KodeCabang        sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
	MasaPerjanjian    sql.NullString `db:"MASA_PERJANJIAN" json:"masa_perjanjian"`
	PeriodePerjanjian sql.NullString `db:"PERIODE_PERJANJIAN" json:"periode_perjanjian"`
}

type ResGetListCA struct {
	Produk
	KodeCabang sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
}

type LimitProduk2 struct {
	KodeLimit       string  `db:"KODE_LIMIT" json:"kode_limit,omitempty"`
	JenisLimitasi   string  `db:"JENIS_LIMITASI" json:"jenis_limitasi,omitempty"`
	Periode         string  `db:"PERIODE" json:"periode,omitempty"`
	JumlahTransaksi int64   `db:"JUMLAH_TRANSAKSI" json:"jumlah_transaksi,omitempty"`
	NilaiTransaksi  float64 `db:"NILAI_TRANSAKSI" json:"nilai_transaksi,omitempty"`
}

type ResGetRestrictCode struct {
	IdCfgroup   int32  `db:"ID_CFGROUP" json:"id_cfgroup"`
	CfCode      string `db:"CF_CODE" json:"cf_code"`
	Description string `db:"DESCRIPTION" json:"description"`
}

type GetSknTrxcodeListRes struct {
	KodeTxSkn sql.NullString `db:"KODE_TX_SKN"`
	NamaTxSkn sql.NullString `db:"NAMA_TX_SKN"`
}

type TieringBagiHasil struct {
	Itemid          int64   `db:"ITEMID" json:"itemid,omitempty"`
	Lowerlimitvalue float64 `db:"LOWERLIMITVALUE" json:"lowerlimitvalue,omitempty"`
	Upperlimitvalue float64 `db:"UPPERLIMITVALUE" json:"upperlimitvalue,omitempty"`
	Amount1         float64 `db:"AMOUNT1" json:"amount1,omitempty"`
}

type TieringBiayaADMProduk struct {
	Itemid          int64   `db:"ITEMID" json:"itemid,omitempty"`
	Lowerlimitvalue float64 `db:"LOWERLIMITVALUE" json:"lowerlimitvalue,omitempty"`
	Upperlimitvalue float64 `db:"UPPERLIMITVALUE" json:"upperlimitvalue,omitempty"`
	Amount1         float64 `db:"AMOUNT1" json:"amount1,omitempty"`
}
