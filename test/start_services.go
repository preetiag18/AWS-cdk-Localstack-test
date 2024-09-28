package test

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

type services struct {
    localstack *testcontainers.DockerContainer
}

func startServices(ctx context.Context, t *testing.T) services {
    compose, err := tc.NewDockerCompose("./docker-compose.yml")
    failOnError(t, err, "New docker compose")
    t.Cleanup(func() {
        failOnError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "Docker compose down")
    })
    ctx, cancel := context.WithCancel(ctx)
    t.Cleanup(cancel)
    err = compose.
        Up(ctx, tc.Wait(true))
    failOnError(t, err, "Docker compose up")
    localstack, err := compose.ServiceContainer(ctx, "localstack")
    failOnError(t, err, "Localstack service container")
    return services{
        localstack: localstack,
    }
}
