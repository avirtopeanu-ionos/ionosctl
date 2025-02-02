## This Makefile contains operations
## for CloudApiDBaaSPgsql resources:
## Tests, Mocks, Documentation
DOCS_OUT_DBAAS_POSTGRES?=$(shell pwd)/docs/subcommands/database-as-a-service/postgres/

.PHONY: dbaas_postgres_test_unit
dbaas_postgres_test_unit:
	@echo "--- Run unit tests for CloudApi DBaaS Postgres ---"
	@go test -cover ./commands/dbaas/postgres/... ./services/dbaas-postgres/...
	@echo "DONE"

.PHONY: dbaas_postgres_test
dbaas_postgres_test: dbaas_postgres_test_unit

.PHONY: dbaas_postgres_mocks_update
dbaas_postgres_mocks_update:
	@echo "--- Update mocks for CloudApi DBaaS Postgres ---"
	@tools/dbaas-postgres/regenerate_mocks.sh
	@echo "DONE"

.PHONY: dbaas_postgres_docs_update
dbaas_postgres_docs_update:
	@echo "--- Generate Markdown documentation for DBaaS Postgres in ${DOCS_OUT_DBAAS_POSTGRES} ---"
	@mkdir -p ${DOCS_OUT_DBAAS_POSTGRES}
	@DOCS_OUT_DBAAS_POSTGRES=${DOCS_OUT_DBAAS_POSTGRES} go run tools/dbaas-postgres/doc.go
	@echo "DONE"
