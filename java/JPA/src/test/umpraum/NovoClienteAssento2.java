package test.umpraum;

import infra.DAO;
import model.umpraum.Assento;
import model.umpraum.Cliente;

public class NovoClienteAssento2 {

    public static void main(String[] args) {

        Assento assento = new Assento("4D");
        Cliente cliente = new Cliente("André", assento);

        DAO<Cliente> dao = new DAO<>(Cliente.class);
        /*
           Será possível executar a operação abaixo por causa da anotation
           OneToOne na classe Cliente com atributo "cascade" que fará um
           operação em cascata. Ou seja, o JPA vai entender que o assento
           não está registrado no banco e vai efetuar o cadastro do assento
           juntamente com o cadastro do cliente no banco de dados.
        */
        dao.incluirAtomico(cliente);
    }
}
