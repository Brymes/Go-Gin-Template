package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	MongoClient *mongo.Database
	DBClient    *gorm.DB
	RedisClient *redis.Client
)

func InitDB() {
	InitMongo()
	InitMySQLDb()
	InitRedis()
}
func InitMongo() {
	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		uri = "mongodb://127.0.0.1:27017"
		//uri = "mongodb://host.docker.internal:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverApiOptions := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverApiOptions))

	if err != nil {
		log.Println("Error connecting to MongoDB 🤕")
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println("Error Connecting to Mongo")
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB 🛢️")

	MongoClient = client.Database("App-Name")
	//MissionCollection = MongoClient.Collection("missions")
}

// GetCollection function to get database collections
func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Collection(collectionName)
}

func InitMySQLDb() {
	var err error
	uri := os.Getenv("DATABASE_URI")

	if uri == "" {
		uri = "root:localpass@tcp(127.0.0.1:3306)/local?charset=utf8mb4&parseTime=True&loc=Local"
		//uri = "root:mypass@tcp(host.docker.internal:3306)/rewards_cwa?charset=utf8mb4&parseTime=True&loc=Local"
	}

	DBClient, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Println("Error Connecting to Database")
		log.Fatal(err)
	}
}

func InitRedis() {
	ctx := context.Background()
	uri := os.Getenv("REDIS_URL")

	if uri == "" {
		uri = "localhost:6379"
		// uri = "host.docker.internal:6379"
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     uri,
			Password: "",
			DB:       0,
		})

	} else {
		// works for external Redis Clients
		opt, _ := redis.ParseURL(uri)
		RedisClient = redis.NewClient(opt)
	}

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Println(pong)
		log.Println(uri)
		log.Println("Error Connecting to Redis")
		log.Fatal(err)
	}
	//RedisClient.FlushAll(context.TODO())
}
