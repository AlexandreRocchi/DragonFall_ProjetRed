package main

type monstre struct {
	nom_monstre       string
	hpmax_monstre     int
	hpactuels_monstre int
	attack_monstre    int
	tour_monstre      int
}

func (v *monstre) Init(nom_monstre string, hpmax_monstre int, hpactuels_monstre int, attack_monstre int, tour_monstre int) {
	v.nom_monstre = nom_monstre
	v.hpmax_monstre = hpmax_monstre
	v.hpactuels_monstre = hpactuels_monstre
	v.attack_monstre = attack_monstre
	v.tour_monstre = tour_monstre

}
