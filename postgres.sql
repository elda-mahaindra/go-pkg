-- name: GetDataLainSavPlan-main
select rl.kode_program, rl.kode_marketing_referensi, rl.status_kelengkapan 
		from %[1]s.rekeningliabilitas rl  
		where nomor_rekening = $1

-- name: GetAccountingDayByDate-main
select isworkday, periode_status
		from %[1]s.accountingday
		where datevalue = to_timestamp($1, 'yyyy-mm-dd')

-- name: GetListLabaRugi-main
select
			ac.account_code, ac.account_name, ac.account_level, 
			ac.account_type, ac.is_detail, ' ' as project_no, 
			(case when db.balance is null then 0.0 else db.balance end) as balance,
			(case when db.balance_ekuiv is null then 0.0 else db.balance_ekuiv end) as balance_ekuiv,
			(case when ac.account_type = 'I' then 'PENDAPATAN' else 'BIAYA' end) as account_group
		from %[1]s.account ac
		left outer join
			(select h.fl_parentaccountcode as account_code,  
				sum(d1.balancecumulative) as balance, 
				sum(d1.balancecumulative_ekuiv) as balance_ekuiv
			from %[1]s.accounthierarchy h, %[1]s.accountinstance ai, %[1]s.dailybalance d1,
				(select accountinstance_id, max(datevalue) as maxdate
				from %[1]s.dailybalance where datevalue <= to_timestamp($1, 'yyyy-mm-dd')
				group by accountinstance_id) d2
			where h.fl_childaccountcode = ai.account_code 
				and ai.branch_code = $2 
				and ai.currency_code = $3
				and ai.accountinstance_id = d1.accountinstance_id
				and d1.accountinstance_id = d2.accountinstance_id
				and d1.datevalue = d2.maxdate
				%[2]s
			group by h.fl_parentaccountcode
		) db on ac.account_code = db.account_code
		where ac.account_type in ('I','X') and ac.is_detail='F'  
		union  
		select ac.account_code, ac.account_name, ac.account_level, 
			ac.account_type, ac.is_detail, db.project_no,
			(case when db.balance is null then 0.0 else db.balance end) as balance,
			(case when db.balance_ekuiv is null then 0.0 else db.balance_ekuiv end) as balance_ekuiv,
			(case when ac.account_type = 'I' then 'PENDAPATAN' else 'BIAYA' end) as account_group
		from %[1]s.account ac
		inner join %[1]s.project p on 1=1                  
		left outer join
			(select h.fl_parentaccountcode as account_code,  
				d1.project_no,
				sum(d1.balancecumulative) as balance, 
				sum(d1.balancecumulative_ekuiv) as balance_ekuiv
			from %[1]s.accounthierarchy h, %[1]s.accountinstance ai, %[1]s.dailyprojectbalance d1,
				(select accountinstance_id, project_no, max(datevalue) as maxdate
				from %[1]s.dailyprojectbalance where datevalue <= to_timestamp($1, 'yyyy-mm-dd')
				group by accountinstance_id, project_no) d2
			where h.fl_childaccountcode = ai.account_code 
				and ai.branch_code = $2 
				and ai.currency_code = $3
				and ai.accountinstance_id = d1.accountinstance_id
				and d1.accountinstance_id = d2.accountinstance_id
				and d1.project_no = d2.project_no
				and d1.datevalue = d2.maxdate
				%[2]s
				group by h.fl_parentaccountcode, d1.project_no
		) db on ac.account_code = db.account_code and p.project_no = db.project_no
		where ac.account_type in ('I', 'X') 
		and ac.is_detail='T'
		and ( db.project_no = '0000' or ( db.project_no <> '0000' and db.balance <> 0.0) )           
		order by account_code

-- name: GetListLabaRugi-filter-CompanyCode
and ac.company_code = $4

-- name: GetListNeraca-subQueryDailyBalanceFirstDate
SELECT
			ai.accountinstance_id, '%[2]s' as p_date,
				ai.account_code, ai.branch_code, ai.currency_code, ai.balance_sign,
				case when db.debit is null then 0.0 else db.debit end as debit,
				case when db.debit_ekuiv is null then 0.0 else db.debit_ekuiv end as debit_ekuiv,
	            case when db.datevalue is null then 0.0 else db.datevalue end as datevalue,
				case when db.credit is null then 0.0 else db.credit end as credit,
				case when db.credit_ekuiv is null then 0.0 else db.credit_ekuiv end as credit_ekuiv,
				case when db.balance is null then 0.0 else db.balance end as balance,
				case when db.balance_ekuiv is null then 0.0 else db.balance_ekuiv end as balance_ekuiv,
				case when db.balancecumulative is null then 0.0 else db.balancecumulative end as balancecumulative,
				case when db.balancecumulative_ekuiv is null then 0.0 else db.balancecumulative_ekuiv end as balancecumulative_ekuiv,
				case when db.balance_ratarata is null then 0.0 else db.balance_ratarata end as balance_ratarata,
				case when db.balance_ratarata_ekuiv is null then 0.0 else db.balance_ratarata_ekuiv end as balance_ratarata_ekuiv
		FROM %[1]s.accountinstance ai
			LEFT OUTER JOIN
			(SELECT d1.* FROM %[1]s.dailybalance d1,
			(SELECT accountinstance_id, MAX(datevalue) as maxdate
			FROM %[1]s.dailybalance WHERE datevalue <= TO_DATE('%[2]s', 'YYYY-MM-DD')
			GROUP BY accountinstance_id) d2
			WHERE d1.accountinstance_id = d2.accountinstance_id
			AND d1.datevalue = d2.maxdate) db
		ON ai.accountinstance_id = db.accountinstance_id

-- name: GetListNeraca-subQueryDailyBalanceLastDate
SELECT
			ai.accountinstance_id, '%[2]s' as p_date,
				ai.account_code, ai.branch_code, ai.currency_code, ai.balance_sign,
				case when db.debit is null then 0.0 else db.debit end as debit,
				case when db.debit_ekuiv is null then 0.0 else db.debit_ekuiv end as debit_ekuiv,
				case when db.credit is null then 0.0 else db.credit end as credit,
				case when db.credit_ekuiv is null then 0.0 else db.credit_ekuiv end as credit_ekuiv,
				case when db.balance is null then 0.0 else db.balance end as balance,
	            case when db.datevalue is null then 0.0 else db.datevalue end as datevalue,
				case when db.balance_ekuiv is null then 0.0 else db.balance_ekuiv end as balance_ekuiv,
				case when db.balancecumulative is null then 0.0 else db.balancecumulative end as balancecumulative,
				case when db.balancecumulative_ekuiv is null then 0.0 else db.balancecumulative_ekuiv end as balancecumulative_ekuiv,
				case when db.balance_ratarata is null then 0.0 else db.balance_ratarata end as balance_ratarata,
				case when db.balance_ratarata_ekuiv is null then 0.0 else db.balance_ratarata_ekuiv end as balance_ratarata_ekuiv
		FROM %[1]s.accountinstance ai
			LEFT OUTER JOIN
			(SELECT d1.* FROM %[1]s.dailybalance d1,
			(SELECT accountinstance_id, MAX(datevalue) as maxdate
			FROM %[1]s.dailybalance WHERE datevalue <= TO_DATE('%[2]s', 'YYYY-MM-DD')
			GROUP BY accountinstance_id) d2
			WHERE d1.accountinstance_id = d2.accountinstance_id
			AND d1.datevalue = d2.maxdate) db
		ON ai.accountinstance_id = db.accountinstance_id

-- name: GetListNeraca-main
select 
			pg.accountgroup_code
			, pg.accountgroup_name
			, sum(x.balancecumulative_ekuiv) balancecumulative_ekuiv
			, sum(x.MonthBeginBalance) MonthBeginBalance
		from %[1]s.accountgroup pg
		inner join (
			select 	
					ag.fl_parent_accountgroup as accountgroup_code
					, ah.fl_parentaccountcode AS account_code
					, a.account_code AS account_code_detail
					, a.account_name AS account_name_detail
					, db1.accountinstance_id
					, db1.datevalue
					, db1.branch_code , db1.currency_code
					, db1.balancecumulative , db1.balancecumulative_ekuiv
					, db1.balance , db1.balance_ekuiv
					, db1.balance_ratarata ,  db1.balance_ratarata_ekuiv
					, db1.balance_sign
					, a.account_type
					, (case when (a.account_type in ('I','X')) then 0.0 else db2.balancecumulative end) as MonthBeginBalance
					, (case when (a.account_type in ('I','X')) then 0.0 else db2.balancecumulative_ekuiv end) MonthBeginBalanceEkuiv
			from %[1]s.accounthierarchy ah
					, (%[2]s) db1
					, (%[3]s) db2
					, %[1]s.account a
					, %[1]s.accountgroup ag
			where ah.fl_childaccountcode = db1.account_code
					and a.account_code = db1.account_code
					and db1.accountinstance_id = db2.accountinstance_id
					and a.account_group_code = ag.accountgroup_code
		) x on x.accountgroup_code = pg.accountgroup_code  
		where x.balancecumulative_ekuiv <> 0.0
		group by pg.fl_parent_accountgroup, pg.accountgroup_code, pg.accountgroup_name 
		order by pg.accountgroup_code, pg.accountgroup_name

-- name: UpdateDataIndividuHAR-main
update %[1]s.individu 
		set nama_lengkap = $1, tempat_lahir = $2, tanggal_lahir = $3, status_perkawinan = $4, kewarganegaraan = $5,
        jenis_identitas = $6, nomor_identitas = $7, alamat_rumah_jalan = $8, alamat_rumah_rt = $9, alamat_rumah_rw = $10,
        alamat_rumah_kelurahan = $11, alamat_rumah_kecamatan = $12, alamat_rumah_kota_kabupaten = $13, alamat_rumah_provinsi = $14,
        alamat_rumah_kode_pos = $15, kode_pekerjaan = $16, telepon_hp_nomor = $17, jenis_identitas_lain = $18,
        nomor_identitas_lain = $19, jenis_kelamin = $20, status_pep = $21, kode_negara_asal = $22
		where id_individu = $23

-- name: GetIDStandingInstruction-main
select nextval('%[1]s.seq_standinginstruction') as id 

-- name: GetIDRecurrentTransaction-main
select nextval('%[1]s.seq_recurrenttransaction') as id 

-- name: CreateRecurrenttrx-main
insert into %[1]s.recurrenttransaction
		(recid, instructionid, nominal, keterangan, status_data, tgl_proses, tgl_proses_berikutnya, tgl_kadaluarsa, periode_proses, 
			jarak_periode_proses, maksimum_gagal_proses, proses_hanya_tglproses, akumulasi_proses_gagal, jumlah_total_proses,
			jumlah_sudah_proses, jumlah_gagal_proses)
		values 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,  $11, $12, $13, $14, $15, $16 )

 -- name: CreateRegRT-main
 insert into %[1]s.registerrt
		(id_register, description, recid, nomor_rekening_kredit)
		values 
		($1, $2, $3, $4)

 -- name: UpdateRegLayIsPassT-main
 UPDATE %[1]s.registerlayanan SET
			NOMOR_REKENING_LAYANAN = $1,
			STATUS_REGISTER_LAYANAN = $2,
			JENIS_REGISTER_LAYANAN = $3,
			TANGGAL_REGISTRASI= $4
		WHERE id_register = $5

-- name: DeleteListIndividu-main
DELETE FROM %[1]s.listindividurekening
		WHERE nomor_rekening = $1

-- name: UpdateRegPassbook-main
update %[1]s.registerpassbook set IS_BARU_REGISTER = $1
		where ID_REGISTER = $2

-- name: GetDetailSavPlan-main
select rl.nomor_rekening ,rt.nama_rekening ,rt.keterangan ,rl.nomor_nasabah ,n.jenis_nasabah ,to_char(rl.tanggal_buka ,'YYYY-MM-DD') as tanggal_buka
			,to_char(rl.tanggal_tutup, 'YYYY-MM-DD') as tanggal_tutup ,rl.kode_produk ,p.nama_produk ,rl.is_sedang_ditutup ,rt.kode_cabang ,c.nama_cabang
			,rl.is_join_account ,rt.kode_valuta ,rt.nama_valuta ,rl.status_restriksi ,rl.kode_status_retriksi ,rl.is_blokir ,rl.is_dapat_bagi_hasil ,rl.nisbah_spesial
			,rl.persentase_zakat_bagi_hasil ,rl.tarif_pajak ,rl.nisbah_dasar ,rl.is_tiering_nisbahbonus ,rt.saldo ,rl.saldo_ditahan ,rl.saldo_float ,p.saldo_minimum
			,rt.saldo - (coalesce(rl.saldo_float, 0.0) + coalesce(rl.saldo_ditahan, 0.0) + coalesce(p.saldo_minimum, 0.0)) as saldo_efektif, rl.is_blokir_debet ,rl.is_blokir_kredit 
			,rl.is_cetak_nota ,rl.is_status_passbook ,rl.is_tidak_dormant ,rl.is_biaya_rekening_dormant ,rl.is_kena_biayalayananumum ,rl.is_biaya_saldo_minimum 
			,rl.is_biaya_atm ,rl.kode_multicifrelation ,rl.user_update ,to_char(rt.tanggal_update, 'YYYY-MM-DD') as tanggal_update ,rl.status_kelengkapan ,rl.alamat_kirim 
			,to_char(rl.tgl_trans_cabang_terakhir, 'YYYY-MM-DD') as tgl_trans_cabang_terakhir ,to_char(rl.tgl_trans_echannel_terakhir, 'YYYY-MM-DD') as tgl_trans_echannel_terakhir 
			,to_char(rl.tgl_transaksi_terakhir, 'YYYY-MM-DD') as tgl_transaksi_terakhir ,rl.kode_program ,rl.kode_marketing_referensi ,rl.kode_marketing_pertama 
			,rl.kode_marketing_current ,rr.nomor_rekening_induk ,rr.nomor_rekening_pencairan ,to_char(rr.tanggal_jatuh_tempo, 'YYYY-MM-DD') as tanggal_jatuh_tempo ,rr.setoran_rutin 
			,rr.tanggal_setoran_rutin ,rr.jangka_waktu ,rr.jumlah_kewajiban_jt  ,rr.jumlah_realisasi ,rr.jumlah_tunggakan ,rl.jenis_rekening_liabilitas, rl.kode_sumber_dana 
			,rl.kode_tujuan_rekening ,rl.nisbah_bagi_hasil ,rt.user_input ,rl.user_otorisasi ,rt.status_rekening ,to_char(rt.tanggal_aktifitas_terakhir, 'YYYY-MM-DD') as tanggal_aktifitas_terakhir,
			case
				when rt.status_rekening = 1 then 'Aktif'
				when rt.status_rekening = 2 then 'Dormant'
				when rt.status_rekening = 3 then 'Tutup'
			end as status_rekening_desc
		from %[1]s.rekeningliabilitas rl
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = rl.nomor_rekening
			inner join %[1]s.produk p on p.kode_produk = rl.kode_produk
			inner join %[1]s.rekeningrencana rr on rr.nomor_rekening = rt.nomor_rekening
			inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah 
			inner join %[2]s.cabang c on c.kode_cabang = rt.kode_cabang 
		where rl.nomor_rekening = $1

-- name: GetValidasiRekeningRT-main
select rrt.nomor_rekening_kredit ,rl.status_register_layanan
		from %[1]s.registerrt rrt
	 	inner join %[1]s.registerlayanan rl on rl.id_register = rrt.id_register
		where nomor_rekening_kredit = $1

-- name: GetListRencanaLayanan-main
select id_layanan from %[1]s.listrencanalayanan
		where id_layanan = $1 and kode_produk = $2

-- name: GetAccountForclose-main
select rt.nama_rekening,  p.kode_produk, p.nama_produk, concat(concat(rt.kode_cabang, '-'), c.nama_cabang) as cabang, 
			rt.kode_cabang, c.nama_cabang, rt.saldo, rt.nomor_rekening, concat(concat(rt.kode_valuta, '-'), cu.description) as valuta, 
			rl.saldo_ditahan, rl.is_blokir, rl.saldo_float, NVL(rl.is_sedang_ditutup, 'F') is_sedang_ditutup, p.biaya_penutupan_rekening,
			coalesce(rl.is_tutup_kartu_atm, 'F') as is_tutup_kartu_atm
		from %[1]s.rekeningliabilitas rl
			inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
			inner join %[1]s.produk p on rl.kode_produk = p.kode_produk
			left join %[1]s.currency cu on rt.kode_valuta = cu.currency_code 
			inner join %[2]s.cabang c on rt.kode_cabang = c.kode_cabang 
		where rl.nomor_rekening = $1

--name: IsAfiliasiRekeningRencana-main
select rr.nomor_rekening, nomor_rekening_induk
		from %[1]s.rekeningrencana rr
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = rr.nomor_rekening 
		where rr.nomor_rekening_induk = $1
			and rt.status_rekening != 3


-- name: IsAfiliasiJatuhTempoDeposito-main
select d.nomor_rekening, d.rekening_disposisi
		from %[1]s.deposito d 
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening
		where d.rekening_disposisi =  $1
			and rt.status_rekening != 3
		order by d.nomor_rekening

-- name: IsAfiliasiBagiHasilDeposito-main
select rl.nomor_rekening_disposisi, d.nomor_rekening
		from %[1]s.deposito d 
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening 
			inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = d.nomor_rekening 
		where rl.nomor_rekening_disposisi = $1
			and rt.status_rekening != 3
		order by d.nomor_rekening

-- name: IsAfiliasiFinancing-main
select fa.nomor_rekening, fa.norek_paymentsrc
		from %[2]s.finaccount fa
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = fa.nomor_rekening 
		where rt.status_rekening = 1
			and fa.norek_paymentsrc = $1

-- name: IsAfiliasiPencairanRekeningRencana-main
select rr.nomor_rekening, nomor_rekening_pencairan
		from %[1]s.rekeningrencana rr
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = rr.nomor_rekening 
		where rr.nomor_rekening_pencairan = $1
			and rt.status_rekening != 3

-- name: IsAfiliasiAtsPerTanggal-main
select si.debitaccountno, si.creditaccountno
		from %[1]s.standinginstruction si
			inner join %[1]s.recurrenttransaction rt on  si.instructionid = rt.instructionid
			inner join %[1]s.registerrt rr on rt.recid = rr.recid
			inner join %[1]s.registerlayanan rl on rr.id_register = rl.id_register
			inner join %[1]s.rekeningtransaksi rdb on rdb.nomor_rekening = si.debitaccountno
			inner join %[1]s.rekeningtransaksi rcr on rcr.nomor_rekening = si.creditaccountno
		where si.debitaccountno = $1
			and rl.status_register_layanan = 'A'

-- name: IsAfiliasiAtsPerSaldo-main
select rsw.nomor_rekening_tujuan
		from %[1]s.registerautosweep rsw 
			inner join %[1]s.registerlayanan rl on rsw.id_register = rl.id_register
		where rl.nomor_rekening_layanan = $1
			and rl.status_register_layanan = 'A' 

-- name: GetAcctAccessList-main
select i.id_individu , i.nama_lengkap
		, i.jenis_identitas , i.nomor_identitas
		, i.jenis_identitas_lain , i.nomor_identitas_lain
		, i.Telepon_HP_Nomor
		, i.alamat_rumah_jalan
		, i.tempat_lahir
		, i.tanggal_lahir
		, i.jenis_kelamin
		, ir.keterangan
	from  %[1]s.listindividurekening ir
		inner join %[1]s.individu i on i.id_individu = ir.id_individu
	where ir.nomor_rekening = $1
		and is_aktif='T'
		and tipeindividurekening is null

-- name: GetPendingClrDepositList-main
select a.id_group ,
			a.tanggal_input ,
			a.tanggal_efektif ,
			a.status_notadebit ,
			a.nomor_rekening ,
			a.nomor_aplikasi ,
			a.keterangan ,
			a.kode_cabang ,
			a.kode_valuta ,
			a.inputer ,
			a.is_otorisasi ,
			a.tipe_grup ,
			a.tanggal_eksekusi_titipan ,
			a.user_eksekusi_titipan ,
			a.jam_eksekusi_titipan ,
			b.Nama_Rekening ,
			(SELECT Sum(nominal) FROM %[1]s.notadebitoutward nd WHERE nd.id_group = a.id_group) AS total_nominal,
            (SELECT count(nominal) FROM %[1]s.notadebitoutward nd WHERE nd.id_group = a.id_group) AS total_warkat             
		from %[1]s.NotaDebitOutwardGroup a
			inner join %[1]s.RekeningTransaksi b  on a.Nomor_Rekening = b.Nomor_Rekening               
			and a.Is_Otorisasi = 'T'
			and a.Status_NotaDebit = 'T' 
			%[2]s              			            
		order by a.Tanggal_Efektif

-- name: GetPendingClrDepositList-filter-KodeCabang
and a.Kode_Cabang = $1

-- name: GetPendingClrDepositList-filter-Tanggal
and a.Tanggal_Efektif = to_date($2, 'yyyy-mm-dd')

-- name: GetAcctSavplanCertificate-main
SELECT 
		A.NOMOR_REKENING, 
		B.IS_BLOKIR_DEBET, 
		B.IS_BLOKIR_KREDIT, 
		D.SALDO, 
		to_char(B.TANGGAL_BUKA, 'YYYY-MM-DD') tanggal_buka, 
		A.TANGGAL_JATUH_TEMPO, 
		A.NOMOR_REKENING, 
		D.NAMA_REKENING, 
		A.NOMOR_REKENING_INDUK, 
		RI.NAMA_REKENING as NAMA_REKENING_INDUK,
		A.NOMOR_REKENING_PENCAIRAN,  
		RP.NAMA_REKENING as NAMA_REKENING_PENCAIRAN,
		A.SETORAN_RUTIN, 
		A.TANGGAL_SETORAN_RUTIN, 
		B.NOMOR_NASABAH,
		Nvl(B.Alamat_Kirim,'N') alamat_kirim ,
		N.jenis_nasabah
	FROM %[1]v.REKENINGRENCANA A
		INNER JOIN %[1]v.REKENINGLIABILITAS B ON A.Nomor_Rekening = B.Nomor_Rekening 
		INNER JOIN %[1]v.REKENINGCUSTOMER C ON B.Nomor_Rekening = C.Nomor_Rekening 
		INNER JOIN %[1]v.REKENINGTRANSAKSI D ON C.Nomor_Rekening = D.Nomor_Rekening
		INNER JOIN %[1]v.NASABAH N ON N.Nomor_Nasabah = B.Nomor_Nasabah
        LEFT JOIN %[1]v.REKENINGTRANSAKSI RI on RI.nomor_Rekening = A.NOMOR_REKENING_INDUK
        LEFT JOIN %[1]v.REKENINGTRANSAKSI RP on Rp.nomor_Rekening = A.NOMOR_REKENING_PENCAIRAN
	WHERE A.NOMOR_REKENING = $1

-- name: GetAcctSavplanCertificateStandingInstruction-main
SELECT count(1) as standing_instruction
	FROM
	%[1]v.STANDINGINSTRUCTION A
	WHERE (A.DEBITACCOUNTNO = $1 AND A.CREDITACCOUNTNO = $2)

-- name: GetAcctSavplanCertificateAlamat-main-A
select nomor_rekening, alamat_jalan || ' ' || alamat_kecamatan|| ' ' || alamat_rtrw|| ' ' || alamat_kelurahan|| ' ' || alamat_provinsi|| ' ' || alamat_kode_pos|| ' ' || alamat_kota_kabupaten as AS alamat
			from %[1]s.alamatalternatif 
			where nomor_rekening = $1 

-- name: GetAcctSavplanCertificateAlamat-main-D
select alamat_surat_Jalan || ' ' || alamat_surat_Kode_Pos || ' ' || alamat_surat_RTRW || ' ' || alamat_surat_Kelurahan || ' ' || alamat_surat_Kecamatan || ' ' || alamat_surat_Kota_Kabupaten || ' ' || alamat_surat_Provinsi as alamat
			from %[1]s.nasabah
			where nomor_nasabah = $1 

-- name: GetAcctSavplanCertificateAlamat-main-NI
select i.Alamat_rumah_Jalan || ' ' || i.Alamat_rumah_Kode_Pos || ' ' || i.Alamat_rumah_RTRW || ' ' || i.Alamat_rumah_Kelurahan || ' ' || i.Alamat_rumah_Kecamatan || ' ' || i.Alamat_rumah_Kota_Kabupaten || ' ' || i.Alamat_rumah_Provinsi as alamat
			FROM %[1]s.nasabahindividu ni
			INNER JOIN individu i ON ni.id_individu = i.id_individu
			WHERE ni.nomor_nasabah = $1

-- name: GetAcctSavplanCertificateAlamat-main-NK
SELECT Alamat_Jalan || ' ' || Alamat_Kode_Pos || ' ' || Alamat_RTRW || ' ' || Alamat_Kelurahan || ' ' || Alamat_Kecamatan || ' ' || Alamat_Kota_Kabupaten || ' ' || Alamat_Provinsi as alamat
			FROM %[1]s.nasabahkorporat
			where nomor_nasabah = $1

-- name: GetPendingClrDepositByGroup-main
select a.id_group ,
			a.tanggal_input ,
			a.tanggal_efektif ,
			a.status_notadebit ,
			a.nomor_rekening ,
			a.nomor_aplikasi ,
			a.keterangan ,
			a.kode_cabang ,
			a.kode_valuta ,
			a.inputer ,
			a.is_otorisasi ,
			a.tipe_grup ,
			a.tanggal_eksekusi_titipan ,
			a.user_eksekusi_titipan ,
			a.jam_eksekusi_titipan ,
			b.Nama_Rekening,
			n.nominal    
		from %[1]s.NotaDebitOutwardGroup a
			inner join %[1]s.NOTADEBITOUTWARD n on n.NOMOR_REKENING = a.NOMOR_REKENING
			inner join %[1]s.RekeningTransaksi b  on a.Nomor_Rekening = b.Nomor_Rekening               
			and a.Is_Otorisasi = 'T'
			and a.Status_NotaDebit = 'T' 
			and a.id_group = $1          			            
		order by a.Tanggal_Efektif

-- name: GetSeqValRepoRekening-main
select  nextval('%[1]s.seq_validasi_reposisi_rekening') 

-- name: GetDataNorekVal-main
select r.nomor_rekening, r.nama_rekening, r.kode_cabang, c.nama_cabang 
		from %[1]s.rekeningtransaksi r 
			left join %[2]s.cabang c on r.kode_cabang = c.kode_cabang 
		where r.nomor_rekening = $1

-- name: InsertTableValRepTemp-main
INSERT INTO %[1]s.VALIDASI_REPOSISI_REKENING 
			(id_session, nomor_rekening, nama_rekening, kode_cabang, nama_cabang, kode_cabang_tujuan, nama_cabang_tujuan, is_valid, invalid_message)
		VALUES(
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)

-- name: GetDataValTemp-main
SELECT tmp.NOMOR_REKENING, tmp.NAMA_REKENING, tmp.KODE_CABANG, tmp.NAMA_CABANG, tmp.KODE_CABANG_TUJUAN, tmp.NAMA_CABANG_TUJUAN , tmp.IS_VALID, tmp.INVALID_MESSAGE 
		FROM %[1]s.validasi_reposisi_rekening tmp
		WHERE tmp.ID_SESSION = $1

-- name: GetSequenceNextValPostAcctTempTable-main
SELECT  nextval('%[1]s.seq_validasi_reposisi_rekening')
		FROM %[1]s.VALIDASI_REPOSISI_REKENING

-- name: DeleteDataValTempyIdSes-main
delete from  %[1]s.validasi_reposisi_rekening tmp where id_session=$1

-- name: HitungTotalDPK-main
select sum(saldo * kurs_buku) as totaldpk
		from %[1]s.rekeningtransaksi a
			inner join %[1]s.rekeningliabilitas b on a.nomor_rekening = b.nomor_rekening 
			inner join %[1]s.currency c on a.kode_valuta = c.currency_code
		where a.status_rekening != 3 
			and nomor_nasabah = $1

-- name: InsertToAcctTempTable-main
INSERT INTO %[1]v.VALIDASI_REPOSISI_REKENING 
			(id_session, nomor_rekening, kode_cabang, kode_cabang_tujuan)
		VALUES(
			$1, $2, $3, $4
		)

-- name: UpdateKodeProduk-main
MERGE INTO %[1]s.validasi_reposisi_rekening tmp 
		USING %[2]s.rekeningliabilitas rl 
		ON (tmp.NOMOR_REKENING  = rl.NOMOR_REKENING )
		WHEN MATCHED THEN UPDATE SET kode_produk = rl.kode_produk
		WHERE id_session = $1

-- name: UpdateKodeAccount-main
MERGE INTO %[1]s.validasi_reposisi_rekening tmp 
		USING %[2]s.glinterface gi
		ON (tmp.kode_produk  = gi.kode_produk AND gi.kode_Interface IN ('Saldo_Plus','glnomi'))
		WHEN MATCHED THEN UPDATE SET kode_account = gi.kode_account		
		WHERE id_session = $1

-- name: UpdateSaldoReposisi-main
MERGE INTO %[1]s.validasi_reposisi_rekening tmp 
		USING %[2]s.rekeningtransaksi rt
		ON (tmp.nomor_rekening  = rt.nomor_rekening )
		WHEN MATCHED THEN UPDATE SET saldo_reposisi = rt.saldo	
		WHERE id_session = $1

-- name: GetDataFromTempTable-main
SELECT tmp.NOMOR_REKENING, tmp.KODE_CABANG, tmp.KODE_CABANG_TUJUAN, tmp.SALDO_REPOSISI, tmp.KODE_PRODUK, tmp.KODE_ACCOUNT
		FROM %[1]s.validasi_reposisi_rekening tmp 
		WHERE tmp.ID_SESSION = $1

-- name: InsertPostAcctIntoReposisiTable-main
INSERT INTO %[1]v.INFOREPOSISIREKENING (id_transaksi, tanggal, nomor_Rekening, cabang_asal, cabang_tujuan, user_input, user_otorisasi, cabang_input, saldo_reposisi, kode_produk)
		VALUES(
			nextval('%[1]v.seq_transaksi'), current_timestamp, $1, $2, $3, $4, $5, $6, $7, $8
		)

-- name: UpdateKodeCabangRekeningReposisi-main
update %[1]v.rekeningtransaksi r 
		set kode_cabang = $1
		where nomor_rekening= $2

-- name: GetDataFromPostAcctChangeTempTable-main
SELECT 
			tmp.NOMOR_REKENING,
			tmp.KODE_PRODUK,
			tmp.KODE_PRODUK_BARU,
			(CASE WHEN tmp.SALDO is NULL THEN 0 ELSE tmp.SALDO END) as SALDO,
			tmp.KODE_CABANG,
			tmp.KODE_ACCOUNT,
			tmp.KODE_ACCOUNT_BARU
		FROM %[1]s.VALIDASI_UBAHPRODUK_REKENING tmp 
		WHERE tmp.ID_SESSION = $1

-- name: InsertToAcctChangeProductInfoTable-main
INSERT INTO %[1]v.INFOUBAHPRODUKREKENING ( 
			id_info, 
			nomor_Rekening, 
			kode_produk_lama, 
			kode_produk_baru, 
			user_input, 
			user_otorisasi, 
			cabang_input, 
			saldo
		) VALUES (
			nextval('seq_transaksi'), 
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7
		)

-- name: UpdateKodeCabangAcctChangeProductInfoTable-main
update %[1]v.rekeningliabilitas r 
			set kode_produk = $1
			where nomor_Rekening = $2

-- name: InsertToAcctChangeProductTempTable-main
INSERT INTO %[1]v.VALIDASI_UBAHPRODUK_REKENING 
			(id_session, nomor_rekening, kode_produk, kode_produk_baru)
		VALUES(
			$1, $2, $3, $4
		)

-- name: DeletePostAcctChangeTempTable-main
delete from  %[1]s.validasi_ubahproduk_rekening tmp where id_session=$1

-- name: UpdateJenisrodukPostAcctChange-main
MERGE INTO %[1]v.validasi_ubahproduk_rekening tmp
		USING %[2]v.produk c 
		ON (c.kode_produk = tmp.kode_produk)
		WHEN MATCHED THEN UPDATE SET jenis_produk = c.jenis_produk
		WHERE ID_SESSION = $1

-- name: UpdateKodeAccountPostAcctChange-main
MERGE INTO %[1]v.validasi_ubahproduk_rekening tmp
		USING %[2]v.glinterface gi
		ON (tmp.kode_produk  = gi.kode_produk AND gi.kode_Interface IN ('Saldo_Plus','glnomi'))
		WHEN MATCHED THEN UPDATE SET kode_account = gi.kode_account
		WHERE ID_SESSION = $1

-- name: UpdateKodeAccountAccruPostAcctChange-main
MERGE INTO %[1]v.validasi_ubahproduk_rekening tmp
		USING %[2]v.glinterface gi
		ON (tmp.kode_produk  = gi.kode_produk AND gi.kode_Interface IN ('Saldo_Plus','glnomi'))
		WHEN MATCHED THEN UPDATE SET kode_account_accru = gi.kode_account
		WHERE ID_SESSION = $1

-- name: UpdateKodeAccountBaruPostAcctChange-main
MERGE INTO %[1]v.validasi_ubahproduk_rekening tmp
		USING %[2]v.glinterface gi
		ON (tmp.kode_produk_baru  = gi.kode_produk AND gi.kode_Interface IN ('Cadangan_BGH','glcadb'))
		WHEN MATCHED THEN UPDATE SET kode_account_baru = gi.kode_account
		WHERE ID_SESSION = $1

-- name: UpdateKodeAccountBaruAccruPostAcctChange-main
MERGE INTO %[1]v.validasi_ubahproduk_rekening tmp
		USING %[2]v.glinterface gi
		ON (tmp.kode_produk_baru  = gi.kode_produk AND gi.kode_Interface IN ('Cadangan_BGH','glcadb'))
		WHEN MATCHED THEN UPDATE SET kode_account_accru_baru = gi.kode_account
		WHERE ID_SESSION = $1

-- name: GetDataSaldoKodeCabangPostAcctChange-main
MERGE INTO %[1]v.validasi_ubahproduk_rekening tmp
		USING %[2]v.rekeningtransaksi rt
		ON (tmp.nomor_rekening  = rt.nomor_rekening )
		WHEN MATCHED THEN UPDATE SET saldo = rt.saldo, kode_cabang = rt.kode_cabang
		WHERE ID_SESSION = $1

-- name: UpdateNoreNotValid-main
UPDATE %[2]s.validasi_reposisi_rekening tmp
	SET is_valid ='F'
	  , invalid_message = invalid_message || 'nomor rekening tidak terdaftar, '
	WHERE ID_SESSION = $1
	  AND NOT EXISTS (SELECT 1 FROM %[1]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening)

-- name: UpdateNorekChangeProd-main
UPDATE %[1]s.validasi_ubahproduk_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'nomor rekening tidak terdaftar, '
	WHERE ID_SESSION = $1
	AND NOT EXISTS (SELECT 1 FROM %[2]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening)

-- name: UpdateNamaValid-main
MERGE INTO %[1]s.validasi_reposisi_rekening tmp
	USING %[2]s.rekeningtransaksi r 
	ON (r.nomor_rekening = tmp.NOMOR_rekening)
	WHEN MATCHED THEN UPDATE SET NAMA_rekening =r.nama_rekening
		,kode_cabang = r.kode_cabang
	WHERE ID_SESSION = $1

-- name: UpdateNamaCabangValid-main
MERGE INTO %[1]s.validasi_reposisi_rekening tmp
	USING %[2]s.cabang c 
	ON (c.kode_cabang = tmp.kode_cabang)
	WHEN MATCHED THEN UPDATE SET NAMA_CABANG = c.nama_cabang
	WHERE ID_SESSION = $1

-- name: UpdateNamaCabangTujuanValid-main
MERGE INTO %[1]s.validasi_reposisi_rekening tmp
	USING %[2]s.cabang c 
	ON (c.kode_cabang = tmp.KODE_CABANG_tujuan)
	WHEN MATCHED THEN UPDATE SET 
	NAMA_CABANG_tujuan  = c.nama_cabang
	WHERE ID_SESSION = $1

-- name: UpdateKodeCabangTujuanValid-main
UPDATE %[1]s.validasi_reposisi_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'kode cabang tujuan tidak terdaftar, '
	WHERE ID_SESSION = $1
	AND NOT EXISTS (SELECT 1 FROM %[2]s.cabang c WHERE c.KODE_CABANG =tmp.KODE_CABANG_tujuan)

-- name: UpdateValRekNotValChangeProd-main
UPDATE %[1]s.validasi_ubahproduk_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'nomor rekening tidak terdaftar, '
	WHERE ID_SESSION = $1
		AND NOT EXISTS (SELECT 1 FROM %[2]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening)

-- name: UpdateValNamaChangeProd-main
MERGE INTO %[1]s.validasi_ubahproduk_rekening tmp
	USING rekeningtransaksi r 
	ON (r.nomor_rekening = tmp.NOMOR_rekening)
	WHEN MATCHED THEN UPDATE SET NAMA_rekening =r.nama_rekening
	WHERE ID_SESSION = $1

-- name: UpdateValCodeProdChangeProd-main
MERGE INTO %[1]s.validasi_ubahproduk_rekening tmp 
	USING %[2]s.rekeningliabilitas rl 
	ON (tmp.NOMOR_REKENING  = rl.NOMOR_REKENING )
	WHEN MATCHED THEN UPDATE SET kode_produk = rl.kode_produk
	WHERE id_session= $1

-- name: UpdateValNamaProdChangeProd-main
MERGE INTO %[1]s.validasi_ubahproduk_rekening tmp
	USING %[2]s.produk c 
	ON (c.kode_produk = tmp.kode_produk)
	WHEN MATCHED THEN UPDATE SET nama_produk = c.nama_produk
				, jenis_produk = c.jenis_produk
	WHERE ID_SESSION = $1

-- name: UpdateValNamaProdBaruChangeProd-main
MERGE INTO %[1]s.validasi_ubahproduk_rekening tmp
	USING %[2]s.produk c 
	ON (c.kode_produk = tmp.kode_produk_baru)
	WHEN MATCHED THEN UPDATE SET 
	nama_produk_baru  = c.nama_produk
	, jenis_produk = c.jenis_produk
	WHERE ID_SESSION = $1

-- name: UpdateValKodCabTujNotValChangeProd-main
UPDATE %[1]s.validasi_ubahproduk_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'kode produk baru tidak terdaftar, '
	WHERE ID_SESSION = $1
		AND NOT EXISTS (SELECT 1 FROM %[2]s.produk c WHERE c.kode_produk =tmp.kode_produk_baru)

-- name: UpdateValKoCabSamaChangeProd-main
UPDATE %[1]s.validasi_ubahproduk_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'kode produk baru tidak boleh sama dgn kode produk lama, '
	WHERE ID_SESSION = $1
	AND tmp.kode_produk =tmp.kode_produk_baru

-- name: UpdateValIsValidChangeProd-main
UPDATE %[1]s.validasi_ubahproduk_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'tipe kode produk baru berbeda dgn kode produk lama, '
	WHERE ID_SESSION = $1
	AND tmp.jenis_produk <> tmp.jenis_produk_baru

-- name: UpdateKodeCabangTujuanNotValid-main
UPDATE %[1]s.validasi_reposisi_rekening tmp
	SET is_valid ='F'
	, invalid_message = invalid_message || 'kode cabang tujuan tidak boleh sama dgn kode cabang asal, '
	WHERE ID_SESSION = $1
	AND tmp.KODE_CABANG =tmp.KODE_CABANG_tujuan

-- name: UpdateFlagIsValid-main
UPDATE %[1]s.validasi_reposisi_rekening tmp
	SET is_valid ='T'
	, invalid_message = ' '
	WHERE ID_SESSION = $1


-- name: GetSeqValRepoRekeningProd-main
SELECT  nextval('%[1]s.seq_val_ubahproduk_rekening')  

-- name: InsertTableValUbahProdukTemp-main
INSERT INTO %[1]s.validasi_ubahproduk_rekening 
			(
				id_session, nomor_rekening, nama_rekening, kode_produk, jenis_produk, 
				kode_produk_baru,nama_produk_baru,is_valid, invalid_message
			)
		VALUES(
			$1, $2, $3, $4, $5, $6, $7, $8, $9)

-- name: GetDataValUbahProdukTemp-main
SELECT 
			NOMOR_REKENING, NAMA_REKENING, KODE_PRODUK, NAMA_PRODUK, 
			JENIS_PRODUK, KODE_PRODUK_BARU, NAMA_PRODUK_BARU,
			IS_VALID, INVALID_MESSAGE
		FROM %[1]s.validasi_ubahproduk_rekening tmp
		WHERE tmp.ID_SESSION = $1

-- name: DeleteDataValUbahProdukTempyIdSes-main
delete from  %[1]s.validasi_ubahproduk_rekening tmp where id_session=$1

-- name: GetListNomorRekening-main
select nomor_rekening
		from %[1]s.listnomorrekening
		where kode_cabang = $1
			and sequence_no >= $2
			and sequence_no <= ($2 + $3 - 1)
			and status = 'A'
		order by sequence_no

-- name: InsertListNorekIfNotExist-main
insert into %[1]s.listnomorrekening 
		(nomor_rekening, kode_cabang, sequence_no, status)
        select *
        from ( 
			select $1 || lpad(rownum + $2 - 1, 6,'0') || mod10($3 || lpad(rownum + $4 - 1, 6,'0')) as nomor_rekening, 
        		$5, rownum + $6 - 1 as sequence_no, 'A' as status
        	
        	connect by rownum <= $7
        ) q
        where not exists (select 1 from %[1]s.listnomorrekening l where l.nomor_rekening = q.nomor_rekening)

-- name: GetDataListAcctNumber-main
select * from %[1]s.listnomorrekening
		where nomor_rekening = $1

-- name: UpdateStatusListAcctNumber-main
update %[1]s.listnomorrekening
		set status = 'P'
		where nomor_rekening = $1

-- name: UpdateStatusListAcctNumberSetActive-main
update %[1]s.listnomorrekening
		set status = 'A'
		where nomor_rekening = $1

-- name: GetNextWorkDate-main
select to_char(min(datevalue), 'yyyy-mm-dd') as next_work_date 
		from %[1]s.accountingday a
		where datevalue >= to_date('2021-01-10','yyyy-mm-dd') 
			and coalesce(isworkday,'F') = 'T'

-- name: AccountingDayAvability-main
select datevalue, periode_status 
		from %[1]s.accountingday accd
		where accd.datevalue = to_date($1, 'YYYY-MM-DD')

-- name: ValidasiRekKasTrx-main
select rt.nomor_rekening, rt.nama_rekening, rt.kode_valuta, rt.kode_cabang, rt.saldo, a.accountinstance_id, rt.kode_account, rt.saldo as saldo_efektif,
			rt.status_rekening 
		from %[1]s.rekeningtransaksi rt
			left join %[1]s.transitkas k on rt.nomor_rekening = k.nomor_rekening
			left join %[1]s.accountinstance a on a.account_code = rt.kode_account
				and rt.kode_valuta = a.currency_code
				and rt.kode_cabang = a.branch_code
		where rt.nomor_rekening = $1

-- name: ValidasiRekGLTrx-main
select rt.nomor_rekening, rt.nama_rekening, rt.kode_valuta, rt.kode_cabang, rt.saldo, a.accountinstance_id, rt.kode_account, ai.normal_balance_type, ai.trx_permit_type, 
			ai.balance_sign, ai.balance
		from %[1]s.rekeningtransaksi rt
			inner join %[1]s.glaccount a on rt.nomor_rekening = a.nomor_rekening
			inner join %[1]s.accountinstance ai on a.accountinstance_id = ai.accountinstance_id
				and rt.kode_valuta = ai.currency_code
				and rt.kode_cabang = ai.branch_code
		where rt.nomor_rekening = $1

-- name: ValidasiRekKodeInterfaceTrx-main
select rt.nomor_rekening, rt.nama_rekening, rt.kode_valuta, rt.kode_cabang, rt.saldo, a.accountinstance_id, rt.kode_account, g.kode_interface
		from %[1]s.rekeningtransaksi rt
			inner join %[1]s.glaccount a on rt.nomor_rekening = a.nomor_rekening
			inner join %[1]s.globalglinterface g on rt.kode_account = g.kode_account
		where rt.kode_cabang = $1
			and g.kode_interface = $2

-- name: ValidasiRekLiabTrx-main
select rt.nomor_rekening, rt.kode_valuta, rt.kode_cabang, rt.saldo, a.accountinstance_id, 
			i.kode_account, a.accountinstance_id, 
			rt.saldo - coalesce(rl.saldo_ditahan, 0.0) - coalesce(p.saldo_minimum, 0.0) - coalesce(rl.saldo_float, 0.0) as saldo_efektif,
			rt.status_rekening
		from %[1]s.rekeningtransaksi rt
			inner join %[1]s.rekeningliabilitas rl on rt.nomor_rekening = rl.nomor_rekening
			inner join %[1]s.produk p on rl.kode_produk = p.kode_produk
			left join %[1]s.glinterface i on p.kode_produk = i.kode_produk and i.kode_interface in ('Saldo_Plus','glnomi')
			left join %[1]s.accountinstance a on rt.kode_valuta = a.currency_code and rt.kode_cabang = a.branch_code and i.kode_account = a.account_code
		where rt.nomor_rekening = $1

-- name: GetBalanceToday-main
select coalesce(sum(amount_credit - amount_debit) * aci.balance_sign,0.0) as total_today
		from %[1]s.journal j 
			left join %[1]s.journalitem ji on j.journal_no = ji.fl_journal
			left join %[1]s.accountinstance aci on aci.accountinstance_id = ji.accountinstance_id
		where j.journal_date = to_date($1,'YYYY-MM-DD')
			and aci.accountinstance_id = $2
		group by aci.balance_sign

-- name: TrxForUpdate-main
select nomor_rekening
		from %[1]s.rekeningtransaksi
		where nomor_rekening = $1
		for update

-- name: UpdateAcctBlock-main
update %[1]s.rekeningliabilitas
		set 
			is_blokir = $1,
			is_blokir_kredit = $2,
			is_blokir_debet = $3,
			kode_sifat_rekening = $4,
			keterangan_blokir = $5,
			sistem_pemblokir = $6,
			user_blokir = $7
		where nomor_rekening = $8

-- name: InsertHistBlokir-main
insert into %[1]s.histblokir
		(id_histblokir, nomor_rekening, keterangan_blokir, user_input, user_dualcontrol, nomor_nasabah, nama_lengkap, status_blokir, 
			tanggal_ubah, sistem_pemblokir, kode_cabang_input, kode_cabang_pemblokir)
		values( nextval('%[1]s.seq_histblokir'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)

-- name: GetHistblokirAktif-main
select nomor_rekening ,status_blokir ,sistem_pemblokir ,user_input ,to_char(tanggal_ubah, 'yyyy-mm-dd') as tanggal_blokir, keterangan_blokir
		from %[1]s.histblokir h
		where status_blokir = 'T'
			and nomor_rekening = $1

-- name: CreateAlamatAlernatif-main
insert into %[1]s.alamatalternatif
		(alamat_jalan, alamat_rtrw, alamat_kelurahan, alamat_provinsi, 
			alamat_kode_pos, nomor_rekening, alamat_kecamatan, alamat_kota_kabupaten, telepon, inputer, tgl_input, status) 
		values 
		($1, $2, $3, $4, 
			$5, $6, $7, $8, $9, $10, current_timestamp, $11)

-- name: DeleteAlamatAlernatif-main
delete from %[1]s.alamatalternatif 
		where nomor_rekening = $1

-- name: UpdateAlamatAlernatif-main
update %[1]s.alamatalternatif 
		set telepon = $1, alamat_jalan = $2, 
			alamat_rtrw = $3, alamat_kelurahan = $4, alamat_kecamatan = $5,
			alamat_kota_kabupaten = $6, alamat_provinsi = $7, 
			alamat_kode_pos = $8 
		where nomor_rekening = $9

-- name: ValidasiRegSweep-main
select ras.nomor_rekening_tujuan, rl.status_register_layanan
		from %[1]s.registerautosweep ras
			inner join %[1]s.registerlayanan rl on rl.id_register = ras.id_register
		where nomor_rekening_tujuan = $1

-- name: GetAutotrfList-main
select si.debitaccountno as rekening_debit, d.nama_rekening as nama_rekening_debit, si.creditaccountno as rekening_kredit, 
			k.nama_rekening as nama_rekening_kredit, to_char(rl.tanggal_registrasi, 'YYYY-MM-DD') as tgl_registrasi, 
			to_char(rt.tgl_proses, 'YYYY-MM-DD') as tgl_proses, to_char(rt.tgl_kadaluarsa, 'YYYY-MM-DD') as tgl_kadaluarsa, 
			si.description as keterangan, to_char(rt.tgl_proses_berikutnya, 'YYYY-MM-DD') as tgl_proses_berikutnya, 
			si.instructiontypecode as kode_tipe_instruksi, rl.id_register, rl.status_register_layanan, 
			case when rl.status_register_layanan = 'A' then 'Aktif' else 'Tidak Aktif' end as status_register_layanan_desc, rt.nominal 
		from %[1]s.standinginstruction si
			inner join %[1]s.recurrenttransaction rt on si.instructionid = rt.instructionid
			inner join %[1]s.registerrt rr on rt.recid = rr.recid
			inner join %[1]s.registerlayanan rl on rr.id_register = rl.id_register
			left join %[1]s.rekeningtransaksi d on d.nomor_rekening = si.debitaccountno
			left join %[1]s.rekeningtransaksi k on k.nomor_rekening = si.creditaccountno
		where rl.id_layanan = $1
			and si.debitaccountno = $2
		order by rl.tanggal_registrasi

-- name: GetAutosaveList-main
select rl.id_register ,to_char(rl.tanggal_registrasi, 'yyyy-mm-dd') as tgl_registrasi ,rl.nomor_rekening_layanan 
			,d.nama_rekening as nama_rekening_layanan ,r.nomor_rekening_tujuan ,k.nama_rekening as nama_rekening_tujuan ,rl.jenis_register_layanan
			,rl.status_register_layanan ,case when rl.status_register_layanan = 'A' then 'Aktif' else 'Tidak Aktif' end as status_register_layanan_desc
			,si.description ,r.saldo_minimal ,r.saldo_maksimal, rl.biaya_layanan
		from %[1]s.registerautosweep r 
			inner join %[1]s.registerlayanan rl on rl.id_register = r.id_register
			inner join %[1]s.standinginstruction si on r.instructionid = si.instructionid
			left join %[1]s.rekeningtransaksi d on d.nomor_rekening = rl.nomor_rekening_layanan
			left join %[1]s.rekeningtransaksi k on (k.nomor_rekening = r.nomor_rekening_tujuan)
		where (rl.nomor_rekening_layanan = $1 or r.nomor_rekening_tujuan = $2)
			and rl.id_layanan = $3
			and rl.jenis_register_layanan = $4

-- name: CreateRegAutosave-main
insert into %[1]s.registerautosweep 
		(id_register, instructionid, instructionid_swap, saldo_minimal, saldo_maksimal, cek_saldo_minimal, 
			cek_saldo_maksimal, nomor_rekening_tujuan) 
		values 
		($1, $2, $3, $4, $5, $6, $7, $8)

-- name: GetKodeBankTrfList-main
select m.kode_member ,nama_member ,nama_singkat, status_peserta_rtgs ,status_peserta_skn, ma.account_no, ma.account_name
		from %[1]s.memberrtgs m 
			inner join %[1]s.memberrtgsaccount ma on m.kode_member = ma.kode_member
		where %[2]s
		order by nama_member

-- name: GetKodeBankTrfList-filter-Default
m.status = 'A' and ma.currency_code = 'IDR'

-- name: GetKodeBankTrfList-filter-JenisMemberSKN
and m.status_peserta_skn = $1

-- name: GetKodeBankTrfList-filter-JenisMemberRTGS
and m.status_peserta_rtgs = $1

-- name: GetKliringKotaList-main
select sandi_kota, sandi_propinsi, nama_kota, keterangan
		from %[1]s.kliringkota k
		order by sandi_kota

-- name: GetKliringKotaSknList-main
select k.sandi_kota ,k.sandi_propinsi ,k.nama_kota ,k.keterangan
		from %[1]s.kliringkota k
			inner join %[1]s.kliringwilayahkota wk on wk.sandi_kota = k.sandi_kota 
			inner join %[1]s.kliringbankwilayah bw on bw.sandi_wilayah = wk.sandi_wilayah 
		where bw.kode_bank = $1
		order by k.sandi_kota

-- name: SelectRtgsMsgCodeList-main
SELECT * FROM %v.TIPEMESSAGE_RTGS tr

-- name: SelectRtgsTrxCodeList-main
SELECT * FROM %v.KODETX_RTGS kr 

-- name: SelectRtgsPriorCodeList-main
select kode_prioritas, nama_prioritas
		from %v.kodeprioritas_rtgs kr
		order by kode_prioritas

-- name: GetBlobDataLength-main
select length(data) blob_data_length
		from %[1]s.blobdata
		where id_data = $1

-- name: GetBlobData-main
select substr(data, 3500, $1) blob_data
		from %[1]s.blobdata
		where id_data = $2

-- name: GetBlobdataId-main
select  nextval('%[1]s.seq_blobdata')
		

-- name: InsertIdBlobData-main
insert into %[1]s.blobdata (id_data)
	values ($1)

-- name: InsertBlobData1-main
update %[1]s.blobdata set data = $1
    where id_data = $2

-- name: InsertBlobData2-main
update %[1]s.blobdata set data = data || $1
        where id_data = $2

-- name: InsertHoldDana-main
insert into %v.holddanarekening
		(id_hold_dana, status, sistem_hold, nomor_rekening, nominal_hold, alasan_hold, tanggal_hold, tanggal_kadaluarsa, is_regulatory, 
			user_hold, user_otorisasi, kode_cabang_hold)
		values (nextval('%v.seq_holddanarekening'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)

-- name: InsertHoldDanaFinfFund-main
insert into %v.holddanarekening
		(id_hold_dana, status, sistem_hold, nomor_rekening, nominal_hold, alasan_hold, tanggal_hold, tanggal_kadaluarsa, is_regulatory, 
			user_hold, user_otorisasi, kode_cabang_hold, nomor_referensi)
		values (nextval('%v.seq_holddanarekening'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)

-- name: GetBlokirDana-main
select h.id_hold_dana, h.nominal_hold, h.alasan_hold, h.nomor_rekening, h.status 
		from %v.holddanarekening h
		where h.id_hold_dana = $1
			and h.nomor_rekening = $2

-- name: UpdateStatusHold-main
update %v.holddanarekening h
		set h.status = $1, h.tanggal_non_aktif = $2
		where h.id_hold_dana = $3
			and h.nomor_rekening = $4

-- name: UpdateSaldoDitahan-main
update  %v.rekeningliabilitas rl 
		set saldo_ditahan = (
			select sum(coalesce(h.nominal_hold,0.0)) saldo_hold
			from %v.holddanarekening h
			where h.status = 'A'
				and h.nomor_rekening = $1
		)
		where rl.nomor_rekening = $1

-- name: GetListHoldDana-main
select 
			h.id_hold_dana, h.nominal_hold, h.alasan_hold, 
			to_char(h.tanggal_hold, 'YYYY-MM-DD') as tanggal_hold, 
			to_char(h.tanggal_kadaluarsa, 'YYYY-MM-DD') as tanggal_kadaluarsa,
			h.user_hold, h.nomor_rekening, h.status, 
			case 
				when h.status ='A' then 'AKTIF' 
				when h.status ='N' then 'NON AKTIF'
				else '' 
			end as status_desc, 
			to_char(h.tanggal_non_aktif, 'YYYY-MM-DD') as tanggal_non_aktif, 
			h.sistem_hold, h.user_otorisasi, h.kode_cabang_hold, h.nomor_referensi, 
			coalesce(is_regulatory,'F') as is_regulatory
		from %[1]s.holddanarekening h
		where h.nomor_rekening = $1
		order by tanggal_hold desc 

-- name: GetListHoldDana-filter-Status
and h.status = $2

-- name: GetHoldDanaFinFund-main
SELECT * FROM %[1]s.holddanarekening
		WHERE nomor_referensi = $1
		AND STATUS = 'A'

-- name: UpdateHoldDanaFinFund-main
UPDATE %[1]s.holddanarekening
		SET 
			TANGGAL_HOLD = $1, 
			TANGGAL_KADALUARSA = TO_DATE($2, '%[2]s'), 
			NOMINAL_HOLD = $3, 
			USER_HOLD = $4
		WHERE
			NOMOR_REFERENSI =  $5

-- name: UpdateHoldDanaNonaktifFinFund-main
UPDATE %[1]s.holddanarekening
		SET 
			TANGGAL_NON_AKTIF = $1
			, Status = 'N'
		WHERE
		NOMOR_REFERENSI =  $2

-- name: GetBukuWarkat-main
select nomor_seri_awal, nomor_seri_akhir, lembar_terpakai, lembar_tersedia, id_bukuwarkat, jenis_warkat, is_otorisasi, tanggal_otorisasi, 
			user_otorisasi, user_input, nomor_rekening, status_buku, tanggal_input, tanggal_aktivasi, user_otorisasi_aktivasi, tanggal_otorisasi_aktivasi, user_aktivasi 
		from %[1]s.bukuwarkat
		where id_bukuwarkat = $1

-- name: UpdateAktivasiWarkat-main
update %[1]s.bukuwarkat
		set user_aktivasi = $1, user_otorisasi_aktivasi = $2,
			status_buku = $3, tanggal_aktivasi = current_timestamp, tanggal_otorisasi_aktivasi = current_timestamp
		where id_bukuwarkat = $4

-- name: UpdateLembarTerpakaiCair-main
update %[1]s.bukuwarkat
		set lembar_terpakai = lembar_terpakai + 1
		where id_bukuwarkat = $1

-- name: UpdateLembarTerpakaiCairReverse-main
update %[1]s.bukuwarkat
		set lembar_terpakai = lembar_terpakai - 1
		where id_bukuwarkat = $1

-- name: GetCabangDetail-main
select kode_cabang, nama_cabang, tipe_cabang, status_akses, case when status_akses ='O' then 'BUKA' else 'TUTUP' end as status_akses_desc, 
			kode_wilayah_kliring , kode_wilayah , is_kordinator_kliring , status_cabang , status_aktif 
		from %[1]s.cabang c
		where c.kode_cabang = $1

-- name: GetCardListByAccNo-main
select savingaccountno ,cardstate ,cardno
		from %[1]s.card c 
		where c.savingaccountno = $1
			and c.cardstate in (1, 5, 6, 7)

-- name: GetCardListByAccountNo-main
select cardid ,cardno ,cardstate ,branchcode ,customerno ,customername ,savingaccountno ,currentaccountno
		from %[1]s.card c 
		where savingaccountno = $1
			or currentaccountno = $2

-- name: GetCardByCardId-main
select savingaccountno ,cardstate ,cardno
		from %[1]s.card c 
		where c.cardid = $1

-- name: GetCardPinByCardId-main
select cardid, pinoffset, pinenc
		from %[1]s.cardpin c 
		where c.cardid = $1

-- name: GetCardByAccountNo-main
select cardid ,cardno ,branchcode ,customerno ,customername ,savingaccountno ,currentaccountno
		from %[1]s.card c 
		where savingaccountno = $1
			or currentaccountno = $2

-- name: GetIdRequest-main
select nextval('%[1]s.seq_cardcustomerrequest') as id_request 

-- name: InsertCardCustomerRequest-main
insert into %[1]s.cardcustomerrequest
		(id_request, request_state, request_type, input_term, auth_term, input_date, auth_date, input_user, auth_user, cardid, kode_cabang, product_code)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)

-- name: InsertCardAccountLinkRequest-main
insert into %[1]s.cardaccountlinkrequest
		(id_request, cardid, account_no, nama_nasabah, nomor_nasabah, product_code, is_removal, unlink_account)
		values($1, $2, $3, $4, $5, $6, $7, $8)

-- name: InsertCardpin-main
insert into %[1]s.cardpin
		(cardid, pinoffset)
		values($1, $2)

-- name: HistoryCardStatus-main
insert into %[1]s.historycardstatus
		(id_historycard, cardid, customername, customerno, cardstate_before, cardstate_after, changedate, description, user_change)
		values( nextval('%[1]s.seq_historycardstatus'), $1, $2, $3, $4, $5, current_timestamp, $6, $7)

-- name: UpdateCard-main
update %[1]s.card
		set CUSTOMERNAME = $1, CUSTOMERNO = $2, SAVINGACCOUNTNO = $3, CARDSTATE = $4, 
			UPDATE_DATE = $5, OTOR_UPDATE = $6, PRODUCT_CODE = $7, BRANCHCODE = $8
		where cardid = $9

-- name: UpdateCardState-main
update %[1]s.card
		set cardstate = $1
		where cardid = $2

-- name: InsertCardBlockRequest-main
insert into %[1]s.cardblockrequest
		(id_request, block_reason, reference_no, is_block)
		values($1, $2, $3, $4)

-- name: ValidasiPostChq-main
select b.jenis_warkat ,b.nomor_seri_awal ,b.nomor_rekening ,b.status_buku 
		from %[1]s.bukuwarkat b
		where b.jenis_warkat = $1
			and b.nomor_seri_awal = $2

-- name: GetChequeID-main
select nextval('%[1]s.seq_bukuwarkat') as id_bukuwarkat 

-- name: GetChequeList-main
select b.id_bukuwarkat, b.nomor_rekening,
			case 
                when b.jenis_warkat = 'B' then 'Bilyet Giro'
                when b.jenis_warkat = 'C' then 'Cek'
            end as jenis_warkat,
            b.is_otorisasi, to_char(b.tanggal_otorisasi, 'yyyy-mm-dd') as tanggal_otorisasi, to_char(b.tanggal_input, 'yyyy-mm-dd') as tanggal_input,
			to_char(b.tanggal_aktivasi, 'yyyy-mm-dd') as tanggal_aktivasi, to_char(b.tanggal_otorisasi_aktivasi, 'yyyy-mm-dd') as tanggal_otorisasi_aktivasi,
			b.lembar_tersedia, b.nomor_seri_awal, b.nomor_seri_akhir, b.lembar_terpakai, b.user_input, b.user_otorisasi,
			case 
                when b.status_buku = 'A' then 'Aktif'
                when b.status_buku = 'N' then 'Belum Aktivasi'
                else ''
            end as status_buku
        from %[1]s.bukuwarkat b 
		%[2]s
        order by b.id_bukuwarkat desc

-- name: GetChequeList-filter-Default
where nomor_rekening = $1

-- name: GetChequeList-filter-Keyword
and (upper(b.nomor_seri_awal) like $2 
	or upper(b.nomor_seri_akhir) like $2)

-- name: GetChequeList-filter-StatusBuku
and (upper(b.status_buku) = $3)

-- name: GetChequeDetail-main
select b.id_bukuwarkat, b.nomor_rekening, n.nomor_seri, 
			case
				when n.status_warkat = 'A' then 'AKTIF'
				when n.status_warkat = 'B' then 'DIBATALKAN'
				when n.status_warkat = 'C' then 'CAIR'
				when n.status_warkat = 'H' then 'HILANG'
				when n.status_warkat = 'P' then 'PROSES'
				when n.status_warkat = 'S' then 'BLOKIR'
				else ''
			end as status_warkat_desc, 
			n.used_counter, n.status_warkat, n.jenis_warkat, 
			case
				when n.jenis_warkat = 'B' then 'Bilyet Giro'
				when n.jenis_warkat = 'C' then 'Cek'
				else ''
			end as jenis_warkat_desc, 
			to_char(n.tanggal_pakai_warkat, 'YYYY-MM-DD') as tanggal_pakai_warkat,
			to_char(n.tanggal_perubahan_status, 'YYYY-MM-DD') as tanggal_perubahan_status,
			n.user_pengubah, b.user_input, 
			to_char(b.tanggal_aktivasi, 'YYYY-MM-DD') as tanggal_aktivasi
		from %[1]s.bukuwarkat b 
			inner join %[1]s.notadebetinternal n on n.id_bukuwarkat = b.id_bukuwarkat 
		where b.id_bukuwarkat = $1
		order by n.nomor_seri

-- name: GetChequeBySerialNum-main
select b.id_bukuwarkat, n.nomor_seri, n.status_warkat, n.jenis_warkat, 
			case
				when n.jenis_warkat = 'B' then 'Bilyet Giro'
				when n.jenis_warkat = 'C' then 'Cek'
				else ''
			end as jenis_warkat_desc, 
			b.status_buku, b.nomor_rekening, ev.enum_description as status_warkat_desc, rt.nama_rekening, rt.kode_cabang, rt.kode_valuta, rl.kode_produk, p.nama_produk, c.nama_cabang 
		from %[1]s.bukuwarkat b 
			inner join %[1]s.notadebetinternal n on n.id_bukuwarkat = b.id_bukuwarkat
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = b.nomor_rekening
			inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rt.nomor_rekening 
			left join %[1]s.produk p on p.kode_produk = rl.kode_produk
			left join %[2]s.cabang c on rt.kode_cabang = c.kode_cabang
			left join %[1]s.enum_varchar ev on ev.enum_name ='eStatusWarkat' and ev.enum_value = n.status_warkat 
		where n.nomor_seri = $1

-- name: PostChqBook-main
insert into %[1]s.bukuwarkat
		(id_bukuwarkat, nomor_seri_awal, nomor_seri_akhir, lembar_terpakai, lembar_tersedia, jenis_warkat,  is_otorisasi, tanggal_otorisasi, 
			user_otorisasi, user_input, nomor_rekening, status_buku, tanggal_input) 
		values 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

-- name: PostNotaDebetInternal-main
insert into %[1]s.notadebetinternal 
		(id_notadebetinternal, jenis_warkat, id_bukuwarkat, nomor_seri, status_warkat, tanggal_perubahan_status, user_pengubah, nomor_rekening, used_counter) 
		values ( nextval('%[1]s.seq_notadebetinternal'), $1, $2, $3, $4, current_timestamp, $5, $6, $7)

-- name: GetLbvData-main
select 
		ai.accountinstance_id
		, a.account_code
		, a.account_name
		, ai.branch_code
		, ai.currency_code
		, dbr.balance_nominative
		, dbl.balancecumulative
		, dbr.balance_nominative - dbl.balancecumulative as selisih
	from %[1]v.account a
		, accountinstance ai
		left outer join 
		(select accountinstance_id, sum(balance) as balance_nominative
			from
				DailyBalanceRekening db1
				, ( select nomor_rekening, max(balance_date) as balance_date
					from DailyBalanceRekening
					where balance_date <= TO_DATE($1, 'YYYY-MM-DD')
						and balance_field = 'saldo'
					group by nomor_rekening
					) db2
			where db1.Nomor_Rekening = db2.Nomor_Rekening
					and db1.balance_field = 'saldo'
					and db1.balance_date = db2.balance_date
			group by accountinstance_id
		) dbr
		on ai.accountinstance_id = dbr.accountinstance_id
		left outer join
			(select
				d1.accountinstance_id, d1.balancecumulative
			from
				DailyBalance d1,
				(select accountinstance_id, max(datevalue) as maxdate
				from DailyBalance where datevalue <= TO_DATE($1, 'YYYY-MM-DD')
				group by accountinstance_id) d2
			where
				d1.accountinstance_id = d2.accountinstance_id and
				d1.datevalue = d2.maxdate) dbl
			on ai.accountinstance_id = dbl.accountinstance_id
	where
		a.account_code = ai.account_code
		and ( ai.Branch_Code = $2 or  'ALL' = '%[2]v')  
		and ( ai.Currency_code = $3 or 'ALL' = '%[3]v')
		and 
		( 
			exists (
				select 1 
				from produk p 
					, glinterface g 
				where  p.kode_produk = g.kode_produk
					and kode_interface in ('Saldo_Plus', 'glnomi') 
					and kode_account = a.account_code
					and p.jenis_produk in ('T', 'G', 'D', 'R')
			)
			
		)
	order by ai.branch_code, a.account_code, ai.currency_code

-- name: GetSaldoRataRata-main
select nomor_rekening ,balance
		from %[1]s.dailybalancerekening d
		where is_latest = 'T'
			and balance_field = 'saldo_rata_bulanan'
			and nomor_rekening = $1

-- name: GetDepList-main
select rl.nomor_nasabah, rt.nomor_rekening, rt.nama_rekening, rt.kode_cabang
			, p.kode_produk, p.nama_produk,
			case
				when rt.status_rekening = 3 then 0
				else rt.saldo - p.saldo_minimum - rl.saldo_ditahan 
			end as saldo_efektif,
			rt.status_rekening,
			case 
				when rt.status_rekening = 1 then 'Aktif'
				when rt.status_rekening = 2 then 'Dormant'
				when rt.status_rekening = 3 then 'Tutup'
				else ''
			end as status_rekening_desc,
			case 
				when rl.is_blokir = 'T' then 'blokir_debet_credit'
				when rl.is_blokir_debet = 'T' then 'blokir_debet'
				when rl.is_blokir_kredit = 'T' then 'blokir_credit'
				else ''
			end jenis_blokir,
			rl.is_blokir_debet, rl.is_blokir_kredit, rl.is_sedang_ditutup
		from %[1]s.deposito d
			inner join %[1]s.rekeningliabilitas rl on d.nomor_rekening = rl.nomor_rekening
			inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
			inner join %[1]s.produk p on p.kode_produk = rl.kode_produk
		where rl.jenis_rekening_liabilitas in ('D') and rt.kode_jenis in ('DEP')
		%[2]s
		OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY

-- name: GetDepList-filter-KodeCabang
and rt.kode_cabang = $3

-- name: GetDepList-filter-StatusRekening
and rt.status_rekening = $4

-- name: GetDepList-filter-Keyword
and ((upper(rt.nama_rekening) like $5 or upper(rt.nomor_rekening) like $5))

-- name: GetDepList-filter-NomorNasabah
and rl.nomor_nasabah = $6

-- name: GetDetailDep-main
select rl.nomor_rekening, rl.nomor_nasabah, rt.kode_valuta, rt.kode_cabang, c.nama_cabang, rt.user_input, 
			rt.nama_rekening, rt.nama_valuta, rt.keterangan, rl.is_bagi_hasil_khusus, rl.is_tiering_nisbahbonus, d.jenis_alamat_bilyet as alamat_bilyet,
			rl.kode_marketing_current, rl.kode_marketing_pertama, rl.kode_marketing_referensi, rl.kode_produk, rl.kode_program, 
			rl.kode_sumber_dana, rl.status_kelengkapan, rl.tarif_pajak, rl.user_otorisasi, round(rl.nisbah_spesial, 2) as nisbah_spesial, round(rl.nisbah_bagi_hasil, 2) as nisbah_bagi_hasil, 
			rl.nisbah_dasar, rl.kode_tujuan_rekening, rl.nomor_rekening_disposisi, rt.saldo, rt.saldo as saldo_deposito,
			to_char(rl.tanggal_buka, 'YYYY-MM-DD') as tanggal_buka, rl.status_restriksi, rl.kode_status_retriksi, rl.penerima_transfer_bagi_hasil, 
			rl.disposisi_bagi_hasil, rl.kode_member_nominal, rl.kode_member_nominal_desc, rl.kode_member_bagihasil_desc, rl.baghas_alamat_transfer, 
			rl.baghas_sandikota_transfer, rl.baghas_sandikota_transfer_desc, rl.kode_member_bagihasil, rl.alamat_kirim, rl.is_join_account,
			rl.is_dapat_bagi_hasil, rl.is_blokir, to_char(d.tanggal_jatuh_tempo_terakhir, 'YYYY-MM-DD') as tanggal_jatuh_tempo_terakhir, 
			to_char(d.tanggal_jatuh_tempo_berikutnya, 'YYYY-MM-DD') as tanggal_jatuh_tempo_berikutnya, d.nomor_bilyet, d.jenis_deposito, 
			d.rekening_disposisi, d.disposisi_nominal, d.penerima_transfer_disposisi, p.nama_produk, 
			round(rl.persentase_zakat_bagi_hasil, 2) as persentase_zakat_bagi_hasil, rl.kode_multicifrelation, d.is_sudah_disetor, d.is_sedang_dicairkan, rl.saldo_accrual_bagihasil, rt.status_rekening , 
			case 
				when rt.status_rekening = 1 then 'Aktif'
				when rt.status_rekening = 2 then 'Dormant'
				when rt.status_rekening = 3 then 'Tutup'
			end as status_rekening_desc,
			rl.saldo_ditahan , d.cadangan_bagihasil_kapitalisir, d.periode_perjanjian ,d.masa_perjanjian, to_char(rl.tanggal_bagi_hasil_terakhir, 'yyyy-mm-dd') as tanggal_bagi_hasil_terakhir, 
			to_char(rl.tanggal_bagi_hasil_berikutnya, 'yyyy-mm-dd') as tanggal_bagi_hasil_berikutnya, to_char(rl.tanggal_blokir, 'yyyy-mm-dd') as tanggal_blokir,
			rl.keterangan_blokir, rl.sistem_pemblokir, rl.user_blokir, rl.nisbah_dasar 
		from %[1]s.deposito d
			inner join %[1]s.rekeningliabilitas rl on d.nomor_rekening = rl.nomor_rekening
			inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
			inner join %[1]s.produk p on p.kode_produk = rl.kode_produk
			inner join %[2]s.cabang c on c.kode_cabang = rt.kode_cabang
		where rl.jenis_rekening_liabilitas = 'D'
			and rt.kode_jenis = 'DEP' 
			and rl.nomor_rekening = $1

-- name: CreateDeposito-main
insert into %[1]s.deposito(
			nomor_rekening, tanggal_jatuh_tempo_terakhir, cadangan_bagihasil_kapitalisir, masa_perjanjian,
			periode_perjanjian, nomor_bilyet, tanggal_jatuh_tempo_berikutnya, disposisi_nominal,
			rekening_disposisi, penerima_transfer_disposisi, is_aro, is_sudah_disetor ,is_sedang_dicairkan ,
			is_sudah_cetak_bilyet, jenis_alamat_bilyet
		) 
		values (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 , $13, $14, $15 
		)

-- name: GetDepositoByNoBilyet-main
select d.nomor_rekening, d.nomor_bilyet 
		from %[1]s.deposito d 
		where d.nomor_bilyet = $1

-- name: GetDepositoBilyetPrint-main
select rt.nomor_rekening, rt.nama_rekening, rt.kode_cabang, c.nama_cabang, coalesce(rl.alamat_kirim, 'N') alamat_kirim, n.jenis_nasabah,
			n.nomor_nasabah, d.is_sudah_cetak_bilyet, d.is_sudah_disetor, to_char(rl.tanggal_buka, 'yyyy-mm-dd') as tanggal_buka,
			to_char(d.tanggal_jatuh_tempo_berikutnya, 'yyyy-mm-dd') as tanggal_jatuh_tempo, rl.nisbah_bagi_hasil as bagi_hasil_deposan, 
			100 - rl.nisbah_bagi_hasil as bagi_hasil_bank, d.disposisi_nominal, rl.disposisi_bagi_hasil, d.is_aro, d.nomor_bilyet, 
			d.rekening_disposisi as nomor_rekening_pokok, d.penerima_transfer_disposisi as nama_rekening_pokok, 
			d.kode_bank_nominal as kode_bank_pokok, rl.nomor_rekening_disposisi as nomor_rekening_bagihasil, 
			rl.penerima_transfer_bagi_hasil as nama_rekening_bagihasil, d.masa_perjanjian, d.periode_perjanjian, rt.saldo
		from %[1]s.rekeningtransaksi rt 
			inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rt.nomor_rekening
			inner join %[1]s.deposito d on d.nomor_rekening = rl.nomor_rekening
			inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah
			inner join %[2]s.cabang c on c.kode_cabang = rt.kode_cabang
		where rt.nomor_rekening = $1

-- name: UpdateBilyetPrintStatus-main
UPDATE %[1]s.deposito
		SET is_sudah_cetak_bilyet = $1 
		WHERE nomor_rekening = $2

-- name: CreateHistBilyetDepoPrint-main
INSERT INTO %[1]s.histcetakbilyetdepo (ID_HISTORI, NOMOR_REKENING, USER_CETAK, USER_OVERRIDE, TANGGAL_CETAK)
		VALUES ( nextval('%[1]s.seq_histcetakbilyetdepo'), $1, $2, $3, SYSDATE)

-- name: GetInterbankDepositList-main
select 
		n.Id_NotaDebitOutward 
		, n.Tanggal_Input 
		, n.Tanggal_Efektif 
		, n.Nomor_Rekening 
		, n.Kode_Bank 
		, n.Nomor_Referensi 
		, n.Nomor_Rekening_Nota 
		, n.Nama_Pemilik_Nota 
		, n.Nominal 
		, n.Kurs_Transaksi 
		, n.Nomor_Aplikasi 
		, n.Inputer 
		, n.Jenis_Warkat
		, c.Kode_Cabang
		, c.Nama_Cabang
		, r.NAMA_REKENING
		, m.NAMA_MEMBER as nama_bank                    
	from 
		%[1]v.NotaDebitOutward n
		inner join %[2]v.Cabang c on c.KODE_CABANG = n.KODE_CABANG 
		inner join %[1]v.REKENINGTRANSAKSI r on r.NOMOR_REKENING = n.NOMOR_REKENING 
		inner join %[1]v.memberrtgs m on n.kode_bank = m.kode_member
	where
		n.Tanggal_Efektif <=  to_date($1, 'yyyy-mm-dd')
		and n.Status = 'S' 
		and n.Is_Otorisasi = 'T' 
		and n.Kode_Cabang = $2 
	order by Tanggal_Efektif, n.Nomor_Rekening

-- name: ModAcctDepRT-main
update %[1]s.rekeningtransaksi 
		set nama_rekening = $1, keterangan = $2, 
			tanggal_perbarui = current_timestamp, tanggal_update = $3, 
			user_perbarui = $4 
		where nomor_rekening= $5

-- name: ModAcctDepRL-main
update %[1]s.rekeningliabilitas 
		set nama_singkat = $1, kode_marketing_referensi = $2, kode_marketing_pertama = $3, 
			kode_marketing_current = $4, kode_program = $5, kode_sumber_dana = $6, 
			kode_tujuan_rekening = $7, status_kelengkapan = $8, is_dapat_bagi_hasil = $9, 
			kode_multicifrelation = $10, nisbah_spesial = $11, persentase_zakat_bagi_hasil = $12, 
			tarif_pajak = $13, disposisi_bagi_hasil = $14, nomor_rekening_disposisi = $15, 
			penerima_transfer_bagi_hasil = $16, status_restriksi = $17, kode_status_retriksi = $18, 
			is_bagi_hasil_khusus = $19, is_tiering_nisbahbonus = $20, is_join_account = $21, 
			is_kena_pajak = $22, is_kena_zakat_bagi_hasil = $23, nisbah_bagi_hasil = $24, 
			nisbah_dasar = $25, user_update = $26
		where nomor_rekening = $27

-- name: ModAcctDepDeposito-main
update %[1]s.deposito 
		set nomor_bilyet = $1, disposisi_nominal = $2, rekening_disposisi = $3, 
			penerima_transfer_disposisi = $4, jenis_alamat_bilyet = $5, is_aro = $6
		where nomor_rekening = $7

-- name: ModDepForclose-main
update %[1]s.deposito 
		set is_sedang_dicairkan = $1
		where nomor_rekening = $2

-- name: InsertInfoPencairanAwalDeposito-main
insert into %[1]s.infopencairanawaldeposito 
		(id_transaksi, nomor_rekening, tanggal_proses, kode_valuta, disposisi_nominal_ori, kode_bank_disposisi_ori, rekening_disposisi_ori, 
			penerima_disposisi_ori, nilai_pencairan, user_input, kode_cabang, saldo_accrual_bagihasil, cadangan_bagihasil_kapitalisir)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

-- name: DeleteInfoPencairanAwalDeposito-main
delete from %[1]s.infopencairanawaldeposito 
		where id_transaksi = $1

-- name: UpdateDepositoPencairanAwalDeposito-main
update %[1]s.deposito 
		set is_aro = $1, disposisi_nominal = $2, rekening_disposisi = $3, 
			penerima_transfer_disposisi = $4, cadangan_bagihasil_kapitalisir = $5
		where nomor_rekening = $6

-- name: GetNotaDebitOutwardDelList-main
select 
		count (ID_NOTADEBITOUTWARD) AS "JUMLAH_TERHAPUS"               
	from 
		%[1]v.NotaDebitOutward n
	where
		n.Id_Group = $1 
	GROUP BY 
		n.ID_GROUP

-- name: DelNotaDebitOutward-main
DELETE FROM %[1]s.NOTADEBITOUTWARD 
		WHERE id_group = $1

-- name: DelNotaDebitOutwardGroup-main
DELETE FROM %[1]s.NOTADEBITOUTWARDGROUP 
		WHERE id_group = $1

-- name: GetDepForClose-main
select rt.nomor_rekening ,rl.nomor_nasabah ,rt.nama_rekening  ,rt.kode_cabang ,rt.kode_valuta ,pd.is_kena_penalty, pd.nominal_penalty
			,to_char(rl.tanggal_buka, 'yyyy-mm-dd') AS tanggal_buka ,to_char(d.tanggal_jatuh_tempo_terakhir, 'yyyy-mm-dd') AS tanggal_jatuh_tempo_terakhir 
			,d.jumlah_aro ,to_char(d.tanggal_jatuh_tempo_berikutnya, 'yyyy-mm-dd') AS tanggal_jatuh_tempo_berikutnya ,rt.saldo ,d.disposisi_nominal 
			,d.rekening_disposisi ,d.penerima_transfer_disposisi, rt.status_rekening, d.is_sudah_disetor, d.is_sedang_dicairkan, d.cadangan_bagihasil_kapitalisir
			,rl.is_kena_pajak ,rl.tarif_pajak, rl.is_kena_zakat_bagi_hasil, rl.persentase_zakat_bagi_hasil, rl.saldo_accrual_bagihasil
		from %[1]s.rekeningtransaksi rt
			inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rt.nomor_rekening 
			inner join %[1]s.deposito d on d.nomor_rekening = rt.nomor_rekening 
			inner join %[1]s.produk p on p.kode_produk = rl.kode_produk 
			inner join %[1]s.produkdeposito pd on pd.kode_produk = p.kode_produk 
		where rt.nomor_rekening = $1

-- name: GetDepWithdrawListToday-main
select t.id_transaksi, t.id_batch_transaksi, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, i.nomor_rekening, 
		rt.nama_rekening, i.nilai_pencairan, i.kode_valuta, t.kode_cabang_transaksi, t.status_otorisasi, 
		case
			when t.status_otorisasi = 0 then 'Belum otorisasi'
			when t.status_otorisasi = 1 then 'Sudah otorisasi'
			when t.status_otorisasi = 2 then 'Sudah direject'
		end as status_otorisasi_desc,
		'T' as is_trx_today
	from %[1]s.transaksi t  
		inner join %[1]s.infopencairanawaldeposito i on t.id_transaksi = i.id_transaksi 
		inner join %[1]s.rekeningtransaksi rt on i.nomor_rekening = rt.nomor_rekening 
	where t.kode_cabang_transaksi = $1
		and t.status_otorisasi = 1
		and coalesce(t.is_reversed, 'F') != 'T'
		%[2]s
	order by tanggal_transaksi asc

-- name: GetDepWithdrawListToday-filter-UserInput
and t.user_input = $2

-- name: GetDepWithdrawListHist-main
select ht.id_transaksi, ht.id_batch_transaksi, to_char(ht.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, i.nomor_rekening, 
			rt.nama_rekening, i.nilai_pencairan, i.kode_valuta, ht.kode_cabang_transaksi, ht.status_otorisasi, 
			case
				when ht.status_otorisasi = 0 then 'Belum otorisasi'
				when ht.status_otorisasi = 1 then 'Sudah otorisasi'
				when ht.status_otorisasi = 2 then 'Sudah direject'
			end as status_otorisasi_desc,
			'F' as is_trx_today
		from %[1]s.histtransaksi ht 
			inner join %[1]s.infopencairanawaldeposito i on ht.id_transaksi = i.id_transaksi
			inner join %[1]s.rekeningtransaksi rt on i.nomor_rekening = rt.nomor_rekening 
		where ht.kode_cabang_transaksi = $1
			and ht.status_otorisasi = 1
			and coalesce(ht.is_reversed, 'F') != 'T'
			and i.tanggal_proses >= to_timestamp($2,'yyyy-mm-dd') 
			and i.tanggal_proses <= to_timestamp($3,'yyyy-mm-dd')   
			%[2]s 
		order by tanggal_transaksi asc

-- name: GetDepWithdrawListHist-filter-UserInput
and ht.user_input = $4

-- name: GetDepWithdrawToday-main
select t.id_transaksi, t.id_batch_transaksi, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, i.nomor_rekening, 
		rt.nama_rekening, i.nilai_pencairan, i.kode_valuta, t.kode_cabang_transaksi, t.status_otorisasi, 
		case
			when t.status_otorisasi = 0 then 'Belum otorisasi'
			when t.status_otorisasi = 1 then 'Sudah otorisasi'
			when t.status_otorisasi = 2 then 'Sudah direject'
		end as status_otorisasi_desc,
		'T' as is_trx_today, t.is_reversed, t.nomor_referensi, t.keterangan, t.jenis_aplikasi, t.kode_valuta_transaksi, t.kurs_manual,
		t.nilai_transaksi, t.biaya, i.disposisi_nominal_ori, i.kode_bank_disposisi_ori, i.rekening_disposisi_ori, i.penerima_disposisi_ori,
		i.saldo_accrual_bagihasil, i.cadangan_bagihasil_kapitalisir
	from %[1]s.transaksi t  
		inner join %[1]s.infopencairanawaldeposito i on t.id_transaksi = i.id_transaksi 
		inner join %[1]s.rekeningtransaksi rt on i.nomor_rekening = rt.nomor_rekening 
	where t.id_transaksi = $1

-- name: GetDepWithdrawHist-main
select ht.id_transaksi, ht.id_batch_transaksi, to_char(ht.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, i.nomor_rekening, 
		rt.nama_rekening, i.nilai_pencairan, i.kode_valuta, ht.kode_cabang_transaksi, ht.status_otorisasi, 
		case
			when ht.status_otorisasi = 0 then 'Belum otorisasi'
			when ht.status_otorisasi = 1 then 'Sudah otorisasi'
			when ht.status_otorisasi = 2 then 'Sudah direject'
		end as status_otorisasi_desc,
		'F' as is_trx_today, ht.is_reversed, ht.nomor_referensi, ht.keterangan, ht.jenis_aplikasi, ht.kode_valuta_transaksi, ht.kurs_manual,
		ht.nilai_transaksi, ht.biaya, i.disposisi_nominal_ori, i.kode_bank_disposisi_ori, i.rekening_disposisi_ori, i.penerima_disposisi_ori,
		i.saldo_accrual_bagihasil, i.cadangan_bagihasil_kapitalisir
	from %[1]s.histtransaksi ht 
		inner join %[1]s.infopencairanawaldeposito i on ht.id_transaksi = i.id_transaksi
		inner join %[1]s.rekeningtransaksi rt on i.nomor_rekening = rt.nomor_rekening 
	where ht.id_transaksi = $1

-- name: GetDepBagihasilList-main
select id_report, nomor_rekening, to_char(tanggal, 'yyyy-mm-dd') as tanggal, round(nominal_total, 2) as nominal_total, round(pajak, 2) as pajak, round(zakat, 2) as zakat,
			disposisi, nomor_disposisi, to_char(periode_awal, 'yyyy-mm-dd') as periode_awal, to_char(periode_akhir, 'yyyy-mm-dd') as periode_akhir,
        	round(nisbah_dasar, 2) as nisbah_dasar, round(nisbah_spesial, 2) as nisbah_spesial
		from %[1]s.bagihasil_deposito p
		where p.nomor_rekening = $1
			and p.tanggal >= to_date($2, 'YYYY-MM-DD HH24::MI::SS')
			and p.tanggal < to_timestamp($3, 'YYYY-MM-DD HH24::MI::SS') + 1

-- name: GetCurrentTimeStamp-main
SELECT 
			TO_CHAR(CURRENT_TIMESTAMP, 'YYYY-MM-DD') AS tanggal,
			TO_CHAR(CURRENT_TIMESTAMP, 'HH24MISS') AS jam
		

-- name: GetTrxDetailDtk-main
select d.id_detil_transaksi, d.id_transaksi, d.jenis_mutasi, d.kode_valuta, d.kode_valuta_rak, to_char(d.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi,
			d.nilai_mutasi, d.keterangan, d.kode_cabang, d.jenis_detil_transaksi, d.nilai_kurs_manual, d.saldo_awal, d.nilai_ekuivalen,
			d.nomor_referensi, d.kode_account, d.nomor_rekening, d.kode_jurnal, d.id_parameter_transaksi, d.flag_proses, d."SEQUENCE",
			d.kode_tx_class
		from %v.detiltransaksi d
		where d.id_transaksi = $1
			and d.jenis_mutasi = $2
		order by d.id_detil_transaksi desc

-- name: GetListDHN-main
select m.no_nasabah as nomor_nasabah,
			m.nama_nasabah,
			m.status_cif,
			esc.enum_description as e_status_cif,
			m.gelar,
			m.jenis_nasabah,
			m.nomor_identitas,
			m.tgl_lahir as tanggal_lahir,
			m.nomor_npwp,
			m.is_dhib,
			case when m.is_dhn='T' then m.is_dhn || ' #' || coalesce(m.ctr_match_dhn,0) else m.is_dhn end as is_dhn,
			m.ctr_match_dhn,
			d.no_dhn,
			d.batas_sanksi,
			d.nomor_ref as nomor_ref_pelapor,
			ejn.enum_description AS e_jenis_nasabah,
			n.kode_cabang_input AS kode_cabang_cif
		from %[1]s.dhib m
			left join %[1]s.dhn d on d.id_dhn = m.id_dhn
			left join %[1]s.enum_varchar ejn on ejn.enum_name='eJenisNasabah' and ejn.enum_value = m.jenis_nasabah
			left join %[1]s.enum_varchar esc on esc.enum_name='eStatusCIF' and esc.enum_value = m.status_cif
			left join %[2]s.nasabah n on n.nomor_nasabah = m.no_nasabah
		where %[3]s
		order by m.nama_nasabah, m.tgl_lahir, d.no_dhn desc
		offset $1 rows fetch next $2 rows only

-- name: GetListDHN-filter-Default
rownum < 30

-- name: GetListDHN-filter-Branch
and n.kode_cabang_input = $3

-- name: GetListDHN-filter-Nasabah
and upper(m.nama_nasabah) LIKE $4

-- name: GetListDHN-filter-TanggalLahir
and to_date(to_char(m.tgl_lahir, 'yyyy-mm-dd'), 'yyyy-mm-dd') = to_date($5, 'yyyy-mm-dd')

-- name: GetListDHN-filter-DHNExpiredTrue
and to_date(to_char(d.batas_sanksi, 'yyyy-mm-dd'), 'yyyy-mm-dd') < to_date($6, 'yyyy-mm-dd')

-- name: GetListDHN-filter-DHNExpiredFalse
and to_date(to_char(d.batas_sanksi, 'yyyy-mm-dd'), 'yyyy-mm-dd') >= to_date($7, 'yyyy-mm-dd')

-- name: GetListDHN-filter-Npwp
and upper(m.nomor_npwp) = $8

-- name: GetListDHN-filter-StatusCif
and upper(m.status_cif) = $9

-- name: GetListDHN-filter-Keyword
and (upper(m.no_nasabah) like $10
	or upper(m.nama_nasabah) like $10
	or upper(m.nomor_identitas) like $10
	or upper(m.nomor_npwp) like $10
	or upper(m.gelar) like $10
	or upper(d.no_dhn) like $10)

-- name: GetListDHN-filter-IsDHIB
and upper(m.is_dhib) = $11

-- name: GetListDHN-filter-IsDHN
and upper(m.is_dhn) = $12

-- name: GetDetilDHIBData-main
select d.id_dhn ,d.no_nasabah ,d.nama_nasabah ,d.jenis_nasabah ,d.nomor_identitas ,to_char(d.tgl_lahir, 'yyyy-mm-dd') as tgl_lahir ,d.no_nasabah ,d.gelar 
			,d.is_bbl ,d.alamat_jalan ,d.alamat_rt ,d.alamat_rw ,d.alamat_kelurahan ,d.alamat_kecamatan ,d.alamat_kota_kabupaten ,d.alamat_provinsi ,d.alamat_kode_pos 
		from %[1]s.dhib d 
		where d.no_nasabah = $1

-- name: GetDetilDHNData-main
select d2.id_dhn ,d2.no_dhn ,d2.no_urut ,d2.nomor_ref ,d2.sandi_bank_asal ,d2.nama_bank_asal ,d2.nama_nasabah ,d2.status_usaha ,d2.jenis_nasabah ,d2.nomor_identitas ,d2.nomor_npwp 
			,to_char(d2.tgl_lahir, 'yyyy-mm-dd') as tgl_lahir, d2.keterangan ,to_char(d2.batas_sanksi, 'yyyy-mm-dd') as batas_sanksi, d2.alamat_jalan ,d2.alamat_rt ,d2.alamat_rw 
			,d2.alamat_kelurahan ,d2.alamat_kecamatan ,d2.alamat_kota_kabupaten ,d2.alamat_provinsi ,d2.alamat_kode_pos 
		from %[1]s.dhn d2
		where d2.id_dhn = $1

-- name: GetDetiltransgroupclass-main
SELECT * FROM %[1]s.detiltransgroupclass 
    WHERE classgroup = $1

-- name: GetAccountDetailed-main
SELECT 
		a.nomor_rekening, a.facility_no, a.targeted_eqv_rate, a.period_count, a.period_unit,
		a.opening_date, a.dropping_due_date, a.contract_number, a.description,a.norek_paymentsrc,
		a.norek_costpayment, a.drop_counter, a.dropping_model, a.loan_amount, a.id_schedule,
		a.write_off_status, b.nama_rekening, b.kode_cabang, b.status_rekening, d.nomor_nasabah,
		c.contracttype_code, c.description as nama_akad, pr.product_code, pr.product_name, pr.currency_code, 
		pr.kode_pof, pr.is_restricted, ag.agentcode, ag.agentname, a.limited_without_bc, a.limited_for_enterprises
	FROM %[2]s.finaccount a
		INNER JOIN %[1]s.rekeningtransaksi b ON a.nomor_rekening=b.nomor_rekening
		INNER JOIN %[1]s.rekeningcustomer d ON a.nomor_rekening=d.nomor_rekening	
		INNER JOIN %[2]s.contracttype c ON a.contracttype_code=c.contracttype_code	
		INNER JOIN %[2]s.finproduct pr ON a.product_code=pr.product_code
		LEFT JOIN %[1]s.fundingagent ag ON a.amcode=ag.agentcode
	WHERE a.nomor_rekening = $1

-- name: GetActiveSchedule-main
SELECT 
		ps.nomor_rekening, ps.id_schedule, ps.completion_status 
	FROM %[1]s.finpaymentschedule ps 
	WHERE nomor_rekening = $1 

-- name: GetCollateralDone-main
SELECT
		CASE fp.use_collateral WHEN 'F' THEN 1
		ELSE (
		SELECT count(*) as coll_count 
		FROM %[1]s.fincollateralaccount f 
		WHERE f.norek_finaccount = $1
		) END as coll_count 
	FROM %[1]s.finaccount fa, %[1]s.finproduct fp 
	WHERE 
	fa.nomor_rekening = $1 
	AND fa.product_code = fp.product_code

-- name: GetNeedFindataComplete-main
SELECT 
		coalesce(need_findata_complete, 'F') as need_findata_complete
	FROM %[1]s.finaccount a
	LEFT JOIN %[1]s.finproduct c on a.product_code=c.product_code
	WHERE a.NOMOR_REKENING= $1 

-- name: GetCustDataCount-main
SELECT 
		coalesce(count(*), 0) as cust_data_count 
	FROM %[1]s.fincustadditionaldata 
	WHERE nomor_nasabah = $1 

-- name: GetAccDataCount-main
SELECT 
		coalesce(count(*), 0) as acc_data_count 
	FROM %[1]s.finaccadditionaldata 
	WHERE nomor_rekening = $1 

-- name: UpdateDroppingFinAccount-main
UPDATE %[1]s.finaccount SET
			drop_counter = $1,
			dropping_amount = $2,
			norek_disbursaldest = $3,
			overall_col_level = $4,
			performance_col_level = $5,
			prospect_col_level = $6,
			repayment_col_level = $7,
			so_col_level = $8,
			manual_col_level = $9,
			coll_deduct_recalc = $10,
			col_eval_required = $11,
			dropping_date = $12,
			due_date = $13,
			predicted_margin_amount = $14
		WHERE nomor_rekening = $15

-- name: UpdateFinAccount-main
UPDATE %[1]s.finaccount SET
			id_schedule = $1
			, period_count = $2
			, due_date = $3
		WHERE nomor_rekening = $4

-- name: CreateCollateralAsset-main
INSERT INTO %[1]s.fincollateralasset (
			NOMOR_REKENING, LAST_VALUATION_DATE, LAST_APPRAISAL_TYPE, LAST_APPRAISAL_NAME,
			DOCUMENT_NO, IS_VALUATED, VALUATION, LIQUIDATION_VALUE, PHOTO_ID, DASAR_NILAI,
			PCT_USE, VALUE_USE, DOC_STORAGE, ASSET_TYPE, NOMOR_NASABAH, AGENT_CODE,
			COLASSET_STATUS_UPDATE, REF_PERINGKAT_SB, REF_JENIS_PENGIKATAN, REF_DATI2,
			PEMILIK_AGUNAN, BUKTI_KEPEMILIKAN, ALAMAT_AGUNAN, PARIPASU, ASURANSI,
			RESERVE_DEDUCT_VALUE, IS_CASH_COLLATERAL, ALAMAT_JALAN1, ALAMAT_JALAN2,
			ALAMAT_KELURAHAN, ALAMAT_KECAMATAN, ALAMAT_KODEPOS, REF_PROPINSI, GROUP_CODE,
			GROUP_ASSET_TYPE, RECORD_DATE, EXPIRE_DATE, MARKET_VALUE, NOTES_AMOUNT,
			NOTES_CURRENCY_CODE, COL_DOC_STATUS, COL_DOC_NO, SECURITIESTYPE, SECURITIES_ISSUED_DATE,
			SECURITIES_NO, SECURITIES_OWNER, LOKASIJAMINAN, TGL_DOKUMEN, JTPEMBAHARUAN,
			TGLJTPEMBAHARUAN, RELEASE_DATE, REF_DATI2_JAMINAN, LOKASI_JAMINAN, RESERVE_DEDUCT_RATIO,
			ID_AGUNAN, LAST_UPDATE, KODE_CABANG, JNS_PRODUK, JENIS_AGUNAN
		)
		VALUES (
			$1, TO_DATE($2, '%[2]s'), $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34,
			$35, TO_DATE($36,'%[2]s'), TO_DATE($37, '%[2]s'), $38, $39, $40, $41, $42, $43, TO_DATE($44, '%[2]s'),
			$45, $46, $47, TO_DATE($48, '%[2]s'), $49, TO_DATE($50, '%[2]s'), TO_DATE($51, '%[2]s'), $52, $53, $54,
			$55, $56, $57, $58, $59
		)

-- name: GetDetailCollateralAsset-main
SELECT
			NOMOR_REKENING, ASSET_TYPE, GROUP_ASSET_TYPE, 
			DASAR_NILAI, COLASSET_STATUS_UPDATE, LAST_UPDATE, 
			NOMOR_NASABAH, KODE_CABANG, JENIS_AGUNAN, 
			PEMILIK_AGUNAN, ALAMAT_JALAN1, ALAMAT_JALAN2, 
			ALAMAT_KELURAHAN, ALAMAT_KECAMATAN, ALAMAT_KODEPOS,
			REF_DATI2, BUKTI_KEPEMILIKAN, ASURANSI, 
			JNS_PRODUK, IS_VALUATED, LAST_VALUATION_DATE, 
			LAST_APPRAISAL_TYPE, AGENT_CODE, LAST_APPRAISAL_NAME, 
			VALUATION, MARKET_VALUE, LIQUIDATION_VALUE, 
			DOCUMENT_NO, COL_DOC_NO, COL_DOC_STATUS, 
			EXPIRE_DATE, DOC_STORAGE
		FROM %[1]s.FINCOLLATERALASSET
		WHERE NOMOR_REKENING = $1

-- name: GetNextColAccountNumber1-main
SELECT kode_cabang_input as branch_code
		FROM %[1]s.nasabah
		WHERE nomor_nasabah = $1

-- name: GetNextColAccountNumber2-main
SELECT
			CASE
				WHEN REGEXP_LIKE(f.NOMOR_REKENING, '_') THEN
					TO_NUMBER(SUBSTR(SUBSTR(f.NOMOR_REKENING, 1, 15), -3))
				ELSE
					TO_NUMBER(SUBSTR(f.NOMOR_REKENING, -3))
			END AS last_id
		FROM %[1]s.FINCOLLATERALASSET f
		WHERE nomor_rekening LIKE ('%%'||$1||'%%')||'%%'

-- name: GetHoldDanaRekening-main
SELECT * 
		FROM %[1]s.HOLDDANAREKENING 
		WHERE %[2]s

-- name: GetHoldDanaRekening-filter-IdHoldDana
ID_HOLD_DANA = $1

-- name: GetHoldDanaRekening-filter-NomorReferensiHold
NOMOR_REFERENSI = $2

-- name: UpdateSaldoHoldReleaseAgunan-main
UPDATE %[1]s.REKENINGLIABILITAS
		SET SALDO_DITAHAN = SALDO_DITAHAN - $1
		WHERE NOMOR_REKENING = $2

-- name: UpdateFinMusyarakahAccount-main
UPDATE %[1]s.finmusyarakahaccount SET
			total_projected_income = $1,
			pricing_model = $2,
			is_balloon_payment = $3
		WHERE nomor_rekening = $4

-- name: CreateFinPaymentSchedule-main
INSERT INTO %[1]s.finpaymentschedule (
			ID_SCHEDULE
			, CREATION_DATE
			, CUSTOMER_NAME
			, IS_ACTIVE
			, INITIAL_PRINCIPAL_AMT
			, INITIAL_PROFIT_AMT
			, PRODUCT_CODE
			, USER_ID
			, NOMOR_REKENING
			, NOMOR_NASABAH
			, IS_TEMPORARY
			, COMPLETION_STATUS
			, ACTIVE_SEQ
			, PERIOD
			, FIRST_INSTALL_DATE
			, DROPPING_DATE
			, ID_PREV_SCHEDULE
			, NORMATIVE_SEQ
			, MUKASAH_SEQ_LIMIT
			, PERIOD_INITIAL
			, ACCRUE_SEQ
			, RBH_PBH_80
			, RBH_PBH_50
			, MUKASAH_PERIOD
			, MUKASAH_RATE
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
		)

-- name: CreateFinPaymentSchedDetail-main
INSERT INTO %[1]s.finpaymentscheddetail (
			ID_SCHEDDETAIL
			, ID_SCHEDULE
			, SCHED_DATE
			, SEQ_NUMBER
			, PRINCIPAL_AMOUNT
			, PROFIT_AMOUNT
			, REALIZED
			, PAYMENT_DATE
			, ID_DETILTRANSGROUP
			, REALIZED_MARGIN
			, OWNERSHIP_FIX_AMOUNT
			, REALIZED_OWNERSHIP_FIX_AMOUNT
			, ARREAR
			, ID_DETILTRANSGROUP_TUNGGAKAN
			, MUKASAH_AMOUNT
			, PRINCIPAL_OUTSTANDING
			, MUKASAH_PCT
			, PROFIT_RATIO
			, PROFIT_AMOUNT_WHOLE
			, DEPRECATED
			, PRINCIPAL_AMOUNT_NEGATIF
			, PROFIT_AMOUNT_NEGATIF
			, PRINCIPAL_AMOUNT_0
			, PROFIT_AMOUNT_0
			, REALIZED_MUKASAH
			, ACCRUED
			, ID_DETILTRANSGROUP_ACCR
			, REALIZED_PROFIT_RATIO
			, REALIZED_PROFIT_AMOUNT_WHOLE
			, REALIZED_PENALTY
			, ACCRUED_AMOUNT
			, PENALTY_TAWIDH
			, PENALTY_TAZIR
			, PENALTY_COUNTER
			, ID_HOLD_DANA
			, PBH_VS_RBH
			, PBH_VS_RBH_AKUM
			, EQV_RATE
		) VALUES (
			 nextval('%[1]s.seq_finpaymentscheddetail'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
			, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
			, $31, $32, $33, $34, $35, $36, $37
		)

-- name: UpdateFinPaymentSchedule-main
UPDATE %[1]s.finpaymentschedule SET
			PERIOD = $1,
			INITIAL_PROFIT_AMT = $2,
			MUKASAH_RATE = $3,
			MUKASAH_PERIOD = $4
		WHERE ID_SCHEDULE = $5

-- name: GetListTabGiro-main
SELECT 
		n.nomor_nasabah,
		n.is_wic,
		l.nomor_rekening,
		t.nama_rekening,
		t.kode_cabang,
		p.kode_produk,
		p.nama_produk,
		l.jenis_rekening_liabilitas,
		case 
			when l.jenis_rekening_liabilitas = 'T' then 'Tabungan'
			when l.jenis_rekening_liabilitas = 'G' then 'Giro'
			when l.jenis_rekening_liabilitas = 'R' then 'Rencana'
			else ''
		end as jenis_rekening_liabilitas_desc,
		case
			when t.status_rekening = 3 then 0
			else t.saldo - p.saldo_minimum - coalesce(l.saldo_ditahan, 0)
		end as saldo_efektif,
		t.status_rekening,
		case 
			when t.status_rekening = 1 then 'Aktif'
			when t.status_rekening = 2 then 'Dormant'
			when t.status_rekening = 3 then 'Tutup'
			else ''
		end as status_rekening_desc,
		case 
			when l.is_blokir = 'T' then 'blokir_debet_credit'
			when l.is_blokir_debet = 'T' then 'blokir_debet'
			when l.is_blokir_kredit = 'T' then 'blokir_credit'
			else ''
		end jenis_blokir,
		l.is_blokir_debet, 
		l.is_blokir_kredit, 
		l.is_sedang_ditutup,
		coalesce(l.is_blokir, 'F') is_blokir,
		n.alamat,
		n.jenis_nasabah_kode jenis_nasabah,
		n.jenis_nasabah jenis_nasabah_desc
	FROM (
		SELECT  
		n.nomor_nasabah,
		n.is_wic,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.nama_lengkap ELSE nk.nama_perusahaan END) AS nama_lengkap,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.id_individu ELSE 0 END) AS id_individu,
		(CASE n.jenis_nasabah WHEN 'I' 
			THEN 
			i.alamat_rumah_jalan||' '||i.alamat_rumah_rt||' '||i.alamat_rumah_rw||' '||
			i.alamat_rumah_kelurahan||' '||i.alamat_rumah_kecamatan||' '||i.alamat_rumah_kota_kabupaten||' '||
			i.alamat_rumah_provinsi||' '||i.alamat_rumah_kode_pos
			ELSE 
			nk.alamat_jalan||' '||nk.alamat_rt||' '||nk.alamat_rw||' '||
			nk.alamat_kelurahan||' '||nk.alamat_kecamatan||' '||nk.alamat_kota_kabupaten||' '||
			nk.alamat_provinsi||' '||nk.alamat_kode_pos 
		END) AS alamat,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.nama_ibu_kandung ELSE '' END) AS nama_ibu_kandung,
		(CASE n.jenis_nasabah WHEN 'I' 
			THEN trim(i.telepon_rumah_kode_negara)||'-'||trim(i.telepon_rumah_kode_area)||'-'||trim(i.telepon_rumah_nomor)
			ELSE trim(nk.telp_kantor1_kode_negara)||'-'||trim(nk.telp_kantor1_kode_area)||'-'||trim(nk.telp_kantor1_nomor)
		END) AS telepon1,
		(CASE n.jenis_nasabah WHEN 'I' 
			THEN trim(i.telepon_rumah2_kode_area)||'-'||trim(i.telepon_rumah2_nomor)
			ELSE trim(nk.telp_kantor2_kode_negara)||'-'||trim(nk.telp_kantor2_kode_area)||'-'||trim(nk.telp_kantor2_nomor)
		END) AS telepon2,
		(CASE n.jenis_nasabah WHEN 'I' 
			THEN ' '
			ELSE trim(nk.telp_kantor3_kode_negara)||'-'||trim(nk.telp_kantor3_kode_area)||'-'||trim(nk.telp_kantor3_nomor)
		END) AS telepon3,
		(CASE n.jenis_nasabah WHEN 'I' 
			THEN trim(i.telepon_hp_kode_negara)||'-'||trim(i.telepon_hp_nomor)
			ELSE ' '
		END) AS handphone,
		(CASE n.jenis_nasabah WHEN 'I' THEN 'Individu' ELSE 'Korporat' END) AS jenis_nasabah,
		n.jenis_nasabah AS jenis_nasabah_kode,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.jenis_identitas ELSE '' END) AS jenis_identitas,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.nomor_identitas ELSE nk.akte_pendirian_nomor END) AS nomor_identitas,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.tanggal_lahir ELSE nk.akte_pendirian_tanggal END) AS tgl_lahir,
		(CASE n.jenis_nasabah WHEN 'I' THEN '' ELSE nk.nomor_siupp END) AS nomor_siupp,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.nomor_npwp ELSE nk.nomor_npwp END) AS nomor_npwp_ori,
		(CASE n.jenis_nasabah WHEN 'I' THEN i.nomor_npwp_plain ELSE nk.nomor_npwp_plain END) AS nomor_npwp_plain
		FROM 
		%[1]s.nasabah n
		LEFT JOIN %[1]s.nasabahindividu ni ON ni.nomor_nasabah = n.nomor_nasabah
		LEFT JOIN %[1]s.nasabahkorporat nk ON nk.nomor_nasabah = n.nomor_nasabah
		LEFT JOIN %[1]s.individu i ON i.id_individu = ni.id_individu
	) n
	INNER JOIN %[1]s.rekeningliabilitas l ON l.nomor_nasabah = n.nomor_nasabah
	LEFT JOIN %[1]s.rekeningtransaksi t ON t.nomor_rekening = l.nomor_rekening
	LEFT JOIN %[1]s.produk p ON p.kode_produk = l.kode_produk                    
	LEFT JOIN %[1]s.deposito d ON t.nomor_rekening = d.nomor_rekening
	WHERE l.jenis_rekening_liabilitas in ('T', 'G', 'R')
		and t.kode_jenis in ('SAV', 'CA') 
		%[2]s
	OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY

-- name: GetListTabGiro-filter-KodeCabang
AND UPPER(t.kode_cabang) = $3

-- name: GetListTabGiro-filter-JenisRekeningLiabilitas
and l.jenis_rekening_liabilitas = $4

-- name: GetListTabGiro-filter-StatusRekening
and t.status_rekening = $5

-- name: GetListTabGiro-filter-NomorNasabah
AND UPPER(n.nomor_nasabah) = $6

-- name: GetListTabGiro-filter-NomorRekening
AND upper(t.nomor_rekening) like $7

-- name: GetListTabGiro-filter-JenisNasabah
AND UPPER(n.jenis_nasabah_kode) = $8

-- name: GetListTabGiro-filter-Npwp
AND (UPPER(n.nomor_npwp_ori) LIKE $9
	OR UPPER(n.nomor_npwp_plain) LIKE $9)

-- name: GetListTabGiro-filter-NomorIdentitas
AND UPPER(n.nomor_identitas) LIKE $10

-- name: GetListTabGiro-filter-NamaNasabah
AND UPPER(n.nama_lengkap) LIKE $11

-- name: GetListTabGiro-filter-LandPhone
AND (
	UPPER(n.telepon1) LIKE $12
	OR UPPER(n.telepon2) LIKE $12
	OR UPPER(n.telepon3) LIKE $12
)

-- name: GetListTabGiro-filter-Handphone
AND UPPER(n.handphone) LIKE $13

-- name: GetListTabGiro-filter-TanggalLahir
AND TO_DATE(TO_CHAR(n.tgl_lahir, '%[1]s'), '%[1]s') = TO_DATE($14, '%[1]s')

-- name: GetListTabGiro-filter-IbuKandung
AND UPPER(n.nama_ibu_kandung) LIKE $15

-- name: GetListTabGiro-filter-Address
AND UPPER(n.alamat) LIKE $16

-- name: GetListTabGiro-filter-Keyword
and ((upper(t.nama_rekening) like $17 or upper(t.nomor_rekening) like $17))

-- name: GetRekeningTabgirDetail-main
select rl.nomor_nasabah, n.jenis_nasabah, rl.nomor_rekening, rt.nama_rekening, rt.saldo, rt.kode_cabang, c.nama_cabang,
			to_char(rl.tanggal_buka, 'YYYY-MM-DD') as tanggal_buka, to_char(rl.tanggal_tutup, 'YYYY-MM-DD') as tanggal_tutup, 
			rl.nisbah_spesial, rl.user_update, to_char(rt.tanggal_update, 'YYYY-MM-DD') as tanggal_update,
			rl.nisbah_dasar, rl.tarif_pajak, rl.persentase_zakat_bagi_hasil, rl.kode_sumber_dana, rl.kode_tujuan_rekening, 
			rl.kode_marketing_current, rl.kode_marketing_pertama,
			rt.saldo - (coalesce(rl.saldo_float, 0.0) + coalesce(rl.saldo_ditahan, 0.0) + coalesce(p.saldo_minimum, 0.0)) as saldo_efektif,
			case 
				when rl.jenis_rekening_liabilitas = 'T' then 'Tabungan'
				when rl.jenis_rekening_liabilitas = 'R' then 'Rencana'
			end as jenis_rekening_liabilitas_desc,
			rl.saldo_float, rl.saldo_deposito, rl.saldo_ditahan, p.saldo_minimum, rl.saldo_tunggakan, rl.jenis_rekening_liabilitas, rl.kode_produk, 
			p.nama_produk, rl.is_kena_biayalayananumum, rl.is_status_passbook, rl.alamat_kirim, rt.status_rekening, rt.user_input, rl.user_otorisasi,
			rt.keterangan, rl.is_join_account, rt.kode_valuta, rt.nama_valuta, rl.is_sedang_ditutup,rl.kode_program, rl.kode_marketing_referensi, 
			rl.status_kelengkapan, rl.is_blokir, rl.is_blokir_debet, rl.is_blokir_kredit, rl.is_cetak_nota, rl.is_dapat_bagi_hasil,
			rl.is_tidak_dormant,  rl.is_biaya_rekening_dormant, rl.is_biaya_saldo_minimum,rl.is_biaya_atm, rl.nisbah_bagi_hasil, 
			rl.is_tiering_nisbahbonus, rl.saldo_accrual_bagihasil,to_char(rl.tgl_trans_cabang_terakhir, 'YYYY-MM-DD') as tgl_trans_cabang_terakhir,
			to_char(rl.tgl_trans_echannel_terakhir, 'YYYY-MM-DD') as tgl_trans_echannel_terakhir,to_char(rl.tgl_transaksi_terakhir, 'YYYY-MM-DD') as tgl_transaksi_terakhir, 
			rl.jenis_register_layanan, rl.status_restriksi, coalesce(rl.kode_multicifrelation, rd.reference_code) AS kode_multicifrelation
		from %[1]s.rekeningliabilitas rl
			inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
			inner join %[1]s.produk p on p.kode_produk = rl.kode_produk
			inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah 
			inner join %[2]s.cabang c on c.kode_cabang = rt.kode_cabang 
			left join %[1]s.registerlayanan rl on rl.nomor_rekening_layanan = rt.nomor_rekening
			LEFT JOIN %[2]s.referencedata rd ON rd.refdata_id = rl.id_multicifrelation
		where rl.jenis_rekening_liabilitas in ('T', 'G', 'R') 
			and rt.kode_jenis in ('SAV', 'CA') 
			and rl.nomor_rekening = $1

-- name: GetRekeningTabgirDetail-filter-StatusRekening
and rt.status_rekening = $2

-- name: UpdateIsSedangTutup-main
update %[1]s.rekeningliabilitas
		set is_sedang_ditutup = $1
		where nomor_rekening = $2

-- name: UpdateIsTutupKartuAtm-main
update %[1]s.rekeningliabilitas
		set is_tutup_kartu_atm = $1
		where nomor_rekening = $2

-- name: GetAlamatAlternatif-main
select a.alamat_jalan, a.alamat_rtrw , a.alamat_kelurahan, a.alamat_kecamatan, 
			a.alamat_kota_kabupaten, a.alamat_provinsi, a.alamat_kode_pos, a.telepon
		from %[1]s.alamatalternatif a
		where a.nomor_rekening = $1

-- name: GetHakAksesRekening-main
select l.id_individu, i.nama_lengkap, i.tempat_lahir, to_char(i.tanggal_lahir, 'yyyy-mm-dd') as tanggal_lahir, i.status_perkawinan, 
			i.kewarganegaraan,i.jenis_identitas, i.nomor_identitas, i.jenis_identitas, i.alamat_rumah_jalan, i.alamat_rumah_rt, i.alamat_rumah_rw, 
			i.alamat_rumah_kelurahan, i.alamat_rumah_kecamatan, i.alamat_rumah_kota_kabupaten, i.alamat_rumah_provinsi, i.alamat_rumah_kode_pos, 
			coalesce(i.kode_pekerjaan, kp.reference_desc) as kode_pekerjaan, i.telepon_hp_nomor, i.jenis_identitas_lain, i.nomor_identitas_lain, i.status_pep, 
			l.kode_sumber_dana, l.is_aktif, l.keterangan, i.jenis_kelamin, coalesce(i.kode_negara_asal, kn.reference_desc) as kode_negara_asal,
			case when i.jenis_kelamin = 'P' then 'Laki-laki' when i.jenis_kelamin = 'W' then 'Perempuan' end jenis_kelamin_desc, kn.reference_desc kode_negara_asal_desc
		from %[1]s.individu i
			inner join %[1]s.listindividurekening l on l.id_individu = i.id_individu 
			inner join %[1]s.rekeningtransaksi r  on l.nomor_rekening = r.nomor_rekening   
			-- left join %[2]s.referencedata kn on kn.refdata_id = i.id_negara 
			left join %[2]s.referencedata kn on kn.reference_code = i.KODE_NEGARA_ASAL
			left join %[2]s.referencedata kp on kp.refdata_id = i.id_pekerjaan 
			where l.nomor_rekening = $1

-- name: GetJoinAccount-main
select n.nomor_nasabah ,n.nama_nasabah,
			case
				when n.jenis_nasabah = 'K' then nk.alamat_jalan || ' ' || nk.alamat_rt || ' ' || nk.alamat_rw || ' ' || nk.alamat_kelurahan || ' ' || nk.alamat_kecamatan || ' ' ||
					nk.alamat_kota_kabupaten || ' ' || nk.alamat_provinsi || ' ' || nk.alamat_kode_pos 
				when n.jenis_nasabah = 'I' then i.alamat_rumah_jalan || ' ' || i.alamat_rumah_rt || ' ' || i.alamat_rumah_rw || ' ' || i.alamat_rumah_kelurahan || ' ' || i.alamat_rumah_kecamatan 
					|| ' ' || i.alamat_rumah_kota_kabupaten || i.alamat_rumah_provinsi || ' ' || i.alamat_rumah_kode_pos
			end as alamat,
			case 
				when n.jenis_nasabah = 'I' then i.nama_ibu_kandung
				else ''
			end as nama_ibu_kandung,
			case 
				when n.jenis_nasabah = 'K' then nk.nomor_siupp
				else ''
			end as nomor_siup,
			case 
				when n.jenis_nasabah = 'K' then nk.telp_kantor1_kode_area || '-' || nk.telp_kantor1_nomor
				when n.jenis_nasabah = 'I' then i.telepon_rumah_kode_area || '-' || i.telepon_rumah_nomor
			end as telepon1,
			case 
				when n.jenis_nasabah = 'K' then nk.telp_kantor2_kode_area || '-' || nk.telp_kantor2_nomor
				when n.jenis_nasabah = 'I' then i.telepon_rumah2_kode_area || '-' || i.telepon_rumah2_nomor
			end as telepon2  
		from %[1]s.multiciflink m
			inner join %[1]s.nasabah n on n.nomor_nasabah = m.nomor_nasabah
			left join %[1]s.nasabahindividu ni on ni.nomor_nasabah = n.nomor_nasabah 
			left join %[1]s.individu i on i.id_individu = ni.id_individu 
			left join %[1]s.nasabahkorporat nk on nk.nomor_nasabah = n.nomor_nasabah
		where nomor_rekening = $1
		order by n.nomor_nasabah 

-- name: GetBagiHasil-main
select rl.is_blokir_debet, rl.is_blokir_kredit, rl.is_cetak_nota , rl.is_dapat_bagi_hasil,
			rl.is_tidak_dormant, rl.is_biaya_rekening_dormant, rl.is_biaya_saldo_minimum,
			rl.is_biaya_atm, rl.nisbah_bagi_hasil, rl.is_dapat_bagi_hasil, rl.is_tiering_nisbahbonus 
		from %[1]s.rekeningliabilitas rl  
		where nomor_rekening = $1

-- name: GetTellerDetail-main
select rt.nomor_rekening, rt.nama_rekening, rt.kode_cabang, rt.saldo, rt.kode_valuta
			, case
				when tl.is_open = 'T' then 'Open'
				when tl.is_open = 'F' then 'Close'
				else ''
			end as is_open_desc
			, rt.status_rekening
			, case
				when rt.status_rekening = 1 then 'Aktif'
				when rt.status_rekening = 2 then 'Dormant'
				when rt.status_rekening = 3 then 'Tutup'
				else ''
			end as status_rekening_desc
			, rt.tanggal_update
		from %[1]s.rekeningtransaksi rt
			inner join %[1]s.kasteller tl on tl.nomor_rekening = rt.nomor_rekening
		where
			tl.id_user = $1
			and rt.kode_cabang = $2
			and rt.kode_valuta = $3

-- name: GetTieringBagiHasil-main
SELECT ITEMID , LOWERLIMITVALUE , UPPERLIMITVALUE , amount1 
		FROM %[1]s.TIERINGITEM t
		where tieringid = $1

-- name: GetTieringBiayaADM-main
SELECT ITEMID , LOWERLIMITVALUE , UPPERLIMITVALUE , amount1 
		FROM %[1]s.TIERINGITEM t
		where tieringid = $1

-- name: GetTrxList-main
select t.id_transaksi, to_char(t.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi, t.nomor_referensi,
			t.kode_cabang_transaksi, t.nilai_transaksi, t.keterangan
		from %v.transaksi t
		%v
		order by t.jam_input
		offset %d rows fetch next %d rows only

-- name: GetTrxList-filter-Default
where 1=1

-- name: GetTrxList-filter-NomorReferensi
and t.nomor_referensi = $1

-- name: GetTrxList-filter-KodeCabangTransaksi
and t.kode_cabang_transaksi = $2

-- name: GetTrxByNoref-main
select t.id_transaksi ,t.nomor_referensi
		from %[1]s.transaksi t
		where t.nomor_referensi = $1

-- name: GetTrxByIdTrx-main
select t.id_transaksi, t.nomor_referensi, t.keterangan, t.jenis_aplikasi, t.kode_valuta_transaksi, t.kurs_manual, t.nilai_transaksi,
			t.kode_cabang_transaksi, t.is_sudah_dijurnal, t.kode_transaksi, t.status_otorisasi, t.nomor_seri, t.journal_no, t.nomor_counter,
			t.user_input, t.user_otorisasi, t.tanggal_transaksi, t.tanggal_otorisasi, t.jam_input, t.jam_otorisasi, t.tanggal_input,
			t.terminal_input, t.terminal_otorisasi, t.channelname, t.is_reversed, t.biaya, t.is_reverse_allowed
		from %[1]s.transaksi t
		where t.id_transaksi = $1

-- name: GetTrxDetail-main
select d.id_detil_transaksi, d.id_transaksi, d.jenis_mutasi, d.kode_valuta, to_char(d.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi,
			d.nilai_mutasi, d.keterangan, d.kode_cabang, d.jenis_detil_transaksi, d.nilai_kurs_manual, d.saldo_awal, d.nilai_ekuivalen,
			d.nomor_referensi, d.kode_account, d.nomor_rekening, d.kode_jurnal, d.id_parameter_transaksi, d.flag_proses, d."SEQUENCE"
		from %v.detiltransaksi d
		where d.id_transaksi = $1

-- name: GetTrxDetailRev-main
select d.id_detil_transaksi, d.id_transaksi, d.jenis_mutasi, d.kode_valuta, d.kode_valuta_rak, to_char(d.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi,
			d.nilai_mutasi, d.keterangan, d.kode_cabang, d.jenis_detil_transaksi, d.nilai_kurs_manual, d.saldo_awal, d.nilai_ekuivalen,
			d.nomor_referensi, d.kode_account, d.nomor_rekening, d.kode_jurnal, d.id_parameter_transaksi, d.flag_proses, d."SEQUENCE",
			d.kode_tx_class, d.nomor_seri_warkat, d.flag_code
		from %v.detiltransaksi d
		where d.id_transaksi = $1
		order by d.id_detil_transaksi desc

-- name: GetTrxDetailTransfer-main
select d.id_detil_transaksi
		from %v.detiltransaksi d
		where d.id_transaksi = $1
			and d.is_create_remittance = $2

-- name: GetHistTrxListByNorek-main
select hd.id_detil_transaksi, hd.id_transaksi, to_char(ht.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, hd.id_parameter_transaksi, hd.kode_jurnal, 
			ht.kode_transaksi, hd.keterangan, hd.jenis_mutasi, hd.saldo_awal, round(hd.nilai_mutasi, 2) as nilai_mutasi, hd.kode_valuta, hd.nilai_kurs_manual, hd.nilai_ekuivalen, 
			ht.nomor_referensi, ht.kode_cabang_transaksi, ht.user_input, ht.user_otorisasi, to_char(ht.tanggal_input, 'yyyy-mm-dd') as tanggal_input, 
			to_char(ht.jam_input, 'HH24::MI::SS') as jam_input, ht.keterangan_tambahan, ht.journal_no, ht.status_otorisasi,
			case
				when ht.status_otorisasi = 0 then 'Menunggu Otorisasi'
				when ht.status_otorisasi = 1 then 'Sudah Otorisasi'
				when ht.status_otorisasi = 2 then 'Sudah Direject'
				else ''
			end as status_otorisasi_desc
		from %[1]s.histtransaksi ht
			inner join %[1]s.histdetiltransaksi hd on hd.id_transaksi = ht.id_transaksi
		where ht.tanggal_transaksi >= to_date($1, 'yyyy-mm-dd') 
			and ht.tanggal_transaksi < to_date($2, 'yyyy-mm-dd') + 1
			and hd.nomor_rekening = $3
		order by hd.id_detil_transaksi

-- name: GetTrxListByNorek-main
select d.id_detil_transaksi, d.id_transaksi, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, d.id_parameter_transaksi, d.kode_jurnal, 
			t.kode_transaksi, d.keterangan, d.jenis_mutasi, d.saldo_awal, round(d.nilai_mutasi, 2) as nilai_mutasi, d.kode_valuta, d.nilai_kurs_manual, d.nilai_ekuivalen, 
			t.nomor_referensi, t.kode_cabang_transaksi, t.user_input, t.user_otorisasi, to_char(t.tanggal_input, 'yyyy-mm-dd') as tanggal_input, 
			to_char(t.jam_input, 'hh::mm::ss') as jam_input, t.keterangan_tambahan, t.journal_no, t.status_otorisasi,
			case
				when t.status_otorisasi = 0 then 'Menunggu Otorisasi'
				when t.status_otorisasi = 1 then 'Sudah Otorisasi'
				when t.status_otorisasi = 2 then 'Sudah Direject'
				else ''
			end as status_otorisasi_desc
		from %[1]s.transaksi t
			inner join %[1]s.detiltransaksi d on d.id_transaksi = t.id_transaksi
		where d.nomor_rekening = $1
		order by d.id_detil_transaksi

-- name: GetAcctStatementHeaderPrint-main
select rl.nomor_rekening, rt.nama_rekening, c.kode_cabang, c.nama_cabang, rl.alamat_kirim, n.jenis_nasabah, rl.nomor_nasabah
		from %[1]s.rekeningliabilitas rl 
			inner join %[1]s.nasabah n on rl.nomor_nasabah = n.nomor_nasabah 
			left join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening 
			left join %[2]s.cabang c on rt.kode_cabang = c.kode_cabang 
		where rl.nomor_rekening = $1

-- name: GetSumTrxAcctStatementHeader-main
select 
			sum(case
				when dt.kode_transaksi not in ('SC','BC','SD','SDZ','SDP','3002','6007','8002') then decode(dt.jenis_mutasi, 'D', dt.nilai_mutasi, 0)
				else 0
			end) as debet,
			sum(case
				when dt.kode_transaksi not in ('SC','BC','SD','SDZ','SDP','3002','6007','8002') then decode(dt.jenis_mutasi, 'C', dt.nilai_mutasi, 0)
				else 0
			end) as kredit,
			sum(decode(dt.jenis_mutasi, 'D', 1, 0)) AS jml_debet,
			sum(decode(dt.jenis_mutasi, 'C', 1, 0)) AS jml_kredit,
			sum(case
				when dt.kode_transaksi in ('SC','BC') then decode(dt.jenis_mutasi, 'D', dt.nilai_mutasi, 0)
				else 0
			end) as admin,
			sum(case
				when dt.kode_transaksi in ('SD','3002') then decode(dt.jenis_mutasi, 'C', dt.nilai_mutasi, 0)
				else 0
			end) as bagi_hasil,
			sum(case
				when dt.kode_transaksi in ('SDZ','SDP','6007','8002') then decode(dt.jenis_mutasi, 'D', dt.nilai_mutasi, 0)
				else 0
			end) as beban
		from %[1]s.rekeningtransaksi rtrans
			join %[1]s.rekeningliabilitas rlib on rtrans.nomor_rekening = rlib.nomor_rekening
			join (
				select t.kode_transaksi, d.id_detil_transaksi, d.nomor_rekening, d.jenis_mutasi, d.nilai_mutasi
				from %[1]s.histtransaksi t
				join %[1]s.histdetiltransaksi d on t.id_transaksi = d.id_transaksi
				where t.tanggal_transaksi >= to_timestamp($1, 'YYYY-MM-DD HH24::MI::SS')
				and t.tanggal_transaksi < to_timestamp($2, 'YYYY-MM-DD HH24::MI::SS') + 1
					and d.kode_tx_class is null
				union all
				select t.kode_transaksi, d.id_detil_transaksi, d.nomor_rekening, d.jenis_mutasi, d.nilai_mutasi
				from %[1]s.transaksi t
				join %[1]s.detiltransaksi d on t.id_transaksi = d.id_transaksi
				where t.tanggal_transaksi >= to_timestamp($1, 'YYYY-MM-DD HH24::MI::SS')
					and t.tanggal_transaksi < to_timestamp($2, 'YYYY-MM-DD HH24::MI::SS') + 1
					and d.kode_tx_class is null
			) dt on rtrans.nomor_rekening = dt.nomor_rekening
		where  rtrans.nomor_rekening = $3
		group by rtrans.nomor_rekening

-- name: GetHistTrxList-main
select hd.id_detil_transaksi, to_char(ht.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, hd.id_parameter_transaksi, hd.kode_jurnal, 
			ht.kode_transaksi, hd.keterangan, hd.jenis_mutasi, hd.saldo_awal, hd.nilai_mutasi, ht.nomor_referensi, ht.status_otorisasi, 
			case
				when ht.status_otorisasi = 0 then 'Menunggu Otorisasi'
				when ht.status_otorisasi = 1 then 'Sudah Otorisasi'
				when ht.status_otorisasi = 2 then 'Sudah Direject'
				else ''
			end as status_otorisasi_desc,
			ht.kode_cabang_transaksi, ht.user_otorisasi, ht.user_input, to_char(ht.tanggal_input, 'yyyy-mm-dd') as tanggal_input, 
			to_char(ht.jam_input, 'HH24::MI::SS') as jam_input, ht.keterangan_tambahan
		from %[1]s.histtransaksi ht
			inner join %[1]s.histdetiltransaksi hd on hd.id_transaksi = ht.id_transaksi
		%[2]s
		order by hd.id_detil_transaksi
		offset $1 rows fetch next $2 rows only

-- name: GetHistTrxList-filter-Default
where ht.tanggal_transaksi >= to_date($3, 'yyyy-mm-dd') and ht.tanggal_transaksi < to_date($4, 'yyyy-mm-dd')

-- name: GetHistTrxList-filter-NomorReferensi
and ht.nomor_referensi = $5

-- name: GetHistTrxList-filter-KodeCabangTransaksi
and ht.kode_cabang_transaksi = $6

-- name: GetHistTrxList-filter-NomorRekening
and hd.nomor_rekening = $7

-- name: GetSaldoByDate-main
select round(balance, 2) as balance
		from %[1]s.dailybalancerekening d1
			inner join (
				select nomor_rekening, max(balance_date) as balance_date 
				from %[1]s.dailybalancerekening 
				where nomor_rekening = $1
					and balance_date <= to_date($2, 'yyyy-mm-dd')
					and balance_field = 'saldo'
				group by nomor_rekening
			) d2 on (d1.balance_date =  d2.balance_date) and (d1.nomor_rekening = d2.nomor_rekening)
		where d1.nomor_rekening = $1    
			and d1.balance_field = 'saldo'
			and d1.balance_date <= to_date($2, 'yyyy-mm-dd')

-- name: GetHisttrxDetail-main
select d.id_detil_transaksi, d.id_transaksi , d.jenis_mutasi, d.kode_valuta, d.kode_valuta_rak,
			to_char(d.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi, d.nilai_mutasi, d.keterangan,
			d.kode_tx_class
		from %v.histdetiltransaksi
		d where d.id_transaksi = $1

-- name: GetHisttrxDetailRev-main
select d.id_detil_transaksi, d.id_transaksi , d.jenis_mutasi, d.kode_valuta, d.kode_valuta_rak,
			to_char(d.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi, d.nilai_mutasi, d.keterangan,
			d.kode_tx_class
		from %v.histdetiltransaksi
		d where d.id_transaksi = $1
		order by d.id_detil_transaksi desc

-- name: PostTransaksi-main
insert into %[1]s.transaksi
		(id_transaksi, nomor_referensi, keterangan, jenis_aplikasi, kode_valuta_transaksi, kurs_manual, nilai_transaksi, nilai_ekuivalen,
			kode_cabang_transaksi, is_sudah_dijurnal, kode_transaksi, status_otorisasi, nomor_seri, journal_no, nomor_counter, user_input, biaya,
			user_otorisasi, tanggal_transaksi, tanggal_otorisasi,  jam_input, jam_otorisasi, tanggal_input, terminal_input, terminal_otorisasi, channelname,
			is_reverse_allowed)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, current_timestamp, current_timestamp, current_timestamp, $21, $22, $23, $24)

-- name: PostDetailTrx1-main
insert into %[1]s.detiltransaksi
		(id_detil_transaksi, tanggal_transaksi, jenis_mutasi, nilai_mutasi, keterangan, kode_cabang, kode_valuta, jenis_detil_transaksi,
			nilai_kurs_manual, saldo_awal, nilai_ekuivalen, nomor_referensi, kode_account, kode_valuta_rak, kode_jurnal, id_transaksi, nomor_rekening,
			jenis_transaksi_liabilitas, id_parameter_transaksi, flag_proses, is_create_remittance, sequence, kode_tx_class, nomor_seri_warkat,
			flag_code, kode_rc)
		values ( nextval('%[1]s.seq_detiltransaksi'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)

-- name: PostDetailTrx2-main
insert into %[1]s.detiltransaksi
		(id_detil_transaksi, tanggal_transaksi, jenis_mutasi, nilai_mutasi, keterangan, kode_cabang, kode_valuta, jenis_detil_transaksi,
			nilai_kurs_manual, saldo_awal, nilai_ekuivalen, nomor_referensi, kode_account, kode_valuta_rak, kode_jurnal, id_transaksi, nomor_rekening, 
			jenis_transaksi_liabilitas, id_parameter_transaksi, flag_proses, is_create_remittance, sequence, kode_tx_class, nomor_seri_warkat,
			flag_code, kode_rc)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)

-- name: UpdateIsReversedTransaksi-main
update %[1]s.transaksi
		set is_reversed = 'T'
		where id_transaksi = $1

-- name: InsertInfoTransaksi-main
insert into %[1]s.infotransaksi
		(id_transaksi, nomor_rekening_debet, nomor_rekening_kredit, nama_rekening_debet, nama_rekening_kredit, is_tunai, kode_bank_kliring, kode_member_rtgs,
			nominal_transaksi, nominal_biaya, account_biaya, nomor_rekening_transaksi, nama_rekening_transaksi, nomor_rekening_pair, nama_rekening_pair, sumber_biaya,
			jenis_mutasi_transaksi, jenis_mutasi_pair, is_confidential, nomor_nasabah_wic, id_notadebitoutwardgroup, nominal_tunai_fisik,
			is_tunai_fisik, id_sumber_dana_trx, id_tujuan_trx, kd_sumber_dana_trx, kd_tujuan_trx, ket_sumber_dana_trx, ket_tujuan_trx, tipe_akses_transaksi, listnama_akses_transaksi)
		values
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31)

-- name: GetInfotransaksi-main
select id_transaksi, nomor_rekening_debet, nomor_rekening_kredit, nama_rekening_debet, nama_rekening_kredit, is_tunai, kode_bank_kliring, 
			kode_member_rtgs, nominal_transaksi, nominal_biaya, account_biaya, nomor_rekening_transaksi, nama_rekening_transaksi, 
			nomor_rekening_pair, nama_rekening_pair, sumber_biaya, jenis_mutasi_transaksi, jenis_mutasi_pair, counter_cetak_validasi, is_confidential, 
			nomor_nasabah_wic, id_notadebitoutwardgroup, nominal_tunai_fisik, is_tunai_fisik, id_sumber_dana_trx, id_tujuan_trx, kd_sumber_dana_trx, 
			kd_tujuan_trx, ket_sumber_dana_trx, ket_tujuan_trx, tipe_akses_transaksi, listnama_akses_transaksi 
		from %[1]s.infotransaksi i
		where id_transaksi = $1

-- name: InsertListAksesTransaksi-main
insert into %[1]s.listaksestransaksi
		(id_transaksi, id_detil_transaksi, id_individu, nomor_cif)
		values
		($1, $2, $3, $4)

-- name: GetTrxDetailByParams-main
select d.id_detil_transaksi
		from %v.detiltransaksi d
		where d.id_transaksi = $1
			and d.jenis_mutasi = $2
			and d.nomor_rekening = $3
			and d.nilai_mutasi = $4

-- name: GetListTrxDetailByParams-main
select d.id_detil_transaksi, d.nomor_rekening, d.nilai_mutasi
		from %v.detiltransaksi d
		where d.id_transaksi = $1
			and d.flag_code = $2

-- name: GetTrxDetailTolakanKliring-main
select dko.id_notadebitoutward
		from %[1]s.detiltransaksi d
			inner join %[1]s.transaksi t on t.id_transaksi = d.id_transaksi
			inner join %[1]s.dtkliringoutward dko on dko.id_detil_transaksi = d.id_detil_transaksi
		where d.id_transaksi = $1
			and d.jenis_mutasi = $2
			and d.nomor_rekening = $3
			and d.nilai_mutasi = $4

-- name: GetDetailTrxTDN-main
select d.id_detil_transaksi, d.id_transaksi, d.nomor_rekening, dko.id_notadebitoutward
		from %[1]s.detiltransaksi d
			inner join %[1]s.transaksi t on t.id_transaksi = d.id_transaksi
			inner join %[1]s.dtkliringoutward dko on dko.id_detil_transaksi = d.id_detil_transaksi
		where d.id_transaksi = $1

-- name: GetSequenceId-main
SELECT nextval('%v.seq_val_transaksi_rekening') 

-- name: InsertValidasiTransRek-main
INSERT
	INTO
	%v.validasi_transaksi_rekening (id_session,
		nomor_rekening ,
		nominal,
		nominal_sign,
		keterangan,
		is_valid)
	VALUES ($1, $2, $3, $4, $5, 'T')

-- name: UpdateJenisMutasi-main
UPDATE
	%v.validasi_transaksi_rekening tmp
	SET
		jenis_mutasi =
	CASE
			WHEN nominal_sign = '+' THEN 'C'
			ELSE 'D'
		END
	WHERE
		ID_SESSION = $1

-- name: UpdateIsValidRekening-main
UPDATE
	%v.validasi_transaksi_rekening tmp
	SET
		is_valid = 'F',
		invalid_message = invalid_message || 'nomor rekening tidak terdaftar, '
	WHERE
		ID_SESSION = $1
		AND NOT EXISTS (
		SELECT
			1
		FROM
			%[1]s.REKENINGTRANSAKSI_TMP n
		WHERE
			n.NOMOR_rekening = tmp.NOMOR_rekening)

-- name: UpdateNama-main
MERGE INTO %[1]v.validasi_transaksi_rekening tmp
	USING %[2]v.rekeningtransaksi r 
	ON (r.nomor_rekening = tmp.NOMOR_rekening)
	WHEN MATCHED THEN UPDATE SET NAMA_rekening = r.nama_rekening
	WHERE ID_SESSION = $1

-- name: UpdateRekTutup-main
UPDATE
	%[1]v.validasi_transaksi_rekening tmp
	SET
		is_valid = 'F'
	,
		invalid_message = invalid_message || 'nomor rekening sudah ditutup, '
	WHERE
		ID_SESSION = $1
		AND EXISTS (
		SELECT
			1
		FROM
			%[2]v.rekeningtransaksi n
		WHERE
			n.NOMOR_rekening = tmp.NOMOR_rekening
			AND n.status_rekening = 3)

-- name: UpdateRekDormant-main
UPDATE
	%[1]v.validasi_transaksi_rekening tmp
	SET
		is_valid = 'F'
	,
		invalid_message = invalid_message || 'nomor rekening sedang dormant, '
	WHERE
		ID_SESSION = $1
		AND EXISTS (
		SELECT
			1
		FROM
			%[2]v.rekeningtransaksi n
		WHERE
			n.NOMOR_rekening = tmp.NOMOR_rekening
			AND n.status_rekening = 2)

-- name: UpdateRekTutupLiabilitas-main
UPDATE
	%[1]v.validasi_transaksi_rekening tmp
	SET
		is_valid = 'F'
	,
		invalid_message = invalid_message || 'nomor rekening sedang ditutup, '
	WHERE
		ID_SESSION = $1
		AND EXISTS (
		SELECT
			1
		FROM
			%[2]v.rekeningliabilitas n
		WHERE
			n.NOMOR_rekening = tmp.NOMOR_rekening
			AND n.is_sedang_ditutup = 'T')

-- name: UpdateRekBlokirLiabilitas-main
UPDATE
	%[1]v.validasi_transaksi_rekening tmp
SET
	is_valid = 'F'
  ,
	invalid_message = invalid_message || 'nomor rekening sedang diblokir, '
WHERE
	ID_SESSION = $1
	AND EXISTS (
	SELECT
		1
	FROM
		%[2]v.rekeningliabilitas n
	WHERE
		n.NOMOR_rekening = tmp.NOMOR_rekening
		AND n.is_blokir = 'T')

-- name: SelectValTrxTemp-main
select * from
	%v.validasi_transaksi_rekening 
	where id_session = $1

-- name: DeleteValTrxTemp-main
delete from  %v.validasi_transaksi_rekening tmp where id_session = $1

-- name: InsertValTrxUploadGL-main
INSERT INTO %[1]v.VALIDASI_TRANSAKSI_GL 
		(ID_SESSION, IS_VALID, TANGGAL_INPUT, TANGGAL_POSTING,
		KODE_CABANG, KODE_ACCOUNT, KODE_VALUTA, JENIS_MUTASI,
		NOMOR_REFERENSI, NOMINAL, NOMINAL_EKUIVALEN, KODE_AREA,
		KODE_PAIR, DESCRIPTION)
	VALUES
		($1, 'T', TO_DATE($2, 'YYYY-MM-DD'), TO_DATE($3, 'YYYY-MM-DD'), $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

-- name: ValAccTidakValidTrxUploadGL-main
UPDATE 
			%[1]v.validasi_transaksi_gl tmp
		SET 
			is_valid ='F', invalid_message = invalid_message || 'account GL tidak terdaftar, '
		WHERE 
			ID_SESSION = $1
			AND NOT EXISTS (SELECT 1 FROM %[2]v.account a WHERE a.account_code=tmp.kode_account)

-- name: ValCabangTrxUploadGL-main
UPDATE 
			%[1]v.validasi_transaksi_gl tmp
		SET 
			is_valid ='F', invalid_message = invalid_message || 'kode cabang tidak terdaftar, '
		WHERE 
			ID_SESSION = $1
			AND NOT EXISTS (SELECT 1 FROM %[2]v.cabang c WHERE c.KODE_CABANG =tmp.kode_cabang)

-- name: ValValutaTrxUploadGL-main
UPDATE 
			%[1]v.validasi_transaksi_gl tmp
		SET 
			is_valid ='F', invalid_message = invalid_message || 'kode valuta tidak terdaftar, '
		WHERE 
			ID_SESSION = $1
			AND NOT EXISTS (SELECT 1 FROM %[2]v.currency c WHERE c.currency_code =tmp.kode_valuta)

-- name: ValAccTransaksionalTrxUploadGL-main
UPDATE 
			%[1]v.validasi_transaksi_gl tmp
		SET 
			is_valid ='F', invalid_message = invalid_message || 'account GL tidak terdaftar, '
		WHERE 
			ID_SESSION = $1
			AND EXISTS (SELECT 1 FROM %[2]v.account a WHERE a.account_code=tmp.kode_account and is_detail='F')

-- name: GetResponseValTrxUploadGL-main
select * from
		%[1]v.validasi_transaksi_gl 
		where id_session = $1

-- name: DelTrxUploadGLTemp-main
delete from  %[1]v.validasi_transaksi_gl tmp 
		where id_session = $1

-- name: GetListTrxPayroll-main
SELECT T.ID_TRANSAKSI, T.TANGGAL_TRANSAKSI, USER_INPUT, T.KETERANGAN, T.NOMOR_REFERENSI, T.ID_TRANSAKSI, T.KODE_TRANSAKSI 
	, SUM(D.NILAI_MUTASI) AS TOTAL_NILAI_MUTASI
	, COUNT(NILAI_MUTASI) AS JUMLAH_REKENING  
	FROM %[1]s.TRANSAKSI T 
		INNER JOIN %[1]s.DETILTRANSAKSI D ON T.ID_TRANSAKSI = D.ID_TRANSAKSI  AND JENIS_MUTASI='C'
	WHERE T.KODE_TRANSAKSI IN ('PYROL')
	AND T.TANGGAL_TRANSAKSI >= TO_DATE($1,'YYYY-MM-DD') AND T.TANGGAL_TRANSAKSI <= TO_DATE($2,'YYYY-MM-DD') 
	AND T.USER_INPUT = $3
	GROUP BY T.ID_TRANSAKSI, T.TANGGAL_TRANSAKSI, USER_INPUT, T.KETERANGAN, T.NOMOR_REFERENSI, KODE_TRANSAKSI 
	UNION ALL
	SELECT T.ID_TRANSAKSI, T.TANGGAL_TRANSAKSI, USER_INPUT, T.KETERANGAN, T.NOMOR_REFERENSI, T.ID_TRANSAKSI, T.KODE_TRANSAKSI 
		, SUM(D.NILAI_MUTASI) AS TOTAL_NILAI_MUTASI
		, COUNT(NILAI_MUTASI) AS JUMLAH_REKENING  
	FROM %[1]s.HISTTRANSAKSI T 
		INNER JOIN %[1]s.HISTDETILTRANSAKSI D ON T.ID_TRANSAKSI = D.ID_TRANSAKSI  AND JENIS_MUTASI='C'
	WHERE T.KODE_TRANSAKSI IN ('PYROL')
	GROUP BY T.ID_TRANSAKSI, T.TANGGAL_TRANSAKSI, USER_INPUT, T.KETERANGAN, T.NOMOR_REFERENSI, KODE_TRANSAKSI

-- name: GetSequenceTempPostMultiJournal-main
select nextval('%[1]s.seq_val_transaksi_umum') 

-- name: InsertToPostMultiJournalTempTable-main
insert into %[1]v.validasi_transaksi_umum 
		(id_session, kode_tipe, nomor_rekening, jenis_mutasi, nominal, keterangan, kode_rc, is_valid)
  		values 
		($1, $2 , $3 , $4 , $5 , $6 , $7, $8)

-- name: CheckBalancePostTrxMultiJournalTempTable-main
select jenis_mutasi, sum(nominal) as total_nominal 
		from %[1]v.validasi_transaksi_umum 
		where id_session = $1
		group by jenis_mutasi

-- name: UpdateNamaCabangValutaAccountPostTrxMultiJournalTempTable-main
update %[1]v.validasi_transaksi_umum tmp
set 
	nama_rekening = r.nama_rekening,
	kode_cabang = r.kode_cabang,
	kode_valuta = r.kode_valuta,
	kode_account = r.kode_account,
	saldo = case when r.saldo is null then 0.0 else r.saldo end
from %[2]v.rekeningtransaksi r 
where r.nomor_rekening = tmp.nomor_rekening
	and id_session = $1

-- name: UpdateNamaCabangSaldoPostTrxMultiJournalTempTable-main
update %[1]v.validasi_transaksi_umum tmp
set nama_rekening =r.nama_rekening, 
	kode_cabang = r.kode_cabang,
	saldo = case when r.saldo is null then 0.0 else r.saldo end
from %[2]v.rekeningtransaksi r 
where r.nomor_rekening = tmp.nomor_rekening
	and id_session = $1

-- name: UpdateProdukFloatDitahanPostTrxMultiJournalTempTable-main
update %[1]v.validasi_transaksi_umum tmp
set 
	kode_produk = r.kode_produk, 
	saldo_float = case when r.saldo_float is null then 0.0 else r.saldo_float end, 
	saldo_ditahan = case when r.saldo_ditahan is null then 0.0 else r.saldo_ditahan end
from %[2]v.rekeningliabilitas r 
where r.nomor_rekening = tmp.nomor_rekening
  and id_session = $1 and kode_tipe = '002'

-- name: UpdateSaldoMinimumPostTrxMultiJournalTempTable-main
update %[1]v.validasi_transaksi_umum tmp
set 
  saldo_minimum = case when p.saldo_minimum is null then 0.0 else p.saldo_minimum end
from %[2]v.produk p 
where p.kode_produk = tmp.kode_produk
  and id_session = $1 and kode_tipe ='002'

-- name: UpdateKodeAccountPostTrxMultiJournalTempTable-main
update %[1]v.validasi_transaksi_umum tmp
set kode_account = gi.kode_account
from %[2]v.glinterface gi
where tmp.kode_produk  = gi.kode_produk and gi.kode_interface in ('Saldo_Plus','glnomi')
   and id_session = $1 and kode_tipe ='002'

-- name: CheckMinimumBalancePostTrxMultiJournalTempTable-main
select nomor_rekening 
		from %[1]v.validasi_transaksi_umum tmp
		where id_session = $1
			and jenis_mutasi = 'D'
			and tmp.nominal > tmp.saldo - (tmp.saldo_ditahan + tmp.saldo_float + tmp.saldo_minimum)
			and kode_tipe = '002'

-- name: GetDataFromPostTrxTempTable-main
select kode_tipe,nomor_rekening,jenis_mutasi,nominal,kode_cabang,keterangan,kode_account,kode_valuta
		from %[1]s.validasi_transaksi_umum
		where id_session = $1

-- name: DeletePostTrxTempTable-main
delete from %[1]s.validasi_transaksi_umum
		where id_session = $1

-- name: GetSequenceTempTrxPayroll-main
SELECT nextval('%[1]s.seq_val_transaksi_payroll') 

-- name: InsertValidasiTempTrxPayroll-main
INSERT INTO %[1]s.VALIDASI_TRANSAKSI_PAYROLL (ID_SESSION, NOMOR_REKENING ,NOMINAL, KETERANGAN, IS_VALID)
		VALUES (:ID_SESSION, :NOMOR_REKENING, :NOMINAL, :KETERANGAN, 'T')

-- name: CheckTmpPayrollRekeningValid-main
UPDATE %[1]s.VALIDASI_TRANSAKSI_PAYROLL TMP
		SET IS_VALID ='F'
		, INVALID_MESSAGE = INVALID_MESSAGE || 'nomor rekening tidak terdaftar, '
		WHERE ID_SESSION = $1
		AND NOT EXISTS (SELECT 1 FROM %[2]s.REKENINGTRANSAKSI N WHERE N.NOMOR_REKENING=TMP.NOMOR_REKENING)

-- name: UpdateTmpPayrollNameValid-main
MERGE INTO %[1]s.VALIDASI_TRANSAKSI_PAYROLL TMP
		USING %[2]s.REKENINGTRANSAKSI R 
		ON (R.NOMOR_REKENING = TMP.NOMOR_REKENING)
		WHEN MATCHED THEN UPDATE SET TMP.NAMA_REKENING = R.NAMA_REKENING
		WHERE ID_SESSION = $1

-- name: CheckTmpPayrollRekeningClosed-main
UPDATE %[1]s.VALIDASI_TRANSAKSI_PAYROLL TMP
	SET IS_VALID ='F'
	  , INVALID_MESSAGE = INVALID_MESSAGE || 'nomor rekening sudah ditutup, '
	WHERE ID_SESSION = $1
	  AND EXISTS (SELECT 1 FROM %[2]s.REKENINGTRANSAKSI N WHERE N.NOMOR_REKENING=TMP.NOMOR_REKENING AND N.STATUS_REKENING = 3)

-- name: GetAllTmpTrxPayroll-main
SELECT NOMOR_REKENING, NAMA_REKENING, KETERANGAN,
		NOMINAL, IS_VALID, INVALID_MESSAGE 
		FROM %[1]s.validasi_transaksi_payroll tmp
		WHERE tmp.ID_SESSION = $1

-- name: DeleteTempTrxPayroll-main
DELETE FROM  %[1]s.VALIDASI_TRANSAKSI_PAYROLL TMP WHERE ID_SESSION = $1

-- name: GetSequenceTempTrxMulti-main
SELECT nextval('%[1]s.seq_val_transaksi_umum') 

-- name: InsertValidasiTempTrxMulti-main
insert into %[1]s.validasi_transaksi_umum 
		(id_session, kode_tipe, nomor_rekening, jenis_mutasi, nominal, keterangan, kode_rc, is_valid)
		values ($1, $2, $3, $4, $5, $6, $7, 't')

-- name: CheckTmpTrxMultiCodeTypeValid-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'kode tipe rekening tidak valid, '
		WHERE kode_tipe not in ('001','002')
		and ID_SESSION = $1

-- name: CheckTmpTrxMultiRekExist-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening tidak terdaftar, '
		WHERE ID_SESSION = $1
		AND NOT EXISTS (SELECT 1 FROM %[2]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening)

-- name: CheckTmpTrxMultiRekExistByTypeCode-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'jenis rekening harus tabungan,giro,deposito atau rekening perantara, '
		WHERE ID_SESSION = $1
		AND NOT EXISTS (SELECT 1 FROM %[2]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and n.kode_jenis in ('SAV','CA','RAB','DEP'))

-- name: CheckTmpTrxMultiByMutasi-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'jenis mutasi tidak valid, '
		WHERE jenis_mutasi not in ('D','C')
		and ID_SESSION = $1

-- name: UpdateRekeningTrxTmpTrxMulti-main
MERGE INTO %[1]s.validasi_transaksi_umum tmp
		USING %[2]s.rekeningtransaksi r 
		ON (r.nomor_rekening = tmp.NOMOR_rekening)
		WHEN MATCHED THEN UPDATE SET NAMA_rekening = r.nama_rekening
		, kode_cabang = r.kode_cabang
		, saldo = case when r.saldo is null then 0.0 else r.saldo end
		WHERE ID_SESSION = $1

-- name: UpdateRekeningLiabilitasTmpTrxMulti-main
MERGE INTO %[1]s.validasi_transaksi_umum tmp
			USING %[2]s.rekeningliabilitas r 
			ON (r.nomor_rekening = tmp.NOMOR_rekening)
			WHEN MATCHED THEN UPDATE SET kode_produk =r.kode_produk
			, saldo_float = case when r.saldo_float is null then 0.0 else r.saldo_float end
			, saldo_ditahan = case when r.saldo_ditahan is null then 0.0 else r.saldo_ditahan end
			WHERE ID_SESSION = $1 and tmp.kode_tipe ='002'

-- name: UpdateProdukTmpTrxMulti-main
MERGE INTO %[1]s.validasi_transaksi_umum tmp
		USING %[2]s.produk p 
		ON (p.kode_produk = tmp.kode_produk)
		WHEN MATCHED THEN UPDATE SET saldo_minimum = case when p.saldo_minimum is null then 0.0 else p.saldo_minimum end
		WHERE ID_SESSION = $1 and tmp.kode_tipe ='002'

-- name: UpdateRekTutupByRekTrxTmpMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening sudah ditutup, '
		WHERE ID_SESSION = $1
		AND EXISTS (SELECT 1 FROM %[2]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and n.status_rekening = 3)

-- name: UpdateDormantRekTmpMultiTrx-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening sedang dormant, '
		WHERE ID_SESSION = $1
		AND EXISTS (SELECT 1 FROM %[2]s.rekeningtransaksi n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and n.status_rekening = 2)

-- name: ValidasiNominalTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nominal transaksi tidak valid (nominal <= 0.0)  , '
		WHERE ID_SESSION = $1
		AND nominal <= 0.0

-- name: UpdateRekTutupByLiabilityTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening sedang ditutup, '
		WHERE ID_SESSION = $1
		AND EXISTS (SELECT 1 FROM %[2]s.rekeningliabilitas n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and n.is_sedang_ditutup = 'T')

-- name: UpdateRekBlokirByLiabilityTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening sedang diblokir, '
		WHERE ID_SESSION = $1
		AND EXISTS (SELECT 1 FROM %[2]s.rekeningliabilitas n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and n.is_blokir='T')

-- name: UpdateRekBlokirKreditByLiabilityTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening sedang diblokir kredit, '
		WHERE ID_SESSION = $1
		AND JENIS_MUTASI ='C'
		AND EXISTS (SELECT 1 FROM %[2]s.rekeningliabilitas n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and coalesce(n.is_blokir_kredit,'F') ='T')

-- name: UpdateRekBlokirDebitByLiabilityTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'nomor rekening sedang diblokir debet, '
		WHERE ID_SESSION = $1
		AND JENIS_MUTASI ='D'
		AND EXISTS (SELECT 1 FROM %[2]s.rekeningliabilitas n WHERE n.NOMOR_rekening=tmp.NOMOR_rekening and coalesce(n.is_blokir_debet,'F') ='T')

-- name: ValidasiSaldoEfektifTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'saldo efektif rekening tidak cukup, '
		WHERE ID_SESSION = $1
			AND JENIS_MUTASI ='D'
			AND tmp.nominal > tmp.saldo - (tmp.saldo_ditahan + tmp.saldo_float + tmp.saldo_minimum)
			AND kode_tipe = '002'

-- name: ValidasiKodeRCTmpTrxMulti-main
UPDATE %[1]s.validasi_transaksi_umum tmp
		SET is_valid ='F'
		, invalid_message = invalid_message || 'kode rc tidak valid, '
		WHERE ID_SESSION = $1
		AND NOT EXISTS (SELECT 1 FROM %[2]s.project p WHERE p.project_no=tmp.kode_rc)
		AND tmp.kode_rc is not null

-- name: GetAllTmpMultiTrx-main
SELECT KODE_TIPE, NOMOR_REKENING, NAMA_REKENING,
		KODE_CABANG, JENIS_MUTASI, KETERANGAN, KODE_RC,
		NOMINAL, IS_VALID, INVALID_MESSAGE 
		FROM %[1]s.validasi_transaksi_umum tmp
		WHERE tmp.ID_SESSION = $1

-- name: GetTotalNominalTmpMultiTrx-main
select jenis_mutasi, sum(nominal) as total_nominal 
		from
		%[1]s.validasi_transaksi_umum 
		where id_session = $1
		group by jenis_mutasi

-- name: DeleteTempMultiTrx-main
delete from  %[1]s.validasi_transaksi_umum tmp where id_session= $1

-- name: GetRekeningLayanan-main
SELECT regl.NOMOR_REKENING_LAYANAN , regw.NOMOR_REKENING_TUJUAN, regw.SALDO_MINIMAL
		FROM %[1]v.REGISTERLAYANAN regl
		INNER JOIN RegisterAutoSweep regw ON regl.ID_REGISTER = regw.ID_REGISTER
		WHERE regl.STATUS_REGISTER_LAYANAN = 'A' AND regw.NOMOR_REKENING_TUJUAN = $1

-- name: UpdateNilaiKurs-main
update %[1]s.detiltransaksi dt
		set nilai_kurs_manual = 1,
			nilai_ekuivalen = nilai_mutasi
		where dt.kode_valuta = :DEFAULT_VALUTA 
			and exists(
				select 1 
				from %[1]s.pendingtransactionjournal ptj
				where ptj.id_transaksi = dt.id_transaksi
					and %[2]s
			)

-- name: UpdateNilaiKurs-filter-IdTransaksiNotZero
dt.id_transaksi = :ID_TRANSAKSI

-- name: UpdateNilaiKurs-filter-IdTransaksiZero
dt.flag_sent is null

-- name: UpdateNilaiEkuivalen-main
update %[1]s.detiltransaksi dt
		set nilai_ekuivalen = nilai_mutasi * nilai_kurs_manual
		where dt.kode_valuta = $1 
			and exists(
				select 1 
				from %[1]s.pendingtransactionjournal ptj
				where ptj.id_transaksi = dt.id_transaksi
					and %[2]s
			)

-- name: UpdateNilaiEkuivalen-filter-IdTransaksiNotZero
dt.id_transaksi = $2

-- name: UpdateNilaiEkuivalen-filter-IdTransaksiZero
dt.flag_sent is null

-- name: GetTrxDetailByUidkey-main
select i.uidkey, t.id_transaksi ,d.kode_jurnal, t.kode_cabang_transaksi, rt.nama_rekening 
	,d.nomor_rekening ,d.keterangan ,d.kode_valuta, d.jenis_mutasi,
		case
			when d.jenis_mutasi = 'D' then 'Debet'
			when d.jenis_mutasi = 'C' then 'Kredit'
		end as jenis_mutasi_desc,
		d.nilai_mutasi ,d.nilai_kurs_manual ,d.nilai_ekuivalen ,d.kode_rc
	from %[2]v.infouid i
		inner join %[2]v.detiltransaksi d on d.id_transaksi = i.keyint
		inner join %[2]v.transaksi t on t.id_transaksi = d.id_transaksi
		inner join %[2]v.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening
	where action_type = $1
	%[1]v
	order by d.id_transaksi ,d.id_detil_transaksi

-- name: GetTrxDetailByUidkey-filter-TrxTypeDefault
and uidkey = $2

-- name: GetTrxDetailByUidkey-filter-TrxType
and uidkey in ($2, $3)

-- name: GetHistTrxDetailByUidkey-main
select i.uidkey, t.id_transaksi ,d.kode_jurnal, t.kode_cabang_transaksi, rt.nama_rekening 
	,d.nomor_rekening ,d.keterangan ,d.kode_valuta, d.jenis_mutasi,
		case
			when d.jenis_mutasi = 'D' then 'Debet'
			when d.jenis_mutasi = 'C' then 'Kredit'
		end as jenis_mutasi_desc,
		d.nilai_mutasi ,d.nilai_kurs_manual ,d.nilai_ekuivalen ,d.kode_rc
	from %[2]v.infouid i
		inner join %[2]v.histdetiltransaksi d on d.id_transaksi = i.keyint
		inner join %[2]v.histtransaksi t on t.id_transaksi = d.id_transaksi
		inner join %[2]v.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening
	where action_type = $1
	%[1]v
	order by d.id_transaksi ,d.id_detil_transaksi

-- name: GetHistTrxDetailByUidkey-filter-TrxTypeDefault
and uidkey = $2

-- name: GetHistTrxDetailByUidkey-filter-TrxType
and uidkey in ($2, $3)

-- name: GetDetailTrxDetail-main
select 
			dt.id_detil_transaksi,
			t.kode_transaksi,
			t.keterangan as keterangan_transaksi, 
			dt.keterangan as keterangan_detil,
			dt.kode_valuta kode_valuta,
			dt.nilai_kurs_manual,
			dt.nilai_ekuivalen,
			to_char(t.tanggal_transaksi, 'yyyy-mm-dd') tanggal_transaksi,
			to_char(t.tanggal_input, 'yyyy-mm-dd') tanggal_input,
			to_char(t.jam_input, 'HH24-MI-SS') jam_input,
			'' as nomor_batch,
			'' as blok_jurnal,
			t.nomor_referensi,
			t.user_input,
			t.user_otorisasi,
			t.kode_cabang_transaksi,
			t.journal_no
		from %[1]s.detilTransaksi dt
		inner join %[1]s.transaksi t on dt.id_transaksi = t.id_transaksi
		where dt.id_detil_transaksi = $1

-- name: GetDetailTrxDetailHist-main
select 
			dt.id_detil_transaksi,
			t.kode_transaksi,
			t.keterangan as keterangan_transaksi, 
			dt.keterangan as keterangan_detil,
			dt.kode_valuta kode_valuta,
			dt.nilai_kurs_manual,
			dt.nilai_ekuivalen,
			to_char(t.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
			to_char(t.tanggal_input, 'DD-MM-YYYY') tanggal_input,
			to_char(t.jam_input, 'HH24-MI-SS') jam_input,
			'' as nomor_batch,
			'' as blok_jurnal,
			t.nomor_referensi,
			t.user_input,
			t.user_otorisasi,
			t.kode_cabang_transaksi,
			t.journal_no
		from %[1]s.histdetiltransaksi dt
		inner join %[1]s.histtransaksi t on dt.id_transaksi = t.id_transaksi
		where dt.id_detil_transaksi = :id_detil_transaksi

-- name: GetTrxDetailByNoref-main
select 
			d.nomor_referensi, t.id_transaksi ,d.kode_jurnal, t.kode_cabang_transaksi, rt.nama_rekening,
			d.nomor_rekening, d.keterangan, d.kode_valuta, d.jenis_mutasi,
			case
				when d.jenis_mutasi = 'D' then 'Debet'
				when d.jenis_mutasi = 'C' then 'Kredit'
			end as jenis_mutasi_desc,
			d.nilai_mutasi, d.nilai_kurs_manual, d.nilai_ekuivalen, d.kode_rc
		from %[1]s.detiltransaksi d
			inner join %[1]s.transaksi t on t.id_transaksi = d.id_transaksi
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening
		where d.nomor_referensi = $1
		order by d.id_transaksi, d.id_detil_transaksi

-- name: GetHistTrxDetailByNoref-main
select 
			d.nomor_referensi, t.id_transaksi ,d.kode_jurnal, t.kode_cabang_transaksi, rt.nama_rekening,
			d.nomor_rekening, d.keterangan, d.kode_valuta, d.jenis_mutasi,
			case
				when d.jenis_mutasi = 'D' then 'Debet'
				when d.jenis_mutasi = 'C' then 'Kredit'
			end as jenis_mutasi_desc,
			d.nilai_mutasi, d.nilai_kurs_manual, d.nilai_ekuivalen, d.kode_rc
		from %[1]s.histdetiltransaksi d
			inner join %[1]s.histtransaksi t on t.id_transaksi = d.id_transaksi
			inner join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening
		where d.nomor_referensi = $1
		order by d.id_transaksi, d.id_detil_transaksi

-- name: CreateTransaksiExtInfo-main
insert into %[1]v.transaksiextinfo
		(
			id_transaksi, kode_sistem_ext, rekening_source, rekening_benef, id_transaksi_ext,
			nominal, nominal_biaya, kode_transaksi_ext
		) 
		values 
		(
			$1, $2, $3, $4, $5, $6, $7, $8
		)

-- name: InsertTransaksiTransitKas-main
insert into %[1]s.transaksitransitkas 
		(id_detil_transaksi, nominal_selisih_kurs, kode_posting) 
		values ($1, $2, $3)

-- name: InsertTransfer-main
insert into %[1]s.transfer
		(jenis_transfer, nama_penerima, nama_pengirim, nomor_rekening_penerima, id_transfer, kode_bank, id_detil_transaksi, id_transaksi, 
			nomor_rekening_pengirim, status_pengirim, jenis_penduduk, kode_account_titipan, kode_cabang_titipan, jenis_pengirim, jenis_penerima, 
			jenispenduduk_pengirim, jenispenduduk_penerima, noidentitas_pengirim, noidentitas_penerima, alamat_pengirim, alamat_penerima, 
			jenis_channel, sandikota_pengirim, sandikota_penerima, kode_tx_skn, kode_tx_rtgs, nomor_referensi_retur, jenisnasabah_pengirim, 
			jenisnasabah_penerima, nomor_nasabah, kode_prioritas_rtgs, tipe_message_rtgs, nomor_account_bank, no_rek_intern_bi, 
			nama_rek_intern_bi, pvp_cp_sender_correspondent, pvp_cp_receiver_correspondent, pvp_ordering_institution, pvp_beneficiary_institution)
		values
		($1, $2, $3, $4,  nextval('%[1]s.seq_transfer'), $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, 
		$15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38)

-- name: RptTrialBalance-main
select ac.account_code
    , ac.account_name, ai.branch_code
    , 
		%[4]s,
		case 
        when db.balancecumulative is null then 0.0 
        else db.balancecumulative 
    end as balance
from %[1]s.account ac
    inner JOIN %[1]s.accountinstance ai on ai.account_code = ac.account_code
    left outer join %[1]s.enum_varchar etype on (etype.enum_value = ac.account_type and etype.enum_name='eAccountType')
    left outer join (
   SELECT 
		ai.accountinstance_id, TO_DATE($1, 'YYYY-MM-DD') - 1 as p_date,
		ai.account_code, ai.branch_code, ai.currency_code, 
		case when db.debit is null then 0.0 else db.debit end as debit, 
		case when db.debit_ekuiv is null then 0.0 else db.debit_ekuiv end as debit_ekuiv, 
		case when db.credit is null then 0.0 else db.credit end as credit, 
		case when db.credit_ekuiv is null then 0.0 else db.credit_ekuiv end as credit_ekuiv, 
		case when db.balance is null then 0.0 else db.balance end as balance, 
		case when db.balance_ekuiv is null then 0.0 else db.balance_ekuiv end as balance_ekuiv, 
		case when db.balancecumulative is null then 0.0 else db.balancecumulative end as balancecumulative, 
		case when db.balancecumulative_ekuiv is null then 0.0 else db.balancecumulative_ekuiv end as balancecumulative_ekuiv, 
		case when db.balance_ratarata is null then 0.0 else db.balance_ratarata end as balance_ratarata, 
		case when db.balance_ratarata_ekuiv is null then 0.0 else db.balance_ratarata_ekuiv end as balance_ratarata_ekuiv
			FROM %[1]s.accountinstance ai
			LEFT OUTER JOIN
				(SELECT d1.* FROM %[1]s.dailybalance d1,
					(SELECT accountinstance_id, MAX(datevalue) as maxdate 
						FROM %[1]s.dailybalance WHERE datevalue <= TO_DATE($2, 'YYYY-MM-DD') - 1
						GROUP BY accountinstance_id) d2
					WHERE d1.accountinstance_id = d2.accountinstance_id
						AND d1.datevalue = d2.maxdate) db
				ON ai.accountinstance_id = db.accountinstance_id
			) db ON ai.accountinstance_id = db.accountinstance_id
    left outer join (
 	SELECT 
					ai.accountinstance_id, TO_DATE($2, 'YYYY-MM-DD')  as p_date,
					ai.account_code, ai.branch_code, ai.currency_code, 
					case when db.debit is null then 0.0 else db.debit end as debit, 
					case when db.debit_ekuiv is null then 0.0 else db.debit_ekuiv end as debit_ekuiv, 
					case when db.credit is null then 0.0 else db.credit end as credit, 
					case when db.credit_ekuiv is null then 0.0 else db.credit_ekuiv end as credit_ekuiv, 
					case when db.balance is null then 0.0 else db.balance end as balance, 
					case when db.balance_ekuiv is null then 0.0 else db.balance_ekuiv end as balance_ekuiv, 
					case when db.balancecumulative is null then 0.0 else db.balancecumulative end as balancecumulative, 
					case when db.balancecumulative_ekuiv is null then 0.0 else db.balancecumulative_ekuiv end as balancecumulative_ekuiv, 
					case when db.balance_ratarata is null then 0.0 else db.balance_ratarata end as balance_ratarata, 
					case when db.balance_ratarata_ekuiv is null then 0.0 else db.balance_ratarata_ekuiv end as balance_ratarata_ekuiv
				FROM %[1]s.accountinstance ai
				LEFT OUTER JOIN
					(SELECT d1.* FROM %[1]s.dailybalance d1,
						(SELECT accountinstance_id, MAX(datevalue) as maxdate 
						FROM %[1]s.dailybalance WHERE datevalue <= TO_DATE($2, 'YYYY-MM-DD')
						GROUP BY accountinstance_id) d2
					WHERE d1.accountinstance_id = d2.accountinstance_id
						AND d1.datevalue = d2.maxdate) db
				ON ai.accountinstance_id = db.accountinstance_id
			) db2 ON ai.accountinstance_id = db2.accountinstance_id
    left outer join (SELECT 
					ai.accountinstance_id, 
					TO_DATE($1, 'YYYY-MM-DD')  AS begindate, 
					TO_DATE($2, 'YYYY-MM-DD')  AS enddate,
					ai.account_code, 
					ai.branch_code, 
					ai.currency_code,
					ai.balance_sign,
					COALESCE(db.debit, 0.0) AS debit,
					COALESCE(db.debit_ekuiv, 0.0) AS debit_ekuiv,
					COALESCE(db.credit, 0.0) AS credit,
					COALESCE(db.credit_ekuiv, 0.0) AS credit_ekuiv,
					COALESCE(db.balance, 0.0) AS balance,
					COALESCE(db.balance_ekuiv, 0.0) AS balance_ekuiv
				FROM %[1]s.accountinstance ai
				LEFT OUTER JOIN (
					SELECT 
						d.accountinstance_id,
						SUM(d.credit) AS credit,
						SUM(d.credit_ekuiv) AS credit_ekuiv,
						SUM(d.debit) AS debit,
						SUM(d.debit_ekuiv) AS debit_ekuiv,
						SUM(d.balance) AS balance,
						SUM(d.balance_ekuiv) AS balance_ekuiv
					FROM %[1]s.dailybalance d
					WHERE d.datevalue >= TO_DATE($1, 'YYYY-MM-DD') 
					AND d.datevalue <= TO_DATE($2, 'YYYY-MM-DD') 
					GROUP BY d.accountinstance_id
					) db ON ai.accountinstance_id = db.accountinstance_id
				) mut ON ai.accountinstance_id = mut.accountinstance_id
    left outer join
        ( SELECT
                ai.accountinstance_id
                , sum(amount_debit) as debit
                , sum(amount_credit) as credit
                , sum((amount_credit-amount_debit) * ai.balance_sign) as balance
            FROM %[1]s.journal j
                    inner join %[1]s.journalitem ji on j.journal_no = ji.fl_journal
                    inner join %[1]s.accountinstance ai on ai.accountinstance_id = ji.accountinstance_id
            WHERE j.journal_date = TO_DATE($2, 'YYYY-MM-DD') 
                and ai.branch_code = $3
            GROUP BY ai.accountinstance_id            
            UNION            
            SELECT
                    ai.accountinstance_id
                    , sum(amount_debit) as debit
                    , sum(amount_credit) as credit
                    , sum(amount_credit-amount_debit) as balance
            FROM %[1]s.journal j
                    inner join %[1]s.journalitem ji on j.journal_no = ji.fl_journal
                    inner join %[1]s.accountinstance ai1 on ai1.accountinstance_id = ji.accountinstance_id
                    inner join %[1]s.accountinstance ai on ai.accountinstance_id = ai1.fl_cpa_accountinstance
            WHERE j.journal_date =  TO_DATE($2, 'YYYY-MM-DD') 
                and ai.branch_code = $3
            GROUP BY ai.accountinstance_id
        ) j on ai.accountinstance_id = j.accountinstance_id
    where ac.is_detail = 'T'
    and ac.account_type in ('A', 'L' , 'E', 'I','X','M')    
    and ai.branch_code = $3
    and ai.currency_code = $4
		%[2]s
		%[3]s
    order by branch_code, account_code

-- name: RptTrialBalance-filter-HasValueOnly
and (
    case when mut.debit is null then 0.0 else mut.debit end <> 0.0
    or case when mut.credit is null then 0.0 else mut.credit end <> 0.0
    or case when j.debit is null then 0.0 else j.debit end <> 0.0
    or case when j.credit is null then 0.0 else j.credit end <> 0.0
    or case when j.balance is null then 0.0 else j.balance end <> 0.0
    or case when db.balancecumulative is null then 0.0 else db.balancecumulative end <> 0.0
    or case when db2.balancecumulative is null then 0.0 else db2.balancecumulative end <> 0.0
)

-- name: RptTrialBalance-filter-CompanyCode
AND ac.company_code = $5

-- name: GetAccDayByPeriode-main
select datevalue from %[1]s.accountingday where periode_status = 'A'

-- name: GetAccount-main
select rt.nama_rekening
		, rt.saldo
		, rt.status_rekening
		, rt.kode_jenis
		, rt.saldo - (coalesce(rl.saldo_float, 0.0) + coalesce(rl.saldo_ditahan, 0.0) + coalesce(p.saldo_minimum, 0.0)) as saldo_efektif
		, rt.saldo - (coalesce(rl.saldo_float, 0.0) + coalesce(rl.saldo_ditahan, 0.0)) as saldo_override
		, p.kode_produk
		, p.nama_produk
		, rl.jenis_rekening_liabilitas
	from %[1]s.rekeningliabilitas rl
		, %[1]s.rekeningtransaksi rt
		, %[1]s.produk p
	where rl.nomor_rekening = rt.nomor_rekening
		and rl.kode_produk = p.kode_produk
		and rl.nomor_rekening = $1

-- name: GetDc-main
select * 
	from %[1]s.jenisdc 
	where network_id = $1

-- name: GetFeeDc-main
select 
		fee_nominal feenom, id_feetransaksidc feeid 
	from %[1]s.feetransaksidc 
	where jenis_dc = $1 and tipe_transaksi = $2

-- name: FeeGlCredit-main
select * 
	from %[1]s.rekeningdc 
	where 
		jenis_dc = $1 and 
		trx_code = $2 and
		is_fee = 'T' and 
		mnemonic = 'C'

-- name: GetPairGl-main
select * 
	from %[1]s.rekeningdc 
	where 
		jenis_dc = $1 and 
		trx_code = $2 and 
		is_fee = 'F' and 
		mnemonic = $3

-- name: GetVa-main
select 
		va.accountname
		, coalesce(vap.admcosttype, 'F') as admcosttype
		, coalesce(vap.admcost, 0) as admcost
		, vap.paymenttype
		, vap.nomor_rekening
		, rt.kode_cabang
	from %[1]s.virtualaccount va
		, %[1]s.girovaproduct vap
		, %[1]s.rekeningtransaksi rt
	where vap.productid = va.productid
		and vap.nomor_rekening = rt.nomor_rekening
		and va.virtualaccountid = $1

-- name: GetVaBilling-main
select vbi.billamount
		, vbi.paymentstatus
		, vbi.billingid
		, gvp.status
	from %[1]s.vabillingitem vbi
		, %[1]s.vabillingsession vbs
		, %[1]s.girovaproduct gvp
	where vbs.sessionid = vbi.sessionid
		and gvp.productid = vbi.productid
		and vbs.status = 'A'
		and vbi.virtualaccountid = $1

-- name: GetVaPaid-main
select sum(billtotalpayment) as totalpayment
	from %[1]s.detiltransva
	where billingid = $1

-- name: GetVaHistPaid-main
select sum(billtotalpayment) as totalpayment
	from %[1]s.histdetiltransva
	where billingid = $1

-- name: GetDataAgentLiabilitas-main
select rl.nomor_rekening, rl.kode_marketing_current, kmc.agentname as kode_marketing_current_desc, 
    rl.kode_marketing_pertama, kmp.agentname as kode_marketing_pertama_desc,
    rl.kode_marketing_referensi, kmr.agentname as kode_marketing_referensi_desc
from %[1]s.rekeningliabilitas rl
    left join %[1]s.fundingagent kmc on kmc.agentcode = rl.kode_marketing_current
    left join %[1]s.fundingagent kmp on kmp.agentcode = rl.kode_marketing_pertama
    left join %[1]s.fundingagent kmr on kmr.agentcode = rl.kode_marketing_referensi
where rl.nomor_rekening = $1

-- name: GetListAccountOfficer-main
select agentcode ,agentname ,statusactive 
from %[1]s.fundingagent f 
where %[2]s

-- name: GetListAccountOfficer-filter-Default
1=1

-- name: GetListAccountOfficer-filter-AgentCode
and f.agentcode = $1

-- name: GetListAccountOfficer-filter-AgentName
and f.agentname LIKE  '%' || $2 || '%'

-- name: GetListAccountOfficer-filter-StatusActive
and f.statusactive = $3

-- name: GetFundingAgent-main
select agentcode ,agentname ,statusactive
from %[1]s.fundingagent
where agentcode = $1

-- name: GetInfoUid-main
select uidkey ,to_char(tanggal_input, 'yyyy-mm-dd') as tanggal_input ,action_type ,user_input ,user_otor, keterangan ,info1 ,info2 
    ,kode_cabang, keyint, 'T' as is_exists, is_reversed
from %[1]s.infouid i 
where uidkey = $1
    and action_type = $2

-- name: GetInfoUidForRev-main
select uidkey ,to_char(tanggal_input, 'yyyy-mm-dd') as tanggal_input ,action_type ,user_input ,user_otor, keterangan ,info1 ,info2 
    ,kode_cabang, keyint, 'T' as is_exists, is_reversed
from %[1]s.infouid i 
where uidkey = $1
    and action_type = $2

-- name: InsertInfoUid-main
insert into %[1]s.infouid 
(uidkey, tanggal_input, action_type, user_input, user_otor, keterangan, info1, info2, kode_cabang, keyint)
values($1, current_timestamp, $2, $3, $4, $5, $6, $7, $8, $9)

-- name: GetPod-main
select to_char(nilai_parameter_tanggal, 'YYYY-MM-DD') as today
from %[1]s.parameterglobal
where kode_parameter = $1

-- name: GetAccountNo-main
select counter from %[1]s.rekeninggenerator 
where grup = 'REKENING' and kode_generator = $1

-- name: GetAccountNo-sqlInsertRG
insert into %[1]s.rekeninggenerator 
(kode_generator ,counter ,grup)
values($1 ,1 ,'REKENING')

-- name: GetAccountNo-sqlUpdateRG
update %[1]s.rekeninggenerator 
set counter = counter + 1 
where grup = 'REKENING' and kode_generator = $1

-- name: GetCifPortfolioByNomorNasabah-main
select 
rl.nomor_nasabah, rt.nomor_rekening, rt.nama_rekening, rt.kode_valuta, rt.kode_cabang, c.nama_cabang, abs(round(rt.saldo, 2)) as saldo, 
case 
    when rt.status_rekening = 1 then 'Aktif'
    when rt.status_rekening = 2 then 'Dormant'
    when rt.status_rekening = 3 then 'Tutup'
end as status_rekening,
case 
    when rl.jenis_rekening_liabilitas = 'T' then 'Tabungan'
    when rl.jenis_rekening_liabilitas = 'R' then 'Rencana'
    when rl.jenis_rekening_liabilitas = 'G' then 'Giro'
    when rl.jenis_rekening_liabilitas = 'D' then 'Deposito'
end as jenis_rekening,
p.kode_produk, p.nama_produk
from %[1]s.rekeningtransaksi rt
inner join %[2]s.cabang c on rt.kode_cabang = c.kode_cabang
inner join %[1]s.rekeningcustomer rc on rc.nomor_rekening = rt.nomor_rekening
inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rc.nomor_rekening
inner join %[1]s.produk p on p.kode_produk = rl.kode_produk
where rc.nomor_nasabah = $1
order by rt.status_rekening ,rt.nomor_rekening

-- name: GetIDIdv-main
select nextval('%[1]s.seq_individu') as id_idv 

-- name: GetIDRegLayanan-main
select nextval('%[1]s.seq_registerlayanan') as id 

-- name: GetListParameterBiaya-main
select kode_akun_ssl as account_code ,nominal_biaya
from %[1]s.parameterbiaya p 
    inner join %[1]s.grupparameterbiaya gpb on gpb.id_grup_biaya = p.id_grup_biaya 
where gpb.kode_grup_biaya = $1

-- name: GetTransaksiID-main
select nextval('%[1]s.seq_transaksi') as id_transaksi 

-- name: GetTransaksiDetailID-main
select nextval('%[1]s.seq_detiltransaksi') as id_detil_transaksi 

-- name: GetCustomIdgen-main
select idcode,
case 
    when last_id is null then 0
    else ( nextval('%[1]s.seq_custom_idgen') - last_id + 1) 
end as last_idgen 
from %[1]s.custom_idgen
where idname = $1 and idCode = $2

-- name: InsertCustomIdgen-main
insert into %[1]s.custom_idgen
(idname, idcode, last_id, locked)
values ($1, $2,  nextval('%[1]s.seq_custom_idgen'), $3)

-- name: GetGLInterface-main
select gi.kode_interface ,gi.kode_account ,gi.kode_produk
from %[1]s.glinterface gi
where gi.kode_interface = $1
and gi.kode_produk = $2

-- name: GetGlAccount-main
select * from %[1]s.glaccount where nomor_rekening = $1

-- name: GetGlobalGlInterface-main
select kode_account, kode_interface, nama_account
from %[1]s.globalglinterface 
where kode_interface = $1

-- name: InsertRakOtomatis-main
insert into %[1]s.journalitem 
(fl_journal, journalitem_no, fl_account, amount_debit, amount_credit, description, branch_code, valuta_code, fl_project_no, 
    userid_create, userid_posting, dateposting, datecreate, journalitemtype, nilai_kurs, accountinstance_id, subaccountcode, subaccountname)
select ubitems.fl_journal,  nextval('%[1]s.seq_channeljournalitemno'), $1,  
    case when ubitems.amount_credit > ubitems.amount_debit then ubitems.amount_credit - ubitems.amount_debit else 0.0 end,  
    case when ubitems.amount_debit > ubitems.amount_credit then ubitems.amount_debit - ubitems.amount_credit else 0.0 end, 
    'TRANSAKSI RAK OTOMATIS', ubitems.branch_code, ubitems.valuta_code, null, 
    $2, $3, current_timestamp, current_timestamp, 'N', 1, ai.accountinstance_id, $1 || '-' || ubitems.branch_code || '-' || ubitems.valuta_code, $4
from
(
    select ji.fl_journal, ji.valuta_code, ji.branch_code, sum(amount_debit) as amount_debit, sum(amount_credit) as amount_credit
    from %[1]s.journalitem ji
        inner join %[1]s.account acc on acc.account_code = ji.fl_account
    where ji.fl_journal = $5
        and acc.account_type <> 'M'
    group by ji.fl_journal, ji.valuta_code, ji.branch_code
) ubitems
    inner join %[1]s.accountinstance ai on (ai.branch_code = ubitems.branch_code) and (ai.currency_code = ubitems.valuta_code)
where abs(ubitems.amount_debit - ubitems.amount_credit) > 0.001
    and ai.account_code = $1

-- name: GetGLList-main
select acc.account_code, acc.account_name, acc.account_type  
		from %[1]s.glaccount g
			inner join %[1]s.accountinstance ai on g.accountinstance_id = ai.accountinstance_id and ai.currency_code = $1 and ai.branch_code = $2 and ai.isbolehinput = 'T'
			inner join %[1]s.account acc on acc.account_code = ai.account_code
		where %[2]s
		group by acc.account_code, acc.account_name, acc.account_type
		order by acc.account_code, acc.account_name, acc.account_type

--name: GetGLList-filter-Default
1=1

--name: GetGLList-filter-Keyword
and (acc.account_code like $3 or acc.account_name like $3)

-- name: GetGLJurnal-main
SELECT 
    j.journal_no, 
    j.serial_no, 
    j.is_posted, 
    j.journal_state, 
    TO_CHAR(j.transaction_date, 'YYYY-MM-DD') as transaction_date, 
    TO_CHAR(j.journal_date, 'YYYY-MM-DD') as journal_date, 
    TO_CHAR(j.journal_date_posting, 'YYYY-MM-DD') as journal_date_posting, 
    TO_CHAR(j.journal_date_last_modified, 'YYYY-MM-DD') as journal_date_last_modified, 
    j.branch_code, 
    j.description, 
    j.userid_create, 
    j.userid_posting, 
    j.userid_last_modified 
    , TO_CHAR(t.tanggal_input, 'YYYY-MM-DD') AS tanggal_input
    , TO_CHAR(t.tanggal_otorisasi, 'YYYY-MM-DD') AS tanggal_otorisasi
    , t.status_otorisasi
FROM %[1]s.journal j
LEFT JOIN ibcoreprod.transaksi t ON t.journal_no = j.journal_no
WHERE %[2]s
ORDER BY journal_date desc 
OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY

-- name: GetGLJurnal-filter-Default
(1 = $3 or branch_code = $4 )
    AND journal_date >= TO_DATE($5, 'YYYY-MM-DD') 
	AND journal_date < TO_DATE($6, 'YYYY-MM-DD')

-- name: GetGLJurnal-filter-JurnalNoGTZero
and journal_no LIKE ( '%%' || $7 || '%%' )

-- name: DetailGLJurnal-main
SELECT 
    j.journal_no
    , TO_CHAR(j.transaction_date, 'YYYY-MM-DD') as transaction_date
    , TO_CHAR(j.journal_date, 'YYYY-MM-DD') as journal_date
    , TO_CHAR(j.journal_date_posting, 'YYYY-MM-DD') as journal_date_posting
    , TO_CHAR(j.journal_date_last_modified, 'YYYY-MM-DD') as journal_date_last_modified
    , j.branch_code 
    , j.description 
    , j.userid_create 
    , j.userid_posting 
    , j.is_posted 
    , j.journal_state 
    , TO_CHAR(t.tanggal_input, 'YYYY-MM-DD') AS tanggal_input
    , TO_CHAR(t.tanggal_otorisasi, 'YYYY-MM-DD') AS tanggal_otorisasi
    , t.status_otorisasi
FROM %[1]s.journal j 
LEFT JOIN ibcoreprod.transaksi t ON t.journal_no = j.journal_no
WHERE j.journal_no = $1

-- name: DetailGLJurnalItem-main
SELECT 
    JournalItem_No as Nomor,
    ai.Account_Code as Kode_Rekening,
    a.Account_Name as Nama_Rekening,
    ai.branch_code as Kode_Kantor,
    bl.Nama_Cabang as Nama_Kantor,
    ai.Currency_Code as Valuta,
    Kode_Kurs,
    Nilai_Kurs,
    p.name as Nama_RC,
    jb.ID_JournalBlock as No_Blok_Jurnal,
    jb.subsystemcode,
    subaccountcode,
    subaccountname,
    amount_debit as Debet,
    (amount_debit * nilai_kurs) as debet_ekuivalen,
    amount_credit as Kredit,
    (amount_credit * nilai_kurs) as kredit_ekuivalen,
    j.userid_create as User_ID_Create,
    ji.description as Keterangan
FROM %[1]s.JournalItem ji
INNER JOIN %[1]s.journal j on j.journal_no = ji.fl_journal 
INNER JOIN %[1]s.accountinstance ai on ai.accountinstance_id = ji.accountinstance_id 
INNER JOIN %[1]s.account a on a.account_code = ai.account_code 
LEFT JOIN %[1]s.journalblock jb on jb.id_journalblock = ji.id_journalblock 
LEFT JOIN %[1]s.branchlocation bl on bl.Kode_Cabang = ai.branch_code
LEFT JOIN %[1]s.project p on p.project_no = ji.rc_code 
WHERE ji.fl_journal = $1 
    AND ji.fl_journal not LIKE 'POD.%%'
order by nomor

-- name: DetailGLJurnalBlock-main
SELECT 
    SubSystemCode as Kode_Sub_Sistem,
    class_id,
    key_id,
    ID_JournalBlock,
    keterangan
FROM %[1]s.journalblock j 
WHERE journal_no = $1 
    AND journal_no not LIKE 'POD.%%'
ORDER BY ID_JournalBlock

-- name: CoaList-main
SELECT 
    account_code
    , account_name
    , account_type
    , account_level 
    , is_detail
    , ev.enum_description AS account_type_desc
FROM %[1]s.account 
LEFT JOIN ibcoreprod.enum_varchar ev ON ev.enum_value = account.account_type AND ev.enum_name = 'eAccountType'
ORDER BY account_code

-- name: CoaListMapping-main
SELECT m.account_code_lama, m.account_code_baru, a.account_name
FROM %[1]s.mappingcoa m, %[1]s.account a
WHERE m.account_code_baru = a.account_code
ORDER BY m.account_code_lama

-- name: GetGlinterface-main
select kode_interface ,kode_account ,deskripsi ,id_glinterface ,kode_produk 
from %[1]s.glinterface g
where kode_produk = $1
    and kode_interface = $2

-- name: GetGlinterfaceTabungan-main
select id_glinterface, kode_interface, deskripsi, kode_account  
from %[1]s.glinterface g
where kode_produk = $1

-- name: InsertHistUbahData-main
insert into %[1]s.historiubahdata
(id_histori, nomor_rekening, tanggal_ubah, user_ubah, nama_class, keterangan_ubah)
values( nextval('%[1]s.seq_historiubahdata'), $1, current_timestamp, $2, $3, $4)

-- name: CreateDataIndividu-main
insert into %[1]s.individu
(alamat_rumah_rt, alamat_rumah_provinsi, kode_pekerjaan, jenis_identitas_lain, nama_lengkap, 
    tempat_lahir, status_perkawinan, nomor_identitas, nomor_identitas_lain, alamat_rumah_rw, 
    alamat_rumah_kecamatan, alamat_rumah_kode_pos, telepon_hp_nomor, id_individu, tanggal_lahir, 
    kewarganegaraan, alamat_rumah_kota_kabupaten, jenis_identitas, alamat_rumah_jalan, 
    alamat_rumah_kelurahan, status_pep, jenis_kelamin, kode_negara_asal)
values 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
    
-- name: GetInfoTTRByIdTrx-main
select nomor_rekening, id_transaksi, to_char(tanggal_proses, 'yyyy-mm-dd') as tanggal_proses, nilai_transaksi, keterangan, user_otorisasi, userotorisasi, userinput
from %[1]s.infopenutupanrekening
where id_transaksi = $1

-- name: DelInfoTTR-main
delete from %[1]s.infopenutupanrekening
where nomor_rekening = $1

-- name: InsertInfoPenutupanRekening-main
insert into %[1]s.infopenutupanrekening
(nomor_rekening, id_transaksi ,tanggal_proses ,nilai_transaksi, keterangan ,userinput ,userotorisasi)
values
($1, $2 ,$3 ,$4, $5 ,$6 ,$7)

-- name: InsertJournal-main
insert into %[1]s.journal
(journal_no, journal_date_posting, description, userid_create, userid_posting, is_posted, journal_type, journal_date, reference_no, reference_no_fwd,
    lastitemno, is_partlychecked, branch_code, transaction_date, journal_state, serial_no, periode_code, fl_accountingperiode, fl_accountingyear)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, current_timestamp, $14, $15, $16, $17, $18)

-- name: InsertJournalItem-main
insert into %[1]s.journalitem
(fl_journal, journalitem_no, fl_account, amount_debit, amount_credit, description, branch_code, valuta_code, userid_create, journalitemstatus, 
    userid_posting, dateposting, datecreate, journalitemtype, nilai_kurs, subaccountcode, accountinstance_id, subaccountname
)
values ($1,  nextval('%[1]s.seq_channeljournalitemno'), $2, $3, $4, $5, $6, $7, $8, $9, $10, current_timestamp, current_timestamp, $11, $12, $13, $14, $15)

-- name: InsertPendingTrxJournal-main
insert into %[1]s.pendingtransactionjournal 
(id_transaksi) 
values ($1)

-- name: CheckValiditasDataJournal-main
select dt.id_detil_transaksi, dt.id_transaksi, dt.nomor_rekening, dt.jenis_mutasi,dt.nilai_mutasi, dt.kode_tx_class, t.nomor_seri
from %[1]s.detiltransaksi dt
    inner join %[1]s.transaksi t on dt.id_transaksi = t.id_transaksi
    inner join %[1]s.pendingtransactionjournal ptj on ptj.id_transaksi = t.id_transaksi
where %[2]s
    and (
        dt.kode_account is null
        or dt.kode_cabang is null
        or dt.kode_valuta is null
        or (
            dt.kode_valuta = $1
            and (dt.nilai_kurs_manual <= 0 or dt.nilai_kurs_manual is null)
        )
    )

-- name: CheckValiditasDataJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $2

-- name: CheckValiditasDataJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: CheckAccountInstanceJournal-main
select dt.id_detil_transaksi, dt.id_transaksi, dt.nomor_rekening, dt.kode_account, dt.kode_cabang, dt.kode_valuta
from %[1]s.detiltransaksi dt
    inner join %[1]s.transaksi t on dt.id_transaksi = t.id_transaksi
    inner join %[1]s.pendingtransactionjournal ptj on ptj.id_transaksi = t.id_transaksi
where %[2]s
and not exists (
    select 1 
    from %[1]s.accountinstance ai
    where ai.account_code = dt.kode_account
        and ai.branch_code = dt.kode_cabang
        and ai.currency_code = dt.kode_valuta
)

-- name: CheckAccountInstanceJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $1

-- name: CheckAccountInstanceJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: CheckBalanceJournal-main
select id_transaksi, kode_valuta 
from (
    select dt.id_transaksi, dt.kode_valuta,
        sum(
            case
                when dt.jenis_mutasi = 'D' then dt.nilai_mutasi
                else 0.0
            end
        ) as amount_debit,
        sum(
            case
                when dt.jenis_mutasi = 'C' then dt.nilai_mutasi
                else 0.0
            end
        ) as amount_credit
    from %[1]s.detiltransaksi dt
        inner join %[1]s.transaksi t on dt.id_transaksi = t.id_transaksi
        inner join %[1]s.pendingtransactionjournal ptj on ptj.id_transaksi = t.id_transaksi
        inner join %[1]s.account acc on dt.kode_account = acc.account_code
    where %[2]s 
        and acc.account_type <> 'M'
    group by dt.id_transaksi, dt.kode_valuta
)
where abs(amount_debit - amount_credit) > 0.001

-- name: CheckBalanceJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $1

-- name: CheckBalanceJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: MergeJournalNumber-main
update %[1]s.pendingtransactionjournal target
set journal_no = upd.journal_no
from  %[1]s.transaksi upd
where upd.id_transaksi = target.id_transaksi and %[2]s

--name: MergeJournalNumber-filter-IdTransaksiNotZero
target.id_transaksi = $1

--name: MergeJournalNumber-filter-IdTransaksiZero
target.flag_sent is null

-- name: InsertSelectJournal-main
insert into %[1]s.journal (
    journal_no, journal_date_posting, description, userid_create, userid_posting, is_posted, journal_type, journal_date, lastitemno, 
    is_partlychecked, branch_code, transaction_date, journal_state, serial_no, periode_code, fl_accountingperiode, fl_accountingyear,
    reference_no, reference_no_fwd
)
select ptj.journal_no, t.tanggal_transaksi, t.keterangan, t.user_input as userid_create, t.user_input, 'T', t.kode_transaksi, t.tanggal_transaksi, 0, 
    'T', t.kode_cabang_transaksi, t.tanggal_input, 'C', '', accd.periode_code, accd.fl_accountingperiode, accd.fl_accountingyear, t.nomor_referensi,
    t.nomor_referensi
from %[1]s.pendingtransactionjournal ptj,
    %[1]s.accountingday accd,
    %[1]s.transaksi t
where %[2]s
    and ptj.journal_no is not null
    and ptj.id_transaksi = t.id_transaksi
    and accd.datevalue = to_date($1, 'yyyy-mm-dd') 

-- name: InsertSelectJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $2

-- name: InsertSelectJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: InsertSelectJournalItem-main
insert into %[1]s.journalitem (
    fl_journal, journalitem_no, fl_account, amount_debit, amount_credit, description, branch_code, valuta_code, fl_project_no, 
    userid_create, userid_posting, userid_check, dateposting, datecreate , journalitemtype, kode_kurs, nilai_kurs, accountinstance_id,
    subaccountcode, subaccountname
)
select ptj.journal_no,  nextval('%[1]s.seq_journalitem'), dt.kode_account,
    case 
        when dt.jenis_mutasi = 'D' then dt.nilai_mutasi
        else 0.0
    end as amount_debit,
    case 
        when dt.jenis_mutasi = 'C' then dt.nilai_mutasi
        else 0.0
    end as amount_credit, 
    dt.keterangan, dt.kode_cabang, dt.kode_valuta, null, 
    t.user_input, t.user_input, t.user_input, t.tanggal_input, t.tanggal_input, 'N', dt.kode_kurs, dt.nilai_kurs_manual, ai.accountinstance_id,
    dt.nomor_rekening, rt.nama_rekening
from %[1]s.pendingtransactionjournal ptj
    inner join %[1]s.transaksi t on ptj.id_transaksi = t.id_transaksi
    inner join %[1]s.detiltransaksi dt on dt.id_transaksi = t.id_transaksi
    inner join %[1]s.accountinstance ai on (ai.account_code = dt.kode_account and ai.branch_code = dt.kode_cabang and ai.currency_code = dt.kode_valuta)
    left join %[1]s.rekeningtransaksi rt on rt.nomor_rekening = dt.nomor_rekening
where %[2]s

-- name: InsertSelectJournalItem-filter-IdTransaksiNotZero
ptj.id_transaksi = $1

-- name: InsertSelectJournalItem-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: BalancingRAKOtomatisJournal-main
insert into %[1]s.journalitem (
    fl_journal, journalitem_no, fl_account, amount_debit, amount_credit, description, branch_code, valuta_code, fl_project_no, 
    userid_create, userid_posting, userid_check, dateposting, datecreate, journalitemtype, kode_kurs, nilai_kurs, accountinstance_id,
    subaccountcode, subaccountname
)
select 
    ubitems.fl_journal,  nextval('%[1]s.seq_journalitem'), $1, ubitems.amount_credit, ubitems.amount_debit, 'BALANCING RAK OTOMATIS', ubitems.branch_code, ubitems.valuta_code, null, 
    t.user_input, t.user_input, t.user_input, t.tanggal_input, t.tanggal_input, 'N', null, 1, ai.accountinstance_id, $2 || '-' ||ubitems.branch_code||'-'||ubitems.valuta_code, $3
FROM 
(
    select ptj.id_transaksi, ji.fl_journal, ji.valuta_code, ji.branch_code, sum(amount_debit) as amount_debit, sum(amount_credit) as amount_credit
    from %[1]s.journalitem ji
        inner join %[1]s.account acc on acc.account_code = ji.fl_account 
        inner join %[1]s.pendingtransactionjournal ptj on ptj.journal_no = ji.fl_journal
    where acc.account_type <> 'M'
        and %[2]s
    group by ptj.id_transaksi, ji.fl_journal, ji.valuta_code, ji.branch_code
) ubitems
    inner join %[1]s.transaksi t on ubitems.id_transaksi = t.id_transaksi
    inner join %[1]s.accountinstance ai on (ai.branch_code = ubitems.branch_code and ai.currency_code = ubitems.valuta_code)
where abs(ubitems.amount_debit - ubitems.amount_credit) > 0.001
    and ai.account_code = $4 

-- name: BalancingRAKOtomatisJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $5

-- name: BalancingRAKOtomatisJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: UpdateTrxStatusJournal-main
update %[1]s.transaksi t
set is_sudah_dijurnal = 'T'
from (
    select * 
    from %[1]s.pendingtransactionjournal ptj
    where %[2]s
) upd 
where upd.id_transaksi = t.id_transaksi

-- name: UpdateTrxStatusJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $1

-- name: UpdateTrxStatusJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: InsertCompletedJournal-main
insert into %[1]s.completedtransactionjournal
(id_transaksi, journal_no)
select id_transaksi, journal_no 
from %[1]s.pendingtransactionjournal ptj
where %[2]s

-- name: InsertCompletedJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $1

-- name: InsertCompletedJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: DeletePendingJournal-main
delete from %[1]s.pendingtransactionjournal ptj
where %[2]s

-- name: DeletePendingJournal-filter-IdTransaksiNotZero
ptj.id_transaksi = $1

-- name: DeletePendingJournal-filter-IdTransaksiZero
ptj.flag_sent is null

-- name: GetDenominasitransaksikas-main
select id_denominasitransaksikas, id_denominasi_valuta, id_detil_transaksi, jumlah_denominasi, jumlah_denominasi_tle
from %[1]s.denominasitransaksikas d
where id_detil_transaksi = $1

-- name: UpdateDenominasiRekening-main
update %[1]s.denominasikas x
set jumlah_denominasi = x.jumlah_denominasi %[2]s $1
where id_denominasi_valuta = $2
    and nomor_rekening = $3

-- name: InsertDenominasitransaksikas-main
select d.jumlah_denominasi, d.id_denominasi_valuta, d.nomor_rekening, d.id_denominasikas, d.jumlah_denominasi_tle,
    dv.nilai_denominasi, dv.kode_valuta, dv.jenis_denominasi
from %[1]s.denominasikas d
    inner join %[1]s.denominasivaluta dv on dv.id_denominasi_valuta = d.id_denominasi_valuta 
where d.id_denominasi_valuta = $1 
    and nomor_rekening = $2

-- name: GetDetilDenominasiKas-main
select d.jumlah_denominasi, d.id_denominasi_valuta, d.nomor_rekening, d.id_denominasikas, d.jumlah_denominasi_tle,
			dv.nilai_denominasi, dv.kode_valuta, dv.jenis_denominasi
		from %[1]s.denominasikas d
			inner join %[1]s.denominasivaluta dv on dv.id_denominasi_valuta = d.id_denominasi_valuta 
		where d.id_denominasi_valuta = $1 
			and nomor_rekening = $2


-- name: GetIdGroup-main
select  nextval('%[1]s.seq_notadebitoutwardgroup') as id_group 

-- name: GetIdNotadebitoutward-main
select  nextval('%[1]s.seq_notadebitoutward') as id_notadebitoutward 

-- name: InsertNotaDebitOutwardGroup-main
insert into %[1]s.notadebitoutwardgroup(
    id_group, tanggal_input, tanggal_efektif, status_notadebit, nomor_rekening, nomor_aplikasi, keterangan, kode_cabang, 
    kode_valuta, inputer, is_otorisasi, tipe_grup, tanggal_eksekusi_titipan, user_eksekusi_titipan, jam_eksekusi_titipan
) 
values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)

-- name: InsertNotaDebitOutward-main
insert into %[1]s.notadebitoutward(
    nomor_referensi, tanggal_setoran, kode_valuta, nominal, nama_pemilik_nota, status, jenis_warkat, tanggal_input, id_notadebitoutward, 
    is_otorisasi, kurs_transaksi, nominal_ekuivalen, nomor_rekening_nota, tanggal_efektif, inputer, kode_cabang, biaya, biaya_ekuivalen, 
    sumber_biaya, nomor_aplikasi, keterangan, nomor_batch, kode_bank, nomor_rekening, id_group, sandi_wilayah_kliring, status_wilayah_kliring, 
    kode_account_titipan, tanggal_eksekusi_efektif, waktu_eksekusi_efektif, user_eksekusi_efektif
) 
values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, 
    $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31
)

-- name: GetNotadebitoutwardByIdGroup-main
select n.nomor_referensi, to_char(n.tanggal_setoran, 'yyyy-mm-dd') as tanggal_setoran, n.kode_valuta, n.nominal, n.nama_pemilik_nota, 
    n.status, n.jenis_warkat, to_char(n.tanggal_input, 'yyyy-mm-dd') as tanggal_input, n.id_notadebitoutward, n.is_otorisasi, 
    n.kurs_transaksi, n.nominal_ekuivalen, n.nomor_rekening_nota, to_char(n.tanggal_efektif, 'yyyy-mm-dd') as tanggal_efektif, 
    n.inputer, n.kode_cabang, n.biaya, n.biaya_ekuivalen, n.sumber_biaya, n.nomor_aplikasi, n.keterangan, n.nomor_batch, n.kode_bank, 
    n.nomor_rekening, n.id_group, n.sandi_wilayah_kliring, n.status_wilayah_kliring, n.kode_account_titipan, 
    to_char(n.tanggal_eksekusi_efektif, 'yyyy-mm-dd') as tanggal_eksekusi_efektif, to_char(n.waktu_eksekusi_efektif, 'yyyy-mm-dd') as waktu_eksekusi_efektif, 
    n.user_eksekusi_efektif 
from %[1]s.notadebitoutward n
where n.id_group = $1

-- name: GetNotadebitoutwardGroupByIdGroup-main
select n.id_group, to_char(n.tanggal_input, 'yyyy-mm-dd') as tanggal_input, to_char(n.tanggal_efektif, 'yyyy-mm-dd') as tanggal_efektif, 
    n.status_notadebit, n.nomor_rekening, n.nomor_aplikasi, n.keterangan, n.kode_cabang, n.kode_valuta, n.inputer, n.is_otorisasi, 
    n.tipe_grup, to_char(n.tanggal_eksekusi_titipan, 'yyyy-mm-dd') as tanggal_eksekusi_titipan, n.user_eksekusi_titipan, 
    to_char(n.jam_eksekusi_titipan, 'yyyy-mm-dd') as jam_eksekusi_titipan
from %[1]s.notadebitoutwardgroup n
where n.id_group = $1

-- name: GetNDOGByIdTrx-main
select distinct ndog.id_group, to_char(ndog.tanggal_input, 'yyyy-mm-dd') as tanggal_input, to_char(ndog.tanggal_efektif, 'yyyy-mm-dd') as tanggal_efektif, 
    ndog.status_notadebit, ndog.nomor_rekening, ndog.nomor_aplikasi, ndog.keterangan, ndog.kode_cabang, ndog.kode_valuta, ndog.inputer, ndog.is_otorisasi, 
    ndog.tipe_grup, to_char(ndog.tanggal_eksekusi_titipan, 'yyyy-mm-dd') as tanggal_eksekusi_titipan, ndog.user_eksekusi_titipan, 
    to_char(ndog.jam_eksekusi_titipan, 'yyyy-mm-dd') as jam_eksekusi_titipan
from %[1]s.dtkliringoutward dko
    inner join %[1]s.detiltransaksi d on d.id_detil_transaksi = dko.id_detil_transaksi 
    inner join %[1]s.transaksi t on t.id_transaksi = d.id_transaksi 
    inner join %[1]s.notadebitoutward ndo on ndo.id_notadebitoutward = dko.id_notadebitoutward 
    inner join %[1]s.notadebitoutwardgroup ndog on ndog.id_group = ndo.id_group 
where t.id_transaksi = $1

-- name: GetDKOByIdTrx-main
select dko.id_detil_transaksi, dko.id_notadebitoutward, dko.id_wic, d.nomor_rekening
from %[1]s.dtkliringoutward dko
    inner join %[1]s.detiltransaksi d on d.id_detil_transaksi = dko.id_detil_transaksi 
    inner join %[1]s.transaksi t on t.id_transaksi = d.id_transaksi
where t.id_transaksi = $1

-- name: GetDKTByIdTrx-main
select dkt.id_detil_transaksi, dkt.id_detil_tolakan, dkt.jenis_warkat, dkt.nomor_warkat, d.nomor_rekening
from %[1]s.dtkliringtarikan dkt
    inner join %[1]s.detiltransaksi d on d.id_detil_transaksi = dkt.id_detil_transaksi 
    inner join %[1]s.transaksi t on t.id_transaksi = d.id_transaksi
where t.id_transaksi = $1

-- name: UpdateStatusNotadebitoutwardgroup-main
update %[1]s.notadebitoutwardgroup 
set status_notadebit = $1
where id_group = $2

-- name: UpdateStatusNotadebitoutward-main
update %[1]s.notadebitoutward 
set status = $1
where id_group = $2

-- name: UpdateStatusNotaDebetOutward-main
UPDATE 
    %[1]v.NOTADEBITOUTWARD 
SET 
    status = 'E'
WHERE
    Id_NotaDebitOutward = $1

-- name: UpdateRecountSaldoFloat-main
MERGE
INTO
    %[1]v.REKENINGLIABILITAS rl
USING (
    SELECT
        rl.nomor_rekening,
        sum(nominal) total_saldo_float
    FROM
        %[1]v.RekeningLiabilitas rl
    LEFT JOIN %[1]v.NotaDebitOutward n ON
        rl.nomor_rekening = n.nomor_rekening
        AND n.status = 'S'
        AND n.is_otorisasi = 'T'
    WHERE
        rl.nomor_rekening = $1
    GROUP BY
        rl.nomor_rekening
    ) q 
ON
    (q.nomor_rekening = rl.nomor_rekening)
WHEN MATCHED THEN
    UPDATE
    SET
        saldo_float = coalesce(q.total_saldo_float, 0.0)
    WHERE
        nomor_rekening = $1

-- name: GetNotadebitoutwardByIdDbOW-main
select n.id_notadebitoutward, n.nomor_referensi, n.nomor_aplikasi , n.keterangan, n.nominal, n.tanggal_efektif, n.nomor_rekening 
from %[1]s.notadebitoutward n
where n.id_notadebitoutward = $1

-- name: UpdateStatusNotadebitoutwardIdb-main
update %[1]s.notadebitoutward 
set status = $1
where id_notadebitoutward = $2

-- name: InsertInfotarikankliring-main
insert into %[1]s.infotarikankliring (
    id_detil_transaksi, no_referensi_tpk, no_batch_tpk, nama_pemegang, noidentitas_pemegang, norekening_pemegang
) 
values (
    $1, $2, $3, $4, $5, $6
)

-- name: InsertDtkliringtarikan-main
insert into %[1]s.dtkliringtarikan (
    id_detil_transaksi, id_detil_tolakan, jenis_warkat, nomor_warkat
) 
values (
    $1, $2, $3, $4
)

-- name: DeleteInfotarikankliringByIDTrx-main
delete from %[1]s.infotarikankliring i
where exists (
    select 1 from %[1]s.detiltransaksi d 
    where d.id_detil_transaksi = i.id_detil_transaksi 
        and d.id_transaksi = $1
)

-- name: GetClrRejectReasonList-main
SELECT ID_PARAMETER_TOLAKAN, KETERANGAN, BIAYA_TOLAKAN FROM %[1]s.PARAMETERTOLAKANKLIRING

-- name: GetCodeGlInterface-main
SELECT gl.KODE_INTERFACE  
FROM %[1]s.GlobalGLInterface gl WHERE gl.KODE_INTERFACE = $1

-- name: InsertInfotolaankliring-main
INSERT INTO %[1]s.InfoTolakTarikanKliring
(
    ID_DETIL_TRANSAKSI,
    NOMOR_REKENING_GIRO,
    JENIS_WARKAT,
    NOMOR_WARKAT,
    JENIS_REKENING_TITIPAN,
    NOMOR_REKENING_TITIPAN,
    KODE_CABANG_TITIPAN,
    NOMINAL_TRANSAKSI,
    ID_PARAMETER_TOLAKAN,
    KODE_BANK_PENARIK,
    IS_SET_TARIKANKLIRING,
    ID_DETIL_TARIKANKLIRING,
    NAMA_TTD_WARKAT,
    NAMA_PENARIK,
    NOMOR_REKENING_PENARIK,
    NOMOR_IDENTITAS_PENARIK,
    NOMOR_REF_TPK_ORI,
    DEBET_NAMA,
    DEBET_NOMORIDENTITAS,
    DEBET_JENISNASABAH,
    DEBET_JENISPENDUDUK,
    DEBET_ALAMATJALAN,
    DEBET_ALAMATRT,
    DEBET_ALAMATRW,
    DEBET_ALAMATKOTA,
    DEBET_ALAMATPROPINSI,
    DEBET_ALAMATKODEPOS,
    DEBET_TANGGALLAHIR,
    DEBET_TEMPATLAHIR,
    DEBET_NPWP
)
VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
    $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
)

-- name: GetClrWithdrawList-main
SELECT
    C.TANGGAL_TRANSAKSI,
    A.JENIS_WARKAT,
    D.NOMOR_REFERENSI,
    C.NOMOR_REKENING,
    I.NAMA_REKENING,
    I.KODE_VALUTA,
    J.FULL_NAME,
    I.KODE_CABANG,
    C.NILAI_MUTASI,
    A.ID_DETIL_TRANSAKSI,
    C.ID_TRANSAKSI
FROM
%[1]s.DTKLIRINGTARIKAN A
    inner join %[1]s.DETILTRANSAKSI C on c.id_detil_transaksi = a.id_detil_transaksi
    inner join %[1]s.TRANSAKSI D on c.id_transaksi = d.id_transaksi
    inner join %[1]s.REKENINGLIABILITAS G on C.NOMOR_REKENING = G.NOMOR_REKENING
    inner join %[1]s.REKENINGCUSTOMER H on G.Nomor_Rekening = H.Nomor_Rekening
    inner join %[1]s.REKENINGTRANSAKSI I on H.Nomor_Rekening = I.Nomor_Rekening
    inner join %[1]s.CURRENCY J on I.KODE_VALUTA = J.CURRENCY_CODE
    LEFT JOIN %[2]s.CABANG F ON D.KODE_CABANG_TRANSAKSI = F.KODE_CABANG
WHERE D.TANGGAL_TRANSAKSI >= to_date($1, 'YYYY-MM-DD') AND D.TANGGAL_TRANSAKSI <= (to_date($1, 'YYYY-MM-DD') + 1)
    AND D.STATUS_OTORISASI = 1
    AND A.ID_DETIL_TOLAKAN IS NULL
    and C.JENIS_MUTASI = 'D'
ORDER BY
    A.ID_DETIL_TRANSAKSI ASC

-- name: InsertDTKliringOutward-main
insert into %[1]s.dtkliringoutward (
    id_detil_transaksi, id_notadebitoutward, id_wic
) 
values (
    $1, $2, $3
)

-- name: GetListSetoranKliring-main
select 
    NOMOR_REKENING,
    to_char(tanggal_setoran, 'YYYY-MM-DD') tanggal_setoran,
    to_char(tanggal_efektif, 'YYYY-MM-DD') tanggal_efektif,
    nominal,
    kode_cabang,
    nomor_referensi
from %[1]s.NotaDebitOutward
where nomor_rekening = $1
    and status = 'S'
    and tanggal_setoran >= TO_DATE($2, '%[2]s')
    and tanggal_setoran < TO_DATE($3, '%[2]s')
order by tanggal_setoran asc, id_notadebitoutward asc

-- name: GetLayananByNamaLayanan-main
select nama_layanan, deskripsi, jenis_layanan, id_layanan, periode_proses_layanan, tanggal_proses_berikutnya, tanggal_proses_terakhir,
    jumlah_maksimal_register, is_ada_proses_periodik, tanggal_proses_pertama, is_layanan_system, is_override_tanggal_proses, 
    nomor_rekening_biaya, jarak_periode_proses, status_biaya_proses, status_data, tag_layanan, biaya_layanan, id_event, user_pembuat, 
    user_pengubah 
from %[1]s.layanan l 
where nama_layanan = $1

-- name: CreateListIndRekening-main
insert into %[1]s.listindividurekening
(id_individu, nomor_rekening, keterangan, is_aktif, kode_sumber_dana)
values 
($1, $2, $3, $4, $5)

-- name: UpdateListIndividuRekByIDIndividu-main
update %[1]s.listindividurekening
set keterangan = $1, is_aktif = $2, kode_sumber_dana = $3
where nomor_rekening = $4
and id_individu = $5

-- name: UpdateListIndividuRekByListIndividu-main
update %[1]s.listindividurekening
set is_aktif = 'F'
where nomor_rekening = $1
and id_individu not in $2

-- name: UpdateListIndividuRekByNoRek-main
update %[1]s.listindividurekening
set is_aktif = 'F'
where nomor_rekening = $1

-- name: UpdateStatusListRek-main
update %[1]s.listnomorrekening
set status = 'U'
where nomor_rekening = $1

-- name: GetListNomorRekeningByNoRek-main
select nomor_rekening, kode_cabang, sequence_no, status
from %[1]s.listnomorrekening
where nomor_rekening = $1

-- name: InsertListNomorRekening-main
insert into %[1]s.listnomorrekening
(nomor_rekening, kode_cabang, sequence_no, status)
values($1, $2, $3, $4)

-- name: CreateJointAccount-main
insert into %[1]s.multiciflink
(LINKID, NOMOR_REKENING, NOMOR_NASABAH)
values
( nextval('%[1]s.seq_multiciflink'), $1, $2)

-- name: DeleteMultiCIFLink-main
delete from %[1]s.multiciflink
where nomor_rekening = $1

-- name: GetNasabah-main
select n.nomor_nasabah, n.nama_nasabah, n.jenis_nasabah, n.status_data
from %[1]s.nasabah n 
where n.nomor_nasabah = $1

-- name: GetNasabahIndividu-main
select n.nomor_nasabah, n.nama_nasabah, n.jenis_nasabah, n.status_data, ni.id_individu
from %[1]s.nasabah n 
    left join %[1]s.nasabahindividu ni on ni.nomor_nasabah = n.nomor_nasabah
where n.nomor_nasabah = $1

-- name: GetNasabahInfo-main
select n.nomor_nasabah ,n.nama_nasabah,
case
    when n.jenis_nasabah = 'I' then to_char(i.tanggal_lahir, 'yyyy-mm-dd')
    else ''
end as tanggal_lahir,
case 
    when n.jenis_nasabah = 'I' then i.nama_ibu_kandung
    else ''
end as nama_ibu_kandung, 
case 
    when n.jenis_nasabah = 'K' then nk.nomor_siupp
    else ''
end as nomor_siup,
case
    when n.jenis_nasabah = 'K' then nk.alamat_jalan || ' ' || nk.alamat_rt || ' ' || nk.alamat_rw || ' ' || nk.alamat_kelurahan || ' ' || nk.alamat_kecamatan || ' ' ||
        nk.alamat_kota_kabupaten || ' ' || nk.alamat_provinsi || ' ' || nk.alamat_kode_pos 
    when n.jenis_nasabah = 'I' then i.alamat_rumah_jalan || ' ' || i.alamat_rumah_rt || ' ' || i.alamat_rumah_rw || ' ' || i.alamat_rumah_kelurahan || ' ' || i.alamat_rumah_kecamatan 
        || ' ' || i.alamat_rumah_kota_kabupaten || i.alamat_rumah_provinsi || ' ' || i.alamat_rumah_kode_pos
end as alamat,
case 
    when n.jenis_nasabah = 'K' then ''
    when n.jenis_nasabah = 'I' then i.telepon_rumah_kode_area || '-' || i.telepon_rumah_nomor
end as telepon_rumah_nomor,
case 
    when n.jenis_nasabah = 'K' then ''
    when n.jenis_nasabah = 'I' then i.telepon_hp_nomor
end as telepon_hp_nomor,
case 
    when n.jenis_nasabah = 'K' then nk.telp_kantor1_kode_area || '-' || nk.telp_kantor1_nomor
    when n.jenis_nasabah = 'I' then ''
end as telp_kantor1_nomor,
case 
    when n.jenis_nasabah = 'K' then nk.telp_kantor2_kode_area || '-' || nk.telp_kantor2_nomor
    when n.jenis_nasabah = 'I' then ''
end as telp_kantor2_nomor
from %[1]s.nasabah n
left join %[1]s.nasabahindividu ni on ni.nomor_nasabah = n.nomor_nasabah 
left join %[1]s.individu i on i.id_individu = ni.id_individu 
left join %[1]s.nasabahkorporat nk on nk.nomor_nasabah = n.nomor_nasabah 
where n.nomor_nasabah = $1

-- name: UpdateStatusByIdBukuwarkat-main
update %[1]s.notadebetinternal
set status_warkat = $1
where id_bukuwarkat = $2

-- name: UpdateStatusWarkat-main
update %[1]s.notadebetinternal
set status_warkat = $1,
    user_pengubah = $2,
    tanggal_perubahan_status = current_timestamp
where nomor_seri = $3
    and jenis_warkat = $4
    and nomor_rekening = $5

-- name: UpdateStatusWarkatCair-main
update %[1]s.notadebetinternal
set status_warkat = $1,
    user_pengubah = $2,
    tanggal_perubahan_status = current_timestamp,
    tanggal_pakai_warkat = $3,
    used_counter = coalesce(used_counter, 0) + 1
where nomor_seri = $4
    and jenis_warkat = $5
    and nomor_rekening = $6

-- name: ReverseStatusWarkatCair-main
update %[1]s.notadebetinternal
set status_warkat = 'A',
    tanggal_pakai_warkat = null
where nomor_seri = $1

-- name: UpdateCHQStatusWarkat-main
update %[1]s.notadebetinternal
set status_warkat = $1,
tanggal_perubahan_status = $2
where nomor_seri = $3

-- name: GetWarkatDetail-main
select n.nomor_seri, n.jenis_warkat, to_char(n.tanggal_pakai_warkat, 'yyyy-mm-dd') as tanggal_pakai_warkat, n.no_seri_detil_transaksi, n.id_notadebetinternal, n.status_warkat, 
    case
        when n.status_warkat = 'A' then 'AKTIF'
        when n.status_warkat = 'B' then 'DIBATALKAN'
        when n.status_warkat = 'C' then 'CAIR'
        when n.status_warkat = 'H' then 'HILANG'
        when n.status_warkat = 'P' then 'PROSES'
        when n.status_warkat = 'S' then 'BLOKIR'
        else ''
    end as status_warkat_desc,
    to_char(n.tanggal_perubahan_status, 'yyyy-mm-dd') as tanggal_perubahan_status, n.user_pengubah, n.id_bukuwarkat, n.m_nomrek, n.nomor_rekening, n.used, n.used_counter
from %[1]s.notadebetinternal n
    inner join %[1]s.bukuwarkat b on b.id_bukuwarkat = n.id_bukuwarkat
where n.nomor_seri = $1
    and n.jenis_warkat = $2
    and n.nomor_rekening = $3

-- name: GetValidationLimitTrx-main
select ptu.kode_transaksi, ptu.keterangan, ptu.jenis_limit_transaksi, lt.id_user, lt.nilai_limit 
from %[1]s.parametertransaksiumum ptu
    inner join %[2]s.limittransaksi lt on lt.jenis_limit = ptu.jenis_limit_transaksi 
where kode_transaksi = $1
    and lt.id_user = $2

-- name: GetMaxDbHarian-main
select count(*) as total_limit
from %[1]s.limittransaksirekening lt
where nomor_rekening = $1 
    and periode = 'H' 
    and jenis_mutasi = 'D'

-- name: GetLimitDebetHarian-main
select sum(nilai_transaksi) as total_nilai_transaksi
from %[1]s.limittransaksirekening lt
where nomor_rekening = $1
    and periode = 'H' 
    and jenis_mutasi = 'D'

-- name: GetCounterDebetHarian-main
select sum(nilai_transaksi) as total_nilai_transaksi
from %[1]s.countertransaksirekening lt
where nomor_rekening = $1 
    and periode = 'H' 
    and jenis_mutasi = 'D'

-- name: GetGlobalParam-main
select kode_parameter, tipe_parameter, nilai_parameter, deskripsi, is_parameter_system, 
    to_char(nilai_parameter_tanggal, 'yyyy-mm-dd') as nilai_parameter_tanggal, nilai_parameter_string, locked
from %[1]s.parameterglobal 
where kode_parameter = $1

-- name: GetGlobalparamRem-main
select kode_parameter, deskripsi, is_parameter_system, 
    to_char(nilai_parameter_tanggal, 'yyyy-mm-dd') as nilai_parameter_tanggal, nilai_parameter_string
from %[1]s.parameterglobal 
where kode_parameter = $1

-- name: CreatePassbookHistBalance-main
insert into %[1]s.histcetaksaldopassbook 
(id_histori, id_register, waktu_cetak, user_cetak, saldo_passbook, no_buku_passbook, halaman_passbook, baris_passbook, alasan, user_override) 
values( nextval('%[1]s.seq_histcetaksaldopassbook'), $1, SYSDATE, $2, $3, $4, $5, $6, $7, $8)

-- name: GetPassbookType-main
SELECT kode_passbook, nama_passbook 
FROM %[1]s.jenispassbook

-- name: GetAccountPassbook-main
select regp.id_register, rt.nomor_rekening, rt.nama_rekening, rt.kode_cabang, c.nama_cabang, regp.nomorregisterbuku_aktif, coalesce(rl.alamat_kirim,'N') alamat_kirim, 
    n.jenis_nasabah, n.nomor_nasabah, to_char(rl.tanggal_buka, 'yyyy-mm-dd') as tanggal_buka, to_char(regl.tanggal_registrasi, 'yyyy-mm-dd') as tanggal_registrasi, 
    regp.halaman_cetak_terakhir, regp.baris_cetak_terakhir, regp.nomor_buku_terakhir, round(regp.saldo_cetak_terakhir, 2) as saldo_cetak_terakhir, n.nama_nasabah, regp.is_ganti_buku, regp.is_cetak_cover,
    regp.id_detil_transaksi
from %[1]s.rekeningtransaksi rt 
    inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rt.nomor_rekening
    inner join %[2]s.cabang c on c.kode_cabang = rt.kode_cabang
    inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah
    inner join %[1]s.registerlayanan regl on regl.nomor_rekening_layanan = rt.nomor_rekening
    inner join %[1]s.registerpassbook regp on regl.id_register = regp.id_register
where rt.nomor_rekening = $1

-- name: CekNomorRegisterBuku-main
select nomor_register_buku, nomor_rekening 
from %[1]v.registerbukutabungan r 
where nomor_register_buku = $1 

-- name: GetAlternateAddress-main
select nomor_rekening, alamat_jalan, alamat_kecamatan, alamat_rtrw, alamat_kelurahan, alamat_provinsi, alamat_kode_pos, alamat_kota_kabupaten
from %[1]s.alamatalternatif 
where nomor_rekening = $1

-- name: GetCustomerAddress-main
select alamat_surat_jalan, alamat_surat_kode_pos, alamat_surat_rtrw, alamat_surat_kelurahan, alamat_surat_kecamatan, 
    alamat_surat_kota_kabupaten, alamat_surat_provinsi
from %[1]s.nasabah
where nomor_nasabah = $1

-- name: GetCustomerPersonalAddress-main
select i.alamat_rumah_jalan, i.alamat_rumah_kode_pos, i.alamat_rumah_rt || '/' || i.alamat_rumah_rw as alamat_rumah_rtrw, 
    i.alamat_rumah_kelurahan, i.alamat_rumah_kecamatan, i.alamat_rumah_kota_kabupaten, i.alamat_rumah_provinsi
from %[1]s.nasabahindividu ni
    inner join %[1]s.individu i on ni.id_individu = i.id_individu
where ni.nomor_nasabah = $1

-- name: GetCustomerCorporateAddress-main
select alamat_jalan, alamat_kode_pos, alamat_rt || '/' || alamat_rw as alamat_rtrw, 
    alamat_kelurahan, alamat_kecamatan, alamat_kota_kabupaten, alamat_provinsi
from %[1]s.nasabahkorporat
where nomor_nasabah = $1

-- name: CreatePassbookHistPrint-main
insert into %[1]s.historycetaknamapassbook 
(id_history, user_override, nomor_rekening, user_input, tgl_input, tgl_sistem, nomor_register_buku, nomor_register_buku_lama, 
    alasan_cetak, kode_cabang_input, ip_input
)
values ( nextval('%[1]s.seq_historycetaknamapassbook'), $1, $2, $3, sysdate, $4, $5, $6, $7, $8, $9 )

-- name: CreateRegisterSavingBook-main
insert into %[1]s.registerbukutabungan 
(nomor_register_buku, nomor_register_buku_lama, nomor_rekening, tanggal_input, tanggal_sistem, user_input, user_override, alasan_cetak, kode_cabang_input, 
    kode_passbook)
values ($1, $2, $3, sysdate, sysdate, $4, $5, $6, $7, $8)

-- name: UpdateRegisterPassbook-main
update %[1]s.registerpassbook
set is_baru_register = $1, 
    nomorregisterbuku_aktif = $2,
    halaman_cetak_terakhir = $3, 
    baris_cetak_terakhir = $4, 
    nomor_buku_terakhir = NOMOR_BUKU_TERAKHIR + 1, 
    is_ganti_buku = $5, 
    is_cetak_cover = $6, 
    kode_passbook_terakhir = $7
where id_register = $8

-- name: GetValNorekPassbook-main
select t.status_otorisasi, d.nomor_rekening, t.kode_transaksi, t.kode_cabang_transaksi 
from %[1]s.detiltransaksi d
    inner join %[1]s.transaksi t on d.id_transaksi  = t.id_transaksi
where d.nomor_rekening = $1 
    and t.status_otorisasi = 0
    
-- name: GetDataTransaksi-main
select d.id_detil_transaksi, d.nilai_mutasi, d.Jenis_Mutasi, d.Keterangan, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, t.kode_transaksi
from %[1]s.histdetiltransaksi d
    inner join %[1]s.histtransaksi t on d.id_transaksi = t.id_transaksi
where nomor_rekening = $1 
    and (hide_passbook = 'F' or hide_passbook is null) 
    and not exists (
        select 1 
        from %[1]s.detiltransaksipassbook dt
        where dt.id_detil_transaksi = d.id_detil_transaksi
    )
union
select d.id_detil_transaksi, d.nilai_mutasi, d.Jenis_Mutasi, d.Keterangan, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, t.kode_transaksi
from %[1]s.detiltransaksi d
    inner join %[1]s.transaksi t on d.id_transaksi = t.id_transaksi
where nomor_rekening = $1
    and (hide_passbook = 'F' or hide_passbook is null)
    and not exists (
        select 1 
        from %[1]s.detiltransaksipassbook dt
        where dt.id_detil_transaksi = d.id_detil_transaksi
    )
order by tanggal_transaksi, id_detil_transaksi

-- name: GetSeqInfoCetakPassbook-main
select  nextval('%[1]s.seq_infocetakpassbook')


-- name: GetRegPassbookLast-main
select id_cetak_passbook, baris, nilai_mutasi, jenis_mutasi, halaman, to_char(tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, to_char(tanggal_cetak, 'yyyy-mm-dd') as tanggal_cetak, 
    is_sudah_cetak, is_gtu, nomor_rekening, kode_transaksi, keterangan, user_cetak, no_buku, nomor_register_buku
from %[1]s.infocetakpassbook i 
where nomor_rekening = $1  
    and id_cetak_passbook = (
        select max(id_cetak_passbook) 
        from %[1]s.infocetakpassbook
        where nomor_rekening = $1
            and is_sudah_cetak = 'T'
    )

-- name: GetCollectionDataTrx-main
select d.id_detil_transaksi, to_char(t.tanggal_transaksi, 'dd-mm-yy') as tanggal_transaksi, d.keterangan, d.jenis_mutasi, d.nilai_mutasi, t.id_transaksi, 
    t.user_input, t.kode_transaksi, p.id_cetak_passbook, i.baris, i.halaman
from %[1]s.histtransaksi t
    inner join %[1]s.histdetiltransaksi d on t.id_transaksi = d.id_transaksi
    inner join %[1]s.detiltransaksipassbook p on d.id_detil_transaksi = p.id_detil_transaksi
    inner join %[1]s.infocetakpassbook i on p.id_cetak_passbook = i.id_cetak_passbook
where (d.hide_passbook = 'F' or d.hide_passbook is null)
    and d.nomor_rekening = $1
    %[2]s
union all
select d.id_detil_transaksi, to_char(t.tanggal_transaksi, 'dd-mm-yy') as tanggal_transaksi, d.keterangan, d.jenis_mutasi, d.nilai_mutasi, t.id_transaksi, 
    t.user_input, t.kode_transaksi, p.id_cetak_passbook, i.baris, i.halaman
from %[1]s.Transaksi t
    inner join %[1]s.detiltransaksi d on t.id_transaksi=d.id_transaksi
    inner join %[1]s.detiltransaksipassbook p on d.id_detil_transaksi = p.id_detil_transaksi
    inner join %[1]s.infocetakpassbook i on p.id_cetak_passbook = i.id_cetak_passbook
where (d.hide_passbook = 'F' or d.hide_passbook is null)
    and d.nomor_rekening= $1
    %[3]s
order by id_cetak_passbook

-- name: GetCollectionDataTrx-filter-CetakUlangTrue
and i.is_sudah_cetak = 'T'

-- name: GetCollectionDataTrx-filter-CetakUlangFalse
and i.is_sudah_cetak = 'F'

-- name: GetCollectionDataTrx-filter-AllDataTrue1
and i.id_cetak_passbook >= $1
	and i.id_cetak_passbook <= $2

-- name: GetCollectionDataTrx-filter-AllDataTrue2
and i.id_cetak_passbook >= $1
	and i.id_cetak_passbook <= $2

-- name: GetBalance-main
select balance
from %[1]s.dailybalancerekening d1
    inner join (
        select nomor_rekening, max(balance_date) as balance_date 
        from %[1]s.dailybalancerekening 
        where nomor_rekening = $1
            and balance_date <= to_date($2, 'yyyy-mm-dd') - INTERVAL '1' DAY
            and balance_field = 'saldo'
        group by nomor_rekening
    ) d2 on (
        d1.balance_date =  d2.balance_date 
        and d1.nomor_rekening = d2.nomor_rekening)
where d1.nomor_rekening = $1   
    and balance_field = 'saldo'

-- name: InsertInfoCetakPassbook-main
insert into %[1]s.infocetakpassbook 
(id_cetak_passbook, is_sudah_cetak, is_gtu, jenis_mutasi, nilai_mutasi, tanggal_transaksi, kode_transaksi, keterangan, nomor_rekening) 
values
($1, $2, $3, $4, $5, $6, $7, $8, $9)

-- name: InsertDetailTransaksiPassbook-main
insert into %[1]s.detiltransaksipassbook 
(id_detil_transaksi, id_cetak_passbook) 
values
($1, $2)

-- name: GetjenisPasbook-main
select jp.total_baris, jp.jumlah_halaman 
from %[1]s.jenispassbook jp 
where jp.kode_passbook = $1

-- name: GetDicMapCodeTrx-main
select m.kode_transaksi, k.kode_trans_passbook 
from %[1]s.mapkodetranspassbook m
    inner join %[1]s.kodetranspassbook k on m.id_kodetrans = k.id_kodetrans
where kode_passbook = $1

-- name: GetTglByIdDetailTrx-main
select to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi
from %[1]s.detiltransaksi dt
    inner join %[1]s.transaksi t on dt.id_transaksi = t.id_transaksi
where dt.id_detil_transaksi = $1
union
select to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi
from %[1]s.histdetiltransaksi dt
    inner join %[1]s.histtransaksi t on dt.id_transaksi = t.id_transaksi
where dt.id_detil_transaksi = $1

-- name: GetDataNTrx-main
select coalesce(sum(case when jenis_mutasi = 'D' then -nilai_mutasi else nilai_mutasi end), 0.0) total_mutasi
from (
    select jenis_mutasi, nilai_mutasi 
    from %[1]s.histtransaksi t
        inner join %[1]s.histdetiltransaksi dt on t.id_transaksi = dt.id_transaksi
    where nomor_rekening = $1 
        and t.tanggal_transaksi >= to_date($2, 'yyyy-mm-dd')
        and dt.Id_Detil_Transaksi < $3 
    union all
    select jenis_mutasi, nilai_mutasi 
    from %[1]s.transaksi t
        inner join %[1]s.detiltransaksi dt on t.id_transaksi = dt.id_transaksi
    where nomor_rekening = $1
        and t.tanggal_transaksi >= to_date($2, 'yyyy-mm-dd')
        and dt.Id_Detil_Transaksi < $3
) nom

-- name: UpdateInfoCetakPassbook-main
update %[1]s.infocetakpassbook
set baris = $1, is_sudah_cetak = 'T', tanggal_cetak = current_timestamp, halaman = $2, no_buku = $3, user_cetak = $4, nomor_register_buku = $5
where id_cetak_passbook = $6

-- name: GetDataInfoCetakPassbook-main
select i.baris , i.halaman , i.no_buku 
from %[1]s.infocetakpassbook i 
where i.id_cetak_passbook = $1 

-- name: UpdateInfoCetakPassbookByIdReg-main
update %[1]s.infocetakpassbook
set baris = $1, is_sudah_cetak = 'T', tanggal_cetak = current_timestamp, halaman = $2, no_buku = $3
where id_cetak_passbook = $4

-- name: UpdatePassbookTxUnprint-main
update %[1]s.registerpassbook
set tanggal_cetak_terakhir = current_timestamp
    , id_detil_transaksi = $1
    , saldo_cetak_terakhir = $2
    , is_baru_register = $3
    , halaman_cetak_terakhir = $4
    , baris_cetak_terakhir = $5
    , is_ganti_buku = $6
    , is_cetak_cover = $7
where id_register = $8

-- name: GetPassbookLastprint-main
SELECT 
    regp.id_register as nomor_register_buku, 
    regp.halaman_cetak_terakhir, saldo_cetak_terakhir,
    baris_cetak_terakhir, rt.nama_rekening
FROM %[1]s.registerlayanan regl 
    INNER JOIN %[1]s.registerpassbook regp ON regl.id_register = regp.id_register
    INNER JOIN %[1]s.rekeningtransaksi rt on rt.nomor_rekening = regl.nomor_rekening_layanan
where nomor_rekening_layanan = $1

-- name: GetDataTrxUnPrint-main
select d.id_detil_transaksi, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, d.keterangan, d.jenis_mutasi, 
    case when d.jenis_mutasi = 'C' then 'Kredit' when d.jenis_mutasi = 'D' then 'Debet' else '' end as jenis_mutasi_desc,
    d.nilai_mutasi, t.id_transaksi, i.halaman, t.user_input, t.kode_transaksi, p.id_cetak_passbook, i.baris, i.no_buku 
from %[1]s.histtransaksi t
    inner join %[1]s.histdetiltransaksi d on t.id_transaksi = d.id_transaksi 
    inner join %[1]s.detiltransaksipassbook p on d.id_detil_transaksi = p.id_detil_transaksi
    inner join %[1]s.infocetakpassbook i on p.id_cetak_passbook = i.id_cetak_passbook
where (d.hide_passbook = 'F' or d.hide_passbook is null) 
    and d.nomor_rekening= $1
    and i.is_sudah_cetak = 'T'
    and t.tanggal_transaksi >= to_date($2, 'yyyy-mm-dd')
    and t.tanggal_transaksi < to_date($3, 'yyyy-mm-dd')
union all
select d.id_detil_transaksi, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, d.keterangan, d.jenis_mutasi, 
    case when d.jenis_mutasi = 'C' then 'Kredit' when d.jenis_mutasi = 'D' then 'Debet' else '' end as jenis_mutasi_desc,
    d.nilai_mutasi, t.id_transaksi, i.halaman, t.user_input, t.kode_transaksi, p.id_cetak_passbook, i.baris, i.no_buku 
from %[1]s.transaksi t
    inner join %[1]s.detiltransaksi d on t.id_transaksi=d.id_transaksi
    inner join %[1]s.detiltransaksipassbook p on d.id_detil_transaksi = p.id_detil_transaksi
    inner join %[1]s.infocetakpassbook i on p.id_cetak_passbook = i.id_cetak_passbook
where (d.hide_passbook = 'F' or d.hide_passbook is null) 
    and d.nomor_rekening= $1 
    and i.is_sudah_cetak = 'T'
    and t.tanggal_transaksi >= to_date($2, 'yyyy-mm-dd')
    and t.tanggal_transaksi < to_date($3, 'yyyy-mm-dd')
order by id_cetak_passbook

-- name: GetPassbookTotalBaris-main
select jp.total_baris from %[1]s.jenispassbook jp where jp.kode_passbook = '01'

-- name: UpdateInfoCetakPassWhenAllData-main
update %[1]s.infocetakpassbook
set is_sudah_cetak = 'T'
where id_cetak_passbook < $1
    and is_sudah_cetak = 'F'
    and nomor_rekening = $2

-- name: GetProduk-main
select kode_produk, nama_produk, is_onbalancesheetmode, saldo_minimum_tidak_aktif, jumlah_hari_jadi_tidak_aktif, saldo_minimum, saldo_maksimum, 
    is_backdated, jumlah_bulan_maks_backdated, is_kena_zakat_bagi_hasil, presentase_zakat_bagi_hasil, is_kena_pajak, tarif_pajak, 
    disposisi_bagi_hasil, periode_bagi_hasil, jenis_produk, kode_valuta, status, is_produk_berasuransi, is_produk_partnership, 
    is_override_bagihasilkhusus, is_override_tarifpajak, is_override_zakatbagihasil, is_override_disposisi_bgh, jarak_periode_bagihasil, 
    is_override_backdated, id_rencanabagihasildefault, is_kena_biayalayananumum, id_biayalayananumum, nama_valuta, is_produkkartu, nama_asuransi, 
    id_partner, nama_grup, insplanid, psplanid, biaya_adm_bulanan, id_tiering_biayaadm, id_tiering_nisbahbonus, is_tiering_biayaadm, 
    is_tiering_nisbahbonus, jenis_akad, kode_rencana_asuransi, round(nisbah_bonus_dasar, 2) as nisbah_bonus_dasar, setoran_awal_minimum, id_asuransi, is_param_saldo_tidak_aktif, 
    is_override_biaya_penutupan, biaya_penutupan_rekening, sumber_biaya_adm, jumlah_hari_pertahun, kode_digit, is_blokir_debet_internal, 
    is_blokir_debet_eksternal, is_blokir_kredit_internal, is_blokir_kredit_eksternal, is_biaya_saldo_minimum, is_biaya_rekening_dormant, 
    is_biaya_atm, is_cetak_nota, is_status_passbook, kolektibilitas, is_dapat_bagi_hasil, is_tidak_dormant, tanggal_acuan_dormant, 
    biaya_rekening_dormant, biaya_saldo_minimum, jumlah_hari_tutup_otomatis, is_blokir_debet, is_blokir_kredit, is_tutup_otomatis_dormant, 
    biaya_adm_atm, biaya_kartu_atm, ekuivalen_rate, saldo_minimum_bagihasil, is_proses_cadangan, persentase_cadangan, order_mb, 
    is_biaya_materai, is_override_biayaatm
from %[1]s.produk 
where kode_produk = $1

-- name: GetProdukTabungan-main
select 
    kode_produk, nama_produk, is_onbalancesheetmode, saldo_minimum_tidak_aktif, jumlah_hari_jadi_tidak_aktif, 
    saldo_minimum, saldo_maksimum, is_backdated, jumlah_bulan_maks_backdated, is_kena_zakat_bagi_hasil, 
    presentase_zakat_bagi_hasil, is_kena_pajak, tarif_pajak, disposisi_bagi_hasil, periode_bagi_hasil, 
    jenis_produk, kode_valuta, status, is_produk_berasuransi, is_produk_partnership, is_override_bagihasilkhusus, 
    is_override_tarifpajak, is_override_zakatbagihasil, is_override_disposisi_bgh, jarak_periode_bagihasil, 
    is_override_backdated, id_rencanabagihasildefault, is_kena_biayalayananumum, id_biayalayananumum, nama_valuta, 
    is_produkkartu, nama_asuransi, id_partner, nama_grup, insplanid, psplanid, biaya_adm_bulanan, 
    id_tiering_biayaadm, id_tiering_nisbahbonus, is_tiering_biayaadm, is_tiering_nisbahbonus, jenis_akad, 
    kode_rencana_asuransi, nisbah_bonus_dasar, setoran_awal_minimum, id_asuransi, is_param_saldo_tidak_aktif, 
    is_override_biaya_penutupan, biaya_penutupan_rekening, sumber_biaya_adm, jumlah_hari_pertahun, kode_digit, 
    is_biaya_saldo_minimum, is_biaya_rekening_dormant, is_biaya_atm, is_cetak_nota, is_status_passbook, 
    kolektibilitas, is_dapat_bagi_hasil, is_tidak_dormant, tanggal_acuan_dormant, biaya_rekening_dormant, 
    biaya_saldo_minimum, jumlah_hari_tutup_otomatis, is_blokir_debet, is_blokir_kredit, is_tutup_otomatis_dormant, 
    biaya_adm_atm, biaya_kartu_atm, is_biaya_materai, is_override_biayaatm, kode_pof_default  
from %[1]s.produk 
where kode_produk = $1

-- name: GetProdukRencana-main
select kode_produk, jangka_waktu_minimal, jangka_waktu_maksimal, jenis_rencana, minimal_setoran_rutin, minimal_target_nominal, 
    status_penarikan_saldo, biaya_penutupan, biayatutup_bawahtarget, biayatutup_atastarget, maksimal_gagal_debet, 
    akumulasi_proses_gagaldebet, debet_tanggal_ulangbulan, setor_dengan_si, target_nominal, biaya_gagal_debet, biaya_tutup_otomatis, 
    graceperiod_gagal_debet, is_autocreate_si
from %[1]s.PRODUKRENCANA 
where kode_produk = $1

-- name: GetProdukDeposito-main
select p.kode_produk, p.nama_produk, p.nisbah_bonus_dasar, p.is_backdated, coalesce(p.is_dapat_bagi_hasil,'F') as is_dapat_bagi_hasil, coalesce(p.tarif_pajak, 0.0) as tarif_pajak, coalesce(p.presentase_zakat_bagi_hasil, 0.0) as presentase_zakat_bagi_hasil, p.kode_valuta, coalesce(pd.masa_perjanjian, 0) as masa_perjanjian, coalesce(pd.periode_perjanjian, 'B') as periode_perjanjian, p.is_kena_biayalayananumum, p.is_biaya_atm, p.jenis_produk, is_biaya_rekening_dormant, is_biaya_saldo_minimum, is_blokir_debet, is_blokir_kredit, is_cetak_nota, is_status_passbook,  is_tidak_dormant
from %[1]s.coreproduct c
    inner join %[1]s.produk p on p.kode_produk = c.productcode
    left join %[1]s.produkdeposito pd on pd.kode_produk = p.kode_produk
where c.status = 'A' and p.jenis_produk='D' and p.kode_produk = $1

-- name: InsertGlInterface-main
INSERT INTO %[1]s.GLINTERFACE
(ID_GLINTERFACE, KODE_INTERFACE, DESKRIPSI, KODE_ACCOUNT)
VALUES 
( nextval('%[1]s.seq_glinterface'), $1, $2, $3)

-- name: UpdateGlInterface-main
UPDATE %[1]s.GLINTERFACE
SET
    KODE_INTERFACE = $1,
    DESKRIPSI = $2,
    KODE_ACCOUNT = $3
WHERE ID_GLINTERFACE = $4

-- name: InsertTieringItem-main
INSERT INTO %[1]s.TIERINGITEM
(ITEMID , TIERINGID, LOWERLIMITVALUE , UPPERLIMITVALUE , AMOUNT1)
VALUES 
( nextval('%[1]s.seq_tieringitem') ,$1, $2, $3, $4)

-- name: ModTieringItem-main
UPDATE %[1]s.TIERINGITEM
SET
    TIERINGID = $1,
    LOWERLIMITVALUE = $2,
    UPPERLIMITVALUE = $3,
    AMOUNT1 = $4
WHERE ITEMID = $5	

-- name: InsertListLimitProduk-main
INSERT INTO %[1]s.KETENTUAN
(ID_KETENTUAN, KODE_LIMIT ,JENIS_LIMITASI , PERIODE , JUMLAH_TRANSAKSI , NILAI_TRANSAKSI, KODE_PRODUK)
VALUES 
( nextval('%[1]s.seq_ketentuan'), $1 , $2 , $3 , $4 , $5 , $6)

-- name: ModLimitProduk-main
UPDATE %[1]s.KETENTUAN
SET
    KODE_LIMIT = $1,
    JENIS_LIMITASI = $2,
    PERIODE = $3,
    JUMLAH_TRANSAKSI = $4,
    NILAI_TRANSAKSI = $5,
    KODE_PRODUK = $6
WHERE ID_KETENTUAN = $7

-- name: ModListLimitProduk-main
UPDATE %[1]s.KETENTUAN
SET
    KODE_LIMIT = $1,
    JENIS_LIMITASI = $2,
    PERIODE = $3,
    JUMLAH_TRANSAKSI = $4,
    NILAI_TRANSAKSI = $5,
    KODE_PRODUK = $6
WHERE ID_KETENTUAN = $7	

-- name: InsertDataProdukSav-main
INSERT INTO %[1]s.PRODUK (
    KODE_PRODUK, NAMA_PRODUK, IS_ONBALANCESHEETMODE, SALDO_MINIMUM_TIDAK_AKTIF,
    JUMLAH_HARI_JADI_TIDAK_AKTIF, SALDO_MINIMUM, SALDO_MAKSIMUM, IS_BACKDATED,
    JUMLAH_BULAN_MAKS_BACKDATED, IS_KENA_ZAKAT_BAGI_HASIL,
    PRESENTASE_ZAKAT_BAGI_HASIL, IS_KENA_PAJAK, TARIF_PAJAK, DISPOSISI_BAGI_HASIL,
    PERIODE_BAGI_HASIL, JENIS_PRODUK, KODE_VALUTA, STATUS, IS_PRODUK_BERASURANSI,
    IS_PRODUK_PARTNERSHIP, IS_OVERRIDE_BAGIHASILKHUSUS, IS_OVERRIDE_TARIFPAJAK,
    IS_OVERRIDE_ZAKATBAGIHASIL, IS_OVERRIDE_DISPOSISI_BGH, JARAK_PERIODE_BAGIHASIL,
    IS_OVERRIDE_BACKDATED, ID_RENCANABAGIHASILDEFAULT, IS_KENA_BIAYALAYANANUMUM,
    ID_BIAYALAYANANUMUM, NAMA_VALUTA, IS_PRODUKKARTU, NAMA_ASURANSI, ID_PARTNER,
    NAMA_GRUP, INSPLANID, PSPLANID, BIAYA_ADM_BULANAN, ID_TIERING_BIAYAADM,
    ID_TIERING_NISBAHBONUS, IS_TIERING_BIAYAADM, IS_TIERING_NISBAHBONUS, JENIS_AKAD,
    KODE_RENCANA_ASURANSI, NISBAH_BONUS_DASAR, SETORAN_AWAL_MINIMUM, ID_ASURANSI,
    IS_PARAM_SALDO_TIDAK_AKTIF, IS_OVERRIDE_BIAYA_PENUTUPAN, BIAYA_PENUTUPAN_REKENING,
    SUMBER_BIAYA_ADM, JUMLAH_HARI_PERTAHUN, KODE_DIGIT, IS_BIAYA_SALDO_MINIMUM,
    IS_BIAYA_REKENING_DORMANT, IS_BIAYA_ATM, IS_CETAK_NOTA, IS_STATUS_PASSBOOK,
    KOLEKTIBILITAS, IS_DAPAT_BAGI_HASIL, IS_TIDAK_DORMANT, TANGGAL_ACUAN_DORMANT,
    BIAYA_REKENING_DORMANT, BIAYA_SALDO_MINIMUM, JUMLAH_HARI_TUTUP_OTOMATIS,
    IS_BLOKIR_DEBET, IS_BLOKIR_KREDIT, IS_TUTUP_OTOMATIS_DORMANT, BIAYA_ADM_ATM,
    BIAYA_KARTU_ATM, IS_BIAYA_MATERAI, IS_OVERRIDE_BIAYAATM, KODE_POF_DEFAULT
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, 'T', $16, $17, $18, $19, $20, $21, $22, $23, $24,
    $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48,
    $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71
)

-- name: InsertDataProdukDep-main
INSERT INTO %[1]s.PRODUK (
    KODE_PRODUK, NAMA_PRODUK, IS_ONBALANCESHEETMODE, SALDO_MINIMUM_TIDAK_AKTIF,
    JUMLAH_HARI_JADI_TIDAK_AKTIF, SALDO_MINIMUM, SALDO_MAKSIMUM, IS_BACKDATED,
    JUMLAH_BULAN_MAKS_BACKDATED, IS_KENA_ZAKAT_BAGI_HASIL,
    PRESENTASE_ZAKAT_BAGI_HASIL, IS_KENA_PAJAK, TARIF_PAJAK, DISPOSISI_BAGI_HASIL,
    PERIODE_BAGI_HASIL, JENIS_PRODUK, KODE_VALUTA, STATUS, IS_PRODUK_BERASURANSI,
    IS_PRODUK_PARTNERSHIP, IS_OVERRIDE_BAGIHASILKHUSUS, IS_OVERRIDE_TARIFPAJAK,
    IS_OVERRIDE_ZAKATBAGIHASIL, IS_OVERRIDE_DISPOSISI_BGH, JARAK_PERIODE_BAGIHASIL,
    IS_OVERRIDE_BACKDATED, ID_RENCANABAGIHASILDEFAULT, IS_KENA_BIAYALAYANANUMUM,
    ID_BIAYALAYANANUMUM, NAMA_VALUTA, IS_PRODUKKARTU, NAMA_ASURANSI, ID_PARTNER,
    NAMA_GRUP, INSPLANID, PSPLANID, BIAYA_ADM_BULANAN, ID_TIERING_BIAYAADM,
    ID_TIERING_NISBAHBONUS, IS_TIERING_BIAYAADM, IS_TIERING_NISBAHBONUS, JENIS_AKAD,
    KODE_RENCANA_ASURANSI, NISBAH_BONUS_DASAR, SETORAN_AWAL_MINIMUM, ID_ASURANSI,
    IS_PARAM_SALDO_TIDAK_AKTIF, IS_OVERRIDE_BIAYA_PENUTUPAN, BIAYA_PENUTUPAN_REKENING,
    SUMBER_BIAYA_ADM, JUMLAH_HARI_PERTAHUN, KODE_DIGIT, IS_BIAYA_SALDO_MINIMUM,
    IS_BIAYA_REKENING_DORMANT, IS_BIAYA_ATM, IS_CETAK_NOTA, IS_STATUS_PASSBOOK,
    KOLEKTIBILITAS, IS_DAPAT_BAGI_HASIL, IS_TIDAK_DORMANT, TANGGAL_ACUAN_DORMANT,
    BIAYA_REKENING_DORMANT, BIAYA_SALDO_MINIMUM, JUMLAH_HARI_TUTUP_OTOMATIS,
    IS_BLOKIR_DEBET, IS_BLOKIR_KREDIT, IS_TUTUP_OTOMATIS_DORMANT, BIAYA_ADM_ATM,
    BIAYA_KARTU_ATM, IS_BIAYA_MATERAI, IS_OVERRIDE_BIAYAATM, KODE_POF_DEFAULT
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, 'D', $16, $17, $18, $19, $20, $21, $22, $23, $24,
    $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48,
    $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71
)

-- name: InsertDataProdukGiro-main
INSERT INTO %[1]s.PRODUK (
		KODE_PRODUK, NAMA_PRODUK, IS_ONBALANCESHEETMODE, SALDO_MINIMUM_TIDAK_AKTIF,
    JUMLAH_HARI_JADI_TIDAK_AKTIF, SALDO_MINIMUM, SALDO_MAKSIMUM, IS_BACKDATED,
    JUMLAH_BULAN_MAKS_BACKDATED, IS_KENA_ZAKAT_BAGI_HASIL,
    PRESENTASE_ZAKAT_BAGI_HASIL, IS_KENA_PAJAK, TARIF_PAJAK, DISPOSISI_BAGI_HASIL,
    PERIODE_BAGI_HASIL, JENIS_PRODUK, KODE_VALUTA, STATUS, IS_PRODUK_BERASURANSI,
    IS_PRODUK_PARTNERSHIP, IS_OVERRIDE_BAGIHASILKHUSUS, IS_OVERRIDE_TARIFPAJAK,
    IS_OVERRIDE_ZAKATBAGIHASIL, IS_OVERRIDE_DISPOSISI_BGH, JARAK_PERIODE_BAGIHASIL,
    IS_OVERRIDE_BACKDATED, ID_RENCANABAGIHASILDEFAULT, IS_KENA_BIAYALAYANANUMUM,
    ID_BIAYALAYANANUMUM, NAMA_VALUTA, IS_PRODUKKARTU, NAMA_ASURANSI, ID_PARTNER,
    NAMA_GRUP, INSPLANID, PSPLANID, BIAYA_ADM_BULANAN, ID_TIERING_BIAYAADM,
    ID_TIERING_NISBAHBONUS, IS_TIERING_BIAYAADM, IS_TIERING_NISBAHBONUS, JENIS_AKAD,
    KODE_RENCANA_ASURANSI, NISBAH_BONUS_DASAR, SETORAN_AWAL_MINIMUM, ID_ASURANSI,
    IS_PARAM_SALDO_TIDAK_AKTIF, IS_OVERRIDE_BIAYA_PENUTUPAN, BIAYA_PENUTUPAN_REKENING,
    SUMBER_BIAYA_ADM, JUMLAH_HARI_PERTAHUN, KODE_DIGIT, IS_BIAYA_SALDO_MINIMUM,
    IS_BIAYA_REKENING_DORMANT, IS_BIAYA_ATM, IS_CETAK_NOTA, IS_STATUS_PASSBOOK,
    KOLEKTIBILITAS, IS_DAPAT_BAGI_HASIL, IS_TIDAK_DORMANT, TANGGAL_ACUAN_DORMANT,
    BIAYA_REKENING_DORMANT, BIAYA_SALDO_MINIMUM, JUMLAH_HARI_TUTUP_OTOMATIS,
    IS_BLOKIR_DEBET, IS_BLOKIR_KREDIT, IS_TUTUP_OTOMATIS_DORMANT, BIAYA_ADM_ATM,
    BIAYA_KARTU_ATM, IS_BIAYA_MATERAI, IS_OVERRIDE_BIAYAATM, KODE_POF_DEFAULT
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, 'G', $16, $17, $18, $19, $20, $21, $22, $23, $24, $25,
    $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49,
    $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71
)

-- name: InsertProdukDep-main
INSERT INTO %[1]s.PRODUKDEPOSITO (
    KODE_PRODUK, MASA_PERJANJIAN,MASA_PERJANJIAN_MAKSIMAL ,MASA_PERJANJIAN_MINIMAL ,
    IS_MASA_PERJANJIAN_BEBAS, IS_KENA_PENALTY ,NOMINAL_PENALTY 
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)

-- name: UpdateProdukDep-main
UPDATE %[1]s.PRODUKDEPOSITO
SET MASA_PERJANJIAN = $1,
    MASA_PERJANJIAN_MAKSIMAL = $2,
    MASA_PERJANJIAN_MINIMAL = $3,
    IS_MASA_PERJANJIAN_BEBAS = $4,
    IS_KENA_PENALTY = $5,
    NOMINAL_PENALTY = $6
WHERE KODE_PRODUK = $7

-- name: GetSeqIdTieringNisbahbonus-main
select  nextval('%[1]s.seq_tieringplan') as id_tiering_nisbahbonus


-- name: GetSeqIdTieringBiayaAdmin-main
select  nextval('%[1]s.seq_tieringplan') as id_tiering_biayaadm


-- name: UpdateIdTeringNisbahBonus-main
update %[1]s.produk
set id_tiering_nisbahbonus = $1
where kode_produk = $2

-- name: UpdateIdTieringBiayaAdmin-main
update %[1]s.produk
set id_tiering_biayaadm = $1
where kode_produk = $2

-- name: InsertTieringPlan-main
insert into %[1]s.tieringplan(tieringid, tieringtype)
values($1, $1)

-- name: GetKodeProdukExist-main
SELECT p.KODE_PRODUK 
FROM %[1]s.produk p WHERE p.KODE_PRODUK = $1

-- name: GetKodeProdukDepExist-main
SELECT p.KODE_PRODUK 
FROM %[1]s.PRODUKDEPOSITO p WHERE p.KODE_PRODUK = $1

-- name: ModDataProduk-main
UPDATE %[1]s.PRODUK
SET 
    NAMA_PRODUK = $1, 
    IS_ONBALANCESHEETMODE = $2, 
    SALDO_MINIMUM_TIDAK_AKTIF = $3, 
    JUMLAH_HARI_JADI_TIDAK_AKTIF = $4, 
    SALDO_MINIMUM = $5, 
    SALDO_MAKSIMUM = $6, 
    IS_BACKDATED = $7, 
    JUMLAH_BULAN_MAKS_BACKDATED = $8, 
    IS_KENA_ZAKAT_BAGI_HASIL = $9, 
    PRESENTASE_ZAKAT_BAGI_HASIL = $10, 
    IS_KENA_PAJAK = $11, 
    TARIF_PAJAK = $12, 
    DISPOSISI_BAGI_HASIL = $13, 
    PERIODE_BAGI_HASIL = $14,  
    STATUS = $15, 
    IS_PRODUK_BERASURANSI = $16, 
    IS_PRODUK_PARTNERSHIP = $17, 
    IS_OVERRIDE_BAGIHASILKHUSUS = $18, 
    IS_OVERRIDE_TARIFPAJAK = $19, 
    IS_OVERRIDE_ZAKATBAGIHASIL = $20, 
    IS_OVERRIDE_DISPOSISI_BGH = $21, 
    JARAK_PERIODE_BAGIHASIL = $22, 
    IS_OVERRIDE_BACKDATED = $23, 
    ID_RENCANABAGIHASILDEFAULT = $24, 
    IS_KENA_BIAYALAYANANUMUM = $25, 
    ID_BIAYALAYANANUMUM = $26, 
    NAMA_VALUTA = $27, 
    IS_PRODUKKARTU = $28, 
    NAMA_ASURANSI = $29, 
    ID_PARTNER = $30, 
    NAMA_GRUP = $31, 
    INSPLANID = $32, 
    PSPLANID = $33, 
    BIAYA_ADM_BULANAN = $34, 
    ID_TIERING_BIAYAADM = $35, 
    ID_TIERING_NISBAHBONUS = $36, 
    IS_TIERING_BIAYAADM = $37, 
    IS_TIERING_NISBAHBONUS = $38, 
    JENIS_AKAD = $39, 
    KODE_RENCANA_ASURANSI = $40, 
    NISBAH_BONUS_DASAR = $41, 
    SETORAN_AWAL_MINIMUM = $42, 
    ID_ASURANSI = $43, 
    IS_PARAM_SALDO_TIDAK_AKTIF = $44, 
    IS_OVERRIDE_BIAYA_PENUTUPAN = $45, 
    BIAYA_PENUTUPAN_REKENING = $46, 
    SUMBER_BIAYA_ADM = $47, 
    JUMLAH_HARI_PERTAHUN = $48, 
    KODE_DIGIT = $49, 
    IS_BIAYA_SALDO_MINIMUM = $50, 
    IS_BIAYA_REKENING_DORMANT = $51, 
    IS_BIAYA_ATM = $52, 
    IS_CETAK_NOTA = $53, 
    IS_STATUS_PASSBOOK = $54, 
    KOLEKTIBILITAS = $55, 
    IS_DAPAT_BAGI_HASIL = $56, 
    IS_TIDAK_DORMANT = $57, 
    TANGGAL_ACUAN_DORMANT = $58, 
    BIAYA_REKENING_DORMANT = $59, 
    BIAYA_SALDO_MINIMUM = $60, 
    JUMLAH_HARI_TUTUP_OTOMATIS = $61, 
    IS_BLOKIR_DEBET = $62, 
    IS_BLOKIR_KREDIT = $63, 
    IS_TUTUP_OTOMATIS_DORMANT = $64, 
    BIAYA_ADM_ATM = $65, 
    BIAYA_KARTU_ATM = $66, 
    IS_BIAYA_MATERAI = $67, 
    IS_OVERRIDE_BIAYAATM = $68, 
    KODE_POF_DEFAULT = $69
WHERE KODE_PRODUK = $70

-- name: DelTieringItem-main
DELETE FROM %[1]s.TIERINGITEM t WHERE t.ITEMID = $1

-- name: DelLimitProd-main
DELETE FROM %[1]s.KETENTUAN k WHERE k.ID_KETENTUAN = $1

-- name: GetProfitaccrualList-main
select id_accrual, nomor_rekening, to_char(tanggal, 'YYYY-MM-DD') as tanggal, round(saldo, 2) as saldo, round(amount, 2) as amount,
    round(nisbah, 2) as nisbah, round(gdrval, 2) as gdrval, jenis_akad
from %[1]s.profitaccrual p
where p.nomor_rekening = $1
    and p.tanggal >= to_date($2, 'YYYY-MM-DD HH24::MI::SS')
    and p.tanggal < to_timestamp($3, 'YYYY-MM-DD HH24::MI::SS') + 1

-- name: GetRcCodeList-main
SELECT PROJECT_NO AS rc_code, name AS rc_name FROM %[1]s.project

-- name: GetListRas-main
select c.kode_cabang, r.nomor_rekening , r.saldo , r.nama_rekening , c.nama_cabang
from %[1]s.rekeningtransaksi r
    inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = r.nomor_rekening
    inner join %[2]s.cabang c on c.kode_cabang = r.kode_cabang 
where rl.jenis_rekening_liabilitas = $1
    %[3]s
order by r.nomor_rekening 
offset $2 rows fetch next $3 rows only

-- name: GetListRas-filter-KodeCabang
and r.kode_cabang = $4

-- name: NonactiveRecurrenttrx-main
update  %[1]s.recurrenttransaction 
set status_data = 'N'
where recid= $1

-- name: GetCifCorpRedundantList-main
SELECT n.kode_cabang_input
, n.nomor_nasabah
, n.is_wic
, n.tanggal_buka
, coalesce(qr.pembiayaan, 0) as pembiayaan
, coalesce(qr.tahapan, 0) as tahapan
, coalesce(qr.deposito, 0) as deposito
, coalesce(qr.trib, 0) as trib
, coalesce(qr.giro, 0) as giro
, nk.nama_perusahaan
, nk.nomor_npwp
, nk.nomor_siupp
, nk.tanggal_siupp
, nk.SKDP_Nomor
, nk.Akte_Pendirian_Nomor
, q.groupid
, %[2]v
FROM
    %[1]v.NASABAH n
    INNER JOIN %[1]v.NASABAHKORPORAT nk ON nk.nomor_nasabah = n.nomor_nasabah
    INNER JOIN (
    SELECT %[4]v
        , COUNT(1)
        , DENSE_RANK() OVER (ORDER BY %[3]v) AS groupid
        FROM %[1]v.NASABAHKORPORAT nk
        INNER JOIN %[1]v.NASABAH n ON n.nomor_nasabah = nk.nomor_nasabah
        WHERE status_data = 'A'
        GROUP BY %[3]v
    HAVING COUNT(1) > 1
    ) q ON %[5]v
    LEFT JOIN (
    SELECT rc.nomor_nasabah
        , SUM(
            CASE WHEN rt.kode_jenis IN ('IJR','MDB','MBH','QRD','MSY','PRK','WAD','ISH','KFL')
            THEN 1 ELSE 0 END) AS pembiayaan
        , SUM(
            CASE WHEN rl.jenis_rekening_liabilitas = 'T'
            THEN 1 ELSE 0 END) AS tahapan
        , SUM(
            CASE WHEN rl.jenis_rekening_liabilitas = 'D'
            THEN 1 ELSE 0 END) AS deposito
        , SUM(
            CASE WHEN rl.jenis_rekening_liabilitas = 'R'
            THEN 1 ELSE 0 END) AS trib
        , SUM(
            CASE WHEN rl.jenis_rekening_liabilitas = 'G'
            THEN 1 ELSE 0 END) AS giro
    FROM %[1]v.REKENINGTRANSAKSI rt
        INNER JOIN %[1]v.REKENINGCUSTOMER rc ON rc.nomor_rekening = rt.nomor_rekening
        LEFT JOIN %[1]v.REKENINGLIABILITAS rl ON rl.nomor_rekening = rt.nomor_rekening
    WHERE rt.kode_jenis IN ('IJR','MDB','MBH','QRD','MSY','PRK','WAD','ISH','KFL','SAV','DEP', 'CA')
    GROUP BY rc.nomor_nasabah
    ) qr ON qr.nomor_nasabah = n.nomor_nasabah
WHERE status_data='A'
ORDER BY
    q.groupid, n.nomor_nasabah

-- name: InsertUpdateRefSched-main
insert into %[1]s.update_ref_sched (
    id_sched, action_sched, nomor_rekening, nomor_nasabah, tanggal_update, user_update, status_update, keyint
) 
values (
     nextval('%[1]s.seq_update_ref_sched'), $1, $2, $3, current_timestamp, $4, $5, $6
)

-- name: UpdateIdMulticifrelation-main
update %[1]s.rekeningliabilitas rl
set id_multicifrelation = (
    select rd.refdata_id 
    from %[2]s.referencedata rd
        inner join %[2]s.referencetype rt on rt.reftype_id = rd.reftype_id and rt.reference_name = 'R_RELASI_MULTICIF'
    where rd.reference_code = rl.kode_multicifrelation 
)
where rl.nomor_rekening = $1

-- name: UpdateIdSumberDana-main
update %[1]s.rekeningliabilitas rl
set id_sumber_dana = (
    select rd.refdata_id 
    from %[2]s.referencedata rd
        inner join %[2]s.referencetype rt on rt.reftype_id = rd.reftype_id and rt.reference_name = 'R_SUMBER_DANA_REKENING'
    where rd.reference_code = rl.kode_sumber_dana 
)
where rl.nomor_rekening = $1

-- name: UpdateIdTujuanRekening-main
update %[1]s.rekeningliabilitas rl
set id_tujuan_rekening = (
    select rd.refdata_id 
    from %[2]s.referencedata rd
        inner join %[2]s.referencetype rt on rt.reftype_id = rd.reftype_id and rt.reference_name = 'R_TUJUAN_BUKA_REKENING'
    where rd.reference_code = rl.kode_tujuan_rekening 
)
where rl.nomor_rekening = $1

-- name: UpdateIdCfgroup-main
update %[1]s.rekeningliabilitas rl
set id_cfgroup = (
    select id_cfgroup  
    from %[1]s.confidentialgroup cg
    where cg.cf_code = rl.kode_status_retriksi  
)
where rl.nomor_rekening = $1

-- name: UpdateIdSumberDanaTrx-main
update %[1]s.infotransaksi i
set id_sumber_dana_trx = (
    select rd.refdata_id 
    from %[2]s.referencedata rd
        inner join %[2]s.referencetype rt on rt.reftype_id = rd.reftype_id and rt.reference_name = 'R_SUMBER_DANA_TRX'
    where rd.reference_code = i.kd_sumber_dana_trx 
)
where i.id_transaksi = $1

-- name: UpdateIdTujuanTrx-main
update %[1]s.infotransaksi i
set id_tujuan_trx = (
    select rd.refdata_id 
    from %[2]s.referencedata rd
        inner join %[2]s.referencetype rt on rt.reftype_id = rd.reftype_id and rt.reference_name = 'R_TUJUAN_TRX'
    where rd.reference_code = i.kd_tujuan_trx 
)
where i.id_transaksi = $1

-- name: UpdateIdNegara-main
update %[1]s.individu i
set id_negara = (
    select rd.refdata_id 
    from %[2]s.referencedata rd
        inner join %[2]s.referencetype rt on rt.reftype_id = rd.reftype_id and rt.reference_name = 'R_KODE_NEGARA'
    where rd.reference_code = i.kode_negara_asal 
)
where i.id_individu = $1

-- name: CreateRegisterLayanan-main
insert into %[1]s.registerlayanan
(id_register, id_layanan, biaya_layanan, nomor_rekening_layanan, status_register_layanan, tanggal_registrasi, user_pembuat, 
    user_otorisasi, tanggal_otorisasi, kode_cabang_input, jenis_register_layanan)
values 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)

-- name: GetRegisterLayanan-main
select id_register, nomor_rekening_layanan, status_register_layanan, jenis_register_layanan 
from %[1]s.registerlayanan
where id_register = $1

-- name: UpdateRegisterLayanan-main
update %[1]s.registerlayanan
set status_register_layanan = $1, tanggal_nonaktif = current_timestamp, user_otorisasi_nonaktif = $2, 
    user_input_nonaktif = $3, kode_cabang_nonaktif = $4
where id_register = $5

-- name: NonActiveRegisterLayanan-main
update %[1]s.registerlayanan
set status_register_layanan = $1, tanggal_nonaktif = current_timestamp, user_otorisasi_nonaktif = $2, 
    user_input_nonaktif = $3, kode_cabang_nonaktif = $4
where id_register = $5

-- name: CreateRegPassbook-main
insert into %[1]s.registerpassbook
(id_register, halaman_cetak_terakhir, baris_cetak_terakhir, nomor_buku_terakhir, is_baru_register)
values 
($1, 0, 0, 0, $2)

-- name: GetRegisterRt-main
select recid, id_register, description, nomor_rekening_kredit
from %[1]s.registerrt
where id_register = $1

-- name: GetListProdSav-main
select p.kode_produk, p.nama_produk, p.is_onbalancesheetmode, p.saldo_minimum_tidak_aktif, p.jumlah_hari_jadi_tidak_aktif, p.saldo_minimum, 
    p.saldo_maksimum, p.is_backdated, p.jumlah_bulan_maks_backdated, p.is_kena_zakat_bagi_hasil, p.presentase_zakat_bagi_hasil, 
    p.is_kena_pajak, p.tarif_pajak, p.disposisi_bagi_hasil, p.periode_bagi_hasil, p.jenis_produk, p.kode_valuta, p.nama_valuta, p.status, 
    p.is_produk_berasuransi, p.is_produk_partnership, p.is_override_bagihasilkhusus, p.is_override_tarifpajak, p.is_override_zakatbagihasil, 
    p.is_override_disposisi_bgh, p.jarak_periode_bagihasil, p.is_override_backdated, p.is_kena_biayalayananumum, p.is_produkkartu, 
    p.nama_asuransi, p.id_partner, p.nama_grup, p.insplanid, p.psplanid, p.biaya_adm_bulanan, p.is_tiering_biayaadm, p.is_tiering_nisbahbonus, 
    p.jenis_akad, p.kode_rencana_asuransi, p.nisbah_bonus_dasar, p.setoran_awal_minimum, p.is_param_saldo_tidak_aktif, p.is_override_biaya_penutupan, 
    p.biaya_penutupan_rekening, p.sumber_biaya_adm, p.jumlah_hari_pertahun, p.kode_digit, p.is_blokir_debet_internal, p.is_blokir_debet_eksternal, 
    p.is_blokir_kredit_internal, p.is_blokir_kredit_eksternal, p.is_biaya_saldo_minimum, p.is_biaya_rekening_dormant, p.is_biaya_atm, p.is_cetak_nota, 
    p.is_status_passbook, p.kolektibilitas, coalesce(p.is_dapat_bagi_hasil, 'F') as is_dapat_bagi_hasil, p.is_tidak_dormant, p.tanggal_acuan_dormant, 
    p.biaya_rekening_dormant, p.biaya_saldo_minimum, p.jumlah_hari_tutup_otomatis, p.is_blokir_debet, p.is_blokir_kredit, 
    p.is_tutup_otomatis_dormant, p.biaya_adm_atm, p.biaya_kartu_atm, p.ekuivalen_rate, p.saldo_minimum_bagihasil, p.is_proses_cadangan, 
    p.persentase_cadangan, p.order_mb, p.is_biaya_materai, p.is_override_biayaatm, p.id_rencanabagihasildefault, p.id_biayalayananumum, 
    p.id_tiering_biayaadm, p.id_tiering_nisbahbonus, p.id_asuransi, pc.kode_cabang
from %[1]s.coreproduct c
    inner join %[1]s.produk p on p.kode_produk = c.productcode
    inner join %[1]s.produkcabang pc on pc.kode_produk = p.kode_produk
where c.status = 'A' and p.jenis_produk='T' and pc.kode_cabang = $1
order by p.kode_produk

-- name: GetListSavPlan-main
select p.kode_produk, p.nama_produk, p.is_onbalancesheetmode, p.saldo_minimum_tidak_aktif, p.jumlah_hari_jadi_tidak_aktif, p.saldo_minimum, 
    p.saldo_maksimum, p.is_backdated, p.jumlah_bulan_maks_backdated, p.is_kena_zakat_bagi_hasil, p.presentase_zakat_bagi_hasil, 
    p.is_kena_pajak, p.tarif_pajak, p.disposisi_bagi_hasil, p.periode_bagi_hasil, p.jenis_produk, p.kode_valuta, p.nama_valuta, p.status, 
    p.is_produk_berasuransi, p.is_produk_partnership, p.is_override_bagihasilkhusus, p.is_override_tarifpajak, p.is_override_zakatbagihasil, 
    p.is_override_disposisi_bgh, p.jarak_periode_bagihasil, p.is_override_backdated, p.is_kena_biayalayananumum, p.is_produkkartu, 
    p.nama_asuransi, p.id_partner, p.nama_grup, p.insplanid, p.psplanid, p.biaya_adm_bulanan, p.is_tiering_biayaadm, p.is_tiering_nisbahbonus, 
    p.jenis_akad, p.kode_rencana_asuransi, p.nisbah_bonus_dasar, p.setoran_awal_minimum, p.is_param_saldo_tidak_aktif, p.is_override_biaya_penutupan, 
    p.biaya_penutupan_rekening, p.sumber_biaya_adm, p.jumlah_hari_pertahun, p.kode_digit, p.is_blokir_debet_internal, p.is_blokir_debet_eksternal, 
    p.is_blokir_kredit_internal, p.is_blokir_kredit_eksternal, p.is_biaya_saldo_minimum, p.is_biaya_rekening_dormant, p.is_biaya_atm, p.is_cetak_nota, 
    p.is_status_passbook, p.kolektibilitas, coalesce(p.is_dapat_bagi_hasil, 'F') as is_dapat_bagi_hasil, p.is_tidak_dormant, p.tanggal_acuan_dormant, 
    p.biaya_rekening_dormant, p.biaya_saldo_minimum, p.jumlah_hari_tutup_otomatis, p.is_blokir_debet, p.is_blokir_kredit, 
    p.is_tutup_otomatis_dormant, p.biaya_adm_atm, p.biaya_kartu_atm, p.ekuivalen_rate, p.saldo_minimum_bagihasil, p.is_proses_cadangan, 
    p.persentase_cadangan, p.order_mb, p.is_biaya_materai, p.is_override_biayaatm, p.id_rencanabagihasildefault, p.id_biayalayananumum, 
    p.id_tiering_biayaadm, p.id_tiering_nisbahbonus, p.id_asuransi, pc.kode_cabang
from %[1]s.coreproduct c
    inner join %[1]s.produk p on p.kode_produk = c.productcode
    inner join %[1]s.produkcabang pc on pc.kode_produk = p.kode_produk
where c.status = 'A' and p.jenis_produk='R' and pc.kode_cabang = $1
order by p.kode_produk

-- name: GetListDep-main
select p.kode_produk, 
    p.nama_produk,
    p.nisbah_bonus_dasar,
    p.nama_valuta,
    coalesce(p.is_dapat_bagi_hasil,'F') as is_dapat_bagi_hasil,
    coalesce(p.tarif_pajak, 0.0) as tarif_pajak,
    coalesce(p.presentase_zakat_bagi_hasil, 0.0) as presentase_zakat_bagi_hasil,
    p.kode_valuta,
    p.is_kena_biayalayananumum,
    p.is_biaya_atm,
    p.jenis_produk,
    pc.kode_cabang
from %[1]s.coreproduct c
    inner join %[1]s.produk p on p.kode_produk = c.productcode
    inner join %[1]s.produkcabang pc on pc.kode_produk = p.kode_produk
where c.status = 'A' and p.jenis_produk='T' and pc.kode_cabang = $1
order by p.kode_produk

-- name: GetProdDepList-main
select p.kode_produk, p.nama_produk, p.is_onbalancesheetmode, p.saldo_minimum_tidak_aktif, p.jumlah_hari_jadi_tidak_aktif, p.saldo_minimum, 
    p.saldo_maksimum, p.is_backdated, p.jumlah_bulan_maks_backdated, p.is_kena_zakat_bagi_hasil, p.presentase_zakat_bagi_hasil, 
    p.is_kena_pajak, p.tarif_pajak, p.disposisi_bagi_hasil, p.periode_bagi_hasil, p.jenis_produk, p.kode_valuta, p.nama_valuta, p.status, 
    p.is_produk_berasuransi, p.is_produk_partnership, p.is_override_bagihasilkhusus, p.is_override_tarifpajak, p.is_override_zakatbagihasil, 
    p.is_override_disposisi_bgh, p.jarak_periode_bagihasil, p.is_override_backdated, p.is_kena_biayalayananumum, p.is_produkkartu, 
    p.nama_asuransi, p.id_partner, p.nama_grup, p.insplanid, p.psplanid, p.biaya_adm_bulanan, p.is_tiering_biayaadm, p.is_tiering_nisbahbonus, 
    p.jenis_akad, p.kode_rencana_asuransi, p.nisbah_bonus_dasar, p.setoran_awal_minimum, p.is_param_saldo_tidak_aktif, p.is_override_biaya_penutupan, 
    p.biaya_penutupan_rekening, p.sumber_biaya_adm, p.jumlah_hari_pertahun, p.kode_digit, p.is_blokir_debet_internal, p.is_blokir_debet_eksternal, 
    p.is_blokir_kredit_internal, p.is_blokir_kredit_eksternal, p.is_biaya_saldo_minimum, p.is_biaya_rekening_dormant, p.is_biaya_atm, p.is_cetak_nota, 
    p.is_status_passbook, p.kolektibilitas, coalesce(p.is_dapat_bagi_hasil, 'F') as is_dapat_bagi_hasil, p.is_tidak_dormant, p.tanggal_acuan_dormant, 
    p.biaya_rekening_dormant, p.biaya_saldo_minimum, p.jumlah_hari_tutup_otomatis, p.is_blokir_debet, p.is_blokir_kredit, 
    p.is_tutup_otomatis_dormant, p.biaya_adm_atm, p.biaya_kartu_atm, p.ekuivalen_rate, p.saldo_minimum_bagihasil, p.is_proses_cadangan, 
    p.persentase_cadangan, p.order_mb, p.is_biaya_materai, p.is_override_biayaatm, p.id_rencanabagihasildefault, p.id_biayalayananumum, 
    p.id_tiering_biayaadm, p.id_tiering_nisbahbonus, p.id_asuransi, pc.kode_cabang,
    coalesce(pd.masa_perjanjian, 0) as masa_perjanjian, coalesce(pd.periode_perjanjian, 'B') as periode_perjanjian
from %[1]s.coreproduct c
    inner join %[1]s.produk p on p.kode_produk = c.productcode
    inner join %[1]s.produkcabang pc on pc.kode_produk = p.kode_produk
    left join %[1]s.produkdeposito pd on pd.kode_produk = p.kode_produk
where c.status = 'A' and p.jenis_produk='D' and pc.kode_cabang = $1
order by p.kode_produk

-- name: GetListCA-main
select p.kode_produk, p.nama_produk, p.is_onbalancesheetmode, p.saldo_minimum_tidak_aktif, p.jumlah_hari_jadi_tidak_aktif, p.saldo_minimum, 
    p.saldo_maksimum, p.is_backdated, p.jumlah_bulan_maks_backdated, p.is_kena_zakat_bagi_hasil, p.presentase_zakat_bagi_hasil, 
    p.is_kena_pajak, p.tarif_pajak, p.disposisi_bagi_hasil, p.periode_bagi_hasil, p.jenis_produk, p.kode_valuta, p.nama_valuta, p.status, 
    p.is_produk_berasuransi, p.is_produk_partnership, p.is_override_bagihasilkhusus, p.is_override_tarifpajak, p.is_override_zakatbagihasil, 
    p.is_override_disposisi_bgh, p.jarak_periode_bagihasil, p.is_override_backdated, p.is_kena_biayalayananumum, p.is_produkkartu, 
    p.nama_asuransi, p.id_partner, p.nama_grup, p.insplanid, p.psplanid, p.biaya_adm_bulanan, p.is_tiering_biayaadm, p.is_tiering_nisbahbonus, 
    p.jenis_akad, p.kode_rencana_asuransi, p.nisbah_bonus_dasar, p.setoran_awal_minimum, p.is_param_saldo_tidak_aktif, p.is_override_biaya_penutupan, 
    p.biaya_penutupan_rekening, p.sumber_biaya_adm, p.jumlah_hari_pertahun, p.kode_digit, p.is_blokir_debet_internal, p.is_blokir_debet_eksternal, 
    p.is_blokir_kredit_internal, p.is_blokir_kredit_eksternal, p.is_biaya_saldo_minimum, p.is_biaya_rekening_dormant, p.is_biaya_atm, p.is_cetak_nota, 
    p.is_status_passbook, p.kolektibilitas, coalesce(p.is_dapat_bagi_hasil, 'F') as is_dapat_bagi_hasil, p.is_tidak_dormant, p.tanggal_acuan_dormant, 
    p.biaya_rekening_dormant, p.biaya_saldo_minimum, p.jumlah_hari_tutup_otomatis, p.is_blokir_debet, p.is_blokir_kredit, 
    p.is_tutup_otomatis_dormant, p.biaya_adm_atm, p.biaya_kartu_atm, p.ekuivalen_rate, p.saldo_minimum_bagihasil, p.is_proses_cadangan, 
    p.persentase_cadangan, p.order_mb, p.is_biaya_materai, p.is_override_biayaatm, p.id_rencanabagihasildefault, p.id_biayalayananumum, 
    p.id_tiering_biayaadm, p.id_tiering_nisbahbonus, p.id_asuransi, pc.kode_cabang
from %[1]s.coreproduct c
    inner join %[1]s.produk p on p.kode_produk = c.productcode
    inner join %[1]s.produkcabang pc on pc.kode_produk = p.kode_produk
where c.status = 'A' and p.jenis_produk='G' and pc.kode_cabang = $1
order by p.kode_produk

-- name: GetAccountLiabilitas-main
select rl.nomor_nasabah, rt.nomor_rekening ,rt.nama_rekening, rt.status_rekening, rt.kode_cabang, rl.is_sedang_ditutup, rl.is_blokir, rl.kode_produk, rl.jenis_rekening_liabilitas
from %[1]s.rekeningliabilitas rl
    inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
    inner join %[1]s.produk p on p.kode_produk = rl.kode_produk
where rl.nomor_rekening = $1

-- name: UpdateAcctDormat-main
update %[1]s.rekeningtransaksi 
set status_rekening = 1,
    user_perbarui = $1,
    tanggal_perbarui = current_timestamp
where nomor_rekening = $2

-- name: GetLimitProduk-main
select kode_limit ,jenis_limitasi , periode , jumlah_transaksi , nilai_transaksi   
from %[1]s.ketentuan k
where kode_produk = $1

-- name: GetRekeningTabCollateral-main-D
SELECT
    rl.nomor_rekening,
    rt.nama_rekening,
    rt.saldo
FROM %[1]s.deposito dp, %[1]s.rekeningtransaksi rt, %[1]s.rekeningliabilitas rl
WHERE rl.nomor_rekening = rt.nomor_rekening
    AND rt.nomor_rekening = dp.nomor_rekening
    AND rl.is_blokir <> 'T'
    AND dp.disposisi_nominal = 'A'
    AND rt.status_rekening <> '3'  
    AND %[2]s

-- name: GetRekeningTabCollateral-filter-D
(rl.nomor_rekening LIKE $1 or upper(rt.nama_rekening) LIKE $1)

-- name: GetRekeningTabCollateral-main-GTJ
SELECT
    l.nomor_rekening, rt.nama_rekening,        
    case when l.status_restriksi='T' then 9999999999 else rt.saldo end saldo 
    FROM %[1]s.rekeningliabilitas l, %[1]s.rekeningtransaksi rt
WHERE      
    l.nomor_rekening=rt.nomor_rekening
    AND l.jenis_rekening_liabilitas in ('G', 'T','P')
    AND rt.status_rekening=1
    AND %[2]s

-- name: GetRekeningTabCollateral-filter-GTJ
(upper(l.nomor_rekening) LIKE $1 or upper(rt.nama_rekening) LIKE $1)

-- name: GetRekeningCustomer-main
select nomor_rekening, jenis_rekening_customer, nomor_nasabah
from %[1]s.rekeningcustomer
where nomor_rekening = $1

-- name: CreateRekenigCustomer-main
insert into %[1]s.rekeningcustomer
(nomor_rekening, nomor_nasabah, jenis_rekening_customer, tgl_akad)
values
($1, $2, $3, $4)

-- name: GetRekeningGenerator-main
select 
counter
from %[1]s.rekeninggenerator
where grup = $1 and kode_generator = $2

-- name: InsertRekeningGenerator-main
insert into %[1]s.rekeninggenerator
(kode_generator ,counter ,grup)
values($1 ,1 ,$2)

-- name: UpdateCounterRekeningGenerator-main
update %[1]s.rekeninggenerator
set counter = counter + 1 
where grup = $1 and kode_generator = $2

-- name: GetRekeningNasabah-main
select rt.nomor_rekening, rl.nomor_nasabah, n.nama_nasabah 
from %[1]s.rekeningtransaksi rt
    inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rt.nomor_rekening
    inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah
where rt.nomor_rekening = $1

-- name: CreateRekenigLiabilitas-main
insert into %[1]s.rekeningliabilitas(
    nama_singkat, alamat_kirim, is_bagi_hasil_khusus, is_biaya_atm, is_biaya_rekening_dormant, is_biaya_saldo_minimum, is_blokir_debet, is_blokir_kredit, 
    is_blokir, is_boleh_debet, is_boleh_kredit, is_cetak_nota, is_dapat_bagi_hasil, is_join_account, is_kena_biayalayananumum, is_kena_pajak, 
    is_kena_zakat_bagi_hasil, is_rekening_baru, is_sedang_ditutup, is_status_passbook, is_tidak_dormant, is_tiering_nisbahbonus, 
    jenis_rekening_liabilitas, kode_marketing_current, kode_marketing_pertama, kode_marketing_referensi, kode_multicifrelation, kode_produk, 
    kode_program, kode_sumber_dana, kode_tujuan_rekening, nisbah_bagi_hasil, nisbah_dasar, nisbah_spesial, nomor_nasabah, nomor_rekening, 
    persentase_zakat_bagi_hasil, saldo_deposito, saldo_ditahan, saldo_float, status_kelengkapan, status_restriksi, kode_status_retriksi, tanggal_buka, tarif_pajak, 
    user_otorisasi, user_input, kode_member_bagihasil, kode_member_bagihasil_desc, kode_member_nominal, kode_member_nominal_desc, 
    baghas_alamat_transfer, nomor_rekening_disposisi, penerima_transfer_bagi_hasil, disposisi_bagi_hasil, baghas_sandikota_transfer, baghas_sandikota_transfer_desc,
    is_backdated, tanggal_bagi_hasil_berikutnya, jumlah_baghas
) 
values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, round($32, 2), round($33, 2), round($34, 2), $35, $36, round($37, 2), $38, $39, $40, $41, $42, $43, $44, 
    round($45, 2), $46, $47, $48, $49, $50 , $51, $52, $53, $54, $55, $56, $57, $58, $59, $60
)

-- name: UpdateRekenigLiabilitas-main
update %[1]s.rekeningliabilitas
set nisbah_spesial = round($1, 2), kode_program = $2, is_boleh_debet = $3, is_kena_pajak = $4, 
    kode_sumber_dana = $5, is_boleh_kredit = $6, is_biaya_atm = $7, is_blokir_debet = $8, 
    is_kena_biayalayananumum = $9, alamat_kirim = $10, nisbah_dasar = round($11, 2), 
    is_tidak_dormant = $12, kode_tujuan_rekening = $13, kode_multicifrelation = $14, 
    is_kena_zakat_bagi_hasil = $15, is_biaya_saldo_minimum = $16, status_kelengkapan = $17, 
    nisbah_bagi_hasil = round($18, 2), is_dapat_bagi_hasil = $19, is_blokir_kredit = $20, 
    is_join_account = $21, kode_marketing_current = $22, is_biaya_rekening_dormant = $23, 
    is_cetak_nota = $24, tarif_pajak = round($25, 2), persentase_zakat_bagi_hasil = round($26, 2), 
    kode_marketing_pertama = $27, is_tiering_nisbahbonus = $28, is_bagi_hasil_khusus = $29, 
    kode_marketing_referensi = $30, is_status_passbook = $31, user_update = $32, user_otorisasi = $33
where nomor_rekening = $34

-- name: GetRekeningLiab-main
select rl.nomor_nasabah, rl.nomor_rekening, rt.nama_rekening, rt.saldo, rt.kode_cabang, coalesce(rl.is_tutup_kartu_atm, 'F') as is_tutup_kartu_atm,
    to_char(rl.tanggal_buka, 'YYYY-MM-DD') as tanggal_buka, rl.nisbah_spesial, rl.nisbah_dasar, rl.tarif_pajak, rl.persentase_zakat_bagi_hasil, 
    rl.kode_sumber_dana, rl.kode_tujuan_rekening, rl.kode_marketing_current, rl.kode_marketing_pertama, rpb.nomorregisterbuku_aktif,
    case 
        when rl.jenis_rekening_liabilitas = 'T' then 'Tabungan'
        when rl.jenis_rekening_liabilitas = 'G' then 'Giro'
        when rl.jenis_rekening_liabilitas = 'R' then 'Rencana'
        when rl.jenis_rekening_liabilitas = 'D' then 'Deposito'
        when rl.jenis_rekening_liabilitas = 'P' then 'Perantara'
    end as jenis_rekening_liabilitas_desc,
    case 
        when n.jenis_nasabah = 'I' then i.alamat_rumah_jalan 
        when n.jenis_nasabah = 'K' then nk.alamat_jalan
        else ''
    end as alamat_jalan,
    case 
        when n.jenis_nasabah = 'I' then i.alamat_rumah_kode_pos 
        when n.jenis_nasabah = 'K' then nk.alamat_kode_pos
        else ''
    end as alamat_kode_pos,
    rl.saldo_float, rl.saldo_deposito, rl.saldo_ditahan, p.saldo_minimum, rl.saldo_tunggakan, rl.jenis_rekening_liabilitas, rl.kode_produk, 
    p.nama_produk as kode_produk_desc, rl.is_kena_biayalayananumum, rl.is_status_passbook, rl.alamat_kirim, rt.status_rekening, 
    case
        when rt.status_rekening = 1 then 'Aktif'
        when rt.status_rekening = 2 then 'Dormant'
        when rt.status_rekening = 3 then 'Tutup'
    end as status_rekening_desc,
    rt.user_input, rl.user_otorisasi, rt.keterangan, rt.kode_valuta, rt.nama_valuta, coalesce(rl.is_sedang_ditutup, 'F') is_sedang_ditutup, rl.kode_program, 
    rl.kode_marketing_referensi, rl.status_kelengkapan, rl.is_blokir, rl.is_blokir_debet, coalesce(rl.is_boleh_debet, 'T') as is_boleh_debet, 
    rl.is_blokir_kredit, coalesce(rl.is_boleh_kredit, 'T') as is_boleh_kredit, rl.is_cetak_nota, rl.is_dapat_bagi_hasil, rl.is_tidak_dormant, rl.is_biaya_rekening_dormant, rl.is_biaya_saldo_minimum, 
    rl.is_biaya_atm, rl.nisbah_bagi_hasil, rl.is_tiering_nisbahbonus, rl.is_join_account ,rl.kode_multicifrelation, rt.nomor_rekening,
    rt.saldo - coalesce(rl.saldo_ditahan, 0) - p.saldo_minimum - rl.saldo_float as saldo_efektif, rt.kode_jenis ,status_restriksi, to_char(rl.tanggal_tutup, 'YYYY-MM-DD') as tanggal_tutup
from %[1]s.rekeningliabilitas rl
    inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
    left join %[1]s.produk p on p.kode_produk = rl.kode_produk
    left join %[1]s.registerlayanan reglay on reglay.nomor_rekening_layanan = rt.nomor_rekening
    left join %[1]s.registerpassbook rpb on reglay.id_register = rpb.id_register
    inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah
    left join %[1]s.nasabahindividu ni on ni.nomor_nasabah = n.nomor_nasabah
    left join %[1]s.individu i on i.id_individu = ni.id_individu
    left join %[1]s.nasabahkorporat nk on nk.nomor_nasabah = n.nomor_nasabah 
where rl.nomor_rekening = $1

-- name: UpdateRekLiabBatalTTR-main
update %[1]s.rekeningliabilitas
set tanggal_tutup = $1,
    is_sedang_ditutup = $2
where nomor_rekening = $3

-- name: UpdateRekLiabTTR-main
update %[1]s.rekeningliabilitas
set tanggal_tutup = $1,
    is_sedang_ditutup = $2
where nomor_rekening = $3

-- name: UpdateRekLiabPencairanDeposito-main
update %[1]s.rekeningliabilitas
set tanggal_tutup = $1,
    saldo_accrual_bagihasil = 0
where nomor_rekening = $2

-- name: UpdateRekLiabPencairanDepositoRev-main
update %[1]s.rekeningliabilitas
set tanggal_tutup = null,
    saldo_accrual_bagihasil = $1
where nomor_rekening = $2

-- name: GetSaldoDitahan-main
select rl.saldo_ditahan
from %[1]s.rekeningliabilitas rl
where rl.nomor_rekening = $1

-- name: UpdateSaldoFloatRekenigLiabilitas-main
update %[1]s.rekeningliabilitas 
set saldo_float = (
    select sum(nominal) 
    from %[1]s.notadebitoutward
    where nomor_rekening = $1
        and status = 'S'
)
where nomor_rekening = $1

-- name: UpdateRekLiabilitasCollateral-main
UPDATE %[1]s.rekeningliabilitas
SET IS_BLOKIR = 'T', 
    KETERANGAN_BLOKIR = 'Blokir oleh user financing sebagai jaminan',
    SISTEM_PEMBLOKIR = 'FINANCING'
WHERE NOMOR_REKENING = $1

-- name: UpdateSaldoDitahanLiabiliitas-main
UPDATE %[1]s.rekeningliabilitas
SET SALDO_DITAHAN = SALDO_DITAHAN + $1
WHERE NOMOR_REKENING = $2

-- name: UpdateSaldoUnholdLiabitas-main
UPDATE %[1]s.rekeningliabilitas
SET SALDO_DITAHAN = SALDO_DITAHAN - $1
WHERE NOMOR_REKENING = $2

-- name: UnblokirRekLiabilitasFinFund-main
UPDATE %[1]s.rekeningliabilitas
SET IS_BLOKIR = 'F',
SISTEM_PEMBLOKIR = NULL
WHERE NOMOR_REKENING = $1

-- name: GetJangkaWaktuRencana-main
select 
    to_char(add_months(nilai_parameter_tanggal, $1), 'YYYY-MM-DD') as jw_date,
    to_char(add_months(nilai_parameter_tanggal, 1), 'YYYY-MM-DD') as next_month_date
from %[1]s.parameterglobal
where kode_parameter = 'POD'

-- name: CreateRekeningRencana-main
insert into %[1]s.rekeningrencana
(nomor_rekening ,nomor_rekening_induk ,nomor_rekening_pencairan ,setoran_rutin ,jangka_waktu ,tanggal_setoran_rutin 
    ,tanggal_jatuh_tempo ,id_register ,sudah_setoran_awal)
values 
($1, $2, $3, $4, $5, $6, $7, $8, $9)

-- name: UpdateRekenigRencana-main
update %[1]s.rekeningrencana 
set tanggal_setoran_rutin = $1, setoran_rutin = $2, nomor_rekening_pencairan = $3, 
    nomor_rekening_induk = $4
where nomor_rekening = $5

-- name: GetSIRencana-main
select rrt.nomor_rekening_kredit, rct.instructionid, rrt.id_register, to_char(rct.tgl_proses, 'yyyy-mm-dd') as tgl_proses, 
    to_char(rct.tgl_proses_berikutnya, 'yyyy-mm-dd') as tgl_proses_berikutnya, to_char(rct.tgl_kadaluarsa, 'yyyy-mm-dd') as tgl_kadaluarsa,
    rct.jumlah_sudah_proses
from %[1]s.registerrt rrt  
    inner join %[1]s.recurrenttransaction rct on rct.recid = rrt.recid
where rrt.nomor_rekening_kredit = $1

-- name: UpdateSIRencana-main
update %[1]s.standinginstruction 
set debitaccountno = $1, debitaccountname = $2, amount = $3
where instructionid = $4

-- name: UpdateRCTRencana-main
update %[1]s.recurrenttransaction 
set nominal = $1, tgl_proses = $2, tgl_kadaluarsa = $3
where instructionid = $4

-- name: UpdateRLRencana-main
update %[1]s.registerlayanan 
set nomor_rekening_layanan = $1
where id_register = $2

-- name: GetRekeningTransaksi-main
select rt.nama_rekening, rt.kode_valuta, rt.nama_valuta, rt.keterangan, rt.jenis_rekening_transaksi
from %[1]s.rekeningtransaksi rt
where rt.nomor_rekening = $1

-- name: CreateRekenigTransaksi-main
insert into %[1]s.rekeningtransaksi 
(nama_rekening, keterangan, kode_valuta, nama_valuta, nomor_rekening, kode_cabang, saldo, saldo_pod, jenis_rekening_transaksi, 
    status_rekening, tanggal_aktifitas_terakhir, balance_sign, user_input, kode_jenis, tanggal_input, tanggal_otorisasi) 
values 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, current_timestamp, current_timestamp)

-- name: UpdateRekeningTransaksi-main
update %[1]s.rekeningtransaksi 
set nama_rekening = $1, keterangan = $2, 
    tanggal_perbarui = current_timestamp, tanggal_update = $3, 
    user_perbarui = $4 
where nomor_rekening= $5

-- name: UpdateSaldoRekeningTransaksi-main
update %[1]s.rekeningtransaksi
set saldo = saldo + $1
where nomor_rekening = $2

-- name: UpdateStatusRekeningTransaksi-main
update %[1]s.rekeningtransaksi
set status_rekening = $1
where nomor_rekening = $2

-- name: GetDetailTrxByIdDetail-main
select t.keterangan, to_char(t.tanggal_transaksi, 'yyyy-mm-dd') as tanggal_transaksi, t.nomor_referensi
from %[1]v.transaksi t
    inner join %[1]v.detiltransaksi d on d.id_transaksi = t.id_transaksi
where id_detil_transaksi  = $1

-- name: GetLastId-main
select 
    case 
        when last_id is null then 0 
        else (nextval('%[1]v.seq_doicr_ref') - last_id + 1) 
    end as last_idgen 
from %[1]v.custom_idgen 
where idName = 'DOICREDITREF' 
    and idCode = $1

-- name: InsertIdGen-main
insert into %[1]v.custom_idgen 
(idname, idcode, last_id, locked)
values 
('DOICREDITREF', $1 , nextval('%[1]v.seq_doicr_ref'), 'F') 

-- name: InsertTmpRemCandidate-main
insert into %[1]v.remittance_candidate 
(record_id, status_history_id, id_detil_transaksi, branch_code, created_time, userid, ori_sandi_bank, ori_nomor_rekening, ori_nama_pengirim, 
    nomor_referensi, keterangan, nominal, recp_sandi_bank, recp_nomor_rekening, recp_nama_penerima, remittance_type, queue_acc_code_outgoing, 
    queue_branchcode_outgoing, id_transfer)
select nextval('%[2]v.seq_trx_record'), nextval('%[2]v.seq_trx_status_history'), dt.id_detil_transaksi, t.kode_cabang_transaksi, current_timestamp, t.user_input, $1, r.nomor_rekening_pengirim, substr(r.nama_pengirim, 1, 70), 
    $2, substr($3, 1, 140), dt.nilai_mutasi, r.kode_bank, r.nomor_rekening_penerima, substr(r.nama_penerima, 1, 70), 'SKN', r.kode_account_titipan, 
    r.kode_cabang_titipan, r.id_transfer
from %[3]v.transaksi t 
    inner join %[3]v.detiltransaksi dt on t.id_transaksi = dt.id_transaksi
    inner join %[3]v.transfer r on r.id_detil_transaksi = dt.id_detil_transaksi
    inner join %[4]v.cabang c on c.kode_cabang = t.kode_cabang_transaksi
where dt.id_detil_transaksi = $4

-- name: InsertTmpRemCandidateRtgs-main
insert into %[1]v.remittance_candidate 
(record_id, record_detail_id, status_history_id, id_detil_transaksi, branch_code, created_time, userid, ori_sandi_bank, ori_nomor_rekening, ori_nama_pengirim, 
    nomor_referensi, keterangan, nominal, recp_sandi_bank, recp_nomor_rekening, recp_nama_penerima, value_date, remittance_type, queue_acc_code_outgoing, 
    queue_branchcode_outgoing, id_transfer)
select nextval('%[2]v.seq_rttransaction'), nextval('%[2]v.seq_rttransferdetail'), nextval('%[2]v.seq_rtstatushistory'), dt.id_detil_transaksi, t.kode_cabang_transaksi, current_timestamp, t.user_input, c.sandi_bi, r.nomor_rekening_pengirim, substr(r.nama_pengirim,1,140), 
    $1, substr($2, 1, 96), dt.nilai_mutasi, r.kode_bank, r.nomor_rekening_penerima, substr(r.nama_penerima, 1, 140), t.tanggal_transaksi, 'RTGS', r.kode_account_titipan, 
    r.kode_cabang_titipan, r.id_transfer
from %[3]v.transaksi t 
    inner join %[3]v.detiltransaksi dt on t.id_transaksi = dt.id_transaksi
    inner join %[3]v.transfer r on r.id_detil_transaksi = dt.id_detil_transaksi
    inner join %[4]v.cabang c on c.kode_cabang = t.kode_cabang_transaksi
where dt.id_detil_transaksi = $3

-- name: InsertRemTrxStatusHist-main
insert into %[2]v.trx_status_history 
(trx_record_id, trx_status_history_id, created_date, created_time, created_by, trx_status_id)
select record_id, status_history_id, current_date, current_timestamp, userid, 9  
from %[1]v.remittance_candidate 
where id_detil_transaksi = $1

-- name: InsertRemTrxRecord-main
insert into %[2]v.trx_record 
(trx_record_id, trx_auth_status, trx_type,  branch_code, id_detil_transaksi, trx_status_history_id, queue_acc_code_outgoing, queue_branchcode_outgoing, status_generate_file)
select record_id, 'N', 'DOICR', branch_code, id_detil_transaksi, status_history_id, queue_acc_code_outgoing, queue_branchcode_outgoing , 'T'
from %[1]v.remittance_candidate 
where id_detil_transaksi = $1

-- name: InsertDoiCredit-main
insert into %[2]v.doi_credit 
(trx_record_id, reff_no, ori_kli_code, ori_opt_code, ori_kbi_code, sender_account, sender_name, recp_name, recp_account, recp_province_code, 
    recp_opt_code, bsn_type_code, notes, trx_amount, resident_status_code, citizen_status_code, record_type, dke_code, id_transfer, 
    sender_resident_status, sender_cust_type, sender_cust_address, sender_cust_identityno, sender_city_code, recp_resident_status, 
    recp_cust_type, recp_cust_identityno, recp_cust_address, recp_city_code, trx_code, recp_bank_code, reff_no_retur, channel_type)
select t.record_id, t.nomor_referensi, substr(t.ori_sandi_bank, 1, 8), null, null, t.ori_nomor_rekening, t.ori_nama_pengirim, t.recp_nama_penerima , t.recp_nomor_rekening,
    null, null, null, t.keterangan, t.nominal, '1', '0', '1', '1', t.id_transfer, tr.jenispenduduk_pengirim, tr.jenisnasabah_pengirim, tr.alamat_pengirim, tr.noidentitas_pengirim, 
    tr.sandikota_pengirim, tr.jenispenduduk_penerima, tr.jenisnasabah_penerima, tr.noidentitas_penerima, tr.alamat_penerima, tr.sandikota_penerima, tr.Kode_TX_SKN, tr.kode_Bank, 
    tr.Nomor_Referensi_Retur, tr.Jenis_Channel
from %[1]v.remittance_candidate  t
    inner join %[3]v.transfer tr on tr.id_transfer = t.id_transfer 
where t.id_detil_transaksi = $1

-- name: MergeDoiCredit-main
update %[2]v.doi_credit r
set 
    sender_resident_status = t.jenispenduduk_pengirim
    , sender_cust_type = t.jenisnasabah_pengirim
    , sender_cust_address = t.alamat_pengirim
    , sender_cust_identityno = t.noidentitas_pengirim
    , sender_city_code = t.sandikota_pengirim
    , recp_resident_status = t.jenispenduduk_penerima
    , recp_cust_type = t.jenisnasabah_penerima
    , recp_cust_identityno = t.noidentitas_penerima
    , recp_cust_address = t.alamat_penerima
    , recp_city_code = t.sandikota_penerima
    , trx_code = t.Kode_TX_SKN
    , recp_bank_code = t.kode_Bank
    --, recp_bankcode_operator = t.kode_Bank
    , reff_no_retur = t.Nomor_Referensi_Retur
    , channel_type = t.Jenis_Channel
from  %[3]v.transfer t
where t.id_transfer = r.id_transfer
 and exists ( 
    select 1 from %[1]v.remittance_candidate 
    where trx_record_id = r.trx_record_id
        and id_detil_transaksi = $1
)

-- name: InsertRemTransSkn-main
insert into %[2]v.remittancetransaction
(id_detil_transaksi, trx_record_id, jenis_transfer, is_uploaded)
select id_detil_transaksi, record_id, 'S', 'F'
from %[1]v.remittance_candidate
where id_detil_transaksi = $1  

-- name: DeleteRemTrans-main
delete from %[1]v.remittance_candidate
where id_detil_transaksi = $1   

-- name: CreateRemittanceCandidate-main
insert into %[2]v.remittance_candidate
(record_id, record_detail_id, status_history_id, id_detil_transaksi, branch_code, created_time, userid, ori_sandi_bank, ori_nomor_rekening, ori_nama_pengirim, 
    nomor_referensi, keterangan, nominal, recp_sandi_bank, recp_nomor_rekening, recp_nama_penerima, value_date, remittance_type, queue_acc_code_outgoing, 
    queue_branchcode_outgoing, id_transfer)
select
    nextval('%[1]v.seq_trxrecord'), nextval('%[1]v.seq_rttransferdetail'), nextval('%[1]v.seq_rtstatushistory'), dt.id_detil_transaksi, t.kode_cabang_transaksi, current_timestamp, t.user_input, $1, r.nomor_rekening_pengirim, substr(r.nama_pengirim, 1, 140),
    $2, substr($3, 1, 140), dt.nilai_mutasi, r.kode_bank, r.nomor_rekening_penerima, substr(r.nama_penerima,1,140), t.tanggal_transaksi, 'RTGS', r.kode_account_titipan,
    r.kode_cabang_titipan, r.id_transfer
from %[3]v.transaksi t 
    inner join %[3]v.detiltransaksi dt on t.id_transaksi = dt.id_transaksi
    inner join %[3]v.transfer r on r.id_detil_transaksi = dt.id_detil_transaksi
    inner join %[4]v.cabang c on c.kode_cabang = t.kode_cabang_transaksi
where dt.id_detil_transaksi = $4	

-- name: CreateRTStatusHistory-main
insert into %[1]v.rtstatushistory
(history_id, status, trans_id, id_user, created_date, created_time)
select status_history_id, 'N' , record_id, userid, current_date, created_time
from %[2]v.remittance_candidate
where id_detil_transaksi = $1	

-- name: CreateRTTransaction-main
insert into %[1]v.rttransaction r 
(trans_id, from_member, status, record_type, trn, transaction_code, related_trn, to_member, amount, amount_str, value_date, branch_code, trans_type, transdetail_id, 
    queue_acc_code_outgoing, queue_branchcode_outgoing, id_transfer, processed_date)
select record_id, $1, 'N', 1, $2, '600', $3, recp_sandi_bank, nominal, lpad(to_char(round(nominal, 2) * 100), 1, '0'), TO_CHAR(value_date, 'YYYYMMDD'), branch_code, 'I', record_detail_id, 
    queue_acc_code_outgoing, queue_branchcode_outgoing, id_transfer, value_date
from %[2]v.remittance_candidate rc
where id_detil_transaksi = $4

-- name: MergeRTTransaction-main
update %[1]v.rttransaction s
set trn = q.kode_tx_rtgs, priority_code = q.kode_prioritas_rtgs, message_type = q.tipe_message_rtgs
from (
    select tmp.record_id, r.*
    from %[2]v.remittance_candidate tmp
        inner join %[3]v.transfer r on r.id_transfer = tmp.id_transfer
    where tmp.id_detil_transaksi = $1
) q 
where q.record_id = s.trans_id
  and exists (
    select 1 from %[2]v.remittance_candidate rc 
    where record_id = s.trans_id
    and id_detil_transaksi = $1)

-- name: CreateRTTransferDetail-main
insert into %[1]v.rttransferdetail r
(transdetail_id ,transdetail_type)
select record_detail_id, 'S'
from %[2]v.remittance_candidate rc
where id_detil_transaksi = $1

-- name: CreateRTSingleTransfer-main
insert into %[1]v.rtsingletransfer r
(transdetail_id, member_info, from_account_number, from_account_name, to_account_name, sender_account, originating_name, ultimate_benef_account , ultimate_benef_name, payment_details)
select t.record_detail_id, $1, $2, $3, recp.Nama_Member, t.ori_nomor_rekening , t.ori_nama_pengirim, t.recp_nomor_rekening , t.recp_nama_penerima, t.keterangan
from %[2]v.remittance_candidate t
    inner join %[3]v.memberrtgs recp on t.recp_sandi_bank = recp.kode_member
where t.id_detil_transaksi = $4

-- name: MergeRTSingleTransfer-main
update %[1]v.rtsingletransfer s
set resident_status = (
        case when q.jenis_penduduk = 'B' then 'NR'
        else 'R'
    end),
    citizen_status = q.status_pengirim,
    to_account_number = q.nomor_account_bank,
    bi_benef_account = q.no_rek_intern_bi,
    bi_benef_name = q.nama_rek_intern_bi,
    originating_address = SUBSTR(q.alamat_pengirim, 1, 69),
    originating_city_code = q.sandikota_pengirim,
    ultimate_benef_address = SUBSTR(q.alamat_penerima, 1, 69),
    ultimate_benef_city_code = q.sandikota_penerima,
    pvp_cp_sender_correspondent = q.pvp_cp_sender_correspondent,
    pvp_cp_receiver_correspondent = q.pvp_cp_receiver_correspondent,
    pvp_ordering_institution = q.pvp_ordering_institution,
    pvp_beneficiary_institution = q.pvp_beneficiary_institution
from (
    select tmp.record_detail_id, r.*
    from %[2]v.remittance_candidate tmp
        inner join %[3]v.transfer r on r.id_transfer = tmp.id_transfer
) q 
where q.record_detail_id = s.transdetail_id
 and exists (
    select
        1
    from
        %[2]v.remittance_candidate rc 
    where
        record_detail_id = s.transdetail_id
        AND id_detil_transaksi = $1
)

-- name: CreateRemittanceTransaction-main
insert into %[1]v.remittancetransaction r
(id_detil_transaksi, trx_record_id, jenis_transfer , is_uploaded)
select id_detil_transaksi, record_id, 'R', 'F'
from %[2]v.remittance_candidate rc
where
    id_detil_transaksi = $1

-- name: GetReportCodeList-main
select kode_report, nama_report 
from %[1]s.report
where is_show_bds = 'T'
%[2]s
order by kode_report

-- name: GetReportCodeList-filter-Keyword
and lower(nama_report) like lower($1) or lower(kode_report) like lower($1)

-- name: GetDataReportNasabah-main
select n.nomor_nasabah, 
    n.nama_nasabah, 
    n.jenis_nasabah,
    n.user_input,
    n.user_otorisasi,
    p1.rule_flag as flag_pep_ni,
    p2.rule_flag as flag_nrt_ni,
    p3.rule_flag as flag_nrt_nk,
    n.tanggal_buka
from %[1]s.nasabah n
    left join %[1]s.nasabahindividu ni on (n.nomor_nasabah = ni.nomor_nasabah)
    left join %[1]s.paramrulereferencedata p1 on (p1.refdata_id = ni.id_pekerjaan and p1.rule_code = 'pep_ni' )
    left join %[1]s.paramrulereferencedata p2 on (p2.refdata_id = ni.id_pekerjaan and p2.rule_code = 'nrt_ni' )
    left join %[1]s.nasabahkorporat nk on (n.nomor_nasabah = nk.nomor_nasabah)
    left join %[1]s.paramrulereferencedata p3 on (p3.refdata_id = nk.id_sektor_ekonomi and p3.rule_code = 'nrt_nk' )
where
    n.kode_cabang_input = $1
    and n.tanggal_buka between to_date($2, 'yyyy-mm-dd')  and  to_date($3, 'yyyy-mm-dd')
    order by n.tanggal_buka

-- name: GetDataReportRekeningBaru-main
select
    rt.nomor_rekening, 
    rt.nama_rekening,
    rl.nomor_nasabah,
    rl.kode_program,
    coalesce(rl.is_join_account,'F') as is_join_account,
    to_char(rl.tanggal_buka, 'DD-MM-YYYY') tanggal_buka,
    n.nama_nasabah, 
    p.kode_produk,
    p.nama_produk,
    p.saldo_minimum,
    rt.user_input,
    rl.user_otorisasi,
    1 as data_count
from
    %[1]s.rekeningtransaksi rt
    inner join %[1]s.rekeningliabilitas rl on rt.nomor_rekening = rl.nomor_rekening
    inner join %[1]s.nasabah n on rl.nomor_nasabah = n.nomor_nasabah
    inner join %[1]s.produk p on rl.kode_produk = p.kode_produk
where
    rt.kode_jenis in ('SAV', 'CA')
    and (rt.kode_cabang = $1 or 'ALL' = $2)
    and rl.tanggal_buka between to_timestamp($3, 'yyyy-mm-dd') and to_timestamp($4, 'yyyy-mm-dd')
order by rt.kode_cabang, rl.tanggal_buka, rl.kode_produk, rt.nomor_rekening

-- name: GetDataReportNasabahBaru-main
select
    n.Nomor_Nasabah,
    n.Nama_Nasabah,
    case when n.JENIS_NASABAH = 'I' then 'Individu' when n.JENIS_NASABAH = 'K' then 'Korporat' else '-' end JENIS_NASABAH,
    n.User_Input,
    upper(coalesce(o.User_Otorisasi, n.user_otorisasi)) as user_otorisasi,
    case when n.STATUS_RESIKO = 'H' then 'High' when n.STATUS_RESIKO = 'M' then 'Medium' when n.STATUS_RESIKO = 'L' then 'Low' else '-' end STATUS_RESIKO,
    case when p1.rule_flag = 'T' then 'Tidak' when p1.rule_flag = 'Y' then 'Ya' else '-' end FLAG_PEP_NI,
    case when coalesce(p2.rule_flag, p3.rule_flag) = 'T' then 'Tidak' when p1.rule_flag = 'Y' then 'Ya' else '-' end FLAG_NRT,
    to_char(n.tanggal_buka, 'DD-MM-YYYY') tanggal_buka
from
    %[1]s.nasabah n
    left join %[1]s.nasabahindividu ni on (n.Nomor_Nasabah = ni.Nomor_Nasabah)
    left join %[1]s.paramrulereferencedata p1 on (p1.REFDATA_ID = ni.Id_Pekerjaan and p1.rule_code = 'PEP_NI' )
    left join %[1]s.paramrulereferencedata p2 on (p2.REFDATA_ID = ni.Id_Pekerjaan and p2.rule_code = 'NRT_NI' )
    left join %[1]s.nasabahkorporat nk on (n.Nomor_Nasabah = nk.Nomor_Nasabah)
    left join %[1]s.paramrulereferencedata p3 on (p3.REFDATA_ID = nk.Id_Sektor_Ekonomi and p3.rule_code = 'NRT_NK' )
    left join %[2]s.otorentrihistori o on o.key_string = n.nomor_nasabah and o.kode_entri = 'CRCS001' and o.status_otorisasi = 'T'
where
    ( n.Kode_Cabang_Input = $1 or 'ALL' = $2 )
    and n.TANGGAL_BUKA >= TO_DATE($3, 'YYYY-MM-DD')
    and n.TANGGAL_BUKA < TO_DATE($4, 'YYYY-MM-DD')
order by
    n.TANGGAL_BUKA

-- name: GetDataReportSetoranKliring-main
select
    ND.Nomor_Aplikasi,
    DT.Nilai_Mutasi,
    ND.Nomor_Referensi,
    (case ND.Jenis_Warkat
        when 'C' then '00'
        when 'B' then '10'
        else '40'
    end) as Kode_Warkat,
    ND.Nomor_Rekening_Nota,
    ND.Kode_Bank,
    ND.Nomor_Rekening,
    RT.Nama_Rekening,
    coalesce(qb.Nama_Singkat, m.nama_member) as Nama_Bank,
    T.User_Input,
    T.User_Otorisasi,
    DT.Id_Detil_Transaksi
from
    %[1]s.transaksi T,
    %[1]s.detiltransaksi DT,
    %[1]s.dtkliringoutward DTO,
    %[1]s.notadebitoutward ND
    left join (
        select
            bk.kode_bank,
            k.nama_singkat
        from
            %[1]s.bankkliring bk,
            %[1]s.kliringbank k
        where
            bk.sandi_bi = k.sandi_bank) qb on (qb.kode_bank = ND.kode_bank)
    left join %[1]s.memberrtgs m on (m.kode_member = ND.kode_bank),
    %[1]s.rekeningtransaksi RT
where
    T.id_transaksi = DT.id_transaksi
    and RT.Nomor_Rekening = ND.Nomor_Rekening
    and DTO.id_detil_transaksi = DT.id_detil_transaksi
    and ND.Id_NotaDebitOutward = DTO.Id_NotaDebitOutward
    and ( T.kode_cabang_transaksi = $1 or 'ALL' = $2)
    and T.Tanggal_Transaksi >= TO_DATE($3, 'YYYY-MM-DD')
    and T.Tanggal_Transaksi < TO_DATE($4, 'YYYY-MM-DD')
    and T.kode_transaksi = 'SDN'
    
union all

select
    ND.Nomor_Aplikasi,
    DT.Nilai_Mutasi,
    ND.Nomor_Referensi,
    (case ND.Jenis_Warkat
        when 'C' then '00'
        when 'B' then '10'
        else '40'
    end) as Kode_Warkat,
    ND.Nomor_Rekening_Nota,
    ND.Kode_Bank,
    ND.Nomor_Rekening,
    RT.Nama_Rekening,
    coalesce(qb.Nama_Singkat, m.nama_member) as Nama_Bank,
    T.User_Input,
    T.User_Otorisasi,
    DT.Id_Detil_Transaksi
from
    %[1]s.histtransaksi T,
    %[1]s.histdetiltransaksi DT,
    %[1]s.histdtkliringoutward DTO,
    %[1]s.notadebitoutward ND
    left join (
        select
            bk.kode_bank,
            k.nama_singkat
        from
            %[1]s.bankkliring bk,
            %[1]s.kliringbank k
        where
            bk.sandi_bi = k.sandi_bank) qb on (qb.kode_bank = ND.kode_bank)
    left join %[1]s.memberrtgs m on (m.kode_member = ND.kode_bank),
    %[1]s.rekeningtransaksi RT
where
    T.id_transaksi = DT.id_transaksi
    and RT.Nomor_Rekening = ND.Nomor_Rekening
    and DTO.id_detil_transaksi = DT.id_detil_transaksi
    and ND.Id_NotaDebitOutward = DTO.Id_NotaDebitOutward
    and ( T.kode_cabang_transaksi = $1 or 'ALL' = $2)
    and T.Tanggal_Transaksi >= TO_DATE($3, 'YYYY-MM-DD')
    and T.Tanggal_Transaksi < TO_DATE($4, 'YYYY-MM-DD')
    and T.kode_transaksi = 'SDN'
order by Nomor_Referensi, Id_Detil_Transaksi

-- name: GetDataReportRekeningTutup-main
SELECT 
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    RT.Kode_Valuta,
    N.Nama_Nasabah,
    P.Kode_Produk,
    P.Nama_Produk,
    P.Saldo_Minimum,
    I.UserInput,
    I.UserOtorisasi,
    to_char(I.TANGGAL_PROSES, 'DD-MM-YYYY') tanggal_proses,
    RT.Kode_Cabang,
    1 AS data_count
FROM
    %[1]s.rekeningtransaksi RT,
    %[1]s.rekeningliabilitas RL,
    %[1]s.nasabah N,           
    %[1]s.produk P,
    %[1]s.infopenutupanrekening I
WHERE
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Nomor_Nasabah = N.Nomor_Nasabah
    and RL.Kode_Produk = P.Kode_Produk
    and I.Nomor_Rekening = RL.Nomor_Rekening
    and ( RT.Kode_Cabang = $1 or 'ALL' = $2 )
    and I.TANGGAL_PROSES >= TO_DATE($3, 'YYYY-MM-DD')
    and I.TANGGAL_PROSES < TO_DATE($4, 'YYYY-MM-DD')
ORDER By
    I.TANGGAL_PROSES, RT.Kode_Jenis, RL.Kode_Produk, I.UserInput, RT.Nomor_Rekening

-- name: GetDataReportRekeningBlokir-main
select
    RT.kode_cabang,
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    RT.Kode_Valuta,
    N.Nama_Nasabah,
    P.Kode_Produk,
    P.Nama_Produk,
    HB.User_Input,
    HB.User_DualControl,
    HB.Keterangan_Blokir,
    HB.Sistem_Pemblokir,
    to_char(HB.Tanggal_Ubah, 'DD-MM-YYYY') tanggal_ubah
from
    %[1]s.rekeningtransaksi RT,
    %[1]s.rekeningliabilitas RL,
    %[1]s.histblokir HB,
    %[1]s.nasabah N,
    %[1]s.produk P
where
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Nomor_Nasabah = N.Nomor_Nasabah
    and HB.Nomor_Rekening = RT.Nomor_Rekening
    and RL.Kode_Produk = P.Kode_Produk
    and RT.Kode_Jenis in ('SAV', 'CA')
    and HB.Status_Blokir = 'T'
    and ( RT.Kode_Cabang = $1 or 'ALL' = $2 )
    and HB.Tanggal_Ubah >= TO_DATE($3, 'YYYY-MM-DD')
    and HB.Tanggal_Ubah < TO_DATE($4, 'YYYY-MM-DD')
order by
    RT.Kode_Jenis,
    RL.Kode_Produk,
    RT.Nama_Rekening

-- name: GetDataReportRekeningHoldDana-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    RT.Kode_Valuta,
    N.Nama_Nasabah,
    P.Kode_Produk,
    P.Nama_Produk,
    HD.User_Hold,
    to_char(HD.Tanggal_Kadaluarsa, 'DD-MM-YYYY') Tanggal_Kadaluarsa,
    HD.Nominal_Hold,
    HD.Alasan_Hold
from
    %[1]s.rekeningtransaksi RT,
    %[1]s.rekeningliabilitas RL,
    %[1]s.nasabah N,
    %[1]s.holddanarekening HD,
    %[1]s.produk P
where
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Nomor_Nasabah = N.Nomor_Nasabah
    and HD.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Kode_Produk = P.Kode_Produk
    and RT.Kode_Cabang = $1
    and HD.Tanggal_Hold >= TO_DATE($2, 'YYYY-MM-DD')
    and HD.Tanggal_Hold < TO_DATE($3, 'YYYY-MM-DD')
order by
    RT.Kode_Cabang,
    RT.Kode_Jenis,
    RL.Kode_Produk,
    RT.Nama_Rekening

-- name: GetDataReportRegistrasiAutosave-main
select
    to_char(RL.TANGGAL_REGISTRASI, 'DD-MM-YYYY') TANGGAL_REGISTRASI,
    RT.Nomor_Rekening as nomor_rekening_asal,
    RT.Nama_Rekening as nama_rekening_asal,
    RL.User_Pembuat,
    RA.Nomor_Rekening_Tujuan,
    RT2.Nama_Rekening as nama_rekening_tujuan,
    SI.Description
from
    %[1]s.rekeningtransaksi RT,
    %[1]s.registerlayanan RL,
    %[1]s.registerautosweep RA,
    %[1]s.standinginstruction SI,
    %[1]s.rekeningtransaksi RT2
where
    RT.Nomor_Rekening = RL.Nomor_Rekening_Layanan
    and RL.Id_Register = RA.Id_Register
    and RL.Jenis_Register_Layanan = 'AE'
    and RT2.Nomor_Rekening = RA.Nomor_Rekening_Tujuan
    and RA.instructionid = SI.instructionid
    and (RL.Kode_Cabang_Input = $1 or 'ALL' = $2 )
    and RL.tanggal_registrasi >= TO_DATE($3, 'YYYY-MM-DD')
    and RL.tanggal_registrasi < TO_DATE($4, 'YYYY-MM-DD')
order by
    RT.Nomor_Rekening

-- name: GetDataReportNonAktifAutosave-main
SELECT 
    to_char(RL.Tanggal_NonAktif, 'DD-MM-YYYY') Tanggal_NonAktif,
    RT.Nomor_Rekening as nomor_rekening_asal, 
    RT.Nama_Rekening as nama_rekening_asal, 
    RL.User_Input_NonAktif, 
    RL.User_Otorisasi_NonAktif, 
    RA.Nomor_Rekening_Tujuan, 
    RT2.Nama_Rekening as nama_rekening_tujuan 
FROM
    %[1]s.rekeningtransaksi RT,
    %[1]s.registerlayanan RL,
    %[1]s.registerautosweep RA,
    %[1]s.standinginstruction SI,
    %[1]s.rekeningtransaksi RT2      
WHERE
    RT.Nomor_Rekening = RL.Nomor_Rekening_Layanan
    AND RL.Id_Register = RA.Id_Register
    AND RL.Jenis_Register_Layanan = 'AE'
    AND RT2.Nomor_Rekening = RA.Nomor_Rekening_Tujuan
    AND RA.instructionid = SI.instructionid 
    AND (RL.Kode_Cabang_NonAktif = $1 OR 'ALL' = $2 )
    AND RL.Tanggal_NonAktif >= TO_DATE($3, 'YYYY-MM-DD')
    AND RL.Tanggal_NonAktif < TO_DATE($4, 'YYYY-MM-DD')	
ORDER By
    RT.Nomor_Rekening 

-- name: GetDataReportAlamatAlternatifNasabah-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    N.Nama_Nasabah,
    P.Kode_Produk,
    P.Nama_Produk,
    AA.alamat_jalan || ' - ' || AA.alamat_rtrw || ' - ' || AA.alamat_kelurahan || ' - ' || AA.alamat_kecamatan
        || ' - ' || AA.alamat_kota_kabupaten || ' - ' || AA.alamat_provinsi || ' - ' || AA.alamat_kode_pos as Alamat_Alternatif,
    AA.Alamat_Email,
    AA.Telepon
from
    %[1]s.rekeningtransaksi RT,
    %[1]s.rekeningliabilitas RL,
    %[1]s.alamatalternatif AA,
    %[1]s.nasabah N,
    %[1]s.produk P
where
    AA.Nomor_Rekening = RT.Nomor_Rekening
    and RT.Nomor_Rekening = RL.Nomor_Rekening
    and RT.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Nomor_Nasabah = N.Nomor_Nasabah
    and RL.Kode_Produk = P.Kode_Produk
    and RT.Kode_Jenis in ('SAV', 'CA')
    and RT.Kode_Cabang = $1
order by
    RT.Nama_Rekening

-- name: GetDataReportRekeningUbahProduk-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    to_char(REP.Tanggal, 'DD-MM-YYYY') Tanggal,
    REP.Cabang_Input,
    REP.Saldo,
    REP.User_Input,
    REP.User_Otorisasi,
    REP.Nomor_Jurnal,
    REP.Id_Transaksi,
    REP.Cabang_Rekening || ' - ' || RT.Nama_Rekening AS CABANG_REKENING,
    REP.Kode_Produk_Lama || ' - ' || P1.Nama_Produk AS PRODUK_LAMA,
    REP.Kode_Produk_Baru || ' - ' || P2.Nama_Produk AS PRODUK_BARU,
    CAB.Nama_Cabang as Nama_Cabang,
    1 as count
from
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,
    %[1]v.produk P1,
    %[1]v.produk P2,
    %[2]v.cabang CAB,
    %[1]v.infoubahprodukrekening REP
where
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and REP.Kode_Produk_Lama = P1.Kode_Produk
    and REP.Kode_Produk_Baru = P2.Kode_Produk
    and REP.Nomor_Rekening = RT.Nomor_Rekening
    and REP.Cabang_Rekening = CAB.Kode_Cabang
    and REP.Tanggal >= TO_DATE($1, 'YYYY-MM-DD')
    and REP.Tanggal < TO_DATE($2, 'YYYY-MM-DD')
    and (REP.Cabang_Rekening = $3 or 'ALL' = $4)
order by
    REP.Tanggal,
    REP.Cabang_Input,
    REP.Kode_Produk_Lama,
    RT.Nomor_Rekening

-- name: GetDataReportTabunganBagihasilMudharabah-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    RL.Nisbah_Bagi_Hasil,
    DTB.Saldo,
    DTB.Nisbah,
    (DTB.GDR * 1.2) as GDR,
    (DTB.GDR * DTB.NISBAH) as eqv_rate,
    (DTB.Nominal - DTB.Pajak) as pajak_after_baghas,
    (DTB.Nominal - DTB.Zakat) as zakat_after_baghas,
    DTB.Nominal as Nominal,
    DTB.Pajak,
    DTB.Zakat,
    --(case DTB.Disposisi
    --   when 'H' then 'Dihibahkan'      
    --   when 'K' then 'Dikapitalisir'       
    --   when 'P' then 'Pindah Buku'
    --   when 'R' then 'Transfer' 
    --   when 'G' then 'Transfer'
    --   when 'T' then 'Tunai' 
    --   else 'unknown'
    -- end) as Disposisi_bagi_hasil,
    RL.rekening_disposisi_bagi_hasil,
    RL.kode_produk,
    P.Nama_Produk,
    rank() over (
    order by RT.Kode_Cabang,
    RL.Kode_Produk,
    RT.Nomor_Rekening) as rownumber
from
    %[1]s.rekeningtransaksi RT
    inner join %[1]s.rekeningliabilitas RL on RT.nomor_rekening = RL.nomor_rekening
    inner join %[1]s.produk P on RL.kode_produk = P.kode_produk
    inner join %[2]s.bagihasil_tabgir DTB on DTB.nomor_rekening = RT.nomor_rekening
where
    P.jenis_akad in ('M')
    and (RT.Kode_Cabang = $1 or 'ALL' = $2)
    and (DTB.Tanggal >= TO_DATE($3, 'YYYY-MM-DD') or 'ALL' = $4)
    and (DTB.Tanggal < TO_DATE($5, 'YYYY-MM-DD') or 'ALL' = $6)
order by
    RT.Kode_Cabang,
    RL.Kode_Produk,
    RT.Nomor_Rekening

-- name: GetDataReportTabunganBonusWadiah-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Saldo_Deposito,
    RL.Nomor_Nasabah,
    RL.Nisbah_Bagi_Hasil,
    DTB.Nominal as Nominal,
    DTB.Saldo,
    DTB.Nisbah,
    DTB.Pajak,
    DTB.Zakat,
    --(case DTB.Disposisi
    --   when 'H' then 'Dihibahkan'      
    --   when 'K' then 'Dikapitalisir'       
    --   when 'P' then 'Pindah Buku'
    --   when 'R' then 'Transfer' 
    --   when 'G' then 'Transfer'
    --   when 'T' then 'Tunai' 
    --   else 'unknown'
    -- end) as Disposisi_bagi_hasil,
    RL.rekening_disposisi_bagi_hasil ,
    rank() over (
    order by RT.Kode_Cabang,
    RL.Kode_Produk,
    RT.Nomor_Rekening) as rownumber
from
    %[1]s.rekeningtransaksi RT
    inner join %[1]s.rekeningliabilitas RL on RT.nomor_rekening = RL.nomor_rekening
    inner join %[1]s.produk P on RL.kode_produk = P.kode_produk
    inner join %[2]s.bagihasil_tabgir DTB on DTB.nomor_rekening = RT.nomor_rekening
where
    P.jenis_akad in ('W')
    and (RT.Kode_Cabang = $1 or 'ALL' = $2)
    and (DTB.Tanggal >= TO_DATE($3, 'YYYY-MM-DD') or 'ALL' = $4)
    and (DTB.Tanggal < TO_DATE($5, 'YYYY-MM-DD') or 'ALL' = $6)
order by
    RT.Kode_Cabang,
    RL.Kode_Produk,
    RT.Nomor_Rekening

-- name: GetDataReportBiayaAdministrasi-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    DTB.Nominal as Nominal
from
    %[1]s.rekeningtransaksi RT,
    %[2]s.biaya_adm DTB
where
    RT.Nomor_Rekening = DTB.Nomor_Rekening
    and (rt.kode_cabang = $1 or 'ALL' = $2)
    and (DTB.tanggal >= TO_DATE($3, 'YYYY-MM-DD') or 'ALL' = $4)
    and (DTB.tanggal < TO_DATE($5, 'YYYY-MM-DD') or 'ALL' = $6)
    and tag_biaya = 'ADM'
order by
    RT.Kode_Cabang,
    RT.Nomor_Rekening

-- name: GetDataReportRekeningSaldoHarian-main
select
    RT.Nomor_Rekening,
    case when RL.STATUS_RESTRIKSI = 'F' then RT.NAMA_REKENING else '**********' end NAMA_REKENING,
    case when RL.IS_BLOKIR = 'T' then 1 else 0 end IS_BLOKIR,
    case when RL.IS_BOLEH_DEBET = 'T' then 0 else 1 end IS_BOLEH_DEBET,
    case when RL.IS_BOLEH_KREDIT = 'T' then 0 else 1 end IS_BOLEH_KREDIT,
    case when RL.STATUS_RESTRIKSI = 'F' then to_char(RT.SALDO ) else '**********' end SALDO,
    case when RL.STATUS_RESTRIKSI = 'F' then to_char(P.SALDO_MINIMUM) else '**********' end SALDO_MINIMUM,
    case when RL.STATUS_RESTRIKSI = 'F' then to_char(RL.SALDO_FLOAT ) else '**********' end SALDO_FLOAT,
    case when RL.STATUS_RESTRIKSI = 'F' then to_char(RL.SALDO_DITAHAN)  else '**********' end SALDO_DITAHAN,
    case when RL.STATUS_RESTRIKSI = 'F' then to_char((case
        when (RT.Saldo - (P.Saldo_Minimum + RL.Saldo_Float + RL.Saldo_Ditahan)) < 0.0 then 0.0
        else (RT.Saldo - (P.Saldo_Minimum + RL.Saldo_Float + RL.Saldo_Ditahan))
    end)) else '**********' end SALDO_EFEKTIF,
    RL.Nomor_Nasabah,
    RL.Status_Restriksi,
    RL.Kode_Program,
    N.Nama_Nasabah,
    P.Kode_Produk,
    P.Nama_Produk,
    1 as data_count,
    rank() over (order by RT.Kode_Jenis, RL.Kode_Produk, RL.User_Input, RT.Nomor_Rekening ) as rownumber
from
    %[1]s.rekeningtransaksi RT
    left outer join %[1]s.accountavgbalance AVG on AVG.accountno = RT.Nomor_Rekening
    inner join %[1]s.rekeningcustomer RC on RT.Nomor_Rekening = RC.Nomor_Rekening
    inner join %[1]s.rekeningliabilitas RL on RC.Nomor_Rekening = RL.Nomor_Rekening
    inner join %[1]s.nasabah N on RL.Nomor_Nasabah = N.Nomor_Nasabah
    inner join %[1]s.produk P on RL.Kode_Produk = P.Kode_Produk
where
    RT.Status_Rekening <> 3
    and RT.Kode_Jenis in ('SAV', 'CA')
    and (RT.Kode_Cabang = $1 or 'ALL' = $2)
order by
    RT.Kode_Jenis,
    RL.Kode_Produk,
    RL.User_Input,
    RT.Nomor_Rekening

-- name: GetDataReportRekeningReposisi-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    RL.Kode_Produk || ' - ' || P.Nama_Produk AS PRODUK,
    CAB1.Kode_Cabang || ' - ' || CAB1.Nama_Cabang AS CABANG_ASAL,
    CAB2.Kode_Cabang || ' - ' || CAB2.Nama_Cabang AS CABANG_TUJUAN,
    REP.Cabang_Input,
    REP.Saldo_Reposisi,
    REP.User_Input,
    REP.User_Otorisasi
from
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,
    %[1]v.produk P,
    %[2]v.cabang CAB1,
    %[2]v.cabang CAB2,
    %[1]v.inforeposisirekening REP
where
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Kode_Produk = P.Kode_Produk
    and RT.Kode_Jenis in ('SAV', 'CA')
    and REP.Nomor_Rekening = RT.Nomor_Rekening
    and REP.Cabang_Asal = CAB1.Kode_Cabang
    and REP.Cabang_Tujuan = CAB2.Kode_Cabang
    and REP.Cabang_Input = $1
    and REP.Tanggal = TO_DATE($2, 'YYYY-MM-DD')
order by
    REP.Cabang_Input,
    RL.Kode_Produk,
    RT.Nomor_Rekening

-- name: GetDataReportRekeningSaldoRekap-main
select
    RT.Kode_Cabang,
    P.Kode_Produk,
    P.Nama_Produk,
    round(sum(%[2]v), 2) as Saldo,
    sum(1) as jumlah_total ,
    sum(case when status_rekening <> 3 then 1 else 0 end) as jumlah_aktif,
    sum(case when status_rekening = 3 then 1 else 0 end) as jumlah_tutup
from
    %[1]v.rekeningtransaksi RT
    left outer join getdailybalancerek_bydate(to_date($1, 'yyyy-mm-dd')) db 
        on (db.nomor_rekening = rt.nomor_rekening and db.balance_field = 'saldo')
    inner join %[1]v.rekeningcustomer RC on RT.Nomor_Rekening = RC.Nomor_Rekening
    inner join %[1]v.rekeningliabilitas RL on RC.Nomor_Rekening = RL.Nomor_Rekening
    inner join %[1]v.produk P on RL.Kode_Produk = P.Kode_Produk
where
    RT.Kode_Jenis in ('SAV', 'CA')
    and ( RT.Kode_Cabang = $2
        or 'ALL' = $3)
group by
    RT.Kode_Cabang,
    P.Kode_Produk,
    P.Nama_Produk
order by
    RT.Kode_Cabang,
    P.Kode_Produk,
    P.Nama_Produk

-- name: GetDataReportRekeningSaldoRekap-filter-TanggalAkhirDefault
COALESCE(RT.saldo, 0.0)

-- name: GetDataReportRekeningSaldoRekap-filter-TanggalAkhirBeforeNow
COALESCE(DB.BALANCE, 0.0)

-- name: GetDataReportRekeningTidakAktif-main
select
	RT.Kode_Cabang,
	RT.Nomor_Rekening,
	RT.Nama_Rekening,
	RL.Nomor_Nasabah,
	RT.Kode_Valuta,
	RT.Saldo,
	N.Nama_Nasabah,
	P.Kode_Produk,
	P.Nama_Produk,
	P.Kode_Produk || ' - ' || P.Nama_Produk as Produk,
	P.Saldo_Minimum,
	RL.Tgl_Trans_Cabang_Terakhir,
	RL.Tgl_Trans_Echannel_Terakhir,
	to_char(RL.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka,
	C.Nama_Cabang,
	epr.enum_description as Status_Resiko
from
	%[1]v.rekeningtransaksi RT
	inner join %[1]v.rekeningliabilitas RL on RT.nomor_rekening = RL.nomor_rekening
	inner join %[1]v.produk P on RL.kode_produk = P.kode_produk
	inner join %[2]v.cabang c on c.kode_cabang = RT.kode_cabang
	inner join %[1]v.nasabah N on RL.nomor_nasabah = N.nomor_nasabah
	left join %[1]v.enum_varchar epr on epr.enum_value = N.status_resiko and epr.enum_name = 'eProfilResiko'
where
	RT.status_rekening = 2
	and ( RT.Kode_Cabang = $1 or 'ALL' = $2)
order by
	RT.Kode_Cabang,
	RT.Kode_Jenis,
	Produk,
	RL.User_Input,
	RT.Nomor_Rekening

-- name: GetDataReportBeaMaterai-main
select
	RT.Nomor_Rekening,
	RT.Nama_Rekening,
	DTB.Nominal as Nominal,
	DTB.Total_Saldo as Saldo
from
	%[1]v.rekeningtransaksi RT,
	%[2]v.bea_materai DTB
where
	RT.Nomor_Rekening = DTB.Nomor_Rekening
	and (rt.kode_cabang = $1 or 'ALL' = $2)
	and (DTB.tanggal >= TO_DATE($3, 'YYYY-MM-DD') or 'ALL' = $4)
	and (DTB.tanggal < TO_DATE($5, 'YYYY-MM-DD') or 'ALL' = $6)
order by
	RT.Kode_Cabang,
	RT.Nomor_Rekening

-- name: GetDataReportTabunganCetakKepalaBuku-main
select
	RT.Nomor_Rekening,
	RT.Nama_Rekening,
	RL.Nomor_Nasabah,
	RT.Kode_Valuta,
	N.Nama_Nasabah,
	P.Kode_Produk || ' - ' || P.Nama_Produk as Produk,
	REP.Nomor_Register_Buku,
	REP.Nomor_Register_Buku_Lama,
	REP.User_Input,
	REP.User_Override,
	REP.Tgl_Input,
	REP.Tgl_Sistem,
	e_al.enum_description as alasan_cetak
from
	%[1]v.rekeningtransaksi RT
	inner join %[1]v.rekeningliabilitas RL on (RT.Nomor_Rekening = RL.Nomor_Rekening)
	inner join %[1]v.nasabah N on (RL.Nomor_Nasabah = N.Nomor_Nasabah)
	inner join %[1]v.produk P on (RL.Kode_Produk = P.Kode_Produk)
	inner join %[2]v.historycetaknamapassbook REP on (REP.Nomor_Rekening = RL.Nomor_Rekening)
	inner join %[3]v.userapp U on (U.id_user = REP.User_Input)
	left join %[1]v.enum_varchar e_al on (e_al.enum_value = REP.alasan_cetak and e_al.enum_name = 'eAlasanCetakPassbook')
where
	(RT.kode_cabang = $1 or 'ALL' = $2)
	and REP.Tgl_Sistem >= TO_DATE($3, 'YYYY-MM-DD')
	and REP.Tgl_Sistem < TO_DATE($4, 'YYYY-MM-DD')
order by
	RT.Kode_Jenis,
	RL.Kode_Produk,
	RT.Nama_Rekening ,
	REP.Tgl_Input

-- name: GetDataReportPemblokiranCekBG-main
select
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah,
    N.Nama_Nasabah,
    P.Kode_Produk || ' - ' || P.Nama_Produk as Produk,
    ND.nomor_seri,
    ND.user_pengubah
from
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,
    %[1]v.nasabah N,
    %[1]v.produk P,
    %[1]v.bukuwarkat BW,
    %[1]v.notadebetinternal ND
where
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RT.Nomor_Rekening = BW.Nomor_Rekening
    and BW.id_bukuwarkat = ND.id_bukuwarkat
    and RL.Nomor_Nasabah = N.Nomor_Nasabah
    and RL.Kode_Produk = P.Kode_Produk
    and RT.Kode_Jenis in ('SAV', 'CA')
    and RT.Kode_Cabang = $1
    and ND.TANGGAL_PERUBAHAN_STATUS >= TO_DATE($2, 'YYYY-MM-DD')
    and ND.TANGGAL_PERUBAHAN_STATUS < TO_DATE($3, 'YYYY-MM-DD')
    and RL.JENIS_REKENING_LIABILITAS = 'G'
    and ND.STATUS_WARKAT in ('S', 'H')
order by
    RT.Kode_Jenis,
    RL.Kode_Produk,
    ND.User_Pengubah

-- name: GetDataReportAktivasiWarkat-main
SELECT 
    RT.Kode_Cabang,
    RT.Nomor_Rekening, 
    RT.Nama_Rekening,
    P.Kode_Produk || ' - ' || P.Nama_Produk as Produk,
    (case BW.Jenis_Warkat when 'C' then 'Cek' else 'Bilyet Giro' end) as Jenis_Warkat,
    BW.Lembar_Tersedia,
    BW.Nomor_Seri_Awal,             
    BW.Nomor_Seri_Akhir,
    (case BW.Is_Otorisasi when 'T' then 'Sudah Otorisasi' else 'Belum Otorisasi' end) as Status_Otorisasi,
    (case BW.Status_Buku when 'A' then 'Resi Sudah Kembali' else 'Resi Belum Kembali' end) as Status_Resi,             
    BW.User_Aktivasi,
    BW.User_Otorisasi_Aktivasi
from
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,           
    %[1]v.produk P,
    %[1]v.bukuwarkat BW
WHERE
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RT.Nomor_Rekening = BW.Nomor_Rekening
    and RL.Kode_Produk = P.Kode_Produk
    and BW.Tanggal_Aktivasi >= TO_DATE($1, 'YYYY-MM-DD')
    and BW.Tanggal_Aktivasi < TO_DATE($2, 'YYYY-MM-DD')           
    and RT.Kode_Cabang = $3
ORDER By
    RT.Kode_Cabang, 
    RL.Kode_Produk, 
    RT.Nomor_Rekening

-- name: GetDataReportRekeningDormanOtomatis-main
SELECT 
    RT.Nomor_Rekening, 
    RT.Nama_Rekening, 
    P.Kode_Produk || ' - ' || P.Nama_Produk as Produk,
    TO_CHAR(REP.Tanggal_Buka, 'DD-MM-YYYY') as Tanggal_Buka,
    (CASE WHEN
        REP.Tgl_Trans_Cabang_Terakhir
        IS NULL
        THEN
        '-'
        ELSE 
            TO_CHAR(REP.Tgl_Trans_Cabang_Terakhir, 'YYYY-MM-DD')
    END) AS Tgl_Trans_Cabang_Terakhir,
    (CASE WHEN
        REP.Tgl_Trans_Echannel_Terakhir
        IS NULL
        THEN
            '-'
        ELSE 
            TO_CHAR(REP.Tgl_Trans_Echannel_Terakhir, 'YYYY-MM-DD')
    END) AS Tgl_Trans_Echannel_Terakhir
FROM
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,           
    %[1]v.produk P,
    %[1]v.v_rekeningdormant REP
WHERE
    RT.Nomor_Rekening = RL.Nomor_Rekening
    and RL.Kode_Produk = P.Kode_Produk          
    and REP.Nomor_Rekening = RT.Nomor_Rekening
    and ( RT.Kode_Cabang = $1 or 'ALL' = $2 )
    and REP.Tanggal_Buka >= TO_DATE($3, 'YYYY-MM-DD')
    and REP.Tanggal_Buka < TO_DATE($4, 'YYYY-MM-DD')
ORDER By
    REP.Tanggal_Buka, 
    RT.Kode_Cabang, 
    RL.Kode_Produk, 
    RT.Nomor_Rekening
    
-- name: GetDataReportRegistrasiATSPerTanggal-main
SELECT 
    (S.DEBITACCOUNTNO  || ' - '  || RT.Nama_Rekening ) as DEBITACCOUNTNO,  
    (S.CREDITACCOUNTNO || ' - '  || RT2.Nama_Rekening ) as CREDITACCOUNTNO,
    S.AMOUNT,
    S.BRANCHCODE,
    RC.TGL_PROSES,
    RC.TGL_KADALUARSA,
    S.DESCRIPTION,
    RL.user_pembuat,
    RL.user_otorisasi 
FROM
    %[1]v.REGISTERLAYANAN RL,
    %[1]v.REGISTERRT RR,           
    %[1]v.RECURRENTTRANSACTION RC,
    %[1]v.STANDINGINSTRUCTION S,
    %[1]v.REKENINGTRANSAKSI RT,
    %[1]v.REKENINGTRANSAKSI RT2
WHERE
    RL.id_register = RR.id_register
    AND RR.recid = RC.recid
    AND RC.instructionid = S.instructionid
    AND S.DEBITACCOUNTNO = RT.nomor_rekening
    AND S.CREDITACCOUNTNO = RT2.nomor_rekening
    AND ( RL.Kode_Cabang_Input = $1 OR  'ALL' = $2)
    AND RL.tanggal_registrasi >= TO_DATE($3, 'YYYY-MM-DD')
    AND RL.tanggal_registrasi < TO_DATE($4, 'YYYY-MM-DD')            
ORDER By
    RT.Nomor_Rekening

-- name: GetDataReportNonAktifATSPerTanggal-main
SELECT 
    (S.DEBITACCOUNTNO  || ' - '  || RT.Nama_Rekening ) as DEBITACCOUNTNO,  
    (S.CREDITACCOUNTNO || ' - '  || RT2.Nama_Rekening ) as CREDITACCOUNTNO,
    S.AMOUNT,
    RL.Kode_Cabang_NonAktif,
    to_char(RL.Tanggal_NonAktif, 'DD-MM-YYYY') Tanggal_NonAktif,
    RL.User_Input_NonAktif,
    RL.User_Otorisasi_NonAktif
FROM
    %[1]v.REGISTERLAYANAN RL,
    %[1]v.REGISTERRT RR,           
    %[1]v.RECURRENTTRANSACTION RC,
    %[1]v.STANDINGINSTRUCTION S,
    %[1]v.REKENINGTRANSAKSI RT,
    %[1]v.REKENINGTRANSAKSI RT2
WHERE
    RL.id_register = RR.id_register
    AND RR.recid = RC.recid
    AND RC.instructionid = S.instructionid
    AND S.DEBITACCOUNTNO = RT.nomor_rekening
    AND S.CREDITACCOUNTNO = RT2.nomor_rekening
    AND ( RL.Kode_Cabang_NonAktif = $1 OR  'ALL' = $2)
    AND RL.Tanggal_NonAktif >= TO_DATE($3, 'YYYY-MM-DD')
    AND RL.Tanggal_NonAktif < TO_DATE($4, 'YYYY-MM-DD')             
ORDER By
    RL.Tanggal_NonAktif, RT.Nomor_Rekening

-- name: GetDataReportRekeningJoinAccount-main
SELECT 
    RT.Nomor_Rekening,
    RT.Nama_Rekening,
    RL.Nomor_Nasabah as Nomor_Nasabah_Utama,
    NJ.Nama_Nasabah as Nama_Nasabah_Utama,
    M.Nomor_Nasabah as Nomor_Nasabah_Join,
    N.Nama_Nasabah as Nama_Nasabah_Join,
    (case when 
    RT.Status_Rekening = 1 then 'Aktif' else 
    'Non Aktif' end) as Status
FROM
    %[1]v.REKENINGTRANSAKSI RT
    INNER JOIN %[1]v.REKENINGLIABILITAS RL ON RT.Nomor_Rekening = RL.Nomor_Rekening
    INNER JOIN %[1]v.NASABAH NJ ON RL.Nomor_Nasabah = NJ.Nomor_Nasabah  
    INNER JOIN %[1]v.PRODUK P ON RL.Kode_Produk = P.Kode_Produk
    LEFT JOIN %[1]v.MULTICIFLINK M ON M.Nomor_Rekening = RL.Nomor_Rekening
    LEFT JOIN %[1]v.NASABAH N ON M.Nomor_Nasabah = N.Nomor_Nasabah           
WHERE
    (RT.Kode_Cabang = $1 or 'ALL' = $2)
    and RL.Is_Join_Account = 'T'
    and RT.Kode_Jenis in ('SAV', 'CA', 'DEP')
    AND RL.Tanggal_Buka >= TO_DATE($3, 'YYYY-MM-DD')
    AND RL.Tanggal_Buka < TO_DATE($4, 'YYYY-MM-DD')   
ORDER BY
    RL.TANGGAL_BUKA,
    RT.Nomor_Rekening 

-- name: GetDataReportWICBaru-main
SELECT 
        n.Nomor_Nasabah, 
        n.Nama_Nasabah,
        to_char(n.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka,
        n.Jenis_Nasabah,
        n.User_Input,
        n.User_Otorisasi,
        (case when n.jenis_nasabah = 'I' then e_jid.enum_description else 'NPWP' end) as jenis_identitas,
        (case when n.jenis_nasabah = 'I' then i.nomor_identitas else nk.nomor_npwp end) as nomor_identitas,
        (case when n.jenis_nasabah = 'I' then i.alamat_rumah_jalan else nk.alamat_jalan end) as alamat
FROM
    %[1]v.NASABAH n
    LEFT JOIN %[1]v.NASABAHINDIVIDU ni ON (n.Nomor_Nasabah = ni.Nomor_Nasabah)
    LEFT JOIN %[1]v.INDIVIDU i ON (i.Id_Individu = ni.Id_Individu)
    LEFT JOIN %[1]v.NASABAHKORPORAT nk ON (n.Nomor_Nasabah = nk.Nomor_Nasabah)
    LEFT JOIN %[1]v.ENUM_VARCHAR e_jid ON (e_jid.enum_value = i.jenis_identitas and e_jid.enum_name = 'eJenisIdentitas')
WHERE
    ( n.Kode_Cabang_Input = $1 OR 'ALL' = $2 )
    and n.TANGGAL_BUKA >= TO_DATE($3, 'YYYY-MM-DD')
    and n.TANGGAL_BUKA < TO_DATE($4, 'YYYY-MM-DD')
    and n.IS_WIC = 'T'
ORDER BY 
    n.TANGGAL_BUKA

-- name: GetDataReportDepositoJatuhTempoHariIni-main
SELECT 
    RT.Nomor_Rekening
    , RT.Nama_Rekening
    , RL.Saldo_Deposito
    , RL.Nomor_Nasabah
    , DEP.Nomor_Bilyet
    , to_char(RL.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka
    , to_char(DEP.Tanggal_Jatuh_Tempo_Terakhir, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Terakhir
    , to_char(DEP.Tanggal_Jatuh_Tempo_Berikutnya, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Berikutnya
    , (case JTD.disposisi
        when 'A' then 'Aro'      
        when 'P' then 'Pindah buku'       
        when 'R' then 'Transfer SKN' 
        when 'G' then 'Transfer RTGS' 
        --when 'T' then 'Tunai' 
        else 'unknown'
    end) as Disposisi_Nominal
    , DEP.Rekening_Disposisi
    , DEP.Penerima_Transfer_Disposisi
    , JTD.Nisbah_Bagi_Hasil
    , P.Masa_Perjanjian
    , JTD.Nisbah_Spesial
    , JTD.nomor_disposisi AS Rekening_Disposisi
    , (CASE RL.Disposisi_bagi_hasil
        WHEN 'H' THEN 'Dihibahkan'      
        WHEN 'K' THEN 'Dikapitalisir'       
        WHEN 'P' THEN 'Pindah Buku'
        WHEN 'R' THEN 'Transfer SKN' 
        WHEN 'G' THEN 'Transfer RTGS'
        WHEN 'T' THEN 'Tunai' 
        ELSE 'unknown'
    END) AS Disposisi_bagi_hasil
    , RL.rekening_disposisi_bagi_hasil
    , RL.penerima_transfer_bagi_hasil              
FROM
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,
    %[1]v.DEPOSITO DEP,
    %[2]v.JATUHTEMPO_DEPOSITO JTD,
    %[1]v.PRODUKDEPOSITO P
WHERE
    RT.Nomor_Rekening = RL.Nomor_Rekening
    AND RT.Nomor_Rekening = DEP.Nomor_Rekening
    AND JTD.Nomor_Rekening = DEP.Nomor_Rekening
    AND RL.Kode_Produk = P.Kode_Produk
    AND RT.Kode_Jenis in ('DEP')
    AND RL.Tanggal_Buka < TO_DATE($1, 'YYYY-MM-DD')
    AND (rt.kode_cabang = $2 OR 'ALL' = $3)
    AND (JTD.tanggal >= TO_DATE($4, 'YYYY-MM-DD'))
    AND (JTD.tanggal < TO_DATE($5, 'YYYY-MM-DD'))
ORDER BY
    RT.Kode_Cabang, RL.Kode_Produk, RT.Nomor_Rekening 

-- name: GetDataReportDepositoOutstanding-main
SELECT
    RT.Nomor_Rekening, 
    RT.Nama_Rekening, 
    RL.Nomor_Nasabah,
    DEP.Nomor_Bilyet,
    (CASE DEP.disposisi_nominal 
    WHEN 'A' then 'Aro'      
    WHEN 'P' then 'Pindah buku'       
    WHEN 'R' then 'Transfer SKN' 
    WHEN 'G' then 'Transfer RTGS' 
    ELSE 'unknown'
    END) AS Disposisi_Nominal, 
    (CASE RL.disposisi_bagi_hasil 
    WHEN 'K' then 'Kapitalisir'      
    WHEN 'P' then 'Pindah buku'       
    WHEN 'R' then 'Transfer SKN' 
    WHEN 'G' then 'Transfer RTGS' 
    ELSE 'unknown'
    END) AS Disposisi_Bagi_Hasil, 
    RL.Nomor_Rekening_Disposisi,
    PD.Masa_Perjanjian,
    (CASE PD.periode_perjanjian 
    WHEN 'H' then 'Hari' 
    WHEN 'B' then 'Bulan' 
    WHEN 'T' then 'Tahun' 
    ELSE 'unknown' 
    END) AS Periode_Perjanjian,
    RL.Nisbah_Dasar,
    RL.Nisbah_Spesial,
    DEP.Rekening_Disposisi,
    DEP.masa_perjanjian
FROM
    %[1]v.rekeningtransaksi RT,
    %[1]v.rekeningliabilitas RL,
    %[1]v.DEPOSITO DEP,
    %[1]v.PRODUK P,
    %[1]v.PRODUKDEPOSITO PD
WHERE
    RT.Nomor_Rekening = RL.Nomor_Rekening
    AND RT.Nomor_Rekening = DEP.Nomor_Rekening
    AND RL.Kode_Produk = P.Kode_Produk
    AND P.Kode_Produk = PD.Kode_Produk
    AND RT.Kode_Jenis in ('DEP')
    AND RT.Status_Rekening != 3
    AND (RT.Kode_Cabang = $1 OR 'ALL' = $2)
ORDER BY
    RT.Kode_Cabang, PD.Masa_Perjanjian, RL.Kode_Produk

-- name: GetDataReportDepositoBagiHasil-main
SELECT 
    RT.Nomor_Rekening, 
    RT.Nama_Rekening, 
    RL.Saldo_Deposito,
    RL.Nomor_Nasabah,
    DEP.Nomor_Bilyet,
    --RL.Nisbah_Bagi_Hasil,
    RL.Nomor_Rekening_Disposisi,
    DTB.Nominal_BagiHasil,
    DTB.Nominal_Rate,
    --DTB.Periode_Awal,
    --DTB.Periode_Akhir,        
    DTB.Pajak,
    DTB.Zakat,
    --DTB.Saldo,
    DTB.Ekuivalen_Rate,    
    (DTB.Nominal_Total - DTB.Pajak - DTB.Zakat) AS Nominal_Bersih,
    (CASE DTB.Disposisi
    WHEN 'H' THEN 'Dihibahkan'      
    WHEN 'K' THEN 'Dikapitalisir'       
    WHEN 'P' THEN 'Pindah Buku'
    WHEN 'R' THEN 'Transfer SKN' 
    WHEN 'G' THEN 'Transfer RTGS'
    WHEN 'T' THEN 'Tunai' 
    ELSE 'unknown'
    END) AS Disposisi_bagi_hasil,
    --RL.rekening_disposisi_bagi_hasil,
    --RL.penerima_transfer_bagi_hasil,
    --RL.kode_bank,
    coalesce(BKLI.nama_lengkap,m.nama_member) AS nama_lengkap_bank
FROM
    %[1]v.rekeningtransaksi RT
    INNER JOIN %[1]v.rekeningliabilitas RL ON RL.Nomor_Rekening = RT.Nomor_Rekening
    INNER JOIN %[1]v.DEPOSITO DEP ON DEP.Nomor_Rekening = RT.Nomor_Rekening
    INNER JOIN %[1]v.PRODUKDEPOSITO P ON P.Kode_Produk = RL.Kode_Produk
    INNER JOIN %[2]v.BAGIHASIL_DEPOSITO DTB ON DTB.Nomor_Rekening = RT.Nomor_Rekening
    LEFT JOIN %[1]v.BANKKLIRING BKLI ON BKLI.kode_bank = RL.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m ON m.kode_member = RL.kode_bank
WHERE
    RT.Kode_Jenis in ('DEP')
    AND (rt.Kode_Cabang = $1 OR 'ALL' = $2)
    AND (DTB.tanggal >= TO_DATE($3, 'YYYY-MM-DD'))
    AND (DTB.tanggal < TO_DATE($4, 'YYYY-MM-DD'))
ORDER BY
    RT.Kode_Cabang, RL.Kode_Produk, RT.Nomor_Rekening 

-- name: GetDataReportTplDepositoBaru-main
SELECT RT.Nomor_Rekening,
    to_char(RL.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka,
    RL.Nisbah_Dasar,
    RL.Nisbah_Spesial,
    RL.Nomor_Rekening_Disposisi,
    RL.Penerima_Transfer_Bagi_Hasil, 
    to_char(RL.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka,
    RT.Nama_Rekening,
    RT.Saldo, 
    RL.Kode_Produk,
    P.Nama_Produk,
    P.Nisbah_Bonus_Dasar,
    RL.Saldo_Deposito,
    DEP.Nomor_Bilyet,
    DEP.Rekening_Disposisi,
    DEP.Penerima_Transfer_Disposisi,
    to_char(DEP.Tanggal_Jatuh_Tempo_Terakhir, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Terakhir,
    to_char(DEP.Tanggal_Jatuh_Tempo_Berikutnya, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Berikutnya,
    (case DEP.disposisi_nominal 
    when 'A' then 'Aro'      
    when 'P' then 'Pindah buku'       
    when 'R' then 'Transfer SKN' 
    when 'G' then 'Transfer RTGS' 
    else 'unknown'
    end) as Disposisi_Nominal,
    DEP.Masa_Perjanjian,
    (case DEP.Periode_Perjanjian 
        when 'H' then 'Hari' 
        when 'B' then 'Bulan' 
        when 'T' then 'Tahun' 
        else 'unknown' 
    end
    ) as Periode_Perjanjian,
    DEP.Ekuivalen_Rate,
    (case RL.Disposisi_Bagi_Hasil 
    when 'K' then 'Kapitalisir'      
    when 'P' then 'Pindah buku'       
    when 'R' then 'Transfer SKN' 
    when 'G' then 'Transfer RTGS' 
    else 'unknown'
    end) as Disposisi_Bagi_Hasil,
    RL.Nisbah_Bagi_Hasil,
    RT.User_Input,
    User_Otorisasi,
    DEP.Masa_Perjanjian || ' ' || Periode_Perjanjian as maper_perjanjian,
    RL.Kode_Produk || ' ' || P.Nama_Produk as KODE_NAMA_PRODUK,
    1 as data_count 
FROM
%[1]v.REKENINGTRANSAKSI RT,
%[1]v.REKENINGLIABILITAS RL,
%[1]v.DEPOSITO DEP,
%[1]v.PRODUK P
WHERE
RT.Nomor_Rekening = RL.Nomor_Rekening
and RT.Nomor_Rekening = DEP.Nomor_Rekening
and RL.Kode_Produk = P.Kode_Produk
and RT.Kode_Jenis in ('DEP')
and (RT.Kode_Cabang = $1 OR 'ALL CABANG'= $2)
AND (RL.TANGGAL_BUKA >= TO_DATE($3, 'YYYY-MM-DD') AND RL.TANGGAL_BUKA < TO_DATE($4, 'YYYY-MM-DD'))
and rt.status_rekening = 1
ORDER By
RT.Kode_Jenis, RL.Kode_Produk, RT.Nomor_Rekening 

-- name: GetDataReportTplPerubahanNasabah-main
SELECT 
    C.CUSTOMER_NO,  
    C.ACCOUNT_NO,
    C.ACCOUNT_TYPE,
    C.TABLE_NAME,
    C.FIELD_NAME,
    C.BEFORE_DATA,
    C.AFTER_DATA,
    C.DATA_REMARK,
    C.USER_INPUT,
    C.USER_APPROVE,
    C.INPUT_DATE,
    C.APPROVE_DATE,
    C.FORM_NAME, 
    C.FUNCTION_NAME,
    C.BRANCH_CODE 
FROM
    %[1]v.CUSTCHANGEHISTORI C
WHERE
    C.ACCOUNT_TYPE IN ('C')
    AND (C.BRANCH_CODE = $1 OR 'ALL' = $2 )
    AND C.INPUT_DATE >= TO_DATE($3, 'YYYY-MM-DD')
    AND C.INPUT_DATE < TO_DATE($4, 'YYYY-MM-DD')  
ORDER BY
    C.CUSTOMER_NO  , C.ID_CUSTCHANGEHISTORI 

-- name: GetDataReportTplDepositoBreak-main
SELECT RT.Nomor_Rekening, 
    RT.Nama_Rekening,
    to_char(RL.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka,
    RL.Saldo_Deposito,
    RL.Nomor_Nasabah, 
    RL.Kode_Produk,
    P.Nama_Produk,
    DEP.Nomor_Bilyet,
    to_char(DEP.Tanggal_Jatuh_Tempo_Terakhir, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Terakhir,
    to_char(DEP.Tanggal_Jatuh_Tempo_Berikutnya, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Berikutnya,
    (case DEP.disposisi_nominal 
    when 'A' then 'Aro'      
    when 'P' then 'Pindah buku'       
    when 'R' then 'Transfer' 
    when 'G' then 'Transfer' 
    else 'unknown'
    end) as Disposisi_Nominal,
    DEP.Rekening_Disposisi,
    DEP.Penerima_Transfer_Disposisi,
    DEP.Masa_Perjanjian,
    (case DEP.Periode_Perjanjian 
        when 'H' then 'Hari' 
        when 'B' then 'Bulan' 
        when 'T' then 'Tahun' 
        else 'unknown' 
    end
    ) as Periode_Perjanjian,
    DEP.Masa_Perjanjian || ' ' || Periode_Perjanjian as maper_perjanjian,
    RL.Nisbah_Bagi_Hasil,
    I.Nilai_Pencairan              
FROM
%[1]v.REKENINGTRANSAKSI RT,
%[1]v.REKENINGLIABILITAS RL,
%[1]v.DEPOSITO  DEP,
%[1]v.PRODUK P,
%[1]v.INFOPENCAIRANAWALDEPOSITO I
WHERE
RT.Nomor_Rekening = RL.Nomor_Rekening
and RT.Nomor_Rekening = DEP.Nomor_Rekening
and RL.Kode_Produk = P.Kode_Produk
and DEP.Nomor_Rekening = I.Nomor_Rekening
and RT.Kode_Jenis in ('DEP')
and RT.Kode_Cabang = $1
and I.TANGGAL_PROSES >= TO_DATE($2,'YYYY-MM-DD')
and I.TANGGAL_PROSES < TO_DATE($3,'YYYY-MM-DD')
ORDER BY
RT.Kode_Jenis, RL.Kode_Produk, RT.Nomor_Rekening 

-- name: GetDataReportTplRekRencanaBaru-main
SELECT RT.Nomor_Rekening
    , RT.Nama_Rekening
    , RL.Nomor_Nasabah
    , N.Nama_Nasabah
    , P.Kode_Produk
    , P.Nama_Produk
    , P.Kode_Produk || ' ' || P.Nama_Produk as kode_nama_produk
    , P.Saldo_Minimum
    , RT.User_Input
    , RL.User_Otorisasi
    , RL.Kode_Program
    , to_char(RL.Tanggal_Buka, 'DD-MM-YYYY') Tanggal_Buka
    , to_char(RR.Tanggal_Jatuh_Tempo, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo
FROM
    %[1]v.REKENINGTRANSAKSI RT
, %[1]v.REKENINGLIABILITAS RL
, %[1]v.REKENINGRENCANA RR           
, %[1]v.NASABAH N
, %[1]v.PRODUK P
WHERE
RT.Nomor_Rekening = RL.Nomor_Rekening
and RL.Nomor_Rekening = RR.Nomor_Rekening
and RL.Nomor_Nasabah = N.Nomor_Nasabah
and RL.Kode_Produk = P.Kode_Produk
and RT.Kode_Jenis in ('SAV')
and RL.Jenis_Rekening_Liabilitas = 'R'
and RT.Kode_Cabang = $1
and RL.TANGGAL_BUKA >= TO_DATE($2, 'YYYY-MM-DD')
and RL.TANGGAL_BUKA < TO_DATE($3, 'YYYY-MM-DD')
ORDER By
RT.Kode_Jenis, RL.Kode_Produk, RL.User_Input, RT.Nomor_Rekening 

-- name: GetDataReportTplNasabahReposisi-main
SELECT 
N.Nomor_Nasabah, 
    N.Nama_Nasabah, 
    CAB1.KODE_CABANG || ' - ' || CAB1.NAMA_CABANG AS CABANG_ASAL,
    CAB2.KODE_CABANG || ' - ' || CAB2.NAMA_CABANG AS CABANG_TUJUAN,
    REP.Cabang_Input,
    REP.User_Input,
    REP.User_Otorisasi,
    to_char(REP.Tanggal, 'DD-MM-YYYY') Tanggal, 
    1 as count                                        
FROM
    %[1]v.NASABAH N,
    %[2]v.CABANG CAB1,
    %[2]v.CABANG CAB2,
    %[1]v.INFOREPOSISINASABAH REP
WHERE          
    REP.Nomor_Nasabah = N.Nomor_Nasabah
    and REP.Cabang_Asal = CAB1.Kode_Cabang
    and REP.Cabang_Tujuan = CAB2.Kode_Cabang          
    and (REP.Cabang_Input = $1 or REP.Cabang_Tujuan = $1)
    and REP.Tanggal >= TO_DATE($2,'dd-mm-yyyy')
    and REP.Tanggal < TO_DATE($3,'dd-mm-yyyy')          
ORDER By
    REP.TANGGAL, REP.Cabang_Input 

-- name: GetDataReportTplNasabahGabung-main
SELECT 
    I.KODE_CABANG_INPUT,
    I.NOMOR_NASABAH, 
    I.NOMOR_NASABAH_UTAMA,
    to_char(I.TANGGAL_INPUT, 'DD-MM-YYYY') TANGGAL_INPUT,
    to_char(I.TANGGAL_OTORISASI, 'DD-MM-YYYY') TANGGAL_OTORISASI,
    to_char(I.TANGGAL_PROSES, 'DD-MM-YYYY') TANGGAL_PROSES,
    I.USER_INPUT,
    I.USER_OTORISASI,
    N.NAMA_NASABAH,
    N.KODE_CABANG_INPUT AS CABANG_NASABAH,
    O.NAMA_NASABAH AS NAMA_NASABAH_UTAMA,
    O.KODE_CABANG_INPUT AS CABANG_NASABAH_UTAMA,
    1 AS DATA_COUNT
FROM
    %[1]v.INFOMERGENASABAH I,
    %[1]v.NASABAH N,
    %[1]v.NASABAH O
WHERE
    I.NOMOR_NASABAH = N.NOMOR_NASABAH
    AND I.NOMOR_NASABAH_UTAMA = O.NOMOR_NASABAH
    AND I.TANGGAL_PROSES >=  TO_TIMESTAMP($1,'DD-MM-YYYY')
    AND I.TANGGAL_PROSES <= TO_TIMESTAMP($2,'DD-MM-YYYY')
    AND (I.KODE_CABANG_INPUT = $3 OR 'ALL' = $4)
ORDER BY
    KODE_CABANG_INPUT,NOMOR_NASABAH

-- name: GetDataReportTplDepositoJatuhTempoHariIniTransfer-main
SELECT 
    RT.NOMOR_REKENING, 
    RT.NAMA_REKENING, 
    RL.SALDO_DEPOSITO,
    RL.NOMOR_NASABAH,
    DEP.NOMOR_BILYET,
    to_char(DEP.TANGGAL_JATUH_TEMPO_TERAKHIR, 'DD-MM-YYYY') TANGGAL_JATUH_TEMPO_TERAKHIR,
    to_char(DEP.TANGGAL_JATUH_TEMPO_BERIKUTNYA, 'DD-MM-YYYY') TANGGAL_JATUH_TEMPO_BERIKUTNYA,
    (CASE JTD.DISPOSISI
    --WHEN 'A' THEN 'ARO'      
    --WHEN 'P' THEN 'PINDAH BUKU'       
    WHEN 'R' THEN 'TRANSFER SKN' 
    WHEN 'G' THEN 'TRANSFER RTGS' 
    --WHEN 'T' THEN 'TUNAI' 
    ELSE 'UNKNOWN'
    END) AS DISPOSISI_NOMINAL, 
    P.MASA_PERJANJIAN,
    (CASE P.PERIODE_PERJANJIAN 
        WHEN 'H' THEN 'HARI' 
        WHEN 'B' THEN 'BULAN' 
        WHEN 'T' THEN 'TAHUN' 
        ELSE 'UNKNOWN' 
    END
    ) AS PERIODE_PERJANJIAN,
    RL.NISBAH_BAGI_HASIL,
    to_char(RL.TANGGAL_BUKA, 'DD-MM-YYYY') TANGGAL_BUKA,             
    JTD.NOMOR_DISPOSISI AS REKENING_DISPOSISI,
    DEP.PENERIMA_TRANSFER_DISPOSISI,
    DEP.KODE_BANK,
    NVL(BKLI.NAMA_LENGKAP, M.NAMA_MEMBER) AS NAMA_LENGKAP_BANK              
FROM
    %[1]v.REKENINGTRANSAKSI RT 
    INNER JOIN  %[1]v.REKENINGLIABILITAS  RL ON RL.NOMOR_REKENING = RT.NOMOR_REKENING
    INNER JOIN  %[1]v.DEPOSITO  DEP ON DEP.NOMOR_REKENING = RL.NOMOR_REKENING
    INNER JOIN  %[2]v.JATUHTEMPO_DEPOSITO JTD ON JTD.NOMOR_REKENING = DEP.NOMOR_REKENING
    INNER JOIN  %[1]v.PRODUKDEPOSITO P ON RL.KODE_PRODUK = P.KODE_PRODUK
    LEFT JOIN %[1]v.BANKKLIRING BKLI ON BKLI.KODE_BANK = DEP.KODE_BANK
    LEFT JOIN %[1]v.MEMBERRTGS M ON M.KODE_MEMBER = DEP.KODE_BANK 
WHERE
    RT.KODE_JENIS IN ('DEP')
    AND JTD.DISPOSISI IN ('R','G')
    AND RL.TANGGAL_BUKA < TO_TIMESTAMP($1, 'YYYY-MM-DD')
    AND (RT.KODE_CABANG = $2 OR 'ALL'= $3)
    AND (JTD.TANGGAL >= TO_TIMESTAMP($4, 'YYYY-MM-DD') OR 'ALL' = $5)
    AND (JTD.TANGGAL < TO_TIMESTAMP($1, 'YYYY-MM-DD') OR 'ALL'= $6)
ORDER BY
    RT.KODE_CABANG, RL.KODE_PRODUK, RT.NOMOR_REKENING 

-- name: GetDataReportTplDepositoBlokir-main
SELECT RT.Nomor_Rekening, 
RT.Nama_Rekening,
RL.Nomor_Nasabah,
RT.Kode_Valuta,
DEP.Nomor_Bilyet,
to_char(DEP.Tanggal_Jatuh_Tempo_Terakhir, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Terakhir,             
to_char(DEP.Tanggal_Jatuh_Tempo_Berikutnya, 'DD-MM-YYYY') Tanggal_Jatuh_Tempo_Berikutnya,  
DEP.Disposisi_Nominal,
DEP.Masa_Perjanjian,
DEP.Periode_Perjanjian,
N.Nama_Nasabah, 
P.Kode_Produk,
P.Nama_Produk,
P.Kode_Produk || ' - ' || P.Nama_Produk AS namkod_produk,
HB.User_Input,
HB.User_DualControl,
HB.Keterangan_Blokir,
HB.Sistem_Pemblokir,
to_char(HB.Tanggal_Ubah, 'DD-MM-YYYY') Tanggal_Ubah,             
HB.ID_HISTBLOKIR 
FROM
%[1]v.REKENINGTRANSAKSI RT,
%[1]v.REKENINGLIABILITAS RL,
%[1]v.DEPOSITO DEP,           
%[1]v.HISTBLOKIR HB,
%[1]v.NASABAH N,           
%[1]v.PRODUK P
WHERE
RT.Nomor_Rekening = RL.Nomor_Rekening          
and RL.Nomor_Nasabah = N.Nomor_Nasabah
and RL.Nomor_Rekening = DEP.Nomor_Rekening          
and HB.Nomor_Rekening = RT.Nomor_Rekening
and RL.Kode_Produk = P.Kode_Produk
and RT.Kode_Jenis in ('DEP')
and HB.Status_Blokir = 'T'
and (HB.Kode_Cabang_Input = $1 or 'ALL' = $2)         
and HB.Tanggal_Ubah >= TO_DATE($3, 'YYYY-MM-DD')
and HB.Tanggal_Ubah < TO_DATE($4, 'YYYY-MM-DD')
ORDER BY
RT.Kode_Jenis, RL.Kode_Produk, RT.Nama_Rekening 

-- name: GetDataReportTplDepositoAROHariIni-main
SELECT RT.Nomor_Rekening, 
    RT.Nama_Rekening, 
    RL.Saldo_Deposito,
    RL.Nomor_Nasabah,
    DEP.Nomor_Bilyet,
    (case DEP.disposisi_nominal 
    when 'A' then 'Aro'      
    when 'P' then 'Pindah buku'       
    when 'R' then 'Transfer' 
    when 'G' then 'Transfer' 
    else 'unknown'
    end) as Disposisi_Nominal, 
    P.Masa_Perjanjian,
    (case P.periode_perjanjian 
        when 'H' then 'Hari' 
        when 'B' then 'Bulan' 
        when 'T' then 'Tahun' 
        else 'unknown' 
    end
    ) as Periode_Perjanjian,
    RL.Nisbah_Bagi_Hasil,
    DEP.Rekening_Disposisi
FROM
%[1]v.REKENINGTRANSAKSI RT,
%[1]v.REKENINGLIABILITAS RL,
%[1]v.DEPOSITO DEP,
%[1]v.PRODUKDEPOSITO P
WHERE
RT.Nomor_Rekening = RL.Nomor_Rekening
and RT.Nomor_Rekening = DEP.Nomor_Rekening
and RL.Kode_Produk = P.Kode_Produk
and RT.Kode_Jenis in ('DEP')
and RT.Kode_Cabang = $1
and RL.Tanggal_Buka < TO_DATE($2, 'YYYY-MM-DD') 
and DEP.tanggal_jatuh_tempo_terakhir = TO_DATE($3, 'YYYY-MM-DD')
AND DEP.is_aro = 'T'
ORDER By
RT.Kode_Cabang, RL.Kode_Produk, RT.Nomor_Rekening 

-- name: GetDataReportTplDepositoJatuhTempoBesok-main
SELECT
    RT.NOMOR_REKENING,
    RT.NAMA_REKENING,
    RL.SALDO_DEPOSITO,
    RL.NOMOR_NASABAH,
    DEP.NOMOR_BILYET,
    (CASE
        DEP.DISPOSISI_NOMINAL
    WHEN 'A' THEN 'ARO'
        WHEN 'P' THEN 'PINDAH BUKU'
        WHEN 'R' THEN 'TRANSFER'
        WHEN 'G' THEN 'TRANSFER'
        --WHEN 'T' THEN 'TUNAI' 
        ELSE 'UNKNOWN'
    END) AS DISPOSISI_NOMINAL,
    P.MASA_PERJANJIAN,
    (CASE
        P.PERIODE_PERJANJIAN 
    WHEN 'H' THEN 'HARI'
        WHEN 'B' THEN 'BULAN'
        WHEN 'T' THEN 'TAHUN'
        ELSE 'UNKNOWN'
    END) AS PERIODE_PERJANJIAN,
    RL.NISBAH_BAGI_HASIL,
    DEP.REKENING_DISPOSISI AS REKENING_DISPOSISI,
    to_char(DEP.TANGGAL_JATUH_TEMPO_BERIKUTNYA, 'DD-MM-YYYY') TGL_JATUH_TEMPO,             
    1 AS DATA_COUNT
FROM
    %[1]v.REKENINGTRANSAKSI RT,
    %[1]v.REKENINGLIABILITAS RL,
    %[1]v.DEPOSITO DEP,
    %[1]v.PRODUKDEPOSITO  P
WHERE
    RT.NOMOR_REKENING = RL.NOMOR_REKENING
    AND RT.NOMOR_REKENING = DEP.NOMOR_REKENING
    AND RL.KODE_PRODUK = P.KODE_PRODUK
    AND RT.KODE_JENIS IN ('DEP')
    AND RL.TANGGAL_BUKA <= TO_TIMESTAMP($1, 'YYYY-MM-DD')
    AND (RT.KODE_CABANG = $2 OR 'ALL' = $3)
    AND (DEP.TANGGAL_JATUH_TEMPO_BERIKUTNYA >= TO_TIMESTAMP($4, 'YYYY-MM-DD') OR 'ALL' = $5)
    AND (DEP.TANGGAL_JATUH_TEMPO_BERIKUTNYA < TO_TIMESTAMP($1, 'YYYY-MM-DD') OR 'ALL' = $6)
    AND RT.STATUS_REKENING != 3
ORDER BY
    DEP.TANGGAL_JATUH_TEMPO_BERIKUTNYA,
    RT.KODE_CABANG,
    RL.KODE_PRODUK,
    RT.NOMOR_REKENING

-- name: GetDataReportTplDepositoBagHasTransfer-main
SELECT 
    RT.NOMOR_REKENING, 
    RT.NAMA_REKENING, 
    RL.SALDO_DEPOSITO,
    RL.NOMOR_NASABAH,
    DEP.NOMOR_BILYET,
    RL.NISBAH_BAGI_HASIL,
    RL.NOMOR_REKENING_DISPOSISI,
    DTB.NOMINAL_TOTAL,
    DTB.PAJAK,
    DTB.ZAKAT,
    (DTB.NOMINAL_TOTAL - DTB.PAJAK - DTB.ZAKAT) AS NOMINAL_BERSIH,
    (CASE DTB.DISPOSISI
    --WHEN 'H' THEN 'DIHIBAHKAN'      
    --WHEN 'K' THEN 'DIKAPITALISIR'       
    --WHEN 'P' THEN 'PINDAH BUKU'
    WHEN 'R' THEN 'TRANSFER' 
    WHEN 'G' THEN 'TRANSFER'
    --WHEN 'T' THEN 'TUNAI' 
    ELSE 'UNKNOWN'
    END) AS DISPOSISI_BAGI_HASIL,
    RL.REKENING_DISPOSISI_BAGI_HASIL,
    RL.PENERIMA_TRANSFER_BAGI_HASIL,
    RL.KODE_BANK,
    NVL(BKLI.NAMA_LENGKAP,M.NAMA_MEMBER) AS NAMA_LENGKAP_BANK
FROM
    %[1]v.REKENINGTRANSAKSI RT
    INNER JOIN %[1]v.REKENINGLIABILITAS RL ON RL.NOMOR_REKENING = RT.NOMOR_REKENING
    INNER JOIN %[1]v.DEPOSITO DEP ON DEP.NOMOR_REKENING = RT.NOMOR_REKENING
    INNER JOIN %[1]v.PRODUKDEPOSITO P ON P.KODE_PRODUK = RL.KODE_PRODUK
    INNER JOIN %[2]v.BAGIHASIL_DEPOSITO DTB ON DTB.NOMOR_REKENING = RT.NOMOR_REKENING
    LEFT JOIN %[1]v.BANKKLIRING BKLI ON BKLI.KODE_BANK = RL.KODE_BANK
    LEFT JOIN %[1]v.MEMBERRTGS M ON M.KODE_MEMBER = RL.KODE_BANK
WHERE
    RT.KODE_JENIS IN ('DEP')
    --AND T.KODE_TRANSAKSI = 'POBHD'
    AND DTB.DISPOSISI IN ('R','G')
    --AND DTB.PAJAK > 0.0          
    AND (RT.KODE_CABANG = $1 OR 'ALL' = $2)
    AND (DTB.TANGGAL >= TO_TIMESTAMP($3, 'YYYY-MM-DD') OR 'ALL' = $4)
    AND (DTB.TANGGAL < TO_TIMESTAMP($5, 'YYYY-MM-DD') OR 'ALL' = $6)
ORDER BY
    RT.KODE_CABANG, RL.KODE_PRODUK, RT.NOMOR_REKENING 

-- name: GetDataReportTplDepositoSaldoRekap-main
SELECT RT.Kode_Cabang,
    P.Kode_Produk,
    P.Nama_Produk, 
    sum(RT.Saldo) as Saldo,
    sum(1) as jumlah_total ,
    sum(case when status_rekening <> 3 then 1 else 0 end) as jumlah_aktif,
    sum(case when status_rekening = 3 then 1 else 0 end) as jumlah_tutup
FROM
%[1]v.REKENINGTRANSAKSI RT
LEFT OUTER JOIN GETDAILYBALANCEREK_BYDATE(TO_DATE($1, 'YYYY-MM-DD') + 1) db
    on (db.nomor_rekening = rt.nomor_rekening and db.balance_field='saldo')
INNER JOIN %[1]v.REKENINGCUSTOMER RC ON RT.Nomor_Rekening = RC.Nomor_Rekening
INNER JOIN %[1]v.REKENINGLIABILITAS RL ON RC.Nomor_Rekening = RL.Nomor_Rekening
INNER JOIN %[1]v.PRODUK  P ON RL.Kode_Produk = P.Kode_Produk
WHERE
RT.Kode_Jenis in ('DEP')
and (RT.Kode_Cabang = $2 OR 'ALL' = $3)
GROUP BY
RT.Kode_Cabang, P.Kode_Produk, P.Nama_Produk      
ORDER By
RT.Kode_Cabang, P.Kode_Produk, P.Nama_Produk

-- name: GetDataReportTransaksiSendiri-main
SELECT 		t.ID_TRANSAKSI, 
    RT.Nomor_Rekening, 
    RT.Nama_Rekening,
    RT.Kode_Cabang as Kode_Cabang_Rek,
    T.Kode_Cabang_Transaksi,
    T.Keterangan,
    T.Nomor_Referensi,
    T.Id_Transaksi,
    T.User_Overide,
    T.User_Otorisasi,
    T.Kode_Transaksi,
    TO_CHAR(T.Tanggal_Transaksi, 'DD-MM-YYYY') || ' ' || TO_CHAR(T.Jam_Input, 'HH24::MI::SS') as Jam_Input,
    DT.Kode_Valuta,
    (case DT.Jenis_Mutasi
    when 'D' then DT.Nilai_Mutasi
    else 0.0
    end) as Nominal_Debet,
    (case DT.Jenis_Mutasi
    when 'C' then DT.Nilai_Mutasi
    else 0.0
    end) as Nominal_Kredit,
    DT.Jenis_Mutasi,
    DT.Id_Detil_Transaksi,
    DT.Kode_RC,
    A.Account_Code, 
    A.Account_Name,
    I.User_Override_Passbook,
    I.Is_Passbook_Wajib,
    I.Is_Passbook_Dicetak,
    I.Kd_Sumber_Dana_Trx,
    I.Kd_Tujuan_Trx,
    I.Ket_Sumber_Dana_Trx,
    I.Ket_Tujuan_Trx,
    I.Tipe_Akses_Transaksi,
    I.ListNama_Akses_Transaksi,
    DT.SEQUENCE,
    A.Account_Code|| ' - ' || RT.Nama_Rekening AS Nama_No_Rek
FROM %[1]v.TRANSAKSI t 
    INNER JOIN %[1]v.DETILTRANSAKSI DT on T.Id_Transaksi = DT.Id_Transaksi
    INNER JOIN %[1]v.REKENINGTRANSAKSI RT on DT.Nomor_Rekening = RT.Nomor_Rekening
    LEFT OUTER JOIN %[1]v.INFOTRANSAKSI I on (I.Id_Transaksi = T.Id_Transaksi)
    LEFT OUTER JOIN %[1]v.ACCOUNT A on (A.Account_Code = DT.Kode_Account)
WHERE 
    t.USER_INPUT = $1
    AND T.TANGGAL_TRANSAKSI >= TO_TIMESTAMP($2,'YYYY-MM-DD') 
    AND T.TANGGAL_TRANSAKSI < TO_TIMESTAMP($3,'YYYY-MM-DD') + 1
ORDER By
    Kode_Cabang_Transaksi, Jam_Input, t.Id_Transaksi, Sequence, Nominal_Debet desc

-- name: GetDataReportTransaksiJurnal-main
SELECT '1' as fgrouping,
    j.journal_no ,
    j.description,
    j.journal_type,
    j.reference_no, 
    ji.journalitem_no,
    ji.journalitemstatus as status,
    acc.account_code,
    ji.subaccountcode,
    ji.subaccountname,
    acc.account_name,
    acc.account_type,
    j.branch_code,
    j.journal_date,
    ji.branch_code as branch_code_account,
    ji.Valuta_Code,
    ji.Nilai_Kurs,
    ji.ID_JournalBlock,
    ji.amount_debit,
    ji.amount_credit,
    ji.description as ji_description,
    CASE WHEN t.nomor_referensi IS NULL THEN (CASE WHEN ht.nomor_referensi IS NULL THEN '-' ELSE ht.nomor_referensi END) ELSE t.nomor_referensi END AS nomor_referensi,
    CASE WHEN t.jam_input IS NULL THEN ht.jam_input ELSE t.jam_input END AS jam_input,
    CASE WHEN t.user_otorisasi IS NULL THEN (CASE WHEN ht.user_otorisasi IS NULL THEN '-' ELSE ht.user_otorisasi END) ELSE t.user_otorisasi END AS user_otorisasi,
    CASE WHEN t.user_overide IS NULL THEN (CASE WHEN ht.user_overide IS NULL THEN '-' ELSE ht.user_overide END) ELSE t.user_overide END AS user_overide,
    CASE WHEN t.user_input IS NULL THEN (CASE WHEN ht.user_input IS NULL THEN '-' ELSE ht.user_input END) ELSE t.user_input END AS user_input
FROM 
    %[1]v.JOURNALITEM ji, 
    %[1]v.ACCOUNT acc,
    %[1]v.JOURNAL j
    LEFT JOIN %[1]v.TRANSAKSI t on (j.journal_no = t.journal_no)
    LEFT JOIN %[1]v.HISTTRANSAKSI ht on (j.journal_no = ht.journal_no)
WHERE ji.fl_account = acc.account_code
    and ji.fl_journal = j.journal_no
    and j.journal_date >= TO_TIMESTAMP($1,'YYYY-MM-DD')  
    and j.journal_date < TO_TIMESTAMP($2,'YYYY-MM-DD')  
    and ( j.branch_code = $3 or 'ALL' = $4)            
ORDER By
    fgrouping, j.branch_code, j.journal_no, ji.branch_code, amount_debit desc, amount_credit desc 

-- name: GetDataTrxStandingInstruction-main
SELECT
    st.processdateplan,
    st.processdate,
    st.transdescription,
    st.amount,
    si.debitaccountno,
    rt1.nama_rekening as nama_rekening_debet,
    si.creditaccountno,
    rt2.nama_rekening as nama_rekening_kredit,
    st.process_desc,
    st.id_transaksi,
    st.instructionid,
    DECODE(status, 1, 'Belum Diproses', 2, 'Berhasil', 3, 'Gagal') as status,
    1 as data_count
FROM
    %[1]v.scheduledtransaction st,
    %[1]v.standinginstruction si,
    %[1]v.rekeningtransaksi rt1,
    %[1]v.rekeningtransaksi rt2
WHERE
    st.instructionid = si.instructionid
    AND rt1.nomor_rekening = si.debitaccountno
    AND rt2.nomor_rekening = si.creditaccountno
    AND st.amount > 0
    AND st.processdate >= TO_DATE($1 , 'YYYY-MM-DD')
    AND st.processdate < TO_DATE($2 , 'YYYY-MM-DD')
    AND ( si.branchcode = $3 or 'ALL' = $4)
ORDER BY processdate

-- name: GetCTRValue-main
SELECT NILAI_PARAMETER  FROM %v.parameterglobal
WHERE kode_parameter='LIMIT_CTR'

-- name: GetDataReportTplAMLTransKeuanganTunai-main
SELECT mainQ.* FROM 
  (
    SELECT
	to_char(t.tanggal_transaksi, 'DD-MM-YYYY') as tanggal_transaksi,    
      t.kode_cabang_transaksi AS kode_cabang_transaksi,
      n.kode_cabang_input AS kode_cabang,
      n.nomor_nasabah AS nomor_nasabah, 
      n.nama_nasabah AS nama_nasabah, 
      rt.nomor_rekening AS nomor_rekening,	
      k.Kode_Grup_CTR, 
      (case when k.Kode_Grup_CTR = 'ST' then 'SETORAN TUNAI' else 'TARIK TUNAI' end) AS jenis_transaksi, 
      t.keterangan AS berita,
      t.nomor_referensi,
      d.nilai_mutasi,
      i.is_tunai_fisik,
      i.nominal_tunai_fisik,
      1 AS data_count
    FROM
      %[1]v.rekeningtransaksi rt,
      %[1]v.REKENINGLIABILITAS rl,
      %[1]v.nasabah n,
      %[1]v.DETILTRANSAKSI d,
      %[1]v.PARAMETERTRANSAKSIUMUM p,
      %[1]v.KODETRANSTRCASH k,
      %[1]v.TRANSAKSI t
      LEFT JOIN %[1]v.INFOTRANSAKSI i on i.id_transaksi = t.id_transaksi
    WHERE 
      rt.nomor_rekening = rl.nomor_rekening
      AND rl.nomor_nasabah = n.nomor_nasabah
      AND d.nomor_rekening = rl.nomor_rekening
      AND t.id_transaksi = d.id_transaksi
      AND p.kode_transaksi = t.kode_transaksi
      AND k.kode_transaksi = t.kode_transaksi
      AND (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
      AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[4]v')
      AND (d.id_parameter_transaksi not in ('20','30','31') or d.id_parameter_transaksi is null)
      and (t.is_reversed is null or t.is_reversed = 'F')
      and t.status_otorisasi = 1
      and t.kode_transaksi <> 'RV'  
    UNION ALL
  
    SELECT
	to_char(t.tanggal_transaksi, 'DD-MM-YYYY') as tanggal_transaksi,
      t.kode_cabang_transaksi AS kode_cabang_transaksi,
      n.kode_cabang_input AS kode_cabang,
      n.nomor_nasabah AS nomor_nasabah, 
      n.nama_nasabah AS nama_nasabah, 
      rt.nomor_rekening AS nomor_rekening,
      k.Kode_Grup_CTR, 
      (case when k.Kode_Grup_CTR = 'ST' then 'SETORAN TUNAI' else 'TARIK TUNAI' end) AS jenis_transaksi,  
      t.keterangan AS berita,
      t.nomor_referensi,
      d.nilai_mutasi,
      i.is_tunai_fisik,
      i.nominal_tunai_fisik,
      1 AS data_count
    FROM
      %[1]v.rekeningtransaksi rt,
      %[1]v.REKENINGLIABILITAS rl,
      %[1]v.nasabah n,
      %[1]v.HISTDETILTRANSAKSI d,
	  %[1]v.PARAMETERTRANSAKSIUMUM p,
      %[1]v.KODETRANSTRCASH k,
      %[1]v.HISTTRANSAKSI t
      LEFT JOIN %[1]v.INFOTRANSAKSI i on i.id_transaksi = t.id_transaksi
    WHERE 
      rt.nomor_rekening = rl.nomor_rekening
      AND rl.nomor_nasabah = n.nomor_nasabah
      AND d.nomor_rekening = rl.nomor_rekening
      AND t.id_transaksi = d.id_transaksi
      AND p.kode_transaksi = t.kode_transaksi
      AND k.kode_transaksi = t.kode_transaksi
      AND (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
      AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[4]v')
      AND (d.id_parameter_transaksi not in ('20','30','31') or d.id_parameter_transaksi is null)
      and (t.is_reversed is null or t.is_reversed = 'F')
      and t.status_otorisasi = 1
      and t.kode_transaksi <> 'RV'
  ) mainQ
  WHERE 
  EXISTS (
    SELECT 1 FROM (
      SELECT 
	    dtAkumCIF.tanggal_transaksi,         
        dtAkumCIF.nomor_nasabah,
        dtAkumCIF.kode_cabang_nasabah,
        dtAkumCIF.kode_grup_ctr,
        SUM(dtAkumCIF.nilai_mutasi) AS nilai_mutasi
      FROM (
        SELECT
          d.id_detil_transaksi,
		  to_char(t.tanggal_transaksi, 'DD-MM-YYYY') as tanggal_transaksi, 
          n.nomor_nasabah,
          n.kode_cabang_input as kode_cabang_nasabah,
          k.kode_grup_ctr, 
          d.nilai_mutasi
        FROM
          %[1]v.rekeningtransaksi rt,
          %[1]v.rekeningliabilitas rl,
          %[1]v.nasabah n,
          %[1]v.transaksi t,
          %[1]v.detiltransaksi d,
          %[1]v.parametertransaksiumum p,
          %[1]v.KodeTranSTRCash k
        WHERE 
          rt.nomor_rekening = rl.nomor_rekening
          AND rl.nomor_nasabah = n.nomor_nasabah
          AND d.nomor_rekening = rl.nomor_rekening
          AND t.id_transaksi = d.id_transaksi
          AND t.kode_transaksi = p.kode_transaksi
          AND (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
          AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[4]v')
          AND (d.id_parameter_transaksi not in ('20','30','31') or d.id_parameter_transaksi is null)                  
          AND k.kode_transaksi = t.kode_transaksi
          and (t.is_reversed is null or t.is_reversed = 'F')
          and t.status_otorisasi = 1
          and t.kode_transaksi <> 'RV'
        UNION ALL
      
        SELECT
          d.id_detil_transaksi,
		  to_char(t.tanggal_transaksi, 'DD-MM-YYYY') as tanggal_transaksi,       
          n.nomor_nasabah,
          n.kode_cabang_input as kode_cabang_nasabah,
          k.kode_grup_ctr, 
          d.nilai_mutasi
        FROM
          %[1]v.rekeningtransaksi rt,
          %[1]v.rekeningliabilitas rl,
          %[1]v.nasabah n,
          %[1]v.histtransaksi t,
          %[1]v.histdetiltransaksi d,
          %[1]v.parametertransaksiumum p,
          %[1]v.KodeTranSTRCash k
        WHERE 
          rt.nomor_rekening = rl.nomor_rekening
          AND rl.nomor_nasabah = n.nomor_nasabah
          AND d.nomor_rekening = rl.nomor_rekening
          AND t.id_transaksi = d.id_transaksi
          AND t.kode_transaksi = p.kode_transaksi
          AND (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
          AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[4]v')
          AND (d.id_parameter_transaksi not in ('20','30','31') or d.id_parameter_transaksi is null)                  
          AND k.kode_transaksi = t.kode_transaksi
          and (t.is_reversed is null or t.is_reversed = 'F')
          and t.status_otorisasi = 1
          and t.kode_transaksi != 'RV'
      ) dtAkumCIF
      HAVING SUM(dtAkumCIF.nilai_mutasi) >= %[2]v
      GROUP BY
        dtAkumCIF.tanggal_transaksi,
        dtAkumCIF.nomor_nasabah,
        dtAkumCIF.kode_cabang_nasabah,
        dtAkumCIF.Kode_Grup_CTR
    ) dtExistsAkum
    WHERE
      dtExistsAkum.nomor_nasabah = mainQ.nomor_nasabah
      and dtExistsAkum.tanggal_transaksi = mainQ.tanggal_transaksi 
      and (dtExistsAkum.kode_cabang_nasabah = $3 OR 'ALL' = '%[5]v')
      and dtExistsAkum.Kode_Grup_CTR = mainQ.Kode_Grup_CTR
  )
    
  ORDER BY tanggal_transaksi ASC, nomor_nasabah ASC, jenis_transaksi

-- name: GetDataReportAMLTotalCIF-main
SELECT
    coalesce(kode_jenis, 'NOREK') AS kode_jenis,
    COUNT(DISTINCT a.nomor_nasabah) jumlah
FROM
    %[1]v.nasabah a
LEFT JOIN %[1]v.rekeningliabilitas b ON
    (a.nomor_nasabah = b.nomor_nasabah)
LEFT JOIN %[1]v.rekeningtransaksi c ON
    (b.nomor_rekening = c.nomor_rekening)
WHERE
    (a.kode_cabang_input = $1 OR 'ALL' = $2)
    AND a.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
    AND a.is_wic = 'F'
    AND EXISTS(
    SELECT
        1
    FROM
    %[1]v.rekeningtransaksi rt,
    %[1]v.rekeningliabilitas rl
    WHERE
        rt.nomor_rekening = rl.nomor_rekening
        AND a.nomor_nasabah = rl.nomor_nasabah
        AND rl.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
        AND ( rl.tanggal_tutup >= to_date ($3, 'yyyy-mm-dd')
            OR rl.tanggal_tutup IS NULL )
                    )
    AND NOT EXISTS (
    SELECT
        1
    FROM
    %[1]v.parameterglobal
    WHERE
        nilai_parameter_string = a.nomor_nasabah
        AND kode_parameter = 'COMPANYCIF'
                )
GROUP BY
    c.kode_jenis
UNION
            SELECT
    'TOTAL',
    COUNT(DISTINCT nomor_nasabah) jumlah
FROM
%[1]v.nasabah a
WHERE
    (a.kode_cabang_input = $1
        OR 'ALL' = $2)
    AND a.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
    AND a.is_wic = 'F'
    AND NOT EXISTS (
    SELECT
        1
    FROM
    %[1]v.parameterglobal
    WHERE
        nilai_parameter_string = a.nomor_nasabah
        AND kode_parameter = 'COMPANYCIF'
                )
UNION
            SELECT
    'DEP ONLY',
    COUNT(DISTINCT a.nomor_nasabah) jumlah
FROM
    nasabah a
WHERE
    (a.kode_cabang_input = $1 OR 'ALL' = $2)
    AND a.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
    AND a.is_wic = 'F'
    AND EXISTS(
    SELECT
        1
    FROM
        rekeningtransaksi rt,
        rekeningliabilitas rl
    WHERE
        rt.nomor_rekening = rl.nomor_rekening
        AND a.nomor_nasabah = rl.nomor_nasabah
        AND rl.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
        AND ( rl.tanggal_tutup >= to_date ($3, 'yyyy-mm-dd')
            OR rl.tanggal_tutup IS NULL )
        AND rt.kode_jenis = 'DEP'
                    )
    AND NOT EXISTS (
    SELECT
        1
    FROM
        parameterglobal
    WHERE
        nilai_parameter_string = a.nomor_nasabah
        AND kode_parameter = 'COMPANYCIF'
                )
UNION
            SELECT
    'BLM OTOR',
    COUNT(DISTINCT nomor_nasabah) jumlah
FROM
    nasabah a
WHERE
    (a.kode_cabang_input = $1
        OR 'ALL' =$2)
    AND a.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
    AND a.is_wic = 'F'
    AND a.is_otorisasi <> 'T'
    AND NOT EXISTS (
    SELECT
        1
    FROM
        parameterglobal
    WHERE
        nilai_parameter_string = a.nomor_nasabah
        AND kode_parameter = 'COMPANYCIF'
                )
UNION
            SELECT
    'GABUNGAN',
    COUNT(DISTINCT a.nomor_nasabah) jumlah
FROM
    nasabah a
WHERE
    (a.kode_cabang_input = $1
        OR 'ALL' = $2)
    AND a.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
    AND a.is_wic = 'F'
    AND EXISTS(
    SELECT
        1
    FROM
        rekeningtransaksi rt,
        rekeningliabilitas rl
    WHERE
        rt.nomor_rekening = rl.nomor_rekening
        AND a.nomor_nasabah = rl.nomor_nasabah
        AND rl.tanggal_buka < to_date ($3, 'yyyy-mm-dd')
        AND ( rl.tanggal_tutup >= to_date ($3, 'yyyy-mm-dd')
            OR rl.tanggal_tutup IS NULL )
                    )
    AND NOT EXISTS (
    SELECT
        1
    FROM
        parameterglobal
    WHERE
        nilai_parameter_string = a.nomor_nasabah
        AND kode_parameter = 'COMPANYCIF'
                ) 

-- name: GetDataTrxCabRTGSTransferIn-main
SELECT
    to_char(t.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,             
    rec.amount,
    rec.from_member,
    m.nama_member,
    concat(concat(rec.from_member, ' - '), m.nama_member) as bank_pengirim,
    rec.branch_code AS cabang,
    st.originating_name,
    st.ultimate_benef_account,
    st.ultimate_benef_name,
    st.payment_details,
    rt.nama_rekening,
    1 as data_count
FROM
    %[2]v.rttransaction rec
    LEFT JOIN %[2]v.rtsession h on h.session_id = rec.session_id
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r on r.trx_record_id = rec.trans_id
    LEFT JOIN %[1]v.DETILTRANSAKSI d on r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[1]v.TRANSAKSI t on t.id_transaksi = d.id_transaksi
    LEFT JOIN %[2]v.rtsingletransfer st on st.transdetail_id = rec.transdetail_id
    LEFT JOIN %[1]v.MemberRTGS m on m.kode_member = rec.from_member
    LEFT JOIN %[1]v.RekeningTransaksi rt on rt.nomor_rekening = d.nomor_rekening
WHERE
    (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= $2)
    AND (t.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'= $4)
    AND h.hdr_message_type='O'
    AND t.kode_transaksi = 'RTINC'
    AND rec.branch_code = $5
UNION
SELECT
    to_char(t.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
    rec.amount,
    rec.from_member,
    m.nama_member,
    concat(concat(rec.from_member, ' - '), m.nama_member) as bank_pengirim,
    rec.branch_code AS cabang,
    st.originating_name,
    st.ultimate_benef_account,
    st.ultimate_benef_name,
    st.payment_details,
    rt.nama_rekening,
    1 as data_count
FROM
    %[2]v.rttransaction rec
    LEFT JOIN %[2]v.rtsession h on h.session_id = rec.session_id
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r on r.trx_record_id = rec.trans_id
    LEFT JOIN %[1]v.HistDetilTransaksi d on r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[1]v.HistTransaksi t on t.id_transaksi = d.id_transaksi
    LEFT JOIN %[2]v.rtsingletransfer st on st.transdetail_id = rec.transdetail_id
    LEFT JOIN %[1]v.MemberRTGS m on m.kode_member = rec.from_member
    LEFT JOIN %[1]v.RekeningTransaksi rt on rt.nomor_rekening = d.nomor_rekening
    WHERE
        (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= $2)
        AND (t.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'= $4)
        AND h.hdr_message_type='O'
        AND t.kode_transaksi = 'RTINC'
        AND rec.branch_code = $5

-- name: GetDataReportAMLNasabahOverLimitTrans-main
SELECT
    limitData.kode_cabang AS kode_cabang,
    limitData.nomor_nasabah AS nomor_nasabah,
    limitData.nama_nasabah AS nama_nasabah,
    limitData.j_nasabah AS j_nasabah,
    limitData.tunai_nominal_debet AS tunai_nominal_debet,
    limitData.tunai_frek_debet AS tunai_frek_debet,
    limitData.tunai_nominal_kredit AS tunai_nominal_kredit,
    limitData.tunai_frek_kredit AS tunai_frek_kredit,
    limitData.nontunai_nominal_debet AS nontunai_nominal_debet,
    limitData.nontunai_frek_debet AS nontunai_frek_debet,
    limitData.nontunai_nominal_kredit AS nontunai_nominal_kredit,
    limitData.nontunai_frek_kredit AS nontunai_frek_kredit,
    trx.tunai_nominal_debet AS tunai_nominal_debet1,
    trx.tunai_frek_debet AS tunai_frek_debet1,
    trx.tunai_nominal_kredit AS tunai_nominal_kredit1,
    trx.tunai_frek_kredit AS tunai_frek_kredit1,
    trx.nontunai_nominal_debet AS nontunai_nominal_debet1,
    trx.nontunai_frek_debet AS nontunai_frek_debet1,
    trx.nontunai_nominal_kredit AS nontunai_nominal_kredit1,
    trx.nontunai_frek_kredit AS nontunai_frek_kredit1,
    trx.bulan,
    trx.tahun,
    trx.periode
FROM
    (
    SELECT
        n.kode_cabang_input AS kode_cabang,
        n.nomor_nasabah AS nomor_nasabah,
        n.nama_nasabah AS nama_nasabah,
        t.periode,
        t.bulan,
        t.tahun,
        (CASE
            n.jenis_nasabah WHEN 'I' THEN 'Individu'
            ELSE 'Korporat'
        END) AS j_nasabah,
        SUM(CASE
            WHEN t.jenis_mutasi = 'D' AND t.is_cash = 'T'
            THEN t.nilai_mutasi ELSE 0.0 END) AS tunai_nominal_debet,
        SUM(CASE
            WHEN t.jenis_mutasi = 'D' AND t.is_cash = 'T'
            THEN 1 ELSE 0 END) AS tunai_frek_debet,
        SUM(CASE
            WHEN t.jenis_mutasi = 'C' AND t.is_cash = 'T'
            THEN t.nilai_mutasi ELSE 0.0 END) AS tunai_nominal_kredit,
        SUM(CASE
            WHEN t.jenis_mutasi = 'C' AND t.is_cash = 'T'
            THEN 1 ELSE 0 END) AS tunai_frek_kredit,
        SUM(CASE
            WHEN t.jenis_mutasi = 'D' AND t.is_cash = 'F'
            THEN t.nilai_mutasi ELSE 0.0 END) AS nontunai_nominal_debet,
        SUM(CASE
            WHEN t.jenis_mutasi = 'D' AND t.is_cash = 'F'
            THEN 1 ELSE 0 END) AS nontunai_frek_debet,
        SUM(CASE
            WHEN t.jenis_mutasi = 'C' AND t.is_cash = 'F'
            THEN t.nilai_mutasi ELSE 0.0 END) AS nontunai_nominal_kredit,
        SUM(CASE
            WHEN t.jenis_mutasi = 'C' AND t.is_cash = 'F'
            THEN 1 ELSE 0 END) AS nontunai_frek_kredit
    FROM
        %[1]v.rekeningtransaksi rt,
        %[1]v.rekeningliabilitas rl,
        %[1]v.nasabah n,
        (
        SELECT
            t.id_transaksi
                    ,
            t.kode_transaksi
                    ,
            EXTRACT(MONTH FROM t.tanggal_transaksi) AS bulan
                    ,
            EXTRACT(YEAR FROM t.tanggal_transaksi) AS tahun
                    ,
            TO_CHAR(t.tanggal_transaksi, 'MM-YYYY') AS periode                    
                    ,
            d.nomor_rekening ,
            d.jenis_mutasi,
            d.nilai_mutasi
                    ,
            ( CASE
                WHEN k.kode_transaksi IS NOT NULL THEN 'T'
                ELSE 'F'
            END) AS is_cash
        FROM
        %[1]v.detiltransaksi d
                ,
                %[1]v.transaksi t
        LEFT JOIN KodeTranSTRCash k
                    ON
            (k.kode_transaksi = t.kode_transaksi)
        WHERE
            t.id_transaksi = d.id_transaksi
            AND d.kode_tx_class IS NULL
            AND t.tanggal_transaksi >= to_date($1, 'yyyy-mm-dd')
            AND t.tanggal_transaksi < to_date($2, 'yyyy-mm-dd')
            AND (d.id_parameter_transaksi <> '20'
                OR d.id_parameter_transaksi IS NULL)
            AND NOT EXISTS (
            SELECT
                1
            FROM
                KodeTranSTRExclude
            WHERE
                kode_transaksi = t.kode_transaksi )
            AND (t.is_reversed IS NULL
                OR t.is_reversed = 'F')
            AND t.status_otorisasi = 1
            AND t.kode_transaksi <> 'RV'
    UNION ALL
        SELECT
            t.id_transaksi
                    ,
            t.kode_transaksi
                    ,
            EXTRACT(MONTH FROM t.tanggal_transaksi) AS bulan
                    ,
            EXTRACT(YEAR FROM t.tanggal_transaksi) AS tahun
                    ,
            TO_CHAR(t.tanggal_transaksi, 'MM-YYYY') AS periode
                    ,
            d.nomor_rekening ,
            d.jenis_mutasi,
            d.nilai_mutasi
                    ,
            ( CASE
                WHEN k.kode_transaksi IS NOT NULL THEN 'T'
                ELSE 'F'
            END) AS is_cash
        FROM
        %[1]v.histdetiltransaksi d
                ,
                %[1]v.histtransaksi t
        LEFT JOIN KodeTranSTRCash k
                    ON
            (k.kode_transaksi = t.kode_transaksi)
        WHERE
            t.id_transaksi = d.id_transaksi
            AND d.kode_tx_class IS NULL
            AND t.tanggal_transaksi >= to_date($1, 'yyyy-mm-dd')
            AND t.tanggal_transaksi < to_date($2, 'yyyy-mm-dd')
            AND (d.id_parameter_transaksi <> '20'
                OR d.id_parameter_transaksi IS NULL)
            AND NOT EXISTS (
            SELECT
                1
            FROM
            %[1]v.KodeTranSTRExclude
            WHERE
                kode_transaksi = t.kode_transaksi )
            AND (t.is_reversed IS NULL
                OR t.is_reversed = 'F')
            AND t.status_otorisasi = 1
            AND t.kode_transaksi <> 'RV') t
    WHERE
        rt.nomor_rekening = rl.nomor_rekening
        AND rl.nomor_nasabah = n.nomor_nasabah
        AND t.nomor_rekening = rl.nomor_rekening
        AND (n.kode_cabang_input = $3
            OR 'ALL' = $4)
        AND rl.jenis_rekening_liabilitas IN ('T', 'G')
        AND NOT EXISTS (
        SELECT
            1
        FROM
        %[1]v.parameterglobal
        WHERE
            nilai_parameter_string = n.nomor_nasabah
            AND kode_parameter = 'COMPANYCIF'
            )
    GROUP BY
        n.kode_cabang_input,
        n.nomor_nasabah,
        n.nama_nasabah,
        periode,
        t.bulan,
        t.tahun,
        (CASE
            n.jenis_nasabah WHEN 'I' THEN 'Individu'
            ELSE 'Korporat'
        END)
    ) trx,
    (
    SELECT
        n.kode_cabang_input AS kode_cabang,
        n.nomor_nasabah AS nomor_nasabah,
        n.nama_nasabah AS nama_nasabah,
        (CASE
            n.jenis_nasabah WHEN 'I' THEN 'Individu'
            ELSE 'Korporat'
        END) AS j_nasabah,
        n.limit_nom_tarik_tunai AS tunai_nominal_debet,
        n.limit_frek_tarik_tunai AS tunai_frek_debet,
        n.limit_nom_setor_tunai AS tunai_nominal_kredit,
        n.limit_frek_setor_tunai AS tunai_frek_kredit,
        n.limit_nom_tarik_nontunai AS nontunai_nominal_debet,
        n.limit_frek_tarik_nontunai AS nontunai_frek_debet,
        n.limit_nom_setor_nontunai AS nontunai_nominal_kredit,
        n.limit_frek_setor_nontunai AS nontunai_frek_kredit
    FROM
    %[1]v.nasabah n
    ) limitData
WHERE
    trx.nomor_nasabah = limitData.nomor_nasabah
    AND (
        trx.tunai_nominal_debet > limitData.tunai_nominal_debet
        OR trx.tunai_frek_debet > limitData.tunai_frek_debet
        OR trx.tunai_nominal_kredit > limitData.tunai_nominal_kredit
        OR trx.tunai_frek_kredit > limitData.tunai_frek_kredit
        OR trx.nontunai_nominal_debet > limitData.nontunai_nominal_debet
        OR trx.nontunai_frek_debet > limitData.nontunai_frek_debet
        OR trx.nontunai_nominal_kredit > limitData.nontunai_nominal_kredit
        OR trx.nontunai_frek_kredit > limitData.nontunai_frek_kredit
        )
ORDER BY
    periode,
    kode_cabang,
    nomor_nasabah

-- name: GetDataReportTplAMLTransKeuanganTunaiSummary-main
SELECT qListCTR.tanggal_transaksi
, qListCTR.nomor_nasabah
, qListCTR.kode_transaksi
, qListCTR.nilai_mutasi
, n.kode_cabang_input as kode_cabang
, n.nama_nasabah 
FROM
%[1]v.NASABAH n
, ( SELECT 
    to_char(dtAkumCIF.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
    dtAkumCIF.nomor_nasabah,
    dtAkumCIF.kode_transaksi, 
    SUM(dtAkumCIF.nilai_mutasi) AS nilai_mutasi
FROM (
    SELECT
        t.tanggal_transaksi AS tanggal_transaksi,
        n.nomor_nasabah AS nomor_nasabah,
        (case when k.Kode_Grup_CTR = 'ST' then 'SETORAN TUNAI' else 'TARIK TUNAI' end) AS kode_transaksi,
        d.nilai_mutasi
    FROM
        %[1]v.rekeningtransaksi rt,
        %[1]v.rekeningliabilitas rl,
        %[1]v.nasabah n,
        %[1]v.transaksi t,
        %[1]v.detiltransaksi d,
        %[1]v.parametertransaksiumum p,
        %[1]v.KodeTranSTRCash k
    WHERE 
        rt.nomor_rekening = rl.nomor_rekening
        AND rl.nomor_nasabah = n.nomor_nasabah
        AND d.nomor_rekening = rl.nomor_rekening
        AND t.id_transaksi = d.id_transaksi
        AND p.kode_transaksi = t.kode_transaksi                  
        AND (d.id_parameter_transaksi not in ('20','30','31') or d.id_parameter_transaksi is null)
        AND (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
        AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[4]v')
        AND k.kode_transaksi = t.kode_transaksi
        and (t.is_reversed is null or t.is_reversed = 'F')
        and t.status_otorisasi = 1
        and t.kode_transaksi <> 'RV'
    UNION ALL
    
    SELECT
        t.tanggal_transaksi AS tanggal_transaksi,
        n.nomor_nasabah AS nomor_nasabah,                  
        (case when k.Kode_Grup_CTR = 'ST' then 'SETORAN TUNAI' else 'TARIK TUNAI' end) AS kode_transaksi,
        d.nilai_mutasi
    FROM
    %[1]v.rekeningtransaksi rt,
    %[1]v.rekeningliabilitas rl,
    %[1]v.nasabah n,
    %[1]v.histtransaksi t,
    %[1]v.histdetiltransaksi d,
    %[1]v.parametertransaksiumum p,
    %[1]v.KodeTranSTRCash k
    WHERE 
        rt.nomor_rekening = rl.nomor_rekening
        AND rl.nomor_nasabah = n.nomor_nasabah
        AND d.nomor_rekening = rl.nomor_rekening
        AND t.id_transaksi = d.id_transaksi
        AND p.kode_transaksi = t.kode_transaksi                  
        AND (d.id_parameter_transaksi not in ('20','30','31') or d.id_parameter_transaksi is null)
        AND (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
        AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[4]v')
        AND k.kode_transaksi = t.kode_transaksi
        and (t.is_reversed is null or t.is_reversed = 'F')
        and t.status_otorisasi = 1
        and t.kode_transaksi <> 'RV'
) dtAkumCIF
HAVING SUM(dtAkumCIF.nilai_mutasi) >= %[2]v
GROUP BY
dtAkumCIF.tanggal_transaksi,
dtAkumCIF.nomor_nasabah,
dtAkumCIF.kode_transaksi
) qListCTR
WHERE 
qListCTR.NOMOR_NASABAH = n.NOMOR_NASABAH
AND (n.kode_cabang_input = $3 OR 'ALL'= '%[4]v')
ORDER BY qListCTR.tanggal_transaksi ASC, qListCTR.nomor_nasabah ASC, n.kode_cabang_input ASC

-- name: GetDataTrxCabRTGSTransferOut-main
SELECT
    detran.nilai_mutasi,
    h.session_userid,
    h.session_time,
    trf.kode_bank,
    trf.Nomor_Rekening_Penerima,
    trf.Nama_Penerima,
    trf.Nama_Pengirim,
    m.nama_member,
    tran.kode_cabang_transaksi,
    tran.keterangan,
    tran.user_input,
    tran.user_overide,
    tran.user_otorisasi,
    tran.jam_input,
    tran.jam_otorisasi,
    tran.nomor_referensi,
    tran.status_otorisasi,
    to_char(tran.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
    tran.id_transaksi,
    tran.is_reversed,
    infot.nomor_rekening_debet,
    infot.nominal_biaya,
    enum.enum_description,
    1 AS data_count,
    case when tran.is_reversed = 'T' then concat(enum.enum_description,'(REV)') else enum.enum_description end description
FROM
    %[1]v.Transaksi tran
    INNER JOIN %[1]v.detiltransaksi detran ON tran.id_transaksi = detran.id_transaksi
    INNER JOIN %[1]v.Transfer trf ON trf.id_detil_transaksi = detran.id_detil_transaksi
    LEFT JOIN %[1]v.MemberRTGS m  ON m.kode_member = trf.kode_bank
    LEFT JOIN %[1]v.InfoTransaksi infot ON infot.id_transaksi = tran.id_transaksi
    LEFT JOIN %[1]v.RemittanceTransaction remtran ON detran.id_detil_transaksi = remtran.id_detil_transaksi
    LEFT JOIN %[2]v.rttransaction rec ON remtran.trx_record_id = rec.trans_id
    LEFT JOIN %[2]v.rttransferdetail recdet ON rec.transdetail_id = recdet.transdetail_id
    LEFT JOIN %[2]v.RTSingleTransfer strans ON strans.transdetail_id = recdet.transdetail_id
    LEFT JOIN %[2]v.rtsession h ON h.session_id = rec.session_id
    LEFT JOIN %[1]v.enum_int enum ON tran.status_otorisasi = enum.enum_value AND enum.enum_name = 'eStatusOtorisasiTransaksi'
WHERE
    (tran.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL' = $2)
    AND (tran.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') + 1 OR 'ALL' = $4)
    AND (tran.kode_cabang_transaksi = $5 OR 'ALL' = $6)
    AND tran.kode_transaksi = 'RTGS'
    AND detran.Is_Create_Remittance = 'T'
UNION
SELECT
    detran.nilai_mutasi,
    h.session_userid,
    h.session_time,
    trf.kode_bank,
    trf.Nomor_Rekening_Penerima,
    trf.Nama_Penerima,
    trf.Nama_Pengirim,
    m.nama_member,
    tran.kode_cabang_transaksi,
    tran.keterangan,
    tran.user_input,
    tran.user_overide,
    tran.user_otorisasi,
    tran.jam_input,
    tran.jam_otorisasi,
    tran.nomor_referensi,
    tran.status_otorisasi,
    to_char(tran.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
    tran.id_transaksi,
    tran.is_reversed,
    infot.nomor_rekening_debet,
    infot.nominal_biaya,
    enum.enum_description,
    1 AS data_count,
    case when tran.is_reversed = 'T' then concat(enum.enum_description,'(REV)') else enum.enum_description end description
FROM
    %[1]v.HistTransaksi tran
    INNER JOIN %[1]v.HistDetilTransaksi detran ON tran.id_transaksi = detran.id_transaksi
    INNER JOIN %[1]v.Transfer trf ON trf.id_detil_transaksi = detran.id_detil_transaksi
    LEFT JOIN %[1]v.MemberRTGS m  ON m.kode_member = trf.kode_bank
    LEFT JOIN %[1]v.InfoTransaksi infot ON infot.id_transaksi = tran.id_transaksi
    LEFT JOIN %[1]v.RemittanceTransaction remtran ON detran.id_detil_transaksi = remtran.id_detil_transaksi
    LEFT JOIN %[2]v.rttransaction rec ON remtran.trx_record_id = rec.trans_id
    LEFT JOIN %[2]v.rttransferdetail recdet ON rec.transdetail_id = recdet.transdetail_id
    LEFT JOIN %[2]v.RTSingleTransfer strans ON strans.transdetail_id = recdet.transdetail_id
    LEFT JOIN %[2]v.rtsession h ON h.session_id = rec.session_id
    LEFT JOIN %[1]v.enum_int enum ON tran.status_otorisasi = enum.enum_value AND enum.enum_name = 'eStatusOtorisasiTransaksi'
WHERE
    (tran.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL' = $2)
    AND (tran.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') + 1 OR 'ALL' = $4)
    AND (tran.kode_cabang_transaksi = $5 OR 'ALL' = $6)
    AND tran.kode_transaksi = 'RTGS'
    AND detran.Is_Create_Remittance = 'T'
ORDER BY tanggal_transaksi desc, kode_cabang_transaksi, id_transaksi

-- name: GetDataReportSKNTransferKeluar-main
SELECT 
    ts.session_user_id, 
    cr.reff_no,     
    trf.kode_bank,
    trf.Nomor_Rekening_Penerima,
    trf.Nama_Penerima,
    trf.Nama_Pengirim,     
    d.nilai_mutasi,
    TO_CHAR(ts.session_time, 'HH24::MI::SS') session_time,
    t.nomor_referensi,    
    to_char(t.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
    t.user_input,    
    t.user_otorisasi, 
    TO_CHAR(T.jam_otorisasi, 'HH24::MI::SS') jam_otorisasi,
    TO_CHAR(T.jam_input, 'HH24::MI::SS') jam_input,  
    t.kode_cabang_transaksi,
    t.keterangan,
    t.id_transaksi,    
    t.is_reversed,    
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    t.user_overide,
    infot.nomor_rekening_debet,
    infot.nominal_biaya,
    enum.enum_description,
    1 AS data_count
FROM                                                         
    %[1]v.TRANSAKSI t 
    INNER JOIN %[1]v.DETILTRANSAKSI d ON t.id_transaksi = d.id_transaksi 
    INNER JOIN %[1]v.TRANSFER trf ON trf.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on trf.kode_bank = b.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m on trf.kode_bank = m.kode_member
    LEFT JOIN %[1]v.INFOTRANSAKSI infot ON infot.id_transaksi = t.id_transaksi        
    LEFT JOIN %[1]v.REMITTANCETRANSACTION r ON r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[2]v.TRX_RECORD  rec ON r.trx_record_id = rec.trx_record_id
    LEFT JOIN %[2]v.DOI_CREDIT  cr ON rec.trx_record_id = cr.trx_record_id    
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id 
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id 
    LEFT JOIN %[1]v.ENUM_INT enum ON t.status_otorisasi = enum.enum_value AND enum.enum_name = 'eStatusOtorisasiTransaksi' 
WHERE
    t.kode_transaksi = 'SKN'
    AND (t.kode_cabang_transaksi = $1 OR 'ALL'= $2)
    AND (t.tanggal_transaksi >= TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'= $4)
    AND (t.tanggal_transaksi < TO_DATE($5, 'YYYY-MM-DD') + 1 OR 'ALL'= $6)
    AND d.Is_Create_Remittance = 'T'  
UNION  
SELECT
    ts.session_user_id, 
    cr.reff_no,     
    trf.kode_bank,
    trf.Nomor_Rekening_Penerima,
    trf.Nama_Penerima,     
    trf.Nama_Pengirim,
    d.nilai_mutasi,
    TO_CHAR(ts.session_time, 'HH24::MI::SS') session_time,
    t.nomor_referensi,    
    to_char(t.tanggal_transaksi, 'DD-MM-YYYY') tanggal_transaksi,
    t.user_input,    
    t.user_otorisasi, 
    TO_CHAR(T.jam_otorisasi, 'HH24::MI::SS') jam_otorisasi,
    TO_CHAR(T.jam_input, 'HH24::MI::SS') jam_input,    
    t.kode_cabang_transaksi,
    t.keterangan ,
    t.id_transaksi,  
    t.is_reversed,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    t.user_overide,
    infot.nomor_rekening_debet,
    infot.nominal_biaya,
    enum.enum_description,  
    1 AS data_count
FROM
    %[1]v.HISTTRANSAKSI  t 
    INNER JOIN %[1]v.HISTDETILTRANSAKSI d ON t.id_transaksi = d.id_transaksi 
    INNER JOIN %[1]v.TRANSFER trf ON trf.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on trf.kode_bank = b.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m on trf.kode_bank = m.kode_member
    LEFT JOIN %[1]v.INFOTRANSAKSI infot ON infot.id_transaksi = t.id_transaksi        
    LEFT JOIN %[1]v.REMITTANCETRANSACTION  r ON r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[2]v.TRX_RECORD  rec ON r.trx_record_id = rec.trx_record_id
    LEFT JOIN %[2]v.DOI_CREDIT cr ON rec.trx_record_id = cr.trx_record_id    
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id 
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id 
    LEFT JOIN %[1]v.ENUM_INT enum ON t.status_otorisasi = enum.enum_value AND enum.enum_name = 'eStatusOtorisasiTransaksi'       
WHERE
    t.kode_transaksi = 'SKN'
    AND (t.kode_cabang_transaksi = $1 OR 'ALL'= $2)
    AND (t.tanggal_transaksi >= TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'= $4)
    AND (t.tanggal_transaksi < TO_DATE($5, 'YYYY-MM-DD') + 1 OR 'ALL'= $6)
    AND d.Is_Create_Remittance = 'T'  
ORDER BY tanggal_transaksi, kode_cabang_transaksi, id_transaksi

-- name: GetDataReportSKNTransferMasuk-main
SELECT 
    h.batch_no,
    ts.session_user_id, 
    t.tanggal_transaksi, 
    t.nomor_referensi, 
    cr.ori_kli_code AS kode_bank,
    cr.descr,
    d.nilai_mutasi, 
    t.user_otorisasi, 
    t.jam_input, 
    ts.session_time, 
    rec.branch_code,
    cr.recp_account,
    cr.recp_name,
    cr.sender_name,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    1 AS data_count
FROM
    %[2]v.DIE_CREDIT cr
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = cr.trx_record_id  
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id  
    LEFT JOIN %[1]v.REMITTANCETRANSACTION r ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[1]v.DETILTRANSAKSI d ON r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.TRANSAKSI t ON t.id_transaksi = d.id_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on cr.ori_kli_code = b.kode_bank    
    LEFT JOIN %[1]v.MEMBERRTGS m on cr.ori_kli_code = m.kode_member
WHERE
    (rec.branch_code = $1 OR 'ALL'= $2)
    AND (t.tanggal_transaksi >= TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'= $4)
    AND (t.tanggal_transaksi < TO_DATE($5, 'YYYY-MM-DD') OR 'ALL'= $6)
    AND t.kode_transaksi = 'INSKN'

UNION

SELECT
    h.batch_no,
    ts.session_user_id,     
    t.tanggal_transaksi,
    t.nomor_referensi, 
    cr.ori_kli_code AS kode_bank,
    cr.descr,     
    d.nilai_mutasi,
    t.user_otorisasi, 
    t.jam_input, 
    ts.session_time, 
    rec.branch_code,
    cr.recp_account,
    cr.recp_name,
    cr.sender_name,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    1 AS data_count
FROM
    %[2]v.DIE_CREDIT cr 
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = cr.trx_record_id  
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id    
    LEFT JOIN %[1]v.REMITTANCETRANSACTION r ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[1]v.HISTDETILTRANSAKSI d ON r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.HISTTRANSAKSI t ON t.id_transaksi = d.id_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on cr.ori_kli_code = b.kode_bank        
    LEFT JOIN %[1]v.MEMBERRTGS m on cr.ori_kli_code = m.kode_member   
WHERE
    (rec.branch_code = $1 OR 'ALL'= $2)
    AND (t.tanggal_transaksi >= TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'= $4)
    AND (t.tanggal_transaksi < TO_DATE($5, 'YYYY-MM-DD') OR 'ALL'= $6)
    AND t.kode_transaksi = 'INSKN'
    
ORDER BY branch_code, tanggal_transaksi

-- name: GetDataReportTplTransaksiCabSKNTarikanKliring-main
SELECT 
    h.batch_no, 
    ts.session_date,  
	db.ori_kli_code || ' - ' || coalesce(b.Nama_Lengkap, m.nama_member) as sandi_bank,
    db.ori_kli_code AS kode_bank,
    db.serial_no,    
    db.trx_amount, 
    db.account_no,
    db.trx_code,
	CASE
	WHEN db.trx_code >= '00' AND db.trx_code <= '09' THEN 'Cek'
	WHEN db.trx_code >= '10' AND db.trx_code <= '19' THEN 'Bilyet Giro'
	WHEN db.trx_code >= '20' AND db.trx_code <= '29' THEN 'Wesel'
	WHEN db.trx_code >= '30' AND db.trx_code <= '39' THEN 'Warkat Lainnya'
	WHEN db.trx_code >= '40' AND db.trx_code <= '49' THEN 'Nota Debet'
	ELSE NULL
	END AS ket_warkat,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    rt.nama_rekening,
    ts.session_time, 
    1 AS data_count
  FROM
    %[2]v.DIE_DEBET db
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = db.trx_record_id
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id
    LEFT JOIN %[1]v.BANKKLIRING b ON b.kode_bank = db.ori_kli_code    
    LEFT JOIN %[1]v.MEMBERRTGS m on m.kode_member = db.ori_kli_code
    LEFT JOIN %[1]v.REKENINGTRANSAKSI rt ON rt.nomor_rekening = db.account_no
  WHERE
    (ts.session_date >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
    AND (ts.session_date < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND (rec.branch_code = $3 OR 'ALL'= '%[5]v')
    AND rec.trx_auth_status = 'A'  
  ORDER BY session_date       

-- name: GetDataReportTplTransaksiCabSKNTransferRetur-main
SELECT 
    h.batch_no,
    ts.session_user_id, 
    t.tanggal_transaksi,  
    cr.ori_kli_code AS kode_bank,
    cr.descr,
    cr.retur_reff_no,
    cr.retur_account,
    rt.nama_rekening,        
    d.nilai_mutasi,    
    rec.branch_code,
	cr.ori_kli_code|| ' - ' || coalesce(b.Nama_Lengkap, m.nama_member) as bank_pengirim,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    1 AS data_count
  FROM
    %[2]v.DIE_CREDIT cr 
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = cr.trx_record_id    
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id      
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[1]v.DETILTRANSAKSI d ON r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.TRANSAKSI t ON t.id_transaksi = d.id_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on cr.ori_kli_code = b.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m on cr.ori_kli_code = m.kode_member
    LEFT JOIN %[1]v.REKENINGTRANSAKSI rt on cr.retur_account = rt.nomor_rekening        
  WHERE
    (d.kode_cabang = $1 OR 'ALL'= '%[5]v')
    AND (t.tanggal_transaksi >= TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
    AND (t.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND t.kode_transaksi = 'INRET'
  UNION
  SELECT
    h.batch_no,
    ts.session_user_id,     
    t.tanggal_transaksi, 
    cr.ori_kli_code AS kode_bank,
    cr.descr,
    cr.retur_reff_no,
    cr.retur_account,
    rt.nama_rekening,             
    d.nilai_mutasi,   
    rec.branch_code,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
	cr.ori_kli_code || ' - ' || coalesce(b.Nama_Lengkap, m.nama_member) as bank_pengirim,
    1 AS data_count
  FROM
    %[2]v.DIE_CREDIT cr 
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = cr.trx_record_id
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id 
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[1]v.HISTDETILTRANSAKSI d ON r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.HISTTRANSAKSI t ON t.id_transaksi = d.id_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on cr.ori_kli_code = b.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m on cr.ori_kli_code = m.kode_member
    LEFT JOIN %[1]v.REKENINGTRANSAKSI rt on cr.retur_account = rt.nomor_rekening           
  WHERE
    (d.kode_cabang = $1 OR 'ALL'= '%[5]v')
    AND (t.tanggal_transaksi >= TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[3]v')
    AND (t.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND t.kode_transaksi = 'INRET'

-- name: GetDataReportTransaksiCabSKNDebet-main
SELECT
    'ALL' as Grouping,  
    t.tanggal_transaksi, 
    t.nomor_referensi,
    t.user_input,     
    t.user_otorisasi,
    t.user_overide, 
    t.jam_input,
    t.kode_cabang_transaksi AS cabang,
    t.is_reversed,
    db.reff_no,
    db.warkat_no, 
    ndo.kode_bank,
    (case ndo.Jenis_Warkat
    when 'C' then 'Cek'
    when 'B' then 'Bilyet Giro'
    when 'N' then 'Nota Debit'
    else ''
    end) as Jenis_Warkat,
    ndo.Nomor_Rekening_Nota,
    ndo.Nama_Pemilik_Nota,      
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    d.nilai_mutasi,
    enum.enum_description,
    ts.session_time,
    ts.session_user_id,         
    1 AS data_count
FROM
    %[1]v.TRANSAKSI t  
    INNER JOIN %[1]v.DETILTRANSAKSI d ON t.id_transaksi = d.id_transaksi
    INNER JOIN %[1]v.DTKLIRINGOUTWARD dto ON  dto.id_detil_transaksi = d.id_detil_transaksi
    INNER JOIN %[1]v.NOTADEBITOUTWARD ndo ON  dto.Id_NotaDebitOutward = ndo.Id_NotaDebitOutward
    LEFT JOIN %[1]v.BANKKLIRING b ON ndo.kode_bank = b.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m on ndo.kode_bank = m.kode_member
    LEFT JOIN %[1]v.REMITTANCETRANSACTION  r ON r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[2]v.TRX_RECORD  rec ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[2]v.DOI_DEBET  db ON rec.trx_record_id = db.trx_record_id 
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id
    LEFT JOIN %[1]v.ENUM_INT enum ON t.status_otorisasi = enum.enum_value AND enum.enum_name = 'eStatusOtorisasiTransaksi'  
WHERE
    (t.kode_cabang_transaksi = $1 OR 'ALL' = $2)
    AND (t.tanggal_transaksi >= TO_DATE($3, 'YYYY-MM-DD') OR 'ALL' = $4)
    AND (t.tanggal_transaksi < TO_DATE($5, 'YYYY-MM-DD') OR 'ALL' = $6)
    AND t.kode_transaksi = 'SDN'
UNION
SELECT 
    'ALL' as Grouping,  
    t.tanggal_transaksi,
    t.nomor_referensi,
    t.user_input,     
    t.user_otorisasi,
    t.user_overide, 
    t.jam_input,
    t.kode_cabang_transaksi AS cabang,
    t.is_reversed,
    db.reff_no,
    db.warkat_no, 
    ndo.kode_bank,
    (case ndo.Jenis_Warkat
    when 'C' then 'Cek'
    when 'B' then 'Bilyet Giro'
    when 'N' then 'Nota Debit'
    else ''
    end) as Jenis_Warkat,
    ndo.Nomor_Rekening_Nota,
    ndo.Nama_Pemilik_Nota,      
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    d.nilai_mutasi,
    enum.enum_description,
    ts.session_time,
    ts.session_user_id, 
    1 AS data_count
FROM
    %[1]v.HISTTRANSAKSI t 
    INNER JOIN %[1]v.HISTDETILTRANSAKSI d ON t.id_transaksi = d.id_transaksi
    INNER JOIN %[1]v.HISTDTKLIRINGOUTWARD dto ON  dto.id_detil_transaksi = d.id_detil_transaksi 
    INNER JOIN %[1]v.NOTADEBITOUTWARD ndo ON  dto.Id_NotaDebitOutward = ndo.Id_NotaDebitOutward
    LEFT JOIN %[1]v.BANKKLIRING b ON ndo.kode_bank = b.kode_bank
    LEFT JOIN %[1]v.MEMBERRTGS m on ndo.kode_bank = m.kode_member
    LEFT JOIN %[1]v.REMITTANCETRANSACTION r ON r.id_detil_transaksi = d.id_detil_transaksi  
    LEFT JOIN %[2]v.TRX_RECORD rec ON r.trx_record_id = rec.trx_record_id
    LEFT JOIN %[2]v.DOI_DEBET db ON rec.trx_record_id = db.trx_record_id  
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id
    LEFT JOIN %[1]v.ENUM_INT enum ON t.status_otorisasi = enum.enum_value AND enum.enum_name = 'eStatusOtorisasiTransaksi'       
WHERE
    (t.kode_cabang_transaksi = $1 OR 'ALL' = $2)
    AND (t.tanggal_transaksi >= TO_DATE($3, 'YYYY-MM-DD') OR 'ALL' = $4)
    AND (t.tanggal_transaksi < TO_DATE($5, 'YYYY-MM-DD') OR 'ALL' = $6)
    AND t.kode_transaksi = 'SDN'
ORDER BY tanggal_transaksi, cabang 

-- name: GetDataReportTransaksiCabSKNReturDebet-main
SELECT 
    h.batch_no , 
    ts.session_date ,  
    rt.ori_kli_code AS kode_bank ,
    rt.warkat_no ,    
    rt.trx_amount ,
    rt.account_no ,
    rt.trx_code ,
    rt.reason ,
    p.keterangan , 
    ts.session_time , 
    1 AS data_count
FROM
    %[2]v.DIE_RETUR rt
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = rt.trx_record_id     
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id
    LEFT JOIN %[1]v.PARAMETERTOLAKANKLIRING p ON p.Id_Parameter_Tolakan = rt.reason 
WHERE
    (ts.session_date >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL' = $2)
    AND (ts.session_date < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL' = $4)
    AND (rec.branch_code = $5  or 'ALL' = $6)
ORDER BY session_date , rec.branch_code 

-- name: GetDataReportTplTransaksiCabSKNTransferInFailed-main
SELECT 
    h.batch_no,
    ts.session_user_id, 
    t.tanggal_transaksi, 
    t.nomor_referensi, 
    cr.ori_kli_code AS kode_bank,
    cr.descr,
    d.nilai_mutasi, 
    t.user_otorisasi, 
    t.jam_input, 
    ts.session_time, 
    rec.branch_code,
    cr.recp_account,
    cr.recp_name,
    cr.sender_name,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    rec.failed_description,        
	cr.ori_kli_code || ' - ' || coalesce(b.Nama_Lengkap, m.nama_member) as bank_pengirim,
    1 AS data_count
  FROM
    %[2]v.DIE_CREDIT cr 
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = cr.trx_record_id      
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id  
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[1]v.DETILTRANSAKSI d ON r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.TRANSAKSI t ON t.id_transaksi = d.id_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on cr.ori_kli_code = b.kode_bank        
    LEFT JOIN %[1]v.MEMBERRTGS m on cr.ori_kli_code = m.kode_member
  WHERE
     (rec.branch_code = $1 OR 'ALL'= '%[5]v')
     AND (t.tanggal_transaksi >= TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[3]v')
     AND (t.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
     AND t.kode_transaksi = 'INSKR'
  UNION
  SELECT
    h.batch_no,
    ts.session_user_id,     
    t.tanggal_transaksi,
    t.nomor_referensi, 
    cr.ori_kli_code AS kode_bank,
    cr.descr,     
    d.nilai_mutasi,
    t.user_otorisasi, 
    t.jam_input, 
    ts.session_time, 
    rec.branch_code,
    cr.recp_account,
    cr.recp_name,
    cr.sender_name,
    coalesce(b.Nama_Lengkap, m.nama_member) as nama_bank,
    rec.failed_description,    
	cr.ori_kli_code || ' - ' || coalesce(b.Nama_Lengkap, m.nama_member) as bank_pengirim,        
    1 AS data_count
  FROM
    %[2]v.DIE_CREDIT cr 
    LEFT JOIN %[2]v.TRX_RECORD rec ON rec.trx_record_id = cr.trx_record_id  
    LEFT JOIN %[2]v.TRX_HEADER h ON h.trx_session_id = rec.trx_session_id
    LEFT JOIN %[2]v.TRX_SESSION ts ON ts.trx_session_id = h.trx_session_id
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r ON r.trx_record_id = rec.trx_record_id 
    LEFT JOIN %[1]v.HISTDETILTRANSAKSI d ON r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.HISTTRANSAKSI t ON t.id_transaksi = d.id_transaksi
    LEFT JOIN %[1]v.BANKKLIRING b on cr.ori_kli_code = b.kode_bank        
    LEFT JOIN %[1]v.MEMBERRTGS m on cr.ori_kli_code = m.kode_member       
  WHERE
    (rec.branch_code = $1 OR 'ALL'= '%[5]v')
     AND (t.tanggal_transaksi >= TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[3]v')
     AND (t.tanggal_transaksi < TO_DATE($3, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
     AND t.kode_transaksi = 'INSKR'
  ORDER BY branch_code, tanggal_transaksi

-- name: GetDataReportTplTransaksiCabRTGSTransferRetur-main
SELECT
    t.tanggal_transaksi, 
    rec.amount,    
    rec.from_member,
    m.nama_member, 
    rec.branch_code AS cabang,
    rec.failed_description,    
    st.originating_name,
    st.ultimate_benef_account,   
    st.ultimate_benef_name,
    st.payment_details,    
    rec.from_member || ' - ' || m.nama_member as nama_bank,        
    1 AS data_count
FROM
    %[2]v.RTTRANSACTION rec
    LEFT JOIN %[2]v.RTSESSION h on h.session_id = rec.session_id  
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r on r.trx_record_id = rec.trans_id 
    LEFT JOIN %[1]v.DETILTRANSAKSI d on r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.TRANSAKSI t  on t.id_transaksi = d.id_transaksi
    LEFT JOIN %[2]v.RTSINGLETRANSFER st on st.transdetail_id = rec.transdetail_id
    LEFT JOIN %[1]v.MEMBERRTGS m on m.kode_member = rec.from_member  
WHERE    
    (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
    AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND h.hdr_message_type='O'
    AND t.kode_transaksi = 'RTRET'
    AND rec.branch_code = $3
UNION
SELECT
    t.tanggal_transaksi,
    rec.amount,    
    rec.from_member,
    m.nama_member,     
    rec.branch_code AS cabang,
    rec.failed_description,
    st.originating_name,
    st.ultimate_benef_account,   
    st.ultimate_benef_name,
    st.payment_details,    
    rec.from_member || ' - ' || m.nama_member as nama_bank,       
    1 AS data_count
FROM
    %[2]v.RTTRANSACTION rec                                                          
    LEFT JOIN %[2]v.RTSESSION h on h.session_id = rec.session_id
    LEFT JOIN %[2]v.REMITTANCETRANSACTION r on r.trx_record_id = rec.trans_id
    LEFT JOIN %[1]v.HISTDETILTRANSAKSI d on r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[1]v.HISTTRANSAKSI t on t.id_transaksi = d.id_transaksi
    LEFT JOIN %[2]v.RTSINGLETRANSFER st on st.transdetail_id = rec.transdetail_id
    LEFT JOIN %[1]v.MEMBERRTGS m on m.kode_member = rec.from_member    
WHERE
    (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'='%[3]v')
    AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND h.hdr_message_type='O'
    AND t.kode_transaksi = 'RTRET'
    AND rec.branch_code = $3
ORDER BY cabang, tanggal_transaksi

-- name: GetDataReportTplTransaksiCabRTGSTransferInFailed-main
SELECT
    t.tanggal_transaksi, 
    rec.amount,    
    rec.from_member,
    m.nama_member, 
    rec.branch_code AS cabang,
    rec.failed_description,    
    st.originating_name,
    st.ultimate_benef_account,   
    st.ultimate_benef_name,
    st.payment_details,    
	rec.from_member || ' - ' || m.nama_member as bank_pengirim,       
    1 AS data_count
  FROM
    %[2]v.RTTRANSACTION rec
    LEFT JOIN %[2]v.RTSESSION h on h.session_id = rec.session_id  
    LEFT JOIN %[1]v.REMITTANCETRANSACTION r on r.trx_record_id = rec.trans_id 
    LEFT JOIN %[1]v.DETILTRANSAKSI d on r.id_detil_transaksi = d.id_detil_transaksi 
    LEFT JOIN %[1]v.TRANSAKSI t on t.id_transaksi = d.id_transaksi
    LEFT JOIN %[2]v.RTSINGLETRANSFER st on st.transdetail_id = rec.transdetail_id
    LEFT JOIN %[1]v.MEMBERRTGS m on m.kode_member = rec.from_member  
  WHERE    
    (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'= '%[3]v')
    AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND h.hdr_message_type='O'
    AND t.kode_transaksi = 'RTSKR'
    AND rec.branch_code = $3
  UNION
  SELECT
    t.tanggal_transaksi,
    rec.amount,    
    rec.from_member,
    m.nama_member,     
    rec.branch_code AS cabang,
    rec.failed_description,
    st.originating_name,
    st.ultimate_benef_account,   
    st.ultimate_benef_name,
    st.payment_details,    
	rec.from_member || ' - ' || m.nama_member as bank_pengirim,       
    1 AS data_count
  FROM
    %[2]v.RTTRANSACTION rec                                                          
    LEFT JOIN %[2]v.RTSESSION h on h.session_id = rec.session_id
    LEFT JOIN %[1]v.REMITTANCETRANSACTION r on r.trx_record_id = rec.trans_id
    LEFT JOIN %[1]v.HISTDETILTRANSAKSI d on r.id_detil_transaksi = d.id_detil_transaksi
    LEFT JOIN %[1]v.HISTTRANSAKSI t on t.id_transaksi = d.id_transaksi
    LEFT JOIN %[2]v.RTSINGLETRANSFER  st on st.transdetail_id = rec.transdetail_id
    LEFT JOIN %[1]v.MEMBERRTGS m on m.kode_member = rec.from_member    
  WHERE
    (t.tanggal_transaksi >= TO_DATE($1, 'YYYY-MM-DD') OR 'ALL'='%[3]v')
    AND (t.tanggal_transaksi < TO_DATE($2, 'YYYY-MM-DD') OR 'ALL'='%[4]v')
    AND h.hdr_message_type='O'
    AND t.kode_transaksi = 'RTSKR'
    AND rec.branch_code = $3
  ORDER BY cabang, tanggal_transaksi

-- name: GetDataReportTransaksiJurnalSendiri-main
SELECT 
		'1' as fgrouping,
		j.journal_no ,
		j.description,
		j.journal_type,
		j.reference_no, 
		ji.journalitem_no,
		ji.journalitemstatus as status,
		acc.account_code,
		ji.subaccountcode,
		ji.subaccountname,
		acc.account_name,
		acc.account_type,
		j.branch_code,
		j.journal_date,
		j.Is_Confidential,
		( case when (j.Is_Confidential = 'T' and j.Id_Confidential in ('-1')) then 'T'
		else 'F'
		end) as ConfidentialAllowed,
		ji.branch_code as branch_code_account,
		ji.Valuta_Code,
		ji.Nilai_Kurs,
		ji.ID_JournalBlock,
		ji.amount_debit,
		ji.amount_credit,
		ji.fl_project_no,
		ji.description as ji_description,
		t.nomor_referensi as nomor_referensi_trans,
		t.jam_input as jam_input_trans,
		t.user_otorisasi as user_otor_trans,
		t.user_overide as user_ovr_trans,
		t.user_input as user_input_trans,
		ht.nomor_referensi as nomor_referensi_histtrans,
		ht.jam_input as jam_input_histtrans,
		ht.user_otorisasi as user_otor_histtrans,
		ht.user_overide as user_ovr_histtrans,
		ht.user_input as user_input_histtrans  
	FROM 
		%[1]v.JOURNALITEM ji, 
		%[1]v.ACCOUNT acc ,
		%[1]v.JOURNAL j
		LEFT JOIN %[1]v.TRANSAKSI t on (j.journal_no = t.journal_no)
		LEFT JOIN %[1]v.HISTTRANSAKSI ht on (j.journal_no = ht.journal_no)
	WHERE ji.fl_account = acc.account_code
		and ji.fl_journal = j.journal_no
		and j.journal_date >= TO_DATE($1, 'YYYY-MM-DD')  
		and j.journal_date < TO_DATE($2, 'YYYY-MM-DD') 
		and j.UserId_Create  = $3                     
	ORDER By
		fgrouping, j.branch_code, j.journal_no, ji.branch_code, amount_debit desc, amount_credit desc 

-- name: GetDataReportRekapTransaksiSendiri-main
SELECT min(Kode_Valuta) as kode_valuta, 
    Kode_Cabang_Transaksi, 
    Kode_Transaksi, 
    COUNT(Kode_Transaksi) as jumlah, 
    sum(non_cash_Debet) as non_cash_Debet, 
    sum(non_cash_Kredit) as non_cash_Kredit,
    sum(cash_Debet) as cash_Debet,
    sum(cash_Kredit) as cash_Kredit,
    (sum(non_cash_Debet) - sum(non_cash_Kredit) + sum(cash_Debet) -sum(cash_Kredit)) as ending_balance
FROM   (
    SELECT T.Kode_Cabang_Transaksi,
            T.Kode_Transaksi,
            DT.Kode_Valuta,
            (case 
            when DT.Jenis_Mutasi = 'D' and (I.is_tunai is null or I.is_tunai= 'F') then DT.Nilai_Mutasi
            else 0.0
            end) as non_cash_Debet,
            (case
            when DT.Jenis_Mutasi = 'C' and (I.is_tunai is null or I.is_tunai= 'F') then DT.Nilai_Mutasi
            else 0.0
            end) as non_cash_Kredit,
            (case
            when DT.Jenis_Mutasi = 'D' and I.is_tunai = 'T' then DT.Nilai_Mutasi
            else 0.0
            end) as cash_Debet,
            (case 
            when DT.Jenis_Mutasi = 'C' and I.is_tunai = 'T' then DT.Nilai_Mutasi
            else 0.0
            end) as cash_Kredit,
            DT.Jenis_Mutasi,
            DT.Nilai_Mutasi,
            I.is_tunai,
            T.Id_Transaksi,
            TO_CHAR(T.Tanggal_Transaksi, 'DD-MM-YYYY') || ' '|| TO_CHAR(T.Jam_Input, 'HH24::MI::SS') as Jam_Input                         
    FROM
        %[1]s.rekeningtransaksi RT,
        %[1]s.histdetiltransaksi DT,
        %[1]s.histtransaksi T
        LEFT JOIN %[1]s.INFOTRANSAKSI I ON I.id_transaksi = T.id_transaksi
    WHERE
        T.Id_Transaksi = DT.Id_Transaksi
        and DT.Nomor_Rekening = RT.Nomor_Rekening          
        and T.Tanggal_Transaksi >= TO_TIMESTAMP($1,'YYYY-MM-DD')
        and T.Tanggal_Transaksi < TO_TIMESTAMP($2,'YYYY-MM-DD') + 1
        and T.User_Input = $3
    UNION ALL
    SELECT T.Kode_Cabang_Transaksi,
            T.Kode_Transaksi,
            DT.Kode_Valuta,
            (case 
            when DT.Jenis_Mutasi = 'D' and (I.is_tunai is null or I.is_tunai= 'F') then DT.Nilai_Mutasi
            else 0.0
            end) as non_cash_Debet,
            (case
            when DT.Jenis_Mutasi = 'C' and (I.is_tunai is null or I.is_tunai= 'F') then DT.Nilai_Mutasi
            else 0.0
            end) as non_cash_Kredit,
            (case
            when DT.Jenis_Mutasi = 'D' and I.is_tunai = 'T' then DT.Nilai_Mutasi
            else 0.0
            end) as cash_Debet,
            (case 
            when DT.Jenis_Mutasi = 'C' and I.is_tunai = 'T' then DT.Nilai_Mutasi
            else 0.0
            end) as cash_Kredit,
            DT.Jenis_Mutasi,
            DT.Nilai_Mutasi,
            I.is_tunai,
            T.Id_Transaksi,
            TO_CHAR(T.Tanggal_Transaksi, 'DD-MM-YYYY') || ' '|| TO_CHAR(T.Jam_Input, 'HH24::MI::SS') as Jam_Input                
    FROM
        %[1]s.rekeningtransaksi RT,
        %[1]s.detiltransaksi DT,
        %[1]s.transaksi T
        LEFT JOIN %[1]s.INFOTRANSAKSI I ON I.id_transaksi = T.id_transaksi
    WHERE
        T.Id_Transaksi = DT.Id_Transaksi
        and DT.Nomor_Rekening = RT.Nomor_Rekening          
        and T.Tanggal_Transaksi >= TO_TIMESTAMP($1,'YYYY-MM-DD')
        and T.Tanggal_Transaksi < TO_TIMESTAMP($2,'YYYY-MM-DD') + 1      
        and T.User_Input = $3

    ORDER By
        Kode_Cabang_Transaksi, Jam_Input, Id_Transaksi, Jenis_Mutasi
)
GROUP BY Kode_Transaksi,Kode_Cabang_Transaksi

-- name: GetDataReportTransaksiReversalOtorisasi-main
SELECT 
    RT.Nomor_Rekening, 
        RT.Nama_Rekening,
        RT.Kode_Cabang as Kode_Cabang_Rek,
        T.Keterangan,
        DT.Keterangan as Keterangan_Detail,
        T.Nomor_Referensi,
        T.Id_Transaksi,
        T.Kode_Transaksi,
        T.User_Overide,
        T.User_Otorisasi,
        TO_CHAR(T.Tanggal_Transaksi, 'DD-MM-YYYY') as Tanggal_Transaksi,
        TO_CHAR(T.Jam_Input, 'HH24::MI::SS') as Jam_Input,
        TO_CHAR(T.Tanggal_Transaksi, 'DD-MM-YYYY') || ' / ' || to_char(t.jam_input, 'HH24::MI::SS') tanggal_transaksi_jam_input,
        TO_CHAR(TRV.Tanggal_Transaksi, 'DD-MM-YYYY') as Tanggal_Reverse,              
        TO_CHAR(TRV.Jam_Input, 'HH24::MI::SS') as Jam_Reverse,
        TO_CHAR(TRV.Tanggal_Transaksi, 'DD-MM-YYYY') || ' / ' || to_char(TRV.jam_input, 'HH24::MI::SS') tanggal_reverse_jam_reverse,
        TRV.User_Input as User_Reverse,
        TRV.User_Overide as User_Overide_Reverse,
        DT.Kode_Valuta,
        (case DT.Jenis_Mutasi
        when 'D' then DT.Nilai_Mutasi
        else 0.0
        end) as Nominal_Debet,
        (case DT.Jenis_Mutasi
        when 'C' then DT.Nilai_Mutasi
        else 0.0
        end) as Nominal_Kredit,
        1 as data_count                
FROM
    %[1]s.rekeningtransaksi RT,
    %[1]s.detiltransaksi DT,
    %[1]s.transaksi T,
    %[1]s.reversetransaction RV,
    %[1]s.transaksi TRV
WHERE
    T.Id_Transaksi = DT.Id_Transaksi
    and DT.Nomor_Rekening = RT.Nomor_Rekening          
    and T.Id_Transaksi = RV.Id_Transaksi
    and TRV.Id_Transaksi = RV.Reverse_Id
    and T.Status_Otorisasi = 1
    AND TRV.TANGGAL_TRANSAKSI >= TO_TIMESTAMP($1,'YYYY-MM-DD') 
    AND TRV.TANGGAL_TRANSAKSI < TO_TIMESTAMP($2,'YYYY-MM-DD') + 1
    and (T.Kode_Cabang_Transaksi = $3 or 'ALL' = $4)  
ORDER By
    T.Kode_Cabang_Transaksi, T.Id_Transaksi, DT.Jenis_Mutasi

-- name: GetReportDetail-main
SELECT
    KODE_REPORT, NAMA_REPORT, TEMPLATE_NAME, TAG_REPORT, 
    IS_EOD_EXECUTE, RECIPIENT, RETENSI, IS_SHOW_BDS
FROM %[1]s.REPORT
WHERE KODE_REPORT = $1

-- name: GetRestrictCode-main
SELECT id_cfgroup, cf_code, description 
FROM %[1]s.confidentialgroup

-- name: UpdateAccountRestrictCode-main
update %[1]s.rekeningliabilitas
set status_restriksi = $1,
    id_cfgroup = $2
where nomor_rekening = $3

-- name: UpdateIsReversed-main
update %[1]s.infouid
set is_reversed = $1
where uidkey = $2
    and keyint = $3

-- name: InsertReverseTransaction-main
insert into %[1]s.reversetransaction 
(id_transaksi ,tanggal_reverse ,user_reverse ,reverse_id) 
values ($1, $2, $3, $4)

-- name: GetRevTrxByIdTrx-main
select id_transaksi, tanggal_reverse, user_reverse, reverse_id
from %[1]s.reversetransaction
where id_transaksi = $1

-- name: GetSknTrxcodeList-main
SELECT
kode_tx_skn, nama_tx_skn
FROM %[1]v.kodetx_skn

-- name: CreateStandingInstruction-main
insert into %[1]s.standinginstruction
(instructionid, instructiontypecode, creditaccounttype, debitaccounttype, registerdate, 
    branchcode, description, standinginstructiontype, debitaccountno, debitaccountname, creditaccountno, creditaccountname, amount)
values 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

-- name: GetDataHeaderSvs-main
select rt.nomor_rekening, rt.nama_rekening, rl.nomor_nasabah, cd.remark, cd.additionalremark
from %[1]s.rekeningtransaksi rt 
inner join %[1]s.rekeningliabilitas rl on rl.nomor_rekening = rt.nomor_rekening
left join %[1]s.custaccountimagedesc cd on cd.accountno = rt.nomor_rekening
where rt.nomor_rekening = $1

-- name: GetDataListImageSvs-main
select  imageid, imagetag, remark, id_data
from %[1]s.custaccountimage
where accountno = $1

-- name: CheckSpecimentAcctList-main
SELECT accountno FROM %[1]s.custaccountimagedesc WHERE accountno = $1

-- name: UpdateSpecimentAcctList-main
UPDATE %[1]s.custaccountimagedesc
SET remark = $1, additionalremark = $2
WHERE accountno = $3

-- name: PostSpecimentAcctList-main
INSERT INTO %[1]s.custaccountimagedesc (accountno, remark, additionalremark)
VALUES ($1, $2, $3)

-- name: InsertCustAccountImage-main
insert into %[1]s.custaccountimage (IMAGEID, IMAGETAG, REMARK, ACCOUNTNO, ID_DATA, TAGSEQUENCE, DATATYPE)
VALUES ( nextval('%[1]s.seq_custaccountimage'), $1, $2, $3, $4, $5, 1)

-- name: GetMaxSequenceImageTag-main
SELECT IMAGETAG ,MAX(c.TAGSEQUENCE) AS MAXSEQUENCE
FROM %[1]s.CUSTACCOUNTIMAGE c 
WHERE c.ACCOUNTNO = $1
GROUP BY IMAGETAG, IMAGETAG 

-- name: DeleteSpecimentAcct-main
delete from %[1]s.custaccountimage where imageid in ($1)
 