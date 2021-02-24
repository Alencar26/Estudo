using System;

namespace Modificadores_parametro_ref_e_out
{
    class Program
    {
        static void Main(string[] args)
        {
            int num = 10;
            Calculator.Multiplica(ref num);
            Console.WriteLine(num);

            //=================================

            // modificador out, não exige que a variável seja iniciada.

            int entrada = 10;
            int resultado;

            // o método abaixo vai retornar o resultado por meio de parâmetro para a variável "resultado"
            Calculator.Triplica(entrada, out resultado);
            Console.WriteLine(resultado);
        }
    }
}
