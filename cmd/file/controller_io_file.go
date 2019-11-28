package file

import (
	"github.com/sevenMonkey/gen/util"
	"os"
	"strings"
)

type ControllerIoFile struct {
	ModelName string
	Fields []string
	Tpl       string
	Content   string
}

func NewControllerIoFile(modelName string, fields []string) *ControllerIoFile {
	ioFile := ControllerIoFile{}

	ioFile.ModelName = modelName
	ioFile.Fields = fields

	ioFile.Tpl = getDefaultControllerIoTpl()

	return &ioFile
}

func (m *ControllerIoFile) Write() error {
	filePath := "./controller/" + m.ModelName + "_io.go"
	f, err := os.Create(filePath)
	m.Generate()

	f.Write([]byte(m.Tpl))

	util.FormatSourceCode(filePath)

	return err
}

func (m *ControllerIoFile) Generate() {
	structText := ""
	updateText := ""

	for _, field := range m.Fields {
		fieldSplit := strings.Split(field, ":")
		fieldName := fieldSplit[0]
		fieldType := fieldSplit[1]
		oneLineText := "	" + strings.Title(fieldName) + "    " + fieldType + "    " + "`json:" + "\"" + fieldName + "\"" + "  binding:\"required\"" + "`"
		structText += oneLineText
		structText += `
`
	}

	upFields := append([]string{"id:string"}, m.Fields...)
	for _, field := range upFields {
		fieldSplit := strings.Split(field, ":")
		fieldName := fieldSplit[0]
		fieldType := fieldSplit[1]
		oneLineText := "	" + strings.Title(fieldName) + "    " + fieldType + "    " + "`json:" + "\"" + fieldName + "\"" + "  binding:\"required\"" + "`"
		updateText += oneLineText
		updateText += `
`
	}
	m.Tpl = strings.Replace(m.Tpl, "{{ModelName}}", strings.Title(m.ModelName), -1)
	m.Tpl = strings.Replace(m.Tpl, "{{createStruct}}", structText, -1)
	m.Tpl = strings.Replace(m.Tpl, "{{patchStruct}}", structText, -1)
	m.Tpl = strings.Replace(m.Tpl, "{{updateStruct}}", updateText, -1)

}

func getDefaultControllerIoTpl() string {
	return `//generate by gen
package controller
	type create{{ModelName}} struct{
		{{createStruct}}
    }

	type update{{ModelName}} struct{
		{{updateStruct}}
	}

	type patch{{ModelName}} struct{
		{{patchStruct}}
	}
`
}
