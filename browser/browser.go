package browser

import (
	"encoding/json"
	"net/url"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/xiaohongshu-mcp/cookies"
)

// defaultUserAgent 与原 headless_browser 默认 UA 保持一致。
const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"

// Browser 封装 rod 浏览器与 launcher。
// 自建引导(替代 headless_browser 薄封装)以支持新无头模式 --headless=new，增强反检测。
type Browser struct {
	browser  *rod.Browser
	launcher *launcher.Launcher
}

type browserConfig struct {
	binPath string
}

type Option func(*browserConfig)

func WithBinPath(binPath string) Option {
	return func(c *browserConfig) {
		c.binPath = binPath
	}
}

// maskProxyCredentials 日志中隐藏代理的账号密码。
func maskProxyCredentials(proxyURL string) string {
	u, err := url.Parse(proxyURL)
	if err != nil || u.User == nil {
		return proxyURL
	}
	if _, hasPassword := u.User.Password(); hasPassword {
		u.User = url.UserPassword("***", "***")
	} else {
		u.User = url.User("***")
	}
	return u.String()
}

// NewBrowser 创建浏览器。
//   - headless=false: 有头模式
//   - headless=true, headlessNew=true: 新无头模式(--headless=new，反检测更好)
//   - headless=true, headlessNew=false: 旧无头模式
func NewBrowser(headless, headlessNew bool, options ...Option) *Browser {
	cfg := &browserConfig{}
	for _, opt := range options {
		opt(cfg)
	}

	l := launcher.New().
		Set("--no-sandbox").
		Set("user-agent", defaultUserAgent)

	switch {
	case !headless:
		l = l.Headless(false)
	case headlessNew:
		l = l.HeadlessNew(true)
	default:
		l = l.Headless(true)
	}

	if cfg.binPath != "" {
		l = l.Bin(cfg.binPath)
	}

	// 从环境变量读取代理
	if proxy := os.Getenv("XHS_PROXY"); proxy != "" {
		l = l.Proxy(proxy)
		logrus.Infof("Using proxy: %s", maskProxyCredentials(proxy))
	}

	controlURL := l.MustLaunch()
	b := rod.New().ControlURL(controlURL).MustConnect()

	loadCookies(b)

	return &Browser{browser: b, launcher: l}
}

// NewPage 创建启用 stealth 反检测的新页面。
func (b *Browser) NewPage() *rod.Page {
	return stealth.MustPage(b.browser)
}

// Close 关闭浏览器并清理 launcher 临时目录。
func (b *Browser) Close() {
	b.browser.MustClose()
	b.launcher.Cleanup()
}

// loadCookies 从 cookies 文件加载登录态到浏览器。
func loadCookies(b *rod.Browser) {
	cookiePath := cookies.GetCookiesFilePath()
	data, err := cookies.NewLoadCookie(cookiePath).LoadCookies()
	if err != nil {
		logrus.Warnf("failed to load cookies: %v", err)
		return
	}

	var cks []*proto.NetworkCookie
	if err := json.Unmarshal(data, &cks); err != nil {
		logrus.Warnf("failed to unmarshal cookies: %v", err)
		return
	}
	if len(cks) > 0 {
		b.MustSetCookies(cks...)
	}
}
