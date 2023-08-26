module github.com/evgsrkn/go-microservices-example/auth

go 1.21.0

require (
	github.com/evgsrkn/go-microservices-example/gateway v0.0.0-20230823214357-d48c18067741
	github.com/evgsrkn/go-microservices-example/user v0.0.0-20230823214357-d48c18067741
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/golang/mock v1.6.0
	github.com/jackc/pgx/v5 v5.4.3
	github.com/lgu-elo/auth v0.0.0-20230426230559-1c22e77e7116
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.25.0
	golang.org/x/crypto v0.12.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/sirupsen/logrus v1.9.2
	github.com/stretchr/testify v1.8.2
	go.uber.org/fx v1.20.0
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
)
