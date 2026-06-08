# Go Port Documentation Set

本目錄收錄 Go 版本移植文件。Go runtime 現在是 `master` 的主要 production runtime；Java/Spring Boot legacy runtime 保留在 `legacy-java` 分支維護。

## 文件清單

- `SRS.md`：Software Requirements Specification
- `SDS.md`：Software Design Specification
- `API_SPEC.md`：HTTP API 規格
- `DATA_MODEL.md`：資料模型與持久化規格
- `TEST_PLAN.md`：測試計畫與驗收方式
- `DEPLOYMENT_RUNBOOK.md`：部署與營運文件
- `MIGRATION_PLAN.md`：由 Java/Spring Boot 移植到 Go 的執行計畫
- `RISK_REGISTER.md`：主要風險與因應策略

## 文件目的

這一組文件的目的是把現有 Spring Boot + MongoDB + 靜態前端的匿名回覆機器人，完整定義成一個可實作、可驗證、可部署的 Go 專案。後續 Go 程式碼必須以這些文件為準，而不是邊寫邊猜。
