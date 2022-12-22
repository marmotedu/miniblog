# ==============================================================================
# 用来进行代码生成的 Makefile
#

.PHONY: gen.add-copyright
gen.add-copyright: ## 添加版权头信息.
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

.PHONY: gen.ca
gen.ca: ## 生成 CA 文件.
	@mkdir -p $(OUTPUT_DIR)/cert
	@openssl genrsa -out $(OUTPUT_DIR)/cert/ca.key 1024 # 生成根证书私钥
	@openssl req -new -key $(OUTPUT_DIR)/cert/ca.key -out $(OUTPUT_DIR)/cert/ca.csr \
		-subj "/C=CN/ST=Guangdong/L=Shenzhen/O=devops/OU=it/CN=127.0.0.1/emailAddress=nosbelm@qq.com" # 2. 生成请求文件
	@openssl x509 -req -in $(OUTPUT_DIR)/cert/ca.csr -signkey $(OUTPUT_DIR)/cert/ca.key -out $(OUTPUT_DIR)/cert/ca.crt # 3. 生成根证书
	@openssl genrsa -out $(OUTPUT_DIR)/cert/server.key 1024 # 4. 生成服务端私钥
	@openssl rsa -in $(OUTPUT_DIR)/cert/server.key -pubout -out $(OUTPUT_DIR)/cert/server.pem # 5. 生成服务端公钥
	@openssl req -new -key $(OUTPUT_DIR)/cert/server.key -out $(OUTPUT_DIR)/cert/server.csr \
		-subj "/C=CN/ST=Guangdong/L=Shenzhen/O=serverdevops/OU=serverit/CN=127.0.0.1/emailAddress=nosbelm@qq.com" # 6. 生成服务端向 CA 申请签名的 CSR
	@openssl x509 -req -CA $(OUTPUT_DIR)/cert/ca.crt -CAkey $(OUTPUT_DIR)/cert/ca.key \
		-CAcreateserial -in $(OUTPUT_DIR)/cert/server.csr -out $(OUTPUT_DIR)/cert/server.crt # 7. 生成服务端带有 CA 签名的证书

.PHONY: gen.protoc
gen.protoc: tools.verify.protoc-gen-go ## 编译 protobuf 文件.
	@echo "===========> Generate protobuf files"
	@protoc                                            \
		--proto_path=$(APIROOT)                          \
		--proto_path=$(ROOT_DIR)/third_party             \
		--go_out=paths=source_relative:$(APIROOT)        \
		--go-grpc_out=paths=source_relative:$(APIROOT)   \
		$(shell find $(APIROOT) -name *.proto)

.PHONY: gen.deps
gen.deps: tools.verify ## 安装依赖，例如：生成需要的代码等.
	@go generate $(ROOT_DIR)/...
