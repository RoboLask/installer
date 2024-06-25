#!/bin/bash

printf "Select your preferred desktop environment:\n1.Plasma\n2.GNOME\n3.Hyprland(Wayland compositor)\nenter your choice(in number):"
read user_input

if [ "$user_input" -eq "1" ]; then
    sudo pacman -Rns --noconfirm gnome
    sudo pacman -S --noconfirm kde plasma sddm
    systemctl enable sddm
elif [ "$user_input" -eq "2"]; then
    systemctl enable gdm
elif [ "$user_input" -eq "3" ]; then
    sudo pacman -Rns --noconfirm gnome
    sudo pacman -S --noconfirm hyprland sddm
    systemctl enable sddm
else
    echo "Wrong number typed, Aboring..."
fi
