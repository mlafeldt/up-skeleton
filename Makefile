ENV = staging

deploy: up.json
	@up deploy $(ENV) -v
	@up url $(ENV)

stage: ENV=staging
stage: deploy

prod: ENV=production
prod: deploy

destroy: up.json
	@up stack delete --async --force

local: up.json
	@$(shell jq -r '.environment | to_entries | .[] | .key + "=" + .value' up.json) go run *.go

up.json: up/config.json up/local.json
	@jq -s '.[0] * .[1]' $^ > $@

up/local.json:
	@echo '{"environment":{}}' > $@
