package main

func (v *personnage) Init(nom string, classe string, niveau int, hpmax int, hpactuels int, inventaire []string, argent int, sorts []string, equipement []string, xp int, attaque int, inventairemax int, inventairemegamax int, magie int, argentmax int) {
	v.nom = nom
	v.classe = classe
	v.niveau = niveau
	v.hpmax = hpmax
	v.hpactuels = hpactuels
	v.inventaire = inventaire
	v.argent = argent
	v.equipement = equipement
	v.xp = xp
	v.attaque = attaque
	v.inventairemax = inventairemax
	v.inventairemegamax = inventairemegamax
	v.magie = magie
	v.argentmax = argentmax
}
