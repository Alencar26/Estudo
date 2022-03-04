import java.util.Arrays;
import java.util.List;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.function.UnaryOperator;

public class DesafioMap {

    public static void main(String[] args) {

        Consumer<Integer> print = System.out::println;
        List<Integer> nums = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9);

        /*
        * 1. Número com string binária... 6 => "110"
        * 2. Inverter a string... "110" => "011"
        * 3. Converter de volta para integer => "011" => 3
        */

        Function<Integer, String> converterInteiroParaBinario = n -> Integer.toBinaryString(n);
        UnaryOperator<String> inverterBinario = b -> new StringBuilder(b).reverse().toString();
        Function<String, Integer> converterBinarioParaInteiro = s -> Integer.parseInt(s, 2);

        nums.stream()
                .map(converterInteiroParaBinario)
                .map(inverterBinario)
                .map(converterBinarioParaInteiro)
                .forEach(print);


    }
}
