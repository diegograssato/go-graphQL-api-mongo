package mongo_api

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	db                  *mongo.Client
	collection          *mongo.Collection
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Name                string             `bson:"name"`
	ProgrammingLanguage string             `bson:"programmingLanguage"`
}

func InitMongo(db *mongo.Client, collection *mongo.Collection) *Account {
	return &Account{db: db, collection: collection}

}

func (a *Account) GetAll() ([]Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//collection := a.db.Database().Collection("graphQL")
	filter := bson.M{}

	cur, err := a.collection.Find(ctx, filter)
	if err != nil {
		log.Fatal("** 4 **", err)
	}
	defer cur.Close(ctx)

	var accounts []Account
	for cur.Next(ctx) {
		var tmp Account
		if err := cur.Decode(&tmp); err != nil {
			log.Fatal("** 5 **", err)
		}

		accounts = append(accounts, tmp)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("** 6 **", err)
	}

	return accounts, nil
}

func (a *Account) Insert(name, programmingLanguage string) (Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//collection := a.db.Database("api").Collection("graphQL")
	insert := Account{
		ID:                  primitive.NewObjectID(),
		Name:                name,
		ProgrammingLanguage: programmingLanguage,
	}

	_, err := a.collection.InsertOne(ctx, insert)
	if err != nil {
		return Account{}, err
	}

	return insert, nil
}
