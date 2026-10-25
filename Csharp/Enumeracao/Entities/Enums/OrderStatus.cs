using System;

namespace Enumeracao.Entities.Enums
{
    public enum OrderStatus : int
    {
        PendingPayment,
        Processing,
        Shipped,
        Delivered
    }
}