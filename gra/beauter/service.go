package beauter

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service *Service) SayHello(name string) string {
	return "Hello " + name
}
