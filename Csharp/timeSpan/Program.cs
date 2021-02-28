using System;

namespace timeSpan
{
    class Program
    {
        static void Main(string[] args)
        {
            // construtor passando por parâmetro (hora, minuto, segundo)
            TimeSpan t = new TimeSpan(0, 1, 30);

            // construtor vazio.
            TimeSpan t2 = new TimeSpan();
           
           // contrutor passando por parâmetro a qnt de 'ticks' (nanosegundos)
           TimeSpan t3 = new TimeSpan(900000000L);
           
            // contrutor passando por parâmetro (dia, hora, minuto, segundo)
            TimeSpan t4 = new TimeSpan(1, 2, 25, 58);

            // contrutor passando por parâmetro (dia, hora, minuto, segundo, milisegundos)
            TimeSpan t5 = new TimeSpan(1, 2, 25, 58, 321);


            //=================== MÉTODO FROM ==================================

            // retorna um TimeSpan de um dia e meio
            TimeSpan t6 = TimeSpan.FromDays(1.5);

            // retorna um timeSpan de uma hora e meia
            TimeSpan t7 = TimeSpan.FromHours(1.5);

            // retorna um TimeSpan de um minuto e meio
            TimeSpan t8 = TimeSpan.FromMinutes(1.5);
            // outros...
            TimeSpan t9 = TimeSpan.FromSeconds(1.5);
            TimeSpan t10 = TimeSpan.FromMilliseconds(1.5);
            TimeSpan t11 = TimeSpan.FromTicks(90032554L);

           //===================== PRINTS ==================================
           

            Console.WriteLine(t);   
            // representação da hora em nanosegundos
            Console.WriteLine(t.Ticks);
            Console.WriteLine(t2);
            Console.WriteLine(t3);
            Console.WriteLine(t4);
            Console.WriteLine(t5);
            Console.WriteLine(t6);
            Console.WriteLine(t7);
            Console.WriteLine(t8);
            Console.WriteLine(t9);
            Console.WriteLine(t10);
            Console.WriteLine(t11);
            
            //==================== PROPRIEDADES E OPERAÇÕES =======================

            Console.WriteLine("================= PROPRIEDADE E OPERAÇÕES =====================");
            

            TimeSpan t12 = TimeSpan.MaxValue;
            TimeSpan t13 = TimeSpan.MinValue;
            TimeSpan zero = TimeSpan.Zero;

            Console.WriteLine(t12);
            Console.WriteLine(t13);
            Console.WriteLine(zero);
            Console.WriteLine($"\n=======================\n");
            
            //usando a variável t5
            Console.WriteLine($"Completo: {t5}");

            // devolvem números interiros:
            Console.WriteLine($"Dias: {t5.Days}");
            Console.WriteLine($"Horas: {t5.Hours}");
            Console.WriteLine($"Minutos: {t5.Minutes}");
            Console.WriteLine($"Milisegundos: {t5.Milliseconds}");
            Console.WriteLine($"Segundos: {t5.Seconds}");
            Console.WriteLine($"Ticks: {t5.Ticks}");
            // devolvem números quebrados:
            Console.WriteLine($"Total de Dias: {t5.TotalDays}");
            Console.WriteLine($"Total de Horas: {t5.TotalHours}");
            Console.WriteLine($"Total de Minutos: {t5.TotalMinutes}");
            Console.WriteLine($"Total de Segundos: {t5.TotalSeconds}");
            Console.WriteLine($"Total de Milisegundos: {t5.Milliseconds}");
            
            
            Console.WriteLine("================= OPERAÕES ====================");

            // construtor (hora, minuto, segundo)
            TimeSpan tempo1 = new TimeSpan(1,30,10);
            TimeSpan tempo2 = new TimeSpan(0,10,5);

            Console.WriteLine($"\nTempo 1: {tempo1}");
            Console.WriteLine($"Tempo 2: {tempo2}");
            Console.WriteLine("\n===========================\n");
            
            TimeSpan soma = tempo1.Add(tempo2);
            TimeSpan diferenca = tempo1.Subtract(tempo2);
            // recebe um double como argumento
            TimeSpan multiplicacao = tempo2.Multiply(2.0);
            TimeSpan divisao = tempo2.Divide(2.0);

            Console.WriteLine($"Soma: {soma}");
            Console.WriteLine($"Difereça: {diferenca}");
            Console.WriteLine($"Multiplicação: {multiplicacao}");
            Console.WriteLine($"Divisão: {divisao}");
            
            
            
            
            
            
            
            
            
            
        }
    }
}
