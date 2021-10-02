package lambdas;

import java.util.function.UnaryOperator;

public class PeradorUnario {

    public static void main(String[] args) {

        UnaryOperator<Integer> maisDois = n -> n + 2;
        UnaryOperator<Integer> vezesDois = n -> n * 2;
        UnaryOperator<Integer> aoQuadrado = n -> n * n;

        int resultado = maisDois.andThen(vezesDois).andThen(aoQuadrado).apply(0);
        System.out.println(resultado);

        //executa antes o que vem no parâmetro (começa do final)
        resultado = aoQuadrado.compose(vezesDois).compose(maisDois).apply(0);
        System.out.println(resultado);
    }
}
