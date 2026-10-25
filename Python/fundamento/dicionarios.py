"""
DICIONARIOS: Coleção NÃO ordenada, mutável e que não permite valores duplicados. 
"""

dicionario  = {
    "key1": "value1", 
    "key2": "value2", 
    "key3": 12345, 
    "key4": True
}
#obtendo valores
print(dicionario)
print(dicionario["key2"]) 
print(dicionario.get("key3"))
#obeter todos os valores
print(dicionario.values())
#obtendo as chaves do dicionario
print(dicionario.keys())
#obtem itens com chave e valor
print(dicionario.items())

#---

if "key4" in dicionario:
    print("Minha chave existe:", dicionario["key4"])
    
#---

#alteração dos valores
dicionario["key2"] = False
print(dicionario)

dicionario.update({"key1": "New Value"})
print(dicionario)

#---

#adicionar nova chave
dicionario["newKey"] = hex(45)
print(dicionario)

dicionario.update({"newKey2": "00110010"})
print(dicionario)

#---

#remove
dicionario.pop("key1")
dicionario.popitem() #dependendo da versão do python essa fnução apaga o ultimo valor ou uma valor aleatório
print(dicionario)

del dicionario["key2"]
del dicionario #apaga tudo

#dicionario.clear() #apaga tudo

#---

person = {
    "name": "Andrew",
    "age": 27,
    "country": "Brazil",
    "occupation": "QA"
}

for k in person:
    print(k, '-', person[k])
    
for v in person.values():
    print(v)
    
for k in person.keys():
    print(k)
    
for k, v in person.items():
    print(k, v)


#--- CÓPIA

person2 = person.copy()
print(person2)

person3 = dict(person)
print(person3)

person["alteracao"] = "Esse dicionário foi alterado"
print(person)
print(person2)
print(person3)

#--- CONERTER LISTA/TUPLA ---> DICIONARIO

chaves = ("fruta1", "fruta2", "fruta3")
valores = "Melancia"
sacola_frutas = dict.fromkeys(chaves, valores)
print(sacola_frutas)

#--- OBJETO

escola = {
    "professores": {
        "Prof 5 ano": "João",
        "Prof 4 ano": "Pedro",
        "Prof 6 ano": "Marta"
    },
    "Matérias": {
        "Física": {
            "Professor": "Rosa",
            "Período": 5,
            "Férias": {
                "Inicio": "2024-01-02",
                "Fim": "2024-02-01"
            }
        },
        "Química": None,
        "Matemática": {
            "Professor": "Rosa",
            "Período": 5
        }
    }
}

print(escola)