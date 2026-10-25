package lambdas;

import java.util.function.Function;

public class Funcao {

    public static void main(String[] args) {

        Function<Integer, String> parOuImpar = numero -> numero % 2 == 0 ? "Par" : "Impar";
        System.out.println(parOuImpar.apply(30));

        Function<String, String> oResultadoE = valor -> "O resultado é: " + valor;
        System.out.println(parOuImpar.andThen(oResultadoE).apply(30));
    }
}
