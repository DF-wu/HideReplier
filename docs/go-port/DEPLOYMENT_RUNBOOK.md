# Deployment Runbook — HideReplier Go Port

## 1. 執行環境

- Fly.io shared-cpu-1x
- 1 CPU
- 256MB RAM

## 2. 必要環境變數

- `PORT`（預設 8082）
- `BOT_VERSION`
- `HOST_URL`
- `DC_WEBHOOK_URL`
- `MONGO_URI`
- `MONGO_DATABASE`

其中 `DC_WEBHOOK_URL` 與 `MONGO_URI` 為必填；缺少時 Go 服務必須啟動失敗而不是使用隱藏預設值。

## 3. 啟動條件

服務需：
- 成功連接 MongoDB
- 可回應 `/actuator/health`

## 4. Fly.io 相容需求

- 內部 port 維持 `8082`
- 健康檢查路徑維持 `/actuator/health`
- 單 binary 啟動
- 映像檔應使用 multi-stage build

## 5. 觀察項目

- 啟動時間
- 記憶體佔用
- webhook 成功/失敗率
- history 與 counter 是否一致

## 6. 回滾策略

若 Go 版部署失敗：
1. 將流量切回 Java 版穩定映像
2. 保留同一 MongoDB 資料
3. 重新檢查 webhook 送出與持久化順序
