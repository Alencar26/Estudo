using System;
using InterfaceExample.Entities;

namespace InterfaceExample.Services
{
    class RentalService
    {
        public double PricePerHour { get; private set; }
        public double PricePerDay { get; private set; }

        private ITaxService _taxService;

        public RentalService(double pricePerHour, double pricePerDay, ITaxService taxService)
        {
            PricePerHour = pricePerHour;
            PricePerDay = pricePerDay;
            _taxService = taxService;
        }

        public void ProcessInvoice(CarRental carRental)
        {
            double basicPayment = DurationValue(carRental.Start, carRental.Finish);
            double tax = _taxService.Tax(basicPayment);

            carRental.Invoice = new Invoice(basicPayment, tax);
        }

        private double DurationValue(DateTime start, DateTime finish)
        {
            TimeSpan duration = finish.Subtract(start);
            if (duration.TotalHours > 12)
                return Math.Ceiling(duration.TotalDays) * PricePerDay;
            else
                return Math.Ceiling(duration.TotalHours) * PricePerHour;
        }

        //Nota: Math.Ceiling() << arredonda o valor para cima.
    }
}
