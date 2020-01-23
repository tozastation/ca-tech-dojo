package user

type (
	UpdateRequest struct {
		Name string `json:"name"`
	}
	UpdateUsecase interface {
		Execute(req UpdateRequest, token string) error
	}
	UpdateInteractor struct {
		UserRepo Repository
	}
)

func NewUpdateInteractor(userRepo Repository) UpdateUsecase {
	return &UpdateInteractor{UserRepo:userRepo}
}

func (u UpdateInteractor) Execute(req UpdateRequest, token string)  error {
	user, err := u.UserRepo.Get(token)
	if err != nil {
		return err
	}
	user.Name = req.Name
	if err := u.UserRepo.Update(*user); err != nil {
		return err
	}
	return nil
}