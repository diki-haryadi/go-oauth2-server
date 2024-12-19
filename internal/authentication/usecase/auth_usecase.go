package authUseCase

import (
	sampleExtServiceDomain "github.com/diki-haryadi/go-micro-template/external/sample_ext_service/domain"
	authDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain"
)

type useCase struct {
	repository              authDomain.Repository
	kafkaProducer           authDomain.KafkaProducer
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
	allowedRoles            []string
}

func NewUseCase(
	repository authDomain.Repository,
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase,
	kafkaProducer authDomain.KafkaProducer,
) authDomain.UseCase {
	return &useCase{
		repository:              repository,
		kafkaProducer:           kafkaProducer,
		sampleExtServiceUseCase: sampleExtServiceUseCase,
		allowedRoles:            []string{Superuser, User},
	}
}
