package util

import "strconv"

const DatabaseSQL string = " create databse if not exists locks  ;"

const TableSQL string = " create table if not exists locks.sequence (" +
	" id int primary key , " +
	" maxseq bigint primary key ," +
	" updatetime timestamp(6) default current_timestamp ," +
	" createtime timestamp(6) default current_timestamp );"

const RecordSQL string = " insert into locks.sequence (id,maxSeq,updateTime,createTime ) " +
	" values (1,1,now(),now());"

const GetMaxSQL string = "select maxseq from locks.sequence where id =1 for update ;"

//const UpdateMaxSQL string = "update locks.sequence set maxseq + "

func GeneralUpdateMax(CacheNum int32) (res string) {
	res = "update locks.sequence set maxseq = maxseq +  "
	res += strconv.Itoa(int(CacheNum))
	res += " where id =1 ;"
	return res
}
