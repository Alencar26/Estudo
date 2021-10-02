package lambdas;

import java.util.function.BiFunction;
import java.util.function.BinaryOperator;
import java.util.function.Function;

public class OperadorBinario {

    public static void main(String[] args) {

        BinaryOperator<Double> media = (n1, n2) -> (n1 + n2) / 2;
        System.out.println(media.apply(2.2,6.3));

        // função mais flexivel, permite escolher o tipo das entradas e o tipo do retorno
        BiFunction<Double, Double, String> resultado =
                (n1, n2) -> ((n1 + n2) / 2) >= 7.0? "Aprovado" : "Reprovado";
        System.out.println(resultado.apply(4.5,7.0));

        //conceito composição de função
        Function<Double, String> conceitoMedia = m -> m >= 7 ? "Aprovado" : "Reprovado";

        System.out.println(
                media.andThen(conceitoMedia)
                        .apply(8.5, 3.8)
        );
    }
}
