import java.util.Arrays;
import java.util.List;
import java.util.function.BinaryOperator;

public class ReduceExemplo1 {

    public static void main(String[] args) {

        List<Integer> nums = Arrays.asList(2, 6, 7, 14, 35);

        //Exemplos passando BinaryOparator direto no parâmetro
        Integer result = nums.stream().reduce(Integer::sum).get();
        Integer result2 = nums.stream().reduce((a, b) -> a + b).get();

        System.out.println(result);
        System.out.println(result2);

        //Separando as operações
        BinaryOperator<Integer> soma = Integer::sum;
        BinaryOperator<Integer> soma2 = (a, b) -> a + b;

        // se passar um valor antes como parâmetro ele considera esse valor na hora de somar
        Integer result3 = nums.stream().reduce(0 ,soma);
        Integer result4 = nums.stream().reduce(soma2).get();

        System.out.println(result3);
        System.out.println(result4);

        // Usando filter junto
        nums.stream()
                .filter(n -> n > 7)
                .reduce(soma)
                .ifPresent(System.out::println); // se o resultado estiver presente então faça isso...
    }
}
