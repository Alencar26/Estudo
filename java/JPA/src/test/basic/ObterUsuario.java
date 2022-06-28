package test.basic;

import model.basic.Usuario;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

public class ObterUsuario {

    public static void main(String[] args) {

        //configurações para manipular o banco
        EntityManagerFactory emF = Persistence
                .createEntityManagerFactory("JPA"); //informe o mesmo nome do arquivo persistence.xml
        EntityManager em = emF.createEntityManager();

        Usuario usuario = em.find(Usuario.class, 3L); // o id é um long por isso o L.
        System.out.println(usuario.getNome());

        em.close();
        emF.close();
    }
}
