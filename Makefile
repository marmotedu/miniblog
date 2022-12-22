# 默认执行 all 目标
.DEFAULT_GOAL := all

# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: gen.add-copyright go.format go.lint go.cover go.build

# ==============================================================================
# Includes

# 确保 `include common.mk` 位于第一行，common.mk 中定义了一些变量，后面的子 makefile 有依赖
include scripts/make-rules/common.mk 
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/generate.mk

# ==============================================================================
# Usage

define USAGE_OPTIONS

Options:
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build/build.multiarch
                   Example: make build BINS="miniblog test"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

## --------------------------------------
## Generate / Manifests
## --------------------------------------

##@ generate:

.PHONY: add-copyright
add-copyright: ## 添加版权头信息.
	@$(MAKE) gen.add-copyright

.PHONY: ca
ca: ## 生成 CA 文件.
	@$(MAKE) gen.ca

.PHONY: protoc
protoc: ## 编译 protobuf 文件.
	@$(MAKE) gen.protoc

.PHONY: deps
deps: ## 安装依赖，例如：生成需要的代码、安装需要的工具等.
	@$(MAKE) gen.deps

## --------------------------------------
## Binaries
## --------------------------------------

##@ build:

.PHONY: build
build: go.tidy  ## 编译源码，依赖 tidy 目标自动添加/移除依赖包.
	@$(MAKE) go.build

## --------------------------------------
## Cleanup
## --------------------------------------

##@ clean:

.PHONY: clean
clean: ## 清理构建产物、临时文件等.
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)


## --------------------------------------
## Lint / Verification
## --------------------------------------

##@ lint and verify:

.PHONY: lint
lint: ## 执行静态代码检查.
	@$(MAKE) go.lint


## --------------------------------------
## Testing
## --------------------------------------

##@ test:

.PHONY: test 
test: ## 执行单元测试.
	@$(MAKE) go.test

.PHONY: cover 
cover: ## 执行单元测试，并校验覆盖率阈值.
	@$(MAKE) go.cover


## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

.PHONY: format
format:  ## 格式化 Go 源码.
	@$(MAKE) go.format

.PHONY: swagger
swagger: tools.verify.swagger ## 启动 swagger 在线文档（监听端口：65534）.
	@swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml

.PHONY: tidy
tidy: ## 自动添加/移除依赖包.
	@$(MAKE) go.tidy

.PHONY: help
help: Makefile ## 打印 Makefile help 信息.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo -e "$$USAGE_OPTIONS"
