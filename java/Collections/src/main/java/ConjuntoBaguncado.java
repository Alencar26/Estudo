import java.util.HashSet;
import java.util.Set;

public class ConjuntoBaguncado {

    // Anotação para tirar sinais de warnings no código <- caso não se lembre, tirar para ver a diferença.
    @SuppressWarnings({"rawtypes", "unchecked"})
    public static void main(String[] args) {

        // Nessa situação o HashSet aceita qualquer tipo, menos primitivo.

        HashSet conjunto = new HashSet();

        conjunto.add(1.2);
        conjunto.add(true);
        conjunto.add("teste");
        conjunto.add(2);
        conjunto.add('a');

        System.out.println(conjunto);
        System.out.println("Tamanho é: " + conjunto.size());
        System.out.println(conjunto.contains('a') + " <-- Isso me retorna um boolean (contains())");
        System.out.println(conjunto.remove(2) + " <-- Isso tmbém retorna um boolean (remove())");

        Set nums = new HashSet();

        nums.add(1);
        nums.add(2);
        nums.add(3);
        nums.add(4);

        System.out.println(nums);
        nums.addAll(conjunto);
        nums.clear();

    }
}
