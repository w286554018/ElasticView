// 中间件层
package middleware

import (
	"time"

	"github.com/1340691923/ElasticView/pkg/jwt"
	"github.com/1340691923/ElasticView/pkg/my_error"
	"github.com/1340691923/ElasticView/pkg/response"
	"github.com/1340691923/ElasticView/pkg/util"
	"github.com/1340691923/ElasticView/service/gm_user"
	fiber "github.com/gofiber/fiber/v2"
)

var res response.Response

func JwtMiddleware(c *fiber.Ctx) error {

	var err error
	var claims *jwt.Claims
	token := util.GetToken(c)
	if _, logoff := util.TokenBucket.Load(token); logoff {
		err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_FAIL)
		return res.Error(c, err)
	}

	if util.GetToken(c) == "" {
		err = my_error.NewBusiness(TOKEN_ERROR, INVALID_PARAMS)
		return res.Error(c, err)
	} else {

		var service gm_user.GmUserService
		claims, err = jwt.ParseToken(token)
		if err != nil {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_FAIL)
			return res.Error(c, err)
		} else if time.Now().Unix() > claims.ExpiresAt {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			return res.Error(c, err)
		} else {
			isExitUser, err := service.IsExitUser(claims)
			if err != nil {
				return res.Error(c, err)
			}
			if !isExitUser {
				err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
				return res.Error(c, err)
			}
		}
	}
	return c.Next()

}
