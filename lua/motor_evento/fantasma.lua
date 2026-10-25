local fantasma = {}

function fantasma.novo(nome)
  local _fantasma = {
    nome = nome,
    ouvintes = {}
  }

  local executarCallbacks = function(evento)
    for _, v in pairs(_fantasma.ouvintes[evento]) do
      for _, callback in pairs(v) do
        callback()
      end
    end
  end

  -- Adicionar ouvinte
  function _fantasma:adicionarOuvinte(evento, quem, callback)
    if not self.ouvintes[evento] then
      self.ouvintes[evento] = {}
    end

    if not self.ouvintes[evento][quem] then
      self.ouvintes[evento][quem] = {}
    end

    table.insert(self.ouvintes[evento][quem], callback)
  end

  -- Remover ouvinte
  -- Normalmente chamado: addEventListener
  function _fantasma:removerOuvinte(quem)
    table.remove(self.ouvintes["onStarted"][quem])
    table.remove(self.ouvintes["onPoweredUp"][quem])
    table.remove(self.ouvintes["onPoweredDown"][quem])
  end


  -- Adicionar ouvinte
  function _fantasma:fugir()
    print("FANTASMA: " .. self.nome .. " está fugindo!")
  end

  -- Remover ouvinte
  function _fantasma:atacar()
    print("FANTASMA: " .. self.nome .. " está atacando!")
  end

  -- Ações do usuário
  function _fantasma:morrer()
    print("FANTASMA: " .. self.nome .. " morreu!")
    executarCallbacks("onDied")
  end

  return _fantasma
end

return fantasma
