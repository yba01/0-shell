# Variables
GO = go
SRC_DIR = src
MAIN_FILE = main.go
BINARY_NAME = 0-shell

# Cible par défaut : construire le binaire
all: build

# Cible pour compiler le projet
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_FILE)

# Cible pour exécuter le projet
run: build
	./$(BINARY_NAME)


# Cible pour nettoyer les fichiers compilés
clean:
	rm -f $(BINARY_NAME)

# Cible pour vérifier le formatage du code (go fmt)
fmt:
	$(GO) fmt $(SRC_DIR)

# Cible pour vérifier les dépendances (go mod tidy)
tidy:
	$(GO) mod tidy

# Cible pour afficher les informations du projet (version Go, etc.)
info:
	$(GO) version
	$(GO) env
