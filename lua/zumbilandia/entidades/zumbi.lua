local inimigo = require("inimigo")
local zumbi = {}

function zumbi.novo()
  local zumbi = inimigo.novo(50, "zumbis")
  zumbi.come_cerebro = true
  zumbi.explode = false
  return zumbi
end

function zumbi.novo_bomber()
  local zumbi = zumbi.novo()
  zumbi.come_cerebro = false
  zumbi.explode = true
  return zumbi
end

return zumbi
