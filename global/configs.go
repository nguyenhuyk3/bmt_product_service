package global

import (
	"bmt_product_service/pkgs/settings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Config     settings.Config
	Postgresql *pgxpool.Pool
)
