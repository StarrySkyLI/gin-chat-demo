package conf

import (
	"chat/model"
	"context"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo" //MongoDB的Go驱动包
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	MongoDBClient *mongo.Client
	AppMode       string
	HttpPort      string
	Db            string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassWord    string
	DbName        string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string

	MongoDBName string
	MongoDBAddr string
	MongoDBPwd  string
	MongoDBPort string
)

func Init() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadMongoDB(file)
	loadRabbitMQ(file)

	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err) //日志内容
		panic(err)
	}
	//MySQL
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	//MongoDB
	MongoDB()
}

func MongoDB() {
	// 设置mongoDB客户端连接信息
	clientOptions := options.Client().ApplyURI("mongodb://" + MongoDBAddr + ":" + MongoDBPort)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
	}
	err = MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		logging.Info(err)
	}
	logging.Info("MongoDB Connect")
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
func loadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost ").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()

}
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadMongoDB(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").String()
}
