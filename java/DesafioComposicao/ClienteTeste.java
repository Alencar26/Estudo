package DesafioComposicao;

import java.util.ArrayList;
import java.util.List;

public class ClienteTeste {

    public static void main(String[] args) {

        List<Item> carrinhoMercado = new ArrayList<>();
        List<Item> carrinhoParaguay = new ArrayList<>();

        carrinhoMercado.add(new Item(new Produto("arroz", 5.99), 10));
        carrinhoMercado.add(new Item(new Produto("feijão", 3.89), 11));
        carrinhoMercado.add(new Item(new Produto("carne", 49.99), 2));
        carrinhoMercado.add(new Item(new Produto("salada", 2.70), 4));
        carrinhoMercado.add(new Item(new Produto("frango", 11.50), 5));

        carrinhoParaguay.add(new Item(new Produto("Celular", 556.99), 3));
        carrinhoParaguay.add(new Item(new Produto("Fone", 30.89), 7));
        carrinhoParaguay.add(new Item(new Produto("Carregador", 49.99), 2));
        carrinhoParaguay.add(new Item(new Produto("TV", 1500.50), 4));
        carrinhoParaguay.add(new Item(new Produto("Notebook", 3500.89), 1));

        List<Compra> comprasDoAlberto = new ArrayList<>();

        comprasDoAlberto.add(new Compra(carrinhoMercado));
        comprasDoAlberto.add(new Compra(carrinhoParaguay));

        Cliente cliente = new Cliente("Alberto Roberto", comprasDoAlberto);

        double quantoRobertoGastou = cliente.gastos();

        System.out.format("%.2f", quantoRobertoGastou);

    }
}
