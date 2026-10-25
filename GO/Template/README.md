# 📌 Lista de Anotações (Sintaxe) para Templates no Go

O pacote `text/template` e `html/template` do Go permitem usar **placeholders** e **estruturas de controle** dentro de templates externos.

Abaixo está uma lista completa das **possibilidades de "anotações"** que podem ser usadas em arquivos `.tmpl` ou `.html`.

---

## ✅ 1. Impressão de Variáveis

Para exibir valores de uma variável no template:

```go
{{ .Nome }}
{{ .Idade }}
{{ .CargaHoraria }}
```

**Exemplo:**

```html
<p>Nome: {{ .Nome }}</p>
<p>Idade: {{ .Idade }}</p>
```

---

## ✅ 2. Condições (`if`, `else`, `eq`, `ne`, `lt`, `gt`, etc.)

### 2.1. `if` e `else`

```go
{{ if .Ativo }}
  <p>O usuário está ativo.</p>
{{ else }}
  <p>O usuário está inativo.</p>
{{ end }}
```

---

### 2.2. Comparação (`eq`, `ne`, `lt`, `gt`, `le`, `ge`)

- `eq` (igual a)
- `ne` (diferente de)
- `lt` (menor que)
- `gt` (maior que)
- `le` (menor ou igual)
- `ge` (maior ou igual)

**Exemplo:**

```go
{{ if eq .CargaHoraria 40 }}
  <p>Carga horária completa.</p>
{{ else }}
  <p>Carga horária parcial.</p>
{{ end }}
```

---

## ✅ 3. Laços de Repetição (`range` e `with`)

### 3.1. `range` (Iteração sobre listas/mapas)

```go
{{ range .Usuarios }}
  <p>Nome: {{ .Nome }}, Idade: {{ .Idade }}</p>
{{ end }}
```

Se `Usuarios` for um **mapa** (exemplo: `map[string]int`), pode-se acessar a **chave e o valor**:

```go
{{ range $key, $value := .Produtos }}
  <p>Produto: {{ $key }}, Preço: {{ $value }}</p>
{{ end }}
```

---

### 3.2. `with` (Mudança de Contexto)

```go
{{ with .Endereco }}
  <p>Rua: {{ .Rua }}</p>
  <p>Cidade: {{ .Cidade }}</p>
{{ end }}
```

🔹 **Sem `with`** → Teria que escrever `{{ .Endereco.Rua }}` e `{{ .Endereco.Cidade }}`.

---

## ✅ 4. Definição e Inclusão de Templates

### 4.1. Definir um Bloco (`define` e `end`)

```go
{{ define "cabecalho" }}
  <h1>Bem-vindo, {{ .Nome }}</h1>
{{ end }}
```

---

### 4.2. Incluir um Template (`template` dentro de outro)

```go
{{ template "cabecalho" . }}
```

---

## ✅ 5. Escape de HTML (`html/template`)

Para evitar **ataques XSS**, usar `html/template`, que escapa HTML automaticamente:

```html
{{ .ConteudoSeguro }} <!-- Escapado -->
{{ .ConteudoNaoSeguro | safeHTML }} <!-- Permitido (se adicionar filtro no Go) -->
```

---

## ✅ 6. Pipe (`|`) para Encadear Funções

```go
{{ .Nome | upper }} <!-- Converte para maiúsculas -->
{{ .Preco | printf "%.2f" }} <!-- Formata número com 2 casas decimais -->
```

---

## ✅ 7. Comentários no Template

```go
{{/* Isso é um comentário e não será exibido no HTML gerado */}}
```

---

## 📌 Resumo das "Anotações" (Sintaxe Template Go)

| **Sintaxe** | **Descrição** |
|------------|--------------|
| `{{ .Variavel }}` | Exibir variável |
| `{{ if .Ativo }}...{{ else }}...{{ end }}` | Estrutura condicional |
| `{{ if eq .Campo "valor" }}` | Comparação (`eq`, `ne`, `lt`, `gt`, `le`, `ge`) |
| `{{ range .Lista }}` | Iteração sobre listas/mapas |
| `{{ with .Objeto }}` | Muda o contexto temporariamente |
| `{{ define "nome" }}...{{ end }}` | Definição de template |
| `{{ template "nome" . }}` | Inclusão de template |
| `{{ .Valor \| printf "%.2f" }}` | Encadeamento com `\|` (pipes) |
| `{{/* Comentário */}}` | Comentário |

---

# 📌 Funções em Templates do Go

No Go, podemos usar **funções** dentro de templates para manipular e formatar os dados exibidos. Algumas funções são **embutidas** pelo pacote `text/template`, enquanto outras precisam ser **registradas manualmente** no código Go.

---

## ✅ 1. Funções Embutidas Disponíveis no Go

O pacote `text/template` oferece algumas funções básicas sem necessidade de configuração extra.

| **Função**  | **Descrição** | **Exemplo** |
|------------|-------------|------------|
| `printf`   | Formatação de string semelhante ao `fmt.Sprintf` | `{{ printf "%.2f" .Preco }}` → `19.99` |
| `println`  | Adiciona uma nova linha após o texto | `{{ println "Olá" .Nome }}` |
| `len`      | Retorna o tamanho de um array, slice, string ou mapa | `{{ len .Nomes }}` |
| `index`    | Acessa elementos de um slice ou mapa | `{{ index .Lista 0 }}` |
| `eq` `ne` `lt` `gt` `le` `ge` | Comparações (`==`, `!=`, `<`, `>`, `<=`, `>=`) | `{{ if eq .Idade 18 }} Maior de idade {{ end }}` |

---

## ✅ 2. Funções Personalizadas no Go

O Go não fornece funções como `upper` e `lower` por padrão. Para usá-las, é necessário **registrá-las manualmente** no template.

### 🔹 Funções Comuns para Registro

| **Nome**  | **Descrição** | **Implementação** |
|-----------|-------------|------------------|
| `upper`   | Converte string para maiúsculas | `strings.ToUpper(s)` |
| `lower`   | Converte string para minúsculas | `strings.ToLower(s)` |
| `title`   | Capitaliza cada palavra | `strings.Title(s)` |
| `trim`    | Remove espaços em branco | `strings.TrimSpace(s)` |
| `replace` | Substitui partes da string | `strings.Replace(s, "antigo", "novo", -1)` |
| `split`   | Divide uma string em partes | `strings.Split(s, ",")` |
| `join`    | Junta um slice de strings | `strings.Join(slice, ",")` |
| `contains`| Verifica se uma string contém outra | `strings.Contains(s, "termo")` |
| `hasPrefix` | Verifica se a string começa com um prefixo | `strings.HasPrefix(s, "Go")` |
| `hasSuffix` | Verifica se a string termina com um sufixo | `strings.HasSuffix(s, ".txt")` |

---

### 🔹 Como Registrar Funções Personalizadas no Go

```go
package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	funcs := template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"trim": strings.TrimSpace,
	}

	tmpl := `Nome: {{ .Nome | upper }}, Cidade: {{ .Cidade | lower }}`
	t, _ := template.New("exemplo").Funcs(funcs).Parse(tmpl)

	data := map[string]string{"Nome": "André", "Cidade": "São Paulo"}
	t.Execute(os.Stdout, data)
}
```

**Saída:**
```plaintext
Nome: ANDRÉ, Cidade: são paulo
```

---

## ✅ 3. Aplicação das Funções em Templates

Essas funções podem ser aplicadas a **qualquer campo do template**, desde que os dados tenham o tipo apropriado. Alguns exemplos:

### 3.1. **Manipulação de Strings**
```go
{{ .Nome | upper }}  <!-- ANDRÉ -->
{{ .Texto | trim }}  <!-- Remove espaços extras -->
{{ .Descricao | replace "antigo" "novo" }}  <!-- Substitui palavras -->
```

### 3.2. **Formatação de Números**
```go
{{ printf "%.2f" .Preco }}  <!-- 19.99 -->
{{ printf "%d" .Quantidade }}  <!-- Exibe um número inteiro -->
```

### 3.3. **Manipulação de Listas e Mapas**
```go
{{ len .Produtos }}  <!-- Conta itens na lista -->
{{ index .Produtos 0 }}  <!-- Pega o primeiro item -->
{{ join .Tags ", " }}  <!-- Junta um slice de strings -->
```

---

## 📌 Resumo das Funções em Templates do Go

| **Função** | **Descrição** |
|-----------|--------------|
| `printf` | Formata texto como `fmt.Sprintf` |
| `len` | Retorna o tamanho de listas/mapas |
| `index` | Acessa elementos de um slice/mapa |
| `eq`, `ne`, `lt`, `gt`, `le`, `ge` | Comparações |
| `upper`, `lower` | Converte string para maiúsculas/minúsculas (personalizado) |
| `trim` | Remove espaços extras de uma string (personalizado) |
| `replace` | Substitui partes de uma string (personalizado) |
| `split`, `join` | Divide e junta strings (personalizado) |
| `contains`, `hasPrefix`, `hasSuffix` | Verificações de substring (personalizado) |
