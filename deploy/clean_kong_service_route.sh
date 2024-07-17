#!/bin/bash

# 获取所有路由的ID并删除它们
ROUTE_IDS=$(curl -s http://localhost:8001/routes | jq -r '.data[].id')
for ID in $ROUTE_IDS; do
  curl -i -X DELETE http://localhost:8001/routes/$ID
done

# 获取所有服务的ID并删除它们
SERVICE_IDS=$(curl -s http://localhost:8001/services | jq -r '.data[].id')
for ID in $SERVICE_IDS; do
  curl -i -X DELETE http://localhost:8001/services/$ID
done
