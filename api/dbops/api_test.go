package dbops

import ("testing"
	_ "github.com/go-sql-driver/mysql"
)

//init (dblogin,truncate tables)->	run tests ->cleat data(truncate tables)

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add",testAddUser)
	t.Run("get",testGetUser)
	t.Run("delete",testDeletUser)
	t.Run("reget",testRegetUser)
}

func testAddUser(t *testing.T)  {
	err:=AddUserCredential("root","xiner")
	if err!=nil{
		t.Errorf("AddUser err:%v",err)
	}

}

func testGetUser(t *testing.T)  {
	pwd,err:=GetUserCredential("root")
	if err!=nil||pwd!="xiner"{
		t.Errorf("GetUserCredential err:%v",err)
	}
}

func testDeletUser(t *testing.T) {
	err:=DeletUser("liuzhen","123")
	if err!=nil{
		t.Errorf("DeletUser err: %v",err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd,err:=GetUserCredential("liuzhen")
	if  err!=nil{
		t.Errorf("testRegetUser err: %v",err)
	}

	if pwd!=""{
		t.Errorf("delete err: %v",err)
	}
}

