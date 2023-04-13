
package api
import (
	"xieyuhua/config"
	"xieyuhua/util/jwt"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"time"
	"net/http"
	"strings"
)

// GRouter 全局Router 全局调用的函数会在模块的init()之前执行
var GRouter *gin.Engine  
 
// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	// 配置允许的域名   * 代表所有都允许
	config.AllowOrigins = []string{"http://localhost:8080", "*"}
	config.AllowCredentials = true
	return cors.New(config)
}

type CORSOptions struct {
	Origin string
}

func jwtAbort(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": msg,
	})
	c.Abort()
}

// AuthRequired 需要登录
func AuthRequired(options CORSOptions) gin.HandlerFunc {

	return func(c *gin.Context) {

// 		c.Next()
// 		return 
		//跨域处理
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow any origin domain
		if options.Origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", options.Origin)
		}
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		//token验证
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			jwtAbort(c, "Authorization Failed.")
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "gin") {
			//jwtAbort(c, "Authorization Failed..")
			//return
		}
		// claims, err := jwt.ParseToken(parts[1])
		claims, err := jwt.ParseToken(authHeader)
		if err != nil {
			jwtAbort(c, "无效的Token"+authHeader)
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			jwtAbort(c, "Token已过期")
			return
		}

		//验证账户信息正确性

		c.Next()
	}
}


// GinInit ...
func GinInit(){
	log.Println("启动 gin http服务 :", config.GConf.ServerPort)

	if config.GConf.OutLog {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		os.Mkdir(dir+"/log", 0777)
		f, e := os.OpenFile(dir+"/log/ginlog"+time.Now().Format("2006-01-02 15-04-05")+".log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
		if e != nil {
			log.Println("日志文件开启失败：", e)
		} else {
			gin.DefaultWriter = f
			gin.DefaultErrorWriter = f
		}
	}

	gin.SetMode(gin.ReleaseMode)

	GRouter = gin.Default()
	ginpprof.Wrapper(GRouter)
	 
	groupgo := GRouter.Group("/go")
	// 使用跨域中间件允许跨域
	groupgo.Use(Cors())


	// 加载各个router
	baseInit(groupgo)


    //需要token
	groupgo.Use(AuthRequired(CORSOptions{}))
	initrouter(groupgo)
	
	//启动服务
	go func() {
		if err := GRouter.Run(fmt.Sprintf(":%d", config.GConf.ServerPort)); err != nil {
			log.Println(err.Error())
		}
	}()
	 
}

// 加载各个router
func initrouter(groupgo *gin.RouterGroup) {
// 	baseInit(groupgo)
	zoomdeployapplyInit(groupgo)
    zoomdeploybuildInit(groupgo)
    zoomdeploytaskInit(groupgo)
    zoomprojectInit(groupgo)
    zoomprojectmemberInit(groupgo)
    zoomprojectspaceInit(groupgo)
    zoomserverInit(groupgo)
    zoomservergroupInit(groupgo)
    zoomuserInit(groupgo)
    zoomuserroleInit(groupgo)
    zoomusertokenInit(groupgo)
    
 
}
