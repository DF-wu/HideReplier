# Risk Register — Go Port

## R-01 MongoDB schema mismatch
- 風險：Go BSON/JSON 欄位名稱與既有 Java 寫入格式不一致。
- 影響：舊資料無法正常讀取。
- 緩解：明確定義 struct tag，使用既有欄位名稱。

## R-02 Webhook 成功與持久化狀態不一致
- 風險：送 Discord 失敗但 history 仍被寫入。
- 影響：資料可信度下降。
- 緩解：只有 webhook 成功後才寫 history 與 counter。

## R-03 記憶體超出 256MB
- 風險：Go 版實作不當導致 Fly.io 低配 VM 不穩定。
- 影響：啟動失敗或被 OOM。
- 緩解：使用標準庫與官方 driver、避免重量框架。

## R-04 前端與 API 行為不一致
- 風險：預覽與實際送出 payload 不一致。
- 影響：使用者誤判結果。
- 緩解：預覽與 submit 共用同一套正規化邏輯。

## R-05 外部依賴不穩定
- 風險：IP 查詢服務或 Discord webhook endpoint 不穩定。
- 影響：體驗下降。
- 緩解：前端要有 fallback，後端要有清楚錯誤訊號。
