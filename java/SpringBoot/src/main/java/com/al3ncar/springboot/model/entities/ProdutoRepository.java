package com.al3ncar.springboot.model.entities;

import org.springframework.data.repository.PagingAndSortingRepository;

public interface ProdutoRepository extends PagingAndSortingRepository<Produto, Integer> {

    //Criando assinatura do método seguindo convenção do SpringBoot para que o
    //próprio framework faça implementação do método de forma automática.
    //findBy...Containing (onde tem "..." deve-se colocar o nome de um atributos da sua classe.)
    public Iterable<Produto> findByNomeContainingIgnoreCase(String parteNome);
}
