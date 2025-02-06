package facade

import (
	"log"

	"github.com/aogen-fiber/backend/config"
	"github.com/aogen-fiber/backend/domain"
	"github.com/aogen-fiber/backend/facade/auth"
	"github.com/aogen-fiber/backend/facade/middleware"
	"github.com/aogen-fiber/backend/facade/stripe"
	infrastripe "github.com/aogen-fiber/backend/infrastructure/stripe"
	"github.com/aogen-fiber/backend/repository"
	"github.com/gofiber/fiber/v2"
)

// RegisterHandlers registers all HTTP handlers
func RegisterHandlers(api fiber.Router) {
	// 添加JWT认证中间件
	api.Use(middleware.AuthenticateJWT())

	// 初始化 Stripe 服务
	cfg := config.LoadConfig()
	stripeService := infrastripe.NewService(cfg.StripeSecretKey)
	stripeHandler := stripe.NewHandler(stripeService)

	// 初始化用户相关服务
	userRepo, err := repository.NewUserRepository()
	if err != nil {
		log.Fatalf("Failed to create user repository: %v", err)
	}
	userService := domain.NewUserService(userRepo)

	// 初始化Google认证服务
	authService := auth.NewAuthService(userService, cfg.GoogleClientID)

	// 认证路由
	api.Post("/auth/google-one-tap", handleGoogleOneTapLogin(authService))

	// 需要认证的路由
	api.Post("/auth/refresh_token", middleware.RequireAuth(), handleRefreshToken(authService))

	// Stripe 相关路由
	api.Post("/stripe/create-session", middleware.RequireAuth(), stripeHandler.CreateSession)
	api.Get("/stripe/subscriptions", middleware.RequireAuth(), stripeHandler.GetActiveSubscriptions)
	api.Post("/stripe/subscriptions/:id/cancel", middleware.RequireAuth(), stripeHandler.CancelSubscription)
}

// GoogleOneTapLoginRequest 表示Google One Tap登录请求
type GoogleOneTapLoginRequest struct {
	Credential string `json:"credential"`
}

// GoogleOneTapLoginResponse 表示Google One Tap登录响应
type GoogleOneTapLoginResponse struct {
	Success bool                `json:"success"`
	Data    *auth.LoginResponse `json:"data,omitempty"`
	Error   string              `json:"error,omitempty"`
}

// handleGoogleOneTapLogin 处理Google One Tap登录请求
func handleGoogleOneTapLogin(authService *auth.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req GoogleOneTapLoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(GoogleOneTapLoginResponse{
				Success: false,
				Error:   "Invalid request format",
			})
		}

		if req.Credential == "" {
			return c.Status(fiber.StatusBadRequest).JSON(GoogleOneTapLoginResponse{
				Success: false,
				Error:   "Credential is required",
			})
		}

		loginResp, err := authService.HandleGoogleOneTapLogin(req.Credential)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(GoogleOneTapLoginResponse{
				Success: false,
				Error:   err.Error(),
			})
		}

		return c.JSON(GoogleOneTapLoginResponse{
			Success: true,
			Data:    loginResp,
		})
	}
}

// RefreshTokenRequest 表示刷新token的请求
type RefreshTokenRequest struct {
	UserID string `json:"user_id"`
}

// RefreshTokenResponse 表示刷新token的响应
type RefreshTokenResponse struct {
	Success bool                       `json:"success"`
	Data    *auth.RefreshTokenResponse `json:"data,omitempty"`
	Error   string                     `json:"error,omitempty"`
}

// handleRefreshToken 处理刷新token的请求
func handleRefreshToken(authService *auth.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 从中间件获取用户信息
		user := middleware.GetCurrentUser(c)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(RefreshTokenResponse{
				Success: false,
				Error:   "Unauthorized",
			})
		}

		// 从header获取token用于刷新
		authHeader := c.Get("Authorization")
		resp, err := authService.RefreshToken(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(RefreshTokenResponse{
				Success: false,
				Error:   err.Error(),
			})
		}

		return c.JSON(RefreshTokenResponse{
			Success: true,
			Data:    resp,
		})
	}
}
