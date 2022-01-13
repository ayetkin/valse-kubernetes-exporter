package mongodb

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"strings"
	"valse/pkg/config"
)

type OfficialClient interface {
	NewSession(dbName string) *mongo.Database
	NewSessionWithSecondaryPreferred(dbName string) *mongo.Database
	EnsureIndex(indexKeys []string, isUnique bool, indexName, dbName, collection string) error
}

func NewClient(appConfig *config.AppConfig, logger *logrus.Logger) OfficialClient {

	var (
		client *mongo.Client
		err    error
	)

	hosts := strings.Join(appConfig.MongoDB.Hosts, fmt.Sprintf(":%v,", appConfig.MongoDB.Port))
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", appConfig.MongoDB.Username, appConfig.MongoDB.Password, hosts, appConfig.MongoDB.Port)

	clientOptions := options.Client().ApplyURI(connectionString).SetReplicaSet(appConfig.MongoDB.ReplicaSet)

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), appConfig.MongoDB.TimeoutSeconds)
	defer cancel()

	if client, err = mongo.Connect(ctxWithTimeout, clientOptions); err != nil {
		logger.Fatalf("%v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	return &officialClient{
		client: client,
		logger: logger,
	}
}

type officialClient struct {
	client *mongo.Client
	logger *logrus.Logger
}

func (c *officialClient) NewSessionWithSecondaryPreferred(dbName string) *mongo.Database {

	secondary := readpref.SecondaryPreferred()
	dbOpts := options.Database().SetReadPreference(secondary)

	return c.client.Database(dbName, dbOpts)
}

func (c *officialClient) NewSession(dbName string) *mongo.Database {
	return c.client.Database(dbName)
}

func (c *officialClient) EnsureIndex(indexKeys []string, isUnique bool, indexName, dbName, collection string) error {

	serviceCollection := c.client.Database(dbName).Collection(collection)

	_, err := serviceCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    generateIndexKeys(indexKeys),
		Options: options.Index().SetName(indexName).SetUnique(isUnique)})

	if err != nil {
		return err
	}

	return nil
}

func generateIndexKeys(arr []string) bsonx.Doc {

	var keys bsonx.Doc

	for _, s := range arr {
		keys = append(keys, bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		})
	}

	return keys
}
