package middleware

import (
	"github.com/aogen-fiber/backend/facade/auth"
	"github.com/gofiber/fiber/v2"
)

// AuthUser 表示认证后的用户信息
type AuthUser struct {
	UserID string
	Email  string
}

// AuthenticateJWT 验证JWT token并解析用户信息
func AuthenticateJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			// 如果没有token，继续处理请求
			return c.Next()
		}

		// 提取token
		tokenString, err := auth.ExtractTokenFromHeader(authHeader)
		if err != nil {
			// token格式错误，继续处理请求
			return c.Next()
		}

		// 解析token
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			// token无效，继续处理请求
			return c.Next()
		}

		// 将用户信息存储在上下文中
		c.Locals("user", &AuthUser{
			UserID: claims.UserID,
			Email:  claims.Email,
		})

		return c.Next()
	}
}

// RequireAuth 要求用户必须登录
func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := GetCurrentUser(c)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Authentication required",
			})
		}
		return c.Next()
	}
}

// GetCurrentUser 从上下文中获取当前用户信息
func GetCurrentUser(c *fiber.Ctx) *AuthUser {
	user, ok := c.Locals("user").(*AuthUser)
	if !ok {
		return nil
	}
	return user
}
