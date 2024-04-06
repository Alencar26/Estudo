import datetime

dia_hoje = datetime.datetime.today()
print(dia_hoje)
print(dia_hoje.date())

ano = dia_hoje.year
mes = dia_hoje.month
dia = dia_hoje.day

print(dia, mes, ano)

data_futura = datetime.date(2024,4,20)
print('data criada', data_futura)

primeiro_dia_ano = datetime.date(2024,1,1)

print(f"""
      Quantos dias se passram desde o inicio da ano ate hoje?
      R: {dia_hoje.date() - primeiro_dia_ano}
      """)

print(f"Somar mais 30 dias a contar de hoje: {dia_hoje + datetime.timedelta(days=30)}")

timestamp = datetime.datetime.timestamp(datetime.datetime.now())
print("Timestamp:",int(timestamp))
