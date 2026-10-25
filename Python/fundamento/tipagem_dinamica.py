from typing import Final, Union, Iterable, Optional, Any, Callable, TypeAlias, Literal

#definindo constantes
DATABASE: Final = "database.db"
#---

nome: str = "João"
idade: int = 20
altura: float = 1.75
e_maior: bool = idade > 18
hobies: list[str] = ["Jogar", "Futebol"]
trabalhos: tuple[str, str] = ("Programador", "Analista")
atividades: set[str] = {"Jogar", "Futebol"}
pessoa: dict[str, str] = {"nome": "João", "idade": "20"}
#---

#essa variavel aceita mais de um tipo de valor
idade2: Union[int, str] = 20
idade2 = "20"
#ou
cor: str|None = None
cor = "Azul"

#---

#podemos informar que a variável pode receber qlqr tipo ou ser nula
valor: Optional[int] = None
valor = 10
#---

#se vamos receber uma variável do tipo tupla, lista, dicionário, set, etc...
#podemos apenas informar que essa variável pode receber um tipo interável
def lista_hobies(hobies: Iterable) -> None:
    for hobie in hobies:
        print(hobie)
        
#ou informar que esse interável é de strings
def lista_hobies2(hobies: Iterable[str]) -> None:
    for hobie in hobies:
        print(hobie)
#---

#podemos informar que a variável pode ser de qualquer tipo

variavel: Any = "teste"
variavel = 10
variavel = True

def recebe_tudo(valor: Any) -> Any:
    return valor
#---

#podemos definir a tipagem Callable que é um argumento que pode receber qualquer função
import time

def calc_time(func: Callable):
    start = time.time()
    func()
    end = time.time()
    print(f"Tempo de execução: {end - start}s")
    
calc_time(lambda: print("Essa função será executada dentro de outra função"))
#---

#podemos dar apelidos para nossa tipagem para ficar mais legível
ListaEstudos: TypeAlias = dict[int, str]
estudos: ListaEstudos = {1: "Python", 2: "Java", 3: "C++"}

def retorna_minha_lista_estudos(estudos: ListaEstudos) -> ListaEstudos:
    return estudos
#---

weekends: Literal["Sábado", "Domingo"] = "Sábado"

#valores permitidos
weekends = "Domingo"
weekends = "Sábado"

#valores não permitidos
weekends = "Segunda"
weekends = "Terça"

