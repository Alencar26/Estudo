<?php

use Alura\Pdo\Domain\Model\Student;

require 'vendor/autoload.php';

$databasePath = __DIR__.'/db.sqlite';
$pdo = new PDO('sqlite:'.$databasePath);

$student = new Student(1, "Andre", new DateTimeImmutable("1996-07-26"));

#$sqlInsert = "INSERT INTO students (id, name, birth_date) VALUES ({$student->id()} ,'{$student->name()}', '{$student->birthDate()->format('Y-m-d')}');";

$sqlInsert = "INSERT INTO students (name, birth_date) VALUES (:name,:birth_date);";
$sqlPreparada = $pdo->prepare($sqlInsert);

$sqlPreparada->bindValue(':name', $student->name());
$sqlPreparada->bindValue(':birth_date', $student->birthDate()->format('Y-m-d'));

if ($sqlPreparada->execute()) {
    echo 'Aluno incluído com sucesso';
}

// echo $sqlInsert;
// var_dump($pdo->exec($sqlInsert)); 
 