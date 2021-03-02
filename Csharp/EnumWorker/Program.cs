using System;
using System.Globalization;
using EnumWorker.Entities;
using EnumWorker.Entities.Enums;

namespace EnumWorker
{
    class Program
    {
        static void Main(string[] args)
        {
             Console.Write("Entre com o nome do departamento: ");
             string nomeDepartamento = Console.ReadLine();
             Console.WriteLine("Entre com os dados do funcionário: ");
             Console.Write("Name: ");
             string nome = Console.ReadLine();
             Console.Write("Level (Junior, MidLevel, Senior): ");
             WorkerLevel level = Enum.Parse<WorkerLevel>(Console.ReadLine());
             Console.Write("Salário base: ");
             double salario = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
             
            Department dept = new Department(nomeDepartamento);
            Worker trabalhador = new Worker(nome, level, salario, dept);

            Console.WriteLine("Quantos contratos tem o trabalhador?");
            Console.Write("R: ");
            int numContratos = int.Parse(Console.ReadLine());

            for (var i = 1; i <= numContratos; i++)
            {
                Console.WriteLine($"Entre com os dados do contrato #{i}:");
                Console.Write("Data (DD/MM/YYYY): ");
                DateTime data = DateTime.Parse(Console.ReadLine());
                Console.Write("Valor por hora: ");
                double valorPorHora = double.Parse(Console.ReadLine());
                Console.Write("Duração: ");
                int duracao = int.Parse(Console.ReadLine());

                HourContract contrato = new HourContract(data, valorPorHora, duracao);
                trabalhador.AddContract(contrato);
            }
            
            Console.Write("Entre com o mês e o ano para calcular o ganho (MM/YYYY): ");
            string dataConsulta = Console.ReadLine();
            int mes = int.Parse(dataConsulta.Substring(0,2));
            int ano = int.Parse(dataConsulta.Substring(3));

            Console.WriteLine($"Nome: {trabalhador.Name}");
            Console.WriteLine($"Departamento: {trabalhador.Department.Name}");
            Console.WriteLine($"Ganhos para {dataConsulta}: {trabalhador.Income(ano, mes)}");
            
        }
    }
}
