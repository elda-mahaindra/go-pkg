package registration

var QUERY_NAMES_FUNDTXBATCH = []string{
	QUERY__GetListTrxPayroll_main,
	QUERY__GetInfoUid_main,
	QUERY__GetSequenceTempPostMultiJournal_main,
	QUERY__InsertToPostMultiJournalTempTable_main,
	QUERY__CheckBalancePostTrxMultiJournalTempTable_main,
	QUERY__UpdateNamaCabangValutaAccountPostTrxMultiJournalTempTable_main,
	QUERY__UpdateProdukFloatDitahanPostTrxMultiJournalTempTable_main,
	QUERY__UpdateSaldoMinimumPostTrxMultiJournalTempTable_main,
	QUERY__UpdateKodeAccountPostTrxMultiJournalTempTable_main,
	QUERY__CheckMinimumBalancePostTrxMultiJournalTempTable_main,
	QUERY__GetDataFromPostTrxTempTable_main,
	QUERY__DeletePostTrxTempTable_main,
	QUERY__InsertInfoUid_main,
	QUERY__GetSequenceTempTrxPayroll_main,
	QUERY__InsertValidasiTempTrxMulti_main,
	QUERY__CheckTmpTrxMultiCodeTypeValid_main,
	QUERY__CheckTmpTrxMultiRekExist_main,
	QUERY__CheckTmpTrxMultiRekExistByTypeCode_main,
	QUERY__CheckTmpTrxMultiByMutasi_main,
	QUERY__UpdateRekeningTrxTmpTrxMulti_main,
	QUERY__UpdateRekeningLiabilitasTmpTrxMulti_main,
	QUERY__UpdateProdukTmpTrxMulti_main,
	QUERY__UpdateRekTutupByRekTrxTmpMulti_main,
	QUERY__UpdateDormantRekTmpMultiTrx_main,
	QUERY__ValidasiNominalTmpTrxMulti_main,
	QUERY__UpdateRekTutupByLiabilityTmpTrxMulti_main,
	QUERY__UpdateRekBlokirByLiabilityTmpTrxMulti_main,
	QUERY__UpdateRekBlokirKreditByLiabilityTmpTrxMulti_main,
	QUERY__UpdateRekBlokirDebitByLiabilityTmpTrxMulti_main,
	QUERY__ValidasiSaldoEfektifTmpTrxMulti_main,
	QUERY__ValidasiKodeRCTmpTrxMulti_main,
	QUERY__GetAllTmpMultiTrx_main,
	QUERY__GetTotalNominalTmpMultiTrx_main,
	QUERY__DeleteTempMultiTrx_main,
	QUERY__InsertValidasiTempTrxPayroll_main,
	QUERY__CheckTmpPayrollRekeningValid_main,
	QUERY__UpdateTmpPayrollNameValid_main,
	QUERY__CheckTmpPayrollRekeningClosed_main,
	QUERY__GetAllTmpTrxPayroll_main,
	QUERY__DeleteTempTrxPayroll_main,
	QUERY__GetSequenceId_main,
	QUERY__InsertValidasiTransRek_main,
	QUERY__UpdateJenisMutasi_main,
	QUERY__UpdateIsValidRekening_main,
	QUERY__UpdateNama_main,
	QUERY__UpdateRekTutup_main,
	QUERY__UpdateRekDormant_main,
	QUERY__UpdateRekTutupLiabilitas_main,
	QUERY__UpdateRekBlokirLiabilitas_main,
	QUERY__SelectValTrxTemp_main,
	QUERY__DeleteValTrxTemp_main,
	QUERY__GetSequenceNextValPostAcctTempTable_main,
	QUERY__InsertValTrxUploadGL_main,
	QUERY__ValAccTidakValidTrxUploadGL_main,
	QUERY__ValCabangTrxUploadGL_main,
	QUERY__ValValutaTrxUploadGL_main,
	QUERY__ValAccTransaksionalTrxUploadGL_main,
	QUERY__GetResponseValTrxUploadGL_main,
	QUERY__DelTrxUploadGLTemp_main,
	QUERY__GetTransaksiID_main,
	QUERY__GetPod_main,
	QUERY__GetCustomIdgen_main,
	QUERY__InsertCustomIdgen_main,
	QUERY__PostTransaksi_main,
	QUERY__ValidasiRekKasTrx_main,
	QUERY__TrxForUpdate_main,
	QUERY__UpdateSaldoRekeningTransaksi_main,
	QUERY__ValidasiRekGLTrx_main,
	QUERY__ValidasiRekKodeInterfaceTrx_main,
	QUERY__ValidasiRekLiabTrx_main,
	QUERY__GetTransaksiDetailID_main,
	QUERY__PostDetailTrx2_main,
	QUERY__InsertTransaksiTransitKas_main,
	QUERY__InsertPendingTrxJournal_main,
	QUERY__GetGlobalParam_main,
	QUERY__GetGlobalGlInterface_main,
	QUERY__UpdateNilaiKurs_main,
	QUERY__UpdateNilaiEkuivalen_main,
	QUERY__AccountingDayAvability_main,
	QUERY__CheckValiditasDataJournal_main,
	QUERY__CheckAccountInstanceJournal_main,
	QUERY__CheckBalanceJournal_main,
	QUERY__MergeJournalNumber_main,
	QUERY__InsertSelectJournal_main,
	QUERY__InsertSelectJournalItem_main,
	QUERY__BalancingRAKOtomatisJournal_main,
	QUERY__UpdateTrxStatusJournal_main,
	QUERY__InsertCompletedJournal_main,
	QUERY__DeletePendingJournal_main,
}
