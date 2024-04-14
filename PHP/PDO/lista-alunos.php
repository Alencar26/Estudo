<?php
use Alura\Pdo\Domain\Model\Student;

require_once 'vendor/autoload.php';

$databasePath = __DIR__.'/db.sqlite';
$pdo = new PDO('sqlite:'.$databasePath);

#PEGANDO TODOS VALORES DO BANCO DE DADOS;

$sqlSelectAll = 'SELECT * FROM students;';

$listaDadosEstudantes = $pdo->query($sqlSelectAll)->fetchAll(PDO::FETCH_ASSOC);
$listaEstudantes = array();

foreach ($listaDadosEstudantes as $dadoEstudante) {
    $listaEstudantes[] = new Student($dadoEstudante['id'], 
                                    $dadoEstudante['name'],
                                    new DateTimeImmutable($dadoEstudante['birth_date']));
}

var_dump($listaEstudantes);

#PEGANDO SÓ UMA VALOR DO BANCO DE DADOS;

$sqlSelectId = 'SELECT * FROM students WHERE id = 3;';
$aluno3 = $pdo->query($sqlSelectId)->fetch(PDO::FETCH_ASSOC);

var_dump($aluno3);


#ALTERNATIVA PARA PEGAR TODOS ELEMENTOS SEM ESTOURAR A MEMÓRIA

while ($aluno3 = $pdo->query($sqlSelectAll)->fetch(PDO::FETCH_ASSOC)) {
    $aluno = new Student(
        $aluno3['id'],
        $aluno3['name'],
        new DateTimeImmutable($aluno3['birth_date'])
    );

    echo $aluno->age() . PHP_EOL;
}
exit();