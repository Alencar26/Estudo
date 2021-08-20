package DesafioComposicao;

import java.util.List;

public class Cliente {

    String nome;
    List<Compra> compras;

    public Cliente(String nome, List<Compra> compras) {
        this.nome = nome;
        this.compras = compras;
    }

    double gastos(){
        double gastos = 0;
        for (Compra compra : compras) {
            gastos += compra.obterValorTotal();
        }
        return gastos;
    }
}
