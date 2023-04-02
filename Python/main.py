import requests

cep = input("Digite seu CEP: ")

url = f"https://viacep.com.br/ws/{cep}/json/"

response = requests.get(url)

if response.status_code == 200:
    data = response.json()
    print(data)
else:
    print("Erro ao consultar CEP")
