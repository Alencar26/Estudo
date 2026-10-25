package model.heranca;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Inheritance;
import javax.persistence.InheritanceType;

@Entity
/*
*   Na anotation Inheritace pode-se escolher 3 (três)
*   tipos.
*
*   *Se não houver essa anotation por padrão o JPA vai
*   juntar as classes numa única tabela.
*
*   SINGLE_TABLE: Junta tudo numa única tabela (default).
*   Porém ele cria uma coluna com o discriminador para identificar
*   a tupla (linha da tabela).
*   Caso user essa opção é precio usar a anotation @DiscriminationColumn
*   nessa anotation escolher name, length e discriminatorType.
*
*   TABLE_PER_CLASS: Para cada classe concreta haverá
*   uma tabela no banco de dados.
*
*   JOINED: Cada classe terá uma tabela, porém os dados comuns
*   não ficam duplicados. Eles ficam na tabela principal correspon-
*   dente à classe super. Nas colunas haverá uma coluna para associar
*    o id da tabela pai para fazera relação de Um para Um.
*
*/
@Inheritance(strategy = InheritanceType.SINGLE_TABLE)
public class Aluno  {

    @Id
    private Long matricula;
    private String nome;

    public Aluno() {
    }

    public Aluno(Long matricula, String nome) {
        this.matricula = matricula;
        this.nome = nome;
    }

    public Long getMatricula() {
        return matricula;
    }

    public void setMatricula(Long matricula) {
        this.matricula = matricula;
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }
}
