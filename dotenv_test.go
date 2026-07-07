package main

import (
	"strings"
	"testing"
)

func TestParseDotEnv(t *testing.T) {
	in := `
# 注释行
export APIFY_TOKEN="apify_api_abc123"
XHS_PROXY = 'http://user:pass@host:8080'
EMPTY=
无等号的坏行
`
	env := parseDotEnv(strings.NewReader(in))

	if env["APIFY_TOKEN"] != "apify_api_abc123" {
		t.Fatalf("APIFY_TOKEN = %q, want apify_api_abc123", env["APIFY_TOKEN"])
	}
	if env["XHS_PROXY"] != "http://user:pass@host:8080" {
		t.Fatalf("XHS_PROXY = %q, want 去引号去空格后的值", env["XHS_PROXY"])
	}
	if v, ok := env["EMPTY"]; !ok || v != "" {
		t.Fatalf("EMPTY 应为空字符串, got ok=%v v=%q", ok, v)
	}
	if _, ok := env["无等号的坏行"]; ok {
		t.Fatal("无等号的行不应被解析")
	}
}
