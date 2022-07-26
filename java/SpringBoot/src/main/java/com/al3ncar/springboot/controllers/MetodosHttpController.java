package com.al3ncar.springboot.controllers;

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/metodos")
public class MetodosHttpController {

    @GetMapping
    public String get() {
        return "requisição GET";
    }

    @PostMapping
    public String post() {
        return "requisição POST";
    }

    @PutMapping
    public String put() {
        return "requisição PUT";
    }

    @PatchMapping
    public String patch() {
        return "requisição PATCH";
    }

    @DeleteMapping
    public String delete() {
        return "requisição DELETE";
    }
}
