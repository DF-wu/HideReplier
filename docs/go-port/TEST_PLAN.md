# Test Plan — HideReplier Go Port

## 1. 目標

驗證 Go 版本是否在功能、部署、錯誤處理與前端互動上等價或優於現有版本。

## 2. 測試分層

### 2.1 Unit Tests
- 顏色正規化
- embed 建構
- BBCode 去除
- config 解析

### 2.2 Integration Tests
- Mongo 初始化與 counter 載入
- 成功 webhook 後寫入 history
- 失敗 webhook 不寫入 history
- version route
- health route

### 2.3 End-to-End Tests
- 首頁載入
- 預覽互動
- 成功提交
- 失敗提交

## 3. 驗收案例

### TC-01 健康檢查
- 呼叫 `/actuator/health`
- 預期 `200` 與 `{"status":"UP"}`

### TC-02 版本查詢
- 呼叫 `/HideBot/discord/version`
- 預期回傳版本字串

### TC-03 成功送出
- webhook 指向可回 200 的 sandbox endpoint
- 預期 API 回 200
- 預期 history 出現新紀錄

### TC-04 失敗送出
- webhook 指向可回 500 的 sandbox endpoint
- 預期 API 回非 2xx
- 預期 history 不增加

### TC-05 靜態資產
- `/`, `/index.css`, `/index.js`, `/thumbs/01.svg`
- 預期皆可正常提供

### TC-06 瀏覽器驗證
- 在 sandbox 瀏覽器環境檢查畫面層級、互動與送出狀態

## 4. 完成標準

所有關鍵案例通過後，Go port 才能視為完成。
