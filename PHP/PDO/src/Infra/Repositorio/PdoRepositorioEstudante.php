<?php

namespace Alura\Pdo\Infra\Repositorio;
use Alura\Pdo\Domain\Repositorio\RepositorioEstudante;
use Alura\Pdo\Infra\Database\FabricaDeConexaoDB;
use Alura\Pdo\Domain\Model\Student;
use Alura\Pdo\Domain\Model\Phone;
use \PDO;

class PdoRepositorioEstudante implements RepositorioEstudante {

    /*
        DAO VS REPOSITORY

        DAO:
        get; create; update e delete;

        REPOSITORY:
        all, findById, add, remove, etc;
    */


    private PDO $conexaoDB;

    public function __construct(PDO $connection) {
        $this->conexaoDB = $connection;
    }

    public function todosEstudantes(?string $where = null): array {
        $querySQL = "";
        if ($where === null) {
            $querySQL = "SELECT * FROM students;";
        } else {
            $querySQL = "SELECT * FROM students WHERE".$where.";";
        }
        $queryPrepareda = $this->conexaoDB->prepare($querySQL);
        $queryPrepareda->setFetchMode(PDO::FETCH_ASSOC);
        $queryPrepareda->execute();
        
        $listaDadosAlunos = $queryPrepareda->fetchAll();
        $listaAlunos = [];

        foreach ($listaDadosAlunos as $dadoAluno) {
            $student = new Student(
                $dadoAluno["id"],
                $dadoAluno["name"],
                new \DateTimeImmutable($dadoAluno["birth_date"])
            );

            // $this->fillPhoneOf($student);
            $listaAlunos[] = $student;
        }
        return $listaAlunos;
    }

    // private function fillPhoneOf(Student $student): void {
    //     $querySQL = 'SELECT id, area_code, number FROM phones WHERE student_id = ?';
    //     $queryPrepareda = $this->conexaoDB->prepare($querySQL);
    //     $queryPrepareda->bindValue(1, $student->id(), PDO::PARAM_INT);
    //     $queryPrepareda->execute();

    //     $phoneDataList = $queryPrepareda->fetchAll();
    //     foreach ($phoneDataList as $dataPhone) {
    //         $phone = new Phone(
    //             $dataPhone['id'],
    //             $dataPhone['area_code'],
    //             $dataPhone['number'],
    //         );
    //         $student->addPhone($phone);
    //     }
    // }

    public function estudantesComTelefone(): array {
        $querySQL = 'SELECT students.id,
                            students.name,
                            students.birth_date,
                            phones.id AS phone_id,
                            phones.area_code,
                            phones.number
                    FROM students
                    JOIN phones ON students.id = phones.student_id';
        $stmt = $this->conexaoDB->query( $querySQL );
        $resultQuery = $stmt->fetchAll();
        $listaEstudantes = [];
        foreach ($resultQuery as $resultado) {
            if (!array_key_exists($resultado['id'], $listaEstudantes)) {
                $listaEstudantes[$resultado['id']] = new Student(
                    $resultado['id'],
                    $resultado['name'],
                    new \DateTimeImmutable($resultado['birth_date'])
                );
            }
            $phone = new Phone($resultado['phone_id'], $resultado['area_code'], $resultado['number']);
            $listaEstudantes[$resultado['id']]->addPhone($phone);
        }
        return $listaEstudantes;
    } 

    public function estudanteAniversarioEm(\DateTimeImmutable $birth_date): array { 
        return $this->todosEstudantes("birth_date =".$birth_date->format("Y-m-d"));
    }

    public function obtemEstudatePorId(int $id): Student {
        $AlunoArray = $this->todosEstudantes("id =".$id);
        //return new Student($AlunoArray["id"], $AlunoArray["name"], $AlunoArray["birthDate"]);
        return $AlunoArray();
    }
    
    public function salvar(Student $estudante): bool {
        
        if($estudante->id() === null) {
            return $this->inserir($estudante);
        }
        
        return $this->atualizar($estudante);
    }

    public function inserir(Student $estudante): bool {
        $querySQL = "INSERT INTO students (name, birth_date) VALUES (:name,:birth_date);";
        $queryPreparada = $this->conexaoDB->prepare($querySQL);

        if($queryPreparada == false) {
            throw new \RuntimeException($this->conexaoDB->errorInfo());
        }

        $sucesso = $queryPreparada->execute([
            ":name" => $estudante->name(),
            ":birth_date" => $estudante->birthDate()->format('Y-m-d')
        ]);

        if($sucesso) {
            $estudante->defineId($this->conexaoDB->lastInsertId());
        }

        return $sucesso;
    }

    public function atualizar(Student $estudante): bool {
        $querySQL = "UPDATE students SET name = :name, birth_date = :birth_date WHERE id = :id;";
        $queryPreparada = $this->conexaoDB->prepare($querySQL);
        $sucesso = $queryPreparada->execute([
            ":name" => $estudante->name(),
            ":birth_date" => $estudante->birthDate()->format("Y-m-d"),
            ":id" => $estudante->id()
        ]);

        return $sucesso;
    }

    public function remover(Student $estudante): bool { 
        $querySQL = 'DELETE FROM students WHERE id = :id;';
        $queryPreparada = $this->conexaoDB->prepare($querySQL);
        $queryPreparada->bindValue(':id', $estudante->id(), PDO::PARAM_INT);

        return $queryPreparada->execute();
    }
}
