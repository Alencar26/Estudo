from datetime import date

print("Informe a primeira data:")
dia1 = int(input("Informe o dia: "))
mes1 = int(input("Informe o mês: "))
ano1 = int(input("Informe o ano: "))

print("Informe a segunda data:")
dia2 = int(input("Informe o dia: "))
mes2 = int(input("Informe o mês: "))
ano2 = int(input("Informe o ano: "))

data1 = date(ano1, mes1, dia1)
data2 = date(ano2, mes2, dia2)

diferenca = abs((data2 - data1).days)

print(f"A diferença entre as datas é de {diferenca} dias.")