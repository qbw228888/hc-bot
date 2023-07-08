package api

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"hc-bot/config"
	"sync"
)

// CharacterId 人物id与群号和qq号的映射
type CharacterId struct {
	GroupId string `json:"group_id" bson:"group_id"`
	UserId  string `json:"user_id" bson:"user_id"`
	CharId  string `json:"char_id" bson:"char_id"`
}

// Attributes 人物表的对象
type Attributes struct {
	charId string         `json:"char_id" bson:"char_id"`
	Status map[string]int `json:"status" bson:"status"`
}

type GreatSuccessAndFail struct {
	GroupId      string `json:"group_id" bson:"group_id"`
	GreatSuccess int    `json:"great_success" bson:"great_success"`
	GreatFail    int    `json:"great_fail" bson:"great_fail"`
}

type IsOn struct {
	GroupId string `json:"group_id" bson:"group_id"`
	IfOn    bool   `json:"if_on" bson:"if_on"`
}

var database *mongo.Database

// 用于懒加载的锁
var mutex sync.Mutex = sync.Mutex{}

func initDatabase() {
	mutex.Lock()
	if database == nil {
		option := options.Client().ApplyURI(config.MongoUrl)
		client, err := mongo.Connect(context.Background(), option)
		if err != nil {
			fmt.Println("MongoDB连接失败")
		}
		err = client.Ping(context.Background(), nil) //ping通才代表连接成功
		if err != nil {
			fmt.Println("MongoDBping失败")
		}
		database = client.Database("qqbot")
	}
	mutex.Unlock()
}

func mapUnion(map1, map2 map[string]int) map[string]int {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

// InsertPlayerAttributes 插入玩家人物数据
func InsertPlayerAttributes(groupId string, userId string, status map[string]int) {
	if database == nil {
		initDatabase()
	}
	charIdCollection := database.Collection("char_id")
	var charId CharacterId
	err := charIdCollection.FindOne(context.TODO(), bson.D{{"group_id", groupId}, {"user_id", userId}}).Decode(&charId)
	charCollection := database.Collection("character")
	if err != nil {
		uuid := uuid.NewV4().String()
		charIdCollection.InsertOne(context.TODO(), bson.M{"group_id": groupId, "user_id": userId, "char_id": uuid})
		charCollection.InsertOne(context.TODO(), bson.M{"char_id": uuid, "status": status})
	} else {
		var attributes Attributes
		err := charCollection.FindOne(context.TODO(), bson.D{{"char_id", charId.CharId}}).Decode(&attributes)
		if err != nil {
			charCollection.InsertOne(context.TODO(), bson.M{"char_id": charId.CharId, "status": status})
		} else {
			charCollection.UpdateOne(context.TODO(), bson.M{"char_id": charId.CharId}, bson.M{"$set": bson.M{"char_id": charId.CharId, "status": mapUnion(attributes.Status, status)}})
		}
	}
}

// FindPlayerAttributes 查询玩家任务数据
func FindPlayerAttributes(groupId string, userId string) map[string]int {
	if database == nil {
		initDatabase()
	}
	charIdCollection := database.Collection("char_id")
	characterCollection := database.Collection("character")
	var charId CharacterId
	err := charIdCollection.FindOne(context.TODO(), bson.D{{"group_id", groupId}, {"user_id", userId}}).Decode(&charId)
	if err != nil {
		fmt.Println(groupId)
		return nil
	}
	var attributes Attributes
	err = characterCollection.FindOne(context.TODO(), bson.D{{"char_id", charId.CharId}}).Decode(&attributes)
	if err != nil {
		fmt.Println(charId.CharId)
		return nil
	}
	return attributes.Status
}

// GetGreatSuccess 获取本群大成功数据
func GetGreatSuccess(groupId string) int {
	if database == nil {
		initDatabase()
	}
	collection := database.Collection("great")
	if collection == nil {
		return -1
	}
	var greatSF GreatSuccessAndFail
	err := collection.FindOne(context.TODO(), bson.D{{"group_id", groupId}}).Decode(&greatSF)
	if err != nil {
		return -1
	}
	if greatSF.GreatSuccess == 0 {
		return -1
	}
	return greatSF.GreatSuccess
}

// SetGreatSuccess 设置本群大成功
func SetGreatSuccess(groupId string, greatSuccess int) {
	if database == nil {
		initDatabase()
	}
	collection := database.Collection("great")
	var greatSF GreatSuccessAndFail
	err := collection.FindOne(context.TODO(), bson.D{{"group_id", groupId}}).Decode(&greatSF)
	if err != nil {
		collection.InsertOne(context.TODO(), bson.M{"group_id": groupId, "great_success": greatSuccess})
	} else {
		collection.UpdateOne(context.TODO(), bson.M{"group_id": groupId}, bson.M{"$set": bson.M{"group_id": groupId, "great_success": greatSuccess, "great_fail": greatSF.GreatFail}})
	}
}

// GetGreatFail 获取本群大失败数据
func GetGreatFail(groupId string) int {
	if database == nil {
		initDatabase()
	}
	collection := database.Collection("great")
	if collection == nil {
		return -1
	}
	var greatSF GreatSuccessAndFail
	err := collection.FindOne(context.TODO(), bson.D{{"group_id", groupId}}).Decode(&greatSF)
	if err != nil {
		return -1
	}
	if greatSF.GreatFail == 0 {
		return -1
	}
	return greatSF.GreatFail
}

// SetGreatFail 设置本群大失败
func SetGreatFail(groupId string, greatFail int) {
	if database == nil {
		initDatabase()
	}
	collection := database.Collection("great")
	var greatSF GreatSuccessAndFail
	err := collection.FindOne(context.TODO(), bson.D{{"group_id", groupId}}).Decode(&greatSF)
	if err != nil {
		collection.InsertOne(context.TODO(), bson.M{"group_id": groupId, "great_fail": greatFail})
	} else {
		collection.UpdateOne(context.TODO(), bson.M{"group_id": groupId}, bson.M{"$set": bson.M{"group_id": groupId, "great_success": greatSF.GreatSuccess, "great_fail": greatFail}})
	}
}

func GetGreatSuccessAndGreatFail(groupId string) (int, int) {
	greatSuccess := GetGreatSuccess(groupId)
	if greatSuccess == -1 {
		greatSuccess = config.GreatSuccess
	}
	greatFail := GetGreatFail(groupId)
	if greatFail == -1 {
		greatFail = config.GreatFail
	}
	return greatSuccess, greatFail
}

func GetIsOn(groupId string) bool {
	if database == nil {
		initDatabase()
	}
	collection := database.Collection("is_on")
	var isOn IsOn
	err := collection.FindOne(context.TODO(), bson.D{{"group_id", groupId}}).Decode(&isOn)
	if err != nil {
		return true
	}
	return isOn.IfOn
}

func SetIsOn(groupId string, ifOn bool) {
	if database == nil {
		initDatabase()
	}
	collection := database.Collection("is_on")
	var isOn IsOn
	err := collection.FindOne(context.TODO(), bson.D{{"group_id", groupId}}).Decode(&isOn)
	if err != nil {
		collection.InsertOne(context.TODO(), bson.M{"group_id": groupId, "if_on": ifOn})
	} else {
		collection.UpdateOne(context.TODO(), bson.M{"group_id": groupId}, bson.M{"$set": bson.M{"group_id": groupId, "if_on": ifOn}})
	}
}
