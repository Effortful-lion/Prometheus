# 启动监控系统
Write-Host "Starting monitoring system..."

# 启动 Prometheus 和 Grafana 服务
docker-compose up -d

# 等待服务启动
Write-Host "Waiting for services to start..."
Start-Sleep -Seconds 10

# 显示访问地址
Write-Host "Monitoring system started successfully!"
Write-Host "---------------------------------------"
Write-Host "Prometheus: http://localhost:9090"
Write-Host "Grafana: http://localhost:3000"
Write-Host "---------------------------------------"
Write-Host "Grafana login credentials:"
Write-Host "  Username: admin"
Write-Host "  Password: admin"
Write-Host "---------------------------------------"
Write-Host "To run the application, open a new terminal and execute: go run main.go"
Write-Host "---------------------------------------"
Write-Host "To stop the system, run: docker-compose down"