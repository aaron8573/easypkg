# log pkg

## 异步写日志

通过配置，日志写入队列，然后再从队列存储到内存buffer中，最后从buffer按块落到磁盘。



## 支持：
1、日志按天、小时分割，默认不分割
2、支持异步按buffer大小落磁盘。


## 例子：
```
import "easypkg/log"

logs = log.New(LogConfig{
        Type:         WRITE_LOG_TYPE_ASYNC, // 异步写文本
        QueueSize:    1000000, // 队列大小
        BufferSize:   1 * 1024 * 1024, // 1MB
        FileFullPath: "log.log", // 日志文件
        SplitLogType: async_file.SPLIT_LOG_TYPE_NORMAL, // 分割日志方式 SPLIT_LOG_TYPE_NORMAL -- 默认不分割
                                                        // SPLIT_LOG_TYPE_DAY -- 按天分割
                                                        // SPLIT_LOG_TYPE_HOUR -- 按小时分割
        Level:        0, // 日志级别  0-Debug,1-Info,2-Warn,3-Error,4-Fatal,5-Panic
        Flag:         L_Time | L_LEVEL | L_SHORT_FILE, // 日志标记
                                                        // L_Time ——— 日志时间
                                                        // L_LEVEL ———— 日志级别
                                                        // L_SHORT_FILE ———— 短日志文件
                                                        // L_LONG_FILE ———— 长日志文件
    })

logs.Info("test write log")

// 程序退出时，通知日志队列退出
logs.AsyncQuite()
```