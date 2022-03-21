import java.util.Arrays;
import java.util.List;

public class OutrosMetodos {

    public static void main(String[] args) {

        Aluno a1 = new Aluno("Ana", 7.1);
        Aluno a2 = new Aluno("Bia", 7.2);
        Aluno a3 = new Aluno("Daniel", 8.1);
        Aluno a4 = new Aluno("Gui", 10);
        Aluno a5 = new Aluno("Ana", 7.1);
        Aluno a6 = new Aluno("Pedro", 6.2);
        Aluno a7 = new Aluno("Daniel", 8.1);
        Aluno a8 = new Aluno("Maria", 10);

        List<Aluno> alunos = Arrays.asList(a1, a2, a3, a4, a5, a6, a7, a8);

        // distinct funciona muito similar ao SQL
        // Para o distinct funcionar de forma correta, pois ele está pegando endereço de memória
        // é preciso implementar o "equals" e  o "hashCode" na classe Aluno
        System.out.println("distinct...");
        alunos.stream().distinct().forEach(System.out::println);

        //pular 2 elementos e colocar limite para 2 resultados
        System.out.println("\nSkip/Limit");
        alunos.stream()
                .distinct()
                .skip(2)
                .limit(2)
                .forEach(System.out::println);

        //pegue elementos enquanto uma determinada condição não for atendida
        System.out.println("\ntakeWhile");
        alunos.stream()
                .distinct()
                .skip(2)
                .takeWhile(aluno -> aluno.nota >= 7)
                .forEach(System.out::println);
    }
}
