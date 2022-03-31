
import net.datastructures.ListaSimplesmenteEncadeada;
import org.w3c.dom.ls.LSInput;

import java.util.List;


public class TesteListaSimplesmenteEncadeada {

    private void testarAdicionarNode(ListaSimplesmenteEncadeada l){
        l.addNewNode("primeiro item ");
        l.addNewNode("Segundo item ");
        l.addNewNode("Terceiro item ");
        l.addNewNode("Quarto item ");
        l.addNewNode("Quinto item ");
        l.addNewNode("Sexto item ");
        l.addNewNode("Sétimo item ");


        System.out.println("Lista após as inserções" );
        System.out.println("List: "+ l);

    }

    private void testarRemover1Elemento(ListaSimplesmenteEncadeada  l){
        System.out.println("+--------------------------------------+");
        System.out.println("Removendo um ítem da lista" );
        l.removeNode();
        System.out.println("List: "+ l);
    }

    private void testarRemoverTodosElementos( ListaSimplesmenteEncadeada l){
        System.out.println("Removendo todos os itens" );
        l.removeAll();
        System.out.println("List: "+ l);
    }



    public void execute(){
        ListaSimplesmenteEncadeada l = new ListaSimplesmenteEncadeada();
        testarAdicionarNode(l);

        testarRemover1Elemento(l);

        testarRemoverTodosElementos(l);
    }

    public static void main(String[] args){
        TesteListaSimplesmenteEncadeada T = new TesteListaSimplesmenteEncadeada();
        T.execute();


    }


}


