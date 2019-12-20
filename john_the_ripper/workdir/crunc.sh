#!/bin/bash

crunch 13 13 -t Verifikacija% -o wordlist_1.txt
crunch 13 13 -t verifikacija% -o wordlist_2.txt
crunch 14 14 -t %Verifikacija% -o wordlist_3.txt
crunch 14 14 -t %verifikacija% -o wordlist_4.txt
crunch 15 15 -t %Verifikacija%% -o wordlist_5.txt
crunch 15 15 -t %verifikacija%% -o wordlist_6.txt
crunch 16 16 -t %%Verifikacija%% -o wordlist_7.txt
crunch 16 16 -t %%verifikacija%% -o wordlist_8.txt

