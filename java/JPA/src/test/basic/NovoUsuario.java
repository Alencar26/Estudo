package test.basic;

import model.basic.Usuario;

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
        Usuario user = new Usuario("Andre", "andre@email.com");
        //add no banco
        em.getTransaction().begin();
        em.persist(user);
        em.getTransaction().commit();
        //final da inclusão no banco.

        System.out.println("O Id gerado foi: " + user.getId());

        em.close();
        emF.close();
    }
}
