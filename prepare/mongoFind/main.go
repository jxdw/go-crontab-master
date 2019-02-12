package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	//"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"time"
)

var (
	client     *mongo.Client
	err        error
	database   *mongo.Database
	collection *mongo.Collection
	cond       *FindByJobName
	cursor     *mongo.Cursor
	record     *LogRecord
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

// jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"` // jobName赋值为job10
}

func main() {
	//1, 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", options.Client().SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	//2, 选择数据库cron
	database = client.Database("cron")

	//3, 选择表log
	collection = database.Collection("log")

	//4, 按照jobName字段过滤，想找出jobName=job10,找出5条
	cond = &FindByJobName{JobName: "job10"}

	var findOption = options.Find()
	//5, 查询(过滤+翻页参数)
	if cursor, err = collection.Find(context.TODO(), cond, findOption.SetSkip(0), findOption.SetLimit(10)); err != nil {
		fmt.Println(err)
		return
	}

	//延迟释放游标
	defer cursor.Close(context.TODO())

	//6, 遍历结果集
	for cursor.Next(context.TODO()) {
		//定义一个日志对象
		record = &LogRecord{}

		//反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}

		//把日志行打印出来
		fmt.Println(*record)
	}
}
