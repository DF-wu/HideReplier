# Migration Plan — Spring Boot to Go

## 1. 目標

把現有專案從 Spring Boot 移植為 Go 專案，同時保留 API、資料、前端與 Fly.io 部署目標。

## 2. 階段

### Phase 1 — 文件化
- 完成 SRS/SDS/API/DATA/TEST/DEPLOYMENT/RISK 文件

### Phase 2 — Go 骨架
- 建立 `go.mod`
- 建立 config/server/router/service/store 模組

### Phase 3 — API 與 Discord 邏輯
- 實作 `/HideBot/discord` POST/GET
- 實作 `/HideBot/discord/version`
- 實作 `/actuator/health`

### Phase 4 — 靜態前端與資產
- 提供 `src/main/resources/static/` 與 `thumbs/`

### Phase 5 — 驗證
- build/test
- 成功/失敗 webhook 路徑
- sandbox 瀏覽器驗證

## 3. 移植原則

1. 先保留功能，再談優化。
2. 不新增產品範圍。
3. 所有外部行為變更需有文件與驗證支撐。

## 4. 完成條件

- Go 版本可作為主要執行路徑
- 文件與實作一致
- Fly.io 可部署
