package DesafioComposicao;

import java.util.List;

public class Compra {

    List<Item> itens;

    public Compra(List<Item> itens) {
        this.itens = itens;
    }

    double obterValorTotal(){
      double valorTotal = 0;
      for(Item item : itens){
          valorTotal += (item.quantidade * item.produto.preco);
      }
      return valorTotal;
    }
}
