local map = {}

function map.load(name)
  local file = io.open(name, "r")
  local data = file:read("*a")
  io.close(file)
  return data
end

return map
