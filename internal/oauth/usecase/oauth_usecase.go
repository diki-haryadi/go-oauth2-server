package oauthUseCase

import (
	sampleExtServiceDomain "github.com/diki-haryadi/go-micro-template/external/sample_ext_service/domain"
	articleDomain "github.com/diki-haryadi/go-micro-template/internal/oauth/domain"
)

type useCase struct {
	repository              articleDomain.Repository
	kafkaProducer           articleDomain.KafkaProducer
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
	allowedRoles            []string
}

func NewUseCase(
	repository articleDomain.Repository,
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase,
	kafkaProducer articleDomain.KafkaProducer,
) articleDomain.UseCase {
	return &useCase{
		repository:              repository,
		kafkaProducer:           kafkaProducer,
		sampleExtServiceUseCase: sampleExtServiceUseCase,
		allowedRoles:            []string{Superuser, User},
	}
}
