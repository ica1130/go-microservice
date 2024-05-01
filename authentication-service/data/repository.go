package data

// Repository is the interface that defines the methods that must be implemented by a
// repository. This is used to allow us to switch between different storage solutions
// without changing the code that uses the repository.
type Repository interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Update(user User) error
	DeleteByID(id int) error
	Insert(user User) (int, error)
	ResetPassword(password string, user User) error
	PasswordMatches(plainText string, user User) (bool, error)
}
