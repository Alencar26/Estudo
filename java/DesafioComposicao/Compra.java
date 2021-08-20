package DesafioComposicao;

import java.util.ArrayList;
import java.util.List;

public class Compra {

    List<Item> itens = new ArrayList<>();

    void adicionaNoCarrinho(Item item){
        itens.add(item);
    }

    double obterValorTotal(){
      double valorTotal = 0;
      for(Item item : itens){
          valorTotal += (item.quantidade * item.produto.preco);
      }
      return valorTotal;
    }
}
