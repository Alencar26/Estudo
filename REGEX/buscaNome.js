const fs = require("fs");
const bancoCsv = "database.csv";
const banco = fs.readFileSync(bancoCsv, "utf8");

// À-ÿ --> Obter letras com acentos
// () --> indica que estou agrupando essas regex
// ^ --> indica que a ocorrência precisa estar no ínicio de uma linha
// ?: --> diz que não quero que a ocorrência seguinte seja capturada nesse caso o '\s'
// (()) --> podemos ter grupos aninhados  ou subgrupos
// /{REGEX}/gm --> o 'gm' indica que é multilinhas
const regex = /^([A-Za-zÀ-ÿ]+)(?:\s([A-Za-zÀ-ÿ]+))+/gm;

const matchRegex = banco.match(regex);
console.log("Nomes:");
console.log(matchRegex);
