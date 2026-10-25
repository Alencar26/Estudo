# Cobra CLI

Install: `go install github.com/spf13/cobra-cli@latest`

Para rodar o combra depois de instalado usa-se o comando:

```bash
cobra-cli

#Para iniciar um projeto:
cobra-cli init
```

Temos **três elementos principais** ao trabalhar com cobra:
1. Chamada de Execução
2. Comando
3. Flags configuradas na inicialização - `Feito atravez de uma função Init()`

PS1: Flags só são visíveis no comando onde ela foi criada 
PS2: As PersistentFlags também aparecem em subcomandos ou comandos filhos


### Criar Comandos
 
Para adicionar comando ao nosso CLI usamos o comando abaixo:

```bash
cobra-cli add nomeComando
#exemplo, uma comando de ping

cobra-cli add ping
```

O comando acima vai gerar uma estrutura de código já com o padrão do cobra para nós implementarmos nossa funcionalidade de forma fácil e rápida.
Nassa estrutura dentro de `cmd/nomeComando.go`, é possível adicionar nossa funcionalidade.

### Criar Comando filhos

Para criar comando filhos devemos usar:
```bash
#precisa do Cmd ao final
cobra-cli add nomeComandoFilho -p 'nomecomandopaiCmd'
```
