import java.util.Arrays;
import java.util.List;
import java.util.function.Predicate;
import java.util.function.Function;

public class Match {

    public static void main(String[] args) {

        Aluno a1 = new Aluno("Ana", 7.1);
        Aluno a2 = new Aluno("Bia", 7.2);
        Aluno a3 = new Aluno("Daniel", 8.1);
        Aluno a4 = new Aluno("Gui", 10);

        List<Aluno> alunos = Arrays.asList(a1, a2, a3, a4);

        Predicate<Aluno> aprovado = a -> a.nota >= 7;
        //negar o predicado
        Predicate<Aluno> reprovado = aprovado.negate();

        //todos os alunos foram aprovados?
        System.out.println("todos alunos aprovados? - R: " + alunos.stream().allMatch(aprovado));
        //algum aluno foi aprovado?
        System.out.println("algum aluno aprovado? - R: " + alunos.stream().anyMatch(aprovado));
        //nenhum aluno foi aprovado?
        System.out.println("nenhum aluno aprovado? - R: " + alunos.stream().noneMatch(aprovado));
        //nenhum aluno foi aprovado? - (negar o predicado)
        System.out.println("nenhum aluno aprovado(negado)? - R: " + alunos.stream().noneMatch(aprovado.negate()));
        //nenhum aluno foi aprovado? - (negar o predicado)
        System.out.println("nenhum aluno reprovado? - R: " + alunos.stream().noneMatch(reprovado));
        
    }
}
