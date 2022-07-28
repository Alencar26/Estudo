package com.al3ncar.springboot.controllers;

import com.al3ncar.springboot.model.entities.Produto;
import com.al3ncar.springboot.model.entities.ProdutoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequestMapping("/api/produtos")
public class ProdutoController {

    @Autowired
    // anotation para injeção de dependência para implementação da classe abixo
    private ProdutoRepository produtoRepository;

    @PostMapping
    public @ResponseBody Produto novoProduto(@Valid Produto produto) {
        produtoRepository.save(produto);
        return produto;
    }
}
