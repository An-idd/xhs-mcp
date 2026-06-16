package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/xiaohongshu-mcp/browser"
	"github.com/xpzouying/xiaohongshu-mcp/xiaohongshu"
)

// maxConcurrentPages 同时打开的页面数上限，避免并发查询打爆内存。
const maxConcurrentPages = 4

// loginCacheTTL 登录态缓存有效期，突发查询时跳过重复探测。
const loginCacheTTL = 60 * time.Second

// browserManager 复用一个长生命周期浏览器：按请求开关页面，做并发控制与登录探测。
// 查询类操作共享同一浏览器（cookie 在创建时加载），登录/删除 cookie 后调用 reset 让其重建。
type browserManager struct {
	mu  sync.Mutex
	b   *browser.Browser
	sem chan struct{}

	loginMu       sync.Mutex
	loginVerified bool
	loginAt       time.Time
}

func newBrowserManager() *browserManager {
	return &browserManager{sem: make(chan struct{}, maxConcurrentPages)}
}

// browser 懒加载共享浏览器。
func (m *browserManager) browser() *browser.Browser {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.b == nil {
		m.b = newBrowser()
	}
	return m.b
}

// reset 关闭并丢弃共享浏览器及登录缓存，下次查询会用最新 cookie 重建。
func (m *browserManager) reset() {
	m.mu.Lock()
	b := m.b
	m.b = nil
	m.mu.Unlock()

	m.loginMu.Lock()
	m.loginVerified = false
	m.loginMu.Unlock()

	if b != nil {
		b.Close()
	}
}

// withPage 占用一个并发槽，在共享浏览器上新开页面执行 fn，页面用完即关。
// 浏览器若已崩溃，丢弃并重建一次。
func (m *browserManager) withPage(ctx context.Context, fn func(*rod.Page) error) error {
	select {
	case m.sem <- struct{}{}:
		defer func() { <-m.sem }()
	case <-ctx.Done():
		return ctx.Err()
	}

	for attempt := 0; attempt < 2; attempt++ {
		page, err := m.newPage()
		if err != nil {
			logrus.Warnf("新建页面失败(第%d次)，重建浏览器: %v", attempt+1, err)
			m.reset()
			continue
		}
		return runWithPage(page, fn)
	}
	return fmt.Errorf("浏览器不可用，重建后仍失败")
}

// newPage 在共享浏览器上新建页面，recover go-rod 的 panic 转为 error。
func (m *browserManager) newPage() (page *rod.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return m.browser().NewPage(), nil
}

// ensureLoggedIn 用短超时探测登录态；未登录立即返回明确错误，避免长时间卡死。
// 成功结果在 loginCacheTTL 内缓存，突发查询不重复探测、也不额外导航。
func (m *browserManager) ensureLoggedIn(ctx context.Context, page *rod.Page) error {
	m.loginMu.Lock()
	cached := m.loginVerified && time.Since(m.loginAt) < loginCacheTTL
	m.loginMu.Unlock()
	if cached {
		return nil
	}

	ok, err := xiaohongshu.QuickCheckLogin(ctx, page)
	if err != nil {
		return fmt.Errorf("登录态探测失败: %w", err)
	}
	if !ok {
		return fmt.Errorf("登录态已失效，请用 get_login_qrcode 重新登录（或更新 cookies.json）")
	}

	m.loginMu.Lock()
	m.loginVerified = true
	m.loginAt = time.Now()
	m.loginMu.Unlock()
	return nil
}

// runWithPage 执行 fn 并保证关闭页面，recover panic 转为 error。
func runWithPage(page *rod.Page, fn func(*rod.Page) error) (err error) {
	defer func() {
		closePage(page)
		if r := recover(); r != nil {
			err = fmt.Errorf("查询执行异常: %v", r)
		}
	}()
	return fn(page)
}

// closePage 关闭页面，忽略关闭过程中的 panic。
func closePage(page *rod.Page) {
	defer func() { _ = recover() }()
	_ = page.Close()
}
