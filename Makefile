.PHONY: install

install:
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	
	# sqlboiler
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

gen-api:
	oapi-codegen -config api/server.config.yaml api/schema/openapi.yaml
	oapi-codegen -config api/types.config.yaml api/schema/openapi.yaml

gen-model:
	sqlboiler mysql