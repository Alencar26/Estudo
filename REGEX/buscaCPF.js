const fs = require("fs");
const bancoCsv = "database.csv";
const banco = fs.readFileSync(bancoCsv, "utf8");


/*
[.-] --> classe que agrupa meta caracteres, ou seja depois de 3 dígitos estou esperando '.' ou '-'
? --> metacaracter que informa que a [.-] é opcional. Ele pode aparece 1 ou mais vezes ou nem aparecer.

*/
const regexCPF = /\d{3}[.-]?\d{3}[.-]?\d{3}[.-]?\d{2}/g;

const matchRegex = banco.match(regexCPF);
console.log(matchRegex);
