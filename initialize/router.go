package initialize

import (
	"github.com/gin-gonic/gin"
	"groqai2api/middlewares"
	"groqai2api/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors)

	Router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "https://niansuhai-llms.hf.space")
	})

	Router.GET("/ping", func(c *gin.Context) {
		c.String(200, "(.•snort! It's all your fault\n(`ȏ´) I don't want to coax you either\n(〃′o`) I want to cry so much that I beat your chest, you bad guy! ! ! \n(｡•ˇ‸ˇ•｡)Hmph! It's all your fault\n(`ȏ´)I don't want to coax anyone either\n(〃′o`)I want to cry so much, I'll beat your chest, you bad guy! ! ! \n(￣^￣)ゞ咩QAQ hits your chest, you are so annoying! \n(￣^￣)ゞ咩QAQ hits your chest, you are so annoying! \n(=ﾟωﾟ)ﾉIf you want to give me a hug, let me hug you, if you want to give me a hug, I will hit you on the chest with a small fist! ! ! \n(=ﾟωﾟ)ﾉIf you want to give me a hug, let me hug you, if you want to give me a hug, I will hit you on the chest with a small fist! ! ! \n(｡• ︿•̀｡)Big bad guy, I’ll beat you to death(つд⊂)")
	})
	v1Group := Router.Group("/v1/")
	router.InitRouter(v1Group)

	return Router
}
