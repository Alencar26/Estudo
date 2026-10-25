import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

public class MinMax {

    public static void main(String[] args) {

        Aluno a1 = new Aluno("Ana", 7.1);
        Aluno a2 = new Aluno("Bia", 7.2);
        Aluno a3 = new Aluno("Daniel", 8.1);
        Aluno a4 = new Aluno("Gui", 10);

        List<Aluno> alunos = Arrays.asList(a1, a2, a3, a4);

        Comparator<Aluno> melhorAluno = (aluno1, aluno2) -> {
            if(aluno1.nota > aluno2.nota) return 1;
            if(aluno1.nota < aluno2.nota) return -1;
            return 0;
        };

        //melhor nota
        System.out.println(alunos.stream().max(melhorAluno).get());
        //pior nota
        System.out.println(alunos.stream().min(melhorAluno).get());
    }
}
