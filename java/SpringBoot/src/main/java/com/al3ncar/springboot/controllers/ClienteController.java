package com.al3ncar.springboot.controllers;

import com.al3ncar.springboot.model.entities.Cliente;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(path = "/clientes")
public class ClienteController {

    @GetMapping(path = "/qualquer")
    public Cliente obterCliente() {
        return new Cliente(28, "André", "123.456.789-00");
    }

    @GetMapping("/{id}")
    public Cliente obterClientePorId1(@PathVariable("id") int id) {
        return new Cliente(id, "Maria", "987.654.321-00");
    }

    //Forma mais comum de passar parâmetros na URL.
    @GetMapping
    public Cliente obterClientePorId2(@RequestParam(name = "id") int id) {
        return new Cliente(id, "Julio", "444.333.222-11");
    }
}
