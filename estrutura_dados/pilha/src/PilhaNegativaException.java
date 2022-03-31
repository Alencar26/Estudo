public class PilhaNegativaException extends Exception{

    public  PilhaNegativaException(String msg) {
        super(msg);
    }

    public PilhaNegativaException(String msg, Throwable causa) {
        super(msg, causa);
    }
}