package model.heranca;

import javax.persistence.Entity;

@Entity
public class AlunoBolsista extends Aluno{

    private double valorBolda;

    public AlunoBolsista(Long matricula, String nome, double valorBolda) {
        super(matricula, nome);
        this.valorBolda = valorBolda;
    }

    public double getValorBolda() {
        return valorBolda;
    }

    public void setValorBolda(double valorBolda) {
        this.valorBolda = valorBolda;
    }
}
