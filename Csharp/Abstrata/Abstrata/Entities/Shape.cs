using Abstrata.Entities.Enums;
using System;
using System.Collections.Generic;
using System.Text;

namespace Abstrata.Entities
{
    public abstract class Shape
    {
        public Color Color { get; set; }

        public Shape() { }

        public Shape(Color color)
        {
            Color = color;
        }

        public abstract double Area();
    }
}
