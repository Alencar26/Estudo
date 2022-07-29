package com.al3ncar.springboot.controllers;

import com.al3ncar.springboot.model.entities.Produto;
import com.al3ncar.springboot.model.entities.ProdutoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.util.Optional;

@RestController
@RequestMapping("/api/produtos")
public class ProdutoController {

    @Autowired
    // anotation para injeção de dependência para implementação da classe abixo
    private ProdutoRepository produtoRepository;

    @RequestMapping(method = {RequestMethod.POST, RequestMethod.PUT}) //usando dois verbos Http no mesmo método.
    public @ResponseBody Produto salvarProduto(@Valid Produto produto) {
        produtoRepository.save(produto);
        return produto;
    }

    @GetMapping
    public Iterable<Produto> obterProdutos() {
        return produtoRepository.findAll();
    }

    @GetMapping(path = "/nome/{parteDoNome}")
    public Iterable<Produto> obterProdutosPorNome(@PathVariable String parteDoNome) {
        return produtoRepository.findByNomeContainingIgnoreCase(parteDoNome); // método customizado na interface. (ler descrição lá)
    }

    //Efetuar consulta com paginação
    @GetMapping(path = "/pagina/{numeroPagina}")
    public Iterable<Produto> obterPordutosPorPagina(@PathVariable int numeroPagina) {
        Pageable page = PageRequest.of(numeroPagina - 1, 5); //numero da página está com -1 pois começa em zerp.
        return produtoRepository.findAll(page);
    }

    @GetMapping(params = "id")
    public Optional<Produto> obterPorId(@RequestParam(name = "id") int id) {
        return produtoRepository.findById(id);
    }

//    @PutMapping
//    public Produto alterarProduto(@Valid Produto produto) {
//        produtoRepository.save(produto);
//        return produto;
//    }

    @DeleteMapping(path = "/{id}")
    public String excluirProduto(@PathVariable(name = "id") int id) {
        produtoRepository.deleteById(id);
        return "Deletado com sucesso";
    }
}
