# xiaohongshu-mcp

MCP for 小红书 / xiaohongshu.com。让你的 AI 助手直接查询小红书数据。

> **本分支为「纯查询」精简版**：仅保留只读的数据查询能力（搜索、推荐列表、笔记详情+评论、用户主页），已移除发布、评论、点赞、收藏等写操作及 REST HTTP API，仅通过 MCP 协议对外提供服务。

### 🚀 快速开始：选择最适合你的版本

> [!IMPORTANT]
> #### 🔥 方案 A：Openclaw 深度集成 (推荐给开发者)
> - **Openclaw 太火啦 🔥🔥🔥 ，新增 Openclaw 支持，分为两种，请各位按需使用：**
> - [xiaohongshu-mcp-skills](https://github.com/autoclaw-cc/xiaohongshu-mcp-skills)（适用于已部署完本项目的用户）
> - [xiaohongshu-skills](https://github.com/autoclaw-cc/xiaohongshu-skills)（开箱即用版）

> [!TIP]
> #### ✨ 方案 B：x-mcp 浏览器插件版 (推荐给非技术同学 / 追求极简的用户)
> - **不想折腾 Docker 或部署环境？试试：[xpzouying/x-mcp](https://github.com/xpzouying/x-mcp)**
> - **零配置**：安装插件即用，无需任何代码、代理或复杂的环境配置。
> - **安全稳定**：直接在常用浏览器 (Chrome/Edge) 及本地网络运行，无服务器 IP 风险，且能解决 90% 的部署报错。

### 📖 相关资源

- **我的博客文章**：[haha.ai/xiaohongshu-mcp](https://www.haha.ai/xiaohongshu-mcp)

### 🛠️ 疑难杂症

如果您在部署传统 Docker 版本时遇到问题，**务必先查看：[各种疑难杂症 (Issues #56)](https://github.com/xpzouying/xiaohongshu-mcp/issues/56)**。

> *提示：如果环境排查太耗时，切换到 [x-mcp 插件版](https://github.com/xpzouying/x-mcp) 通常是更高效的选择。*

## 项目简介

**主要功能**

<details>
<summary><b>1. 登录和检查登录状态</b></summary>

第一步必须，小红书需要进行登录。可以检查当前登录状态。

</details>

<details>
<summary><b>2. 搜索内容</b></summary>

根据关键词搜索小红书内容。

</details>

<details>
<summary><b>3. 获取推荐列表</b></summary>

获取小红书首页推荐内容列表。

</details>

<details>
<summary><b>4. 获取帖子详情（包括互动数据和评论）</b></summary>

获取小红书帖子的完整详情，包括：

- 帖子内容（标题、描述、图片等）
- 用户信息
- 互动数据（点赞、收藏、分享、评论数）
- 评论列表及子评论

**⚠️ 重要提示：**

- 需要提供帖子 ID 和 xsec_token（两个参数缺一不可）
- 这两个参数可以从 Feed 列表或搜索结果中获取
- 必须先登录才能使用此功能

</details>

<details>
<summary><b>5. 获取用户个人主页</b></summary>

获取小红书用户的个人主页信息，包括用户基本信息和笔记内容。

**功能说明：**

- 获取用户基本信息（昵称、简介、头像等）
- 获取关注数、粉丝数、获赞量统计
- 获取用户发布的笔记内容列表

**⚠️ 重要提示：**

- 需要先登录才能使用此功能
- 需要提供用户 ID 和 xsec_token
- 这些参数可以从 Feed 列表或搜索结果中获取

**返回信息包括：**

- 用户基本信息：昵称、简介、头像、认证状态
- 统计数据：关注数、粉丝数、获赞量、笔记数
- 笔记列表：用户发布的所有公开笔记

</details>

**使用须知**

- **（非常重要）小红书的同一个账号不允许在多个网页端登录**，如果你登录了当前 xiaohongshu-mcp 后，就不要再在其他的网页端登录该账号，否则就会把当前 MCP 的账号“踢出登录”。你可以使用移动 App 端查看当前账号信息。
- 本项目仅做只读查询，不涉及发帖等写操作；但高频查询仍可能触发风控，建议合理控制调用频率，必要时配置代理。

**风险说明**

1. 该项目是在自己的另外一个项目的基础上开源出来的，原来的项目稳定运行一年多，没有出现过封号的情况，只有出现过 Cookies 过期需要重新登录。
2. 如果账号没有实名认证，特别是新号，一般会触发 **实名认证** 的消息提醒（参见下图）。⚠️ 这个不是封号，不用 MCP 也会要求实名认证。实名认证后，账号就正常了。建议使用该项目前就先实名。

该项目是基于学习的目的，禁止一切违法行为。

## 1. 使用教程

### 1.1. 快速开始（推荐）

**方式一：下载预编译二进制文件**

直接从 [GitHub Releases](https://github.com/xpzouying/xiaohongshu-mcp/releases) 下载对应平台的二进制文件：

**主程序（MCP 服务）：**

- **macOS Apple Silicon**: `xiaohongshu-mcp-darwin-arm64`
- **macOS Intel**: `xiaohongshu-mcp-darwin-amd64`
- **Windows x64**: `xiaohongshu-mcp-windows-amd64.exe`
- **Linux x64**: `xiaohongshu-mcp-linux-amd64`

**登录工具：**

- **macOS Apple Silicon**: `xiaohongshu-login-darwin-arm64`
- **macOS Intel**: `xiaohongshu-login-darwin-amd64`
- **Windows x64**: `xiaohongshu-login-windows-amd64.exe`
- **Linux x64**: `xiaohongshu-login-linux-amd64`

使用步骤：

```bash
# 1. 首先运行登录工具
chmod +x xiaohongshu-login-darwin-arm64
./xiaohongshu-login-darwin-arm64

# 2. 然后启动 MCP 服务
chmod +x xiaohongshu-mcp-darwin-arm64
./xiaohongshu-mcp-darwin-arm64
```

**⚠️ 重要提示**：首次运行时会自动下载无头浏览器（约 150MB），请确保网络连接正常。后续运行无需重复下载。

**方式二：源码编译**

<details>
<summary>源码编译安装详情</summary>

依赖 Golang 环境，安装方法请参考 [Golang 官方文档](https://go.dev/doc/install)。

设置 Go 国内源的代理，

```bash
# 配置 GOPROXY 环境变量，以下三选一

# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct
```

</details>

Windows 遇到问题首先看这里：[Windows 安装指南](./docs/windows_guide.md)

### 1.2. 登录

第一次需要手动登录，需要保存小红书的登录状态。

**使用二进制文件**：

```bash
# 运行对应平台的登录工具
./xiaohongshu-login-darwin-arm64
```

**使用源码**：

```bash
go run cmd/login/main.go
```

### 1.3. 启动 MCP 服务

启动 xiaohongshu-mcp 服务。

**使用二进制文件**：

```bash
# 默认：无头模式，没有浏览器界面
./xiaohongshu-mcp-darwin-arm64

# 非无头模式，有浏览器界面
./xiaohongshu-mcp-darwin-arm64 -headless=false
```

**使用源码**：

```bash
# 默认：无头模式，没有浏览器界面
go run .

# 非无头模式，有浏览器界面
go run . -headless=false
```

**配置代理（可选）**：

如果需要通过代理访问，可以设置 `XHS_PROXY` 环境变量：

```bash
# 设置代理后启动
XHS_PROXY=http://user:pass@proxy:port ./xiaohongshu-mcp-darwin-arm64

# 或使用源码
XHS_PROXY=http://proxy:port go run .
```

支持 HTTP/HTTPS/SOCKS5 代理，日志中会自动隐藏代理的认证信息。

## 1.4. 验证 MCP

```bash
npx @modelcontextprotocol/inspector
```

运行后，打开红色标记的链接，配置 MCP inspector，输入 `http://localhost:18060/mcp` ，点击 `Connect` 按钮。

**注意：** 左侧边框中的选项是否正确。

按照上面配置 MCP inspector 后，点击 `List Tools` 按钮，查看所有的 Tools。

## 2. MCP 客户端接入

本服务支持标准的 Model Context Protocol (MCP)，可以接入各种支持 MCP 的 AI 客户端。

### 2.1. 快速开始

#### 启动 MCP 服务

```bash
# 启动服务（默认无头模式）
go run .

# 或者有界面模式
go run . -headless=false
```

服务将运行在：`http://localhost:18060/mcp`

#### 验证服务状态

```bash
# 测试 MCP 连接
curl -X POST http://localhost:18060/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"initialize","params":{},"id":1}'
```

#### Claude Code CLI 接入

```bash
# 添加 HTTP MCP 服务器
claude mcp add --transport http xiaohongshu-mcp http://localhost:18060/mcp

# 检查 MCP 是否添加成功（确保 MCP 已经启动的前提下，运行下面命令）
claude mcp list
```

### 2.2. 支持的客户端

<details>
<summary><b>Claude Code CLI</b></summary>

官方命令行工具，已在上面快速开始部分展示：

```bash
# 添加 HTTP MCP 服务器
claude mcp add --transport http xiaohongshu-mcp http://localhost:18060/mcp

# 检查 MCP 是否添加成功（确保 MCP 已经启动的前提下，运行下面命令）
claude mcp list
```

</details>

<details>
<summary><b>Open Code CLI</b></summary>

使用交互式命令添加 MCP Server：

```bash
opencode mcp add
```

以添加 `xiaohongshu-mcp` 为例：

```
┌  Add MCP server
│
◇  Enter MCP server name
│  xiaohongshu-mcp
│
◇  Select MCP server type
│  Remote
│
◇  Enter MCP server URL
│  http://localhost:18060/mcp
│
◇  Does this server require OAuth authentication?
│  No
│
◆  MCP server "xiaohongshu-mcp" added to C:\Users\admin\.config\opencode\opencode.json
│
└  MCP server added successfully
```

验证是否添加成功（确保 MCP 已启动的前提下）：

```bash
opencode mcp list
```

```
┌  MCP Servers
│
●  ✓ xiaohongshu-mcp connected
```

</details>

<details>
<summary><b>Cursor</b></summary>

#### 配置文件的方式

创建或编辑 MCP 配置文件：

**项目级配置**（推荐）：
在项目根目录创建 `.cursor/mcp.json`：

```json
{
  "mcpServers": {
    "xiaohongshu-mcp": {
      "url": "http://localhost:18060/mcp",
      "description": "小红书内容查询服务 - MCP Streamable HTTP"
    }
  }
}
```

**全局配置**：
在用户目录创建 `~/.cursor/mcp.json` (同样内容)。

#### 使用步骤

1. 确保小红书 MCP 服务正在运行
2. 保存配置文件后，重启 Cursor
3. 在 Cursor 聊天中，工具应该自动可用
4. 可以通过聊天界面的 "Available Tools" 查看已连接的 MCP 工具

</details>

<details>
<summary><b>VSCode</b></summary>

#### 方法一：使用命令面板配置

1. 按 `Ctrl/Cmd + Shift + P` 打开命令面板
2. 运行 `MCP: Add Server` 命令
3. 选择 `HTTP` 方式。
4. 输入地址： `http://localhost:18060/mcp`，或者修改成对应的 Server 地址。
5. 输入 MCP 名字： `xiaohongshu-mcp`。

#### 方法二：直接编辑配置文件

**工作区配置**（推荐）：
在项目根目录创建 `.vscode/mcp.json`：

```json
{
  "servers": {
    "xiaohongshu-mcp": {
      "url": "http://localhost:18060/mcp",
      "type": "http"
    }
  },
  "inputs": []
}
```

1. 确认运行状态。
2. 查看 `tools` 是否正确检测。

</details>

<details>
<summary><b>Google Gemini CLI</b></summary>

在 `~/.gemini/settings.json` 或项目目录 `.gemini/settings.json` 中配置：

```json
{
  "mcpServers": {
    "xiaohongshu": {
      "httpUrl": "http://localhost:18060/mcp",
      "timeout": 30000
    }
  }
}
```

更多信息请参考 [Gemini CLI MCP 文档](https://google-gemini.github.io/gemini-cli/docs/tools/mcp-server.html)

</details>

<details>
<summary><b>MCP Inspector</b></summary>

调试工具，用于测试 MCP 连接：

```bash
# 启动 MCP Inspector
npx @modelcontextprotocol/inspector

# 在浏览器中连接到：http://localhost:18060/mcp
```

使用步骤：

- 使用 MCP Inspector 测试连接
- 测试 Ping Server 功能验证连接
- 检查 List Tools 是否返回 7 个工具

</details>

<details>
<summary><b>Cline</b></summary>

Cline 是一个强大的 AI 编程助手，支持 MCP 协议集成。

#### 配置方法

在 Cline 的 MCP 设置中添加以下配置：

```json
{
  "xiaohongshu-mcp": {
    "url": "http://localhost:18060/mcp",
    "type": "streamableHttp",
    "autoApprove": [],
    "disabled": false
  }
}
```

#### 使用步骤

1. 确保小红书 MCP 服务正在运行（`http://localhost:18060/mcp`）
2. 在 Cline 中打开 MCP 设置
3. 添加上述配置到 MCP 服务器列表
4. 保存配置并重启 Cline
5. 在对话中可以直接使用小红书相关功能

#### 配置说明

- `url`: MCP 服务地址
- `type`: 使用 `streamableHttp` 类型以获得更好的性能
- `autoApprove`: 可配置自动批准的工具列表（留空表示手动批准）
- `disabled`: 设置为 `false` 启用此 MCP 服务

#### 使用示例

配置完成后，可以在 Cline 中直接使用自然语言查询小红书：

```
帮我检查小红书登录状态
```

```
搜索小红书上关于"美食"的内容
```

```
帮我打开这条笔记的详情并总结评论区观点
```

</details>
<details>
<summary><b>OpenClaw（通过 MCPorter）</b></summary>

> 使用前请确保 xiaohongshu-mcp 已完成本地部署。**不建议**将 GitHub 链接直接丢给 OpenClaw 让其代为部署。

由于 OpenClaw 目前不原生支持 MCP，官方推荐通过 **MCPorter** 来调用 MCP 服务。

> 💡 **提示：** MCPorter 并非调用 MCP 的最佳方案，使用过程中可能出现一些兼容性问题，请知悉。

#### 安装与配置步骤

直接一次性将一下三行命令丢给 OpenClaw（可以是 Control UI、Telegram、Feishu等方式），Openclaw 会代为部署 MCPorter。

```
npm i -g mcporter
npx mcporter config add xiaohongshu-mcp http://localhost:18060/mcp
npx mcporter list xiaohongshu-mcp
```

完成上述步骤后，即可在 OpenClaw 中通过自然语言调用 xiaohongshu-mcp 的所有功能。

</details>
<details>
<summary><b>其他支持 HTTP MCP 的客户端</b></summary>

任何支持 HTTP MCP 协议的客户端都可以连接到：`http://localhost:18060/mcp`

基本配置模板：

```json
{
  "name": "xiaohongshu-mcp",
  "url": "http://localhost:18060/mcp",
  "type": "http"
}
```

</details>

### 2.3. 可用 MCP 工具

连接成功后，可使用以下 MCP 工具：

- `check_login_status` - 检查小红书登录状态（无参数）
- `get_login_qrcode` - 获取登录二维码，返回 Base64 图片和超时时间（无参数）
- `delete_cookies` - 删除 cookies 文件，重置登录状态，删除后需要重新登录（无参数）
- `list_feeds` - 获取小红书首页推荐列表（无参数）
- `search_feeds` - 搜索小红书内容（必需：keyword）
  - `filters`: 筛选选项（可选）
    - `sort_by`: 排序依据 - `综合`（默认）| `最新` | `最多点赞` | `最多评论` | `最多收藏`
    - `note_type`: 笔记类型 - `不限`（默认）| `视频` | `图文`
    - `publish_time`: 发布时间 - `不限`（默认）| `一天内` | `一周内` | `半年内`
    - `search_scope`: 搜索范围 - `不限`（默认）| `已看过` | `未看过` | `已关注`
    - `location`: 位置距离 - `不限`（默认）| `同城` | `附近`
- `get_feed_detail` - 获取帖子详情，包括互动数据和评论（必需：feed_id, xsec_token）
  - `load_all_comments`: 是否加载全部评论（可选），默认 false 仅返回前 10 条一级评论
  - `limit`: 限制加载的一级评论数量（可选），仅当 load_all_comments=true 时生效，默认 20
  - `click_more_replies`: 是否展开二级回复（可选），仅当 load_all_comments=true 时生效，默认 false
  - `reply_limit`: 跳过回复数过多的评论（可选），仅当 click_more_replies=true 时生效，默认 10
  - `scroll_speed`: 滚动速度（可选），`slow` | `normal` | `fast`，仅当 load_all_comments=true 时生效
- `user_profile` - 获取用户个人主页信息（必需：user_id, xsec_token）

### 2.4. 使用示例

使用 Claude Code 查询小红书内容：

**示例 1：检查登录状态**

```
帮我检查一下小红书的登录状态。
```

**示例 2：搜索内容**

```
用 xiaohongshu-mcp 搜索小红书上关于"露营装备"的笔记，按最多点赞排序，只看图文。
```

**示例 3：获取笔记详情与评论**

```
用 xiaohongshu-mcp 打开这条笔记的详情，并加载全部评论，帮我总结评论区的主要观点。
（提供从搜索/推荐结果中拿到的 feed_id 和 xsec_token）
```

**示例 4：查询用户主页**

```
用 xiaohongshu-mcp 查询这个用户的主页，看看他的粉丝数和最近发布的笔记。
```

### 2.5. 💬 MCP 使用常见问题解答

---

> ⚠️ 以下是使用 OpenClaw + MCPorter 时的已知风险，使用前请充分了解：

- OpenClaw 的 AI 自动部署行为不在本项目的维护范围内，部署结果无法保证
- MCPorter 作为中间层可能引入额外的兼容性问题，与 xiaohongshu-mcp 本身无关
- 若遇到连接失败、工具调用异常等问题，请先排查 MCPorter 自身的配置，而非提交 Issue
- 在提问社区或群组前，请先确认问题是否能在**不使用 OpenClaw** 的情况下复现

如果你没有强烈的 OpenClaw 使用需求，强烈建议改用 [Claude Code CLI](#claude-code-cli)、[Cursor](#cursor) 或 [Cline](#cline) 等原生支持 HTTP MCP 的客户端，体验会更稳定。

---

**Q:** 为什么检查登录用户名显示 `xiaghgngshu-mcp`？
**A:** 用户名是写死的。

---

**Q:** 在设备上运行 MCP 程序出现闪退如何解决？
**A:**

1. 建议 **从源码安装**（`go run .`）。
2. 或参考 [X-MCP 项目页面](https://github.com/xpzouying/x-mcp/) 的浏览器插件版。

---

**Q:** 使用 `http://localhost:18060/mcp` 进行 MCP 验证时提示无法连接？
**A:**

- 在 **Docker 环境** 下，请使用
  👉 [http://host.docker.internal:18060/mcp](http://host.docker.internal:18060/mcp)
- 在 **非 Docker 环境** 下，请使用 **本机 IPv4 地址** 访问。
