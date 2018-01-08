package mssql

import (
	"database/sql"
	"errors"
	"fmt"
	. "github.com/Jagre/go/lib/gmodel"
	_ "github.com/denisenkom/go-mssqldb"
	"io/ioutil"
	"os"
	"strings"
)

//DBConfiguration is
type DBConfiguration struct {
	ConnectionString string
	FilePath         string
	FileName         string
}

func init() {
	fmt.Println("Starting...")
}

//GetGModelType implement the interface IGModel
func GetGModelType(sqlType string) string {
	var goType string
	switch sqlType {
	case "varchar", "nvarchar":
		goType = "string"
	case "bit":
		goType = "bool"
	case "tinyint":
		goType = "byte"
	case "int":
		goType = "int"
	case "bigint":
		goType = "int64"
	case "datetime":
		goType = "time.Time"
	case "varbinary":
		goType = "[]byte"
	default:
		goType = sqlType
	}
	return goType
}

//GetModels implement the interface IGModel
func (config *DBConfiguration) GetModels() []*GModel {
	gmodels := []*GModel{}
	db, e := sql.Open("mssql", config.ConnectionString)
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()
	//Get tables' info
	rows, e := db.Query("Select TABLE_NAME From INFORMATION_SCHEMA.TABLES Where TABLE_NAME <> 'sysdiagrams'")
	if e != nil {
		fmt.Println(e)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		e = rows.Scan(&name)
		if e != nil {
			fmt.Println(e)
		}
		gmodels = append(gmodels, &GModel{ModelName: name, Properties: []*GModelProperty{}})
	}

	//Get columns' info
	for _, m := range gmodels {
		rows, e = db.Query("Select COLUMN_NAME, DATA_TYPE From INFORMATION_SCHEMA.COLUMNS Where TABLE_NAME = ? Order By ORDINAL_POSITION", m.ModelName)
		if e != nil {
			fmt.Println(e)
		}
		defer rows.Close()
		//properties := []gmodel.GModelProperty{}
		for rows.Next() {
			var name, originalType, goType string
			e = rows.Scan(&name, &originalType)
			if e != nil {
				fmt.Println(e)
			}
			goType = GetGModelType(originalType)
			newModel := GModelProperty{
				Name:         name,
				OriginalType: originalType,
				GoType:       goType,
			}

			m.Properties = append(m.Properties, &newModel)
		}
		//m.Properties = properties
	}
	return gmodels
}

//Output will output to file
//TODO: 1. Package; 2. Convert sql's type & golang's type each other
func (config *DBConfiguration) Output() error {
	if len(config.FileName) == 0 {
		return errors.New("Pls specified the file name")
	}
	if len(config.FilePath) == 0 {
		config.FilePath, _ = os.Getwd()
	}
	//Set path end with "/"
	config.FilePath = strings.Replace(config.FilePath, "\\", "/", -1)
	if !strings.HasSuffix(config.FilePath, "/") {
		config.FilePath += "/"
	}
	//Set PackageName
	dirNodes := strings.Split(config.FilePath, "/")
	namespace := "Models"
	if len(dirNodes) > 0 {
		namespace = dirNodes[len(dirNodes)-2]
	}
	modelString := fmt.Sprintf("package %s \n***\n", namespace)

	models := config.GetModels()
	for _, m := range models {
		modelString += m.GModelGenerate()
	}

	if strings.Contains(modelString, "time.Time") {
		modelString = strings.Replace(modelString, "***", "import(\"time\")", 1)
	} else {
		modelString = strings.Replace(modelString, "***", " ", 1)
	}
	data := []byte(modelString)
	ioutil.WriteFile(config.FilePath+config.FileName, data, 777)

	return nil
}
