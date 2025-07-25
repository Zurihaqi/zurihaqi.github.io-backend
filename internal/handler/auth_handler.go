package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zurihaqi.github.io-backend/internal/dto"
	"zurihaqi.github.io-backend/internal/service/interfaces"
)

type AuthHandler struct {
	Service interfaces.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	user, err := h.Service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalError(err.Error()))
		return
	}

	user.Password = ""
	c.JSON(http.StatusCreated, dto.Created("User created successfully", user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Unauthorized(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success("Login successful", gin.H{"token": token}))
}
