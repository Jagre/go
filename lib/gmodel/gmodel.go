package gmodel

import (
	"fmt"
	"strings"
)

//IGModel defined the methods of construct model
type IGModel interface {
	GetModels() *[]GModel
	Output() error
}

//IGModelTypeConvertor defined type convert
type IGModelTypeConvertor interface {
	GetGModelType(sqlType string) string
}

//IGModelGenerator defined Model gnerator
type IGModelGenerator interface {
	GModelGenerate() string
}

//GModelProperty defined coulumns' property
type GModelProperty struct {
	Name         string
	OriginalType string
	GoType       string
}

//GModel defined the model's struct
type GModel struct {
	ModelName  string
	Properties []*GModelProperty
}

//GModelGenerate will generate model struct string
func (m *GModel) GModelGenerate() string {
	s := fmt.Sprintf("type %s struct{\n", m.ModelName)
	for _, p := range m.Properties {
		name := p.Name
		if strings.ToLower(name) == "id" {
			name = "ID"
		}
		s += fmt.Sprintf("%s %s `gorm:\"column:%s\"`\n", name, p.GoType, p.Name)
	}
	return s + fmt.Sprintf("}\n")
}
