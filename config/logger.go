package config

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	//"MServer/globalVar"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

const (
	LogMaxTime      time.Duration = 60 * 24 * time.Hour
	LogRotationTime time.Duration = 3 * 24 * time.Hour
)

func SetupLogger(fileName string) {

	Logger = log.New()
	if fileName == "Stdout" {
		Logger.SetOutput(os.Stdout)
	} else {
		Writer, err := rotatelogs.New(
			fileName+"%Y%m%d.log", //%Y%m%d%H%M
			//在项目根目录下生成软链文件 latest_log.log 指向最新的日志文件。注意！！！必须在管理员权限下开终端启动。
			//rotatelogs.WithLinkName(fileName),
			//日志最大保存时间
			//rotatelogs.WithMaxAge(7*24*time.Hour),
			rotatelogs.WithMaxAge(LogMaxTime),
			////设置日志切割时间间隔(1天)(隔多久分割一次)
			rotatelogs.WithRotationTime(LogRotationTime),
		)
		if err != nil {
			log.Fatal("initial log fail, err:", err)
		}
		//設置同時檔案跟Screen output
		_multiWriter := io.MultiWriter(os.Stdout, Writer)

		//Logger.SetOutput(Writer)
		Logger.SetOutput(_multiWriter)

		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(Writer)
	}
	//顯示行數
	//if strings.ToUpper(globalVar.ENV) == "DETAIL" {
	//	Logger.SetReportCaller(true)
	//}

	Logger.SetFormatter(&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg% \n",
	})
	Logger.SetLevel(log.DebugLevel)
	gin.SetMode(gin.ReleaseMode)

}

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%lvl%]: %time% - %msg% \n"
	defaultTimestampFormat = time.RFC3339
)

type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}
