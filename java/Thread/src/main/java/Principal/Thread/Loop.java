package Principal.Thread;
import Principal.Interface.IProcessoPrincipal;

public class Loop implements Runnable {

    private IProcessoPrincipal processoPrincipoal;

    public Loop(IProcessoPrincipal processoPrincipal) {
        this.processoPrincipoal = processoPrincipal;
    }

    @Override
    public void run() {
        //new Thread().sleep(5000);

        for (int x = 0; x < 10; x++) {
            this.processoPrincipoal.mostrar(Thread.currentThread().getName(), x);
        }
    }
}
