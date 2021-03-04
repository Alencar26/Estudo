using Excecoes.Entities;
using Excecoes.Entities.Exceptions;
using System;

namespace Excecoes
{
    class Program
    {
        static void Main(string[] args)
        {
            try
            {
                Console.Write("Romm Number: ");
                int rN = int.Parse(Console.ReadLine());

                Console.Write("Check-in date (dd/MM/yyyy): ");
                DateTime cI = DateTime.Parse(Console.ReadLine());

                Console.Write("Check-out date (dd/MM/yyyy): ");
                DateTime cO = DateTime.Parse(Console.ReadLine());

                Reservation reservation = new Reservation(rN, cI, cO);
                Console.WriteLine(reservation);

                Console.WriteLine("Entre com os dados para atualizar a reserva: ");
                Console.Write("Check-in date (dd/MM/yyyy): ");
                cI = DateTime.Parse(Console.ReadLine());
                Console.Write("Check-out date (dd/MM/yyyy): ");
                cO = DateTime.Parse(Console.ReadLine());

                reservation.UpdateDates(cI, cO);
                Console.WriteLine(reservation);
            }
            catch (DomainException e)
            {
                Console.WriteLine($"Erro na reserva: {e.Message}");
            }
        }
    }
}
