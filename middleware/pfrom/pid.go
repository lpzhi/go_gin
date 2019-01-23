//设置全局变量pfrom_id
//加载对应配置文件
package pfrom

import (
	"github.com/gin-gonic/gin"
	"strings"
)


func Pid()  gin.HandlerFunc{
	return func(c *gin.Context) {
		//var code int
		//code = e.SUCCESS
		data := strings.Split( c.Request.URL.Path,"/")
		c.Set("PF",data[1])
		c.Next()
	}
}