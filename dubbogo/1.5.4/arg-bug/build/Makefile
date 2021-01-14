# 准备环境
PROJECT_DIR=$(shell pwd)
BASE_DIR := $(PROJECT_DIR)/dist

export GO111MODULE ?= on
export GOPROXY ?= https://goproxy.io,direct
export GOSUMDB ?= sum.golang.org
export GOARCH ?= amd64

OUT_DIR := $(BASE_DIR)/$(GOOS)_$(GOARCH)/$(BUILD_TYPE)

## docker-up: Shutdown dependency services on docker
.PHONY: docker-up
docker-up:
	$(info   >  Starting dependency services with $(PROJECT_DIR)/docker/docker-compose.yml)
	@docker-compose -f $(PROJECT_DIR)/docker-compose.yml up -d

## docker-down: Shutdown dependency services on docker
.PHONY: docker-down
docker-down:
	$(info   >  Stopping dependency services with $(PROJECT_DIR)/docker/docker-compose.yml)
	@docker-compose -f $(PROJECT_DIR)/docker-compose.yml down

## run: Run the application (for client)
.PHONY: run
run: export CONF_CONSUMER_FILE_PATH ?= $(OUT_DIR)/conf/client.yml
run: export CONF_ROUTER_FILE_PATH ?= $(OUT_DIR)/conf/router_config.yml
run: build
	$(info   >  Running application $(PROJECT_NAME), output is redirected to $(LOG_FILE))
	@-$(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME) 2>&1 | tee $(LOG_FILE)