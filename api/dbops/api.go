package dbops

import (	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"../utils"
	"../defs"
	"time"
)


func AddUserCredential(loginName string,pwd string) error {  //增
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name,pwd)VALUES (?,?)") //预编译
	if err != nil {
		return err
	}
	_,err=stmtIns.Exec(loginName, pwd)  //执行
	if err!=nil{
		return  err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string,error)  { //查
	stmtOut,err:=dbConn.Prepare("SELECT pwd FROM users where Login_name=?")
	if err!=nil{
		log.Printf("%s",err)
		return "",err
	}

	var pwd string
	err=stmtOut.QueryRow(loginName).Scan(&pwd) //查询一整行
	if err!=nil&&err!=sql.ErrNoRows{   //sql.ErrNoRows正确错误
		return "",err
	}
	 defer stmtOut.Close()
	return pwd,nil
}

func DeletUser(loginName,pwd string) error {
	stmt,err:=dbConn.Prepare("DElETE FROM users where login_name=? AND pwd=?")
	if err!=nil{
		log.Printf("Delete User %s",err)
		return err
	}
	_,err=stmt.Exec(loginName,pwd)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	return nil
}

func AddNewVideo(aid int,name string)(*defs.VideoInfo,error)  {
	//creatUUID
	vid,err:=utils.NewUUID()
	if err!=nil{
		return nil ,err
	}
	t:=time.Now()
	ctime:=t.Format("2006-01-02 13:04:05")
	stmtIns,err:=dbConn.Prepare("INSERT INTO video_info (vid , authot_id,name,display_ctime,creat_time) VALUES (?,?,?,?)")
	_,err=dbConn.Exec(vid,aid,name,ctime)
	if err!=nil{
		return nil,err
	}

	res:=&defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil
}