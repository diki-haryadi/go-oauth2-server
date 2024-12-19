package authConfigurator

import (
	"context"
	sampleExtServiceUseCase "github.com/diki-haryadi/go-micro-template/external/sample_ext_service/usecase"
	authGrpcController "github.com/diki-haryadi/go-micro-template/internal/authentication/delivery/grpc"
	authHttpController "github.com/diki-haryadi/go-micro-template/internal/authentication/delivery/http"
	authKafkaProducer "github.com/diki-haryadi/go-micro-template/internal/authentication/delivery/kafka/producer"
	authDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain"
	authRepository "github.com/diki-haryadi/go-micro-template/internal/authentication/repository"
	authUseCase "github.com/diki-haryadi/go-micro-template/internal/authentication/usecase"
	authenticationV1 "github.com/diki-haryadi/protobuf-ecomerce/oauth2_server_service/authentication/v1"
	externalBridge "github.com/diki-haryadi/ztools/external_bridge"
	infraContainer "github.com/diki-haryadi/ztools/infra_container"
)

type configurator struct {
	ic        *infraContainer.IContainer
	extBridge *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, extBridge *externalBridge.ExternalBridge) authDomain.Configurator {
	return &configurator{ic: ic, extBridge: extBridge}
}

func (c *configurator) Configure(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.extBridge.SampleExtGrpcService)
	kafkaProducer := authKafkaProducer.NewProducer(c.ic.KafkaWriter)
	repository := authRepository.NewRepository(c.ic.Postgres)
	useCase := authUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// grpc
	grpcController := authGrpcController.NewController(useCase)
	authenticationV1.RegisterAuthServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := authHttpController.NewController(useCase)
	authHttpController.NewRouter(httpController).Register(httpRouterGp)

	// consumers
	//oauthKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// jobs
	//oauthJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}
