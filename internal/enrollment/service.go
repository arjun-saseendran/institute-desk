package enrollment

import "gorm.io/gorm"

type EnrollmentService interface {
	CreateEnrollment(enrolData *InputCreateEnrollment) (*Enrollment, error)
	GetEnrollments() ([]Enrollment, error)
	GetEnrollment(id string) (*Enrollment, error)
	DeleteEnrollment(id string) error
}

type enrollmentService struct {
	db *gorm.DB
}

func NewEnrollmentService(db *gorm.DB) EnrollmentService {
	return &enrollmentService{db}
}

func (es *enrollmentService) CreateEnrollment(enrolData *InputCreateEnrollment) (*Enrollment, error) {
	newEnrollment := &Enrollment{UserId: enrolData.UserId, ClassId: enrolData.ClassId, SessionId: enrolData.SessionId}
	result := es.db.Create(newEnrollment)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEnrollment, nil
}

func (es *enrollmentService) GetEnrollments() ([]Enrollment, error) {
	var enrollments []Enrollment
	result := es.db.Find(&enrollments)
	if result.Error != nil {
		return nil, result.Error
	}
	return enrollments, nil
}

func (es *enrollmentService) GetEnrollment(id string) (*Enrollment, error) {
	singleEnrollment := NewEnrollment()
	result := es.db.First(&singleEnrollment, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return singleEnrollment, nil
}

func (es *enrollmentService) DeleteEnrollment(id string) error {
	deleteEnrolllment := NewEnrollment()
	result := es.db.First(deleteEnrolllment, id)
	if result.Error != nil {
		return result.Error
	}
	es.db.Delete(deleteEnrolllment)
	return nil

}
