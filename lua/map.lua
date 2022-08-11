local map = {}

function map.load(name)
  local file = io.open(name, "r")
  local data = file:read("*a")
  io.close(file)
  return data
end

function map.loadInArray(name)
  local file = io.open(name)
  local data = {}

  for line in file:lines() do
    table.insert(data, line)
  end

  return data
end

return map
