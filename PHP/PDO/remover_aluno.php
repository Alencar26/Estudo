<?php

use Alura\Pdo\Infra\Database\FabricaDeConexaoDB;

require 'vendor/autoload.php';

$pdo = FabricaDeConexaoDB::criarConexao();

$sqlDelete = 'DELETE FROM students WHERE id = :id;';
$sqlDeletePreparado = $pdo->prepare($sqlDelete);

$sqlDeletePreparado->bindValue(':id', 2, PDO::PARAM_INT);

if ($sqlDeletePreparado->execute()) {
    echo 'Deletado com sucesso';
}