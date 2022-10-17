## 腾讯云 Coding Devops 平台 API SDK 与 CLI 工具

### 依赖/特性 Dependencies/Feature

#### 依赖 Dependencies
- Cobra: 命令行的解析等 - parse commandline 
- Viper: 配置文件的解析 - parse config files
- TableWriter: 渲染输出表格结构体等 - render tables structs
- ColorCobra: 彩色命令行输出 - color commandline output
- survey: 友好的命令行交互 - friendly commandline interaction
- logrus: 日志输出 - log output
- go doc: 支持 godoc - support godoc
- gout: 超强的 webrequest 工具库 - webrequest library

#### 特性
- Cli 利用工具 WIP
- API 完整对接 WIP

### 项目结构 Structure

1. OpenApi - API sdk 目录
2. main - cli main 包
3. cmd - cobra 解析处理命令的主目录
4. config - 命令执行时需要读入的一些持久化配置

### 