

public class TesteListaSimplesmenteEncadeada {

    private void testarAdicionarNode(ListSimplesmenteEncadeada l){
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

    private void testarRemover1Elemento(ListSimplesmenteEncadeada  l){
        System.out.println("+--------------------------------------+");
        System.out.println("Removendo um ítem da lista" );
        l.removeNode();
        System.out.println("List: "+ l);
    }

    private void testarRemoverTodosElementos( ListSimplesmenteEncadeada l){
        System.out.println("Removendo todos os itens" );
        l.removeAll();
        System.out.println("List: "+ l);
    }



    public void execute(){
        ListSimplesmenteEncadeada l = new ListSimplesmenteEncadeada();
        testarAdicionarNode(l);

        testarRemover1Elemento(l);

        testarRemoverTodosElementos(l);
    }

    public static void main(String[] args){
        TesteListaSimplesmenteEncadeada T = new TesteListaSimplesmenteEncadeada();
        T.execute();


    }
}
