const fs = require("fs");
const bancoCsv = "database.csv";
const banco = fs.readFileSync(bancoCsv, "utf8");

/*
[/\s.] --> indica que após as sequencia de núemros eu posso ter '/' ou espaço vazio (\s) ou '.' ponto.
? --> metacaracter que informa que a [/\s.] é opcional. Ele pode aparece 1 ou mais vezes ou nem aparecer.
$ ---> indica que essa regex está no final de uma linha
/{REGEX}/gm --> 'gm' indica que é um arquivo com multiplas linhas, sem isso não funciona correto
*/
const regexDataNascimento = /\d{2}[/\s.]?\d{2}[/\s.]?\d{4}$/gm;

const matchRegex = banco.match(regexDataNascimento);
console.log(matchRegex);

const regexDataNascimento2 = /\d{2}[/\s.]?[0-1]{1}\d{1}[/\s.]?[1-2]{1}\d{3}/g;
const matchRegex2 = banco.match(regexDataNascimento2);
console.log(matchRegex2);