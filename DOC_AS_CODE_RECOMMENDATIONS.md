# 开源 Doc as Code 仓库推荐

根据您的需求（Go 语言、尽量不依赖第三方前端框架、支持 Markdown 直接生成静态站点），以下是精选的开源项目推荐。

## 一、静态站点生成器（SSG）

### 1. Hugo ⭐⭐⭐⭐⭐
**仓库**: https://github.com/gohugoio/hugo  
**语言**: Go  
**特点**:
- 全球最快的静态网站生成器，单可执行文件
- 完全用 Go 编写，无运行时依赖
- 使用 Go 标准库的 `html/template`
- 支持 Frontmatter、分类、标签、分页
- 丰富的主题生态
- 适合：博客、文档站点

**学习价值**: 参考其模板系统和构建流程设计

---

### 2. Zola ⭐⭐⭐⭐
**仓库**: https://github.com/getzola/zola  
**语言**: Rust  
**特点**:
- 单可执行文件，零依赖
- 逻辑简洁，生成的代码干净
- 专注于 Markdown → HTML 转换
- 性能极佳
- 适合：追求极简和性能的项目

---

### 3. MkDocs ⭐⭐⭐⭐
**仓库**: https://github.com/mkdocs/mkdocs  
**语言**: Python  
**特点**:
- 专注于文档站点
- 生成的页面简洁，无复杂框架
- 支持主题定制
- 适合：技术文档、API 文档

---

## 二、Go 语言 Markdown 处理库

### 4. Goldmark ⭐⭐⭐⭐⭐
**仓库**: https://github.com/yuin/goldmark  
**语言**: Go  
**特点**:
- Hugo 默认使用的 Markdown 解析器
- 完全兼容 CommonMark 规范
- 易于扩展（代码高亮、数学公式等）
- 性能优秀
- **推荐用于您的项目**

**相关扩展**:
- `github.com/yuin/goldmark-highlighting` - 代码高亮
- `github.com/yuin/goldmark-math` - 数学公式支持

---

### 5. Blackfriday ⭐⭐⭐
**仓库**: https://github.com/russross/blackfriday  
**语言**: Go  
**特点**:
- 老牌 Go Markdown 解析器
- 稳定可靠，广泛使用
- 功能相对基础

---

## 三、轻量级/零框架方案

### 6. Docsify ⭐⭐⭐⭐
**仓库**: https://github.com/docsifyjs/docsify  
**语言**: JavaScript  
**特点**:
- 无需构建过程，运行时渲染 Markdown
- 只需要一个 `index.html` 文件
- 轻量级，无复杂依赖
- 适合：快速搭建文档站点

**注意**: 需要浏览器端 JavaScript，但非常轻量

---

### 7. VitePress ⭐⭐⭐⭐
**仓库**: https://github.com/vuejs/vitepress  
**语言**: Vue.js (但生成静态 HTML)  
**特点**:
- Vue.js 官方文档工具
- 构建时生成静态 HTML
- 支持 Vue 组件（可选）
- 适合：需要交互式组件的文档

---

## 四、参考实现（小型博客生成器）

### 8. go-blog ⭐⭐⭐
**搜索关键词**: "go blog generator"  
**特点**:
- 许多开发者用 Go 实现的小型博客生成器
- 代码简单，易于学习和修改
- 适合：学习如何实现 SSG

**示例仓库**:
- 在 GitHub 搜索 "go static site generator"
- 查看代码量 < 5000 行的项目

---

### 9. Jekyll ⭐⭐⭐⭐
**仓库**: https://github.com/jekyll/jekyll  
**语言**: Ruby  
**特点**:
- GitHub Pages 默认支持
- 成熟的生态系统
- 适合：GitHub 托管的博客

**注意**: 需要 Ruby 环境，不符合"零依赖"要求

---

## 五、内容管理工具

### 10. Contentlayer ⭐⭐⭐
**仓库**: https://github.com/contentlayerdev/contentlayer  
**语言**: TypeScript  
**特点**:
- 将 Markdown 转为类型安全的 JSON
- 适合与 Next.js 等框架集成
- 不符合"不依赖三方框架"要求

---

## 推荐学习路径

### 方案 A：直接使用 Hugo（最快）
1. 研究 Hugo 的目录结构和配置
2. 参考其模板系统设计
3. 如果满足需求，直接使用；否则参考其实现

### 方案 B：自研 + 参考 Goldmark（最灵活）
1. 使用 **Goldmark** 解析 Markdown
2. 使用 Go 标准库 `html/template` 生成 HTML
3. 参考 **Hugo** 的模板和数据流设计
4. 参考 **Zola** 的简洁实现思路

### 方案 C：学习小型实现（最易理解）
1. 在 GitHub 搜索小型 Go SSG 项目
2. 阅读源码，理解核心流程
3. 在此基础上定制开发

## 核心学习点

无论选择哪个项目，重点关注：

1. **Markdown 解析流程**
   - Frontmatter 提取
   - Markdown → HTML 转换
   - 代码块和数学公式处理

2. **模板系统设计**
   - 模板继承/嵌套
   - 数据注入方式
   - 部分模板（partials）管理

3. **文件组织结构**
   - 源文件扫描
   - 输出目录结构
   - 静态资源处理

4. **构建流程优化**
   - 增量构建
   - 并行处理
   - 错误处理

## 快速开始建议

**如果您想快速上手**：
```bash
# 1. 安装 Hugo
brew install hugo  # macOS
# 或下载二进制文件

# 2. 创建新站点
hugo new site my-blog

# 3. 研究其目录结构和模板
cd my-blog
```

**如果您想自研**：
```bash
# 1. 研究 Goldmark 示例
go get github.com/yuin/goldmark

# 2. 查看 Hugo 源码
git clone https://github.com/gohugoio/hugo
# 重点关注：hugolib/ 和 tpl/ 目录
```

## 总结

**最符合您需求的推荐**：
1. **Hugo** - 如果可以直接使用，最省事
2. **Goldmark + Go templates** - 如果自研，这是最佳技术栈
3. **小型 Go SSG 项目** - 如果想快速学习，找代码量小的项目研究

**学习优先级**：
1. Hugo 的模板系统（`html/template` 使用方式）
2. Goldmark 的扩展机制（如何添加代码高亮、数学公式）
3. 小型项目的整体架构（如何组织代码）
