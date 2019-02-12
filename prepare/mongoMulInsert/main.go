package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

var (
	client     *mongo.Client
	err        error
	database   *mongo.Database
	collection *mongo.Collection
	record     *LogRecord
	logArr     []interface{}
	result     *mongo.InsertManyResult
	insertId   interface{}
	docId      objectid.ObjectID
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   //任务名
	Command   string    `bson:"command"`   //shell命令
	Err       string    `bson:"err"`       //脚本错误
	Content   string    `bson:"content"`   //脚本输出
	TimePoint TimePoint `bson:"timePoint"` //执行时间点

}

func main() {
	//1, 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	//2, 选择数据库cron
	database = client.Database("cron")

	//3, 选择表log
	collection = database.Collection("log")

	//4, 插入记录bson
	record = &LogRecord{
		JobName: "job10",
		Command: "echo 10",
		Err:     "",
		Content: "hello 10",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 20,
		},
	}

	//5, 批量插入多条document
	logArr = []interface{}{record, record, record}

	//发起插入
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}

	//
	for _, insertId = range result.InsertedIDs {
		docId = insertId.(objectid.ObjectID)
		fmt.Println("自增ID:", docId.Hex())
	}
}
