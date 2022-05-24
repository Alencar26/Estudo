package model.basic;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

public class NovoUsuario {

    public static void main(String[] args) {

        //configurações para manipular o banco
        EntityManagerFactory emF = Persistence
                .createEntityManagerFactory("JPA"); //informe o mesmo nome do arquivo persistence.xml
        EntityManager em = emF.createEntityManager();

        //criando um objeto
        Usuario user = new Usuario("Pedro", "pe.dro@email.com");
        //add no banco
        em.getTransaction().begin();
        em.persist(user);
        em.getTransaction().commit();
        //final da inclusão no banco.

        em.close();
        emF.close();
    }
}
