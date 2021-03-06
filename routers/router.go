package routers

import (
	"fmt"
	"gin/middleware/pfrom"
	"gin/pkg/setting"
	"gin/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine  {

	r :=gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(pfrom.Pid())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		c.JSON(200,gin.H{
			"message": "test2",
		})
	})

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		apiv1.GET("/tags/export", v1.ExportTag)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	qst := r.Group("qst/api/v1")
	{
		qst.GET("/tags", v1.GetTags)
	}


	return r
}