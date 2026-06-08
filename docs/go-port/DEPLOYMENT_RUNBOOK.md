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
- `DISCORD_TARGETS`
- `MONGO_URI`
- `MONGO_DATABASE`

其中 `MONGO_URI` 為必填。`DC_WEBHOOK_URL` 與 `DISCORD_TARGETS` 至少需設定其中一個；缺少時 Go 服務必須啟動失敗而不是使用隱藏預設值。

`DISCORD_TARGETS` 用於部署到多個 Discord server/channel，格式如下：

```json
[
  {
    "id": "main",
    "label": "Main server / anonymous channel",
    "webhookUrl": "https://discord.com/api/webhooks/...",
    "default": true
  },
  {
    "id": "test",
    "label": "Test server / sandbox channel",
    "webhookUrl": "https://discord.com/api/webhooks/..."
  }
]
```

## 3. 啟動條件

服務需：
- 成功連接 MongoDB
- 可回應 `/actuator/health`
- Docker build 需成功完成 Go binary 與 `web/` Vite 前端建置

## 4. Fly.io 相容需求

- 內部 port 維持 `8082`
- 健康檢查路徑維持 `/actuator/health`
- 單 binary 啟動
- 映像檔應使用 multi-stage build
- production runtime 使用 Go binary；Java/Spring Boot 程式碼保留為 legacy 維護路徑，不參與 Fly production deploy
- Java legacy runtime 維護分支為 `legacy-java`；`master` 以 Go runtime 為主線
- Docker build 會把 `web/dist` 複製成 runtime `STATIC_DIR`，並保留 `/thumbs/*` 本地縮圖資產

## 5. 部署流程

1. 確認 Fly app secrets 已設定：`MONGO_URI`，以及 `DC_WEBHOOK_URL` 或 `DISCORD_TARGETS`；必要時設定 `HOST_URL`、`MONGO_DATABASE`。
2. 在 GitHub environment `hidedbot` 設定 Actions secret `FLY_API_TOKEN`。
3. 到 GitHub Actions 手動執行 `Deploy to Fly.io` workflow，輸入要顯示的 `bot_version`，預設為 `1.0.4`。
4. 若在已登入 Fly CLI 的本機部署，設定版本：`export BOT_VERSION=1.0.4`，再執行 `./build&deploy.sh`。
5. 部署後檢查 `/actuator/health`、首頁、`/HideBot/discord/version` 與 `/HideBot/discord/targets`。

## 6. 觀察項目

- 啟動時間
- 記憶體佔用
- webhook 成功/失敗率
- history 與 counter 是否一致

## 7. 回滾策略

若 Go 版部署失敗：
1. 將流量切回 Java 版穩定映像
2. 保留同一 MongoDB 資料
3. 重新檢查 webhook 送出與持久化順序
