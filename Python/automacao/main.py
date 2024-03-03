"""
# Passo a passo:
# Ler planilha e guardar informações sobre, nome, telefone e data de vencimento.
# Criar link personalizado do whatsapp e enviar menasgem 
# para o cliente com base nos dados da planilha
"""
import openpyxl
from urllib.parse import quote
import webbrowser
import pyautogui as bot
from time import sleep


workbook = openpyxl.load_workbook('contatos.xlsx')
pagina_clientes = workbook['Sheet1']

for linha in pagina_clientes.iter_rows(min_row=2, values_only=True):
    nome, telefone, data_vencimento = linha
    
    mensagem = f"Olá {nome}, sua fatura vence no dia {data_vencimento.strftime('%d/%m/%Y')}. Faça o pagamento o quanto antes."
    try:
        link_whatsapp = f"https://web.whatsapp.com/send?phone={telefone}&text={quote(mensagem)}"
    
        webbrowser.open(link_whatsapp)
        sleep(10)
        
        btn = bot.locateOnScreen("./btn.png")
        sleep(2)
        bot.click(btn[0], btn[1])
        sleep(5)
        
        bot.hotkey("ctrl", "w")
        sleep(5)
    except Exception as e:
        print(f"Não foi possível enviar mensagem para {nome}")
        print(f"Error:\n===>\n{e}\n<===")
        with open("erros.csv", "a", newline='', encoding="utf-8") as arquivo:
            arquivo.write(f"{nome},{telefone},{data_vencimento}\n")
        pass