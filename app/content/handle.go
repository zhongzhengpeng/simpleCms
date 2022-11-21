package content

import (
	"log"
	"net/http"
	"simpleCms/app/common"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {
	id := ctx.Param("id")
	row := &Content{}

	if result := common.DB.Debug().First(row, id); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    row,
	})

}

func POST(ctx *gin.Context) {
	row := &Content{}
	row.Status = "new"

	if result := common.DB.Create(row); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "内容添加失败",
		})
		return
	}

	ctx.JSON(
		http.StatusOK, gin.H{
			"code":    0,
			"message": "",
			"data":    row,
		})
}

func PUT(ctx *gin.Context) {
	id := ctx.Param("id")
	body := Work{}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"messafe": "内容解析失败",
		})
	}
	row := &Content{}
	common.DB.First(&row, id)
	row.Work = body
	if result := common.DB.Save(row); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "更新数据失败",
		})
		return
	}

	ctx.JSON(
		http.StatusOK, gin.H{
			"code":    0,
			"message": "",
			"data":    row,
		})
}

func DELETE(ctx *gin.Context) {
	id := ctx.Param("id")
	if result := common.DB.Delete(&Content{}, id); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "内容删除",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"massage": "",
	})
}
