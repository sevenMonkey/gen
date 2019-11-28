package file

import (
	"os"
	"strings"

	"github.com/sevenMonkey/gen/cmd/vars"
	"github.com/sevenMonkey/gen/util"
)

type ControllerFile struct {
	ModelName string
	Tpl       string
	Content   string
}

func NewControllerFile(modelName string) *ControllerFile {
	controllerFile := ControllerFile{}

	controllerFile.ModelName = modelName
	controllerFile.Tpl = getDefaultControllerTpl()

	return &controllerFile
}

func (m *ControllerFile) Write() error {
	filePath := "./controller/" + m.ModelName + "_controller.go"
	f, err := os.Create(filePath)
	m.Generate()

	f.Write([]byte(m.Tpl))

	util.FormatSourceCode(filePath)

	return err
}

func (m *ControllerFile) Generate() {
	m.Tpl = strings.Replace(m.Tpl, "{{modelName}}", m.ModelName, -1)
	m.Tpl = strings.Replace(m.Tpl, "{{ModelName}}", strings.Title(m.ModelName), -1)
	m.Tpl = strings.Replace(m.Tpl, "{{projectName}}", vars.ProjectName, -1)
}

func getDefaultControllerTpl() string {
	return `//generate by gen
package controller

import (
	"{{projectName}}/model"
	"{{projectName}}/pkg"
	"{{projectName}}/app/logging"
	"{{projectName}}/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type {{ModelName}}Controller struct {
}

// @Summary Create
// @Tags    {{ModelName}}
// @Param body body model.{{ModelName}} true "{{ModelName}}"
// @Success 200 {string} json ""
// @Router /{{modelName}}s [post]
func (ctl *{{ModelName}}Controller) Create(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{
	}

	if err := pkg.ParseRequest(c, &{{modelName}}); err != nil {
		return
	}

	if err := {{modelName}}.Insert(); err != nil {
		logging.Error("Create {{ModelName}}", err)
		app.RespJson(c, e.ERROR, nil)
		return
	}

	app.RespJson(c, e.SUCCESS, {{modelName}})
}

// @Summary  Delete
// @Tags     {{ModelName}}
// @Param  {{modelName}}Id  path string true "{{modelName}}Id"
// @Success 200 {string} json ""
// @Router /{{modelName}}s/{{{modelName}}Id} [delete]
func (ctl *{{ModelName}}Controller) Delete(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")
	err := {{modelName}}.Delete()

	if err != nil {
		logging.Error("Delete {{ModelName}}", err)
		app.RespJson(c, e.ERROR, nil)
		return
	}

	app.RespJson(c, e.SUCCESS, nil)
}

// @Summary Put
// @Tags    {{ModelName}}
// @Param body body model.{{ModelName}} true "{{modelName}}"
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {string} json ""
// @Router /{{modelName}}s/{{{modelName}}Id} [put]
func (ctl *{{ModelName}}Controller) Put(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")

	if err := pkg.ParseRequest(c, &{{modelName}}); err != nil {
		return
	}

	err := {{modelName}}.Update()
	if err != nil {
		logging.Error("Put {{ModelName}}", err)
		app.RespJson(c, e.ERROR, nil)
		return
	}

	app.RespJson(c, e.SUCCESS, nil)
}

// @Summary Patch
// @Tags    {{ModelName}}
// @Param body body model.{{ModelName}} true "{{modelName}}"
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {string} json ""
// @Router /{{modelName}}s/{{{modelName}}Id} [patch]
func (ctl *{{ModelName}}Controller) Patch(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")

	if err := pkg.ParseRequest(c, &{{modelName}}); err != nil {
		return
	}

	err := {{modelName}}.Patch()
	if err != nil {
		logging.Error("Patch {{ModelName}}", err)
		app.RespJson(c, e.ERROR, nil)
		return
	}

	app.RespJson(c, e.SUCCESS, nil)
}

// @Summary List
// @Tags    {{ModelName}}
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} model.{{ModelName}} "{{modelName}} array"
// @Router /{{modelName}}s [get]
func (ctl *{{ModelName}}Controller) List(c *gin.Context) {
	{{modelName}} := &model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")
	var err error

	pageParam := c.DefaultQuery("page", "-1")
	pageSizeParam := c.DefaultQuery("pageSize", "-1")
	rawQuery := c.DefaultQuery("query", "")
	rawOrder := c.DefaultQuery("order", "")

	pageInt, err := strconv.Atoi(pageParam)
	pageSizeInt, err := strconv.Atoi(pageSizeParam)

	offset := pageInt*pageSizeInt - pageSizeInt
	limit := pageSizeInt

	if pageInt < 0 || pageSizeInt < 0 {
		limit = -1
	}

	{{modelName}}s, total, err := {{modelName}}.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		logging.Error("List {{ModelName}}", err)
		app.RespJson(c, e.ERROR, nil)
		return
	}
	app.RespJson(c, http.StatusOK, gin.H{
		"total": total,
		"list": {{modelName}}s,
	})
}


// @Summary Get
// @Tags    {{ModelName}}
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {object} model.{{ModelName}} "{{modelName}} object"
// @Router /{{modelName}}s/{{{modelName}}Id} [get]
func (ctl *{{ModelName}}Controller) Get(c *gin.Context) {
	{{modelName}} := &model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")

	{{modelName}}, err := {{modelName}}.Get()
	if err != nil {
		logging.Error("Get {{ModelName}}", err)
		app.RespJson(c, e.ERROR, nil)
		return
	}

	app.RespJson(c, http.StatusOK, {{modelName}})
}
`
}
