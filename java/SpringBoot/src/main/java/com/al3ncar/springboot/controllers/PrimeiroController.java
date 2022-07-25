package com.al3ncar.springboot.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PrimeiroController {

    @RequestMapping(method = RequestMethod.GET, path = {"ola", "/saudacao"})
    //outra forma:
    //@GetMapping(path = {"/ola", "/saldacoes"})
    public String ola() {
        return "Olá Spring Boot";
    }
}