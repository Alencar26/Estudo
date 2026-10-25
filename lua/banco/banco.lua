local conta = require("conta")  

local banco_brasil = conta.novo(100)
assert(banco_brasil.saldo == 100)

banco_brasil:depositar(500)
assert(banco_brasil.saldo == 600)
