// Package apify 通过 Apify 云端 actor 抓取小红书内容。
// 在 Apify 服务器 + 代理上运行，不使用本地浏览器/账号/cookie，
// 把抓取行为与你的小红书账号彻底解耦，规避账号维度的风控封禁。
package apify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// actorEndpoint run-sync-get-dataset-items：同步运行 actor 并直接返回数据集条目。
const actorEndpoint = "https://api.apify.com/v2/acts/zen-studio~rednote-search-scraper/run-sync-get-dataset-items"

// SearchInput rednote-search-scraper 的输入参数(仅列常用项)。
type SearchInput struct {
	Keywords   []string `json:"keywords"`             // 搜索关键词
	MaxResults int      `json:"maxResults,omitempty"` // 每个关键词返回条数
	SortType   string   `json:"sortType,omitempty"`   // general|popularity_descending|time_descending
	NoteType   string   `json:"noteType,omitempty"`   // all|video|image
	TimeFilter string   `json:"timeFilter,omitempty"` // all|1d|1w|6mo
}

// Note 面向调研分析的精简笔记结构：保留判断"是否高赞 / 是否 AI 红海"所需的字段
// (正文、发布时间、各项互动数)，丢弃图片/视频流等分析用不到的字段。
// ponytail: 互动数按 actor 归一化后的 JSON 数字解析(样例确认为数字，非"1.2万"字符串)。
type Note struct {
	ID        string `json:"id"`
	Type      string `json:"type"` // video|image|...
	Title     string `json:"title"`
	Desc      string `json:"desc"` // 正文，判断内容质量/AI 模板化的关键
	URL       string `json:"url"`
	XsecToken string `json:"xsec_token"`
	Timestamp int64  `json:"timestamp"` // 发布时间(秒)，判断赛道时效/密度
	Keyword   string `json:"keyword"`   // 命中的搜索词

	Engagement struct {
		Liked     int64 `json:"liked_count"`
		Comments  int64 `json:"comments_count"`
		Collected int64 `json:"collected_count"`
		Shared    int64 `json:"shared_count"`
	} `json:"engagement"`

	Author struct {
		UserID   string `json:"userid"`
		Nickname string `json:"nickname"`
		Verified bool   `json:"verified"`
	} `json:"author"`
}

// Search 按关键词调用 Apify actor 搜索小红书笔记，云端运行、不使用本地账号/cookie。
func Search(ctx context.Context, token string, in SearchInput) ([]Note, error) {
	if token == "" {
		return nil, fmt.Errorf("APIFY_TOKEN 未配置(请在 .env 中设置)")
	}

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	endpoint := actorEndpoint + "?token=" + url.QueryEscape(token)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// actor 冷启动 + 抓取需要时间，给足超时(ctx 有更短 deadline 时以 ctx 为准)。
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("调用 Apify 失败: %w", err)
	}
	defer resp.Body.Close()

	// run-sync-get-dataset-items 成功时返回 201(Created)，故接受整个 2xx。
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("Apify 返回 %d: %s", resp.StatusCode, strings.TrimSpace(string(msg)))
	}

	var notes []Note
	if err := json.NewDecoder(resp.Body).Decode(&notes); err != nil {
		return nil, fmt.Errorf("解析 Apify 结果失败: %w", err)
	}
	return notes, nil
}
