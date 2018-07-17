package mssql

//【注】在使用些方法前，请在调用方法的主类文件中引用包："github.com/mattn/go-adodb"
import (
	"database/sql"
	"fmt"
	"strings"
)

//MSSQLConnection is connection model
type MSSQLConnection struct {
	DataSource         string
	InitialCatalog     string
	IsWindowsAuthorize bool
	UserName           string
	Password           string
	Port               int
}

//Open will open the db according by connection string
func (conn *MSSQLConnection) Open() (db *sql.DB, err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+conn.DataSource)
	conf = append(conf, "Initial Catalog="+conn.InitialCatalog)

	// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器
	// (需要在安装sqlserver时候设置)，
	// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
	if conn.IsWindowsAuthorize {
		conf = append(conf, "integrated security=SSPI")
	} else {
		conf = append(conf, "user id="+conn.UserName)
		conf = append(conf, "password="+conn.Password)
		conf = append(conf, "port="+fmt.Sprint(conn.Port))
	}

	db, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

//Open the db by connection string
//connection string sample: "server=192.18.56.111;user id=sa;password=test;port=1433;database=MyDB;"
func Open(connectionString string) (db *sql.DB, err error) {
	//db, err := sql.Open("mssql", config.ConnectionString)
	//NOTICE: the driver name "adodb" was depended the package "github.com/mattn/go-adodb"
	db, err = sql.Open("adodb", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
