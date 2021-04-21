package log

import "github.com/sirupsen/logrus"

func Trace(v ...interface{}) {
	logrus.Trace(v...)
}

func Debug(v ...interface{}) {
	logrus.Debug()
}

func Info(v ...interface{}) {
	logrus.Info(v...)
}

func Warn(v ...interface{}) {
	logrus.Warn(v...)
}

func Error(v ...interface{}) {
	logrus.Error(v...)
}

func Fatal(v ...interface{}) {
	logrus.Fatal(v...)
}

func Panic(v ...interface{}) {
	logrus.Panic(v...)
}

func InitLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 显示调用文件名和行号
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)

	// todo：时间格式化
	// todo：从配置文件读取日志输出级别并配置
	// todo：使用hook将日志(含有什么的前缀)分发到不同位置文件系统/redis/钉钉/mongodb/elasticsearch
}
