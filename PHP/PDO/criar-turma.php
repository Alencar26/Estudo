<?php

require_once 'vendor/autoload.php';

use Alura\Pdo\Domain\Model\Student;
use Alura\Pdo\Infra\Database\FabricaDeConexaoDB;
use Alura\Pdo\Infra\Repositorio\PdoRepositorioEstudante;

$connection = FabricaDeConexaoDB::criarConexao();
$studentRepository = new PdoRepositorioEstudante($connection);

//Transação de tudo ou nada -> só retorna sucesso se todas as querys derem sucesso.
$connection->beginTransaction();

try {
    $estudanteA = new Student(
        null,
        "Linus Torvald", 
        new DateTimeImmutable('1964-06-12'));
    
    $estudanteB = new Student(
        null,
        "Nolan Bushnnel", 
        new DateTimeImmutable('1944-02-24'));
    
    $studentRepository->salvar($estudanteA);
    $studentRepository->salvar($estudanteB);

    $connection->commit();
} catch (\RuntimeException $e) {
    echo "Erro: ".$e->getMessage();
    $connection->rollBack();
}

