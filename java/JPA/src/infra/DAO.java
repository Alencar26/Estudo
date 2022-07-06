package infra;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;
import javax.persistence.TypedQuery;
import java.util.List;

public class DAO<E> {

    private static EntityManagerFactory emF;
    final private EntityManager em;
    final private Class<E> classe;

    //bloco estático (só é inicializado uma vez quando a classe é carregada)
    static {
        try {
            emF = Persistence.createEntityManagerFactory("JPA");
        } catch (Exception e) {
            // log4j -> escolha de aplicação para geração de log
            //pode-se fazer um log do que aconteceu na aplicação
            System.out.println("Erro ao instanciar o Entity Manager Factory.\n Erro: " + e.getMessage());
        }
    }

    public DAO() {
        this(null);
    }

    public DAO(Class<E> classe) {
        this.classe = classe;
        em = emF.createEntityManager();
    }

    public DAO<E> abrirTransacao() {
        em.getTransaction().begin();
        return this;
    }

    public DAO<E> fecharTransacao() {
        em.getTransaction().commit();
        return this;
    }

    public DAO<E> incluir(E entidade) {
        em.persist(entidade);
        return this;
    }

    public DAO<E> incluirAtomico(E entidade) {
        return this.abrirTransacao().incluir(entidade).fecharTransacao();
    }

    public E obterPorId(Object id) {
        return em.find(classe, id);
    }

    public List<E> obterTodos() {
        return this.obterTodos(10,0);
    }

    public List<E> obterTodos(int qtde, int deslocamento) {
        if(classe == null) {
            throw new UnsupportedOperationException("Está classe é nula.");
        }

        String jpql = "select e from" + classe.getName() + " e";
        TypedQuery<E> query = em.createQuery(jpql, classe);
        query.setMaxResults(qtde);
        query.setFirstResult(deslocamento);
        return query.getResultList();
    }
    //Named Query
    // Consulta personalizada com base no arquivo consultas.xml
    public List<E> consultar(String nomeConsulta, Object... params) {
        TypedQuery<E> query = em.createNamedQuery(nomeConsulta, classe);
            for (int i = 0; i < params.length; i += 2) {
                    query.setParameter(params[i].toString(), params[i + 1].toString());
                }
            return query.getResultList();
        }

        //consulta com query nativa
    public E consultarUm(String nomeConsulta, Object... params) {
        List<E> lista = consultar(nomeConsulta, params);
        return lista.isEmpty() ? null : lista.get(0);
    }

        public void fechar(){
            em.close();
        }
    }
