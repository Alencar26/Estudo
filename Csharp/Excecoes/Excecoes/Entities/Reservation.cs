using System;
using Excecoes.Entities.Exceptions;

namespace Excecoes.Entities
{
    class Reservation
    {
        public int RoomNumber { get; set; }
        public DateTime Checkin { get; set; }
        public DateTime Checkout { get; set; }

        public Reservation()
        {
        }

        public Reservation(int roomNumber, DateTime checkin, DateTime checkout)
        {
            if (checkout <= checkin)
            {
                throw new DomainException("A data de checkout é anterior à data de checkin");
            }

            RoomNumber = roomNumber;
            Checkin = checkin;
            Checkout = checkout;
        }

        public int Duration()
        {
            return (int) (Checkout.Day - Checkin.Day);
        }

        public void UpdateDates(DateTime checkin, DateTime checkout)
        {
            DateTime now = DateTime.Now;
            if (checkin < now || checkout < now)
            {
                throw new DomainException("As datas de atualização não podem ser anterior à data de agora.");
            }
            if (checkout <= checkin)
            {
                throw new DomainException("A data de checkout é anterior à data de checkin");
            }

            Checkin = checkin;
            Checkout = checkout;
        }

        public override string ToString()
        {
            return $"ROOM {RoomNumber}, check-in: {Checkin.ToString("dd/MM/yyyy")}," +
                   $" check-out: {Checkout.ToString("dd/MM/yyyy")}, {Duration()} nights";
        }
    }
}
