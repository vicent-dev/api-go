package domain

type UserRepositoryInterface interface {
	Find(id int) (*User, error)
	Save(u *User) error
}

type UserRepository struct{}

func (ur *UserRepository) Find(id int) (*User, error) {
	return nil, nil
}

func (ur *UserRepository) Save(u *User) error {
	return nil
}
