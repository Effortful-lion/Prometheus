# Prometheus 监控集成

## 搭建步骤

1. **安装依赖**
   - 确保安装了 Go 1.20+ 版本
   - 确保安装了 Docker 和 Docker Compose

2. **克隆项目**
   - 克隆本项目到本地

3. **启动监控服务**
   - 运行启动脚本：
     - Windows: `./start-monitoring.ps1`
     - Linux/macOS: `chmod +x start-monitoring.sh && ./start-monitoring.sh`

4. **运行应用程序**
   - 打开新终端，运行：`go run main.go`

5. **访问服务**
   - 应用程序：http://localhost:8080
   - Prometheus：http://localhost:9090
   - Alertmanager：http://localhost:9093
   - Grafana：http://localhost:3000（登录凭据：admin/admin）

## 一般指标

本项目集成了以下 Prometheus 指标：

1. **计数器 (Counter)**
   - `http_requests_total`：HTTP 请求总数，带有 method、path、status 标签

2. **仪表盘 (Gauge)**
   - `current_connections`：当前连接数

3. **直方图 (Histogram)**
   - `response_time_seconds`：响应时间直方图

4. **摘要 (Summary)**
   - `response_time_summary_seconds`：响应时间摘要，包含 50%、90%、99% 分位数

## 一键启动 prometheus 监控服务和 Grafana 可视化界面

### Windows
```powershell
./start-monitoring.ps1
```

### Linux/macOS
```bash
chmod +x start-monitoring.sh
./start-monitoring.sh
```

## Alertmanager 告警配置

### 告警规则

本项目配置了以下告警规则：

1. **应用程序告警**
   - `HighErrorRate`：HTTP 请求错误率超过 10%
   - `HighResponseTime`：95% 分位响应时间超过 1 秒
   - `HighConnectionCount`：当前连接数超过 100

2. **系统告警**
   - `PrometheusTargetDown`：监控目标 down 超过 5 分钟
   - `PrometheusAlertmanagerDown`：Alertmanager down 超过 5 分钟

### 告警通知

默认配置为邮件通知，需要在 `alertmanager.yml` 中修改以下参数：

```yaml
receivers:
- name: 'email'
  email_configs:
  - to: 'alerts@example.com'          # 接收告警的邮箱
    from: 'alertmanager@example.com'  # 发送告警的邮箱
    smarthost: 'smtp.example.com:587' # SMTP 服务器地址
    auth_username: 'alertmanager'     # SMTP 用户名
    auth_password: 'password'         # SMTP 密码
    auth_identity: 'alertmanager@example.com'
    require_tls: true
```

### 访问 Alertmanager

- Alertmanager 界面：http://localhost:9093

## 如何扩展

### 添加新指标

1. **选择指标类型**
   - 根据监控需求选择合适的指标类型：
     - Counter：用于累计值，如请求数、错误数
     - Gauge：用于瞬时值，如连接数、内存使用
     - Histogram：用于分布值，如响应时间分布
     - Summary：用于分位数统计，如响应时间分位数

2. **在对应文件中定义指标**
   - 在 `metrics/counter.go`、`metrics/gauge.go`、`metrics/histogram.go` 或 `metrics/summary.go` 中添加新指标

3. **注册指标**
   - 确保新指标在对应文件的注册函数中被注册

4. **使用指标**
   - 在代码中导入 `Prometheus/metrics` 包
   - 使用对应的指标变量进行操作

### 示例：添加新的计数器指标

1. 在 `metrics/counter.go` 中添加：
   ```go
   // 错误计数
   var ErrorCount = prometheus.NewCounterVec(
       prometheus.CounterOpts{
           Name: "error_count",
           Help: "Total number of errors",
       },
       []string{"error_type"},
   )
   ```

2. 在 `RegisterCounters()` 函数中添加：
   ```go
   prometheus.MustRegister(ErrorCount)
   ```

3. 在代码中使用：
   ```go
   metrics.ErrorCount.WithLabelValues("database").Inc()
   ```

### Grafana 集成

第一次登录默认都是admin，进去需要改密码。

1. **添加数据源**
   - 登录 Grafana：http://localhost:3000
   - 点击左侧菜单 "Configuration" > "Data sources"
   - 点击 "Add data source"，选择 "Prometheus"
   - URL 填写：`http://prometheus:9090`
   - 点击 "Save & Test"

2. **创建仪表板**
   - 点击左侧菜单 "Create" > "Dashboard"
   - 点击 "Add new panel"
   - 在 "Query" 标签页中选择 Prometheus 数据源
   - 输入查询语句，如：`sum(http_requests_total)`
   - 配置面板选项，点击 "Apply"

3. **导入仪表板**
   - 点击左侧菜单 "Dashboard" > "Import"
   - 输入仪表板 ID 或上传 JSON 文件
   - 选择 Prometheus 数据源，点击 "Import"

4. **集成 Alertmanager**
   - 点击左侧菜单 "Alerting" > "Alertmanager"
   - 点击 "Add Alertmanager"
   - URL 填写：`http://alertmanager:9093`
   - 点击 "Save"

