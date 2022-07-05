package test.muitospramuitos;

import infra.DAO;
import model.muitospramuitos.Ator;
import model.muitospramuitos.Filme;

public class NovoFilmeAtor {

    public static void main(String[] args) {

        Filme filme1 = new Filme("Missão Impossível", 8.5);
        Filme filme2 = new Filme("Senhor dos Aneis", 10.0);

        Ator ator1 = new Ator("Tom Cruise");
        Ator ator2 = new Ator("Ian McKellen");

        filme1.adicionarAtor(ator1);
        filme1.adicionarAtor(ator2);

        filme2.adicionarAtor(ator2);

        DAO<Filme> dao = new DAO<>(Filme.class);
        //como nas classes foi colocado a opção cascate
        //qnd efetuar operação com filme, será feito com
        //atores também.
        dao.incluirAtomico(filme1);
        dao.incluirAtomico(filme2);
    }
}
