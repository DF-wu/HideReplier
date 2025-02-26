#!/bin/bash

# 設定錯誤時立即退出
set -e

# 檢查 jar 檔案是否存在
if [ ! -f "app.jar" ]; then
    echo "錯誤: app.jar 不存在"
    exit 1
fi

# JVM options for fly.io low spec environment
# UseContainerSupport: 啟用容器感知功能
# MaxRAMPercentage: 設定最大記憶體使用率為50%
# InitialRAMPercentage: 設定初始記憶體使用率為50%
# Xss256k: 設定每個執行緒的堆疊大小為256KB
# UseSerialGC: 使用串行垃圾回收器以減少記憶體使用
# TieredCompilation: 啟用分層編譯以優化效能

JAVA_OPTS="\
    -XX:+UseContainerSupport \
    -XX:+UseSerialGC \
    -Xss512k


# 確保變數不為空
if [ -z "${JAVA_OPTS}" ]; then
    echo "警告: JAVA_OPTS 為空"
    exit 1
fi

# 輸出啟動信息
echo "正在使用以下 JVM 選項啟動:"
echo "${JAVA_OPTS}"

exec java ${JAVA_OPTS} -Dspring.profiles.active=application -jar app.jar