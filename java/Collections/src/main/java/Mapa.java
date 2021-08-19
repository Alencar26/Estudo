import java.util.HashMap;
import java.util.Map;

public class Mapa {

    public static void main(String[] args) {

        Map<Integer, String> usuarios = new HashMap<>();

        //ADD ou SUBSTITUIR
        usuarios.put(1,"Andre");
        usuarios.put(2, "Ana");
        usuarios.put(3, "Helo");

        System.out.println(usuarios);

        System.out.println(usuarios.get(3));
        System.out.println(usuarios.keySet());
        System.out.println(usuarios.values());
        System.out.println(usuarios.entrySet());

        System.out.println(usuarios.containsKey(2));
        System.out.println(usuarios.containsValue("Ana"));

        //usando a chave como referência
        for (int chave : usuarios.keySet()) {
            System.out.println(chave);
        }

        // usando valor de referência
        for (String valor : usuarios.values()) {
            System.out.println(valor);
        }

        // usando registro de referência chave e valor
        for (Map.Entry<Integer, String> registro : usuarios.entrySet()) {
            System.out.println(registro.getKey());
            System.out.println(registro.getValue());
        }
    }
}
