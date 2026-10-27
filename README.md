# Groupie Tracker : Pokédex

Application web en Go qui affiche les Pokémon de l'API [Tyradex](https://tyradex.app) : liste avec recherche et filtres, fiche détaillée (stats, évolutions, formes) et favoris.

Projet de groupe réalisé en formation.

## Lancer le projet

Go est nécessaire, ainsi qu'un accès internet pour appeler l'API.

```bash
cd src
go run .
```

Le site est ensuite disponible sur http://localhost:8080.

## Structure

- `src/main.go` : point d'entrée
- `src/router/` : routes
- `src/controller/` : handlers et appels à l'API
- `src/struct/` : structures de données
- `src/pages/` : templates HTML
- `src/static/` : CSS, images, icônes
