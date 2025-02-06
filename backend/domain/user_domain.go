package domain

import (
	"time"

	"github.com/google/uuid"
)

// User 表示系统中的用户实体
type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Picture     string    `json:"picture"`
	GoogleID    string    `json:"google_id"`
	CreatedAt   time.Time `json:"created_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

// UserService 定义用户相关的业务逻辑接口
type UserService struct {
	repo UserRepository
}

// UserRepository 定义用户存储接口
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByGoogleID(googleID string) (*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(user *User) error
}

// NewUserService 创建用户服务实例
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetUserByID 通过ID获取用户
func (s *UserService) GetUserByID(id string) (*User, error) {
	return s.repo.GetUserByID(id)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(user *User) error {
	return s.repo.UpdateUser(user)
}

// ObfuscateID 对用户ID进行混淆
func ObfuscateID(googleID string) string {
	// 使用 Google ID 的前8位作为种子
	seed := googleID
	if len(seed) > 8 {
		seed = seed[:8]
	}
	// 生成一个新的 UUID
	id := uuid.NewSHA1(uuid.NameSpaceOID, []byte(seed))
	// 返回混淆后的ID
	return "u_" + id.String()
}

// CreateOrUpdateUser 创建或更新用户信息
func (s *UserService) CreateOrUpdateUser(user *User) error {
	existingUser, err := s.repo.GetUserByGoogleID(user.GoogleID)
	if err == nil && existingUser != nil {
		// 更新现有用户信息
		user.ID = existingUser.ID               // 保持原有ID
		user.CreatedAt = existingUser.CreatedAt // 保持原有创建时间
		user.LastLoginAt = time.Now()
		return s.repo.UpdateUser(user)
	}

	// 创建新用户
	user.ID = ObfuscateID(user.GoogleID)
	user.CreatedAt = time.Now()
	user.LastLoginAt = time.Now()
	return s.repo.CreateUser(user)
}
