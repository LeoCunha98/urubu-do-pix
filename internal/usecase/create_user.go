package usecase

import "github.com/LeoCunha98/urubu-do-pix/internal/domain/entity"

type CreateUserInputDto struct {
	Name string
}

type CreateUserOutputDto struct {
	ID   string
	Name string
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewCreateUserUseCase(userRepository entity.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (u *CreateUserUseCase) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	user := entity.NewUser(input.Name)
	err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutputDto{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
