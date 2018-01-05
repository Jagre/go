package mssql

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	cfg := DBConfiguration{
		ConnectionString: "server=192.18.56.111;user id=sa;password=test;port=1433;database=MyDB;",
		FilePath:         `f:\go\src\jagre\gmodel\mssql\`,
		FileName:         "test.go",
	}
	e := cfg.Output()
	if e != nil {
		fmt.Println(e)
	}

}
