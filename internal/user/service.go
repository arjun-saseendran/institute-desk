package user

import "gorm.io/gorm"

type UserService interface {
	Create(userData *InputCreateUser) (*User, error)
	GetUsers() ([]User, error)
	GetUser(id string) (*User, error)
	UpdateUser(id string, userData *InputUpdateUser) (*User, error)
}

type userService struct{ db *gorm.DB }

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (us *userService) Create(userData *InputCreateUser) (*User, error) {
	newUser := &User{Name: userData.Name, Address: userData.Address, Mobile: userData.Mobile}
	result := us.db.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return newUser, nil
}

func (us *userService) GetUsers() ([]User, error) {
	var users []User
	result := us.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (us *userService) GetUser(id string) (*User, error) {
	singleUser := NewUser()
	result := us.db.First(&singleUser, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return singleUser, nil

}

func (us *userService) UpdateUser(id string, userData *InputUpdateUser) (*User, error) {
	updateUser := NewUser()
	result := us.db.First(updateUser, id)
	if result.Error != nil {
		return nil, result.Error
	}
	us.db.Model(&updateUser).Updates(User{Name: userData.Name, Address: userData.Address, Mobile: userData.Mobile})
	return updateUser, nil
}
