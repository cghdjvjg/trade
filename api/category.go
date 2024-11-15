package api

import (
	"github.com/cghdjvjg/trade/pkg/ctl"
	"github.com/cghdjvjg/trade/pkg/util"
	"github.com/cghdjvjg/trade/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetCategoryService()
		resp, err := s.ShowCategory(c.Request.Context())
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))

	}
}
