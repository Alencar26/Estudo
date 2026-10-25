package DesafioComposicao;

import java.util.ArrayList;
import java.util.List;

public class Cliente {

    String nome;
    List<Compra> compras = new ArrayList<>();

    public Cliente(String nome) {
        this.nome = nome;
    }

    void adicionaCompra(Compra compra){
        compras.add(compra);
    }

    double gastos(){
        double gastos = 0;
        for (Compra compra : compras) {
            gastos += compra.obterValorTotal();
        }
        return gastos;
    }
}
