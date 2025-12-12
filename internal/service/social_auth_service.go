package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/thepantheon/api/internal/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type SocialAuthService struct {
	userService    *UserService
	googleConfig   *oauth2.Config
	facebookConfig *oauth2.Config
}

func NewSocialAuthService(userService *UserService, googleClientID, googleClientSecret, facebookAppID, facebookAppSecret, redirectURL string) *SocialAuthService {
	return &SocialAuthService{
		userService: userService,
		googleConfig: &oauth2.Config{
			ClientID:     googleClientID,
			ClientSecret: googleClientSecret,
			RedirectURL:  redirectURL + "/auth/google/callback",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		facebookConfig: &oauth2.Config{
			ClientID:     facebookAppID,
			ClientSecret: facebookAppSecret,
			RedirectURL:  redirectURL + "/auth/facebook/callback",
			Scopes:       []string{"email", "public_profile"},
			Endpoint:     facebook.Endpoint,
		},
	}
}

func (s *SocialAuthService) GetGoogleAuthURL(state string) string {
	return s.googleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *SocialAuthService) GetFacebookAuthURL(state string) string {
	return s.facebookConfig.AuthCodeURL(state)
}

func (s *SocialAuthService) AuthenticateWithGoogle(code string) (*model.User, error) {
	token, err := s.googleConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	userInfo, err := s.getGoogleUserInfo(token.AccessToken)
	if err != nil {
		return nil, err
	}

	return s.findOrCreateSocialUser(userInfo)
}

func (s *SocialAuthService) AuthenticateWithFacebook(code string) (*model.User, error) {
	token, err := s.facebookConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	userInfo, err := s.getFacebookUserInfo(token.AccessToken)
	if err != nil {
		return nil, err
	}

	return s.findOrCreateSocialUser(userInfo)
}

func (s *SocialAuthService) AuthenticateWithToken(provider, accessToken string) (*model.User, error) {
	var userInfo *model.SocialUserInfo
	var err error

	switch provider {
	case "google":
		userInfo, err = s.getGoogleUserInfo(accessToken)
	case "facebook":
		userInfo, err = s.getFacebookUserInfo(accessToken)
	default:
		return nil, errors.New("unsupported provider")
	}

	if err != nil {
		return nil, err
	}

	return s.findOrCreateSocialUser(userInfo)
}

func (s *SocialAuthService) getGoogleUserInfo(accessToken string) (*model.SocialUserInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("google API error: %s", string(body))
	}

	var data struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &model.SocialUserInfo{
		ID:       data.ID,
		Email:    data.Email,
		Name:     data.Name,
		Picture:  data.Picture,
		Provider: "google",
	}, nil
}

func (s *SocialAuthService) getFacebookUserInfo(accessToken string) (*model.SocialUserInfo, error) {
	url := fmt.Sprintf("https://graph.facebook.com/me?fields=id,name,email,picture&access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("facebook API error: %s", string(body))
	}

	var data struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &model.SocialUserInfo{
		ID:       data.ID,
		Email:    data.Email,
		Name:     data.Name,
		Picture:  data.Picture.Data.URL,
		Provider: "facebook",
	}, nil
}

func (s *SocialAuthService) findOrCreateSocialUser(info *model.SocialUserInfo) (*model.User, error) {
	if info.Email == "" {
		return nil, errors.New("email not provided by social provider")
	}

	// Buscar usuário existente por email
	user, err := s.userService.GetUserByEmail(info.Email)
	if err == nil {
		// Usuário existe, atualizar informações sociais se necessário
		needsUpdate := false

		if user.Provider == "" || user.ProviderID == "" {
			user.Provider = info.Provider
			user.ProviderID = info.ID
			needsUpdate = true
		}

		if user.Avatar == "" && info.Picture != "" {
			user.Avatar = info.Picture
			needsUpdate = true
		}

		if needsUpdate {
			if err := s.userService.repo.Update(user.ID, user); err != nil {
				return nil, fmt.Errorf("failed to update user: %w", err)
			}
		}

		return user, nil
	}

	// Criar novo usuário
	user = &model.User{
		Email:      info.Email,
		FullName:   info.Name,
		Avatar:     info.Picture,
		Provider:   info.Provider,
		ProviderID: info.ID,
		Active:     true,
		Password:   "", // Não tem senha para login social
	}

	if err := s.userService.repo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
