package dbops

import (
	"testing"
)

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("alex", "123asd")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("alex")
	if err != nil || pwd != "123asd" {
		t.Errorf("Error of Getuser")
	}
}

func testDeleteUser(t *testing.T)  {
	err := DeleteUser("alex", "123asd")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("alex")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleteing user test failed")
	}
}
