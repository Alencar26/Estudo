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


            Console.WriteLine("\n========== PROPRIEDADES E OPERAÇÕES ===========");
            
            // ===================== PROPRIEDADES E OPERAÇÕES COM DATETIME ========================

            // contrutor (ano, mês, dia, horas, minutos, segundos, milisegundos)
            DateTime data = new DateTime(2021, 2, 26, 08, 38, 56, 255);


            Console.WriteLine(data);
            Console.WriteLine($"Apenas a data: {data.Date}");
            Console.WriteLine($"Dia do mês: {data.Day}");
            Console.WriteLine($"Dia da semana {data.DayOfWeek}");
            Console.WriteLine($"Dia do ano: {data.DayOfYear}");
            Console.WriteLine($"horas: {data.Hour}");
            Console.WriteLine($"tipo(local ou universal): {data.Kind}");
            Console.WriteLine($"Milisegundos: {data.Millisecond}");
            Console.WriteLine($"Minutos: {data.Minute}");
            Console.WriteLine($"segundos: {data.Second}");
            Console.WriteLine($"Mês: {data.Month}");
            Console.WriteLine($"Segundo: {data.Second}");
            Console.WriteLine($"Ticks: {data.Ticks}");
            Console.WriteLine($"Horário do dia: {data.TimeOfDay}");
            Console.WriteLine($"Ano: {data.Year}");
            
            //==================== FORMATAÇÃO =================================

            Console.WriteLine("\n========================= FORMATAÇÃO ======================");
            
            // converte o tipo DateTime para string. Pode-se guardar em uma variável string normalmente.
            Console.WriteLine($"Data por estenso: {data.ToLongDateString()}");

            string estenso = data.ToLongDateString();
            Console.WriteLine($"Data por estenso (string): {estenso})");
            
            string s1 = data.ToLongTimeString();
            Console.WriteLine($"Apenas a hora em string: {s1}");
            
            string s2 = data.ToShortDateString();
            Console.WriteLine($"Apenas a data convertida para string: {s2}");

            string s3 = data.ToShortTimeString();
            Console.WriteLine($"Apenas a hora convertida para string de forma simplificada: {s3}");
            
            string s4 = data.ToString();
            Console.WriteLine($"Só um To.String: {s4}");
            
            
            Console.WriteLine("\n===================== MÁSCARA DE FORMATAÇÃO ===================");
            
            string mascara = data.ToString("yyyy-MM-dd HH:mm:ss");
            Console.WriteLine($"Com máscara de formatação: {mascara}");
            
            string mascaraMilesegundos = data.ToString("yyyy-MM-dd HH:mm:ss.fff");
            Console.WriteLine($"Com milisegundos: {mascaraMilesegundos}");
            
            //=========================== OPERAÇÕES COM DATETIME =======================

            Console.WriteLine("\n======================= OPERAÇÕES COM DATETIME ====================");
            
            // adicionando 2 horas na data que tinha
            DateTime d2 = data.AddHours(2);
            DateTime d3 = data.AddMinutes(30);
            DateTime d4 = data.AddDays(7);
            //... Pode-se fazer isso com qualquer tipos de DataTime (mês, ano, Ticks, milisegundos...)

            Console.WriteLine($"Sem alteração: {data}");
            Console.WriteLine($"+ 2 horas: {d2}");
            Console.WriteLine($"+ 30 minutos: {d3}");
            Console.WriteLine($"+ 7 dias: {d4}");

            Console.WriteLine("Saber a diferença de uma data para outra:\n");
            
            DateTime data2 = new DateTime(2021, 5, 3, 08, 38, 56, 255);
            
            TimeSpan diferença = (data2.Subtract(data));
            Console.WriteLine($"Diferença: {diferença}");
            
            
            

            Console.WriteLine("\n");
            

        }
    }
}
