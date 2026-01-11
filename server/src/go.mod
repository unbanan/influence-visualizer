module contest-influence/server

go 1.25.5

replace contest-influence/proto => ../../proto/go

require (
	contest-influence/proto v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/samber/lo v1.52.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
