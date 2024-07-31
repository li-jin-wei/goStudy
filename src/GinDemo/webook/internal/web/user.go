package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

// NewUserHandle 使用正则表达式规范邮箱和密码
func NewUserHandle() *UserHandle {
	const (
		//匹配邮箱格式
		emailRegexPattern = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"

		//密码至少是8位，必须包含至少一个数字，一个字母，一个特殊字符
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)

	//使用regexp.MustCompile函数将定义的正则表达式字符串编译成*regexp.Regexp类型的实例
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)

	//返回UserHandle结构体指针
	return &UserHandle{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}

}

func (u *UserHandle) RegisterRoutes(server *gin.Engine) {
	userGroup := server.Group("/user")

	//查询用户信息接口
	userGroup.GET("/profile", u.Profile)

	//注册接口
	userGroup.POST("/signup", u.SignUp)

	//登录接口
	userGroup.POST("/login", u.Login)

	//退出接口
	userGroup.POST("/logout", u.Logout)

	//修改用户信息接口
	userGroup.POST("/edit", u.Edit)
}

func (u *UserHandle) Profile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "profile",
	})
}

// SignUp 接口请求数据
// 使用Bind方法接受请求
// Bind 方法会根据HTTP请求的Content-Type来处理请求
// 请求如果是JSON，Content-Type是application/json，那么Gin就会使用JSON反序列化
// 如果Bind方法发现输入有问题，就会直接返回一个错误响应到前端
// 如果输入无效，它会写入400错误并在响应中设置Content-Type标头“text/plain”
func (u *UserHandle) SignUp(c *gin.Context) {
	type SignUpTRequest struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"` //确认密码
	}
	var request SignUpTRequest
	if err := c.Bind(&request); err != nil {
		return
	}

	//邮箱格式校验
	OK, err := u.emailExp.MatchString(request.Email)
	if err != nil {
		c.String(400, "系统错误")
	}
	if !OK {
		c.String(400, "邮箱格式错误")
		return
	}

	//密码复杂度校验
	OK, err = u.passwordExp.MatchString(request.Password)
	if err != nil {
		c.String(400, "系统错误")
	}
	if !OK {
		c.String(400, "密码至少是8位，必须包含至少一个数字，一个字母，一个特殊字符")
	}

	//密码确认校验
	if request.Password != request.ConfirmPassword {
		c.String(400, "两次密码不一致")
		return
	}

	c.String(200, "注册成功")
	fmt.Printf("请求体为%v", request)
}
func (u *UserHandle) Login(c *gin.Context) {

}
func (u *UserHandle) Logout(c *gin.Context) {

}
func (u *UserHandle) Edit(c *gin.Context) {

}
