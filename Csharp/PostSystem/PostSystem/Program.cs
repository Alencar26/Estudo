using PostSystem.Entities;
using System;

namespace PostSystem
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Adicionar post: ");
            DateTime data = DateTime.Now;
            Console.Write("Título: ");
            string titulo = Console.ReadLine();
            Console.Write("Conteúdo: ");
            string conteudo = Console.ReadLine();
            int likes = 10;

            Post post1 = new Post(data, titulo, conteudo, likes);

            Comment com1 = new Comment("Comentário 1");
            Comment com2 = new Comment("Comentário 2");

            post1.AddComment(com1);
            post1.AddComment(com2);

            Console.WriteLine(post1);
        }
    }
}
