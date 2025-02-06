package repository

import (
	"database/sql"

	"github.com/aogen-fiber/backend/domain"
	"github.com/aogen-fiber/backend/infrastructure"
	"github.com/google/uuid"
)

// UserRepository 实现用户存储接口
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建用户存储实例
func NewUserRepository() (*UserRepository, error) {
	db, err := infrastructure.GetDBManager().GetDB(infrastructure.UsersDB)
	if err != nil {
		return nil, err
	}

	repo := &UserRepository{db: db}
	if err := repo.initTable(); err != nil {
		return nil, err
	}
	return repo, nil
}

// initTable 初始化用户表
func (r *UserRepository) initTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		picture TEXT,
		google_id TEXT UNIQUE,
		created_at DATETIME NOT NULL,
		last_login_at DATETIME NOT NULL
	)`
	_, err := r.db.Exec(query)
	return err
}

// CreateUser 创建新用户
func (r *UserRepository) CreateUser(user *domain.User) error {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	query := `
	INSERT INTO users (
		user_id, email, name, picture, google_id,
		created_at, last_login_at
	) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query,
		user.ID, user.Email, user.Name, user.Picture, user.GoogleID,
		user.CreatedAt, user.LastLoginAt)
	return err
}

// GetUserByEmail 通过邮箱查找用户
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT user_id, email, name, picture, google_id, created_at, last_login_at FROM users WHERE email = ?`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Name, &user.Picture, &user.GoogleID,
		&user.CreatedAt, &user.LastLoginAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

// GetUserByGoogleID 通过GoogleID查找用户
func (r *UserRepository) GetUserByGoogleID(googleID string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT user_id, email, name, picture, google_id, created_at, last_login_at FROM users WHERE google_id = ?`
	err := r.db.QueryRow(query, googleID).Scan(
		&user.ID, &user.Email, &user.Name, &user.Picture, &user.GoogleID,
		&user.CreatedAt, &user.LastLoginAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

// UpdateUser 更新用户信息
func (r *UserRepository) UpdateUser(user *domain.User) error {
	query := `
	UPDATE users SET 
		email = ?, name = ?, picture = ?,
		last_login_at = ?
	WHERE user_id = ?`

	_, err := r.db.Exec(query,
		user.Email, user.Name, user.Picture,
		user.LastLoginAt, user.ID)
	return err
}

// GetUserByID 通过ID查找用户
func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT user_id, email, name, picture, google_id, created_at, last_login_at FROM users WHERE user_id = ?`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Name, &user.Picture, &user.GoogleID,
		&user.CreatedAt, &user.LastLoginAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}
