package io.al3ncar;

/**
 * Singleton "Apressado".
 * @author Alencar26
 **/

public class SingletonEager {

    private static SingletonEager instancia = new SingletonEager();

    private SingletonEager() {
        super();
    }

    public static SingletonEager getInstancia() {
        return instancia;
    }
}
