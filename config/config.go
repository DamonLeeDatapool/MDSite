package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"MServer/funcs"
	"MServer/globalVar"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

const (
	DBMAXLIFETIME int = 300
	DBMAXIDLECONN int = 30
	DBMAXOPENCONN int = 30
)

func LoadEnvAndSetupLogger() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unable to load .env file")
	}

	initGlobalVal()

	SetupLogger(globalVar.LogFileOrig)

	//將加載的環境變量顯示出來
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Println(err)
	}
	for key, value := range myEnv {
		Logger.Printf("name: %s ,value:%v", key, value)
	}

}

func initGlobalVal() {

	globalVar.USE_PORT = funcs.AorB(os.Getenv("USE_PORT") == "", globalVar.USE_PORT, os.Getenv("USE_PORT")).(string)

	//globalVar.ENV = strings.ToUpper(funcs.AorB(os.Getenv("ENVIRONMENT") == "", globalVar.ENV, os.Getenv("ENVIRONMENT")).(string))
}

func NewDatabase() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	Logger.Debugln(URL)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		Logger.Panic("Failed to connect to database")
	}

	MysqlDB, _ := db.DB()
	MysqlDB.SetConnMaxLifetime(time.Duration(DBMAXLIFETIME) * time.Second) //MaxLifetime=300
	MysqlDB.SetMaxIdleConns(DBMAXIDLECONN)                                 //MaxIdleConns=30
	MysqlDB.SetMaxOpenConns(DBMAXOPENCONN)                                 //MaxOpenConns=30

	Logger.Infoln("Database connection established")
	return db
}

func CloseDatabase() {
	MysqlDB, _ := DB.DB()
	MysqlDB.Close()
}

/*
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		DBName:   "rochung",
		Password: "1234",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
*/

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + globalVar.REDIS_PORT,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	//Logger.Println(pong, err)
	if err != nil {
		Logger.Panic("can't connect to Redis")
	}
	return client
}
