// +build wireinject

package pheme

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"cathedral/shared/config"
)

func resolveDependencies(cfg *config.Cathedral, logger *zap.Logger) (*Cathedral, error) {
	wire.Build(
		wire.Struct(new(Cathedral), "*"),
		wire.FieldsOf(
			new(*config.Cathedral),
			"AWS",
			"Store",
		),
		newAWS,
		newStore,
	)
	return new(Cathedral), nil
}
