package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zurihaqi.github.io-backend/internal/dto"
	"zurihaqi.github.io-backend/internal/model"
	"zurihaqi.github.io-backend/internal/service/interfaces"
	"zurihaqi.github.io-backend/internal/utils"
)

type ProjectHandler struct {
	Service interfaces.ProjectService
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.Service.FindAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.StringToUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest("invalid project ID"))
		return
	}

	project, err := h.Service.FindProjectByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project model.Project
	if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	createdProject, err := h.Service.CreateProject(&project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.Created("Project created successfully", createdProject))
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.StringToUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest("invalid project ID"))
		return
	}

	var project model.Project
	if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	updatedProject, err := h.Service.UpdateProject(id, &project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success("Project updated successfully", updatedProject))
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.StringToUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest("invalid project ID"))
		return
	}

	err = h.Service.DeleteProject(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success("Project deleted successfully", nil))
}
