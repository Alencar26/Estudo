<?php

namespace Alura\Pdo\Infra\Repositorio;
use Alura\Pdo\Domain\Repositorio\RepositorioEstudante;
use Alura\Pdo\Infra\Database\FabricaDeConexaoDB;
use Alura\Pdo\Domain\Model\Student;
use \PDO;

class PdoRepositorioEstudante implements RepositorioEstudante {

    private PDO $conexaoDB;

    public function __construct() {
        $this->conexaoDB = FabricaDeConexaoDB::criarConexao();
    }

    public function todosEstudantes(?string $where = null): array {
        $querySQL = "";
        if ($where === null) {
            $querySQL = "SELECT * FROM stadents;";
        } else {
            $querySQL = "SELECT * FROM stadents WHERE".$where.";";
        }
        $querySQL = "SELECT * FROM stadents;";
        $queryPrepareda = $this->conexaoDB->prepare($querySQL);
        $queryPrepareda->setFetchMode(PDO::FETCH_ASSOC);
        $queryPrepareda->execute();
        
        $listaDadosAlunos = $queryPrepareda->fetchAll();
        $listaAlunos = [];

        foreach ($listaDadosAlunos as $dadoAluno) {
            $listaAlunos[] = new Student(
                $dadoAluno["id"],
                $dadoAluno["name"],
                new \DateTimeImmutable($dadoAluno["birth_date"])
            );
        }
        return $listaAlunos;
    }
    public function estudanteAniversarioEm(\DateTimeImmutable $birth_date): array { 
        return $this->todosEstudantes("birth_date =".$birth_date->format("Y-m-d"));
    }

    public function obtemEstudatePorId(int $id): Student {
        $AlunoArray = $this->todosEstudantes("id =".$id);
        return new Student($AlunoArray["id"], $AlunoArray["name"], $AlunoArray["birthDate"]);
    }
    
    public function salvar(Student $estudante): bool { 
        $querySQL = "INSERT INTO students (name, birth_date) VALUES (:name,:birth_date);";
        $queryPreparada = $this->conexaoDB->prepare($querySQL);

        $queryPreparada->bindValue(':name', $estudante->name());
        $queryPreparada->bindValue(':birth_date', $estudante->birthDate()->format('Y-m-d'));

        return $queryPreparada->execute();
    }
    public function remover(Student $estudante): bool { 
        $querySQL = 'DELETE FROM students WHERE id = :id;';
        $queryPreparada = $this->conexaoDB->prepare($querySQL);
        $queryPreparada->bindValue(':id', $estudante->id(), PDO::PARAM_INT);

        return $queryPreparada->execute();
    }
}
