const fs = require("fs");
const bancoCsv = "database.csv";
const banco = fs.readFileSync(bancoCsv, "utf8");

/*
/{AQUI É A REGEX}/g --> o 'g' indica que estou querendo percorrer o arquivo todo
\d+ -> indica que estou pegando um número (+ indica que é uma sequencia)
*/
const regexSequenciaDeNumeros = /\d+/g;

const matchNumeros = banco.match(regexSequenciaDeNumeros);
console.log("Números:");
console.log(matchNumeros);


const regexDDD = /\(\d+\)/g
const matchDDD = banco.match(regexDDD);
console.log("DDD:");
console.log(matchDDD);

/*
\s -> indica um espaço vazio nos caracteres
Nos () teve a necessidade de colocar \ (escape) na frente
*/
const regexTelefone = /\(\d+\)\s\d+-\d+/g;
const matchTelefone = banco.match(regexTelefone);
console.log("Telefones:");
console.log(matchTelefone);

/*
\d{n} -> no lugar de 'n' informar quantos números em sequencia devem aparecer
*/
const regexApenasCelulares = /\(\d{2}\)\s\d{5}-\d{4}/g
const matchApenasCelulares = banco.match(regexApenasCelulares);
console.log("Apenas Celulares:");
console.log(matchApenasCelulares);

const regexApenasFixo = /\(\d{2}\)\s\d{4}-\d{4}/g;
const matchApenasFixo = banco.match(regexApenasFixo);
console.log("Apenas Telefone Fixo:");
console.log(matchApenasFixo);

// \d{4,5} --> indica que a sequencia pode variar de 4 a 5 dígitos
const regexTodosTelefones = /\(\d{2}\)\s\d{4,5}-\d{4}/g;
const matchTodosTelefones = banco.match(regexTodosTelefones);
console.log("Todos Telefone 2º Alternativa de REGEX:");
console.log(matchTodosTelefones);