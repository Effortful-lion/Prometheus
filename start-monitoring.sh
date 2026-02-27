#!/bin/bash

# 启动监控系统
echo "Starting monitoring system..."

# 启动 Prometheus 和 Grafana 服务
docker-compose up -d

# 等待服务启动
echo "Waiting for services to start..."
sleep 10

# 显示访问地址
echo "Monitoring system started successfully!"
echo "---------------------------------------"
echo "Prometheus: http://localhost:9090"
echo "Grafana: http://localhost:3000"
echo "---------------------------------------"
echo "Grafana login credentials:"
echo "  Username: admin"
echo "  Password: admin"
echo "---------------------------------------"
echo "To run the application, open a new terminal and execute: go run main.go"
echo "---------------------------------------"
echo "To stop the system, run: docker-compose down"