# GO ENV

Uso do comando:

```bash
#comando para ver as variáveis de ambientes que o go utiliza.
go env
```

Retorno: 
```bash
GO111MODULE=''
GOARCH='amd64'
GOBIN=''
GOCACHE='/home/al3ncar/.cache/go-build'
GOENV='/home/al3ncar/.config/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFLAGS=''
GOHOSTARCH='amd64'
GOHOSTOS='linux'
GOINSECURE=''
GOMODCACHE='/home/al3ncar/go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='linux'
GOPATH='/home/al3ncar/go'
GOPRIVATE=''
GOPROXY='https://proxy.golang.org,direct'
GOROOT='/usr/local/go'
GOSUMDB='sum.golang.org'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/usr/local/go/pkg/tool/linux_amd64'
GOVCS=''
GOVERSION='go1.23.6'
GODEBUG=''
GOTELEMETRY='local'
GOTELEMETRYDIR='/home/al3ncar/.config/go/telemetry'
GCCGO='gccgo'
GOAMD64='v1'
AR='ar'
CC='gcc'
CXX='g++'
CGO_ENABLED='1'
GOMOD='/dev/null'
GOWORK=''
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
PKG_CONFIG='pkg-config'
GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build668679208=/tmp/go-build -gno-record-gcc-switches'
```

Dentro do diretório do GOPATH sempre haverá três pastas.

1. bin - Onde ficam todos os binários
2. pkg - Arquivos compilados, pré compilação
3. src - Onde se pode criar seus projetos