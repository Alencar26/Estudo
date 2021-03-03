using Abstrata.Entities.Enums;
using System;
using System.Collections.Generic;
using System.Text;

namespace Abstrata.Entities
{
    public class Ractangle : Shape
    {
        public double Width { get; set; }
        public double Height { get; set; }

        public Ractangle(Color color, double width, double height)
            : base(color)
        {
            Width = width;
            Height = height;
        }

        public override double Area()
        {
            return Width * Height;
        }
    }
}
