package aniversarioSurpresa;

public class FestaSurpresa {

    public static void main(String[] args) {

        Namorada namorada = new Namorada();
        Porteiro porteiro = new Porteiro();

        porteiro.registrarObservadores(namorada);
        porteiro.monitorar();
    }
}
