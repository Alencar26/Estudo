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
            
                       
           
           
        }
    }
}
