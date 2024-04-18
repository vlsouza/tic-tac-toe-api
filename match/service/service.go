package service

type Service struct {
}

// New ..
func New() Service {
	return Service{}
}

// Create ...
func (s Service) Create() error {
	return nil
}
