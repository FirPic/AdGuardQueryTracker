#!/bin/bash

# Fonction pour afficher le menu principal
function show_main_menu() {
    echo "Choisissez un système d'exploitation :"
    echo "1. Windows"
    echo "2. Linux"
    echo "3. Quitter"
    echo -n "Entrez votre choix : "
    read main_choice
}

# Fonction pour afficher le menu d'architecture
function show_arch_menu() {
    echo "Choisissez une architecture :"
    echo "1. amd64"
    echo "2. arm64"
    echo "3. Retour au menu principal"
    echo -n "Entrez votre choix : "
    read arch_choice
}

# Boucle principale du menu
while true; do
    show_main_menu
    case $main_choice in
        1)
            selected_os="windows"
            ;;
        2)
            selected_os="linux"
            ;;
        3)
            echo "Au revoir !"
            exit 0
            ;;
        *)
            echo "Choix invalide, veuillez réessayer."
            continue
            ;;
    esac

    # Afficher le menu d'architecture si un système d'exploitation a été sélectionné
    while true; do
        show_arch_menu
        case $arch_choice in
            1)
                selected_arch="amd64"
                if [ "$selected_os" = "windows" ]; then
                    export GOARCH="$selected_arch"
                    export GOOS="$selected_os"
                    go build -o AdGuardQueryTracker.exe
                elif [ "$selected_os" = "linux" ]; then
                    make GOOS="$selected_os" GOARCH="$selected_arch" output=AdGuardQueryTracker
                fi
                break
                ;;
            2)
                selected_arch="arm64"
                if [ "$selected_os" = "windows" ]; then
                    export GOARCH="$selected_arch"
                    export GOOS="$selected_os"
                    go build -o AdGuardQueryTracker.exe
                elif [ "$selected_os" = "linux" ]; then
                    make GOOS="$selected_os" GOARCH="$selected_arch" output=AdGuardQueryTracker
                fi
                break
                ;;
            3)
                break
                ;;
            *)
                echo "Choix invalide, veuillez réessayer."
                continue
                ;;
        esac
    done

    # Afficher la sélection finale et les variables d'environnement exportées
    echo "Vous avez sélectionné :"
    echo "Système d'exploitation : $selected_os"
    echo "Architecture : $selected_arch"
    if [ "$selected_os" = "windows" ]; then
        echo "Variables d'environnement exportées :"
        echo "GOARCH : $GOARCH"
        echo "GOOS : $GOOS"
    fi
    echo ""
done
