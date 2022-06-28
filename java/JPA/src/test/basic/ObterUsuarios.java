package test.basic;

import model.basic.Usuario;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;
import javax.persistence.TypedQuery;
import java.util.List;

public class ObterUsuarios {

    public static void main(String[] args) {

        //configurações para manipular o banco
        EntityManagerFactory emF = Persistence
                .createEntityManagerFactory("JPA"); //informe o mesmo nome do arquivo persistence.xml
        EntityManager em = emF.createEntityManager();

        //JPQL
        //Pegar todos os usuários da tabela Usuário
        String jpql = "SELECT u FROM Usuario u";
        TypedQuery<Usuario> query = em.createQuery(jpql, Usuario.class);
        query.setMaxResults(5); // limitei para ter no m´aximo 5 consultas

        List<Usuario> usuarios = query.getResultList();

        for(Usuario u: usuarios) {
            System.out.println("Nome: " + u.getNome() + " - ID: " + u.getId());
        }

        em.close();
        emF.close();
    }
}
