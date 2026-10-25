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
            

            // =========================== PADRÃO DE FORMATAÇÃO DE DATA ==============================
            
            Console.WriteLine("=============== DATETIMEKIND E PADÃO ISO 8601 PARA FORMATAÇÃO ==============");
            

            /*

                SEMPRE armazene em formato UTC(no: banco de dados, json, xml...)
                somente quando for mostrar ao usuário é que você instancia com a data local

                usando: data.ToLocalTime(); ou data.ToUniversalTime();

            */ 
             
            // instanciando a data com horário local (ano, mês, dia, hora, minutos, segundos, formato)
            DateTime local = new DateTime(2021, 2, 28, 7, 38, 52, DateTimeKind.Local);
            
            // instanciando a data com horário UTC (ano, mês, dia, hora, minutos, segundos, formato)
            DateTime universal = new DateTime(2021, 2, 28, 7, 38, 52, DateTimeKind.Utc);
            // não especificado
            DateTime semEspecificacao = new DateTime(2021, 2, 28, 7, 38, 52, DateTimeKind.Unspecified);


            Console.WriteLine("Converções:");
            Console.WriteLine($"Data: {local}");
            Console.WriteLine($"Kind da data: {local.Kind}");
            Console.WriteLine($"To Local: {local.ToLocalTime()}");
            Console.WriteLine($"To UTC: {local.ToUniversalTime()}");
            
            
            //====================== PADRÃO ISO 8601 ==============================

            Console.WriteLine("\n================= PADRÃO ISO 8601 ======================\n");
            Console.WriteLine($"Padrão mais utilizado:  yyyy-MM-ddTHH:mm:ssZ ");
            Console.WriteLine("*Z indica que a data/hora está em UTC.");
            
            // sem especificação
            DateTime d1 = DateTime.Parse("2021-02-28 07:54:36");
            //usando padrão ISO 8601
            DateTime d6 = DateTime.Parse("2021-02-28T07:54:36Z");
            
            //converção segura. Primeiro garante que está no formato UTC e depois aplica a máscara.
            d6.ToUniversalTime().ToString("yyyy-MM-ddTHH:mm:ssZ");

            Console.WriteLine("\n");
            

        }
    }
}
