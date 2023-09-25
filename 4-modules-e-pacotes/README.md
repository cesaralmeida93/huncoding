# Como funcionam os modules e pacotes no Go

## Estruturas e pastas do GOPATH

/go
  /pkg -> Lugar onde o go guarda todas as dependências já instaladas em seus projetos
  /src -> lugar onde todo o seu código go deve ser guardado
    (ex: /github.com/HunCoding/projeto-test) (ex:/gitlab.com/HunCoding/projeto-test)
  /bin -> lugar onde o go guarda os binários que `go install` compila

## GOPATH vd GOROOT

GOROOT é onde fica guardado o código do Go, literalmente. Código do compilador, das ferramentas e etc, é para lá que o código "go" aponta para executar os comandos do GoLang, assim, ele não é o nosso código fonte(go install, go env, go get apontam para a instalação). Já GOPATH é onde fica seu código, instalações, dependências que você já baixou. o GOROOT geralmente fica em algo como `/usr/local/go`. Já o GOPATH geralmente fica em algo como `$HOME/go`.

## O que são pacotes
Subpastas dentro do seu projeto, que define os pacotes, os pacores são guardados abaixo do diretório `$GOPATH/src`

## O que são modules
Coleção de pacotes Go guardados em um arquivo com o nome e go.mod na raiz do projeto 

go.mod contém:
- Dependências requeriads no projeto
- Caminho do module no repositório

- ex:
**Caminho-do-pacote**   **versão do pacote**
github.com/test/c       v1.3.1
github.com/cde/b        v1.0.0


## Estrutura do go.mod
module <nome>

go <versao-do-go>

require <caminho-do-module> <versao-do-module>
Pacotes que o seu projeto precisa para execuar


replace <caminho-do-module>=> /home/test/test-module
Pacotes que você quer "substituir" por alguma versão que você tenha local daquele mesmo pacote

exclude <caminho-do-module> <versao-do-module>
Pacotes/versões que você não quer que exevutem no seu projeto


```go
module huncoding/minha-primeira-api/validando-dados

go 1.18

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/sys v0.0.0-20220224120231-95c6836cb0e7 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
```

[video do HunCoding](https://www.youtube.com/watch?v=NaoyDqE7WIw&t=946s&ab_channel=HunCoding)