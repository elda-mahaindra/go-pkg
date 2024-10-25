package data

import (
	"database/sql"
	"time"
)

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

type ResGetDepList struct {
	NomorNasabah       sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NomorRekening      sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening       sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeCabang         sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	KodeProduk         sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk         sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
	SaldoEfektif       sql.NullFloat64 `db:"SALDO_EFEKTIF" json:"saldo_efektif"`
	StatusRekening     sql.NullString  `db:"STATUS_REKENING" json:"status_rekening"`
	StatusRekeningDesc sql.NullString  `db:"STATUS_REKENING_DESC" json:"status_rekening_desc"`
	JenisBlokir        sql.NullString  `db:"JENIS_BLOKIR" json:"jenis_blokir"`
	IsBlokirDebet      sql.NullString  `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	IsBlokirKredit     sql.NullString  `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsSedangTutup      sql.NullString  `db:"IS_SEDANG_DITUTUP" json:"is_sedang_ditutup"`
}

type DetailDeposito struct {
	NomorRekening                sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	IsBlokir                     sql.NullString  `db:"IS_BLOKIR" json:"is_blokir"`
	TanggalBlokir                sql.NullString  `db:"TANGGAL_BLOKIR" json:"tanggal_blokir"`
	KeteranganBlokir             sql.NullString  `db:"KETERANGAN_BLOKIR" json:"keterangan_blokir"`
	SistemPemblokir              sql.NullString  `db:"SISTEM_PEMBLOKIR" json:"sistem_pemblokir"`
	UserBlokir                   sql.NullString  `db:"USER_BLOKIR" json:"user_blokir"`
	Alamat                       sql.NullString  `db:"ALAMAT" json:"alamat"`
	AlamatKirim                  sql.NullString  `db:"ALAMAT_KIRIM" json:"alamat_kirim"`
	AlamatBilyet                 sql.NullString  `db:"ALAMAT_BILYET" json:"alamat_bilyet"`
	IsJoinAccount                sql.NullString  `db:"IS_JOIN_ACCOUNT" json:"is_join_account"`
	AlamatTujuanTransfer         sql.NullString  `db:"BAGHAS_ALAMAT_TRANSFER" json:"alamat_tujuan_transfer"`
	DisposisiBagiHasil           sql.NullString  `db:"DISPOSISI_BAGI_HASIL" json:"disposisi_bagi_hasil"`
	DisposisiNominal             sql.NullString  `db:"DISPOSISI_NOMINAL" json:"disposisi_nominal"`
	IbuKandung                   sql.NullString  `db:"NASABAH_IBU_KANDUNG" json:"ibu_kandung"`
	IsBagiHasil                  sql.NullString  `db:"IS_BAGI_HASIL_KHUSUS" json:"is_bagi_hasil"`
	IsDapatBagiHasil             sql.NullString  `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	IsKenaBiayaLayananUmum       sql.NullString  `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum"`
	IsStatusPassbook             sql.NullString  `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook"`
	IsTieringNisbahbonus         sql.NullString  `db:"IS_TIERING_NISBAHBONUS" json:"is_tiering_nisbahbonus"`
	JenisDeposito                sql.NullString  `db:"JENIS_DEPOSITO" json:"jenis_deposito"`
	JenisRekeningLiabilitas      sql.NullString  `db:"JENIS_REKENING_LIABILITAS" json:"jenis_rekening_liabilitas"`
	KodeSandiKota                sql.NullString  `db:"BAGHAS_SANDIKOTA_TRANSFER" json:"kode_sandi_kota"`
	Keterangan                   sql.NullString  `db:"KETERANGAN" json:"keterangan" `
	StatusRekening               sql.NullInt64   `db:"STATUS_REKENING" json:"status_rekening"`
	KodeBankTujuanTransfer       sql.NullString  `db:"KODE_MEMBER_NOMINAL" json:"kode_bank_tujuan_transfer"`
	KodeBankTujuanTransferBh     sql.NullString  `db:"KODE_MEMBER_BAGIHASIL" json:"kode_bank_tujuan_transfer_bh"`
	KodeCabang                   sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	NamaCabang                   sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
	KodeMarketingCurrent         sql.NullString  `db:"KODE_MARKETING_CURRENT" json:"kode_marketing_current"`
	KodeMarketingPertama         sql.NullString  `db:"KODE_MARKETING_PERTAMA" json:"kode_marketing_pertama"`
	KodeMarketingReferensi       sql.NullString  `db:"KODE_MARKETING_REFERENSI" json:"kode_marketing_referensi"`
	KodeProduk                   sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	KodeProgram                  sql.NullString  `db:"KODE_PROGRAM" json:"kode_program"`
	KodeSumberDana               sql.NullString  `db:"KODE_SUMBER_DANA" json:"kode_sumber_dana"`
	KodeTujuanRekening           sql.NullString  `db:"KODE_TUJUAN_REKENING" json:"kode_tujuan_rekening"`
	KodeValuta                   sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	NamabankTujuanTransfer       sql.NullString  `db:"KODE_MEMBER_NOMINAL_DESC" json:"namabank_tujuan_transfer"`
	NamabankTujuanTransferBh     sql.NullString  `db:"KODE_MEMBER_BAGIHASIL_DESC" json:"namabank_tujuan_transfer_bh"`
	NamaNasabah                  sql.NullString  `db:"NAMA_NASABAH" json:"nama_nasabah"`
	NamaProduk                   sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
	NamaRekening                 sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	NamaSandiKota                sql.NullString  `db:"BAGHAS_SANDIKOTA_TRANSFER_DESC" json:"nama_sandi_kota"`
	NamaValuta                   sql.NullString  `db:"NAMA_VALUTA" json:"nama_valuta"`
	NisbahBagiHasil              sql.NullFloat64 `db:"NISBAH_BAGI_HASIL" json:"nisbah_bagi_hasil"`
	NisbahDasar                  sql.NullFloat64 `db:"NISBAH_DASAR" json:"nisbah_dasar"`
	NisbahSpesial                sql.NullFloat64 `db:"NISBAH_SPESIAL" json:"nisbah_spesial"`
	NomorBilyet                  sql.NullString  `db:"NOMOR_BILYET" json:"nomor_bilyet"`
	NominalDeposito              sql.NullFloat64 `db:"SALDO_DEPOSITO" json:"nominal_deposito"`
	Saldo                        sql.NullFloat64 `db:"SALDO" json:"saldo"`
	NomorNasabah                 sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NomorRekeningDebet           sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening_debet"`
	NomorRekeningDeposito        sql.NullString  `db:"NOMORREKENINGDEP" json:"nomor_rekening_deposito"`
	NomorRekeningDisposisi       sql.NullString  `db:"NOMOR_REKENING_DISPOSISI" json:"nomor_rekening_disposisi"`
	RekeningDisposisi            sql.NullString  `db:"REKENING_DISPOSISI" json:"rekening_disposisi"`
	PenerimaTransferBagiHasil    sql.NullString  `db:"PENERIMA_TRANSFER_BAGI_HASIL" json:"penerima_transfer_bagi_hasil"`
	PenerimaTransferDisposisi    sql.NullString  `db:"PENERIMA_TRANSFER_DISPOSISI" json:"penerima_transfer_disposisi"`
	PersentaseZakatBagiHasil     sql.NullFloat64 `db:"PERSENTASE_ZAKAT_BAGI_HASIL" json:"persentase_zakat_bagi_hasil"`
	StatusKelengkapan            sql.NullString  `db:"STATUS_KELENGKAPAN" json:"status_kelengkapan"`
	StatusRestriksi              sql.NullString  `db:"STATUS_RESTRIKSI" json:"status_restriksi"`
	KodeStatusRetriksi           sql.NullString  `db:"KODE_STATUS_RETRIKSI" json:"kode_status_retriksi"`
	TanggalBuka                  sql.NullString  `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	TanggalJatuhTempoBerikutnya  sql.NullString  `db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA" json:"tanggal_jatuh_tempo_berikutnya"`
	TanggalJatuhTempoTerakhir    sql.NullString  `db:"TANGGAL_JATUH_TEMPO_TERAKHIR" json:"tanggal_jatuh_tempo_terakhir"`
	TanggalLahir                 sql.NullString  `db:"NASABAH_TANGGAL_LAHIR" json:"tanggal_lahir"`
	TarifPajak                   sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak"`
	TelpHP                       sql.NullString  `db:"TELP_HP" json:"telp_hp"`
	TelpRumah                    sql.NullString  `db:"TELP_RUMAH" json:"telp_rumah"`
	UserInput                    sql.NullString  `db:"USER_INPUT" json:"user_input"`
	UserOtorisasi                sql.NullString  `db:"USER_OTORISASI" json:"user_otorisasi"`
	IsSudahDisetor               sql.NullString  `db:"IS_SUDAH_DISETOR" json:"is_sudah_disetor"`
	IsSedangDicairkan            sql.NullString  `db:"IS_SEDANG_DICAIRKAN" json:"is_sedang_dicairkan"`
	SaldoAccrualBagihasil        sql.NullFloat64 `db:"SALDO_ACCRUAL_BAGIHASIL" json:"saldo_accrual_bagihasil"`
	StatusRekeningDesc           sql.NullString  `db:"STATUS_REKENING_DESC" json:"status_rekening_desc"`
	SaldoDitahan                 sql.NullFloat64 `db:"SALDO_DITAHAN" json:"saldo_ditahan"`
	CadanganBagihasilKapitalisir sql.NullFloat64 `db:"CADANGAN_BAGIHASIL_KAPITALISIR" json:"cadangan_bagihasil_kapitalisir"`
	KodeMulticifrelation         sql.NullString  `db:"KODE_MULTICIFRELATION" json:"kode_multicifrelation"`
	PeriodePerjanjian            sql.NullString  `db:"PERIODE_PERJANJIAN" json:"periode_perjanjian"`
	MasaPerjanjian               sql.NullInt64   `db:"MASA_PERJANJIAN" json:"masa_perjanjian"`
	TanggalBagiHasilTerakhir     sql.NullString  `db:"TANGGAL_BAGI_HASIL_TERAKHIR" json:"tanggal_bagi_hasil_terakhir"`
	TanggalBagiHasilBerikutnya   sql.NullString  `db:"TANGGAL_BAGI_HASIL_BERIKUTNYA" json:"tanggal_bagi_hasil_berikutnya"`
}

type DataDepositoView struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING"`
	NomorBilyet   sql.NullString `db:"NOMOR_BILYET"`
}

type AccountDepBilyetPrint struct {
	NomorRekening          sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NamaRekening           sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeCabang             sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	NamaCabang             sql.NullString  `db:"NAMA_CABANG" json:"nama_cabang"`
	AlamatKirim            sql.NullString  `db:"ALAMAT_KIRIM" json:"alamat_kirim"`
	JenisNasabah           sql.NullString  `db:"JENIS_NASABAH" json:"jenis_nasabah"`
	NomorNasabah           sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	IsSudahCetakBilyet     sql.NullString  `db:"IS_SUDAH_CETAK_BILYET" json:"is_sudah_cetak_bilyet"`
	IsSudahDisetor         sql.NullString  `db:"IS_SUDAH_DISETOR" json:"is_sudah_disetor"`
	TanggalBuka            sql.NullString  `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	TanggalJatuhTempo      sql.NullString  `db:"TANGGAL_JATUH_TEMPO" json:"tanggal_jatuh_tempo"`
	BagiHasilDeposan       sql.NullFloat64 `db:"BAGI_HASIL_DEPOSAN" json:"bagi_hasil_deposan"`
	BagiHasilBank          sql.NullFloat64 `db:"BAGI_HASIL_BANK" json:"bagi_hasil_bank"`
	DisposisiNominal       sql.NullString  `db:"DISPOSISI_NOMINAL" json:"disposisi_nominal"`
	DisposisiBagiHasil     sql.NullString  `db:"DISPOSISI_BAGI_HASIL" json:"disposisi_bagi_hasil"`
	IsAro                  sql.NullString  `db:"IS_ARO" json:"is_aro"`
	NomorBilyet            sql.NullString  `db:"NOMOR_BILYET" json:"nomor_bilyet"`
	NomorRekeningPokok     sql.NullString  `db:"NOMOR_REKENING_POKOK" json:"nomor_rekening_pokok"`
	NamaRekeningPokok      sql.NullString  `db:"NAMA_REKENING_POKOK" json:"nama_rekening_pokok"`
	KodeBankPokok          sql.NullString  `db:"KODE_BANK_POKOK" json:"kode_bank_pokok"`
	NomorRekeningBagihasil sql.NullString  `db:"NOMOR_REKENING_BAGIHASIL" json:"nomor_rekening_bagihasil"`
	NamaRekeningBagihasil  sql.NullString  `db:"NAMA_REKENING_BAGIHASIL" json:"nama_rekening_bagihasil"`
	MasaPerjanjian         sql.NullInt32   `db:"MASA_PERJANJIAN" json:"masa_perjanjian"`
	PeriodePerjanjian      sql.NullString  `db:"PERIODE_PERJANJIAN" json:"periode_perjanjian"`
	Saldo                  sql.NullFloat64 `db:"SALDO" json:"saldo"`
}

type PostNotaDebetOutwardDeleted struct {
	JumlahTerhapus int32 `db:"JUMLAH_TERHAPUS" json:"jumlah_terhapus"`
}

type DepForCloseRespData struct {
	NomorRekening                sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NomorNasabah                 sql.NullString  `db:"NOMOR_NASABAH" json:"nomor_nasabah"`
	NamaRekening                 sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	KodeCabang                   sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	KodeValuta                   sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	IsKenaPenalty                sql.NullString  `db:"IS_KENA_PENALTY" json:"is_kena_penalty"`
	NominalPenalty               sql.NullFloat64 `db:"NOMINAL_PENALTY" json:"nominal_penalty"`
	TanggalBuka                  sql.NullString  `db:"TANGGAL_BUKA" json:"tanggal_buka"`
	TanggalJatuhTempoTerakhir    sql.NullString  `db:"TANGGAL_JATUH_TEMPO_TERAKHIR" json:"tanggal_jatuh_tempo_terakhir"`
	JumlahAro                    sql.NullInt32   `db:"JUMLAH_ARO" json:"jumlah_aro"`
	TanggalJatuhTempoBerikutnya  sql.NullString  `db:"TANGGAL_JATUH_TEMPO_BERIKUTNYA" json:"tanggal_jatuh_tempo_berikutnya"`
	Saldo                        sql.NullFloat64 `db:"SALDO" json:"saldo"`
	DisposisiNominal             sql.NullString  `db:"DISPOSISI_NOMINAL" json:"disposisi_nominal"`
	RekeningDisposisi            sql.NullString  `db:"REKENING_DISPOSISI" json:"rekening_disposisi"`
	PenerimaTransferDisposisi    sql.NullString  `db:"PENERIMA_TRANSFER_DISPOSISI" json:"penerima_transfer_disposisi"`
	StatusRekening               sql.NullInt32   `db:"STATUS_REKENING" json:"status_rekening"`
	IsSudahDisetor               sql.NullString  `db:"IS_SUDAH_DISETOR" json:"is_sudah_disetor"`
	IsSedangDicairkan            sql.NullString  `db:"IS_SEDANG_DICAIRKAN" json:"is_sedang_dicairkan"`
	CadanganBagihasilKapitalisir sql.NullFloat64 `db:"CADANGAN_BAGIHASIL_KAPITALISIR" json:"cadangan_bagihasil_kapitalisir"`
	IsKenaPajak                  sql.NullString  `db:"IS_KENA_PAJAK" json:"is_kena_pajak"`
	TarifPajak                   sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak"`
	NominalPajak                 sql.NullFloat64 `db:"NOMINAL_PAJAK" json:"nominal_pajak"`
	IsKenaZakatBagiHasil         sql.NullString  `db:"IS_KENA_ZAKAT_BAGI_HASIL" json:"is_kena_zakat_bagi_hasil"`
	PersentaseZakatBagiHasil     sql.NullFloat64 `db:"PERSENTASE_ZAKAT_BAGI_HASIL" json:"persentase_zakat_bagi_hasil"`
	NominalZakat                 sql.NullFloat64 `db:"NOMINAL_ZAKAT" json:"nominal_zakat"`
	NilaiDibayarkan              sql.NullFloat64 `db:"NILAI_DIBAYARKAN" json:"nilai_dibayarkan"`
	SaldoAccrualBagihasil        sql.NullFloat64 `db:"SALDO_ACCRUAL_BAGIHASIL" json:"saldo_accrual_bagihasil"`
}
type DepWithdraw struct {
	IdTransaksi                  sql.NullInt64   `db:"ID_TRANSAKSI" json:"id_transaksi"`
	NomorReferensi               sql.NullString  `db:"NOMOR_REFERENSI" json:"nomor_referensi"`
	Keterangan                   sql.NullString  `db:"KETERANGAN" json:"keterangan"`
	JenisAplikasi                sql.NullString  `db:"JENIS_APLIKASI" json:"jenis_aplikasi"`
	KodeValutaTransaksi          sql.NullString  `db:"KODE_VALUTA_TRANSAKSI" json:"kode_valuta_transaksi"`
	KursManual                   sql.NullFloat64 `db:"KURS_MANUAL" json:"kurs_manual"`
	NilaiTransaksi               sql.NullFloat64 `db:"NILAI_TRANSAKSI" json:"nilai_transaksi"`
	Biaya                        sql.NullFloat64 `db:"BIAYA" json:"biaya"`
	IdBatchTransaksi             sql.NullInt64   `db:"ID_BATCH_TRANSAKSI" json:"id_batch_transaksi"`
	TanggalTransaksi             sql.NullString  `db:"TANGGAL_TRANSAKSI" json:"tanggal_transaksi"`
	NomorRekening                sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NilaiPencairan               sql.NullFloat64 `db:"NILAI_PENCAIRAN" json:"nilai_pencairan"`
	KodeValuta                   sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	KodeCabangTransaksi          sql.NullString  `db:"KODE_CABANG_TRANSAKSI" json:"kode_cabang_transaksi"`
	StatusOtorisasi              sql.NullInt64   `db:"STATUS_OTORISASI" json:"status_otorisasi"`
	StatusOtorisasiDesc          sql.NullString  `db:"STATUS_OTORISASI_DESC" json:"status_otorisasi_desc"`
	NamaRekening                 sql.NullString  `db:"NAMA_REKENING" json:"nama_rekening"`
	IsTrxToday                   sql.NullString  `db:"IS_TRX_TODAY" json:"is_trx_today"`
	IsReversed                   sql.NullString  `db:"IS_REVERSED" json:"is_reversed"`
	DisposisiNominalOri          sql.NullString  `db:"DISPOSISI_NOMINAL_ORI" json:"disposisi_nominal_ori"`
	KodeBankDisposisiOri         sql.NullString  `db:"KODE_BANK_DISPOSISI_ORI" json:"kode_bank_disposisi_ori"`
	RekeningDisposisiOri         sql.NullString  `db:"REKENING_DISPOSISI_ORI" json:"rekening_disposisi_ori"`
	PenerimaDisposisiOri         sql.NullString  `db:"PENERIMA_DISPOSISI_ORI" json:"penerima_disposisi_ori"`
	SaldoAccrualBagihasil        sql.NullFloat64 `db:"SALDO_ACCRUAL_BAGIHASIL" json:"saldo_accrual_bagihasil"`
	CadanganBagihasilKapitalisir sql.NullFloat64 `db:"CADANGAN_BAGIHASIL_KAPITALISIR" json:"cadangan_bagihasil_kapitalisir"`
}
type DsResDepBagihasilList struct {
	IdReport       sql.NullInt64   `db:"ID_REPORT" json:"id_report"`
	Tanggal        sql.NullString  `db:"TANGGAL" json:"tanggal"`
	NomorRekening  sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	NominalTotal   sql.NullFloat64 `db:"NOMINAL_TOTAL" json:"nominal_total"`
	Pajak          sql.NullFloat64 `db:"PAJAK" json:"pajak"`
	Zakat          sql.NullFloat64 `db:"ZAKAT" json:"zakat"`
	Disposisi      sql.NullString  `db:"DISPOSISI" json:"disposisi"`
	NomorDisposisi sql.NullString  `db:"NOMOR_DISPOSISI" json:"nomor_disposisi"`
	PeriodeAwal    sql.NullString  `db:"PERIODE_AWAL" json:"periode_awal"`
	PeriodeAkhir   sql.NullString  `db:"PERIODE_AKHIR" json:"periode_akhir"`
	NisbahDasar    sql.NullFloat64 `db:"NISBAH_DASAR" json:"nisbah_dasar"`
	NisbahSpesial  sql.NullFloat64 `db:"NISBAH_SPESIAL" json:"nisbah_spesial"`
}
type GetCurrentTimeStamp struct {
	Tanggal string `db:"TANGGAL"`
	Jam     string `db:"JAM"`
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
type ListNomorRekening struct {
	NomorRekening sql.NullString `db:"NOMOR_REKENING" json:"nomor_rekening"`
	KodeCabang    sql.NullString `db:"KODE_CABANG" json:"kode_cabang"`
	SequenceNo    sql.NullInt64  `db:"SEQUENCE_NO" json:"sequence_no"`
	Status        sql.NullString `db:"STATUS" json:"status"`
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
type ResProdDeposito struct {
	KodeProduk               sql.NullString  `db:"KODE_PRODUK" json:"kode_produk"`
	NamaProduk               sql.NullString  `db:"NAMA_PRODUK" json:"nama_produk"`
	NisbahBonusDasar         sql.NullFloat64 `db:"NISBAH_BONUS_DASAR" json:"nisbah_bonus_dasar"`
	IsDapatBagiHasil         sql.NullString  `db:"IS_DAPAT_BAGI_HASIL" json:"is_dapat_bagi_hasil"`
	TarifPajak               sql.NullFloat64 `db:"TARIF_PAJAK" json:"tarif_pajak"`
	PresentaseZakatBagiHasil sql.NullFloat64 `db:"PRESENTASE_ZAKAT_BAGI_HASIL" json:"presentase_zakat_bagi_hasil"`
	KodeValuta               sql.NullString  `db:"KODE_VALUTA" json:"kode_valuta"`
	MasaPerjanjian           sql.NullString  `db:"MASA_PERJANJIAN" json:"masa_perjanjian"`
	PeriodePerjanjian        sql.NullString  `db:"PERIODE_PERJANJIAN" json:"periode_perjanjian"`
	IsKenaBiayalayananumum   sql.NullString  `db:"IS_KENA_BIAYALAYANANUMUM" json:"is_kena_biayalayananumum"`
	IsBiayaAtm               sql.NullString  `db:"IS_BIAYA_ATM" json:"is_biaya_atm"`
	JenisProduk              sql.NullString  `db:"JENIS_PRODUK" json:"jenis_produk"`
	KodeCabang               sql.NullString  `db:"KODE_CABANG" json:"kode_cabang"`
	IsBackdated              sql.NullString  `db:"IS_BACKDATED" json:"is_backdated"`
	IsBiayaRekeningDormant   sql.NullString  `db:"IS_BIAYA_REKENING_DORMANT" json:"is_biaya_rekening_dormant"`
	IsBiayaSaldoMinimum      sql.NullString  `db:"IS_BIAYA_SALDO_MINIMUM" json:"is_biaya_saldo_minimum"`
	IsBlokirDebet            sql.NullString  `db:"IS_BLOKIR_DEBET" json:"is_blokir_debet"`
	IsBlokirKredit           sql.NullString  `db:"IS_BLOKIR_KREDIT" json:"is_blokir_kredit"`
	IsCetakNota              sql.NullString  `db:"IS_CETAK_NOTA" json:"is_cetak_nota"`
	IsStatusPassbook         sql.NullString  `db:"IS_STATUS_PASSBOOK" json:"is_status_passbook"`
	IsTidakDormant           sql.NullString  `db:"IS_TIDAK_DORMANT" json:"is_tidak_dormant"`
}
type DsProfitaccrualList struct {
	IdAccrual     sql.NullInt64   `db:"ID_ACCRUAL" json:"id_accrual"`
	NomorRekening sql.NullString  `db:"NOMOR_REKENING" json:"nomor_rekening"`
	Tanggal       sql.NullString  `db:"TANGGAL" json:"tanggal"`
	Saldo         sql.NullFloat64 `db:"SALDO" json:"saldo"`
	Amount        sql.NullFloat64 `db:"AMOUNT" json:"amount"`
	Nisbah        sql.NullFloat64 `db:"NISBAH" json:"nisbah"`
	Gdrval        sql.NullFloat64 `db:"GDRVAL" json:"gdrval"`
	JenisAkad     sql.NullString  `db:"JENIS_AKAD" json:"jenis_akad"`
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
