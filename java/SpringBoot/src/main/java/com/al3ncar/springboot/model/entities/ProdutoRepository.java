package com.al3ncar.springboot.model.entities;

import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.repository.query.Param;

public interface ProdutoRepository extends PagingAndSortingRepository<Produto, Integer> {

    //Criando assinatura do método seguindo convenção do SpringBoot para que o
    //próprio framework faça implementação do método de forma automática.
    //findBy...Containing (onde tem "..." deve-se colocar o nome de um atributos da sua classe.)
    public Iterable<Produto> findByNomeContainingIgnoreCase(String parteNome);

    //MAIS NOMES QUE O SPRING BOOT ACEITA:
    //findByNomeContaining
    //findByNomeIsContaining
    //findByNomeContains

    //findByNomeStartWith
    //findByNomeEndsWith

    //findByNomeNotContaining

    //100% CUSTOMIZADO (está quebrando o código aqui)
    @Query("SELECT p FROM Produto p WHERE p.nome LIKE %:nome%")
    public Iterable<Produto> searchByNameLike(@Param("nome") String nome);
}
