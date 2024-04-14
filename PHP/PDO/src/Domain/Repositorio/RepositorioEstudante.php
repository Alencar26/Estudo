<?php

namespace Alura\Pdo\Domain\Repositorio;
use Alura\Pdo\Domain\Model\Student;

interface RepositorioEstudante {
    public function todosEstudantes(): array;
    public function estudanteAniversarioEm(\DateTimeImmutable $birth_date): array;
    public function obtemEstudatePorId(int $id): Student;
    public function salvar(Student $estudante): bool;
    public function remover(Student $estudante): bool;
}