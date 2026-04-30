package class

import "gorm.io/gorm"

type ClassService interface {
	CreateClass(classData *InputCreateClass) (*Class, error)
	GetClasses() ([]Class, error)
	GetClass(id string) (*Class, error)
	UpdateClass(id string, classData *InputUpdateClass) (*Class, error)
	DeleteClass(id string) error
}

type classService struct {
	db *gorm.DB
}

func NewClassService(db *gorm.DB) ClassService {
	return &classService{db: db}
}

func (cs *classService) CreateClass(classData *InputCreateClass) (*Class, error) {
	newClass := &Class{Title: classData.Title, Category: classData.Category, Price: classData.Price, SessionId: classData.SessionId}
	result := cs.db.Create(newClass)
	if result.Error != nil {
		return nil, result.Error
	}
	return newClass, nil
}

func (cs *classService) GetClasses() ([]Class, error) {
	var classes []Class

	result := cs.db.Find(&classes)
	if result.Error != nil {
		return nil, result.Error
	}
	return classes, nil
}

func (cs *classService) GetClass(id string) (*Class, error) {
	singleClass := NewClass()
	result := cs.db.First(&singleClass, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return singleClass, nil
}

func (cs *classService) UpdateClass(id string, classData *InputUpdateClass) (*Class, error) {
	updateClass := NewClass()
	result := cs.db.First(updateClass, id)
	if result.Error != nil {
		return nil, result.Error
	}
	cs.db.Model(&updateClass).Updates(Class{Title: updateClass.Title, Category: updateClass.Category, Price: updateClass.Price, SessionId: updateClass.SessionId})
	return updateClass, nil
}

func (cs *classService) DeleteClass(id string) error {
	deleteClass := NewClass()
	result := cs.db.First(deleteClass, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
