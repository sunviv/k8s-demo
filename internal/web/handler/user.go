package handler

import (
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sunviv/k8s-demo/internal/domain"
	"github.com/sunviv/k8s-demo/internal/service"
	"github.com/sunviv/k8s-demo/internal/web/middleware"
	"net/http"
)

const (
	emailRegexPattern    = `[\w.-]+@[\w_-]+\w{1,}[\.\w-]+`
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRegexp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	server.POST("/signIn", h.SignInWithSession)
	server.POST("/signUp", h.SignUp)
	ug := server.Group("/users", middleware.AuthnMiddlewareBuilder{}.Authn())
	ug.GET("", h.Profile)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	type Request struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return
	}
	isValidEmail, err := h.emailRegexp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isValidEmail {
		ctx.String(http.StatusOK, "邮箱格式错误")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}

	isValidPassword, err := h.passwordRegexp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isValidPassword {
		ctx.String(http.StatusOK, "密码必须包含字母、数字、特殊字符且不少于8位")
		return
	}

	err = h.svc.SignUp(ctx, domain.User{Email: req.Email, Password: req.Password})
	switch {
	case err == nil:
		ctx.String(http.StatusOK, "注册成功")
	case errors.Is(err, service.ErrEmailDuplicate):
		ctx.String(http.StatusOK, err.Error())
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

}

func (h *UserHandler) SignInWithSession(ctx *gin.Context) {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := h.svc.SignIn(ctx, req.Email, req.Password)
	switch {
	case err == nil:
		sess := sessions.Default(ctx)
		sess.Set("userID", user.ID)
		sess.Options(sessions.Options{
			MaxAge: 60 * 60,
		})
		err = sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}
		ctx.String(http.StatusOK, "登录成功")
	case errors.Is(err, service.ErrUserOrPasswordInvalid):
		ctx.String(http.StatusOK, err.Error())
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "*********Profile********")
}
