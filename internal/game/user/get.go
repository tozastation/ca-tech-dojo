package user

type (
	GetResponse struct {
		Name string `json:"name"`
	}
	GetUsecase interface {
		Execute(token string) (*GetResponse, error)
	}
	GetInteractor struct {
		UserRepo Repository
	}
)

func NewGetInteractor(userRepo Repository) GetUsecase {
	return &GetInteractor{UserRepo:userRepo}
}

func (g GetInteractor) Execute(token string) (*GetResponse, error) {
	user, err := g.UserRepo.Get(token)
	if err != nil {
		return nil, err
	}
	return &GetResponse{Name:user.Name}, nil
}