module tickets-api

go 1.17

replace github.com/azura-labs/common => ./common

replace github.com/azura-labs/controllers => ./controllers

replace github.com/azura-labs/middlewares => ./middlewares

replace github.com/azura-labs/models => ./models

replace github.com/azura-labs/repositories => ./repositories

replace github.com/azura-labs/databases => ./libs/databases

replace github.com/azura-labs/helper => ./helper

require (
	github.com/azura-labs/common v0.0.0-00010101000000-000000000000
	github.com/azura-labs/controllers v0.0.0-00010101000000-000000000000
	github.com/azura-labs/databases v0.0.0-00010101000000-000000000000
	github.com/azura-labs/middlewares v0.0.0-00010101000000-000000000000
	github.com/labstack/echo v3.3.10+incompatible
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/azura-labs/helper v0.0.0-00010101000000-000000000000 // indirect
	github.com/azura-labs/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/azura-labs/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.6 // indirect
)
