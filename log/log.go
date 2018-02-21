package log

import (
    "fmt"
    "runtime"
    "strings"

    "github.com/sirupsen/logrus"
    //"github.com/johntdyer/slackrus"
)

// logrus.Fields is just a map[string]interface{}
type Fields logrus.Fields
type Logger *logrus.Logger

var logger *logrus.Logger = nil

func New() {
    if logger == nil {
        logger = logrus.New()
    }
}

func GetLogger() Logger {
    return logger
}

//https://hooks.slack.com/services/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
// SetLogLevel sets level to logrus
func SetLogLevel(level string) error {
    l, err := logrus.ParseLevel(level)
    if err != nil {
        return err
    }

    logger.Level = l

    return nil
}

//func SetHookSlack(hookurl string, acceptedlevels string, channel string, iconemoji string, username string) {
    //l, err := logrus.ParseLevel(acceptedlevels)
    //if err != nil {
        //l = logrus.InfoLevel
    //}

    //logger.Hooks.Add(&slackrus.SlackrusHook{
        //HookURL: hookurl ,
        //AcceptedLevels: slackrus.LevelThreshold(l),
        //Channel: channel,
        //IconEmoji: iconemoji,
        //Username: username,
    //})
//}

// SetLogFormatter set formatter for logrus
func SetLogFormatter(formatter logrus.Formatter) {
    logger.Formatter = formatter
}

func fileInfo(skip int) string {
    _, file, line, ok := runtime.Caller(skip)
    if !ok {
        file = "<???>"
        line = 1
    } else {
        slash := strings.LastIndex(file, "/")
        if slash >= 0 {
            file = file[slash+1:]
        }
    }
    return fmt.Sprintf("%s:%d", file, line)
}

func SetFileEntry() *logrus.Entry {
    //entry := logger.WithField("file", fileInfo(2))
    entry := logger.WithFields(logrus.Fields{})
    return entry
}

func SetFileEntryWithFields(fields Fields) *logrus.Entry {
    entry := logger.WithFields(logrus.Fields(fields))
    //entry.Data["file"] = fileInfo(2)
    return entry
}

func Debug(args... interface{}) {
    SetFileEntry().Debug(args...)
}

func Debugln(args... interface{}) {
     SetFileEntry().Debugln(args...)
}

func Debugf(format string, args... interface{}) {
     SetFileEntry().Debugf(format, args...)
}

func DebugWithFields(fields Fields, args... interface{}) {
    SetFileEntryWithFields(fields).Debug(args...)
}

func Info(args... interface{}) {
    SetFileEntry().Info(args...)
}

func Infoln(args... interface{}) {
     SetFileEntry().Infoln(args...)
}

func Infof(format string, args... interface{}) {
     SetFileEntry().Infof(format, args...)
}

func InfoWithFields(fields Fields, args... interface{}) {
    SetFileEntryWithFields(fields).Info(args...)
}

func Warn(args... interface{}) {
    SetFileEntry().Warn(args...)
}

func Warnln(args... interface{}) {
     SetFileEntry().Warnln(args...)
}

func Warnf(format string, args... interface{}) {
     SetFileEntry().Warnf(format, args...)
}

func WarnWithFields(fields Fields, args... interface{}) {
    SetFileEntryWithFields(fields).Warn(args...)
}

func Error(args... interface{}) {
    SetFileEntry().Error(args...)
}

func Errorln(args... interface{}) {
     SetFileEntry().Errorln(args...)
}

func Errorf(format string, args... interface{}) {
     SetFileEntry().Errorf(format, args...)
}

func ErrorWithFields(fields Fields, args... interface{}) {
    SetFileEntryWithFields(fields).Error(args...)
}

func Fatal(args... interface{}) {
    SetFileEntry().Fatal(args...)
}

func Fatalln(args... interface{}) {
     SetFileEntry().Fatalln(args...)
}

func Fatalf(format string, args... interface{}) {
     SetFileEntry().Fatalf(format, args...)
}

func FatalWithFields(fields Fields, args... interface{}) {
    SetFileEntryWithFields(fields).Fatal(args...)
}
