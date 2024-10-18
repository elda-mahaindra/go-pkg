-- name: get-gl-list
select acc.account_code, acc.account_name, acc.account_type  
from %[1]s.glaccount g
    inner join %[1]s.accountinstance ai on g.accountinstance_id = ai.accountinstance_id and ai.currency_code = :currency_code and ai.branch_code = :kode_cabang and ai.isbolehinput = 'T'
    inner join %[1]s.account acc on acc.account_code = ai.account_code
where %[2]s
group by acc.account_code, acc.account_name, acc.account_type
order by acc.account_code, acc.account_name, acc.account_type

-- name: get-hist-trx-detail-by-uidkey
select i.uidkey, t.id_transaksi ,d.kode_jurnal, t.kode_cabang_transaksi, rt.nama_rekening,d.nomor_rekening ,d.keterangan ,d.kode_valuta, d.jenis_mutasi,
    case
		when d.jenis_mutasi = 'D' then 'Debet'
		when d.jenis_mutasi = 'C' then 'Kredit'
	end as jenis_mutasi_desc, d.nilai_mutasi ,d.nilai_kurs_manual ,d.nilai_ekuivalen ,d.kode_rc
from %[2]v.infouid i
	inner join %[2]v.histdetiltransaksi d on d.id_transaksi = i.keyint
	inner join %[2]v.histtransaksi t on t.id_transaksi = d.id_transaksi
	inner join %[2]v.rekeningtransaksi rt on rt.nomor_rekening = d.nomor_rekening
where action_type = :action_type
%[1]v
order by d.id_transaksi ,d.id_detil_transaksi

-- name: get-validation-limit-trx
select ptu.kode_transaksi, ptu.keterangan, ptu.jenis_limit_transaksi, lt.id_user, lt.nilai_limit 
from %[1]s.parametertransaksiumum ptu
	inner join %[2]s.limittransaksi lt on lt.jenis_limit = ptu.jenis_limit_transaksi 
where kode_transaksi = :KODE_TRANSAKSI
    and lt.id_user = :USER_ID

-- name: get-rekening-liab
select rl.nomor_nasabah, rl.nomor_rekening, rt.nama_rekening, rt.saldo, rt.kode_cabang, nvl(rl.is_tutup_kartu_atm, 'F') as is_tutup_kartu_atm,
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
	rt.user_input, rl.user_otorisasi, rt.keterangan, rt.kode_valuta, rt.nama_valuta, nvl(rl.is_sedang_ditutup, 'F') is_sedang_ditutup, rl.kode_program, 
	rl.kode_marketing_referensi, rl.status_kelengkapan, rl.is_blokir, rl.is_blokir_debet, nvl(rl.is_boleh_debet, 'T') as is_boleh_debet, 
	rl.is_blokir_kredit, nvl(rl.is_boleh_kredit, 'T') as is_boleh_kredit, rl.is_cetak_nota, rl.is_dapat_bagi_hasil, rl.is_tidak_dormant, rl.is_biaya_rekening_dormant, rl.is_biaya_saldo_minimum, 
	rl.is_biaya_atm, rl.nisbah_bagi_hasil, rl.is_tiering_nisbahbonus, rl.is_join_account ,rl.kode_multicifrelation, rt.nomor_rekening,
	rt.saldo - nvl(rl.saldo_ditahan, 0) - p.saldo_minimum - rl.saldo_float as saldo_efektif, rt.kode_jenis ,status_restriksi, to_char(rl.tanggal_tutup, 'YYYY-MM-DD') as tanggal_tutup
from %[1]s.rekeningliabilitas rl
	inner join %[1]s.rekeningtransaksi rt on rl.nomor_rekening = rt.nomor_rekening
	left join %[1]s.produk p on p.kode_produk = rl.kode_produk
	left join %[1]s.registerlayanan reglay on reglay.nomor_rekening_layanan = rt.nomor_rekening
	left join %[1]s.registerpassbook rpb on reglay.id_register = rpb.id_register
			inner join %[1]s.nasabah n on n.nomor_nasabah = rl.nomor_nasabah
	left join %[1]s.nasabahindividu ni on ni.nomor_nasabah = n.nomor_nasabah
	left join %[1]s.individu i on i.id_individu = ni.id_individu
	left join %[1]s.nasabahkorporat nk on nk.nomor_nasabah = n.nomor_nasabah 
where rl.nomor_rekening = :nomor_rekening

-- name: get-max-db-harian
select count(*) as total_limit
from %[1]s.limittransaksirekening lt
where nomor_rekening = :NOMOR_REKENING 
    and periode = 'H' 
	and jenis_mutasi = 'D'

-- name: get-limit-debet-harian
select sum(nilai_transaksi) as total_nilai_transaksi
from %[1]s.limittransaksirekening lt
where nomor_rekening = :NOMOR_REKENING 
    and periode = 'H' 
	and jenis_mutasi = 'D'

-- name: get-counter-debet-harian
select sum(nilai_transaksi) as total_nilai_transaksi
from %[1]s.countertransaksirekening lt
where nomor_rekening = :NOMOR_REKENING 
	and periode = 'H' 
    and jenis_mutasi = 'D'

-- name: get-rc-code-list
SELECT PROJECT_NO AS rc_code, name AS rc_name FROM %[1]s.project

-- name: get-pod
select to_char(nilai_parameter_tanggal, 'YYYY-MM-DD') as today
from %s.parameterglobal
where kode_parameter = :kode_parameter

-- name: get-saldo-by-date
select round(balance, 2) as balance
from %[1]s.dailybalancerekening d1
	inner join (
		select nomor_rekening, max(balance_date) as balance_date 
		from %[1]s.dailybalancerekening 
		where nomor_rekening = :nomor_rekening1
			and balance_date <= to_date(:tanggal1, 'yyyy-mm-dd')
			and balance_field = 'saldo'
		group by nomor_rekening
	) d2 on (d1.balance_date =  d2.balance_date) and (d1.nomor_rekening = d2.nomor_rekening)
where d1.nomor_rekening = :nomor_rekening2    
	and d1.balance_field = 'saldo'
	and d1.balance_date <= to_date(:tanggal2, 'yyyy-mm-dd')

-- name: get-hist-trx-list-by-norek
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
where ht.tanggal_transaksi >= to_date(:TANGGAL_AWAL, 'yyyy-mm-dd') 
	and ht.tanggal_transaksi < to_date(:TANGGAL_AKHIR, 'yyyy-mm-dd') + 1
	and hd.nomor_rekening = :NOMOR_REKENING
order by hd.id_detil_transaksi

-- name: get-trx-list-by-norek
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
where d.nomor_rekening = :NOMOR_REKENING
order by d.id_detil_transaksi

-- name: get-trx-detail
select d.id_detil_transaksi, d.id_transaksi, d.jenis_mutasi, d.kode_valuta, to_char(d.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi,
	d.nilai_mutasi, d.keterangan, d.kode_cabang, d.jenis_detil_transaksi, d.nilai_kurs_manual, d.saldo_awal, d.nilai_ekuivalen,
	d.nomor_referensi, d.kode_account, d.nomor_rekening, d.kode_jurnal, d.id_parameter_transaksi, d.flag_proses, d."SEQUENCE"
from %v.detiltransaksi d
where d.id_transaksi = :id_transaksi

-- name: get-trx-detail-by-uidkey
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
where action_type = :action_type
%[1]v
order by d.id_transaksi ,d.id_detil_transaksi

-- name: get-trx-list
select t.id_transaksi, to_char(t.tanggal_transaksi, 'YYYY-MM-DD') as tanggal_transaksi, t.nomor_referensi,
	t.kode_cabang_transaksi, t.nilai_transaksi, t.keterangan
from %v.transaksi t
%v
order by t.jam_input
offset :bw1 rows fetch next :bw2 rows only

-- name: detail-gl-jurnal
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
WHERE j.journal_no = :journal_no

-- name: detail-gl-jurnal-item
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
WHERE ji.fl_journal = :journal_no 
	AND ji.fl_journal not LIKE 'POD.%%'
order by nomor

-- name: detail-gl-jurnal-block
SELECT 
	SubSystemCode as Kode_Sub_Sistem,
	class_id,
	key_id,
	ID_JournalBlock,
	keterangan
FROM %[1]s.journalblock j 
WHERE journal_no = :journal_no 
	AND journal_no not LIKE 'POD.%%'
ORDER BY ID_JournalBlock

-- name: get-gl-jurnal
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
OFFSET :offset ROWS FETCH NEXT :limit ROWS ONLY

-- name: get-info-uid
select uidkey ,to_char(tanggal_input, 'yyyy-mm-dd') as tanggal_input ,action_type ,user_input ,user_otor, keterangan ,info1 ,info2 
	,kode_cabang, keyint, 'T' as is_exists, is_reversed
from %[1]s.infouid i 
where uidkey = :uidkey
	and action_type = :action_type

-- name: get-transaksi-id
select %s.seq_transaksi.nextval as id_transaksi from dual