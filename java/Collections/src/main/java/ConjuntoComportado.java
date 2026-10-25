import java.util.HashSet;
import java.util.Set;
import java.util.TreeSet;

public class ConjuntoComportado {

    public static void main(String[] args) {

        // hashSet NÃO respeita ordem
        Set<String> conjunto = new HashSet<>();
        conjunto.add("Ana");
        conjunto.add("Helo");
        conjunto.add("Paçoca");

        System.out.println(conjunto);

        // TreeSet RESPEITA ordem
        Set<String> conjuntoOrdem = new TreeSet<>();
        conjuntoOrdem.add("Ana");
        conjuntoOrdem.add("Helo");
        conjuntoOrdem.add("Paçoca");

        System.out.println(conjuntoOrdem);


        //Coleções não aceitam tipos primitivos
        // exemplo:
        // Set<int> colecao = new TreeSet<>(); <-- isso não é permitido
        Set<Integer> colecao = new TreeSet<>();  // <-- isso é válido

    }
}
