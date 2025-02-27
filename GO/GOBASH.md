# GO Comandos

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

---

### Build Código GO

> Link útil: [Building Go Applications for Different Operating Systems and Architectures](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

GO não faz uso de uma virtual machine (ex: JVM igual java), o Go pega o *Go Runtime* que é o código necessário para funcionar a linguagem e junta ele com o meu código da minha aplicação em um único binário;

`GO Runtime` + `Meu código` = `Binário`

Esse binário é o executável do meu software.

É possíve, escolher para qual SO queremos o binário, seja linux, windows ou MacOS.

Com o comando `go env | grep GOOS` podemos ver qual é o SO default configurado em nossas variáveis do GO.

Podemos ainda escolher a arquitetura se será `arm64` ou `amd64` por exemplo. 

Com o comando `go env | grep GOARCH` é possível ver o valor default.

```bash
#ver qual SO padrão para gerar o binário
go env | grep GOOS

#ver qual arquitetura padrão de build
go env | grep GOARCH

#gerando binário
go build arquivo.go

#para escolher o nome do binário pode usar o parametro -o
go build -o nome_do_binario arquivo.go

#se estiver dentro de uma módulo
go build -o nome_do_binario

#gerando binário para SO específico
GOOS=windows go build arquivo.go

#gerando binário para SO e Arquitetura específica
GOOS=windows GOARCH=amd64 go build arquivo.go
GOOS=windows GOARCH=arm64 go build arquivo.go
#ou
GOOS=linux GOARCH=amd64 go build arquivo.go
GOOS=linux GOARCH=arm go build arquivo.go
GOOS=linux GOARCH=mips go build arquivo.go
```

```bash
#Para ver a lista de combinações possíveis use o comando:
go tool dist list
```

```TXT
COMBINAÇÕES POSSÍVEIS:

aix/ppc64        freebsd/amd64   linux/mipsle   openbsd/386
android/386      freebsd/arm     linux/ppc64    openbsd/amd64
android/amd64    illumos/amd64   linux/ppc64le  openbsd/arm
android/arm      js/wasm         linux/s390x    openbsd/arm64
android/arm64    linux/386       nacl/386       plan9/386
darwin/386       linux/amd64     nacl/amd64p32  plan9/amd64
darwin/amd64     linux/arm       nacl/arm       plan9/arm
darwin/arm       linux/arm64     netbsd/386     solaris/amd64
darwin/arm64     linux/mips      netbsd/amd64   windows/386
dragonfly/amd64  linux/mips64    netbsd/arm     windows/amd64
freebsd/386      linux/mips64le  netbsd/arm64   windows/arm
```

> OBS: Se o meu projeto estiver dentro de um módulo do GO `go.mod`, não há necessidade de passar o arquivo para compilar, basta executar `go build` e automaticamnte ele vai no módulo e pega todas as informações necessário para compilar. O arquivo gerado estará com o nome do módulo.

