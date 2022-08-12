local item = require("item") -- Herança em lua
local pocao = {}

function pocao.novo()
  local pocao = item.novo("pocao", "vida", 50)
  pocao.vida = 100
  return pocao
end

return pocao
