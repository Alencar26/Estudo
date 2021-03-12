using System;
using System.Collections.Generic;
using System.IO;
using IComparableExample.Entities;

namespace IComparableExample
{
    class Program
    {
        static void Main(string[] args)
        {
            // padrão da linguagem para comparações de objetos.
            // verifica se o obj é > ou < ou == à outro obj.

            string path = @"C:\Users\andre\www\ws_csarp\Exercicios_POO\Csharp\IComparableExample\IComparableExample\file.txt";

            try
            {
                using (StreamReader sr = File.OpenText(path))
                {
                    List<Employee> list = new List<Employee>();
                    while (!sr.EndOfStream)
                    {
                        list.Add(new Employee(sr.ReadLine()));
                    }
                    // ordenar a lista
                    list.Sort();

                    foreach (Employee emp in list)
                    {
                        Console.WriteLine(emp);
                    }
                }
            }
            catch (IOException e)
            {
                Console.WriteLine("Houve um erro:");
                Console.WriteLine(e.Message);
            }

        }
    }
}
