package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	animalListURL = "https://www.randomlists.com/data/animals.json"
	imageBaseURL  = "https://www.randomlists.com/img/animals"
	imageSaveDir  = "images"
)

type Response struct {
	RandL struct {
		Items []string `json:"items"`
	} `json:"RandL"`
}

func main() {
	// 创建保存目录
	if err := os.MkdirAll(imageSaveDir, 0755); err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}

	// 获取动物列表
	animals, err := getAnimalList()
	if err != nil {
		fmt.Printf("获取动物列表失败: %v\n", err)
		return
	}

	// 使用 WaitGroup 来等待所有下载完成
	var wg sync.WaitGroup
	// 限制并发数量
	semaphore := make(chan struct{}, 5)

	for _, animal := range animals {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			// 获取信号量
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			if err := downloadImage(name); err != nil {
				fmt.Printf("下载图片 %s 失败: %v\n", name, err)
			} else {
				fmt.Printf("成功下载图片: %s\n", name)
			}
		}(animal)
	}

	wg.Wait()
	fmt.Println("所有图片下载完成！")
}

func getAnimalList() ([]string, error) {
	resp, err := http.Get(animalListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.RandL.Items, nil
}

func downloadImage(name string) error {
	// 将空格替换为下划线
	processedName := strings.ReplaceAll(name, " ", "_")
	imageURL := fmt.Sprintf("%s/%s.webp", imageBaseURL, processedName)
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP错误: %d", resp.StatusCode)
	}

	fileName := filepath.Join(imageSaveDir, name+".webp")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
