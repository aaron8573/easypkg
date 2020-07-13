/**
 * @Author: guomumin <aaron8573@gmail.com>
 * @File:  logs_test.go
 * @Version: 1.0.0
 * @Date: 2020/7/13 下午4:32
 * @Description:
 */

package log

import (
    "easypkg/log/async_file"
    "testing"
)

var (
    log    *Logger
    isInit bool
)

func logInit() {
    log = New(LogConfig{
        Type:         WRITE_LOG_TYPE_ASYNC,
        QueueSize:    1000000,
        BufferSize:   1 * 1024 * 1024, // 1MB
        FileFullPath: "log.log",
        SplitLogType: async_file.SPLIT_LOG_TYPE_NORMAL,
        Level:        0,
        Flag:         L_Time | L_LEVEL | L_SHORT_FILE,
    })
}

func TestNew(t *testing.T) {

    // 1000000
    // BufferSize:1MB 2.50s
    // BufferSize:2MB 2.37s
    logInit()

    for i := 0; i < 1000000; i++ {
        log.Info("test write log")
    }

    log.AsyncQuite()
}
