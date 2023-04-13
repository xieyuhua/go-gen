
package api
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xieyuhua/util/jwt"
)

// 加载路由
func baseInit(groupgo *gin.RouterGroup) {
	  
	// 所有api的开头 都加上/go  例如 /go/ping
	groupgo.GET("ping", Ping)
    groupgo.GET("token", Token)
}

// Ping 测试存活
func Token(c *gin.Context) {
	reply,_ := jwt.GenerateToken(121545454,"1510120461@qq.com")
	c.JSON(http.StatusOK, reply)
}

// Ping 测试存活
func Ping(c *gin.Context) {
	reply := NewReplyOk()
	c.JSON(http.StatusOK, reply)
}

