# SDS — HideReplier Go Port

## 1. 設計目標

Go 版本採用低依賴、低記憶體、容易部署的設計：

- HTTP 層：Go 標準庫 `net/http`
- 路由：標準庫 mux + 明確 handler 綁定
- MongoDB：官方 `go.mongodb.org/mongo-driver`
- 靜態檔案：`http.FileServer`
- Discord webhook：標準庫 `net/http` client

## 2. 模組拆分

### 2.1 `cmd/server`
程式進入點，負責：
- 讀取設定
- 建立 Mongo client
- 組裝 services 與 handlers
- 啟動 HTTP server

### 2.2 `internal/config`
負責解析：
- `PORT`
- `BOT_VERSION`
- `HOST_URL`
- `DC_WEBHOOK_URL`
- `MONGO_URI`
- `MONGO_DATABASE`

### 2.3 `internal/app`
提供 handlers：
- `/HideBot/discord` POST
- `/HideBot/discord` GET
- `/HideBot/discord/version` GET
- `/actuator/health` GET

### 2.4 `internal/service`
封裝商業邏輯：
- 顏色正規化
- embed 建構
- webhook dispatch
- 成功後持久化
- counter 載入/更新

### 2.5 `internal/store`
封裝 Mongo collection 操作：
- history repository
- counter repository

### 2.6 Discord webhook 子系統
目前實作位於 `internal/service` 內，負責 Discord payload 建構與 webhook 傳送。

## 3. 請求流程

### POST /HideBot/discord
1. decode JSON request
2. normalize request values
3. load and increment serial counter in memory
4. build embed payload
5. send webhook to Discord target
6. 若成功：寫入 `DiscordPostCollection`
7. 若成功：更新 `Counter`
8. 回傳前端可用 JSON

### GET /HideBot/discord
1. 從 Mongo 讀取所有紀錄
2. 依 `serialNumber` 排序
3. 轉為 JSON 回傳

## 4. 靜態前端設計

靜態檔案沿用現有 `src/main/resources/static/` 內容，於 Go 版本中直接作為公開目錄提供：

- `/` → `index.html`
- `/index.css`
- `/index.js`
- `/thumbs/*.svg`

## 5. 失敗處理策略

### 5.1 Discord webhook 失敗
- 回傳非 2xx
- 不寫 history
- 不更新 counter

### 5.2 Mongo 失敗
- 回傳 500
- 若發生於 webhook 成功後的持久化階段，需明確記錄 log

### 5.3 前端請求失敗
- 回傳錯誤狀態碼
- 前端顯示 inline error state

## 6. 資源限制設計

為符合 Fly 256MB 目標：

- 避免大型 framework
- 避免 template engine
- 避免高併發 goroutine fan-out
- 使用單一 binary 部署

## 7. 相容性要求

Go 版本必須維持：

- route 路徑不變
- version route 不變
- health route 不變
- frontend 與 API 的互動意圖不變
- Mongo collection 名稱與欄位語意大致不變
