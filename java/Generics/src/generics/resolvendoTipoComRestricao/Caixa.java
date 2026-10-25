package generics.resolvendoTipoComRestricao;

// É mais comum usar apenas uma letra no T genérico
public class Caixa<T> {

    private T coisa;

    public void guardar(T coisa) {
        this.coisa = coisa;
    }

    public T abrir() {
        return this.coisa;
    }
}
