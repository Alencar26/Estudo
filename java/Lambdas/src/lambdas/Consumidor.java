package lambdas;

import java.util.Arrays;
import java.util.List;
import java.util.function.Consumer;

public class Consumidor {

    public static void main(String[] args) {

        Consumer<Produto> imprimir = produto -> System.out.println(produto.nome);

        Produto produto = new Produto("Caneta", 12.34, 0.09);
        imprimir.accept(produto);

        Produto produto1 = new Produto("Caneta", 12.34, 0.09);
        Produto produto2 = new Produto("Caderno", 12.34, 0.09);
        Produto produto3 = new Produto("Lapis", 12.34, 0.09);
        Produto produto4 = new Produto("Borracha", 12.34, 0.09);
        List<Produto> produtos = Arrays.asList(produto1, produto2, produto3, produto4);

        produtos.forEach(imprimir);
        produtos.forEach(p -> System.out.println(p.preco));
        produtos.forEach(System.out::println); // chama o toString do objeto.
    }
}
