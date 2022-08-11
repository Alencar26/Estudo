local map = require("map")

print("Level 1")
print(map.load("map_level1.txt"))

print("Level 2")
print(map.load("map_level2.txt"))

--
-- CARREGANDO OS MAPAS EM UMA TABELA 
--

print("Mapas em tabelas.")
print(map.loadInArray("map_level1.txt"))
print(map.loadInArray("map_level2.txt"))

print("Imprimindo o level 1:")
for chave, valor in pairs(map.loadInArray("map_level1.txt")) do
  print(valor)
end
