package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
)

func (h *Handlers) getUserIDFromRequest(c *gin.Context) (uuid.UUID, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return uuid.Nil, false
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		return uuid.Nil, false
	}

	claims, err := h.authService.ValidateToken(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return uuid.Nil, false
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user id"})
		return uuid.Nil, false
	}

	return userID, true
}

// GetMyModules godoc
// @Summary      Listar modulos
// @Tags         meus-cursos
// @Produce      json
// @Success      200 {array} model.CourseModule
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/modulos [get]
func (h *Handlers) GetMyModules(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	modules, err := h.courseService.GetMyModules(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, modules)
}

// GetCourses godoc
// @Summary      Listar cursos
// @Tags         cursos
// @Produce      json
// @Success      200 {array} model.Course
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos [get]
func (h *Handlers) GetCourses(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	courses, err := h.courseService.GetMyCourses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetCourseCategories godoc
// @Summary      Listar categorias
// @Tags         categorias
// @Produce      json
// @Success      200 {array} model.CourseCategory
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos/categorias [get]
func (h *Handlers) GetCourseCategories(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	categories, err := h.courseService.GetMyCategories(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// CreateCourse godoc
// @Summary      Criar curso
// @Tags         cursos
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCourseRequest true "Curso"
// @Success      201 {object} model.Course
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos [post]
func (h *Handlers) CreateCourse(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	var req model.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.courseService.CreateCourse(userID, &req)
	if err != nil {
		if err.Error() == "module not found" || err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

// UpdateCourse godoc
// @Summary      Atualizar curso
// @Tags         cursos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID do curso"
// @Param        request body model.UpdateCourseRequest true "Curso"
// @Success      200 {object} model.Course
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos/{id} [put]
func (h *Handlers) UpdateCourse(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	courseID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.courseService.UpdateCourse(userID, courseID, &req)
	if err != nil {
		if err.Error() == "category not found" || err.Error() == "module not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

// DeleteCourse godoc
// @Summary      Remover curso
// @Tags         cursos
// @Param        id path string true "ID do curso"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos/{id} [delete]
func (h *Handlers) DeleteCourse(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	courseID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.courseService.DeleteCourse(userID, courseID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateCourseCategory godoc
// @Summary      Criar categoria
// @Tags         categorias
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCourseCategoryRequest true "Categoria"
// @Success      201 {object} model.CourseCategory
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos/categorias [post]
func (h *Handlers) CreateCourseCategory(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	var req model.CreateCourseCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.courseService.CreateCategory(userID, &req)
	if err != nil {
		if err.Error() == "course not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCourseCategory godoc
// @Summary      Atualizar categoria
// @Tags         categorias
// @Accept       json
// @Produce      json
// @Param        id path string true "ID da categoria"
// @Param        request body model.UpdateCourseCategoryRequest true "Categoria"
// @Success      200 {object} model.CourseCategory
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos/categorias/{id} [put]
func (h *Handlers) UpdateCourseCategory(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateCourseCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.courseService.UpdateCategory(userID, categoryID, &req)
	if err != nil {
		if err.Error() == "course not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCourseCategory godoc
// @Summary      Remover categoria
// @Tags         categorias
// @Param        id path string true "ID da categoria"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cursos/categorias/{id} [delete]
func (h *Handlers) DeleteCourseCategory(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.courseService.DeleteCategory(userID, categoryID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateCourseModuleStandalone godoc
// @Summary      Criar modulo
// @Tags         meus-cursos
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCourseModuleRequest true "Modulo"
// @Success      201 {object} model.CourseModule
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/modulos [post]
func (h *Handlers) CreateCourseModuleStandalone(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	var req model.CreateCourseModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module, err := h.courseService.CreateModule(userID, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, module)
}

// UpdateCourseModule godoc
// @Summary      Atualizar modulo
// @Tags         meus-cursos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID do modulo"
// @Param        request body model.UpdateCourseModuleRequest true "Modulo"
// @Success      200 {object} model.CourseModule
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/modulos/{id} [put]
func (h *Handlers) UpdateCourseModule(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	moduleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateCourseModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module, err := h.courseService.UpdateModule(userID, moduleID, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, module)
}

// DeleteCourseModule godoc
// @Summary      Remover modulo
// @Tags         meus-cursos
// @Param        id path string true "ID do modulo"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/modulos/{id} [delete]
func (h *Handlers) DeleteCourseModule(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	moduleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.courseService.DeleteModule(userID, moduleID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetMyItems godoc
// @Summary      Listar itens
// @Tags         meus-cursos
// @Produce      json
// @Success      200 {array} model.CourseItem
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/itens [get]
func (h *Handlers) GetMyItems(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	items, err := h.courseService.GetMyItems(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// CreateCourseItemStandalone godoc
// @Summary      Criar item sem modulo
// @Tags         meus-cursos
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCourseItemRequest true "Item"
// @Success      201 {object} model.CourseItem
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/itens [post]
func (h *Handlers) CreateCourseItemStandalone(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	var req model.CreateCourseItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.courseService.CreateItem(userID, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateCourseItem godoc
// @Summary      Atualizar item
// @Tags         meus-cursos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID do item"
// @Param        request body model.UpdateCourseItemRequest true "Item"
// @Success      200 {object} model.CourseItem
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/itens/{id} [put]
func (h *Handlers) UpdateCourseItem(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateCourseItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.courseService.UpdateItem(userID, itemID, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteCourseItem godoc
// @Summary      Remover item
// @Tags         meus-cursos
// @Param        id path string true "ID do item"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meus-cursos/itens/{id} [delete]
func (h *Handlers) DeleteCourseItem(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.courseService.DeleteItem(userID, itemID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
