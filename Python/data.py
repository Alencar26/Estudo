import datetime

print("=========ATIVIDADE COM DATA E HORAS==========\n")

#entrada do usuário
dia_str = input("Informe o dia: ")
mes_str = input("Informe o mês: ")
ano_str = input("Informe o ano: ")
hora_str = input("Informe a hora: ")
min_str = input("Informe o minuto: ")
seg_str = input("Informe os segundos: ")

#concatenando as variáveis
data_hora_string = dia_str+"/"+mes_str+"/"+ano_str+" "+hora_str+":"+min_str+":"+seg_str

#convertendo a data (string) para data (datetime)
data_hora = datetime.datetime.strptime(data_hora_string, "%d/%m/%Y %H:%M:%S")
print("Você informou a seguinte informação: ",data_hora)

#Menu de opções
alterar = input("Deseja acrescentar algum valor? s(sim)/n(não)")
if alterar.lower() == "s":
    print("""Escolha uma opção:
    1 - Add dia  
    2 - Add mês
    3 - Add ano 
    4 - Add horas 
    5 - Add minuto 
    6 - Add segundo""")
    resp = input("Informe um valor: ")
    valor = int(input("Acrescente um valor: "))
    if resp == "1":
        #adicionando dias na data
        acrescimo_dias = datetime.timedelta(days=valor)
        data_hora = data_hora + acrescimo_dias
        print("Data atualizada:", data_hora)
    elif resp == "2":
        #adicionando mês na data
        acrescimo_mes = datetime.timedelta(months=valor)
        data_hora = data_hora + acrescimo_mes
        print("Data atualizada:", data_hora)
    elif resp == "3":
        #adicionando ano na data
        acrescimo_ano = datetime.timedelta(years=valor)
        data_hora = data_hora + acrescimo_ano
        print("Data atualizada:", data_hora)
    elif resp == "4":
        #adicionando horas na data
        acrescimo_hora = datetime.timedelta(hours=valor)
        data_hora = data_hora + acrescimo_hora
        print("Data atualizada:", data_hora)
    elif resp == "5":
        #adicionando minutos na data
        acrescimo_minuto = datetime.timedelta(minutes=valor)
        data_hora = data_hora + acrescimo_minuto
        print("Data atualizada:", data_hora)
    elif resp == "6":
        #adicionando segundo na data
        acrescimo_segundo = datetime.timedelta(seconds=valor)
        data_hora = data_hora + acrescimo_segundo
        print("Data atualizada:", data_hora)
    else:
        print("Informações inválida!")
else:
    print("Sua data ficou sem alteração!", data_hora)