package configs

var (
	useHeadless = true

	// headlessNew 是否使用 Chrome 新无头模式(--headless=new)。
	// 新无头模式渲染路径接近有头浏览器，反检测能力更好。
	headlessNew = true

	binPath = ""
)

func InitHeadless(h bool) {
	useHeadless = h
}

// IsHeadless 是否无头模式。
func IsHeadless() bool {
	return useHeadless
}

// SetHeadlessNew 设置是否使用新无头模式。
func SetHeadlessNew(b bool) {
	headlessNew = b
}

// IsHeadlessNew 是否使用新无头模式。
func IsHeadlessNew() bool {
	return headlessNew
}

func SetBinPath(b string) {
	binPath = b
}

func GetBinPath() string {
	return binPath
}
