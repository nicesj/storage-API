include envfile
export $(shell sed 's/=.*//' envfile)

all: getmodule

validate: ${SPEC_FILE}
	@swagger validate ${SPEC_FILE}

go.mod:
	@go mod init

generate: ${SPEC_FILE} go.mod
	@swagger generate server -f ${SPEC_FILE} -t ${BASE} --api-package=operations --model-package=models --server-package=restapi

getmodule: generate
	@go get -u -f ./${BASE}/...

doc: validate
	@swagger serve ${SPEC_FILE} -p ${DOC_PORT} --no-open

run:
	@go run internal/cmd-server/main.go #--tls-port=${TLS_PORT} --tls-key=${TLS_KEY} --tls-certificate=${TLS_CERTIFICATE}

clean:
	@rm -rf go.mod
	@rm -rf go.sum
	@rm -rf internal/cmd/
	@rm -rf internal/pkg/
