# HideReplier

-----------------
**本repo code品質是黑歷史= =**
**This repo code quality is not good for mental health. Don't see.**

-----------------


<!-- To be compatible with GitHub, you should maintain a table of content by yourself. Recommend to use markdown extension in vscode. -->
- [HideReplier](#hidereplier)
  - [building chain](#building-chain)
  - [更新紀錄](#更新紀錄)
  - [重要聲明](#重要聲明)
  - [系統介紹](#系統介紹)
  - [簡介](#簡介)
  - [如何使用](#如何使用)
    - [介面介紹](#介面介紹)
      - [最新版介面](#最新版介面)
      - [新版介面(DEP)](#新版介面dep)
      - [舊版介面](#舊版介面)
    - [圖片連結說明](#圖片連結說明)
  - [進度](#進度)
    - [Contribution](#contribution)

---
## building chain
1. compile front-end pack (with tenor api key and Bot version in this stage)
   - api key in dc manager channel. bot version is managing by my hand.....
2. build maven project
3. build as docker image
4. push to docker hub
5. deploy to fly.io platform



## 更新紀錄
- Health action: v1.0.2
  - smaller image and runtime image usage now.


- Amazing announcement: v1.0.1
  - migrate the avatar and icon from discord cdn to my github image pool.
  - optimize the building time.
  - colorful building message.
  - fix some bugs.

- Initial Availability ?:) V1.0.0
  - refactor everything
  - recode everything
  - ~~leaving everything as legacy.~~
  - re-construct ci/cd chain.
- HOTFIX v0.6.0
  - update springboot to v3.0
  - dc webhook token被幹走了，修正漏洞。
- HOTFIX v0.5.0
  - update dependency
  - RIP heroku. migrate to fly.io
- HOTFIX v0.4.1
    - 修正小圖無法如所選的問題
- 隆重推出 v0.4  代號：凱留？
    - 新增小圖選項 現在可以用預設的小圖為你的訊息增添花樣
    - ~~有乾爹要贊助那個欄位嗎XD~~
    - ~~我跟公主連結不熟 圖是群裡抓的~~
- HOTFIX v0.3.1
    - 修正功能異常
        - PR沒看清楚 被陰了
- 隆重推出 v0.3 代號： AKS
    - 前端大更新 感謝@stanley2058的貢獻
    - 現在前端可以即時渲染出discord訊息預覽
    - 增加系統的說明
- 緊急修復v0.2.1.1
    - 修正頭像無功能問題
- 隆重推出v0.2.1 代號：感謝花生
    - 修正avatar功能 已可正常運作
    - 增加連結範圍 現在圖片以及avatar欄位之imgur連結皆會自動轉換至合法格式
        - 是**BBCODE** 請看使用說明的部分
    - 預計逐步完成說明文件以及更新文件~~如果我的隊友沒因為我寫這個機器人而把大專案晾在一旁而把我砍了~~
    - 其實是可以公開到Github了 但是現在repo裡面有db密碼 discordwebhook連結等等敏感資料
        - ~~懶人我有點想直接推上去......~~
- 隆重推出v0.2
    - 增加顏色記憶功能 自動記住你上次的顏色
    - 增加avatarUrl功能 （測試中）
- v0.1
    - 世界的初始

## 重要聲明
- 系統仍在開發中，請不要惡意破壞或惡意測試 DDOS 或大量的request都是不允許的。
- 請大家好好愛惜他 他會慢慢成長~~應該吧~~。
- 歡迎開Issue或發PR給我。

## 系統介紹
- ~~[機器人連結](https://hidedbot.herokuapp.com/)~~
- ↑我怕她被幹壞，還是把連結存在dc吧。

## 簡介
- ~~首先 你要是國立臺灣海洋大學資工系相關人事 不然可以左轉了~~
- 本專案是給 海大資工Discord群組使用
    - [GitHub 連結~~請去按星星~~](https://github.com/DF-wu/NTOUCS_DiscordInvitationPage)
    - ~~按完以後這個也要按~~
- [期末報告投影片連結](https://docs.google.com/presentation/d/1bW8c_qMBIMk7MU57SnitGGbVtT4q9azKrPZhveccIRA/edit)
    - 仍未更新成新版
    - 做的hen爛 因為整個專案到demo前20分鐘才算是能拿上來講.....
- 本機器人的開發動機：
    1. ~~WBSE課程說要寫專案~~
    2. 由於某資工系上Discord開了一個黑特區，很不巧的在這個群組中又是實名制，有時候無形的壓力會造成同學不敢抒發意見。因此開發一個可以匿名回覆的機器人
    3. ~~本來想要做一個discord平台的關鍵字alertor，很遺憾的走偏了變成寫回覆機器人~~

## 如何使用

### 介面介紹


#### 最新版介面
+ heroku涼了，遷移平台
+ 神說要換UI，於是有了新UI
    + ![UI](img/image.png)

#### 新版介面(DEP)
- 神說要有UI，於是有了UI。
- ~~有位朋友看不下去就生了新的UI出來了~~
- 由於是期末個人專案，要說明一下新版UI不是我做的QQ
- **可以及時渲染！！！！！**
    - ![](https://i.imgur.com/q0gZLma.png)

#### 舊版介面
- ~~一圖流~~
- ~~客家風格~~
- 簡潔明了
- ~~沒有時間寫css也不會寫....~~
- ![](https://i.imgur.com/JPOJWJM.jpg)

### 圖片連結說明
- 必須是靜態連結
    - 例如 `http://dfder.tw/87.jpg`
- 為了方便大家使用，針對[Imgur圖床](https://imgur.com)實做了轉換器
    - ~~迷之音：imgur很慢 94ㄌㄙ~~
- 使用步驟：
    1. 到imgur網站，點擊`new Post`![](https://i.imgur.com/rWKAPmW.jpg)
    2. 照指示上傳圖片 ![](https://i.imgur.com/SF1xZF5.png)
    3. 點`get share links`![](https://i.imgur.com/JQvCU3w.png)
    4. 複製`BBcode`連結並貼到回覆機器人之圖片連結欄位裡面即可![](https://i.imgur.com/FTNzuEt.png )

## 進度
- TODO:
    - [x] 歷史記錄功能
    - [ ] 前端顯示歷史記錄
    - [x] 歷史紀錄加入流水號
    - [x] 固定輸出格式
    - [x] ~~Color轉成java物件儲存~~
    - [x] 統一轉為HEX color儲存
    - [x] imgur平台圖片格式問題
    - [x] 前端顯示ip地址資訊（地區、電信業者等）
- TOFIX:
    - [x] 時區修正
    - [x] 把discord wehook的JSON annotation用java bean完整實做出來
    - [x] 找宇得大大救援findAll 的問題
- TOIMPROVE:
    - [ ] 可由流水號產生獨立url 並嵌入message
    - [ ] 改由後端負責ip相關資料解析
    - [ ] 解決各種趕工造成的技術債
    - [ ] 各種安全問題
    - [x] ip位置
    - [x] 把DiscordWebhook的功能完整做出來
    - [ ] 在前端實做 完整Discordwebhook功能操作
    - [ ] 多平台支援（Telegram等）
    - [ ] yude對imgur有意見 說要自己架圖床
    - [ ] ~~回想起一開始是要做alertor的初衷~~
    - [ ] 把敏感資料如DB密碼或webhook密碼從git大海中移除
 

### Contribution
@k3kdude for the DiscordWebhook Java base concept implement
@stanley2058 for the AMAZING web UI improvement.
@penut85420 for cleaning up code and optimize the github README.md 
