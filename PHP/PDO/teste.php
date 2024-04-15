<?php

use Alura\Pdo\Domain\Repositorio\RepositorioEstudante;
use Alura\Pdo\Infra\Repositorio\PdoRepositorioEstudante;

function enviaEmailaniversariante(RepositorioEstudante $repositorioEstudante) {
    $repositorioEstudante->estudanteAniversarioEm(new \DateTimeImmutable());
}

//EXEMPLO DE COMO USAR INTERFACE

enviaEmailaniversariante(new PdoRepositorioEstudante());
