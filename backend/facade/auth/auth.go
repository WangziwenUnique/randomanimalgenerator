package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/aogen-fiber/backend/domain"
	"google.golang.org/api/idtoken"
)

// AuthService 处理认证相关的逻辑
type AuthService struct {
	userService *domain.UserService
	clientID    string
}

// LoginResponse 表示登录响应
type LoginResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}

// NewAuthService 创建认证服务实例
func NewAuthService(userService *domain.UserService, clientID string) *AuthService {
	return &AuthService{
		userService: userService,
		clientID:    clientID,
	}
}

// GoogleUserInfo 表示从Google JWT获取的用户信息
type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

// HandleGoogleOneTapLogin 处理Google One Tap登录
func (s *AuthService) HandleGoogleOneTapLogin(credential string) (*LoginResponse, error) {
	// 验证Google JWT
	payload, err := s.verifyGoogleJWT(credential)
	if err != nil {
		return nil, fmt.Errorf("failed to verify google jwt: %v", err)
	}

	// 创建或更新用户
	user := &domain.User{
		Email:       payload.Claims["email"].(string),
		Name:        payload.Claims["name"].(string),
		Picture:     payload.Claims["picture"].(string),
		GoogleID:    payload.Claims["sub"].(string),
		LastLoginAt: time.Now(),
	}

	if err := s.userService.CreateOrUpdateUser(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	// 生成JWT
	jwtToken, err := GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &LoginResponse{
		User:  user,
		Token: jwtToken,
	}, nil
}

// verifyGoogleJWT 验证Google JWT令牌
func (s *AuthService) verifyGoogleJWT(credential string) (*idtoken.Payload, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, credential, s.clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}
	return payload, nil
}

// RefreshTokenResponse 表示刷新token的响应
type RefreshTokenResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}

// RefreshToken 刷新用户的访问令牌
func (s *AuthService) RefreshToken(authHeader string) (*RefreshTokenResponse, error) {
	// 从header中提取token
	tokenString, err := ExtractTokenFromHeader(authHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to extract token: %v", err)
	}

	// 解析JWT获取用户信息
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// 获取用户信息
	user, err := s.userService.GetUserByID(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// 生成新的JWT
	newToken, err := GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &RefreshTokenResponse{
		User:  user,
		Token: newToken,
	}, nil
}
