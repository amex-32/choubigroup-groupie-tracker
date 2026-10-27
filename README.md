# 🧩 Pokédex Interactif — Projet API Pokémon

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTML](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![API](https://img.shields.io/badge/API-Pokémon-blue?style=for-the-badge)

Ce projet est un **Pokédex interactif** développé autour de l’exploitation d’une **API Pokémon**.  
Il permet de consulter, filtrer et explorer l’ensemble des Pokémon à travers une interface web moderne, tout en manipulant des données complexes côté backend.

---

## 🎯 Objectif du projet

L’objectif principal est de :

- Comprendre et exploiter une **API externe**
- Structurer et manipuler des données complexes en backend
- Générer dynamiquement des pages HTML à partir des données reçues
- Concevoir une interface claire et interactive
- Mettre en place une navigation cohérente entre plusieurs pages

Le projet met l’accent sur la **communication entre le backend Go et le frontend HTML/CSS**.

---

## 🧠 Principe général

Le site agit comme une **interface de consultation de données Pokémon** :

1. Les données sont récupérées depuis une **API Pokémon**
2. Elles sont transformées et stockées dans des structures Go
3. Les pages HTML sont rendues dynamiquement via des templates
4. L’utilisateur interagit avec les données via la navigation, la recherche et les filtres

---

## 📄 Pages principales

Le site est composé de plusieurs pages interconnectées :

- **Accueil**
  - Présentation du projet
  - Affichage d’un Pokémon aléatoire
  - Introduction à l’univers Pokémon

- **Collection**
  - Liste complète des Pokémon
  - Système de recherche par nom
  - Filtres avancés (types, générations, formes spéciales)
  - Cartes interactives avec détails extensibles

- **Ressources**
  - Consultation structurée des Pokémon
  - Accès aux fiches détaillées

- **Détail Pokémon**
  - Informations complètes d’un Pokémon :
    - Types
    - Stats
    - Talents
    - Résistances
    - Évolutions
    - Formes spéciales
    - Sprites (regular, shiny, gigamax)

- **Favoris**
  - Gestion d’une liste de Pokémon favoris

- **Compte**
  - Connexion et inscription utilisateur (interface)

---

## 🔍 Fonctionnalités clés

- 🔎 **Recherche dynamique**
  - Recherche par nom de Pokémon

- 🧪 **Filtres avancés**
  - Types
  - Générations (1 à 9)
  - Formes spéciales (Méga, Gigamax, formes régionales)

- 🖼️ **Affichage détaillé**
  - Sprites multiples
  - Stats complètes
  - Arbre d’évolution
  - Informations biologiques

- ❤️ **Système de favoris**
  - Ajout / suppression depuis les fiches Pokémon

- 🎨 **UI moderne**
  - Design Glassmorphism
  - Overlays (recherche, favoris, authentification)
  - Composants interactifs (`details/summary`)

---

## 🏗️ Architecture technique

### Backend
- Langage : **Go**
- Rôle :
  - Appels à l’API Pokémon
  - Traitement et structuration des données
  - Gestion des routes
  - Rendu des templates HTML

Les données sont mappées dans des structures Go dédiées (Pokémon, stats, évolutions, formes, etc.).

### Frontend
- **HTML** (templates dynamiques)
- **CSS** (design moderne et responsive)
- Aucune dépendance JavaScript lourde
- Interactions gérées via HTML/CSS (checkbox, overlays)

---

## 🧬 Modélisation des données

Les données Pokémon sont organisées autour de structures représentant :

- Identité (nom, ID, génération)
- Sprites (regular, shiny, gigamax)
- Types et talents
- Statistiques
- Résistances
- Évolutions
- Formes régionales et spéciales
- Données biologiques (taille, poids, sexe, groupes d’œufs)

Cette modélisation permet une exploitation claire et évolutive des données issues de l’API.

---

## 🚧 État du projet

- Projet fonctionnel
- Développé dans un cadre pédagogique
- Axé sur la compréhension des API et du rendu dynamique
- Architecture évolutive

---

## 📝 Licence

Projet réalisé à des fins pédagogiques.  
Licence libre.
