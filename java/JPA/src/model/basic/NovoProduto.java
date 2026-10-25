package model.basic;

import infra.DAO;

public class NovoProduto {

    public static void main(String[] args) {

        DAO<Produto> dao = new DAO<>(Produto.class);

        Produto produto = new Produto("Caneta", 7.45);
        Produto produto2 = new Produto("Notebook", 2956.77);

        //incluir com todos os métodos.
        dao.abrirTransacao().incluir(produto).fecharTransacao().fechar();
       // incluir de forma atômica, só fechando o DAO no final.
        dao.incluirAtomico(produto2).fechar();

        System.out.println("ID do produto 1: " + produto.getId());
        System.out.println("ID do produto 2: " + produto2.getId());
    }
}
