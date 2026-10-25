from random import choice

jogador_vitorias = 0
maquina_vitorias = 0

def Opcao_Jogador():
    opcao_jogador = input("Escolha uma opção: Pedra, Papel ou Tesoura: ")
    opcao_jogador.lower()
    if opcao_jogador in ["pedra", "papel", "tesoura"]:
        return opcao_jogador
    else:
        print("Opção inválida")
        return Opcao_Jogador()

def Opcao_Maquina():
    opcao_maquina = choice(["pedra", "papel", "tesoura"])
    return opcao_maquina

while True:
    
    print("-"*30)
    esc_jogador = Opcao_Jogador()
    esc_maquina = Opcao_Maquina()
    print("-"*30)
    
    if (esc_jogador == "pedra") and (esc_maquina == "tesoura") \
        or (esc_jogador == "papel") and (esc_maquina == "pedra") \
        or (esc_jogador == "tesoura") and (esc_maquina == "papel"):
        jogador_vitorias += 1
        print(f"Você ganhou! - (você) [{esc_jogador}] ---> [{esc_maquina}] (máquina)")
    elif esc_jogador == esc_maquina:
        print(f"Empate! - (você) [{esc_jogador}] <--> [{esc_maquina}] (máquina)")
    else:
        maquina_vitorias += 1
        print(f"Você perdeu! - (você) [{esc_jogador}] <--- [{esc_maquina}] (máquina)")
    
    print("-"*30)
    print(f"Vitórias do jogador: {jogador_vitorias}")
    print(f"Vitórias da máquina: {maquina_vitorias}")
    print("-"*30)
    
    esc_jogador = input("Você quer jogar novamente? ")
    if esc_jogador in ["SIM", "sim", "S", "s"]:
        continue
    elif esc_jogador in ["NÃO", "não", "N", "n"]:
        break
    else:
        break