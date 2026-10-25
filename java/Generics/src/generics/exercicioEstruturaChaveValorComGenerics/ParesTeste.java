package generics.exercicioEstruturaChaveValorComGenerics;

public class ParesTeste {

    public static void main(String[] args) {

        Pares<Integer, String> resultadoConcurso = new Pares<>();

        resultadoConcurso.adicionar(1, "Maria");
        resultadoConcurso.adicionar(2, "Ana");
        resultadoConcurso.adicionar(3, "Pedro");
        resultadoConcurso.adicionar(4, "Helo");
        resultadoConcurso.adicionar(3, "Fabricio");

        System.out.println(resultadoConcurso.getValor(1));
        System.out.println(resultadoConcurso.getValor(4));
        System.out.println(resultadoConcurso.getValor(3));
    }
}
