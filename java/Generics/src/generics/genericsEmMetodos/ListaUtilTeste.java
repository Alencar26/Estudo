package generics.genericsEmMetodos;

import java.util.Arrays;
import java.util.List;

public class ListaUtilTeste {

    public static void main(String[] args) {

        List<String> langs = Arrays.asList("Java", "Go", "Rust", "C++");
        List<Integer> nums = Arrays.asList(2,3,5,9);

        // usando a versão 1.0 - PRECISA DO CASTING PARA FUNCIONAR

        String ultimaLinguagem = (String) ListaUtil.getUltimo1(langs);
        System.out.println(ultimaLinguagem);

        int ultimoNumero = (Integer) ListaUtil.getUltimo1(nums);
        System.out.println(ultimoNumero);

        //======================================================================

        // usando a versão 2.0 - NÃO PRECISA DO CASTING

        String ultimaLinguagem2 = ListaUtil.getUltimo2(langs);
        System.out.println(ultimaLinguagem);

        int ultimoNumero2 = ListaUtil.getUltimo2(nums);
        System.out.println(ultimoNumero);
    }
}
