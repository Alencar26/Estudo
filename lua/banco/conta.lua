local conta = {}

function conta.novo(deposito_inicial)
  local conta = {
    saldo = deposito_inicial
  }

  function conta:depositar(valor)
    self.saldo = self.saldo + valor
  end

  return conta
end


return conta
