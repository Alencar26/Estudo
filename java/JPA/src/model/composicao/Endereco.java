package model.composicao;

import javax.persistence.Embeddable;

/*  Pelo fato de estar com anotatio Embeddable não será
    criado uma tabela no banco para endereço, mas todas as
    classes que tiverem Endereco como atributo, será criado as
    colunas de endereço nas tabelas correspondentes às classes.
*/

@Embeddable
public class Endereco {

    private String logradouro;
    private String complemento;

    public String getLogradouro() {
        return logradouro;
    }

    public void setLogradouro(String logradouro) {
        this.logradouro = logradouro;
    }

    public String getComplemento() {
        return complemento;
    }

    public void setComplemento(String complemento) {
        this.complemento = complemento;
    }
}
