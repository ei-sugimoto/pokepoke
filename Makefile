.PHONY ogen:
ogen:
	@oapi-codegen -package openapi openapi.yml > back/openapi/openapi.gen.go

.PHONY schema:
schema:
	@cd back && go run -mod=mod entgo.io/ent/cmd/ent new $(name)

.PHONY entgen:
entgen:
	@cd back && go generate ./ent