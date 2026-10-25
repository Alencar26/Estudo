package test.basic;

import model.basic.Usuario;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

public class AlterarUsuario1 {

    public static void main(String[] args) {

        //configurações para manipular o banco
        EntityManagerFactory emF = Persistence
                .createEntityManagerFactory("JPA"); //informe o mesmo nome do arquivo persistence.xml
        EntityManager em = emF.createEntityManager();

        em.getTransaction().begin();
        Usuario usuario = em.find(Usuario.class, 1L);
        usuario.setNome("Mudei de Nome");
        usuario.setEmail("mudei_@gmail.com");
        //responsável por fazer alteração dos dados já existente no banco
        em.merge(usuario);
        em.getTransaction().commit();

        em.close();
        emF.close();
    }
}
