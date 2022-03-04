import java.util.Arrays;
import java.util.List;

public class Filter {

    public static void main(String[] args) {

        AlunoFilter a1 = new AlunoFilter("Ana", 7.8);
        AlunoFilter a2 = new AlunoFilter("Bia", 5.8);
        AlunoFilter a3 = new AlunoFilter("Daniel", 9.8);
        AlunoFilter a4 = new AlunoFilter("Gui", 6.8);
        AlunoFilter a5 = new AlunoFilter("Rebeca", 7.1);
        AlunoFilter a6 = new AlunoFilter("Pedro", 8.8);


        List<AlunoFilter> alunos = Arrays.asList(a1, a2, a3, a4, a5, a6);

        alunos.stream()
                .filter(a -> a.nota >= 7)
                .map(a -> "Parabéns " + a.nome + "! Você foi aprovado")
                .forEach(System.out::println);
    }
}
