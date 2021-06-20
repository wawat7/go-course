package student

type Service interface {
	List() (responses []FormatResponse)
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
