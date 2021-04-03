package infrastructure

import "go.mongodb.org/mongo-driver/mongo"

type Options func(reps *Repositories)

func WithMongoDB(client *mongo.Client, dbName string) Options {
	return func(reps *Repositories) {
		reps.Persistence.mongoDB = client.Database(dbName)
	}
}
