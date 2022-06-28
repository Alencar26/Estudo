package test.basic;

import model.basic.Usuario;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

public class AlterarUsuario2 {

    public static void main(String[] args) {

        //configurações para manipular o banco
        EntityManagerFactory emF = Persistence
                .createEntityManagerFactory("JPA"); //informe o mesmo nome do arquivo persistence.xml
        EntityManager em = emF.createEntityManager();

        em.getTransaction().begin();
        Usuario usuario = em.find(Usuario.class, 1L);

        em.detach(usuario); // abixo
        // isso faz com que o objeto não fique no estado gerenciado do JPA. (evita alteração não desejada no banco)
        //agora só será alterado no banco se o em.merge() existir.

        usuario.setNome("Mudei de novo");
        usuario.setEmail("mudei_novamente@gmail.com");
        //responsável por fazer alteração dos dados já existente no banco
        em.merge(usuario);
        em.getTransaction().commit();

        em.close();
        emF.close();
    }
}
