using System;
using System.Globalization;

namespace DataComDateTime
{
    class Program
    {
        static void Main(string[] args)
        {
            // instante atual da máquina
            DateTime d = DateTime.Now;

            //instanciando uma data
            DateTime instancia = new DateTime(2021, 2, 25);
            DateTime instanciaComHoras = new DateTime(2021, 2, 25, 21, 39, 03);

            // data de hoje com horário zerado
            DateTime dataSemHoras = DateTime.Today;

            // horário universal 
            DateTime dataComHorarioGlobal = DateTime.UtcNow;

            //converter string para dateTime
            DateTime dataString = DateTime.Parse("2021-02-25");
            DateTime dataString2 = DateTime.Parse("2021/02/25");

            // convertendo string para DateTime com máscara
            DateTime data1 = DateTime.ParseExact("2021-02-25", "yyyy-MM-dd", CultureInfo.InvariantCulture);


            Console.WriteLine("atual: " + d);
            System.Console.WriteLine("instancia: " + instancia);
            System.Console.WriteLine("Data sem horas: " + dataSemHoras);
            System.Console.WriteLine("Data string: " + dataString);
            Console.WriteLine(data1);
                        
            
        }
    }
}
