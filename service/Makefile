export GOOS := windows
export GOARCH := amd64
export PATH := $(PATH):$(GOPATH)/bin

all: wireguard-service.exe

BUILDDIR := .tmp
rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
DOWNSTREAM_FILES := $(filter-out %/.tmp wireguard-go/%,$(call rwildcard,,*.go))
UPSTREAM_FILES := $(filter-out $(addprefix %/,$(DOWNSTREAM_FILES)) %/.tmp %/main_windows.go,$(call rwildcard,wireguard-go/,*.go)) wireguard-go/go.mod wireguard-go/go.sum

define copy-src-to-build
$(BUILDDIR)/$(patsubst $(1)%,%,$(2)): $(2)
	@mkdir -vp "$$(dir $$@)"
	@cp -vp "$$<" "$$@"
$(BUILDDIR)/.prepared: $(BUILDDIR)/$(patsubst $(1)%,%,$(2))
endef

$(foreach FILE,$(UPSTREAM_FILES),$(eval $(call copy-src-to-build,wireguard-go/,$(FILE))))
$(foreach FILE,$(DOWNSTREAM_FILES),$(eval $(call copy-src-to-build,,$(FILE))))

$(BUILDDIR)/.prepared:
	cd "$(BUILDDIR)" && go get
	touch "$@"

wireguard-service.exe: $(BUILDDIR)/.prepared $(BUILDDIR)/resources.syso
	cd $(BUILDDIR) && go build -ldflags="-H windowsgui" -o ../$@ -v

$(BUILDDIR)/resources.syso: ../manager/icon/icon.ico
	rsrc -ico ../manager/icon/icon.ico -arch amd64 -o $@

run: wireguard-service.exe
	wine wireguard-service.exe

clean:
	rm -rf wireguard-service.exe $(BUILDDIR)

.PHONY: run clean
