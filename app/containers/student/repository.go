package student

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	GetAll() []Student
}

func NewStudentRepository(database *mongo.Database) Repository {
	return &repository{Collection: database.Collection("students")}
}

type repository struct {
	Collection *mongo.Collection
}

func (repository *repository) GetAll() []Student {
	var students []Student
	return students
}