
//buscar por URLs
const r = /https?:\/\/[^\s]+/g

//obtem sequencia de numeros depois de letras
const r2 = /\B\d+/g

//corrigindo busca com erros de ortografia
const harryPotter = /\b[HhrR][aA4][rnN][rR][yY]\s*[Pp][oO0][tT][tT][eE][rR]\b/g