#dado uma matriz de m x n chamada board e uma string chamada
# word, retorne True se a Word estiver dentro da matriz.
# A palavra é construída por celulas sequencialmente adjacentes
# Entenda como celulas vizinhas horintais e verticais, cada celula só
# pode ser usada uma vez.

from typing import List, TypeAlias, Tuple

Matriz2D: TypeAlias = List[List[str]]
cordenadas: TypeAlias = List[Tuple[int, int]]

movimentos: cordenadas = [
    (1,0),
    (0, 1),
    (-1, 0),
    (0, -1)
]

board: Matriz2D = [
        ["A", "B", "C", "E"],
        ["S", "F", "C", "S"],
        ["A", "D", "E", "E"]
        ]

word: str = "ABCCED"

def verificar_palavra(board:Matriz2D, word:str):
    for i, linha in enumerate(board):
        for j, letra in enumerate(linha):
            if letra != word[0]:
                continue
            if andar(i,j, board, word, 0):
                print("\nPALAVRA ENCONTRADA!\n")
                print_board(board)
    return False
        

def andar(x:int, y:int, board:Matriz2D, word:str, idx:int) -> bool:
    
    #verifica tamanho para ver se achou a palavra
    if idx >= len(word):
        return True
    
    #verifica se o movimento é valido
    if x < 0 or x >= len(board) or y < 0 or y >= len(board[0]):
        return False
    
    #verifica se a palavra é igual a celula
    celula = board[x][y]
    if celula != word[idx]:
        return False
    
    #marca a celula como usada
    board[x][y] = "*"
    
    #faz o andar
    for movimento in movimentos:
        if andar(x+movimento[0], y+movimento[1], board, word, idx + 1):
            return True
    
    #desmarca a celula
    board[x][y] = celula
    
    return False

def print_board(board:Matriz2D):
    for linha in board:
        print(linha)
    print("\n")

verificar_palavra(board, word)
