package model.basic;

import infra.DAO;

public class ProdutoDAO extends DAO<Produto> {
    // Forma de resolver o DAO generics
    public ProdutoDAO() {
        super(Produto.class);
    }
}
