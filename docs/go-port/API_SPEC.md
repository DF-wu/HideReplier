# API Specification — HideReplier Go Port

## 1. GET /

### Purpose
提供前端頁面。

### Response
- `200 text/html`

## 2. GET /assets/*.css
提供 Vite build 後的前端樣式。

## 3. GET /assets/*.js
提供 Vite build 後的前端邏輯。

## 4. GET /thumbs/{name}.{svg,gif}

### Purpose
提供本地縮圖資產。

### Response
- `200 image/svg+xml`

## 5. GET /actuator/health

### Response
```json
{"status":"UP"}
```

## 6. GET /HideBot/discord/version

### Response
```json
"1.0.4"
```

## 7. GET /HideBot/discord

### Response Shape
```json
[
  {
    "id": "...",
    "timeStamp": 1710000000,
    "serialNumber": 1,
    "discordMessage": {
      "url": "https://hidereplier.fly.dev/",
      "content": "...",
      "username": "...",
      "tts": false,
      "targetId": "main",
      "extras": {
        "imgUrl": "...",
        "thumbnail": "...",
        "color": "8150271",
        "avatar_url": "",
        "ip": "127.0.0.1"
      },
      "avatarUrl": ""
    },
    "posterIp": "127.0.0.1",
    "responseCode": null
  }
]
```

## 8. POST /HideBot/discord

### Request Shape
```json
{
  "username": "ui-success",
  "content": "hello",
  "color": "#7c5cff",
  "avatar_url": "",
  "imgUrl": "https://i.imgur.com/zul9zzl.jpg",
  "ip": "127.0.0.1",
  "thumbnail": "http://127.0.0.1:18082/thumbs/01.svg",
  "targetId": "main"
}
```

### Success Response
- `200 application/json`

### Failure Response
- `500 application/json` when webhook dispatch fails or server-side processing fails.

### Notes
- `color` must be normalized from hex string to base-10 string before persistence/forwarding.
- `avatar_url` and `imgUrl` must follow the same parsed/normalized semantics used in preview.
- `targetId` is optional. When omitted, the configured default Discord target is used.
- webhook URLs are dispatch-only secrets and must not be returned from targets API or newly persisted history rows.
- successful and failed submissions must not share the same persistence outcome.

## 9. GET /HideBot/discord/targets

### Purpose
提供前端可選的 Discord 發送目標清單。回應不可包含 webhook URL。

### Response Shape
```json
[
  {
    "id": "main",
    "label": "Main server / anonymous channel",
    "default": true
  },
  {
    "id": "test",
    "label": "Test server / sandbox channel"
  }
]
```
