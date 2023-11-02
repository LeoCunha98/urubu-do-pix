package usecase

import "github.com/LeoCunha98/urubu-do-pix/internal/domain/entity"

type ListUsersOutputDto struct {
	ID   string
	Name string
}

type ListUsersUseCase struct {
	UserRepository entity.UserRepository
}

func NewListUsersUseCase(userRepository entity.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		UserRepository: userRepository,
	}
}

func (u *ListUsersUseCase) Execute() ([]*ListUsersOutputDto, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var usersOutput []*ListUsersOutputDto
	for _, user := range users {
		usersOutput = append(usersOutput, &ListUsersOutputDto{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return usersOutput, nil
}
