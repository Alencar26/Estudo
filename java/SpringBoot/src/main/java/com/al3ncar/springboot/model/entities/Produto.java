package com.al3ncar.springboot.model.entities;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity(name = "produtos")
public class Produto {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;
    private double preco;
    private double desconso;

    public Produto() { }

    public Produto(String nome ,double preco, double desconso) {
        this.preco = preco;
        this.desconso = desconso;
        this.nome = nome;
    }

    private String nome;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public double getPreco() {
        return preco;
    }

    public void setPreco(double preco) {
        this.preco = preco;
    }

    public double getDesconso() {
        return desconso;
    }

    public void setDesconso(double desconso) {
        this.desconso = desconso;
    }
}
