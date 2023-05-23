-include eng/Makefile

.DEFAULT_GOAL = build
.PHONY: \
	install \
	-install-%

BUILD_VERSION=$(shell git rev-parse --short HEAD)
GO_LDFLAGS=-X 'github.com/Carbonfrost/mechanic/internal/build.Version=$(BUILD_VERSION)'


install: -install-mechanic

-install-%: build -check-env-PREFIX -check-env-_GO_OUTPUT_DIR
	$(Q) eng/install "${_GO_OUTPUT_DIR}/$*" $(PREFIX)/bin
