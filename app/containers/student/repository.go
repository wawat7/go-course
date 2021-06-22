package student

import (
	"go-course/app/config"
	"go-course/app/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAll() (students []Student)
	Insert(student Student) (insertedID primitive.ObjectID)
	GetById(ID primitive.ObjectID) (student Student)
}

func NewStudentRepository(database *mongo.Database) Repository {
	return &repository{Collection: database.Collection("students")}
}

type repository struct {
	Collection *mongo.Collection
}

func (repository *repository) GetAll() (students []Student) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	exception.PanicIfNeeded(err)

	err = cursor.Err()
	exception.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var student Student
		cursor.Decode(&student)
		students = append(students, student)
	}

	return
}

func (repository *repository) Insert(student Student) (insertedID primitive.ObjectID) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	result, err := repository.Collection.InsertOne(ctx, student)
	exception.PanicIfNeeded(err)
	insertedID = result.InsertedID.(primitive.ObjectID)
	return
}

func (repository *repository) GetById(ID primitive.ObjectID) (student Student) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	err := repository.Collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&student)
	exception.PanicIfNeeded(err)
	return
}

