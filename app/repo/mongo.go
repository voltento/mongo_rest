package repo

import (
	"context"
	"errors"
	"github.com/pieterclaerhout/go-log"
	"github.com/voltento/mongo_rest/app/config"
	"github.com/voltento/mongo_rest/app/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Mongo struct {
	client    *mongo.Client
	cfg       config.Config
	healthyCB func(err error)
}

func (m *Mongo) HealthCheck() error {
	return m.client.Ping(context.TODO(), nil)
}

func NewMongo(cfg config.Config) *Mongo {
	mongoHost, err := buildMongoHost(cfg)
	if err != nil {
		log.Fatal(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoHost))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	m := &Mongo{
		client: client,
		cfg:    cfg,
	}

	if err := m.HealthCheck(); err != nil {
		log.Fatal(err)
	}
	return m
}

func buildMongoHost(cfg config.Config) (string, error) {
	const prefix = "mongodb://"
	address := prefix
	if len(cfg.MongoDBCreds) > 0 {
		address += cfg.MongoDBCreds + "@"
	}

	if len(cfg.MongoDBHost) == 0 {
		return "", errors.New("mongo db host is empty")
	}

	address += cfg.MongoDBHost
	return address, nil
}

func buildFindParams(filters *dto.Filters) (*options.FindOptions, bson.M) {
	opts := options.Find()
	opts.Limit = &filters.Limit

	conditions := bson.M{}

	if filters.Found != nil {
		conditions["found"] = *filters.Found
	}

	if filters.Number != nil {
		conditions["number"] = *filters.Number
	}

	if filters.Type != nil {
		conditions["type"] = *filters.Type
	}

	return opts, conditions
}

func (m *Mongo) FindRecords(filters *dto.Filters) []dto.Record {
	// TODO: configure request TTL from config
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	quickstartDatabase := m.client.Database(m.cfg.MongoDBName)
	numbers := quickstartDatabase.Collection(m.cfg.MongoCollectionName)

	findOptions, conditions := buildFindParams(filters)

	c, err := numbers.Find(ctx, conditions, findOptions)
	if err != nil {
		log.Errorf("can not perform find requet %v", err.Error())
	}
	defer c.Close(context.TODO())

	var results []dto.Record
	for c.Next(context.TODO()) {
		var elem dto.Record
		err := c.Decode(&elem)
		if err != nil {
			log.Errorf("can not parse the record %v", err.Error())
		}

		results = append(results, elem)
	}
	return results
}

func (m *Mongo) Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.client.Disconnect(ctx)
}
