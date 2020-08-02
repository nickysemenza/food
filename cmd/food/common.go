package main

import (
	"fmt"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() error {
	endpoint := viper.GetString("JAEGER_ENDPOINT")
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	// Create and install Jaeger export pipeline
	if projectID != "" {
		// google cloud trace
		// env: GOOGLE_CLOUD_PROJECT=xx GOOGLE_APPLICATION_CREDENTIALS=x.json
		exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))
		if err != nil {
			return fmt.Errorf("texporter.NewExporter: %w", err)
		}

		tp, err := sdktrace.NewProvider(sdktrace.WithSyncer(exporter))
		if err != nil {
			return err
		}
		global.SetTraceProvider(tp)
	} else if endpoint != "" {
		_, _, err := jaeger.NewExportPipeline(
			jaeger.WithCollectorEndpoint(endpoint),
			jaeger.WithProcess(jaeger.Process{
				ServiceName: "food",
				Tags: []core.KeyValue{
					key.String("exporter", "jaeger"),
				},
			}),
			jaeger.RegisterAsGlobal(),
			jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		)
		return err
	}

	return nil
}

func setupEnv() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5555)
	viper.SetDefault("DB_USER", "food")
	viper.SetDefault("DB_PASSWORD", "food")
	viper.SetDefault("DB_DBNAME", "food")
	viper.SetDefault("DB_MAX_OPEN_CONNS", 20)
	viper.SetDefault("PORT", 4242)
	viper.SetDefault("HTTP_TIMEOUT", "30s")
	viper.SetDefault("SENTRY_DSN", "https://8220ab8a2b3d4c3c9cf7f636ec183c7a@o83311.ingest.sentry.io/5298706")

	viper.SetDefault("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")
}

func setupMisc() {
	// env vars
	setupEnv()
	viper.AutomaticEnv()

	// tracing
	if err := initTracer(); err != nil {
		log.Fatal(err)
	}
	log.Infof("tracer initialized")

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("SENTRY_DSN"),
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	log.Infof("sentry initialized")
}
