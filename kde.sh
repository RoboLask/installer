#!/bin/bash

echo "youe choice"
read user_input

if [ "$user_input" -eq 1 ]; then
    sudo pacman -S plasma konsole sddm dolphin
    systemctl enable sddm.service
else
    echo "sd"
fi
