package lambdas;

import java.sql.Array;
import java.util.Arrays;
import java.util.List;
import java.util.function.Supplier;

public class Fornecedor {

    public static void main(String[] args) {

        Supplier<List<String>> nameList = () -> Arrays.asList("Lia", "Bia", "Gui", "Mat");

        System.out.println(nameList.get());
    }
}
