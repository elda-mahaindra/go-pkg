package registration

var QUERY_NAMES_FUNDACCOUNT = []string{
	QUERY__UpdateDataIndividuHAR_main,
	QUERY__GetIDStandingInstruction_main,
	QUERY__GetIDRecurrentTransaction_main,
	QUERY__CreateRecurrenttrx_main,
	QUERY__CreateRegRT_main,
	QUERY__GetDetailSavPlan_main,
	QUERY__GetListRencanaLayanan_main,
	QUERY__GetAccountForclose_main,
	QUERY__IsAfiliasiJatuhTempoDeposito_main,
	QUERY__IsAfiliasiBagiHasilDeposito_main,
	QUERY__IsAfiliasiFinancing_main,
	QUERY__IsAfiliasiRekeningRencana_main,
	QUERY__IsAfiliasiPencairanRekeningRencana_main,
	QUERY__IsAfiliasiAtsPerTanggal_main,
	QUERY__IsAfiliasiAtsPerSaldo_main,
	QUERY__GetAcctAccessList_main,
	QUERY__GetListNomorRekening_main,
	QUERY__InsertListNorekIfNotExist_main,
	QUERY__GetDataListAcctNumber_main,
	QUERY__UpdateStatusListAcctNumber_main,
	QUERY__UpdateStatusListAcctNumberSetActive_main,
	QUERY__AccountingDayAvability_main,
	QUERY__ValidasiRekKasTrx_main,
	QUERY__ValidasiRekGLTrx_main,
	QUERY__ValidasiRekKodeInterfaceTrx_main,
	QUERY__ValidasiRekLiabTrx_main,
	QUERY__TrxForUpdate_main,
	QUERY__UpdateAcctBlock_main,
	QUERY__InsertHistBlokir_main,
	QUERY__GetHistblokirAktif_main,
	QUERY__CreateAlamatAlernatif_main,
	QUERY__DeleteAlamatAlernatif_main,
	QUERY__GetSaldoRataRata_main,
	QUERY__GetDataAgentLiabilitas_main,
	QUERY__GetFundingAgent_main,
	QUERY__GetInfoUid_main,
	QUERY__InsertInfoUid_main,
	QUERY__GetPod_main,
	QUERY__GetCifPortfolioByNomorNasabah_main,
	QUERY__GetIDIdv_main,
	QUERY__GetIDRegLayanan_main,
	QUERY__GetTransaksiID_main,
	QUERY__GetTransaksiDetailID_main,
	QUERY__GetCustomIdgen_main,
	QUERY__InsertCustomIdgen_main,
	QUERY__GetGlobalGlInterface_main,
	QUERY__InsertHistUbahData_main,
	QUERY__CreateDataIndividu_main,
	QUERY__InsertPendingTrxJournal_main,
	QUERY__CheckValiditasDataJournal_main,
	QUERY__CheckAccountInstanceJournal_main,
	QUERY__CheckBalanceJournal_main,
	QUERY__MergeJournalNumber_main,
	QUERY__BalancingRAKOtomatisJournal_main,
	QUERY__UpdateTrxStatusJournal_main,
	QUERY__InsertCompletedJournal_main,
	QUERY__DeletePendingJournal_main,
	QUERY__GetLayananByNamaLayanan_main,
	QUERY__CreateListIndRekening_main,
	QUERY__UpdateListIndividuRekByIDIndividu_main,
	QUERY__UpdateListIndividuRekByNoRek_main,
	QUERY__UpdateStatusListRek_main,
	QUERY__GetListNomorRekeningByNoRek_main,
	QUERY__InsertListNomorRekening_main,
	QUERY__CreateJointAccount_main,
	QUERY__DeleteMultiCIFLink_main,
	QUERY__GetNasabah_main,
	QUERY__GetNasabahInfo_main,
	QUERY__GetGlobalParam_main,
	QUERY__GetGlobalparamRem_main,
	QUERY__GetBalance_main,
	QUERY__GetProduk_main,
	QUERY__GetProdukRencana_main,
	QUERY__GetCifCorpRedundantList_main,
	QUERY__UpdateIdMulticifrelation_main,
	QUERY__UpdateIdSumberDana_main,
	QUERY__UpdateIdTujuanRekening_main,
	QUERY__UpdateIdNegara_main,
	QUERY__CreateRegisterLayanan_main,
	QUERY__CreateRegPassbook_main,
	QUERY__GetAccountLiabilitas_main,
	QUERY__UpdateAcctDormat_main,
	QUERY__GetRekeningCustomer_main,
	QUERY__CreateRekenigCustomer_main,
	QUERY__GetRekeningGenerator_main,
	QUERY__InsertRekeningGenerator_main,
	QUERY__UpdateCounterRekeningGenerator_main,
	QUERY__CreateRekenigLiabilitas_main,
	QUERY__UpdateRekenigLiabilitas_main,
	QUERY__GetRekeningLiab_main,
	QUERY__GetJangkaWaktuRencana_main,
	QUERY__CreateRekeningRencana_main,
	QUERY__UpdateRekenigRencana_main,
	QUERY__GetSIRencana_main,
	QUERY__UpdateSIRencana_main,
	QUERY__UpdateRCTRencana_main,
	QUERY__UpdateRLRencana_main,
	QUERY__GetRekeningTransaksi_main,
	QUERY__CreateRekenigTransaksi_main,
	QUERY__UpdateRekeningTransaksi_main,
	QUERY__UpdateSaldoRekeningTransaksi_main,
	QUERY__GetDetailTrxByIdDetail_main,
	QUERY__GetLastId_main,
	QUERY__InsertIdGen_main,
	QUERY__InsertTmpRemCandidate_main,
	QUERY__InsertTmpRemCandidateRtgs_main,
	QUERY__InsertRemTrxStatusHist_main,
	QUERY__InsertRemTrxRecord_main,
	QUERY__InsertDoiCredit_main,
	QUERY__MergeDoiCredit_main,
	QUERY__InsertRemTransSkn_main,
	QUERY__DeleteRemTrans_main,
	QUERY__CreateRemittanceCandidate_main,
	QUERY__CreateRTStatusHistory_main,
	QUERY__CreateRTTransaction_main,
	QUERY__MergeRTTransaction_main,
	QUERY__CreateRTTransferDetail_main,
	QUERY__CreateRTSingleTransfer_main,
	QUERY__MergeRTSingleTransfer_main,
	QUERY__CreateRemittanceTransaction_main,
	QUERY__InsertReverseTransaction_main,
	QUERY__CreateStandingInstruction_main,
	QUERY__GetListTabGiro_main,
	QUERY__GetRekeningTabgirDetail_main,
	QUERY__UpdateIsSedangTutup_main,
	QUERY__UpdateIsTutupKartuAtm_main,
	QUERY__GetAlamatAlternatif_main,
	QUERY__GetHakAksesRekening_main,
	QUERY__GetJoinAccount_main,
	QUERY__PostTransaksi_main,
	QUERY__PostDetailTrx2_main,
	QUERY__UpdateIsReversedTransaksi_main,
	QUERY__GetRekeningLayanan_main,
	QUERY__UpdateNilaiKurs_main,
	QUERY__UpdateNilaiEkuivalen_main,
	QUERY__InsertTransaksiTransitKas_main,
	QUERY__GetAccount_main,
}