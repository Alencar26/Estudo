import java.util.Arrays;
import java.util.List;
import java.util.function.Consumer;
import java.util.stream.Stream;

public class CriandoStreams {

    public static void main(String[] args) {

        Consumer<String> print = System.out::print;
        Consumer<Integer> println = System.out::println;

        Stream<String> lang = Stream.of("Java ", "Lua ", "JS\n");
        lang.forEach(print);

        String[] maisLangs = { "Python ", "Lisp ", "Perl ", "Go\n" };

        Stream.of(maisLangs).forEach(print);
        Arrays.stream(maisLangs).forEach(print);
        //No método abaixo pode-se escolher um intervalo de quais dados da lista serão listados,
        // o último não será mostrado.
        Arrays.stream(maisLangs, 1, 3).forEach(print);

        List<String> outrasLangs = Arrays.asList("C ", "C# ", "PHP\n");
        outrasLangs.parallelStream().forEach(print);

        //loop infinito
        Stream.generate(() -> "a").forEach(print);
        //loop infinito
        Stream.iterate(0, n -> n + 1).forEach(println);
    }
}
