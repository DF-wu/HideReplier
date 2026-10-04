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
- `DISCORD_TARGETS`
- `MONGO_URI`
- `MONGO_DATABASE`

### 2.3 `internal/app`
提供 handlers：
- `/HideBot/discord` POST
- `/HideBot/discord` GET
- `/HideBot/discord/version` GET
- `/HideBot/discord/targets` GET
- `/actuator/health` GET

### 2.4 `internal/service`
封裝商業邏輯：
- 顏色正規化
- embed 建構
- webhook dispatch
- Discord target resolution
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
1. decode JSON request（body 上限 64KB）
2. normalize request values
3. resolve configured Discord target by `targetId`
4. 以 `findOneAndUpdate` + `$inc` 在 Mongo 原子性地保留流水號（單一 round trip，無全域鎖；多實例安全）
5. build embed payload
6. send webhook to Discord target
7. 若成功：寫入 `DiscordPostCollection`
8. 回傳前端可用 JSON

> webhook 失敗時流水號不會回收，會留下空號。這是用「請求可完全並行」換來的；
> 舊版以全域 mutex 把整段 webhook 往返串行化，在單核機器上會把所有發文排成一列。

### GET /HideBot/discord
1. Mongo 端依 `serialNumber` 排序（啟動時會嘗試建立該欄位索引）
2. 以 cursor 逐筆 decode 並直接串流寫出 JSON array，記憶體不隨歷史筆數成長
3. `?limit=N` 時只取最新 N 筆（上限 5000），仍以升冪回傳

### GET /HideBot/discord/targets
1. 從 runtime config 讀取 Discord targets
2. 移除 webhook URL
3. 回傳前端可顯示的 `id`、`label`、`default`

## 4. 靜態前端設計

production Docker build 由 `web/` React/Vite 專案產生 `web/dist`，並於 runtime 設為 `STATIC_DIR`。既有 `src/main/resources/static/thumbs/` 縮圖資產會被複製到 runtime static 目錄以保留 `/thumbs/*` 相容路徑。

- `/` → `index.html`
- `/assets/*.css`
- `/assets/*.js`
- `/thumbs/*.svg`
- `/thumbs/*.gif`
- `/icon.png`

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
