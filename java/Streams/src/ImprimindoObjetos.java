import java.util.Arrays;
import java.util.Iterator;
import java.util.List;
import java.util.stream.Stream;

public class ImprimindoObjetos {

    public static void main(String[] args) {

        List<String> aprovados = Arrays.asList("Lu", "Gui", "Luca", "Ana");

        //usando foreach...
        for (String aprovado : aprovados) {
            System.out.println(aprovado);
        }

        //usando Iterator...
        Iterator<String> iterator = aprovados.iterator();
        while (iterator.hasNext()){
            System.out.println(iterator.next());
        }

        //usando stream
        Stream<String> stream = aprovados.stream();
        stream.forEach(System.out::println); // laço interno
    }
}
