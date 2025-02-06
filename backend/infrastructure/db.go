package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// DBConfig 数据库配置
type DBConfig struct {
	Name string // 数据库名称
	Path string // 数据库路径
}

// DBManager 数据库连接管理器
type DBManager struct {
	connections map[string]*sql.DB
	configs     map[string]DBConfig
	mutex       sync.RWMutex
}

var (
	instance *DBManager
	once     sync.Once
)

// 默认数据库配置
const (
	DefaultDataDir = "./data"
	UsersDB        = "users"
	ConnectionsDB  = "connections"
	HumanizerDB    = "humanizer"
)

// GetDBManager 获取数据库管理器实例
func GetDBManager() *DBManager {
	once.Do(func() {
		instance = &DBManager{
			connections: make(map[string]*sql.DB),
			configs:     make(map[string]DBConfig),
		}
		// 初始化默认配置
		instance.initDefaultConfigs()
	})
	return instance
}

// initDefaultConfigs 初始化默认数据库配置
func (m *DBManager) initDefaultConfigs() {
	m.configs[UsersDB] = DBConfig{
		Name: UsersDB,
		Path: filepath.Join(DefaultDataDir, "users.db"),
	}
	m.configs[ConnectionsDB] = DBConfig{
		Name: ConnectionsDB,
		Path: filepath.Join(DefaultDataDir, "connections.db"),
	}
	m.configs[HumanizerDB] = DBConfig{
		Name: HumanizerDB,
		Path: filepath.Join(DefaultDataDir, "humanizer.db"),
	}
}

// GetDB 获取指定数据库的连接
func (m *DBManager) GetDB(dbName string) (*sql.DB, error) {
	m.mutex.RLock()
	if db, exists := m.connections[dbName]; exists {
		m.mutex.RUnlock()
		return db, nil
	}
	m.mutex.RUnlock()

	// 如果连接不存在，创建新连接
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 双重检查
	if db, exists := m.connections[dbName]; exists {
		return db, nil
	}

	config, exists := m.configs[dbName]
	if !exists {
		return nil, fmt.Errorf("未找到数据库配置: %s", dbName)
	}

	db, err := sql.Open("sqlite3", config.Path)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败 %s: %v", dbName, err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("测试数据库连接失败 %s: %v", dbName, err)
	}

	m.connections[dbName] = db
	return db, nil
}

// CloseAll 关闭所有数据库连接
func (m *DBManager) CloseAll() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for name, db := range m.connections {
		if err := db.Close(); err != nil {
			log.Printf("关闭数据库 %s 失败: %v", name, err)
		}
	}
	m.connections = make(map[string]*sql.DB)
}

// Close 关闭指定的数据库连接
func (m *DBManager) Close(dbName string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if db, exists := m.connections[dbName]; exists {
		if err := db.Close(); err != nil {
			return fmt.Errorf("关闭数据库 %s 失败: %v", dbName, err)
		}
		delete(m.connections, dbName)
	}
	return nil
}
