# TMUX - Cheat Sheet & Quick Reference

[Referências](https://tmuxcheatsheet.com/)

### Comandos terminal

```bash
#iniciar app
tmux

#listar as várias instâncias do tmux
tmux ls


```

### Comandos funcionais dentro do Tmux

```tmux
#Navegação com scroll e pesquisa
C-b + [
#Para sair do modo de naveção use 'q'

#Ainda no mode navegação para pesquisar algo em tela, use:
/ + termo de pesquisa + ENTER

#Para navegar em cada termo use 'n'

#split terminal (vertical)
C-b + %

#split terminal (horizontal)
C-b + "

#Navegar entras as janelas
C-b + (setas) <- ->

#Navegar entre as instâncias de forma visual
#Session and Window Preview
C-b + w

#Renomear sessão
C-b + $

#ajustar tamanho das janelas
C-b + (ctrl segurado) + setas <- ->

#mostrar todas as sessões
C-b + s

#navegar entras as sessões
C-b + ( (anterior)
C-b + ) (próx)

#criar nova janela
C-b + c

#mover janelas de posição
C-b + {
C-b + }

#trocar de janela para a próxima
C-b + n

#zoom
C-b + z

#separar split para uma janela separada
C-b + !

#fechar janela atual
C-b + x


```
