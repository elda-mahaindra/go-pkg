package registration

var QUERY_NAMES_FUNDTX = []string{
	QUERY__ValidasiRekKasTrx_main,
	QUERY__ValidasiRekGLTrx_main,
	QUERY__ValidasiRekKodeInterfaceTrx_main,
	QUERY__ValidasiRekLiabTrx_main,
	QUERY__GetBalanceToday_main,
	QUERY__TrxForUpdate_main,
	QUERY__UpdateLembarTerpakaiCair_main,
	QUERY__UpdateLembarTerpakaiCairReverse_main,
	QUERY__GetChequeBySerialNum_main,
	QUERY__GetNotaDebitOutwardDelList_main,
	QUERY__DelNotaDebitOutward_main,
	QUERY__DelNotaDebitOutwardGroup_main,
	QUERY__GetTrxDetailDtk_main,
	QUERY__GetInfoUid_main,
	QUERY__GetInfoUidForRev_main,
	QUERY__InsertInfoUid_main,
	QUERY__GetPod_main,
	QUERY__GetTransaksiID_main,
	QUERY__GetTransaksiDetailID_main,
	QUERY__GetCustomIdgen_main,
	QUERY__InsertCustomIdgen_main,
	QUERY__GetGlobalGlInterface_main,
	QUERY__GetGLList_main,
	QUERY__GetGLJurnal_main,
	QUERY__DetailGLJurnal_main,
	QUERY__DetailGLJurnalItem_main,
	QUERY__DetailGLJurnalBlock_main,
	QUERY__GetInfoTTRByIdTrx_main,
	QUERY__DelInfoTTR_main,
	QUERY__InsertPendingTrxJournal_main,
	QUERY__CheckValiditasDataJournal_main,
	QUERY__CheckValiditasDataJournal_filter_IdTransaksiNotZero,
	QUERY__CheckValiditasDataJournal_filter_IdTransaksiZero,
	QUERY__CheckAccountInstanceJournal_main,
	QUERY__CheckAccountInstanceJournal_filter_IdTransaksiNotZero,
	QUERY__CheckAccountInstanceJournal_filter_IdTransaksiZero,
	QUERY__CheckBalanceJournal_main,
	QUERY__CheckBalanceJournal_filter_IdTransaksiNotZero,
	QUERY__CheckBalanceJournal_filter_IdTransaksiZero,
	QUERY__MergeJournalNumber_main,
	QUERY__InsertSelectJournal_main,
	QUERY__InsertSelectJournalItem_main,
	QUERY__BalancingRAKOtomatisJournal_main,
	QUERY__BalancingRAKOtomatisJournal_filter_IdTransaksiNotZero,
	QUERY__BalancingRAKOtomatisJournal_filter_IdTransaksiZero,
	QUERY__UpdateTrxStatusJournal_main,
	QUERY__InsertCompletedJournal_main,
	QUERY__DeletePendingJournal_main,
	QUERY__GetDenominasitransaksikas_main,
	QUERY__UpdateDenominasiRekening_main,
	QUERY__InsertDenominasitransaksikas_main,
	QUERY__GetDetilDenominasiKas_main,
	QUERY__GetDKOByIdTrx_main,
	QUERY__GetDKTByIdTrx_main,
	QUERY__UpdateStatusNotadebitoutwardgroup_main,
	QUERY__UpdateStatusNotadebitoutward_main,
	QUERY__UpdateStatusNotadebitoutwardIdb_main,
	QUERY__DeleteInfotarikankliringByIDTrx_main,
	QUERY__InsertDTKliringOutward_main,
	QUERY__UpdateStatusWarkatCair_main,
	QUERY__ReverseStatusWarkatCair_main,
	QUERY__GetWarkatDetail_main,
	QUERY__GetValidationLimitTrx_main,
	QUERY__GetMaxDbHarian_main,
	QUERY__GetLimitDebetHarian_main,
	QUERY__GetCounterDebetHarian_main,
	QUERY__GetGlobalParam_main,
	QUERY__GetRekeningLiab_main,
	QUERY__UpdateRekLiabBatalTTR_main,
	QUERY__UpdateSaldoFloatRekenigLiabilitas_main,
	QUERY__GetRekeningTransaksi_main,
	QUERY__UpdateSaldoRekeningTransaksi_main,
	QUERY__UpdateStatusRekeningTransaksi_main,
	QUERY__CreateRemittanceCandidate_main,
	QUERY__CreateRTStatusHistory_main,
	QUERY__CreateRTTransaction_main,
	QUERY__MergeRTTransaction_main,
	QUERY__CreateRTTransferDetail_main,
	QUERY__CreateRTSingleTransfer_main,
	QUERY__MergeRTSingleTransfer_main,
	QUERY__CreateRemittanceTransaction_main,
	QUERY__InsertReverseTransaction_main,
	QUERY__GetRevTrxByIdTrx_main,
	QUERY__GetTellerDetail_main,
	QUERY__GetTrxList_main,
	QUERY__GetTrxByIdTrx_main,
	QUERY__GetTrxDetail_main,
	QUERY__GetTrxDetailRev_main,
	QUERY__GetTrxListByNorek_main,
	QUERY__GetHistTrxList_main,
	QUERY__GetSaldoByDate_main,
	QUERY__PostTransaksi_main,
	QUERY__PostDetailTrx2_main,
	QUERY__UpdateIsReversedTransaksi_main,
	QUERY__InsertInfoTransaksi_main,
	QUERY__GetInfotransaksi_main,
	QUERY__GetListTrxDetailByParams_main,
	QUERY__GetRekeningLayanan_main,
	QUERY__UpdateNilaiKurs_main,
	QUERY__UpdateNilaiEkuivalen_main,
	QUERY__GetTrxDetailByUidkey_main,
	QUERY__GetHistTrxDetailByUidkey_main,
	QUERY__GetHistTrxDetailByNoref_main,
	QUERY__InsertTransaksiTransitKas_main,
	QUERY__GetHistTrxListByNorek_main,
	QUERY__GetRcCodeList_main,
	QUERY__AccountingDayAvability_main,
}
