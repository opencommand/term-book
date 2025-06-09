package handler

import (
	"net/http"
	"os/exec"
	"strings"
	"termbook/api"
	"time"

	"github.com/gin-gonic/gin"
)

// RunCell godoc
// @Summary 执行单元格命令
// @Description 执行输入的命令并返回输出结果
// @Tags Cell
// @Accept json
// @Produce json
// @Param request body api.RunCellRequest true "命令输入"
// @Success 200 {object} api.ResponseRunCell
// @Failure 400 {object} api.ResponseBad
// @Router /run-cell [post]
func RunCell(ctx *gin.Context) {
	var req api.RunCellRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": "请求参数错误"})
		return
	}
	start := time.Now()
	cmd := exec.Command("cmd", "/C", req.Input)
	raw, err := cmd.CombinedOutput()
	timeElapsed := time.Since(start)
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")

	res := api.RunCellResponse{
		Output:     data,
		OutputTime: time.Now(),
		ExecTime:   timeElapsed.String(),
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "执行失败", "data": res})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "执行成功", "data": res})
}
