package v1

import (
	"fmt"
	"gin/models"
	"gin/pkg/e"
	"gin/pkg/export"
	"gin/pkg/setting"
	"gin/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = 1

	if arg:=c.Query("state");arg !=""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	fmt.Println(maps)
	data["lists"] = models.GetTags(util.GetPage(c),setting.AppSetting.PageSize,maps)
	data["total"] = models.GetTagTotal(maps)

	fmt.Println(c.Get("PF"))
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}

func ExportTag(c *gin.Context)  {

	maps := make(map[string]interface{})
	maps["state"] = 1
	data := models.GetTags(util.GetPage(c),setting.AppSetting.PageSize,maps)


	var exportData [][]string
	for _,v := range data {
		var tmp []string
		tmp = append(tmp,v.Name,com.Int2HexStr(v.State))
		exportData = append(exportData,tmp)
	}

	filePath,err := export.ExportCsv("lupeng",exportData)

	if err !=nil {
		panic(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":e.SUCCESS,
		"msg":"aa",
		"filePath":filePath,
	})
}