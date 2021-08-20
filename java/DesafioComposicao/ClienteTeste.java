package DesafioComposicao;

import java.util.ArrayList;
import java.util.List;

public class ClienteTeste {

    public static void main(String[] args) {

        Compra compraMercado = new Compra();
        Compra compraParaguay = new Compra();

        compraMercado.adicionaNoCarrinho(new Item(new Produto("arroz", 5.99), 10));
        compraMercado.adicionaNoCarrinho(new Item(new Produto("feijão", 3.89), 11));
        compraMercado.adicionaNoCarrinho(new Item(new Produto("carne", 49.99), 2));
        compraMercado.adicionaNoCarrinho(new Item(new Produto("salada", 2.70), 4));
        compraMercado.adicionaNoCarrinho(new Item(new Produto("frango", 11.50), 5));

        compraParaguay.adicionaNoCarrinho(new Item(new Produto("Celular", 556.99), 3));
        compraParaguay.adicionaNoCarrinho(new Item(new Produto("Fone", 30.89), 7));
        compraParaguay.adicionaNoCarrinho(new Item(new Produto("Carregador", 49.99), 2));
        compraParaguay.adicionaNoCarrinho(new Item(new Produto("TV", 1500.50), 4));
        compraParaguay.adicionaNoCarrinho(new Item(new Produto("Notebook", 3500.89), 1));

        Cliente cliente = new Cliente("Alberto Roberto");

        cliente.adicionaCompra(compraMercado);
        cliente.adicionaCompra(compraParaguay);

        System.out.format("%.2f", cliente.gastos());
    }
}
