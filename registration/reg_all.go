package registration

// QUERY_NAMES_ALL is a slice containing all predefined query names.
// This slice is used to iterate over all available queries when loading them,
// and also serves as the source for the GetRandomQueryName function.
var QUERY_NAMES_ALL = []string{
	QUERY__GetDataLainSavPlan_main,
	QUERY__GetAccountingDayByDate_main,
	QUERY__GetListLabaRugi_main,
	QUERY__GetListNeraca_subQueryDailyBalanceFirstDate,
	QUERY__GetListNeraca_subQueryDailyBalanceLastDate,
	QUERY__GetListNeraca_main,
	QUERY__UpdateDataIndividuHAR_main,
	QUERY__GetIDStandingInstruction_main,
	QUERY__GetIDRecurrentTransaction_main,
	QUERY__CreateRecurrenttrx_main,
	QUERY__CreateRegRT_main,
	QUERY__UpdateRegLayIsPassT_main,
	QUERY__DeleteListIndividu_main,
	QUERY__UpdateRegPassbook_main,
	QUERY__GetDetailSavPlan_main,
	QUERY__GetValidasiRekeningRT_main,
	QUERY__GetListRencanaLayanan_main,
	QUERY__GetAccountForclose_main,
	QUERY__IsAfiliasiRekeningRencana_main,
	QUERY__IsAfiliasiJatuhTempoDeposito_main,
	QUERY__IsAfiliasiBagiHasilDeposito_main,
	QUERY__IsAfiliasiFinancing_main,
	QUERY__IsAfiliasiPencairanRekeningRencana_main,
	QUERY__IsAfiliasiAtsPerTanggal_main,
	QUERY__IsAfiliasiAtsPerSaldo_main,
	QUERY__GetAcctAccessList_main,
	QUERY__GetPendingClrDepositList_main,
	QUERY__GetAcctSavplanCertificate_main,
	QUERY__GetAcctSavplanCertificateStandingInstruction_main,
	QUERY__GetAcctSavplanCertificateAlamat_main_A,
	QUERY__GetAcctSavplanCertificateAlamat_main_D,
	QUERY__GetAcctSavplanCertificateAlamat_main_NI,
	QUERY__GetAcctSavplanCertificateAlamat_main_NK,
	QUERY__GetPendingClrDepositByGroup_main,
	QUERY__GetSeqValRepoRekening_main,
	QUERY__GetDataNorekVal_main,
	QUERY__InsertTableValRepTemp_main,
	QUERY__GetDataValTemp_main,
	QUERY__GetSequenceNextValPostAcctTempTable_main,
	QUERY__DeleteDataValTempyIdSes_main,
	QUERY__HitungTotalDPK_main,
	QUERY__InsertToAcctTempTable_main,
	QUERY__UpdateKodeProduk_main,
	QUERY__UpdateKodeAccount_main,
	QUERY__UpdateSaldoReposisi_main,
	QUERY__GetDataFromTempTable_main,
	QUERY__InsertPostAcctIntoReposisiTable_main,
	QUERY__UpdateKodeCabangRekeningReposisi_main,
	QUERY__GetDataFromPostAcctChangeTempTable_main,
	QUERY__InsertToAcctChangeProductInfoTable_main,
	QUERY__UpdateKodeCabangAcctChangeProductInfoTable_main,
	QUERY__InsertToAcctChangeProductTempTable_main,
	QUERY__DeletePostAcctChangeTempTable_main,
	QUERY__UpdateJenisrodukPostAcctChange_main,
	QUERY__UpdateKodeAccountPostAcctChange_main,
	QUERY__UpdateKodeAccountAccruPostAcctChange_main,
	QUERY__UpdateKodeAccountBaruPostAcctChange_main,
	QUERY__UpdateKodeAccountBaruAccruPostAcctChange_main,
	QUERY__GetDataSaldoKodeCabangPostAcctChange_main,
	QUERY__UpdateNoreNotValid_main,
	QUERY__UpdateNorekChangeProd_main,
	QUERY__UpdateNamaValid_main,
	QUERY__UpdateNamaCabangValid_main,
	QUERY__UpdateNamaCabangTujuanValid_main,
	QUERY__UpdateKodeCabangTujuanValid_main,
	QUERY__UpdateValRekNotValChangeProd_main,
	QUERY__UpdateValNamaChangeProd_main,
	QUERY__UpdateValCodeProdChangeProd_main,
	QUERY__UpdateValNamaProdChangeProd_main,
	QUERY__UpdateValNamaProdBaruChangeProd_main,
	QUERY__UpdateValKodCabTujNotValChangeProd_main,
	QUERY__UpdateValKoCabSamaChangeProd_main,
	QUERY__UpdateValIsValidChangeProd_main,
	QUERY__UpdateKodeCabangTujuanNotValid_main,
	QUERY__UpdateFlagIsValid_main,
	QUERY__GetSeqValRepoRekeningProd_main,
	QUERY__InsertTableValUbahProdukTemp_main,
	QUERY__GetDataValUbahProdukTemp_main,
	QUERY__DeleteDataValUbahProdukTempyIdSes_main,
	QUERY__GetListNomorRekening_main,
	QUERY__InsertListNorekIfNotExist_main,
	QUERY__GetDataListAcctNumber_main,
	QUERY__UpdateStatusListAcctNumber_main,
	QUERY__UpdateStatusListAcctNumberSetActive_main,
	QUERY__GetNextWorkDate_main,
	QUERY__AccountingDayAvability_main,
	QUERY__ValidasiRekKasTrx_main,
	QUERY__ValidasiRekGLTrx_main,
	QUERY__ValidasiRekKodeInterfaceTrx_main,
	QUERY__ValidasiRekLiabTrx_main,
	QUERY__GetBalanceToday_main,
	QUERY__TrxForUpdate_main,
	QUERY__UpdateAcctBlock_main,
	QUERY__InsertHistBlokir_main,
	QUERY__GetHistblokirAktif_main,
	QUERY__CreateAlamatAlernatif_main,
	QUERY__DeleteAlamatAlernatif_main,
	QUERY__UpdateAlamatAlernatif_main,
	QUERY__ValidasiRegSweep_main,
	QUERY__GetAutotrfList_main,
	QUERY__GetAutosaveList_main,
	QUERY__CreateRegAutosave_main,
	QUERY__GetKodeBankTrfList_main,
	QUERY__GetKliringKotaList_main,
	QUERY__GetKliringKotaSknList_main,
	QUERY__SelectRtgsMsgCodeList_main,
	QUERY__SelectRtgsTrxCodeList_main,
	QUERY__SelectRtgsPriorCodeList_main,
	QUERY__GetBlobDataLength_main,
	QUERY__GetBlobData_main,
	QUERY__GetBlobdataId_main,
	QUERY__InsertIdBlobData_main,
	QUERY__InsertBlobData1_main,
	QUERY__InsertBlobData2_main,
	QUERY__InsertHoldDana_main,
	QUERY__InsertHoldDanaFinfFund_main,
	QUERY__GetBlokirDana_main,
	QUERY__UpdateStatusHold_main,
	QUERY__UpdateSaldoDitahan_main,
	QUERY__GetListHoldDana_main,
	QUERY__GetHoldDanaFinFund_main,
	QUERY__UpdateHoldDanaFinFund_main,
	QUERY__UpdateHoldDanaNonaktifFinFund_main,
	QUERY__GetBukuWarkat_main,
	QUERY__UpdateAktivasiWarkat_main,
	QUERY__UpdateLembarTerpakaiCair_main,
	QUERY__UpdateLembarTerpakaiCairReverse_main,
	QUERY__GetCabangDetail_main,
	QUERY__GetCardListByAccNo_main,
	QUERY__GetCardListByAccountNo_main,
	QUERY__GetCardByCardId_main,
	QUERY__GetCardPinByCardId_main,
	QUERY__GetCardByAccountNo_main,
	QUERY__GetIdRequest_main,
	QUERY__InsertCardCustomerRequest_main,
	QUERY__InsertCardAccountLinkRequest_main,
	QUERY__InsertCardpin_main,
	QUERY__HistoryCardStatus_main,
	QUERY__UpdateCard_main,
	QUERY__UpdateCardState_main,
	QUERY__InsertCardBlockRequest_main,
	QUERY__ValidasiPostChq_main,
	QUERY__GetChequeID_main,
	QUERY__GetChequeList_main,
	QUERY__GetChequeDetail_main,
	QUERY__GetChequeBySerialNum_main,
	QUERY__PostChqBook_main,
	QUERY__PostNotaDebetInternal_main,
	QUERY__GetLbvData_main,
	QUERY__GetSaldoRataRata_main,
	QUERY__GetDepList_main,
	QUERY__GetDetailDep_main,
	QUERY__CreateDeposito_main,
	QUERY__GetDepositoByNoBilyet_main,
	QUERY__GetDepositoBilyetPrint_main,
	QUERY__UpdateBilyetPrintStatus_main,
	QUERY__CreateHistBilyetDepoPrint_main,
	QUERY__GetInterbankDepositList_main,
	QUERY__ModAcctDepRT_main,
	QUERY__ModAcctDepRL_main,
	QUERY__ModAcctDepDeposito_main,
	QUERY__ModDepForclose_main,
	QUERY__InsertInfoPencairanAwalDeposito_main,
	QUERY__DeleteInfoPencairanAwalDeposito_main,
	QUERY__UpdateDepositoPencairanAwalDeposito_main,
	QUERY__GetNotaDebitOutwardDelList_main,
	QUERY__DelNotaDebitOutward_main,
	QUERY__DelNotaDebitOutwardGroup_main,
	QUERY__GetDepForClose_main,
	QUERY__GetDepWithdrawListToday_main,
	QUERY__GetDepWithdrawListHist_main,
	QUERY__GetDepWithdrawToday_main,
	QUERY__GetDepWithdrawHist_main,
	QUERY__GetDepBagihasilList_main,
	QUERY__GetCurrentTimeStamp_main,
	QUERY__GetTrxDetailDtk_main,
	QUERY__GetListDHN_main,
	QUERY__GetDetilDHIBData_main,
	QUERY__GetDetilDHNData_main,
	QUERY__GetDetiltransgroupclass_main,
	QUERY__GetAccountDetailed_main,
	QUERY__GetActiveSchedule_main,
	QUERY__GetCollateralDone_main,
	QUERY__GetNeedFindataComplete_main,
	QUERY__GetCustDataCount_main,
	QUERY__GetAccDataCount_main,
	QUERY__UpdateDroppingFinAccount_main,
	QUERY__UpdateFinAccount_main,
	QUERY__CreateCollateralAsset_main,
	QUERY__GetDetailCollateralAsset_main,
	QUERY__GetNextColAccountNumber1_main,
	QUERY__GetNextColAccountNumber2_main,
	QUERY__GetHoldDanaRekening_main,
	QUERY__UpdateSaldoHoldReleaseAgunan_main,
	QUERY__UpdateFinMusyarakahAccount_main,
	QUERY__CreateFinPaymentSchedule_main,
	QUERY__CreateFinPaymentSchedDetail_main,
	QUERY__UpdateFinPaymentSchedule_main,
	QUERY__GetListTabGiro_main,
	QUERY__GetRekeningTabgirDetail_main,
	QUERY__UpdateIsSedangTutup_main,
	QUERY__UpdateIsTutupKartuAtm_main,
	QUERY__GetAlamatAlternatif_main,
	QUERY__GetHakAksesRekening_main,
	QUERY__GetJoinAccount_main,
	QUERY__GetBagiHasil_main,
	QUERY__GetTellerDetail_main,
	QUERY__GetTieringBagiHasil_main,
	QUERY__GetTieringBiayaADM_main,
	QUERY__GetTrxList_main,
	QUERY__GetTrxByNoref_main,
	QUERY__GetTrxByIdTrx_main,
	QUERY__GetTrxDetail_main,
	QUERY__GetTrxDetailRev_main,
	QUERY__GetTrxDetailTransfer_main,
	QUERY__GetHistTrxListByNorek_main,
	QUERY__GetTrxListByNorek_main,
	QUERY__GetAcctStatementHeaderPrint_main,
	QUERY__GetSumTrxAcctStatementHeader_main,
	QUERY__GetHistTrxList_main,
	QUERY__GetSaldoByDate_main,
	QUERY__GetHisttrxDetail_main,
	QUERY__GetHisttrxDetailRev_main,
	QUERY__PostTransaksi_main,
	QUERY__PostDetailTrx1_main,
	QUERY__PostDetailTrx2_main,
	QUERY__UpdateIsReversedTransaksi_main,
	QUERY__InsertInfoTransaksi_main,
	QUERY__GetInfotransaksi_main,
	QUERY__InsertListAksesTransaksi_main,
	QUERY__GetTrxDetailByParams_main,
	QUERY__GetListTrxDetailByParams_main,
	QUERY__GetTrxDetailTolakanKliring_main,
	QUERY__GetDetailTrxTDN_main,
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
	QUERY__InsertValTrxUploadGL_main,
	QUERY__ValAccTidakValidTrxUploadGL_main,
	QUERY__ValCabangTrxUploadGL_main,
	QUERY__ValValutaTrxUploadGL_main,
	QUERY__ValAccTransaksionalTrxUploadGL_main,
	QUERY__GetResponseValTrxUploadGL_main,
	QUERY__DelTrxUploadGLTemp_main,
	QUERY__GetListTrxPayroll_main,
	QUERY__GetSequenceTempPostMultiJournal_main,
	QUERY__InsertToPostMultiJournalTempTable_main,
	QUERY__CheckBalancePostTrxMultiJournalTempTable_main,
	QUERY__UpdateNamaCabangValutaAccountPostTrxMultiJournalTempTable_main,
	QUERY__UpdateNamaCabangSaldoPostTrxMultiJournalTempTable_main,
	QUERY__UpdateProdukFloatDitahanPostTrxMultiJournalTempTable_main,
	QUERY__UpdateSaldoMinimumPostTrxMultiJournalTempTable_main,
	QUERY__UpdateKodeAccountPostTrxMultiJournalTempTable_main,
	QUERY__CheckMinimumBalancePostTrxMultiJournalTempTable_main,
	QUERY__GetDataFromPostTrxTempTable_main,
	QUERY__DeletePostTrxTempTable_main,
	QUERY__GetSequenceTempTrxPayroll_main,
	QUERY__InsertValidasiTempTrxPayroll_main,
	QUERY__CheckTmpPayrollRekeningValid_main,
	QUERY__UpdateTmpPayrollNameValid_main,
	QUERY__CheckTmpPayrollRekeningClosed_main,
	QUERY__GetAllTmpTrxPayroll_main,
	QUERY__DeleteTempTrxPayroll_main,
	QUERY__GetSequenceTempTrxMulti_main,
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
	QUERY__GetRekeningLayanan_main,
	QUERY__UpdateNilaiKurs_main,
	QUERY__UpdateNilaiEkuivalen_main,
	QUERY__GetTrxDetailByUidkey_main,
	QUERY__GetHistTrxDetailByUidkey_main,
	QUERY__GetDetailTrxDetail_main,
	QUERY__GetDetailTrxDetailHist_main,
	QUERY__GetTrxDetailByNoref_main,
	QUERY__GetHistTrxDetailByNoref_main,
	QUERY__CreateTransaksiExtInfo_main,
	QUERY__InsertTransaksiTransitKas_main,
	QUERY__InsertTransfer_main,
	QUERY__RptTrialBalance_main,
	QUERY__GetAccDayByPeriode_main,
	QUERY__GetAccount_main,
	QUERY__GetDc_main,
	QUERY__GetFeeDc_main,
	QUERY__FeeGlCredit_main,
	QUERY__GetPairGl_main,
	QUERY__GetVa_main,
	QUERY__GetVaBilling_main,
	QUERY__GetVaPaid_main,
	QUERY__GetVaHistPaid_main,
	QUERY__GetDataAgentLiabilitas_main,
	QUERY__GetListAccountOfficer_main,
	QUERY__GetFundingAgent_main,
	QUERY__GetInfoUid_main,
	QUERY__GetInfoUidForRev_main,
	QUERY__InsertInfoUid_main,
	QUERY__GetPod_main,
	QUERY__GetAccountNo_main,
	QUERY__GetAccountNo_main_sqlInsertRG,
	QUERY__GetAccountNo_main_sqlUpdateRG,
	QUERY__GetCifPortfolioByNomorNasabah_main,
	QUERY__GetIDIdv_main,
	QUERY__GetIDRegLayanan_main,
	QUERY__GetListParameterBiaya_main,
	QUERY__GetTransaksiID_main,
	QUERY__GetTransaksiDetailID_main,
	QUERY__GetCustomIdgen_main,
	QUERY__InsertCustomIdgen_main,
	QUERY__GetGLInterface_main,
	QUERY__GetGlAccount_main,
	QUERY__GetGlobalGlInterface_main,
	QUERY__InsertRakOtomatis_main,
	QUERY__GetGLList_main,
	QUERY__GetGLJurnal_main,
	QUERY__DetailGLJurnal_main,
	QUERY__DetailGLJurnalItem_main,
	QUERY__DetailGLJurnalBlock_main,
	QUERY__CoaList_main,
	QUERY__CoaListMapping_main,
	QUERY__GetGlinterface_main,
	QUERY__GetGlinterfaceTabungan_main,
	QUERY__InsertHistUbahData_main,
	QUERY__CreateDataIndividu_main,
	QUERY__GetInfoTTRByIdTrx_main,
	QUERY__DelInfoTTR_main,
	QUERY__InsertInfoPenutupanRekening_main,
	QUERY__InsertJournal_main,
	QUERY__InsertJournalItem_main,
	QUERY__InsertPendingTrxJournal_main,
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
	QUERY__GetDenominasitransaksikas_main,
	QUERY__UpdateDenominasiRekening_main,
	QUERY__InsertDenominasitransaksikas_main,
	QUERY__GetDetilDenominasiKas_main,
	QUERY__GetIdGroup_main,
	QUERY__GetIdNotadebitoutward_main,
	QUERY__InsertNotaDebitOutwardGroup_main,
	QUERY__InsertNotaDebitOutward_main,
	QUERY__GetNotadebitoutwardByIdGroup_main,
	QUERY__GetNotadebitoutwardGroupByIdGroup_main,
	QUERY__GetNDOGByIdTrx_main,
	QUERY__GetDKOByIdTrx_main,
	QUERY__GetDKTByIdTrx_main,
	QUERY__UpdateStatusNotadebitoutwardgroup_main,
	QUERY__UpdateStatusNotadebitoutward_main,
	QUERY__UpdateStatusNotaDebetOutward_main,
	QUERY__UpdateRecountSaldoFloat_main,
	QUERY__GetNotadebitoutwardByIdDbOW_main,
	QUERY__UpdateStatusNotadebitoutwardIdb_main,
	QUERY__InsertInfotarikankliring_main,
	QUERY__InsertDtkliringtarikan_main,
	QUERY__DeleteInfotarikankliringByIDTrx_main,
	QUERY__GetClrRejectReasonList_main,
	QUERY__GetCodeGlInterface_main,
	QUERY__InsertInfotolaankliring_main,
	QUERY__GetClrWithdrawList_main,
	QUERY__InsertDTKliringOutward_main,
	QUERY__GetListSetoranKliring_main,
	QUERY__GetLayananByNamaLayanan_main,
	QUERY__CreateListIndRekening_main,
	QUERY__UpdateListIndividuRekByIDIndividu_main,
	QUERY__UpdateListIndividuRekByListIndividu_main,
	QUERY__UpdateListIndividuRekByNoRek_main,
	QUERY__UpdateStatusListRek_main,
	QUERY__GetListNomorRekeningByNoRek_main,
	QUERY__InsertListNomorRekening_main,
	QUERY__CreateJointAccount_main,
	QUERY__DeleteMultiCIFLink_main,
	QUERY__GetNasabah_main,
	QUERY__GetNasabahIndividu_main,
	QUERY__GetNasabahInfo_main,
	QUERY__UpdateStatusByIdBukuwarkat_main,
	QUERY__UpdateStatusWarkat_main,
	QUERY__UpdateStatusWarkatCair_main,
	QUERY__ReverseStatusWarkatCair_main,
	QUERY__UpdateCHQStatusWarkat_main,
	QUERY__GetWarkatDetail_main,
	QUERY__GetValidationLimitTrx_main,
	QUERY__GetMaxDbHarian_main,
	QUERY__GetLimitDebetHarian_main,
	QUERY__GetCounterDebetHarian_main,
	QUERY__GetGlobalParam_main,
	QUERY__GetGlobalparamRem_main,
	QUERY__CreatePassbookHistBalance_main,
	QUERY__GetPassbookType_main,
	QUERY__GetAccountPassbook_main,
	QUERY__CekNomorRegisterBuku_main,
	QUERY__GetAlternateAddress_main,
	QUERY__GetCustomerAddress_main,
	QUERY__GetCustomerPersonalAddress_main,
	QUERY__GetCustomerCorporateAddress_main,
	QUERY__CreatePassbookHistPrint_main,
	QUERY__CreateRegisterSavingBook_main,
	QUERY__UpdateRegisterPassbook_main,
	QUERY__GetValNorekPassbook_main,
	QUERY__GetDataTransaksi_main,
	QUERY__GetSeqInfoCetakPassbook_main,
	QUERY__GetRegPassbookLast_main,
	QUERY__GetCollectionDataTrx_main,
	QUERY__GetBalance_main,
	QUERY__InsertInfoCetakPassbook_main,
	QUERY__InsertDetailTransaksiPassbook_main,
	QUERY__GetjenisPasbook_main,
	QUERY__GetDicMapCodeTrx_main,
	QUERY__GetTglByIdDetailTrx_main,
	QUERY__GetDataNTrx_main,
	QUERY__UpdateInfoCetakPassbook_main,
	QUERY__GetDataInfoCetakPassbook_main,
	QUERY__UpdateInfoCetakPassbookByIdReg_main,
	QUERY__UpdatePassbookTxUnprint_main,
	QUERY__GetPassbookLastprint_main,
	QUERY__GetDataTrxUnPrint_main,
	QUERY__GetPassbookTotalBaris_main,
	QUERY__UpdateInfoCetakPassWhenAllData_main,
	QUERY__GetProduk_main,
	QUERY__GetProdukTabungan_main,
	QUERY__GetProdukRencana_main,
	QUERY__GetProdukDeposito_main,
	QUERY__InsertGlInterface_main,
	QUERY__UpdateGlInterface_main,
	QUERY__InsertTieringItem_main,
	QUERY__ModTieringItem_main,
	QUERY__InsertListLimitProduk_main,
	QUERY__ModLimitProduk_main,
	QUERY__ModListLimitProduk_main,
	QUERY__InsertDataProdukSav_main,
	QUERY__InsertDataProdukDep_main,
	QUERY__InsertDataProdukGiro_main,
	QUERY__InsertProdukDep_main,
	QUERY__UpdateProdukDep_main,
	QUERY__GetSeqIdTieringNisbahbonus_main,
	QUERY__GetSeqIdTieringBiayaAdmin_main,
	QUERY__UpdateIdTeringNisbahBonus_main,
	QUERY__UpdateIdTieringBiayaAdmin_main,
	QUERY__InsertTieringPlan_main,
	QUERY__GetKodeProdukExist_main,
	QUERY__GetKodeProdukDepExist_main,
	QUERY__ModDataProduk_main,
	QUERY__DelTieringItem_main,
	QUERY__DelLimitProd_main,
	QUERY__GetProfitaccrualList_main,
	QUERY__GetRcCodeList_main,
	QUERY__GetListRas_main,
	QUERY__NonactiveRecurrenttrx_main,
	QUERY__GetCifCorpRedundantList_main,
	QUERY__InsertUpdateRefSched_main,
	QUERY__UpdateIdMulticifrelation_main,
	QUERY__UpdateIdSumberDana_main,
	QUERY__UpdateIdTujuanRekening_main,
	QUERY__UpdateIdCfgroup_main,
	QUERY__UpdateIdSumberDanaTrx_main,
	QUERY__UpdateIdTujuanTrx_main,
	QUERY__UpdateIdNegara_main,
	QUERY__CreateRegisterLayanan_main,
	QUERY__GetRegisterLayanan_main,
	QUERY__UpdateRegisterLayanan_main,
	QUERY__NonActiveRegisterLayanan_main,
	QUERY__CreateRegPassbook_main,
	QUERY__GetRegisterRt_main,
	QUERY__GetListProdSav_main,
	QUERY__GetListSavPlan_main,
	QUERY__GetListDep_main,
	QUERY__GetProdDepList_main,
	QUERY__GetListCA_main,
	QUERY__GetAccountLiabilitas_main,
	QUERY__UpdateAcctDormat_main,
	QUERY__GetLimitProduk_main,
	QUERY__GetRekeningTabCollateral_main_D,
	QUERY__GetRekeningTabCollateral_main_GTJ,
	QUERY__GetRekeningCustomer_main,
	QUERY__CreateRekenigCustomer_main,
	QUERY__GetRekeningGenerator_main,
	QUERY__InsertRekeningGenerator_main,
	QUERY__UpdateCounterRekeningGenerator_main,
	QUERY__GetRekeningNasabah_main,
	QUERY__CreateRekenigLiabilitas_main,
	QUERY__UpdateRekenigLiabilitas_main,
	QUERY__GetRekeningLiab_main,
	QUERY__UpdateRekLiabBatalTTR_main,
	QUERY__UpdateRekLiabTTR_main,
	QUERY__UpdateRekLiabPencairanDeposito_main,
	QUERY__UpdateRekLiabPencairanDepositoRev_main,
	QUERY__GetSaldoDitahan_main,
	QUERY__UpdateSaldoFloatRekenigLiabilitas_main,
	QUERY__UpdateRekLiabilitasCollateral_main,
	QUERY__UpdateSaldoDitahanLiabiliitas_main,
	QUERY__UpdateSaldoUnholdLiabitas_main,
	QUERY__UnblokirRekLiabilitasFinFund_main,
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
	QUERY__UpdateStatusRekeningTransaksi_main,
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
	QUERY__GetReportCodeList_main,
	QUERY__GetDataReportNasabah_main,
	QUERY__GetDataReportRekeningBaru_main,
	QUERY__GetDataReportNasabahBaru_main,
	QUERY__GetDataReportSetoranKliring_main,
	QUERY__GetDataReportRekeningTutup_main,
	QUERY__GetDataReportRekeningBlokir_main,
	QUERY__GetDataReportRekeningHoldDana_main,
	QUERY__GetDataReportRegistrasiAutosave_main,
	QUERY__GetDataReportNonAktifAutosave_main,
	QUERY__GetDataReportAlamatAlternatifNasabah_main,
	QUERY__GetDataReportRekeningUbahProduk_main,
	QUERY__GetDataReportTabunganBagihasilMudharabah_main,
	QUERY__GetDataReportTabunganBonusWadiah_main,
	QUERY__GetDataReportBiayaAdministrasi_main,
	QUERY__GetDataReportRekeningSaldoHarian_main,
	QUERY__GetDataReportRekeningReposisi_main,
	QUERY__GetDataReportRekeningSaldoRekap_main,
	QUERY__GetDataReportRekeningTidakAktif_main,
	QUERY__GetDataReportBeaMaterai_main,
	QUERY__GetDataReportTabunganCetakKepalaBuku_main,
	QUERY__GetDataReportPemblokiranCekBG_main,
	QUERY__GetDataReportAktivasiWarkat_main,
	QUERY__GetDataReportRekeningDormanOtomatis_main,
	QUERY__GetDataReportRegistrasiATSPerTanggal_main,
	QUERY__GetDataReportNonAktifATSPerTanggal_main,
	QUERY__GetDataReportRekeningJoinAccount_main,
	QUERY__GetDataReportWICBaru_main,
	QUERY__GetDataReportDepositoJatuhTempoHariIni_main,
	QUERY__GetDataReportDepositoOutstanding_main,
	QUERY__GetDataReportDepositoBagiHasil_main,
	QUERY__GetDataReportTplDepositoBaru_main,
	QUERY__GetDataReportTplPerubahanNasabah_main,
	QUERY__GetDataReportTplDepositoBreak_main,
	QUERY__GetDataReportTplRekRencanaBaru_main,
	QUERY__GetDataReportTplNasabahReposisi_main,
	QUERY__GetDataReportTplNasabahGabung_main,
	QUERY__GetDataReportTplDepositoJatuhTempoHariIniTransfer_main,
	QUERY__GetDataReportTplDepositoBlokir_main,
	QUERY__GetDataReportTplDepositoAROHariIni_main,
	QUERY__GetDataReportTplDepositoJatuhTempoBesok_main,
	QUERY__GetDataReportTplDepositoBagHasTransfer_main,
	QUERY__GetDataReportTplDepositoSaldoRekap_main,
	QUERY__GetDataReportTransaksiSendiri_main,
	QUERY__GetDataReportTransaksiJurnal_main,
	QUERY__GetDataTrxStandingInstruction_main,
	QUERY__GetCTRValue_main,
	QUERY__GetDataReportTplAMLTransKeuanganTunai_main,
	QUERY__GetDataReportAMLTotalCIF_main,
	QUERY__GetDataTrxCabRTGSTransferIn_main,
	QUERY__GetDataReportAMLNasabahOverLimitTrans_main,
	QUERY__GetDataReportTplAMLTransKeuanganTunaiSummary_main,
	QUERY__GetDataTrxCabRTGSTransferOut_main,
	QUERY__GetDataReportSKNTransferKeluar_main,
	QUERY__GetDataReportSKNTransferMasuk_main,
	QUERY__GetDataReportTplTransaksiCabSKNTarikanKliring_main,
	QUERY__GetDataReportTplTransaksiCabSKNTransferRetur_main,
	QUERY__GetDataReportTransaksiCabSKNDebet_main,
	QUERY__GetDataReportTransaksiCabSKNReturDebet_main,
	QUERY__GetDataReportTplTransaksiCabSKNTransferInFailed_main,
	QUERY__GetDataReportTplTransaksiCabRTGSTransferRetur_main,
	QUERY__GetDataReportTplTransaksiCabRTGSTransferInFailed_main,
	QUERY__GetDataReportTransaksiJurnalSendiri_main,
	QUERY__GetDataReportRekapTransaksiSendiri_main,
	QUERY__GetDataReportTransaksiReversalOtorisasi_main,
	QUERY__GetReportDetail_main,
	QUERY__GetRestrictCode_main,
	QUERY__UpdateAccountRestrictCode_main,
	QUERY__UpdateIsReversed_main,
	QUERY__InsertReverseTransaction_main,
	QUERY__GetRevTrxByIdTrx_main,
	QUERY__GetSknTrxcodeList_main,
	QUERY__CreateStandingInstruction_main,
	QUERY__GetDataHeaderSvs_main,
	QUERY__GetDataListImageSvs_main,
	QUERY__CheckSpecimentAcctList_main,
	QUERY__UpdateSpecimentAcctList_main,
	QUERY__PostSpecimentAcctList_main,
	QUERY__InsertCustAccountImage_main,
	QUERY__GetMaxSequenceImageTag_main,
	QUERY__DeleteSpecimentAcct_main,
}
