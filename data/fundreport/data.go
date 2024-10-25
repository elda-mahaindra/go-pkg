package data

import "database/sql"

type GetAcctSavplanCertificateRes struct {
	NomorRekening             string          `db:"NOMOR_REKENING"`
	IsBlokirDebet             sql.NullString  `db:"IS_BLOKIR_DEBET"`
	IsBlokirKredit            sql.NullString  `db:"IS_BLOKIR_KREDIT"`
	Saldo                     sql.NullFloat64 `db:"SALDO"`
	TanggalBuka               sql.NullString  `db:"TANGGAL_BUKA"`
	TanggalJatuhTempo         sql.NullString  `db:"TANGGAL_JATUH_TEMPO"`
	NamaRekening              sql.NullString  `db:"NAMA_REKENING"`
	NomorRekeningInduk        string          `db:"NOMOR_REKENING_INDUK"`
	NamaRekeningInduk         sql.NullString  `db:"NAMA_REKENING_INDUK"`
	NomorRekeningPencairan    sql.NullString  `db:"NOMOR_REKENING_PENCAIRAN"`
	NamaRekeningPencairan     sql.NullString  `db:"NAMA_REKENING_PENCAIRAN"`
	Alamat                    sql.NullString  `db:"ALAMAT_KIRIM"`
	SetoranRutin              sql.NullFloat64 `db:"SETORAN_RUTIN"`
	TanggalSetoranRutin       sql.NullInt32   `db:"TANGGAL_SETORAN_RUTIN"`
	NomorNasabah              sql.NullString  `db:"NOMOR_NASABAH"`
	StatusStandingInstruction sql.NullString  `db:"STATUS_STANDING_INSTRUCTION"`
	JenisNasabah              sql.NullString  `db:"JENIS_NASABAH"`
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
type ListDHNRes struct {
	NomorNasabah    sql.NullString `db:"NOMOR_NASABAH"`
	NamaNasabah     sql.NullString `db:"NAMA_NASABAH"`
	StatusCif       sql.NullString `db:"STATUS_CIF"`
	EStatusCif      sql.NullString `db:"E_STATUS_CIF"`
	Gelar           sql.NullString `db:"GELAR"`
	JenisNasabah    sql.NullString `db:"JENIS_NASABAH"`
	NomorIdentitas  sql.NullString `db:"NOMOR_IDENTITAS"`
	TanggalLahir    sql.NullString `db:"TANGGAL_LAHIR"`
	NomorNpwp       sql.NullString `db:"NOMOR_NPWP"`
	IsDhib          sql.NullString `db:"IS_DHIB"`
	IsDhn           sql.NullString `db:"IS_DHN"`
	CtrMatchDhn     sql.NullString `db:"CTR_MATCH_DHN"`
	NoDhn           sql.NullString `db:"NO_DHN"`
	BatasSanksi     sql.NullString `db:"BATAS_SANKSI"`
	NomorRefPelapor sql.NullString `db:"NOMOR_REF_PELAPOR"`
	EJenisNasabah   sql.NullString `db:"E_JENIS_NASABAH"`
	KodeCabangCif   sql.NullString `db:"KODE_CABANG_CIF"`
}
type DhibData struct {
	IdDhn               sql.NullInt64  `db:"ID_DHN"`
	NoNasabah           sql.NullString `db:"NO_NASABAH"`
	NamaNasabah         sql.NullString `db:"NAMA_NASABAH"`
	JenisNasabah        sql.NullInt64  `db:"JENIS_NASABAH"`
	NomorIdentitas      sql.NullString `db:"NOMOR_IDENTITAS"`
	TglLahir            sql.NullString `db:"TGL_LAHIR"`
	Gelar               sql.NullString `db:"GELAR"`
	IsBbl               sql.NullString `db:"IS_BBL"`
	AlamatJalan         sql.NullString `db:"ALAMAT_JALAN"`
	AlamatRt            sql.NullString `db:"ALAMAT_RT"`
	AlamatRw            sql.NullString `db:"ALAMAT_RW"`
	AlamatKelurahan     sql.NullString `db:"ALAMAT_KELURAHAN"`
	AlamatKecamatan     sql.NullString `db:"ALAMAT_KECAMATAN"`
	AlamatKotaKabupaten sql.NullString `db:"ALAMAT_KOTA_KABUPATEN"`
	AlamatProvinsi      sql.NullString `db:"ALAMAT_PROVINSI"`
	AlamatKodePos       sql.NullString `db:"ALAMAT_KODE_POS"`
}
type DhnData struct {
	IdDhn               sql.NullInt64  `db:"ID_DHN"`
	NoDhn               sql.NullString `db:"NO_DHN"`
	NoUrut              sql.NullString `db:"NO_URUT"`
	NomorRef            sql.NullString `db:"NOMOR_REF"`
	SandiBankAsal       sql.NullString `db:"SANDI_BANK_ASAL"`
	NamaBankAsal        sql.NullString `db:"NAMA_BANK_ASAL"`
	NamaNasabah         sql.NullString `db:"NAMA_NASABAH"`
	StatusUsaha         sql.NullString `db:"STATUS_USAHA"`
	JenisNasabah        sql.NullInt64  `db:"JENIS_NASABAH"`
	NomorIdentitas      sql.NullString `db:"NOMOR_IDENTITAS"`
	NomorNpwp           sql.NullString `db:"NOMOR_NPWP"`
	TglLahir            sql.NullString `db:"TGL_LAHIR"`
	Keterangan          sql.NullString `db:"KETERANGAN"`
	BatasSanksi         sql.NullString `db:"BATAS_SANKSI"`
	AlamatJalan         sql.NullString `db:"ALAMAT_JALAN"`
	AlamatRt            sql.NullString `db:"ALAMAT_RT"`
	AlamatRw            sql.NullString `db:"ALAMAT_RW"`
	AlamatKelurahan     sql.NullString `db:"ALAMAT_KELURAHAN"`
	AlamatKecamatan     sql.NullString `db:"ALAMAT_KECAMATAN"`
	AlamatKotaKabupaten sql.NullString `db:"ALAMAT_KOTA_KABUPATEN"`
	AlamatProvinsi      sql.NullString `db:"ALAMAT_PROVINSI"`
	AlamatKodePos       sql.NullString `db:"ALAMAT_KODE_POS"`
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
type AccountPassbookData struct {
	IdRegister             sql.NullInt64   `db:"ID_REGISTER" json:"id_register"`
	NomorNasabah           sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening           sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeCabang             sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	NamaCabang             sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
	NomorRegisterBukuAktif sql.NullString  `db:"NOMORREGISTERBUKU_AKTIF" json:"nomor_register_buku_aktif"`
	AlamatKirim            sql.NullString  `db:"ALAMAT_KIRIM" json:"alamat_kirim"`
	JenisNasabah           sql.NullString  `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	TanggalBuka            sql.NullString  `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	TanggalRegistrasi      sql.NullString  `db:"TANGGAL_REGISTRASI" json:"tanggal_registrasi"`
	HalamanCetakTerakhir   sql.NullInt64   `db:"HALAMAN_CETAK_TERAKHIR" json:"halaman_cetak_terakhir"`
	BarisCetakTerakhir     sql.NullInt64   `db:"BARIS_CETAK_TERAKHIR" json:"baris_cetak_terakhir"`
	NomorBukuTerakhir      sql.NullInt64   `db:"NOMOR_BUKU_TERAKHIR" json:"nomor_buku_terakhir"`
	SaldoCetakTerakhir     sql.NullFloat64 `db:"SALDO_CETAK_TERAKHIR" json:"saldo_cetak_terakhir"`
	NamaNasabah            sql.NullString  `db:"NAMA_NASABAH" json:"nama_nasabah"`
	IsGantiBuku            sql.NullString  `db:"IS_GANTI_BUKU" json:"is_ganti_buku"`
	IsCetakCover           sql.NullString  `db:"IS_CETAK_COVER" json:"is_cetak_cover"`
	IdDetilTransaksi       sql.NullInt64   `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
}
type CekNomorRegisterBuku struct {
	NomorRegisterBuku string `db:"NOMOR_REGISTER_BUKU" json:"nomor_register_buku"`
	NomorRekening     string `db:"NOMOR_REKENING" json:"nomor_rekening"`
}
type AlternateAddress struct {
	NomorRekening       sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorNasabah        sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	AlamatJalan         sql.NullString `db:"ALAMAT_JALAN" json:"alamat_jalan"`
	AlamatKecamatan     sql.NullString `db:"ALAMAT_KECAMATAN" json:"alamat_kecamatan"`
	AlamatRtRw          sql.NullString `db:"ALAMAT_RTRW" json:"alamat_rtrw"`
	AlamatKelurahan     sql.NullString `db:"ALAMAT_KELURAHAN" json:"alamat_kelurahan"`
	AlamatProvinsi      sql.NullString `db:"ALAMAT_PROVINSI" json:"alamat_provinsi"`
	AlamatKodePos       sql.NullString `db:"ALAMAT_KODE_POS" json:"alamat_kode_pos"`
	AlamatKotaKabupaten sql.NullString `db:"ALAMAT_KOTA_KABUPATEN" json:"alamat_kota_kabupaten"`
}
type CustomerAddress struct {
	AlamatSuratJalan         sql.NullString `db:"ALAMAT_SURAT_JALAN" json:"alamat_surat_jalan"`
	AlamatSuratKodePos       sql.NullString `db:"ALAMAT_SURAT_KODE_POS" json:"alamat_surat_kode_pos"`
	AlamatSuratRtRw          sql.NullString `db:"ALAMAT_SURAT_RTRW" json:"alamat_surat_rtrw"`
	AlamatSuratKelurahan     sql.NullString `db:"ALAMAT_SURAT_KELURAHAN" json:"alamat_surat_kelurahan"`
	AlamatSuratKecamatan     sql.NullString `db:"ALAMAT_SURAT_KECAMATAN" json:"alamat_surat_kecamatan"`
	AlamatSuratKotaKabupaten sql.NullString `db:"ALAMAT_SURAT_KOTA_KABUPATEN" json:"alamat_surat_kota_kabupaten"`
	AlamatSuratProvinsi      sql.NullString `db:"ALAMAT_SURAT_PROVINSI" json:"alamat_surat_provinsi"`
}
type CustomerPersonalAddress struct {
	AlamatRumahJalan         sql.NullString `db:"ALAMAT_RUMAH_JALAN" json:"alamat_rumah_jalan"`
	AlamatRumahKodePos       sql.NullString `db:"ALAMAT_RUMAH_KODE_POS" json:"alamat_rumah_kode_pos"`
	AlamatRumahRtRw          sql.NullString `db:"ALAMAT_RUMAH_RTRW" json:"alamat_rumah_rtrw"`
	AlamatRumahKelurahan     sql.NullString `db:"ALAMAT_RUMAH_KELURAHAN" json:"alamat_rumah_kelurahan"`
	AlamatRumahKecamatan     sql.NullString `db:"ALAMAT_RUMAH_KECAMATAN" json:"alamat_rumah_kecamatan"`
	AlamatRumahKotaKabupaten sql.NullString `db:"ALAMAT_RUMAH_KOTA_KABUPATEN" json:"alamat_rumah_kota_kabupaten"`
	AlamatRumahProvinsi      sql.NullString `db:"ALAMAT_RUMAH_PROVINSI" json:"alamat_rumah_provinsi"`
}
type CustomerCorporateAddress struct {
	AlamatJalan         sql.NullString `db:"ALAMAT_JALAN" json:"alamat_jalan"`
	AlamatKodePos       sql.NullString `db:"ALAMAT_KODE_POS" json:"alamat_kode_pos"`
	AlamatRtRw          sql.NullString `db:"ALAMAT_RTRW" json:"alamat_rtrw"`
	AlamatKelurahan     sql.NullString `db:"ALAMAT_KELURAHAN" json:"alamat_kelurahan"`
	AlamatKecamatan     sql.NullString `db:"ALAMAT_KECAMATAN" json:"alamat_kecamatan"`
	AlamatKotaKabupaten sql.NullString `db:"ALAMAT_KOTA_KABUPATEN" json:"alamat_kota_kabupaten"`
	AlamatProvinsi      sql.NullString `db:"ALAMAT_PROVINSI" json:"alamat_provinsi"`
}
type JenisPassbook struct {
	TotalBaris    sql.NullInt64 `db:"TOTAL_BARIS"`
	JumlahHalaman sql.NullInt64 `db:"JUMLAH_HALAMAN"`
}
type ResValNorekPassbook struct {
	StatusOtorisasi     sql.NullFloat64 `db:"STATUS_OTORISASI"`
	NomorRekening       sql.NullString  `db:"NOMOR_REKENING"`
	KodeTransaksi       sql.NullString  `db:"KODE_TRANSAKSI"`
	KodeCabangTransaksi sql.NullString  `db:"KODE_CABANG_TRANSAKSI"`
}
type ResDataTransaksi struct {
	IdDetailTransaksi sql.NullInt64   `db:"ID_DETIL_TRANSAKSI"`
	NilaiMutasi       sql.NullFloat64 `db:"NILAI_MUTASI"`
	JenisMutasi       sql.NullString  `db:"JENIS_MUTASI"`
	Keterangan        sql.NullString  `db:"KETERANGAN"`
	KodeTransaksi     sql.NullString  `db:"KODE_TRANSAKSI"`
	TanggalTransaksi  sql.NullString  `db:"TANGGAL_TRANSAKSI"`
}
type CollectDataTrx struct {
	IdDetailTransaksi sql.NullInt64   `db:"ID_DETIL_TRANSAKSI"`
	TanggalTransaksi  sql.NullString  `db:"TANGGAL_TRANSAKSI"`
	Keterangan        sql.NullString  `db:"KETERANGAN"`
	JenisMutasi       sql.NullString  `db:"JENIS_MUTASI"`
	JenisMutasiDesc   sql.NullString  `db:"JENIS_MUTASI_DESC"`
	NilaiMutasi       sql.NullFloat64 `db:"NILAI_MUTASI"`
	IdTransasksi      sql.NullInt64   `db:"ID_TRANSAKSI"`
	Halaman           sql.NullInt64   `db:"HALAMAN"`
	UserInput         sql.NullString  `db:"USER_INPUT"`
	KodeTransaksi     sql.NullString  `db:"KODE_TRANSAKSI"`
	IdCetakPassbook   sql.NullInt64   `db:"ID_CETAK_PASSBOOK"`
	Baris             sql.NullInt64   `db:"BARIS"`
	NoBuku            sql.NullInt64   `db:"NO_BUKU"`
}
type DictMapCodeTrx struct {
	KodeTrx         sql.NullString `db:"KODE_TRANSAKSI"`
	KodeTrxPassBook sql.NullString `db:"KODE_TRANS_PASSBOOK"`
}
type PassBookLastPrintRes struct {
	NomorRekening        string  `db:"NOMOR_REKENING"`
	NamaRekening         string  `db:"NAMA_REKENING"`
	NomorRegisterBuku    string  `db:"NOMOR_REGISTER_BUKU"`
	HalamanCetakTerakhir int32   `db:"HALAMAN_CETAK_TERAKHIR"`
	BarisCetakTerakhir   int32   `db:"BARIS_CETAK_TERAKHIR"`
	SaldoCetakTerakhir   float64 `db:"SALDO_CETAK_TERAKHIR"`
}
type GetReportCodeListResponse struct {
	KodeReport string `db:"KODE_REPORT"`
	NamaReport string `db:"NAMA_REPORT"`
}
type ResGetDataReportNasabah struct {
	NomorNasabah  sql.NullString `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaNasabah   sql.NullString `db:"NAMA_NASABAH" json:"nama_nasabah"`
	JenisNasabah  sql.NullString `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	UserInput     sql.NullString `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi sql.NullString `db:"USER_OTORISASI" json:"user_otorisasi"`
	FlagPepNi     sql.NullString `db:"FLAG_PEP_NI" json:"flag_pep_ni"`
	FlagNrtNi     sql.NullString `db:"FLAG_NRT_NI" json:"flag_nrt_ni"`
	FlagNrtNk     sql.NullString `db:"FLAG_NRT_NK" json:"flag_nrt_nk"`
	TanggalBuka   sql.NullString `db:"TANGGAL_BUKA" json:"tanggal_buka"`
}
type DataReportRekeningBaru struct {
	NomorRekening string  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening  string  `excel:"C" db:"NAMA_REKENING"`
	NomorNasabah  string  `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah   string  `excel:"E" db:"NAMA_NASABAH"`
	NamaProduk    string  `excel:"F" db:"NAMA_PRODUK"`
	SaldoMinimum  float64 `excel:"G" db:"SALDO_MINIMUM"`
	TanggalBuka   string  `excel:"H" db:"TANGGAL_BUKA"`
	IsJoinAccount string  `excel:"I" db:"IS_JOIN_ACCOUNT"`
	KodeProgram   string  `excel:"J" db:"KODE_PROGRAM"`
	UserInput     string  `excel:"K" db:"USER_INPUT"`
	UserOtorisasi string  `excel:"L" db:"USER_OTORISASI"`

	KodeProduk string `db:"KODE_PRODUK" excel:"-"`
	DataCount  int32  `db:"DATA_COUNT" excel:"-"`
}
type DataReportNasabahBaru struct {
	NomorNasabah  string `excel:"B" db:"NOMOR_NASABAH"`
	NamaNasabah   string `excel:"C" db:"NAMA_NASABAH"`
	JenisNasabah  string `excel:"D" db:"JENIS_NASABAH"`
	FlagPepNi     string `excel:"E" db:"FLAG_PEP_NI"`
	FlagNrt       string `excel:"F" db:"FLAG_NRT"`
	StatusResiko  string `excel:"G" db:"STATUS_RESIKO"`
	UserInput     string `excel:"H" db:"USER_INPUT"`
	UserOtorisasi string `excel:"I" db:"USER_OTORISASI"`
	TanggalBuka   string `db:"TANGGAL_BUKA"`
}
type DataReportSetoranKliring struct {
	NomorAplikasi     string  `excel:"B" db:"NOMOR_APLIKASI"`
	NilaiMutasi       float64 `excel:"C" db:"NILAI_MUTASI"`
	NomorReferensi    string  `excel:"D" db:"NOMOR_REFERENSI"`
	KodeBank          string  `excel:"E" db:"KODE_BANK"`
	NamaBank          string  `excel:"F" db:"NAMA_BANK"`
	NomorRekeningNota string  `excel:"G" db:"NOMOR_REKENING_NOTA"`
	KodeWarkat        string  `excel:"H" db:"KODE_WARKAT"`
	NomorRekening     string  `excel:"I" db:"NOMOR_REKENING"`
	NamaRekening      string  `excel:"J" db:"NAMA_REKENING"`
	UserInput         string  `excel:"K" db:"USER_INPUT"`
	UserOtorisasi     string  `excel:"L" db:"USER_OTORISASI"`
	IdDetilTransaksi  int64   `db:"ID_DETIL_TRANSAKSI"`
}
type DataReportRekeningTutup struct {
	NomorRekening string  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening  string  `excel:"C" db:"NAMA_REKENING"`
	NomorNasabah  string  `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah   string  `excel:"E" db:"NAMA_NASABAH"`
	NamaProduk    string  `excel:"F" db:"NAMA_PRODUK"`
	KodeValuta    string  `excel:"G" db:"KODE_VALUTA"`
	TanggalProses string  `excel:"H" db:"TANGGAL_PROSES"`
	Userinput     string  `excel:"I" db:"USERINPUT"`
	Userotorisasi string  `excel:"J" db:"USEROTORISASI"`
	KodeProduk    string  `db:"KODE_PRODUK"`
	SaldoMinimum  float64 `db:"SALDO_MINIMUM"`
	KodeCabang    string  `db:"KODE_CABANG"`
	DataCount     int32   `db:"DATA_COUNT"`
}
type DataReportRekeningBlokir struct {
	NomorRekening    string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening     string `excel:"C" db:"NAMA_REKENING"`
	NamaProduk       string `excel:"D" db:"NAMA_PRODUK"`
	KodeValuta       string `excel:"E" db:"KODE_VALUTA"`
	KeteranganBlokir string `excel:"F" db:"KETERANGAN_BLOKIR"`
	UserInput        string `excel:"G" db:"USER_INPUT"`
	SistemPemblokir  string `excel:"H" db:"SISTEM_PEMBLOKIR"`
	KodeCabang       string `db:"KODE_CABANG"`
	NamaNasabah      string `db:"NAMA_NASABAH"`
	NomorNasabah     string `db:"NOMOR_NASABAH"`
	KodeProduk       string `db:"KODE_PRODUK"`
	UserDualcontrol  string `db:"USER_DUALCONTROL"`
	TanggalUbah      string `db:"TANGGAL_UBAH"`
}
type DataReportRekeningHoldDana struct {
	NomorRekening     string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening      string `excel:"C" db:"NAMA_REKENING"`
	NamaProduk        string `excel:"D" db:"NAMA_PRODUK"`
	KodeValuta        string `excel:"E" db:"KODE_VALUTA"`
	NominalHold       string `excel:"F" db:"NOMINAL_HOLD"`
	AlasanHold        string `excel:"G" db:"ALASAN_HOLD"`
	TanggalKadaluarsa string `excel:"H" db:"TANGGAL_KADALUARSA"`
	UserHold          string `excel:"I" db:"USER_HOLD"`
	NamaNasabah       string `db:"NAMA_NASABAH"`
	NomorNasabah      string `db:"NOMOR_NASABAH"`
	KodeProduk        string `db:"KODE_PRODUK"`
}
type DataReportRegistrasiAutosave struct {
	TanggalRegistrasi   string `excel:"B" db:"TANGGAL_REGISTRASI"`
	NomorRekeningAsal   string `excel:"C" db:"NOMOR_REKENING_ASAL"`
	NamaRekeningAsal    string `excel:"D" db:"NAMA_REKENING_ASAL"`
	NomorRekeningTujuan string `excel:"E" db:"NOMOR_REKENING_TUJUAN"`
	NamaRekeningTujuan  string `excel:"F" db:"NAMA_REKENING_TUJUAN"`
	Description         string `excel:"G" db:"DESCRIPTION"`
	UserPembuat         string `excel:"H" db:"USER_PEMBUAT"`
}
type DataReportNonAktifAutosave struct {
	TanggalNonAktif string `excel:"B" db:"TANGGAL_NONAKTIF"`
	NoRekAsal       string `excel:"C" db:"NOMOR_REKENING_ASAL"`
	NamaRekAsal     string `excel:"D" db:"NAMA_REKENING_ASAL"`
	NoRekTujuan     string `excel:"E" db:"NOMOR_REKENING_TUJUAN"`
	NamaRekTujuan   string `excel:"F" db:"NAMA_REKENING_TUJUAN"`
	UserInput       string `excel:"G" db:"USER_INPUT_NONAKTIF"`
	UserOtorisasi   string `excel:"H" db:"USER_OTORISASI_NONAKTIF"`
}
type DataReportAlamatAlternatifNasabah struct {
	NomorRekening    string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening     string `excel:"C" db:"NAMA_REKENING"`
	NomorNasabah     string `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah      string `excel:"E" db:"NAMA_NASABAH"`
	KodeProduk       string `excel:"F" db:"KODE_PRODUK"`
	AlamatAlternatif string `excel:"G" db:"ALAMAT_ALTERNATIF"`
	AlamatEmail      string `excel:"H" db:"ALAMAT_EMAIL"`
	Telepon          string `excel:"I" db:"TELEPON"`
	NamaProduk       string `db:"NAMA_PRODUK"`
}
type DataReportRekeningUbahProduk struct {
	NomorRekening  string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening   string `excel:"C" db:"NAMA_REKENING"`
	Tanggal        string `excel:"D" db:"TANGGAL"`
	CabangRekening string `excel:"E" db:"CABANG_REKENING"`
	Saldo          string `excel:"F" db:"SALDO"`
	NomorJurnal    string `excel:"G" db:"NOMOR_JURNAL"`
	ProdukLama     string `excel:"H" db:"PRODUK_LAMA"`
	ProdukBaru     string `excel:"I" db:"PRODUK_BARU"`
	UserInput      string `excel:"J" db:"USER_INPUT"`
	UserOtorisasi  string `excel:"K" db:"USER_OTORISASI"`
	NamaCabang     string `db:"NAMA_CABANG"`
	NomorNasabah   string `db:"NOMOR_NASABAH"`
	CabangInput    string `db:"CABANG_INPUT"`
	IdTransaksi    string `db:"ID_TRANSAKSI"`
	Count          string `db:"COUNT"`
}
type DataReportTabunganBagihasilMudharabah struct {
	NomorRekening              string  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening               string  `excel:"C" db:"NAMA_REKENING"`
	Saldo                      string  `excel:"D" db:"SALDO"`
	Nisbah                     string  `excel:"E" db:"NISBAH"`
	Gdr                        string  `excel:"F" db:"GDR"`
	Nominal                    string  `excel:"G" db:"NOMINAL"`
	Pajak                      string  `excel:"H" db:"PAJAK"`
	RekeningDisposisiBagiHasil string  `excel:"I" db:"REKENING_DISPOSISI_BAGI_HASIL"`
	NomorNasabah               string  `db:"NOMOR_NASABAH"`
	NisbahBagiHasil            float64 `db:"NISBAH_BAGI_HASIL"`
	EqvRate                    float64 `db:"EQV_RATE"`
	PajakAfterBaghas           float64 `db:"PAJAK_AFTER_BAGHAS"`
	ZakatAfterBaghas           float64 `db:"ZAKAT_AFTER_BAGHAS"`
	Zakat                      string  `db:"ZAKAT"`
	KodeProduk                 string  `db:"KODE_PRODUK"`
	NamaProduk                 string  `db:"NAMA_PRODUK"`
	Rownumber                  int32   `db:"ROWNUMBER"`
}
type DataReportTabunganBonusWadiah struct {
	NomorRekening              string  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening               string  `excel:"C" db:"NAMA_REKENING"`
	Saldo                      float64 `excel:"D" db:"SALDO"`
	Nisbah                     string  `excel:"E" db:"NISBAH"`
	Nominal                    float64 `excel:"F" db:"NOMINAL"`
	Pajak                      float64 `excel:"G" db:"PAJAK"`
	Zakat                      float64 `excel:"H" db:"ZAKAT"`
	RekeningDisposisiBagiHasil string  `excel:"I" db:"REKENING_DISPOSISI_BAGI_HASIL"`
	SaldoDeposito              float64 `db:"SALDO_DEPOSITO"`
	NomorNasabah               string  `db:"NOMOR_NASABAH"`
	NisbahBagiHasil            float64 `db:"NISBAH_BAGI_HASIL"`
	Rownumber                  int32   `db:"ROWNUMBER"`
}
type DataReportBiayaAdministrasi struct {
	NomorRekening string  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening  string  `excel:"C" db:"NAMA_REKENING"`
	Nominal       float64 `excel:"D" db:"NOMINAL"`
}
type DataReportRekeningSaldoHarian struct {
	NomorRekening   string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening    string `excel:"C" db:"NAMA_REKENING"`
	KodeProgram     string `excel:"D" db:"KODE_PROGRAM"`
	IsBlokir        string `excel:"E" db:"IS_BLOKIR"`
	IsBolehDebet    string `excel:"F" db:"IS_BOLEH_DEBET"`
	IsBolehKredit   string `excel:"G" db:"IS_BOLEH_KREDIT"`
	Saldo           string `excel:"H" db:"SALDO"`
	SaldoMinimum    string `excel:"I" db:"SALDO_MINIMUM"`
	SaldoFloat      string `excel:"J" db:"SALDO_FLOAT"`
	SaldoDitahan    string `excel:"K" db:"SALDO_DITAHAN"`
	SaldoEfektif    string `excel:"L" db:"SALDO_EFEKTIF"`
	NomorNasabah    string `db:"NOMOR_NASABAH"`
	StatusRestriksi string `db:"STATUS_RESTRIKSI"`
	NamaNasabah     string `db:"NAMA_NASABAH"`
	KodeProduk      string `db:"KODE_PRODUK"`
	NamaProduk      string `db:"NAMA_PRODUK"`
	DataCount       string `db:"DATA_COUNT"`
	Rownumber       string `db:"ROWNUMBER"`
}
type DataReportRekeningReposisi struct {
	NomorRekening string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening  string `excel:"C" db:"NAMA_REKENING"`
	Produk        string `excel:"D" db:"PRODUK"`
	SaldoReposisi string `excel:"E" db:"SALDO_REPOSISI"`
	CabangAsal    string `excel:"F" db:"CABANG_ASAL"`
	CabangTujuan  string `excel:"G" db:"CABANG_TUJUAN"`
	UserInput     string `excel:"H" db:"USER_INPUT"`
	UserOtorisasi string `excel:"I" db:"USER_OTORISASI"`
	NomorNasabah  string `db:"NOMOR_NASABAH"`
	CabangInput   string `db:"CABANG_INPUT"`
}
type DataReportRekeningSaldoRekap struct {
	KodeProduk  string `excel:"B" db:"KODE_PRODUK"`
	NamaProduk  string `excel:"C" db:"NAMA_PRODUK"`
	JumlahAktif string `excel:"D" db:"JUMLAH_AKTIF"`
	JumlahTutup string `excel:"E" db:"JUMLAH_TUTUP"`
	JumlahTotal string `excel:"F" db:"JUMLAH_TOTAL"`
	Saldo       string `excel:"G" db:"SALDO"`
	KodeCabang  string `db:"KODE_CABANG"`
}
type DataReportRekeningTidakAktif struct {
	NomorRekening            string         `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening             string         `excel:"C" db:"NAMA_REKENING"`
	NomorNasabah             string         `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah              string         `excel:"E" db:"NAMA_NASABAH"`
	NamaCabang               string         `excel:"F" db:"NAMA_CABANG"`
	KodeValuta               string         `excel:"G" db:"KODE_VALUTA"`
	TanggalBuka              string         `excel:"H" db:"TANGGAL_BUKA"`
	TglTransCabangTerakhir   sql.NullString `excel:"I" db:"TGL_TRANS_CABANG_TERAKHIR"`
	TglTransEchannelTerakhir sql.NullString `excel:"J" db:"TGL_TRANS_ECHANNEL_TERAKHIR"`
	Saldo                    string         `excel:"K" db:"SALDO"`
	StatusResiko             string         `excel:"L" db:"STATUS_RESIKO"`
	KodeCabang               string         `db:"KODE_CABANG"`
	KodeProduk               string         `db:"KODE_PRODUK"`
	NamaProduk               string         `db:"NAMA_PRODUK"`
	Produk                   string         `db:"PRODUK"`
	SaldoMinimum             string         `db:"SALDO_MINIMUM"`
}
type DataReportBeaMaterai struct {
	NomorRekening string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening  string `excel:"C" db:"NAMA_REKENING"`
	Nominal       string `excel:"D" db:"NOMINAL"`
	Saldo         string `excel:"E" db:"SALDO"`
}
type DataReportTabunganCetakKepalaBuku struct {
	NomorRekening         string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening          string `excel:"C" db:"NAMA_REKENING"`
	Produk                string `excel:"D" db:"PRODUK"`
	NomorRegisterBuku     string `excel:"E" db:"NOMOR_REGISTER_BUKU"`
	NomorRegisterBukuLama string `excel:"F" db:"NOMOR_REGISTER_BUKU_LAMA"`
	AlasanCetak           string `excel:"G" db:"ALASAN_CETAK"`
	UserInput             string `excel:"H" db:"USER_INPUT"`
	UserOverride          string `excel:"I" db:"USER_OVERRIDE"`
	TglInput              string `excel:"J" db:"TGL_INPUT"`
	TglSistem             string `excel:"K" db:"TGL_SISTEM"`
	NomorNasabah          string `db:"NOMOR_NASABAH"`
	NamaNasabah           string `db:"NAMA_NASABAH"`
	KodeValuta            string `db:"KODE_VALUTA"`
}
type DataReportPemblokiranCekBG struct {
	NomorRekening string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening  string `excel:"C" db:"NAMA_REKENING"`
	NomorNasabah  string `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah   string `excel:"E" db:"NAMA_NASABAH"`
	Produk        string `excel:"F" db:"PRODUK"`
	NomorSeri     string `excel:"G" db:"NOMOR_SERI"`
	UserPengubah  string `excel:"H" db:"USER_PENGUBAH"`
}
type DataReportAktivasiWarkat struct {
	NomorRekening   string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening    string `excel:"C" db:"NAMA_REKENING"`
	Produk          string `excel:"D" db:"PRODUK"`
	JenisWarkat     string `excel:"E" db:"JENIS_WARKAT"`
	LembarTersedia  string `excel:"F" db:"LEMBAR_TERSEDIA"`
	NomorSeriAwal   string `excel:"G" db:"NOMOR_SERI_AWAL"`
	NomorSeriAkhir  string `excel:"H" db:"NOMOR_SERI_AKHIR"`
	StatusOtorisasi string `excel:"I" db:"STATUS_OTORISASI"`
	StatusResi      string `excel:"J" db:"STATUS_RESI"`
	UserAktivasi    string `excel:"K" db:"USER_AKTIVASI"`
	UserOtorisasi   string `excel:"L" db:"USER_OTORISASI_AKTIVASI"`
	KodeCabang      string `db:"KODE_CABANG"`
}
type DataReportRekeningDormanOtomatis struct {
	NomorRekening                string `excel:"B" db:"NOMOR_REKENING"`
	Nama                         string `excel:"C" db:"NAMA_REKENING"`
	Produk                       string `excel:"D" db:"PRODUK"`
	TglProses                    string `excel:"E" db:"TANGGAL_BUKA"`
	TglTransaksiCounterTerakhir  string `excel:"F" db:"TGL_TRANS_CABANG_TERAKHIR"`
	TglTransaksiDChannelTerakhir string `excel:"G" db:"TGL_TRANS_ECHANNEL_TERAKHIR"`
}
type DataReportRegistrasiATSPerTanggal struct {
	NomorNamaRekeningDebit  string `excel:"B" db:"DEBITACCOUNTNO"`
	NomorNamaRekeningKredit string `excel:"C" db:"CREDITACCOUNTNO"`
	Nominal                 string `excel:"D" db:"AMOUNT"`
	KodeCabang              string `excel:"E" db:"BRANCHCODE"`
	TglProses               string `excel:"F" db:"TGL_PROSES"`
	Keterangan              string `excel:"G" db:"DESCRIPTION"`
	TanggalKadarluarsa      string `excel:"H" db:"TGL_KADALUARSA"`
	UserInput               string `excel:"I" db:"USER_PEMBUAT"`
	UserOtorisasi           string `excel:"J" db:"USER_OTORISASI"`
}
type DataReportNonAktifATSPerTanggal struct {
	NomorNamaRekeningDebit  string `excel:"B" db:"DEBITACCOUNTNO"`
	NomorNamaRekeningKredit string `excel:"C" db:"CREDITACCOUNTNO"`
	Nominal                 string `excel:"D" db:"AMOUNT"`
	KodeCabang              string `excel:"E" db:"KODE_CABANG_NONAKTIF"`
	TglNonAktif             string `excel:"F" db:"TANGGAL_NONAKTIF"`
	UserInput               string `excel:"G" db:"USER_INPUT_NONAKTIF"`
	UserOtorisasi           string `excel:"H" db:"USER_OTORISASI_NONAKTIF"`
}
type DataReportRekeningJoinAccount struct {
	NoRekening       string `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening     string `excel:"C" db:"NAMA_REKENING"`
	TanggalBuka      string `excel:"D" db:"TANGGAL_BUKA"`
	NoCIFUtama       string `excel:"E" db:"NOMOR_NASABAH_UTAMA"`
	NamaNasabahUtama string `excel:"F" db:"NAMA_NASABAH_UTAMA"`
	NoCIFJoin        string `excel:"G" db:"NOMOR_NASABAH_JOIN"`
	NamaNasabahJoin  string `excel:"H" db:"NAMA_NASABAH_JOIN"`
	StatusRekProduk  string `excel:"I" db:"STATUS"`
}
type DataReportWICBaru struct {
	NomorWIC       string `excel:"B" db:"NOMOR_NASABAH"`
	NamaWIC        string `excel:"C" db:"NAMA_NASABAH"`
	JenisWIC       string `excel:"D" db:"JENIS_NASABAH"`
	JenisIdentitas string `excel:"E" db:"JENIS_IDENTITAS"`
	NomorIdentitas string `excel:"F" db:"NOMOR_IDENTITAS"`
	Alamat         string `excel:"G" db:"ALAMAT"`
	TanggalInput   string `excel:"H" db:"TANGGAL_BUKA"`
	UserInput      string `excel:"I" db:"USER_INPUT"`
	UserOtorisasi  string `excel:"J" db:"USER_OTORISASI"`
}
type DataReportDepositoJatuhTempoHariIni struct {
	NomorDeposito             string `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet               string `excel:"C" db:"NOMOR_BILYET"`
	NamaRekeningDeposito      string `excel:"D" db:"NAMA_REKENING"`
	NomorCIF                  string `excel:"E" db:"NOMOR_NASABAH"`
	JumlahPokok               string `excel:"F" db:"SALDO_DEPOSITO"`
	NisbahDasar               string `excel:"G" db:"NISBAH_BAGI_HASIL"`
	NisbahSpesial             string `excel:"H" db:"NISBAH_SPESIAL"`
	TglBuka                   string `excel:"I" db:"TANGGAL_BUKA"`
	TglMulai                  string `excel:"J" db:"TANGGAL_JATUH_TEMPO_TERAKHIR"`
	TglJatuhTempo             string `excel:"K" db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA"`
	JangkaWaktu               string `excel:"L" db:"MASA_PERJANJIAN"`
	DispPokok                 string `excel:"M" db:"DISPOSISI_NOMINAL"`
	NomorRekeningDisposisi    string `excel:"N" db:"REKENING_DISPOSISI"`
	PenerimaTransferDisposisi string `excel:"O" db:"PENERIMA_TRANSFER_DISPOSISI"`
	DispBagiHasil             string `excel:"P" db:"DISPOSISI_BAGI_HASIL"`
	NomorRekeningBagiHasil    string `excel:"Q" db:"REKENING_DISPOSISI_BAGI_HASIL"`
	NamaRekeningBagiHasil     string `excel:"R" db:"PENERIMA_TRANSFER_BAGI_HASIL"`
	NamaRekening              string `db:"NAMA_REKENING"`
}
type DataReportDepositoBagiHasil struct {
	NomorDeposito          string `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet            string `excel:"C" db:"NOMOR_BILYET"`
	NamaRekeningDeposito   string `excel:"D" db:"NAMA_REKENING"`
	SaldoDeposito          string `excel:"E" db:"SALDO_DEPOSITO"`
	PeriodeBagiHasil       string `excel:"F" db:"PERIODE_BAGI_HASIL"`
	ECR                    string `excel:"G" db:"EKUIVALEN_RATE"`
	BagiHasilNisbah        string `excel:"H" db:"NOMINAL_BAGIHASIL"`
	BagiHasilECR           string `excel:"I" db:"NOMINAL_RATE"`
	Pajak                  string `excel:"J" db:"PAJAK"`
	Zakat                  string `excel:"K" db:"ZAKAT"`
	BagiHasilBersih        string `excel:"L" db:"NOMINAL_BERSIH"`
	DisposisiBagiHasil     string `excel:"M" db:"DISPOSISI_BAGI_HASIL"`
	NomorRekeningDisposisi string `excel:"N" db:"NOMOR_REKENING_DISPOSISI"`
	Bank                   string `excel:"O" db:"NAMA_LENGKAP_BANK"`
	NomorNasarabah         string `db:"NOMOR_NASABAH"`
}

type DataReportTplDepositoBaru struct {
	NomorDeposito             string  `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet               string  `excel:"C" db:"NOMOR_BILYET"`
	NamaNasabah               string  `excel:"D" db:"NAMA_REKENING"`
	JumlahPokok               float64 `excel:"E" db:"SALDO"`
	NisbahProduk              string  `excel:"F" db:"KODE_NAMA_PRODUK"`
	NisbahDasar               float64 `excel:"G" db:"NISBAH_DASAR"`
	NisbahSpesial             float64 `excel:"H" db:"NISBAH_SPESIAL"`
	TanggalBuka               string  `excel:"I" db:"TANGGAL_BUKA"`
	TglJatuhTempo             string  `excel:"J" db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA"`
	JangkaWaktu               string  `excel:"K" db:"MAPER_PERJANJIAN"`
	DisposisiPokok            string  `excel:"L" db:"DISPOSISI_NOMINAL"`
	RekeningDisposisiPokok    string  `excel:"M" db:"REKENING_DISPOSISI"`
	PenerimaTransferDisposisi string  `excel:"N" db:"PENERIMA_TRANSFER_DISPOSISI"`
	DisposisiBagiHasil        string  `excel:"O" db:"DISPOSISI_BAGI_HASIL"`
	NomorRekeningDisposisi    string  `excel:"P" db:"NOMOR_REKENING_DISPOSISI"`
	PenerimaTransferBagiHasil string  `excel:"Q" db:"PENERIMA_TRANSFER_BAGI_HASIL"`
	UserInput                 string  `excel:"R" db:"USER_INPUT"`
	UserOtorisasi             string  `excel:"S" db:"USER_OTORISASI"`

	NamaProduk        string  `db:"NAMA_PRODUK"`
	MasaPerjanjian    string  `db:"MASA_PERJANJIAN"`
	SaldoDeposito     float64 `db:"SALDO_DEPOSITO"`
	TglMulai          string  `db:"TANGGAL_JATUH_TEMPO_TERAKHIR"`
	KodeProduk        string  `db:"KODE_PRODUK"`
	NisbahBonusDasar  float64 `db:"NISBAH_BONUS_DASAR"`
	PeriodePerjanjian string  `db:"PERIODE_PERJANJIAN"`
	EkuivalenRate     float64 `db:"EKUIVALEN_RATE"`
	NisbahBagiHasil   float64 `db:"NISBAH_BAGI_HASIL"`
	DataCount         int64   `db:"DATA_COUNT"`
}

type DataReportTplPerubahanNasabah struct {
	Tipe         string `excel:"B" db:"ACCOUNT_TYPE"`
	NomorCif     string `excel:"C" db:"CUSTOMER_NO"`
	Field        string `excel:"D" db:"FIELD_NAME"`
	Sebelum      string `excel:"E" db:"BEFORE_DATA"`
	Sesudah      string `excel:"F" db:"AFTER_DATA"`
	InfoTambahan string `excel:"G" db:"DATA_REMARK"`
	TanggalUbah  string `excel:"H" db:"INPUT_DATE"`
	UserId       string `excel:"I" db:"USER_INPUT"`
	UserApprove  string `db:"USER_APPROVE"`
	TableName    string `db:"TABLE_NAME"`
	ApproveDate  string `db:"APPROVE_DATE"`
	FormName     string `db:"FORM_NAME"`
	FunctionName string `db:"FUNCTION_NAME"`
	BranchCode   string `db:"BRANCH_CODE"`
	AccountNo    string `db:"ACCOUNT_NO"`
}

type DataReportTplDepositoBreak struct {
	NomorDeposito         string  `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet           string  `excel:"C" db:"NOMOR_BILYET"`
	NamaRekeningDeposito  string  `excel:"D" db:"NAMA_REKENING"`
	NomorCif              string  `excel:"E" db:"NOMOR_NASABAH"`
	JumlahPokok           float64 `excel:"F" db:"NILAI_PENCAIRAN"`
	NisbahBagiHasil       float64 `excel:"G" db:"NISBAH_BAGI_HASIL"`
	TanggalBuka           string  `excel:"H" db:"TANGGAL_BUKA"`
	TanggalMulai          string  `excel:"I" db:"TANGGAL_JATUH_TEMPO_TERAKHIR"`
	TanggalJatuhTempo     string  `excel:"J" db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA"`
	JangkaWaktu           string  `excel:"K" db:"MAPER_PERJANJIAN"`
	DisposisiPokok        string  `excel:"L" db:"DISPOSISI_NOMINAL"`
	NamaPenerima          string  `excel:"M" db:"REKENING_DISPOSISI"`
	NomorRekeningPenerima string  `excel:"N" db:"PENERIMA_TRANSFER_DISPOSISI"`
	MasaPerjanjian        int64   `db:"MASA_PERJANJIAN"`
	SaldoDeposito         float64 `db:"SALDO_DEPOSITO"`
	KodeProduk            string  `db:"KODE_PRODUK"`
	NamaProduk            string  `db:"NAMA_PRODUK"`
	PeriodePerjanjian     string  `db:"PERIODE_PERJANJIAN"`
}

type DataReportTplRekRencanaBaru struct {
	NomorRekening     sql.NullString  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening      sql.NullString  `excel:"C" db:"NAMA_REKENING"`
	NomorCif          sql.NullString  `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah       sql.NullString  `excel:"E" db:"NAMA_NASABAH"`
	KodeNamaProduk    sql.NullString  `excel:"F" db:"KODE_NAMA_PRODUK"`
	TanggalBuka       sql.NullString  `excel:"G" db:"TANGGAL_BUKA"`
	TanggalJatuhTempo sql.NullString  `excel:"H" db:"TANGGAL_JATUH_TEMPO"`
	KodeProgram       sql.NullString  `excel:"I" db:"KODE_PROGRAM"`
	UserInput         sql.NullString  `excel:"J" db:"USER_INPUT"`
	UserOtorisasi     sql.NullString  `excel:"K" db:"USER_OTORISASI"`
	Produk            sql.NullString  `db:"NAMA_PRODUK"`
	KodeProduk        sql.NullString  `db:"KODE_PRODUK"`
	SaldoMinimum      sql.NullFloat64 `db:"SALDO_MINIMUM"`
}

type DataReportTplNasabahReposisi struct {
	NomorCif         string          `excel:"B" db:"NOMOR_NASABAH"`
	Nama             string          `excel:"C" db:"NAMA_NASABAH"`
	TanggalPerubahan string          `excel:"D" db:"TANGGAL"`
	CabangAsal       string          `excel:"E" db:"CABANG_ASAL"`
	CabangTujuan     string          `excel:"F" db:"CABANG_TUJUAN"`
	UserInput        string          `excel:"G" db:"USER_INPUT"`
	UserOtorisasi    string          `excel:"H" db:"USER_OTORISASI"`
	Count            sql.NullFloat64 `db:"COUNT"`
	CabangInput      sql.NullFloat64 `db:"CABANG_INPUT"`
}

type DataReportTplNasabahGabung struct {
	NomorCif          string        `excel:"B" db:"NOMOR_NASABAH"`
	Nama              string        `excel:"C" db:"NAMA_NASABAH"`
	Cabang            string        `excel:"D" db:"CABANG_NASABAH"`
	NomorNasabahUtama string        `excel:"E" db:"NOMOR_NASABAH_UTAMA"`
	NamaNasabahUtama  string        `excel:"F" db:"NAMA_NASABAH_UTAMA"`
	CabangUtama       string        `excel:"G" db:"CABANG_NASABAH_UTAMA"`
	CabangInput       string        `excel:"H" db:"KODE_CABANG_INPUT"`
	TanggalInput      string        `excel:"I" db:"TANGGAL_INPUT"`
	TanggalOtorisasi  string        `excel:"J" db:"TANGGAL_OTORISASI"`
	TanggalProses     string        `excel:"K" db:"TANGGAL_PROSES"`
	UserInput         string        `excel:"L" db:"USER_INPUT"`
	UserOtorisasi     string        `excel:"M" db:"USER_OTORISASI"`
	DataCount         sql.NullInt32 `db:"DATA_COUNT"`
}

type DataReportTplDepositoJatuhTempoHariIniTransfer struct {
	NomorDeposito               string `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet                 string `excel:"C" db:"NOMOR_BILYET"`
	NamaRekeningDeposito        string `excel:"D" db:"NAMA_REKENING"`
	NomorCIF                    string `excel:"E" db:"NOMOR_NASABAH"`
	JumlahPokok                 string `excel:"F" db:"SALDO_DEPOSITO"`
	NisbahBagiHasil             string `excel:"G" db:"NISBAH_BAGI_HASIL"`
	TanggalBuka                 string `excel:"H" db:"TANGGAL_BUKA"`
	TanggalMulai                string `excel:"I" db:"TANGGAL_JATUH_TEMPO_TERAKHIR"`
	MasaPerjanjian              string `excel:"J" db:"MASA_PERJANJIAN"`
	DisposisiPokok              string `excel:"K" db:"DISPOSISI_NOMINAL"`
	NomorRekeningDisposisi      string `excel:"L" db:"REKENING_DISPOSISI"`
	NamaPenerima                string `excel:"M" db:"PENERIMA_TRANSFER_DISPOSISI"`
	NamaBankPenerima            string `excel:"N" db:"NAMA_LENGKAP_BANK"`
	TanggalJatuhTempoBerikutnya string `db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA"`
	PeriodePerjanjian           string `db:"PERIODE_PERJANJIAN"`
	KodeBank                    string `db:"KODE_BANK"`
}
type DataReportTplDepositoBlokir struct {
	NomorRekening               string `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet                 string `excel:"C" db:"NOMOR_BILYET"`
	NamaRekening                string `excel:"D" db:"NAMA_REKENING"`
	NamkodProduk                string `excel:"E" db:"NAMKOD_PRODUK"`
	Valuta                      string `excel:"F" db:"KODE_VALUTA"`
	AlasanBlokir                string `excel:"G" db:"KETERANGAN_BLOKIR"`
	UserInput                   string `excel:"H" db:"USER_INPUT"`
	SistemPemblokir             string `excel:"I" db:"SISTEM_PEMBLOKIR"`
	NamaProduk                  string `db:"NAMA_PRODUK"`
	NomorNasabah                string `db:"NOMOR_NASABAH"`
	TanggalJatuhTempoTerakhir   string `db:"TANGGAL_JATUH_TEMPO_TERAKHIR"`
	TanggalJatuhTempoBerikutnya string `db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA"`
	DisposisiNominal            string `db:"DISPOSISI_NOMINAL"`
	MasaPerjanjian              string `db:"MASA_PERJANJIAN"`
	PeriodePerjanjian           string `db:"PERIODE_PERJANJIAN"`
	NamaNasabah                 string `db:"NAMA_NASABAH"`
	KodeProduk                  string `db:"KODE_PRODUK"`
	UserDualControl             string `db:"USER_DUALCONTROL"`
	TanggalUbah                 string `db:"TANGGAL_UBAH"`
	IdHistBlokir                string `db:"ID_HISTBLOKIR"`
}

type DataReportTplDepositoAROHariIni struct {
	NomorDeposito          string  `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet            string  `excel:"C" db:"NOMOR_BILYET"`
	NamaRekeningDeposito   string  `excel:"D" db:"NAMA_REKENING"`
	NomorCif               string  `excel:"E" db:"NOMOR_NASABAH"`
	JumlahPokok            float64 `excel:"F" db:"SALDO_DEPOSITO"`
	NisbahBagiHasil        float64 `excel:"G" db:"NISBAH_BAGI_HASIL"`
	DisposisiPokok         string  `excel:"H" db:"DISPOSISI_NOMINAL"`
	NomorRekeningDisposisi string  `excel:"I" db:"REKENING_DISPOSISI"`
	MasaPerjanjian         string  `db:"MASA_PERJANJIAN"`
	PeriodePerjanjian      string  `db:"PERIODE_PERJANJIAN"`
}

type DataReportTplDepositoJatuhTempoBesok struct {
	TglJatuhTempo        string `excel:"A" db:"TGL_JATUH_TEMPO"`
	NoDeposito           string `excel:"C" db:"NOMOR_REKENING"`
	NomorBilyet          string `excel:"D" db:"NOMOR_BILYET"`
	NamaRekeningDeposito string `excel:"E" db:"NAMA_REKENING"`
	NomorCIF             string `excel:"F" db:"NOMOR_NASABAH"`
	JumlahPokok          string `excel:"G" db:"SALDO_DEPOSITO"`
	NisbahBagiHasil      string `excel:"H" db:"NISBAH_BAGI_HASIL"`
	DisposisiPokok       string `excel:"I" db:"DISPOSISI_NOMINAL"`
	NoRekeningDisposisi  string `excel:"J" db:"REKENING_DISPOSISI"`
	MasaPerjanjian       string `db:"MASA_PERJANJIAN"`
	PeriodePerjanjian    string `db:"PERIODE_PERJANJIAN"`
	DataCount            string `db:"DATA_COUNT"`
}

type DataReportTplDepositoBagHasTransfer struct {
	NomorDeposito              string `excel:"B" db:"NOMOR_REKENING"`
	NomorBilyet                string `excel:"C" db:"NOMOR_BILYET"`
	NamaRekeningDeposito       string `excel:"D" db:"NAMA_REKENING"`
	NominalBagiHasil           string `excel:"E" db:"NOMINAL_TOTAL"`
	Pajak                      string `excel:"F" db:"PAJAK"`
	Zakat                      string `excel:"G" db:"ZAKAT"`
	BagiHasilBersih            string `excel:"H" db:"NOMINAL_BERSIH"`
	DisposisiBagiHasil         string `excel:"I" db:"DISPOSISI_BAGI_HASIL"`
	NomorRekeningDisposisi     string `excel:"J" db:"NOMOR_REKENING_DISPOSISI"`
	PenerimaBagiHasil          string `excel:"K" db:"PENERIMA_TRANSFER_BAGI_HASIL"`
	Bank                       string `excel:"L" db:"KODE_BANK"`
	SaldoDeposito              string `db:"SALDO_DEPOSITO"`
	NomorNasabah               string `db:"NOMOR_NASABAH"`
	NisbahBagiHasil            string `db:"NISBAH_BAGI_HASIL"`
	RekeningDisposisiBagiHasil string `db:"REKENING_DISPOSISI_BAGI_HASIL"`
	NamaLengkapBank            string `db:"NAMA_LENGKAP_BANK"`
}

type DataReportTplDepositoSaldoRekap struct {
	KodeProduk    string  `excel:"B" db:"KODE_PRODUK"`
	NamaProduk    string  `excel:"C" db:"NAMA_PRODUK"`
	Saldo         float64 `excel:"D" db:"SALDO"`
	JmlRekAktif   int     `excel:"E" db:"JUMLAH_AKTIF"`
	JmlRekTutup   int     `excel:"F" db:"JUMLAH_TUTUP"`
	TotalRekening int     `excel:"G" db:"JUMLAH_TOTAL"`
	KodeCabang    string  `db:"KODE_CABANG"`
}

type DataReportTransaksiSendiri struct {
	IDTransaksi       int64         `db:"ID_TRANSAKSI"`
	NomorRekening     string        `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening      string        `excel:"C" db:"NAMA_REKENING"`
	KodeCabangRek     string        `excel:"G" db:"KODE_CABANG_REK"`
	KodeCabangTrans   string        `db:"KODE_CABANG_TRANSAKSI"`
	Keterangan        string        `db:"KETERANGAN"`
	NomorReferensi    string        `db:"NOMOR_REFERENSI"`
	IdTransaksi       int64         `db:"ID_TRANSAKSI"`
	UserOveride       string        `db:"USER_OVERIDE"`
	UserOtorisasi     string        `db:"USER_OTORISASI"`
	KodeTransaksi     string        `db:"KODE_TRANSAKSI"`
	JamInput          string        `db:"JAM_INPUT"`
	KodeValuta        string        `excel:"F" db:"KODE_VALUTA"`
	NominalDebet      float64       `excel:"D" db:"NOMINAL_DEBET" currency:"IDR"`
	NominalKredit     float64       `excel:"E" db:"NOMINAL_KREDIT" currency:"IDR"`
	JenisMutasi       string        `db:"JENIS_MUTASI"`
	IdDetilTransaksi  int64         `db:"ID_DETIL_TRANSAKSI"`
	KodeRC            string        `db:"KODE_RC"`
	AccountCode       string        `db:"ACCOUNT_CODE"`
	AccountName       string        `db:"ACCOUNT_NAME"`
	UserOverridePass  string        `db:"USER_OVERRIDE_PASSBOOK"`
	IsPassbookWajib   string        `db:"IS_PASSBOOK_WAJIB"`
	IsPassbookDicetak string        `db:"IS_PASSBOOK_DICETAK"`
	KdSumberDanaTrx   string        `db:"KD_SUMBER_DANA_TRX"`
	KdTujuanTrx       string        `db:"KD_TUJUAN_TRX"`
	KetSumberDanaTrx  string        `db:"KET_SUMBER_DANA_TRX"`
	KetTujuanTrx      string        `db:"KET_TUJUAN_TRX"`
	TipeAksesTrx      string        `db:"TIPE_AKSES_TRANSAKSI"`
	ListNamaAksesTrx  string        `db:"LISTNAMA_AKSES_TRANSAKSI"`
	Sequence          sql.NullInt64 `db:"SEQUENCE"`
	NamaNoRek         string        `excel:"H" db:"NAMA_NO_REK"`
}

type DataReportTransaksiJurnal struct {
	FGrouping         sql.NullString `db:"FGROUPING"`
	JournalNo         sql.NullString `db:"JOURNAL_NO"`
	Description       sql.NullString `db:"DESCRIPTION"`
	JournalType       sql.NullString `db:"JOURNAL_TYPE"`
	JournalItemNo     sql.NullString `db:"JOURNALITEM_NO"`
	JournalItemStatus sql.NullString `db:"JOURNALITEMSTATUS"`
	SubAccountCode    sql.NullString `db:"SUBACCOUNTCODE"`
	SubAccountName    sql.NullString `excel:"I" db:"SUBACCOUNTNAME"`
	AmountDebit       float64        `excel:"D" db:"AMOUNT_DEBIT" currency:"IDR"`
	AmountCredit      float64        `excel:"E" db:"AMOUNT_CREDIT" currency:"IDR"`
	ValutaCode        sql.NullString `excel:"F" db:"VALUTA_CODE"`
	NilaiKurs         float64        `db:"NILAI_KURS"`
	IDJournalBlock    sql.NullString `db:"ID_JOURNALBLOCK"`
	ReferenceNo       sql.NullString `excel:"B" db:"REFERENCE_NO"`
	Status            sql.NullString `db:"STATUS"`
	AccountCode       sql.NullString `excel:"H" db:"ACCOUNT_CODE"`
	AccountName       sql.NullString `excel:"C" db:"ACCOUNT_NAME"`
	AccountType       sql.NullString `db:"ACCOUNT_TYPE"`
	BranchCode        sql.NullString `excel:"G" db:"BRANCH_CODE"`
	JournalDate       sql.NullString `db:"JOURNAL_DATE"`
	BranchCodeAcc     sql.NullString `db:"BRANCH_CODE_ACCOUNT"`
	IDJournal         sql.NullString `db:"ID_JOURNALBLOCK"`
	DescriptionJI     sql.NullString `db:"JI_DESCRIPTION"`
	NomorReferensi    sql.NullString `db:"NOMOR_REFERENSI"`
	JamInput          sql.NullString `db:"JAM_INPUT"`
	UserOtorisasi     sql.NullString `db:"USER_OTORISASI"`
	UserOverride      sql.NullString `db:"USER_OVERIDE"`
	UserInput         sql.NullString `db:"USER_INPUT"`
}

type DataTrxStandingInstruction struct {
	TanggalProses          string `excel:"B" db:"PROCESSDATE"`
	TanggalRencanaProsesSI string `excel:"C" db:"PROCESSDATEPLAN"`
	Nominal                string `excel:"D" db:"AMOUNT"`
	RekeningDebet          string `excel:"E" db:"DEBITACCOUNTNO"`
	NamaRekeningDebet      string `excel:"F" db:"NAMA_REKENING_DEBET"`
	RekeningKredit         string `excel:"G" db:"CREDITACCOUNTNO"`
	NamaRekeningKredit     string `excel:"H" db:"NAMA_REKENING_KREDIT"`
	KeteranganTransaksi    string `excel:"I" db:"TRANSDESCRIPTION"`
	StatusProses           string `excel:"J" db:"STATUS"`
	KeteranganProses       string `excel:"K" db:"PROCESS_DESC"`
	IdTransaksi            string `db:"ID_TRANSAKSI"`
	InstructionId          string `db:"INSTRUCTIONID"`
	DataCount              string `db:"DATA_COUNT"`
}
type DataTplAMLTransKeuanganTunai struct {
	TanggalTransaksi  sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	NoCif             sql.NullString  `excel:"C" db:"NOMOR_NASABAH"`
	CabCif            sql.NullString  `excel:"D" db:"KODE_CABANG"`
	NamaNasabah       sql.NullString  `excel:"E" db:"NAMA_NASABAH"`
	NomorRekening     sql.NullString  `excel:"F" db:"NOMOR_REKENING"`
	JenisTransaksi    sql.NullString  `excel:"G" db:"JENIS_TRANSAKSI"`
	Nominal           sql.NullFloat64 `excel:"H" db:"NILAI_MUTASI" currency:"IDR"`
	CabTrans          sql.NullString  `excel:"I" db:"KODE_CABANG_TRANSAKSI"`
	NomorReferensi    sql.NullString  `excel:"J" db:"NOMOR_REFERENSI"`
	FisikTunai        sql.NullString  `excel:"K" db:"IS_TUNAI_FISIK"`
	NominalTunaiFisik sql.NullFloat64 `excel:"L" db:"NOMINAL_TUNAI_FISIK" currency:"IDR"`
	Berita            sql.NullString  `excel:"M" db:"BERITA"`
	DataCount         sql.NullInt64   `db:"DATA_COUNT"`
	KodeGrupCtr       sql.NullString  `db:"KODE_GRUP_CTR"`
}

type DataReportAMLTotalCIF struct {
	KodeJenis sql.NullString `db:"KODE_JENIS"`
	Jumlah    sql.NullString `db:"JUMLAH"`
}

type DataTrxCabRTGSTransferIn struct {
	TanggalTransaksi    string `excel:"B" db:"TANGGAL_TRANSAKSI"`
	BankPengirim        string `excel:"C" db:"BANK_PENGIRIM"`
	NilaiTransfer       string `excel:"D" db:"AMOUNT"`
	KeteranganTransfer  string `excel:"E" db:"PAYMENT_DETAILS"`
	NamaPengirim        string `excel:"F" db:"ORIGINATING_NAME"`
	NomorRekeningTujuan string `excel:"G" db:"ULTIMATE_BENEF_ACCOUNT"`
	NamaRekening        string `excel:"H" db:"NAMA_REKENING"`
	UltimateBenefName   string `db:"ULTIMATE_BENEF_NAME"`
	Cabang              string `db:"CABANG"`
	DataCount           string `db:"DATA_COUNT"`
	FromMemher          string `db:"FROM_MEMBER"`
	NamaMember          string `db:"NAMA_MEMBER"`
}

type DataReportAMLNasabahOverLimitTrans struct {
	Cabang                sql.NullString  `excel:"B" db:"KODE_CABANG"`
	NomorCif              sql.NullString  `excel:"C" db:"NOMOR_NASABAH"`
	NamaNasabah           sql.NullString  `excel:"D" db:"NAMA_NASABAH"`
	JNasabah              sql.NullString  `excel:"E" db:"J_NASABAH"`
	Limit                 sql.NullString  `excel:"F" db:"LIMIT"`
	TunaiNominalDebet     sql.NullFloat64 `excel:"G" db:"TUNAI_NOMINAL_DEBET"`
	TunaiFrekDebet        sql.NullInt32   `excel:"H" db:"TUNAI_FREK_DEBET"`
	TunaiNominalKredit    sql.NullFloat64 `excel:"I" db:"TUNAI_NOMINAL_KREDIT"`
	TunaiFrekKredit       sql.NullInt32   `excel:"J" db:"TUNAI_FREK_KREDIT"`
	NontunaiNominalDebet  sql.NullFloat64 `excel:"K" db:"NONTUNAI_NOMINAL_DEBET"`
	NontunaiFrekDebet     sql.NullInt32   `excel:"L" db:"NONTUNAI_FREK_DEBET"`
	NontunaiNominalKredit sql.NullFloat64 `excel:"M" db:"NONTUNAI_NOMINAL_KREDIT"`
	NontunaiFrekKredit    sql.NullInt32   `excel:"N" db:"NONTUNAI_FREK_KREDIT"`
	Bulan                 sql.NullString  `db:"BULAN"`
	Tahun                 sql.NullString  `db:"TAHUN"`
	Periode               sql.NullString  `db:"PERIODE"`

	TunaiNominalDebet1     sql.NullFloat64 `db:"TUNAI_NOMINAL_DEBET1"`
	TunaiFrekDebet1        sql.NullInt32   `db:"TUNAI_FREK_DEBET1"`
	TunaiNominalKredit1    sql.NullFloat64 `db:"TUNAI_NOMINAL_KREDIT1"`
	TunaiFrekKredit1       sql.NullInt32   `db:"TUNAI_FREK_KREDIT1"`
	NontunaiNominalDebet1  sql.NullFloat64 `db:"NONTUNAI_NOMINAL_DEBET1"`
	NontunaiFrekDebet1     sql.NullInt32   `db:"NONTUNAI_FREK_DEBET1"`
	NontunaiNominalKredit1 sql.NullFloat64 `db:"NONTUNAI_NOMINAL_KREDIT1"`
	NontunaiFrekKredit1    sql.NullInt32   `db:"NONTUNAI_FREK_KREDIT1"`
}

type DataTplAMLTransKeuanganTunaiSummary struct {
	TanggalTransaksi    sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	CabCif              sql.NullString  `excel:"C" db:"KODE_CABANG"`
	NoCif               sql.NullString  `excel:"D" db:"NOMOR_NASABAH"`
	NamaNasabah         sql.NullString  `excel:"E" db:"NAMA_NASABAH"`
	JenisTransaksi      sql.NullString  `excel:"F" db:"KODE_TRANSAKSI"`
	TotalNilaiTransaksi sql.NullFloat64 `excel:"G" db:"NILAI_MUTASI"`
}

type DataTrxCabRTGSTransferOut struct {
	TanggalTransaksi    sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	SandiBankKeterangan sql.NullString  `excel:"C" db:"KODE_BANK"`
	NamaBank            sql.NullString  `excel:"D" db:"NAMA_MEMBER"`
	NilaiTransfer       sql.NullFloat64 `excel:"E" sum:"T" currency:"IDR" db:"NILAI_MUTASI"`
	NoAkunPenerima      sql.NullString  `excel:"F" db:"NOMOR_REKENING_PENERIMA"`
	NamaPenerima        sql.NullString  `excel:"G" db:"NAMA_PENERIMA"`
	NamaPengirim        sql.NullString  `excel:"H" db:"NAMA_PENGIRIM"`
	NomorReferensi      sql.NullString  `excel:"I" db:"NOMOR_REFERENSI"`
	RekeningSumberDana  sql.NullString  `excel:"J" db:"NOMOR_REKENING_DEBET"`
	Biaya               sql.NullFloat64 `excel:"K" currency:"IDR" db:"NOMINAL_BIAYA"`
	UserInput           sql.NullString  `excel:"L" db:"USER_INPUT"`
	UserOveride         sql.NullString  `excel:"M" db:"USER_OVERIDE"`
	UserOtorisasi       sql.NullString  `excel:"N" db:"USER_OTORISASI"`
	StatusOtorisasi     sql.NullString  `excel:"O" db:"DESCRIPTION"`
	JamInput            sql.NullString  `excel:"P" db:"JAM_INPUT"`
	JamOtorisasi        sql.NullString  `excel:"Q" db:"JAM_OTORISASI"`
	UserCreateFile      sql.NullString  `excel:"R" db:"SESSION_USERID"`
	JamGenerateFile     sql.NullString  `excel:"S" db:"SESSION_TIME"`
	KodeCabangTransaksi sql.NullString  `db:"KODE_CABANG_TRANSAKSI"`
	Keterangan          sql.NullString  `db:"KETERANGAN"`
	IdTransaksi         sql.NullString  `db:"ID_TRANSAKSI"`
	IsReversed          sql.NullString  `db:"IS_REVERSED"`
	DataCount           sql.NullString  `db:"DATA_COUNT"`
	EnumDescription     sql.NullString  ` db:"ENUM_DESCRIPTION"`
	StatusOtorisasi2    sql.NullString  `db:"STATUS_OTORISASI"`
}

type DataReportSKNTransferKeluar struct {
	TanggalTransaksi    sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	BankTujuan          sql.NullString  `excel:"C" db:"KODE_BANK"`
	NilaiTransfer       sql.NullFloat64 `excel:"D" sum:"T" currency:"IDR" db:"NILAI_MUTASI"`
	KeteranganTransfer  sql.NullString  `excel:"E" db:"KETERANGAN"`
	NoRekPenerima       sql.NullString  `excel:"F" db:"NOMOR_REKENING_PENERIMA"`
	NamaPenerima        sql.NullString  `excel:"G" db:"NAMA_PENERIMA"`
	NamaPengirim        sql.NullString  `excel:"H" db:"NAMA_PENGIRIM"`
	NomorReferensi      sql.NullString  `excel:"I" db:"NOMOR_REFERENSI"`
	RekeningSumberDana  sql.NullString  `excel:"J" db:"NOMOR_REKENING_DEBET"`
	Biaya               sql.NullFloat64 `excel:"K" currency:"IDR" db:"NOMINAL_BIAYA"`
	UserInput           sql.NullString  `excel:"L" db:"USER_INPUT"`
	UserOveride         sql.NullString  `excel:"M" db:"USER_OVERIDE"`
	UserOtorisasi       sql.NullString  `excel:"N" db:"USER_OTORISASI"`
	StatusOtorisasi     sql.NullString  `excel:"O" db:"ENUM_DESCRIPTION"`
	JamInput            sql.NullString  `excel:"P" db:"JAM_INPUT"`
	JamOtorisasi        sql.NullString  `excel:"Q" db:"JAM_OTORISASI"`
	UserCreateFile      sql.NullString  `excel:"R" db:"SESSION_USER_ID"`
	JamGenerate         sql.NullString  `excel:"S" db:"SESSION_TIME"`
	ReffNo              sql.NullString  `db:"REFF_NO"`
	KodeCabangTransaksi sql.NullString  `db:"KODE_CABANG_TRANSAKSI"`
	IdTransaksi         sql.NullString  `db:"ID_TRANSAKSI"`
	IsReversed          sql.NullString  `db:"IS_REVERSED"`
	NamaBank            sql.NullString  `db:"NAMA_BANK"`
	DataCount           sql.NullString  `db:"DATA_COUNT"`
}

type DataReportSKNTransferMasuk struct {
	TglTransaksi       string         `excel:"B" db:"TANGGAL_TRANSAKSI"`
	BankPengirim       string         `excel:"C" db:"KODE_BANK"`
	NilaiTransfer      float64        `excel:"D" db:"NILAI_MUTASI"`
	KeteranganTransfer sql.NullString `excel:"E" db:"DESCR"`
	NoRekTujuan        string         `excel:"F" db:"RECP_ACCOUNT"`
	NamaRekTujuan      string         `excel:"G" db:"RECP_NAME"`
	NamaPengirim       string         `excel:"H" db:"SENDER_NAME"`
	BatchNo            sql.NullString `db:"BATCH_NO"`
	NamaBank           sql.NullString `db:"NAMA_BANK"`
	SessionUserId      sql.NullString `db:"SESSION_USER_ID"`
	NoReferensi        sql.NullString `db:"NOMOR_REFERENSI"`
	UserOtorisasi      sql.NullString `db:"USER_OTORISASI"`
	JamInput           sql.NullString `db:"JAM_INPUT"`
	JamGenerate        sql.NullString `db:"SESSION_TIME"`
	BranchCode         sql.NullString `db:"BRANCH_CODE"`
	DataCount          sql.NullString `db:"DATA_COUNT"`
}

type DataReportTplTransaksiCabSKNTarikanKliring struct {
	TglTarikan          sql.NullString  `excel:"B" db:"SESSION_DATE"`
	SandiBankPenarik    sql.NullString  `excel:"C" db:"SANDI_BANK"`
	NoWarkat            sql.NullString  `excel:"D" db:"SERIAL_NO"`
	JenisWarkat         sql.NullString  `excel:"E" db:"KET_WARKAT"`
	NilaiTarikanKliring sql.NullFloat64 `excel:"F" db:"TRX_AMOUNT"`
	NoRekeningDebet     sql.NullString  `excel:"G" db:"ACCOUNT_NO"`
	NamaRekeningDebet   sql.NullString  `excel:"H" db:"NAMA_REKENING"`
	KodeBank            sql.NullString  `db:"KODE_BANK"`
	BatchNo             sql.NullString  `db:"BATCH_NO"`
	TrxCode             sql.NullString  `db:"TRX_CODE"`
	NamaBank            sql.NullString  `db:"NAMA_BANK"`
	SessionTime         sql.NullString  `db:"SESSION_TIME"`
	DataCount           sql.NullInt64   `db:"DATA_COUNT"`
}

type DataReportTplTransaksiCabSKNTransferRetur struct {
	TglRetur      sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	BankPengirim  sql.NullString  `excel:"C" db:"BANK_PENGIRIM"`
	NilaiTransfer sql.NullFloat64 `excel:"D" db:"NILAI_MUTASI"`
	AlasanRetur   sql.NullString  `excel:"E" db:"DESCR"`
	NoReff        sql.NullString  `excel:"F" db:"RETUR_REFF_NO"`
	ReturAccount  sql.NullString  `excel:"G" db:"RETUR_ACCOUNT"`
	NamaRekening  sql.NullString  `excel:"H" db:"NAMA_REKENING"`
	NamaBank      sql.NullString  `db:"NAMA_BANK"`
	BatchNo       sql.NullString  `db:"BATCH_NO"`
	SessionUserId sql.NullString  `db:"SESSION_USER_ID"`
	KodeBank      sql.NullString  `db:"KODE_BANK"`
	BranchCode    sql.NullString  `db:"BRANCH_CODE"`
	DataCount     sql.NullInt64   `db:"DATA_COUNT"`
}

type DataReportTransaksiCabSKNDebet struct {
	TglTransaksi        sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	NoAplikasi          sql.NullString  `excel:"C" db:"NOMOR_REFERENSI"`
	NoReferensiKliring  sql.NullString  `excel:"D" db:"REFF_NO"`
	SandiBank           sql.NullString  `excel:"E" db:"KODE_BANK"`
	NilaiSetoranKliring sql.NullFloat64 `excel:"F" db:"NILAI_MUTASI"`
	JenisWarkat         sql.NullString  `excel:"G" db:"JENIS_WARKAT"`
	NomorWarkat         sql.NullString  `excel:"H" db:"WARKAT_NO"`
	NoRekPemilikWarkat  sql.NullString  `excel:"I" db:"NOMOR_REKENING_NOTA"`
	NamaPemilikWarkat   sql.NullString  `excel:"J" db:"NAMA_PEMILIK_NOTA"`
	UserInput           sql.NullString  `excel:"K" db:"USER_INPUT"`
	UserOverride        sql.NullString  `excel:"L" db:"USER_OVERIDE"`
	UserOtorisasi       sql.NullString  `excel:"M" db:"USER_OTORISASI"`
	StatusOtorisasi     sql.NullString  `excel:"N" db:"ENUM_DESCRIPTION"`
	UserCreateFile      sql.NullString  `excel:"O" db:"SESSION_USER_ID"`
	JamGenerateFile     sql.NullString  `excel:"P" db:"SESSION_TIME"`
	Grouping            sql.NullString  `db:"GROUPING"`
	NamaBank            sql.NullString  `db:"NAMA_BANK"`
	JamInput            sql.NullString  `db:"JAM_INPUT"`
	Cabang              sql.NullString  `db:"CABANG"`
	IsReversed          sql.NullString  `db:"IS_REVERSED"`
	DataCount           sql.NullString  `db:"DATA_COUNT"`
}

type DataReportTransaksiCabSKNReturDebet struct {
	TglRetur            string         `excel:"B" db:"SESSION_DATE"`
	SandiBank           string         `excel:"C" db:"KODE_BANK"`
	NoWarkat            string         `excel:"D" db:"WARKAT_NO"`
	JenisWarkat         string         `excel:"E" db:"KET_WARKAT"`
	NilaiSetoranKliring float64        `excel:"F" db:"TRX_AMOUNT"`
	AlasanTolakan       string         `excel:"G" db:"KETERANGAN"`
	NoRekeningNota      string         `excel:"H" db:"ACCOUNT_NO"`
	NamaPemilikNota     string         `excel:"I" db:"NAMA_PEMILIK_NOTA"`
	NomorRekeningKredit string         `excel:"J" db:"NOMOR_REKENING"`
	NamaRekeningKredit  string         `excel:"K" db:"NAMA_REKENING"`
	BatchNo             sql.NullString `db:"BATCH_NO"`
	TrxCode             sql.NullString `db:"TRX_CODE"`
	Reason              sql.NullString `db:"REASON"`
	SessionTime         sql.NullString `db:"SESSION_TIME"`
	DataCount           sql.NullString `db:"DATA_COUNT"`
}

type DataReportTplTransaksiCabSKNTransferInFailed struct {
	TanggalTransaksi   sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	BankPengirim       sql.NullString  `excel:"C" db:"BANK_PENGIRIM"`
	NilaiTransfer      sql.NullFloat64 `excel:"D" db:"NILAI_MUTASI"`
	KeteranganTransfer sql.NullString  `excel:"E" db:"DESCR"`
	NoRekTujuan        sql.NullString  `excel:"F" db:"RECP_ACCOUNT"`
	NamaRekTujuan      sql.NullString  `excel:"G" db:"RECP_NAME"`
	NamaPengirim       sql.NullString  `excel:"H" db:"SENDER_NAME"`
	KeteranganMaintSys sql.NullString  `excel:"I" db:"FAILED_DESCRIPTION"`
	NamaBank           sql.NullString  `db:"NAMA_BANK"`
	BatchNo            sql.NullString  `db:"BATCH_NO"`
	SessionUserId      sql.NullString  `db:"SESSION_USER_ID"`
	KodeBank           sql.NullString  `db:"KODE_BANK"`
	NomorReferensi     sql.NullString  `db:"NOMOR_REFERENSI"`
	UserOtorisasi      sql.NullString  `db:"USER_OTORISASI"`
	JamInput           sql.NullString  `db:"JAM_INPUT"`
	SessionTime        sql.NullString  `db:"SESSION_TIME"`
	BranchCode         sql.NullString  `db:"BRANCH_CODE"`
	DataCount          sql.NullInt64   `db:"DATA_COUNT"`
}

type DataReportTplTransaksiCabRTGSTransferRetur struct {
	TanggalRetur        sql.NullString `excel:"B" db:"TANGGAL_TRANSAKSI"`
	NamaBank            sql.NullString `excel:"C" db:"NAMA_BANK"`
	NilaiTransfer       sql.NullString `excel:"D" db:"AMOUNT"`
	AlasanRetur         sql.NullString `excel:"E" db:"PAYMENT_DETAILS"`
	NamaPengirim        sql.NullString `excel:"F" db:"ORIGINATING_NAME"`
	NomorRekeningTujuan sql.NullString `excel:"G" db:"ULTIMATE_BENEF_ACCOUNT"`
	NamaRekeningTujuan  sql.NullString `excel:"H" db:"ULTIMATE_BENEF_NAME"`
	FailedDescription   sql.NullString `excel:"I" db:"FAILED_DESCRIPTION"`
	Namper              sql.NullString `db:"NAMA_PENGIRIM"`
	FromMember          sql.NullString `db:"FROM_MEMBER"`
	NamaMember          sql.NullString `db:"NAMA_MEMBER"`
	Cabang              sql.NullString `db:"CABANG"`
	DataCount           sql.NullString `db:"DATA_COUNT"`
}

type DataReportTplTransaksiCabRTGSTransferInFailed struct {
	TanggalTransaksi   sql.NullString  `excel:"B" db:"TANGGAL_TRANSAKSI"`
	BankPengirim       sql.NullString  `excel:"C" db:"BANK_PENGIRIM"`
	NilaiTransfer      sql.NullFloat64 `excel:"D" db:"AMOUNT"`
	KeteranganTransfer sql.NullString  `excel:"E" db:"PAYMENT_DETAILS"`
	NamaPengirim       sql.NullString  `excel:"F" db:"ORIGINATING_NAME"`
	NoRekTujuan        sql.NullString  `excel:"G" db:"ULTIMATE_BENEF_ACCOUNT"`
	NamaRekTujuan      sql.NullString  `excel:"H" db:"ULTIMATE_BENEF_NAME"`
	KeteranganSistem   sql.NullString  `excel:"I" db:"FAILED_DESCRIPTION"`
	Cabang             sql.NullString  `db:"CABANG"`
	FromMember         sql.NullString  `db:"FROM_MEMBER"`
	NamaMember         sql.NullString  `db:"NAMA_MEMBER"`
	DataCount          sql.NullInt64   `db:"DATA_COUNT"`
}

type DataReportTransaksiJurnalSendiri struct {
	KodeRekening            sql.NullString  `excel:"B" db:"ACCOUNT_CODE"`
	NamaRekening            sql.NullString  `excel:"C" db:"ACCOUNT_NAME"`
	NominalDebet            sql.NullFloat64 `excel:"D" db:"AMOUNT_DEBIT"`
	NominalKredit           sql.NullFloat64 `excel:"E" db:"AMOUNT_CREDIT"`
	ValutaCode              sql.NullString  `excel:"F" db:"VALUTA_CODE"`
	CabangRekening          sql.NullString  `excel:"G" db:"BRANCH_CODE_ACCOUNT"`
	UserOtorisasi           sql.NullString  `excel:"H" db:"SUBACCOUNTCODE"`
	WaktuInput              sql.NullString  `excel:"I" db:"SUBACCOUNTNAME"`
	FGrouping               sql.NullString  `db:"FGROUPING"`
	JournalNo               sql.NullString  `db:"JOURNAL_NO"`
	JournalType             sql.NullString  `db:"JOURNAL_TYPE"`
	JournalitemNo           sql.NullString  `db:"JOURNALITEM_NO"`
	Status                  sql.NullString  `db:"STATUS"`
	AccountType             sql.NullString  `db:"ACCOUNT_TYPE"`
	BranchCode              sql.NullString  `db:"BRANCH_CODE"`
	JournalDate             sql.NullString  `db:"JOURNAL_DATE"`
	IsConfidential          sql.NullString  `db:"IS_CONFIDENTIAL"`
	ConfidentialAllowed     sql.NullString  `db:"CONFIDENTIALALLOWED"`
	NilaiKurs               sql.NullString  `db:"NILAI_KURS"`
	IdJournalblock          sql.NullString  `db:"ID_JOURNALBLOCK"`
	FlProjectNo             sql.NullString  `db:"FL_PROJECT_NO"`
	JiDescription           sql.NullString  `db:"JI_DESCRIPTION"`
	NomorReferensiTrans     sql.NullString  `db:"NOMOR_REFERENSI_TRANS"`
	NomorReferensiHisttrans sql.NullString  `db:"NOMOR_REFERENSI_HISTTRANS"`
	JamInputHisttrans       sql.NullString  `db:"JAM_INPUT_HISTTRANS"`
	UserOtorHisttrans       sql.NullString  `db:"USER_OTOR_HISTTRANS"`
	UserOvrHisttrans        sql.NullString  `db:"USER_OVR_HISTTRANS"`
	UserInputHisttrans      sql.NullString  `db:"USER_INPUT_HISTTRANS"`

	ReferenceNo    sql.NullString `db:"REFERENCE_NO"`
	Description    sql.NullString `db:"DESCRIPTION"`
	UserInputTrans sql.NullString `db:"USER_INPUT_TRANS"`
	UserOvrTrans   sql.NullString `db:"USER_OVR_TRANS"`
	UserOtorTrans  sql.NullString `db:"USER_OTOR_TRANS"`
	JamInputTrans  sql.NullString `db:"JAM_INPUT_TRANS"`
}

type DataReportRekapTransaksiSendiri struct {
	KodeValuta          sql.NullString  `excel:"B" db:"KODE_VALUTA"`
	KodeCabangTransaksi sql.NullString  `excel:"C" db:"KODE_CABANG_TRANSAKSI"`
	KodeTransaksi       sql.NullString  `excel:"D" db:"KODE_TRANSAKSI"`
	Jumlah              sql.NullInt64   `excel:"E" db:"JUMLAH"`
	NonCashDebet        sql.NullFloat64 `excel:"F" currency:"IDR" db:"NON_CASH_DEBET"`
	NonCashKredit       sql.NullFloat64 `excel:"G" currency:"IDR" db:"NON_CASH_KREDIT"`
	BeginningBalance    string          `excel:"H"`
	CashDebet           sql.NullFloat64 `excel:"I" currency:"IDR" db:"CASH_DEBET"`
	CashKredit          sql.NullFloat64 `excel:"J" currency:"IDR" db:"CASH_KREDIT"`
	EndingBalance       sql.NullFloat64 `excel:"K" currency:"IDR" db:"ENDING_BALANCE"`
}

type DataReportTransaksiReversalOtorisasi struct {
	NomorRekening            sql.NullString  `excel:"B" db:"NOMOR_REKENING"`
	NamaRekening             sql.NullString  `excel:"C" db:"NAMA_REKENING"`
	KodeCabangRek            sql.NullString  `excel:"G" db:"KODE_CABANG_REK"`
	Keterangan               sql.NullString  `db:"KETERANGAN"`
	KeteranganDetail         sql.NullString  `excel:"H" db:"KETERANGAN_DETAIL"`
	NomorReferensi           sql.NullString  `db:"NOMOR_REFERENSI"`
	IdTransaksi              *int64          `db:"ID_TRANSAKSI"`
	KodeTransaksi            sql.NullString  `db:"KODE_TRANSAKSI"`
	UserOveride              sql.NullString  `db:"USER_OVERIDE"`
	UserOtorisasi            sql.NullString  `db:"USER_OTORISASI"`
	TanggalTransaksi         sql.NullString  `db:"TANGGAL_TRANSAKSI"`
	JamInput                 sql.NullString  `db:"JAM_INPUT"`
	TanggalTransaksiJamInput sql.NullString  `db:"TANGGAL_TRANSAKSI_JAM_INPUT"`
	TanggalReverse           sql.NullString  `db:"TANGGAL_REVERSE"`
	JamReverse               sql.NullString  `db:"JAM_REVERSE"`
	TanggalReverseJamReverse sql.NullString  `db:"TANGGAL_REVERSE_JAM_REVERSE"`
	UserReverse              sql.NullString  `db:"USER_REVERSE"`
	UserOverideReverse       sql.NullString  `db:"USER_OVERIDE_REVERSE"`
	KodeValuta               sql.NullString  `excel:"F" db:"KODE_VALUTA"`
	NominalDebet             sql.NullFloat64 `excel:"D" currency:"IDR" db:"NOMINAL_DEBET"`
	NominalKredit            sql.NullFloat64 `excel:"E" currency:"IDR" db:"NOMINAL_KREDIT"`
	DataCount                sql.NullString  `db:"DATA_COUNT"`
}
type GetReportDetail struct {
	KodeReport   sql.NullString `db:"KODE_REPORT"`
	NamaReport   sql.NullString `db:"NAMA_REPORT"`
	TemplateName sql.NullString `db:"TEMPLATE_NAME"`
	ScriptName   sql.NullString `db:"SCRIPT_NAME"`
	TagReport    sql.NullString `db:"TAG_REPORT"`
	IsEodExecute sql.NullString `db:"IS_EOD_EXECUTE"`
	Recipient    sql.NullString `db:"RECIPIENT"`
	Retensi      sql.NullString `db:"RETENSI"`
	IsShowBds    sql.NullString `db:"IS_SHOW_BDS"`
}
type GetListTrxRes struct {
	IdTransaksi         *int32   `db:"ID_TRANSAKSI" json:"id_transaksi,omitempty"`
	TanggalTransaksi    string   `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi,omitempty"`
	NomorReferensi      string   `db:"NOMOR_REFERENSI" json:"nomor_referensi,omitempty"`
	NilaiTransaksi      *float64 `db:"NILAI_TRANSAKSI" json:"nilai_transaksi,omitempty"`
	Keterangan          string   `db:"KETERANGAN" json:"keterangan,omitempty"`
	KodeCabangTransaksi string   `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi,omitempty"`
}
type ListHistTrx struct {
	IdDetilTransaksi     sql.NullInt64   `db:"ID_DETIL_TRANSAKSI" json:"id_detil_transaksi"`
	TanggalTransaksi     sql.NullString  `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	IdParameterTransaksi sql.NullString  `db:"ID_PARAMETER_TRANSAKSI" json:"id_parameter_transaksi"`
	KodeJurnal           sql.NullString  `db:"KODE_JURNAL" json:"kode_jurnal"`
	KodeTransaksi        sql.NullString  `db:"KODE_TRANSAKSI" json:"kode_transaksi"`
	Keterangan           sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	JenisMutasi          sql.NullString  `db:"JENIS_MUTASI" json:"jenis_mutasi"`
	SaldoAwal            sql.NullFloat64 `db:"SALDO_AWAL" json:"saldo_awal"`
	Saldo                sql.NullFloat64 `db:"SALDO" json:"saldo"`
	NilaiMutasi          sql.NullFloat64 `db:"NILAI_MUTASI" json:"nilai_mutasi"`
	NomorReferensi       sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	StatusOtorisasi      sql.NullInt32   `db:"STATUS_OTORISASI" json:"status_otorisasi"`
	StatusOtorisasiDesc  sql.NullString  `db:"STATUS_OTORISASI_DESC" json:"status_otorisasi_desc"`
	KodeCabangTransaksi  sql.NullString  `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi"`
	UserOtorisasi        sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	UserInput            sql.NullString  `db:"USER_INPUT" json:"user_input"`
	TanggalInput         sql.NullString  `db:"TANGGAL_INPUT" json:"tanggal_input"`
	JamInput             sql.NullString  `db:"JAM_INPUT" json:"jam_input"`
	KeteranganTambahan   sql.NullString  `db:"KETERANGAN_TAMBAHAN" json:"keterangan_tambahan"`
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
type HeaderAcctStatement struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING" `
	NamaRekening  sql.NullString `db:"NAMA_REKENING" `
	AlamatKirim   sql.NullString `db:"ALAMAT_KIRIM" `
	KodeCabang    sql.NullString `db:"KODE_CABANG" `
	NamaCabang    sql.NullString `db:"NAMA_CABANG" `
	JenisNasabah  sql.NullString `db:"JENIS_NASABAH" `
	NomorNasabah  sql.NullString `db:"NOMOR_NASABAH"`
}
type GetSumHeaderAcctStatement struct {
	Debit     sql.NullFloat64 `db:"DEBET" `
	Kredit    sql.NullFloat64 `db:"KREDIT" `
	JmlDebet  sql.NullInt32   `db:"JML_DEBET" `
	JmlKredit sql.NullInt32   `db:"JML_KREDIT" `
	Admin     sql.NullFloat64 `db:"ADMIN" `
	BagiHasil sql.NullFloat64 `db:"BAGI_HASIL" `
	Beban     sql.NullFloat64 `db:"BEBAN"`
}
