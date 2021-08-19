import java.util.LinkedList;
import java.util.Queue;

public class Filas {

    public static void main(String[] args) {

        Queue<String> fila = new LinkedList<>();

        //FORMAS DE ADICIONAR
        fila.add("teste1"); // retorna um erro qnd a fila está cheia
        fila.offer("teste2"); // retorna true ou false qnd add na fila. Caso esteja cheia retorna false.

        //FORMAS DE PEGAR O PRÓXIMO ELEMENTO DA FILA (SEM REMOVER)
        System.out.println(fila.peek()); // qnd a fila está vazia retorna NULL
        System.out.println(fila.element()); // qnd a fila está vazia lança uma exceção

        //FORMA DE PEGAR O ELEMENTO DA FILA E REMOVER AO MESMO TEMPO
        System.out.println(fila.poll());

        //limpar a fila
        fila.clear();
        //verificar tamnho
        int tamanho = fila.size();
        //verificar se está vazia
        boolean vazia = fila.isEmpty();

    }
}
