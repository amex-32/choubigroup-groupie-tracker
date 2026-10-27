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

## Hébergement

Le dossier `src` est déployé sur Vercel comme fonction Go (`src/api/index.go`, `src/vercel.json`), avec `src` comme dossier racine du projet.

## Structure

- `src/main.go` : point d'entrée
- `src/router/` : routes
- `src/controller/` : handlers et appels à l'API
- `src/struct/` : structures de données
- `src/pages/` : templates HTML
- `src/static/` : CSS, images, icônes
