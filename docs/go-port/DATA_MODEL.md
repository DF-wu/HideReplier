# Data Model — HideReplier Go Port

## 1. MongoDB Database

- Database name default: `DiscordBotHistoryDB`

## 2. Collection: Counter

### Document
```json
{
  "_id": "...",
  "counter": 3599
}
```

### Purpose
保存目前流水號。

## 3. Collection: DiscordPostCollection

### Document
```json
{
  "_id": "...",
  "timeStamp": 1710000000,
  "serialNumber": 3599,
  "discordMessage": {
    "url": "https://hidereplier.fly.dev/",
    "content": "ui success",
    "username": "ui-success",
    "tts": false,
    "targetId": "main",
    "extras": {
      "imgUrl": "https://i.imgur.com/zul9zzl.jpg",
      "thumbnail": "http://127.0.0.1:18082/thumbs/01.svg",
      "color": "8150271",
      "avatar_url": "",
      "ip": "127.0.0.1"
    },
    "avatarUrl": ""
  },
  "posterIp": "127.0.0.1",
  "responseCode": null
}
```

## 4. Behavioral Rules

1. `serialNumber` 必須只在成功 webhook 後被持久化。
2. `timeStamp` 使用台灣時區 +8 對應的 epoch second。
3. 歷史資料查詢需依 `serialNumber` 升冪排序。

## 5. Go Struct Mapping Principles

- 對外 JSON 欄位名稱維持與現有前端/回應相容。
- Mongo BSON 標記應明確指定，避免欄位名稱漂移。
- `extras` 使用 map 結構保存動態欄位。
