package student

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service interface {
	List() (responses []FormatResponse)
	Create(request FormatRequest) (response FormatResponse)
	FindById(ID primitive.ObjectID) (response FormatResponse)
}

type service struct {
	repository Repository
}

func NewService(repository2 Repository) Service {
	return &service{repository: repository2}
}

func (service *service) List() (responses []FormatResponse) {
	students := service.repository.GetAll()
	for _, student := range students {
		responses = append(responses, FormatResponse{
			ID:        student.Id,
			Name:      student.Name,
			Email:     student.Email,
			Gender:    student.Gender,
			IsActive:  student.IsActive,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
		})
	}
	return
}

func (service *service) Create(request FormatRequest) (response FormatResponse) {

	student := Student{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Gender:      request.Gender,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ID := service.repository.Insert(student)

	response = FormatResponse{
		ID:        ID,
		Name:      student.Name,
		Email:     student.Email,
		Gender:    student.Gender,
		IsActive:  student.IsActive,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}

	return
}

func (service *service) FindById(ID primitive.ObjectID) (response FormatResponse) {
	student := service.repository.GetById(ID)

	response = FormatResponse{
		ID:        student.Id,
		Name:      student.Name,
		Email:     student.Email,
		Gender:    student.Gender,
		IsActive:  student.IsActive,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}
	return
}
