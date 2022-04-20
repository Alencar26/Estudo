package generics.resolvendoTipoNaHeranca;

// É mais comum usar apenas uma letra no tipo genérico
public class Caixa<TIPO> {

    private TIPO coisa;

    public void guardar(TIPO coisa) {
        this.coisa = coisa;
    }

    public TIPO abrir() {
        return this.coisa;
    }
}
