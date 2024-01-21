```Csharp
//using System;
// using System.Collections.Generic;
// using System.Linq;

namespace Biblio
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
            /*else
            {
                Console.WriteLine("Le livre ne nous appartient pas");
            }*/
        }

        public void LivresDisponible()
        {
            int id = 0;
            foreach (Livre livre in Livres.Where(e => e.EstEmpruntable()))
            {
                id++;
                Console.WriteLine(id + ":\n" + "\tTitre: " + livre.Titre + "\n\tAuteur " + livre.Auteur + "\n\tDisponibilité: [" + livre.ExemplaireDisponible + "] Exemplaires\n");
            }

            if (id == 0)
            {
                Console.WriteLine("Il n'y a aucun livre dans la bibliotheque\n");
            }
        }

        public void LivresEmpruntEncours()
        {
            int id = 0;
            foreach (Emprunt emprunt in EmpruntEncours)
            {
                id++;
                Console.WriteLine(id + ":\n" + emprunt.LivresEmprunte.Titre + " - emprunté le " + emprunt.DateEmprunt.ToShortDateString() + " à rendre avant le " + emprunt.DateRetourPrevue.ToShortDateString() + "\n");
            }

            if (id == 0)
            {
                Console.WriteLine("Aucun livre n'a été emprunté\n");
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

        // Opération d'ajout
        static void Ajouter(Bibliotheque bibliotheque)
        {
            Console.Write("Titre: ");
            string titre = Console.ReadLine();

            Console.Write("Auteur: ");
            string auteur = Console.ReadLine();

            Console.Write("ISBN: ");
            string isbn = Console.ReadLine();

            Console.Write("Nombre d'exemplaires: ");
            int exemplaires = Convert.ToInt32(Console.ReadLine());

            Livre nouveauLivre = new Livre { Titre = titre, Auteur = auteur, ISBN = isbn, ExemplaireDisponible = exemplaires };
            bibliotheque.AjouterLivre(nouveauLivre);

            Console.WriteLine("Livre ajouté avec succès!\n");
        }

        // Opération d'emprunt
        static void Emprunter(Bibliotheque bibliotheque)
        {
            Console.Write("Titre: ");
            string titre = Console.ReadLine();

            Livre livre = bibliotheque.Livres.FirstOrDefault(l => l.Titre.Equals(titre, StringComparison.OrdinalIgnoreCase));

            if (livre != null)
            {
                bibliotheque.EmprunterLivre(livre, DateTime.Now);
                Console.WriteLine("Livre emprunté avec succès!\n");
            }
            else
            {
                Console.WriteLine("Le livre n'existe pas.\n");
            }
        }

        // Opération de retour
        static void Retourner(Bibliotheque bibliotheque)
        {
            Console.Write("Titre: ");
            string titre = Console.ReadLine();

            Livre livre = bibliotheque.Livres.FirstOrDefault(l => l.Titre.Equals(titre, StringComparison.OrdinalIgnoreCase));

            if (livre != null)
            {
                bibliotheque.RetournerLivre(livre);
                Console.WriteLine("Livre retourné avec succès!\n");
            }
            else
            {
                Console.WriteLine("Ce n'est pas notre livre\n");
            }
        }

        static void Main(string[] args)
        {
            bool response = true;

            // Instanciation d'un élément de type Bibliotheque
            Bibliotheque bibliotheque = new Bibliotheque();
            while (response)
            {
                // Menu
                Menu();
                Console.Write("Choix: ");

                // Attendre la valeur de l'utilisateur
                string input = Console.ReadLine();
                int choix;

                // Vérifier si la valeur entrée par l'utilisateur est un entier
                if (!(int.TryParse(input, out choix)))
                {
                    choix = -1;
                    //Console.WriteLine("Your value: " + choix);
                }

                // Système d'échange des valeurs entrées par l'utilisateur
                switch (choix)
                {
                    case 0:
                        Console.WriteLine("Vous Avez Quitte Le Programme, A Bientot !!!\n");
                        response = false;
                        break;
                    case 1:
                        Console.WriteLine("============ Ajouter Un Livre ============\n");
                        Ajouter(bibliotheque);
                        break;
                    case 2:
                        Console.WriteLine("============ Emprunter Un Livre ============\n");
                        Emprunter(bibliotheque);
                        break;
                    case 3:
                        Console.WriteLine("============ Retourner Un Livre ============\n");
                        Retourner(bibliotheque);
                        break;
                    case 4:
                        Console.WriteLine("============ Lister Les Livres Disponibles ============\n");
                        bibliotheque.LivresDisponible();
                        break;
                    case 5:
                        Console.WriteLine("============ Lister Les Emprunts En Cours ============\n");
                        bibliotheque.LivresEmpruntEncours();
                        break;
                    default:
                        Console.WriteLine("Mauvais Choix, Reessayez !\n");
                        break;
                }
            }
        }
    }
}
```

### Author
[Franchis Janel MOKOMBA](https://github.com/pro12x)
