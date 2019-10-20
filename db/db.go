package db

import (
	"context"
	"log"

	"github.com/faozimipa/monggo/models"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = func() context.Context {
	return context.Background()
}()

func connect() (*mongo.Database, error) {
	client, err := mongo.NewClient(
		options.Client().
			ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func collection() *mongo.Collection {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	return db.Collection("student")
}

/*GetAllStudent get all
 */
func GetAllStudent() ([]*models.Student, error) {
	findOptions := options.Find()
	findOptions.SetLimit(20)

	var res []*models.Student
	cur, err := collection().
		Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem models.Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &elem)
	}

	cur.Close(context.TODO())

	return res, nil
}

// GetOne returns a single student from the database.
func GetOne(id string) (*models.Student, error) {
	var res models.Student
	// objID, _ := primitive.ObjectIDFromHex(id)
	objID, _ := uuid.FromString(id)
	err := collection().
		FindOne(context.TODO(), bson.M{"uuid": objID}).
		Decode(&res)

	if err != nil {
		log.Fatal(err)
	}

	return &res, err
}

// Save inserts an item to the database.
func Save(item models.Student) error {
	_, err := collection().InsertOne(ctx, item)
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

/*Update update data
 */
func Update(id string, item models.Student) error {

	objID, _ := uuid.FromString(id)
	var selector = bson.M{"uuid": objID}

	_, err := collection().UpdateOne(ctx, selector, bson.M{"$set": item})
	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}

/*Delete student
 */
func Delete(id string) error {

	objID, _ := uuid.FromString(id)
	selector := bson.M{"uuid": objID}
	_, err := collection().DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}
