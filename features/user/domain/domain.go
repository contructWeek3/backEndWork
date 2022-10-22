package domain

type UserCore struct {
	ID       uint
	Username string
	Email    string
	Name     string
	Phone    string
	Address  string
	Images   string
	Password string
}

type Repository interface {
	GetMyUser(userID uint) (UserCore, error)
	Update(updatedUser UserCore, userID uint) (UserCore, error)
	Delete(userID uint) error
	AddUser(newUser UserCore) (UserCore, error)
	GetUser(existUser UserCore) (UserCore, error)
}

type Service interface {
	MyProfile(userID uint) (UserCore, error)
	UpdateProfile(updatedUser UserCore, userID uint) (UserCore, error)
	Deactivate(userID uint) error
	Register(newUser UserCore) (UserCore, error)
	Login(existUser UserCore) (UserCore, error)
}
