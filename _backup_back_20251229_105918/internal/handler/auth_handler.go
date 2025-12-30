package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
)

// Login godoc
// @Summary      Login de usuário
// @Description  Autentica usuário e retorna token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.LoginRequest  true  "Credenciais de login"
// @Success      200      {object}  model.LoginResponse
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /auth/login [post]
func (h *Handlers) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Register godoc
// @Summary      Registrar novo usuário
// @Description  Cria uma nova conta de usuário com nome completo, email, senha e confirmação
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateUserRequest  true  "Dados do novo usuário"
// @Success      201      {object}  model.User
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /auth/register [post]
func (h *Handlers) Register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// RegisterAdmin godoc
// @Summary      Registrar novo administrador
// @Description  Cria uma conta de administrador usando um código secreto
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateAdminRequest  true  "Dados do novo administrador"
// @Success      201      {object}  model.User
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /auth/admin/register [post]
func (h *Handlers) RegisterAdmin(c *gin.Context) {
	var req model.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if h.adminSecret == "" || req.AdminSecret != h.adminSecret {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin secret"})
		return
	}

	userReq := model.CreateUserRequest{
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
		Confirm:  req.Confirm,
	}

	user, err := h.userService.CreateAdmin(&userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handlers) RefreshToken(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := h.authService.ValidateToken(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := h.userService.GetUserByEmail(claims.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	newToken, expiresIn, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      newToken,
		"expires_in": expiresIn,
	})
}
