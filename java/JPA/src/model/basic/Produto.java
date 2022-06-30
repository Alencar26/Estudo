package model.basic;

import javax.persistence.*;

@Entity
@Table(name = "produtos", schema = "curso_java")
public class Produto {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "nome_produto", length = 200, nullable = false)
    private String nome;

    //precision = qnt de números - scale = qnt número casa decimal
    @Column(name = "preco_produto", nullable = false, precision = 11, scale = 2)
    private Double preco;

    public Produto() { }

    public Produto(Long id, String nome, Double preco) {
        this.id = id;
        this.nome = nome;
        this.preco = preco;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public Double getPreco() {
        return preco;
    }

    public void setPreco(Double preco) {
        this.preco = preco;
    }
}
