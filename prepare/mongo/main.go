package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
	"fmt"
)

var (
	client *mongo.Client
	err error
	database *mongo.Database
	collection *mongo.Collection
)

func main()  {
	//1, 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientopt.ConnectTimeout(5*time.Second)); err != nil{
		fmt.Println(err)
		return
	}

	//2, 选择数据库my_db
	database = client.Database("my_db")

	//3, 选择表my_collection
	collection = database.Collection("my_collection")

	//4, 输出表名
	fmt.Println(collection.Name())
}