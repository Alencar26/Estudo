package Principal;
import Principal.Interface.IProcessoPrincipal;
import Principal.Thread.Loop;

public class Principal implements IProcessoPrincipal {

    public static void main(String[] args) throws InterruptedException {
        new Principal().execute();
    }

    private void execute() throws InterruptedException {
        System.out.println("começou");

        Thread threadUm = new Thread( new Loop(this), "THREAD-UM");
        Thread threadDois = new Thread( new Loop(this), "THREAD-DOIS");

        // informa o estado da thread
        ThreadStatus(threadUm);
        ThreadStatus(threadDois);

        threadUm.start();
        threadDois.start();

        ThreadStatus(threadUm);
        ThreadStatus(threadDois);

        // Join suspende o processo principal até que a thread termine
        threadUm.join();

        ThreadStatus(threadUm);
        ThreadStatus(threadDois);
    }

    
    public void ThreadStatus(Thread thread){
        System.out.println(thread.getName() + " "+thread.getState());
    }

    @Override
    public void mostrar(String nome, int ciclo) {
        System.out.println(nome + " " + ciclo);
    }
}

