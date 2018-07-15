GO ?= go

.PHONY: all
all: noteshare

.PHONY: noteshare
noteshare:
	@echo ""
	@echo "==> Compiling and booting noteshare"

	@$(GO) run ./cmd/noteshare/noteshare.go