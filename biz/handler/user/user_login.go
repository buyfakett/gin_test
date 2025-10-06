package user

import (
	"gin_template/biz/dal"
	"gin_template/biz/response"
	"gin_template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username   string `json:"username" binding:"required,min=1,max=255"`
	Password   string `json:"password" binding:"required,min=1,max=255"`
	RememberMe bool   `json:"remember_me"`
}

type LoginData struct {
	Token string `json:"token"`
}

type LoginResp struct {
	Code response.Code `json:"code"`
	Msg  string        `json:"msg"`
	Data *LoginData    `json:"data"`
}

// UserLogin 用户登录
// @Tags 用户
// @Summary 用户登录
// @Description 用户登录
// @Accept application/json
// @Produce application/json
// @Param req body LoginReq true "登录凭证"
// @Success 200 {object} LoginResp
// @router /api/user/login [POST]
func UserLogin(c *gin.Context) {
	req := new(LoginReq)
	if err := c.ShouldBind(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	resp := new(LoginResp)

	userData, err := dal.UserLogin(req.Username)
	if err != nil {
		c.JSON(http.StatusOK, &LoginResp{Code: response.Code_DBErr, Msg: err.Error()})
		return
	}

	if userData.Password != utils.MD5(req.Password) {
		c.JSON(http.StatusOK, &LoginResp{Code: response.Code_PasswordErr, Msg: "密码错误"})
		return
	}

	var token string
	if req.RememberMe {
		token, _ = utils.GenerateToken(userData.ID, req.Username)
	} else {
		//如果没有选记住我就1小时token
		token, _ = utils.GenerateToken(userData.ID, req.Username, 1)
	}

	resp.Code = response.Code_Success
	resp.Msg = "登录成功"
	resp.Data = &LoginData{
		Token: token,
	}
	c.JSON(http.StatusOK, resp)
}
