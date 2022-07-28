package com.al3ncar.springboot.model.entities;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.validation.constraints.Max;
import javax.validation.constraints.Min;
import javax.validation.constraints.NotBlank;

@Entity(name = "produtos")
public class Produto {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @NotBlank //não pode ser vazio
    private String nome;

    @Min(0) //valor minimo
    private double preco;

    @Min(0)
    @Max(1) //valor maximo
    private double desconso;

    public Produto() { }

    public Produto(String nome ,double preco, double desconso) {
        this.preco = preco;
        this.desconso = desconso;
        this.nome = nome;
    }

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
