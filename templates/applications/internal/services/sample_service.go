package services

type SampleService struct {
}

func NewSampleService() *SampleService {
	return &SampleService{}
}

func (s *SampleService) SampleMethod() string {
	return "Sample"
}
