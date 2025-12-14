package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
)

// SocialLogin godoc
// @Summary      Login com provedor social
// @Description  Autentica usuário usando Google ou Facebook OAuth2
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.SocialAuthRequest  true  "Token de acesso do provedor"
// @Success      200      {object}  model.LoginResponse
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /auth/social [post]
func (h *Handlers) SocialLogin(c *gin.Context) {
	var req model.SocialAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.socialAuthService.AuthenticateWithToken(req.Provider, req.AccessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Gerar token JWT
	token, expiresIn, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		Token:     token,
		ExpiresIn: expiresIn,
		User:      *user,
	})
}

// GoogleAuthURL godoc
// @Summary      Obter URL de autenticação do Google
// @Description  Retorna URL para redirecionar usuário ao login do Google
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /auth/google/url [get]
func (h *Handlers) GoogleAuthURL(c *gin.Context) {
	state := "random-state-token" // Em produção, use um token aleatório e salve na sessão
	url := h.socialAuthService.GetGoogleAuthURL(state)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// GoogleCallback godoc
// @Summary      Callback do Google OAuth
// @Description  Processa retorno do Google após autorização
// @Tags         auth
// @Produce      json
// @Param        code   query     string  true  "Authorization code"
// @Param        state  query     string  true  "State token"
// @Success      200    {object}  model.LoginResponse
// @Failure      400    {object}  map[string]string
// @Router       /auth/google/callback [get]
func (h *Handlers) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	// state := c.Query("state") // Validar state em produção

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
		return
	}

	user, err := h.socialAuthService.AuthenticateWithGoogle(code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, expiresIn, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		Token:     token,
		ExpiresIn: expiresIn,
		User:      *user,
	})
}

// FacebookAuthURL godoc
// @Summary      Obter URL de autenticação do Facebook
// @Description  Retorna URL para redirecionar usuário ao login do Facebook
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /auth/facebook/url [get]
func (h *Handlers) FacebookAuthURL(c *gin.Context) {
	state := "random-state-token" // Em produção, use um token aleatório e salve na sessão
	url := h.socialAuthService.GetFacebookAuthURL(state)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// FacebookCallback godoc
// @Summary      Callback do Facebook OAuth
// @Description  Processa retorno do Facebook após autorização
// @Tags         auth
// @Produce      json
// @Param        code   query     string  true  "Authorization code"
// @Param        state  query     string  true  "State token"
// @Success      200    {object}  model.LoginResponse
// @Failure      400    {object}  map[string]string
// @Router       /auth/facebook/callback [get]
func (h *Handlers) FacebookCallback(c *gin.Context) {
	code := c.Query("code")
	// state := c.Query("state") // Validar state em produção

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
		return
	}

	user, err := h.socialAuthService.AuthenticateWithFacebook(code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, expiresIn, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		Token:     token,
		ExpiresIn: expiresIn,
		User:      *user,
	})
}
