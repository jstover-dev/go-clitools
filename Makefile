# Name of the main entrypoint
MAIN   := clitools

BUILD  := build
BIN    := $(BUILD)/bin

# Application names (all except main entrypoint)
APPS         := $(notdir $(wildcard pkg/*))

# Application binaries go to build/bin
APP_TARGETS  := $(foreach x,$(APPS) $(MAIN),$(BIN)/$x)

# Application links link from build -> bin/main
LINK_TARGETS := $(foreach x,$(APPS),$(BUILD)/$x)

# Recursive Wildcard
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

PKG_SOURCES  := $(call rwildcard,pkg,*.go)

MAIN_SOURCES := $(wildcard cmd/$(MAIN)/*) $(PKG_SOURCES)


.PHONY: all clean

all: $(APP_TARGETS) $(LINK_TARGETS)

clean:
	rm -rf $(BIN) $(BUILD)

$(BIN)/$(MAIN): $(MAIN_SOURCES)
	mkdir -p $(BIN)
	cd $(BIN) && go build ../../cmd/$(MAIN)

.SECONDEXPANSION:
$(BIN)/% : cmd/%/* pkg/%/*
	@mkdir -p $(BIN)
	cd $(BIN) && go build ../../cmd/$*

$(BUILD)/% : $(BIN)/$(MAIN)
	@mkdir -p $(BUILD)
	cd $(BUILD) && ln -sfv bin/$(MAIN) $*

$(APPS):
	@$(MAKE) --no-print-directory $(BIN)/$@
	@$(MAKE) --no-print-directory $(BUILD)/$@
