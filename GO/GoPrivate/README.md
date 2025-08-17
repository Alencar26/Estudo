# Go Private

Como trabalhar com repositórios privados no GO?

Primeiramente deve-se adicionar a URL do repositório privado nas suas variáveis de ambiente.

Para verificar suas variáveis do GO, use o comando abaixo:

```bash
go env
```

Para localizar a variável de repositórios privados use o comando:

```bash
go env | grep GOPRIVATE
```

Para adicionar um repositório privado basta informar o endereço dele nessa variável e separar por `,`.

```bash
export GOPRIVATE=github.com/User/Repositorio/Exemplo,github.com/User2/OutroRepo/Exemplo
```

Só isso não basta, agora precisamos de autenticação para ter acesso à esse repositório.
Para isso vamos deve ser editado um arquivo chamado `.netrc` que fica na raiz do usuário.

```bash
vim ~/.netrc
```
 
Nele deve conter o seu login + token de autenticação do github.
Exemplo:

```.netrc
machine github.com
login meu_login_aqui
password token_do_github_aqui
```

Com isso, já é possível efeturar sua autenticação via HTTP.
É possível fazer autenticaão via SSH também, editando o arquivo .git/config.
