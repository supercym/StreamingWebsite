package dbops

import (
	"fmt"
	"testing"
)

func TestSession(t *testing.T)  {
	clearTables()
	t.Run("Ins", testInsertSession1)
	t.Run("Ret", testRetrieveSession)
	t.Run("Del", testDeleteSession)
	t.Run("ReRet", testReRetrieveSession)
	t.Run("Ins", testInsertSession1)
	t.Run("Ins", testInsertSession2)
	t.Run("ReAll", testRetrieveAllSession)
}

func testInsertSession1(t *testing.T)  {
	err := InsertSession("001", 24000, "alex")
	if err != nil {
		t.Errorf("Error of InsertSession: %v", err)
	}
}

func testInsertSession2(t *testing.T)  {
	err := InsertSession("002", 36000, "jack")
	if err != nil {
		t.Errorf("Error of InsertSession: %v", err)
	}
}

func testRetrieveSession(t *testing.T)  {
	session, err := RetrieveSession("001")
	if err != nil {
		t.Errorf("Error of GetSession")
	}

	fmt.Printf("test Retrievesession content: %+v\n", *session)
}

func testDeleteSession(t *testing.T)  {
	err := DeleteSession("001")
	if err != nil {
		t.Errorf("Error of DeleteSession: %v", err)
	}
}

func testReRetrieveSession(t *testing.T)  {
	session, err := RetrieveSession("001")
	if err != nil || session != nil {
		t.Errorf("Error of ReRetrieveSession")
	}

}

func testRetrieveAllSession(t *testing.T)  {
	_, err := RetrieveAllSessions()
	if err != nil {
		t.Errorf("Error of RetrieveAllSession: %v", err)
	}
}


