package pheme

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/erathorus/quickstore"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"cathedral/internal/config"
)

type Cathedral struct {
	Config *config.Cathedral
	Logger *zap.Logger

	AWS   *session.Session
	Store *quickstore.Store
}

func NewCathedral(cfg *config.Cathedral) (*Cathedral, error) {
	if cfg == nil {
		log.Panic("config is nil")
	}

	var logger *zap.Logger
	if cfg.Mode == "production" {
		logger = newProductionLogger()
	} else {
		logger = newDevelopmentLogger()
	}

	cathedral, err := resolveDependencies(cfg, logger)
	if err != nil {
		return nil, err
	}
	return cathedral, nil
}

func (c *Cathedral) Serve() {
	cfg := c.Config
	logger := c.Logger
	router := c.buildRouter()

	logger.Info(fmt.Sprintf("start Cathedral in mode: %s", cfg.Mode))
	logger.Info(fmt.Sprintf("HTTP base path: %s", cfg.HTTP.BasePath))
	logger.Info(fmt.Sprintf("serving at address: %s", cfg.HTTP.Address))
	logger.Panic("serving error", zap.Error(http.ListenAndServe(cfg.HTTP.Address, router)))
}

func newDevelopmentLogger() *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		os.Stderr,
		zap.DebugLevel,
	)
	return zap.New(core, zap.AddStacktrace(zap.WarnLevel), zap.AddCaller())
}

func newProductionLogger() *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.InfoLevel,
	)
	return zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
}

func newAWS(cfg *config.AWS) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func newStore(cfg *config.Store, sess *session.Session) (*quickstore.Store, error) {
	return quickstore.NewStore(dynamodb.New(sess), cfg.DynamoDBTable)
}
