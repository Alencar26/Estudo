package test.consulta;

import infra.DAO;
import model.muitospramuitos.Ator;
import model.muitospramuitos.Filme;

import java.util.List;

public class ObterFilmes {

    public static void main(String[] args) {

        DAO<Filme> dao = new DAO<>(Filme.class);
        List<Filme> filmes = dao.consultar(
                "obterFilmesComNotaMaiorQue", "nota", 8.1);

        for (Filme filme : filmes) {
            System.out.println(filme.getNome() + "=>" + filme.getNota());

            for (Ator ator : filme.getAtores()) {
                System.out.println(ator.getNome());
            }
        }
    }
}
