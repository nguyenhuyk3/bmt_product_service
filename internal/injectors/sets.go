package injectors

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/internal/implementations/message_broker/writer"
	"bmt_product_service/internal/implementations/redis"
	"bmt_product_service/internal/injectors/provider"

	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	provider.ProvidePgxPool,
	writer.NewKafkaWriter,
	sqlc.NewStore,
)

var redisSet = wire.NewSet(
	redis.NewRedisClient,
)
