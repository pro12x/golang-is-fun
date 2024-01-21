```Csharp
using System;
using System.Collections;

namespace Hello
{
    interface IEmpruntable
    {
        bool EstEmpruntable();
    }

    class Livre : IEmpruntable
    {
        public string Titre { get; set; }
        public string Auteur { get; set; }
        public string ISBN { get; set; }
        public int ExemplaireDisponible { get; set; }

        public bool EstEmpruntable()
        {
            return ExemplaireDisponible > 0;
        }
    }

    class Auteur
    {
        public string Nom { get; set; }
        public string Biographie { get; set; }
        public List<Livre> LivresEcrits { get; set; } = new List<Livre>();
    }

    class Emprunt
    {
        public DateTime DateEmprunt { get; set; }
        public DateTime DateRetourPrevue { get; set; }
        public Livre LivresEmprunte { get; set; }
    }

    class Bibliotheque
    {
        public List<Livre> Livres { get; set; } = new List<Livre>();
        public List<Emprunt> EmpruntEncours { get; set; } = new List<Emprunt>();

        public void AjouterLivre(Livre livre)
        {
            Livres.Add(livre);
        }

        public void EmprunterLivre(Livre livre, DateTime dateEmprunt)
        {
            if (livre.EstEmpruntable())
            {
                livre.ExemplaireDisponible--;
                EmpruntEncours.Add(new Emprunt
                {
                    LivresEmprunte = livre,
                    DateEmprunt = dateEmprunt,
                    DateRetourPrevue = dateEmprunt.AddDays(7) // Emprunt pour 7 jours
                });
            }
            else
            {
                Console.WriteLine("Le livre n'est pas disponible pour l'emprunt");
            }
        }

        public void RetournerLivre(Livre livre)
        {
            Emprunt emprunt = EmpruntEncours.FirstOrDefault(e => e.LivresEmprunte == livre);
            if (emprunt != null)
            {
                livre.ExemplaireDisponible++;
                EmpruntEncours.Remove(emprunt);
            }
            else
            {
                Console.WriteLine("Le livre ne nous appartient pas");
            }
        }

        public void LivresDisponible()
        {
            Console.WriteLine("Liste des livres disponibles");
            foreach (Livre livre in Livres.Where(e => e.EstEmpruntable()))
            {
                Console.WriteLine($"- (livre.Titre) écrit par (livre.Auteur)");
            }
        }

        public void LivresEmpruntEncours()
        {
            Console.WriteLine("Emprunts en cours");
            foreach (Emprunt emprunt in EmpruntEncours)
            {
                Console.WriteLine($"- (emprunt.LivreEmprunte.Titre) emprunté par (emprunt.DateEmprunt.ToShortDateString()) à rendre avant le (emprunt.DateRetourPrevue.ToShortDateString())");
            }
        }
    }

    class Program
    {
        static void Menu()
        {
            Console.WriteLine("============== MENU ==============");
            Console.WriteLine("1. Ajouter Un Livre");
            Console.WriteLine("2. Emprunter Un Livre");
            Console.WriteLine("3. Retourner Un Livre");
            Console.WriteLine("4. Lister Les Livres Disponibles");
            Console.WriteLine("5. Lister Les Livres En Cours d'Emprunt");
            Console.WriteLine("0. Quitter\n");
        }

        static void Main(string[] args)
        {
            Menu();
            bool response = true;
            while (response)
            {
                Console.Write("Choix: ");
                string input = Console.ReadLine();
                int choix;
                if (!(int.TryParse(input, out choix)))
                {
                    choix = -1;
                    //Console.WriteLine("Your value: " + choix);
                }

                switch (choix)
                {
                    case 0:
                        Console.WriteLine("Vous Avez Quitte Le Programme\n\n");
                        response = false;
                        break;
                    case 1:
                        Console.WriteLine("Ajouter Un Livre");
                        break;
                    case 2:
                        Console.WriteLine("Emprunter Un Livre");
                        break;
                    case 3:
                        Console.WriteLine("Retourner Un Livre");
                        break;
                    case 4:
                        Console.WriteLine("Lister Les Livres Disponibles");
                        break;
                    case 5:
                        Console.WriteLine("Lister Les Emprunts En Cours");
                        break;
                    default:
                        Console.WriteLine("Mauvais Choix, Reessayez !");
                        break;
                }
            }
        }
    }
}
```

### Author
[Franchis Janel MOKOMBA](https://github.com/pro12x)
