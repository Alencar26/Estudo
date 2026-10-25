import java.util.ArrayDeque;
import java.util.Deque;

public class Pillha {

    public static void main(String[] args) {

        //PILHA DE LIVROS

        Deque<String> livros = new ArrayDeque<>();

        livros.add("O Senhor dos aneis"); // retorna true or flase ao add os livros na pilha
        livros.push("O Hobbit"); // lança exceção caso não consiga

        String retornoLivro = livros.poll(); // retorna NULL qnd a pilha acaba
        String retornoLivro2 =  livros.remove(); // lança exceção qnd a pilha acaba
        String retornoLivro3 = livros.pop(); // parece o remove
    }
}
