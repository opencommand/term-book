package handler

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"termbook/api"
	"termbook/config"
	"termbook/model"

	"github.com/gin-gonic/gin"
)

// TODO 通过绝对路径打开文件
// OpenDocument godoc
// @Summary 打开文档
// @Description 从文件中读取文档内容
// @Tags Document
// @Accept json
// @Produce json
// @Param request body api.OpenDocumentRequest true "文档文件名"
// @Success 200 {object} api.ResponseOpenDocument
// @Failure 400 {object} api.ResponseBad
// @Failure 404 {object} api.ResponseBad
// @Failure 500 {object} api.ResponseBad
// @Router /document/open [post]
func OpenDocument(ctx *gin.Context) {
	var req api.OpenDocumentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	raw, err := os.ReadFile(filepath.Join(config.DataPath, req.Filename+".json"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "文件读取失败"})
		return
	}
	var doc model.Document
	if err := json.Unmarshal(raw, &doc); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "JSON 解析失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "载入成功", "data": doc})
}

// SaveDocument godoc
// @Summary 保存文档
// @Description 将当前文档保存到文件中
// @Tags Document
// @Accept json
// @Produce json
// @Param request body api.SaveDocumentRequest true "要保存的文档内容"
// @Success 200 {object} api.ResponseBad
// @Failure 400 {object} api.ResponseBad
// @Failure 500 {object} api.ResponseBad
// @Router /document/save [post]
func SaveDocument(ctx *gin.Context) {
	var req api.SaveDocumentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	data, err := json.Marshal(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "JSON 序列化失败"})
		return
	}
	err = os.WriteFile(filepath.Join(config.DataPath, req.Filename+".json"), data, 0666)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "文件保存失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "保存成功"})
}

// NewDocument godoc
// @Summary 新建文档
// @Description 创建一个空文档并保存
// @Tags Document
// @Accept json
// @Produce json
// @Param request body api.NewDocumentRequest true "新建文档请求"
// @Success 200 {object} api.ResponseNewDocument
// @Failure 400 {object} api.ResponseBad
// @Failure 500 {object} api.ResponseBad
// @Router /document/new [post]
func NewDocument(ctx *gin.Context) {
	var req api.NewDocumentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	doc := model.Document{Filename: req.Filename}
	data, err := json.Marshal(&doc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "JSON 序列化失败"})
		return
	}
	err = os.WriteFile(filepath.Join(config.DataPath, doc.Filename+".json"), data, 0666)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "文件写入失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "创建成功", "data": doc})
}

// ListDocument godoc
// @Summary 获取所有文档文件名
// @Description 获取所有可打开的文档
// @Tags Document
// @Produce json
// @Success 200 {object} api.ListDocumentResponse "查询成功返回文档名数组"
// @Failure 500 {object} api.ResponseBad "目录遍历失败"
// @Router /document/list [get]
func ListDocument(ctx *gin.Context) {
	files := []string{}
	err := filepath.WalkDir(config.DataPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".json") {
			filename := filepath.Base(path)
			nameWithoutExt := strings.TrimSuffix(filename, ".json")
			files = append(files, nameWithoutExt)
		}
		return nil
	})

	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "目录遍历失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": files,
	})
}
