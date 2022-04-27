#!/usr/bin/perl

#IMIE: Mateusz
#Nazwisko: Redzimski
#Nr. Studenta: 266601
#/Grupa: 3
#Tytuł: Artes
use warnings;
use strict;

my $filename = './slowo.txt';
my $filename1 = './a.txt';
my $filename2 = './wynik.txt';
my $slowo = "x";
my $sukces = 0;
#owarcie pliku ze słowem które podalismy oraz zczytanie go do zmiennej
open(FH, '<', $filename) or die $!;
while(<FH>){
   $slowo = $_;
}
close(FH);
#podzielenie słowa na litery w celu uzyskania 1 litery i sprawdzeniu w ktorym pliku skrypt ma szukać hasła
my @chars = split("",$slowo);
$filename1 = $chars[0] . '.txt';

#otwarcie odpowiedniego pliku w celu znalezienia slowa np dla słowa "Artes" otworzymy plik a.txt
open(FH, '<', $filename1) or die $!;

while(<FH>){
   if($slowo eq $_){
      print $_;
      $sukces = 1;  
   }
}
close(FH);
#otwarcie pliku wyniku i zapisanie do niego rezultatów jezeli slowo znajduje sie w slowniku wynik =1 jezeli nie wynik=0 
open(FH, '>', $filename2) or die $!;
print FH $sukces;
close(FH);  
print $sukces


