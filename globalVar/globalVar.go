package globalVar

//"context"
//"errors"
//"net/http"
//"time"
//"github.com/gin-gonic/gin"

//.env variable
var (
	USE_PORT   string = "8080"
	REDIS_PORT string = "6379"
)

var (
	JWTSecretKey string
	//LogFileOrig string = "Stdout" //"log檔名" or "Stdout" (目前使用log檔名會同時輸出在Stdout)
	LogFileOrig string = "./db_data/log"
	//DeviceConfigurationFile string = "./db/gateways.json"
)
