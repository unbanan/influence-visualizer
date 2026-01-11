module contest-influence/server/tests/unit

go 1.25.5

replace contest-influence/server => ../../src

replace contest-influence/proto => ../../../proto/go

require (
	contest-influence/proto v0.0.0-00010101000000-000000000000
	contest-influence/server v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/samber/lo v1.52.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
