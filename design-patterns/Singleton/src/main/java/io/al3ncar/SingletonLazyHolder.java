package io.al3ncar;

/**
 * Singleton "Lazy Holder".
 *
 * @see <a href="https://stackoverflow.com/a/24018148">Singleton Lazy Holder</a>
 *
 * @author Alencar26
 **/
public class SingletonLazyHolder {

    private static class Holder {
        public static SingletonLazyHolder instancia = new SingletonLazyHolder();
    }

    private SingletonLazyHolder() { super (); }

    public  static SingletonLazyHolder getInstancia() {
        return Holder.instancia;
    }
}
