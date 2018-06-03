package routers

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2"
)

func InitDBConn(t *testing.T) (*mgo.Database, error) {
	session, err := mgo.Dial("mongodb://test:test123@ds035240.mongolab.com:35240/vim_tips")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("Failed to connect to mongo DB...")
	} else {
		t.Log("Connected to mongo DB...")
	}

	return session.DB("vim_tips"), err
}

func TestHandleRandomTxtTip(t *testing.T) {
	db, err := InitDBConn(t)

	if db != nil && err == nil {

	}

}
