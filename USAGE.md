# Docsify 使用说明

## 本地开发

### 方式一：使用 docsify-cli（推荐）

```bash
# 全局安装 docsify-cli
npm i docsify-cli -g

# 在项目根目录启动本地服务器
docsify serve .

# 访问 http://localhost:3000
```

### 方式二：使用 Python 简单服务器

```bash
# Python 3
python3 -m http.server 8000

# 访问 http://localhost:8000
```

### 方式三：使用 Go 简单服务器

```bash
# 使用 Go 内置服务器
go run -m http.server 8000

# 或使用其他 Go HTTP 服务器工具
```

## 部署

### GitHub Pages

1. 将代码推送到 GitHub 仓库
2. 在仓库设置中启用 GitHub Pages
3. 选择主分支（main/master）作为源
4. 访问 `https://yourusername.github.io/repository-name`

### Vercel / Netlify

1. 连接 GitHub 仓库
2. 构建命令留空（docsify 无需构建）
3. 发布目录设置为 `/`（根目录）
4. 部署即可

## 文件说明

- `index.html` - Docsify 入口文件，包含所有配置
- `README.md` - 首页内容
- `_sidebar.md` - 侧边栏导航配置
- `posts/` - Markdown 文章目录
- `.nojekyll` - 告诉 GitHub Pages 不要使用 Jekyll

## 添加新文章

1. 在 `posts/` 目录下创建新的 `.md` 文件
2. 在 `_sidebar.md` 中添加链接
3. 文件中的 frontmatter 会自动被移除（不会显示在页面上）

## 功能特性

- ✅ 自动移除 Markdown 文件的 frontmatter
- ✅ 支持数学公式（KaTeX）：`$...$` 和 `$$...$$`
- ✅ 代码高亮（Prism.js）
- ✅ 全文搜索
- ✅ 响应式设计
- ✅ 图片缩放
- ✅ Emoji 支持

## 自定义配置

编辑 `index.html` 中的 `window.$docsify` 对象来修改配置。

常用配置：
- `name`: 站点名称
- `repo`: GitHub 仓库地址
- `loadSidebar`: 是否加载侧边栏
- `subMaxLevel`: 侧边栏最大层级
- `search`: 搜索配置
