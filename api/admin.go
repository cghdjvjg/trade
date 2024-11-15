package api

import (
	"github.com/cghdjvjg/gin"
	"github.com/cghdjvjg/service"
	"github.com/cghdjvjg/types"
	"github.com/cghdjvjg/util"
	"net/http"
	"strconv"
)

func ShowAllAdminsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNumStr := c.DefaultQuery("pageNum", "1")
		pageSizeStr := c.DefaultQuery("pageSize", "10")
		searchQuery := c.DefaultQuery("searchQuery", "")

		pageNum, err := strconv.Atoi(pageNumStr)
		if err != nil {
			util.LogrusObj.Infoln("Invalid pageNum:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			util.LogrusObj.Infoln("Invalid pageSize:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		req := &types.GetAdminRequest{
			PageNum:     pageNum,
			PageSize:    pageSize,
			SearchQuery: searchQuery,
		}

		s := service.GetAdminService()
		resp, err := s.ShowAllAdmins(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, RespSuccess(c, resp))
	}
}

func AddAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.AddAdminRequest
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetAdminService()
		resp, err := s.AddAdmin(c.Request.Context(), &req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, RespSuccess(c, resp))
	}
}

func UpdateAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UpdateAdminRequest
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		adminIDStr := c.Param("id")
		adminID, err := strconv.Atoi(adminIDStr)
		if err != nil {
			util.LogrusObj.Infoln("Invalid adminID:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		req.AdminID = adminID

		s := service.GetAdminService()
		resp, err := s.UpdateAdmin(c.Request.Context(), &req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, RespSuccess(c, resp))
	}
}

func DeleteAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminIDStr := c.Param("id")
		adminID, err := strconv.Atoi(adminIDStr)
		if err != nil {
			util.LogrusObj.Infoln("Invalid adminID:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		req := &types.DeleteAdminRequest{
			AdminID: adminID,
		}

		s := service.GetAdminService()
		resp, err := s.DeleteAdmin(c.Request.Context(), req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, RespSuccess(c, resp))
	}
}
