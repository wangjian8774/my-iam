.DEFAULT_GOLA := all

.PHONY: all
all: tidy gen add-copyright format lint cover build

include scripts/make-rules/common.mk

define USAGE_OPTIONS

Options:
	DEBUG				where to generate debug symbol.
	BINS				The binarires to build. Default is all of cmd.
	IMAGES				Backend images to make.
	REGISTRY_PREFIX		Docker registry prefix.
	PLATFORMS			The multiple platforms to build. Default is linux_amd74 and linux_arm64.
	VERSION				The version information compiled into binaries.
						The default is obtianed from gsmver or git.
	V					Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

.PHONY: help
help: Makefile
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"