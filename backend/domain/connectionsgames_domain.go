package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aogen-fiber/backend/infrastructure"
	"github.com/sqids/sqids-go"
)

type ConnectionsGame struct {
	ID         int64      `json:"id,omitempty"`
	GameId     string     `json:"gameId"`
	Name       string     `json:"name"`
	Creator    string     `json:"creator"`
	CreatedAt  int64      `json:"created_at"`
	Rating     int        `json:"rating"`
	Played     bool       `json:"played"`
	Categories []Category `json:"categories"`
	Stats      *GameStats `json:"stats,omitempty"`
}

type Category struct {
	Name        string   `json:"title"`
	Color       string   `json:"color"`
	Words       []string `json:"words"`
	Explanation string   `json:"explanation,omitempty"`
	Tip         string   `json:"tip,omitempty"`
}

// NewConnectionsGame creates a new game instance
func NewConnectionsGame(name, creator string, rating int, categories []Category) *ConnectionsGame {
	return &ConnectionsGame{
		Name:       name,
		Creator:    creator,
		CreatedAt:  time.Now().Unix(),
		Rating:     rating,
		Played:     false,
		Categories: categories,
	}
}

// Validate validates game data
func (g *ConnectionsGame) Validate() error {
	if g.Name == "" {
		return errors.New("game name cannot be empty")
	}
	if g.Creator == "" {
		return errors.New("creator cannot be empty")
	}
	if len(g.Categories) != 4 {
		return errors.New("must contain 4 categories")
	}

	for _, cat := range g.Categories {
		if len(cat.Words) != 4 {
			return errors.New("each category must contain 4 words")
		}
	}
	return nil
}

// GameStats 表示游戏统计信息
type GameStats struct {
	GameId            string      `json:"gameId"`
	TotalPlays        int         `json:"totalPlays"`
	PerfectRate       float64     `json:"perfectRate"`
	AvgGuesses        float64     `json:"avgGuesses"`
	GuessDistribution map[int]int `json:"guessDistribution"`
}

// GameReport 表示游戏完成报告
type GameReport struct {
	GuessCount int  `json:"guessCount"`
	IsPerfect  bool `json:"isPerfect"`
}

// ConnectionsGameRepository defines game storage interface
type ConnectionsGameRepository interface {
	Save(game *ConnectionsGame) error
	GetByID(id int64) (*ConnectionsGame, error)
	GetAll() ([]*ConnectionsGame, error)
	UpdateRandomRatings() error
	GetRandomGames(limit int) ([]*ConnectionsGame, error)
	SaveGameReport(gameId int64, report *GameReport) error
	GetGameStats(gameId int64) (*GameStats, error)
}

// ConnectionsGameService defines game service interface
type ConnectionsGameService struct {
	repo   ConnectionsGameRepository
	ollama *infrastructure.OllamaService
	sqids  *sqids.Sqids
}

// NewConnectionsGameService creates a game service instance
func NewConnectionsGameService(repo ConnectionsGameRepository, ollama *infrastructure.OllamaService) *ConnectionsGameService {
	// Initialize Sqids
	options := sqids.Options{
		Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		MinLength: 8,
	}
	s, _ := sqids.New(options)

	return &ConnectionsGameService{
		repo:   repo,
		ollama: ollama,
		sqids:  s,
	}
}

// encodeID converts numeric ID to obfuscated string
func (s *ConnectionsGameService) encodeID(id int64) string {
	encoded, _ := s.sqids.Encode([]uint64{uint64(id)})
	return encoded
}

// decodeID converts obfuscated string back to numeric ID
func (s *ConnectionsGameService) decodeID(encoded string) (int64, error) {
	numbers := s.sqids.Decode(encoded)
	if len(numbers) == 0 {
		return 0, fmt.Errorf("invalid encoded ID")
	}
	return int64(numbers[0]), nil
}

// Save saves the game and generates explanations and tips
func (s *ConnectionsGameService) Save(game *ConnectionsGame) error {
	if err := game.Validate(); err != nil {
		return err
	}

	colors := []string{"yellow", "green", "blue", "purple"}
	for i := range game.Categories {
		if game.Categories[i].Color == "" {
			game.Categories[i].Color = colors[i]
		}
	}

	// Check if explanations and tips need to be generated
	needGenerate := false
	for _, cat := range game.Categories {
		if cat.Explanation == "" || cat.Tip == "" {
			needGenerate = true
			break
		}
	}

	if needGenerate {
		if err := s.generateExplanationsAndTips(game); err != nil {
			return err
		}
	}

	if err := s.repo.Save(game); err != nil {
		return err
	}

	// 设置编码后的 gameId
	game.GameId = s.encodeID(game.ID)
	game.ID = 0 // 隐藏原始 ID

	return nil
}

// generateExplanationsAndTips generates explanations and tips for each category
func (s *ConnectionsGameService) generateExplanationsAndTips(game *ConnectionsGame) error {
	if s.ollama == nil {
		return errors.New("ollama service not initialized")
	}

	for i, cat := range game.Categories {
		if cat.Explanation != "" && cat.Tip != "" {
			continue
		}

		prompt := fmt.Sprintf(`You are a helpful assistant that analyzes word connections.
Given these 4 words: %s
Category name: %s

Your task:
1. Find a meaningful connection between these words, considering the category name as a reference
2. Create an explanation (max 100 chars)
3. Create a helpful tip (max 50 chars)

Rules:
- ALWAYS return a valid JSON response
- NEVER explain your thinking
- NEVER include additional text
- ONLY return the JSON object

Required JSON format:
{"explanation":"[your explanation here]","tip":"[your tip here]"}`, strings.Join(cat.Words, ", "), cat.Name)
		fmt.Println(prompt)
		var fullResponse strings.Builder
		err := s.ollama.GenerateStream("llama3.2:3b", prompt, func(resp infrastructure.OllamaResponse) {
			fullResponse.WriteString(resp.Response)
		})

		if err != nil {
			return fmt.Errorf("failed to generate explanation for category %s: %w", cat.Name, err)
		}

		// 打印原始 LLM 响应
		fmt.Printf("Raw LLM response for category %s:\n%s\n", cat.Name, fullResponse.String())

		// 清理和格式化 JSON 字符串
		jsonStr := fullResponse.String()
		jsonStr = strings.TrimSpace(jsonStr)

		response := struct {
			Explanation string `json:"explanation"`
			Tip         string `json:"tip"`
		}{}

		if err := json.Unmarshal([]byte(jsonStr), &response); err != nil {
			return fmt.Errorf("failed to parse explanation response: %w, raw response: %s", err, jsonStr)
		}

		game.Categories[i].Explanation = response.Explanation
		game.Categories[i].Tip = response.Tip
	}

	return nil
}

// GameListItem 表示游戏列表中的简化游戏信息
type GameListItem struct {
	GameId  string `json:"gameId"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
	Rating  int    `json:"rating"`
}

// GetAll gets all game list
func (s *ConnectionsGameService) GetAll() ([]*GameListItem, error) {
	games, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get games: %w", err)
	}

	// 转换为简化的游戏列表项
	gameList := make([]*GameListItem, len(games))
	for i, game := range games {
		gameList[i] = &GameListItem{
			GameId:  s.encodeID(game.ID),
			Name:    game.Name,
			Creator: game.Creator,
			Rating:  game.Rating,
		}
	}

	return gameList, nil
}

// GetByID gets game by obfuscated ID
func (s *ConnectionsGameService) GetByEncodedID(encodedID string) (*ConnectionsGame, error) {
	id, err := s.decodeID(encodedID)
	if err != nil {
		return nil, fmt.Errorf("invalid game ID: %w", err)
	}

	game, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get game by id: %w", err)
	}

	game.GameId = s.encodeID(game.ID)
	game.ID = 0 // 隐藏原始 ID

	return game, nil
}

// UpdateRandomRatings updates random ratings for all games
func (s *ConnectionsGameService) UpdateRandomRatings() error {
	return s.repo.UpdateRandomRatings()
}

// GetRandomRecommendations 获取随机推荐的游戏
func (s *ConnectionsGameService) GetRandomRecommendations(limit int) ([]*GameListItem, error) {
	games, err := s.repo.GetRandomGames(limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get random games: %w", err)
	}

	// 转换为简化的游戏列表项
	gameList := make([]*GameListItem, len(games))
	for i, game := range games {
		gameList[i] = &GameListItem{
			GameId:  s.encodeID(game.ID),
			Name:    game.Name,
			Creator: game.Creator,
			Rating:  game.Rating,
		}
	}

	return gameList, nil
}

// SaveGameReport saves game completion report
func (s *ConnectionsGameService) SaveGameReport(encodedID string, report *GameReport) error {
	id, err := s.decodeID(encodedID)
	if err != nil {
		return fmt.Errorf("invalid game ID: %w", err)
	}
	return s.repo.SaveGameReport(id, report)
}

// GetGameStats gets game statistics
func (s *ConnectionsGameService) GetGameStats(encodedID string) (*GameStats, error) {
	id, err := s.decodeID(encodedID)
	if err != nil {
		return nil, fmt.Errorf("invalid game ID: %w", err)
	}

	stats, err := s.repo.GetGameStats(id)
	if err != nil {
		return nil, err
	}

	stats.GameId = encodedID
	return stats, nil
}
