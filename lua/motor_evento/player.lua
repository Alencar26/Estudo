local player = {}

function player.novo()
  local _player = {
    ouvintes = {}
  }

  local executarCallbacks = function(evento)
    for _, v in pairs(_player.ouvintes[evento]) do
      for _, callback in pairs(v) do
        callback()
      end
    end
  end

  -- Adicionar ouvinte
  function _player:adicionarOuvinte(evento, quem, callback)
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
  function _player:removerOuvinte(quem)
    table.remove(self.ouvintes["onStarted"][quem])
    table.remove(self.ouvintes["onPoweredUp"][quem])
    table.remove(self.ouvintes["onPoweredDown"][quem])
  end

  -- Ações do usuário
  -- Normalmente chamada: removeEventListener
  function _player:iniciar()
    executarCallbacks("onStarted")
  end

  function _player:poderoso()
    executarCallbacks("onPoweredUp")
  end

  function _player:fraco()
    executarCallbacks("onPoweredDown")
  end

  return _player
end

return player
