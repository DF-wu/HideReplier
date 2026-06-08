# SRS — HideReplier Go Port

## 1. 文件目的

本文件定義 HideReplier Go 版本的功能需求與非功能需求。目標是保留現有產品行為：

- 靜態單頁前端
- 匿名 Discord webhook 發文
- 多 Discord server/channel 目標選擇
- 歷史查詢
- 版本查詢
- 健康檢查

## 2. 系統範圍

Go 版本必須取代目前 Spring Boot 執行路徑，並滿足以下範圍：

1. 提供與目前產品等價的 HTTP API。
2. 提供相同用途的靜態前端頁面與資產。
3. 保留 MongoDB 作為持久化來源。
4. 保留 Fly.io 1 CPU / 256MB 的部署目標。

不在本次範圍內：

- OAuth / 登入系統
- 分散式擴展

## 3. 使用者與情境

### 3.1 一般使用者
- 透過前端頁面填寫機器人名稱、顏色、頭像、圖片、內容。
- 預覽訊息外觀。
- 選擇可用的 Discord 發送目標並送出到 webhook。

### 3.2 維運者
- 透過 Fly.io 部署 Go 版本。
- 透過環境變數設定版本、MongoDB URI、Discord webhook URL 或多目標 webhook 清單。
- 透過健康檢查驗證服務狀態。

## 4. 功能需求

### FR-01 靜態首頁
系統必須在 `/` 提供單頁前端，內容包含：
- 頁首介紹區
- IP/區域資訊區
- 表單輸入區
- Discord 風格預覽區

### FR-02 預覽功能
系統必須允許使用者在送出前預覽以下欄位：
- 機器人名稱
- 顏色
- 頭像
- 圖片
- 縮圖
- 內容文字

### FR-03 匿名送出
系統必須接受 `POST /HideBot/discord` 並：
1. 驗證輸入資料。
2. 正規化顏色值。
3. 建立 Discord embed。
4. 以 webhook 傳送給 Discord。
5. 僅在 webhook 成功時才寫入歷史與序號狀態。

### FR-04 歷史查詢
系統必須在 `GET /HideBot/discord` 回傳歷史資料，並依 `serialNumber` 升冪排序。

### FR-05 版本查詢
系統必須在 `GET /HideBot/discord/version` 回傳目前版本字串。

### FR-06 健康檢查
系統必須在 `GET /actuator/health` 回傳最小健康 JSON，例如 `{"status":"UP"}`。

### FR-07 本地縮圖資產
系統必須提供本地縮圖資產於 `/thumbs/*.svg`，避免前端依賴第三方縮圖來源。

### FR-08 錯誤回報
前端在送出失敗時必須顯示失敗訊息，不能一律顯示成功。

### FR-09 多 Discord 目標
系統必須允許維運者用環境變數設定多個 Discord server/channel webhook 目標，並讓前端在送出前選擇目標。目標清單 API 不得洩漏 webhook URL。

## 5. 非功能需求

### NFR-01 記憶體限制
系統應可在 Fly.io shared-cpu-1x / 256MB RAM 下穩定運行。

### NFR-02 啟動時間
系統應在低規格 VM 上維持可接受啟動時間，足以通過 Fly 健康檢查寬限期。

### NFR-03 可維護性
Go 版本必須使用清楚的分層與文件，避免將 HTTP、Discord、Mongo、前端靜態服務混在單一檔案中。

### NFR-04 相容性
HTTP route、前端互動流程、資料結構語意需與現行版本相容。

### NFR-05 真實狀態一致性
若 webhook 傳送失敗，歷史資料與序號狀態不得被視為成功送出。

## 6. 驗收標準

本任務完成需同時滿足：

1. Go 專案可 build 成功。
2. Go 服務啟動後可提供首頁、API、health、thumbs 資產。
3. 成功 webhook 路徑回 200，失敗 webhook 路徑回非 2xx。
4. 成功提交才會持久化歷史資料。
5. 前端在 sandbox 環境完成實際瀏覽器驗證。
