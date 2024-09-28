package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"
)

type App struct {
}

type GetUserInput struct {
	Name string `json:"name"`
}

func (app *App) handle(_ context.Context, event GetUserInput) (interface{}, error) {
	log.Info().
		Str("Name", event.Name).
		Msg("Get user response for name ...")

	return "ok", nil
}

func main() {
	log.Info().Msg("Initializing get user lambda")
	app := App{}

	lambda.Start(app.handle)
}
