package v1

type Servicer interface {
	HelloService() string
}

type Service struct{}

func NewHelloService() *Service {
	return &Service{}
}

func (s Service) HelloService() string {
	return "hello"
}
