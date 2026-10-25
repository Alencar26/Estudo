<?php

use Alura\Pdo\Domain\Model\Student;
use Alura\Pdo\Infra\Database\FabricaDeConexaoDB;
use Alura\Pdo\Infra\Repositorio\PdoRepositorioEstudante;

require_once 'vendor/autoload.php';

$pdo = FabricaDeConexaoDB::criarConexao();
$repository = new PdoRepositorioEstudante($pdo);

/** @var Alura\Pdo\Domain\Model\Student[] $studentList */
$studentList = $repository->estudantesComTelefone();


echo $studentList[1]->phones()[0]->formattedPhone();
var_dump($studentList);