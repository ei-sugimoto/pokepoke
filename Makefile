.PHONY ogen:
ogen:
	@oapi-codegen -package openapi openapi.yml > back/openapi/openapi.gen.go