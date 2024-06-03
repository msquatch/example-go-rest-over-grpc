# For usage information, invoke this makefile without a target.
.DEFAULT_GOAL = help

all: help

PROTOS = echo service-echo

PROGS = echo_service

CUR_DIR = $(shell /bin/pwd)
PROTO_DIR ?= $(CUR_DIR)/proto

PROTO_FILES = $(foreach proto,$(PROTOS),$(proto).proto)
PROG_DEPS = $(foreach prog,$(PROGS),cmd_$(prog))

GO_SVC_PKG_NAME = my

GO_OUT_DIR ?= $(CUR_DIR)/api/go

GATEWAY_OUT_DIR =$ (GO_OUT_DIR)/$(GO_SVC_PKG_NAME)

GO_OUT_FILES = $(foreach proto,$(PROTO),$(proto).pb.go)

BUF_GEN_CONFIG = $(PROTO_DIR)/buf.gen.yaml
OPENAPI_OUT ?= $(CUR_DIR)

.PHONE: run

# Run the echo service.
run:
	./cmd/echo_service/echo_service


cmd_%:
	cd cmd/$* && go build -a

# Build CLI programs.
progs: go $(PROG_DEPS)

.PHONY: go
.PHONY: godir

godir:
	@mkdir -p $(GO_OUT_DIR)/$(GO_SVC_PKG_NAME)

# %.pb.go: %.proto
# 	protoc \
# 		--proto_path=. \
# 		--go_out=$(GO_OUT_DIR) \
# 		--go-grpc_out=$(GO_OUT_DIR) \
# 		$<

# go: godir $(GO_OUT_FILES)

bufgen:
	cd $(PROTO_DIR) && buf generate --template $(BUF_GEN_CONFIG)


# Compile protobuf specs for Go.
go: godir bufgen
	@if [ ! -e "$(GO_OUT_DIR)/$(GO_SVC_PKG_NAME)/go.mod" ]; then \
		cd "$(GO_OUT_DIR)/$(GO_SVC_PKG_NAME)" && \
			go mod init $(GO_SVC_PKG_NAME) && \
			go mod tidy; \
	fi

.PHONY: help
# Print help message.
help:
	@awk 'BEGIN{ds=1;ks=0;ms=25;c="";td=ENVIRON["TMPDIR"];if(td==""){td="/tmp"}sf=sprintf("%s/sort_%d.txt",td,int(1000000*rand()));}function pr(d,tgt,f,v){v=d[tgt];gsub(/\n/,"\n"ind pad,v);printf(f,tgt,v)}/^[[:blank:]]*#[[:blank:]]?/{tc=$$0;sub(/^[[:blank:]]*#[[:blank:]]?/,"",tc);if(c!=""){c=c"\n"}c=c tc;next}/^[[:blank:]]*[^[:blank:]:.=?%][^[:blank:]:=?%]*[[:blank:]]*:/{tgt=$$0;sub(/[[:blank:]]*:.*$$/,"",tgt);sub(/^[[:blank:]]*/,"",tgt);if(c==""){next}if(length(tgt)>ks){ks=length(tgt)}ca[tgt]=c;print tgt>sf;c="";next}//{c=""}END{close(sf);if(ks>ms){ks=ms}ind="    ";f=ind"%-"ks"s  %s\n";pad=sprintf("%"ks"s  "," ");printf("Usage: make <target>\n\n  Targets:\n");if(ds==1){cmd="sort "sf;while((cmd|getline tgt)>0){pr(ca,tgt,f)}close(cmd)}else{while((getline tgt<sf)>0){pr(ca,tgt,f)}close(sf)}system("/bin/rm "sf)}' $(firstword $(MAKEFILE_LIST))
