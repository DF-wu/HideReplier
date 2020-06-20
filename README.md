---
tags: HideReplier
---

# HideReplier

[TOC]
---
### 重要聲明
- 系統仍在開發中，請不要惡意破壞或惡意測試 DDOS 或 大量的request都是不允許的。
- 請大家好好愛惜他 他會慢慢成長~~應該吧~~。
- 歡迎開issue或發PR給我。

# 系統介紹
- [機器人連結](https://hidedbot.herokuapp.com/)
## 簡介
- ~~首先 你要是國立臺灣海洋大學資工系相關人事 不然可以左轉了~~
- 本專案是給 海大資工discord群組使用 
    - [github連結~~請去按星星~~ ](https://github.com/DF-wu/NTOUCS_DiscordInvitationPage)
    - ~~按完以後這個也要按~~

- [期末報告投影片連結](https://docs.google.com/presentation/d/1bW8c_qMBIMk7MU57SnitGGbVtT4q9azKrPZhveccIRA/edit)
    - 做的hen爛 因為整個專案到demo前20分鐘才算是能拿上來講.....

- 本機器人的開發動機：
    1. ~~WBSE課程說要寫專案~~
    2. 由於某資工系上Discord開了一個黑特區，很不巧的在這個群組中又是實名制，有時候無形的壓力會造成同學不敢抒發意見。因此開發一個可以匿名回覆的機器人
    3. ~~本來想要做一個discord平台的關鍵字alertor，很遺憾的走偏了變成寫回覆機器人~~


## 如何使用

### 介面介紹

#### 新版介面
- 神說要有UI，於是有了UI。
- ~~有位朋友看不下去就生了新的UI出來了~~
- 由於是期末個人專案，要說明一下那不是我做的QQ
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


## TODO
- TODO:
    - [x] 歷史記錄功能
    - [ ] 前端顯示歷史記錄
    - [x] 歷史紀錄加入流水號
    - [x] 固定輸出格式
    - [x] ~~Color轉成java物件儲存~~
    - [x] 統一轉為HEX color儲存
    - [x] imgur平台圖片格式問題
- TOFIX:
    - [x] 時區修正
    - [x] 把discord wehook的JSON annotation用java bean完整實做出來
    - [x] 找宇得大大救援findAll 的問題
- TOIMPROVE:
    - [ ] 解決各種趕工造成的技術債
    - [ ] 各種安全問題
    - [x] ip位置
    - [x] 把DiscordWebhook的功能完整做出來
    - [ ] 在前端實做 完整Discordwebhook功能操作
    - [ ] 多平台支援（Telegram等）
    - [ ] yude對imgur有意見 說要自己架圖床
    - [ ] ~~回想起一開始是要做alertor的初衷~~
    - [ ] 把敏感資料如DB密碼或webhook密碼從git大海中移除

## 更新紀錄

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

### Contribution
@k3kdude for the DiscordWebhook Java base concept implement 

@stanley2058 for the AMAZING web UI improvement.

