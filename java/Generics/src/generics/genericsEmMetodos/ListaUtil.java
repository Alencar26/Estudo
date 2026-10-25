package generics.genericsEmMetodos;

import java.util.List;

public class ListaUtil {

    // versão 1.0
    public static Object getUltimo1(List<?> lista) {
        return lista.get(lista.size() - 1);
    }

    // versão 2.0
    public static <T> T getUltimo2(List<T> lista) {
        return lista.get(lista.size() - 1);
    }
}

