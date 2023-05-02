package groupie

import (
	"strconv"
)

func Filtre(Detail []Detail, Vide *Vide, DetailFull []Detail, DetailTest []Detail, member int, date int, album int, europe string, amerique string, afrique string, asie string, oceanie string) []Detail {

	Temp := DetailTest

	first := 0

	// filtre tout

	DetailTest, first = FilterMember(Detail, DetailTest, Temp, member, first)
	DetailTest, first = FilterDate(Detail, DetailTest, Temp, date, first)
	DetailTest, first = FilterAlbum(Detail, DetailTest, Temp, album, first)
	DetailTest, first = FilterEurope(Detail, DetailTest, Temp, Vide, europe, first)
	DetailTest, first = FilterAmerica(Detail, DetailTest, Temp, Vide, amerique, first)
	DetailTest, first = FilterAfrica(Detail, DetailTest, Temp, Vide, afrique, first)
	DetailTest, first = FilterAsia(Detail, DetailTest, Temp, Vide, asie, first)
	DetailTest, first = FilterOceania(Detail, DetailTest, Temp, Vide, oceanie, first)

	if len(DetailTest) > 0 {
		DetailFull = DetailTest
	} else {
		DetailFull = Detail
	}

	return DetailFull
}

func FilterMember(Detail []Detail, DetailTest []Detail, Temp []Detail, member int, first int) ([]Detail, int) { // filtre les artiste par nombre de membre

	if member > 0 {
		if first == 0 {
			for i := 0; i < len(Detail); i++ {
				if len((Detail)[i].Members) == member {
					DetailTest = append(DetailTest, (Detail)[i])
					first = 1
				}
			}
			return DetailTest, first
		} else {
			if member > 1901 {
				for i := 0; i < len(DetailTest); i++ {
					if len((DetailTest)[i].Members) == member {
						Temp = append(Temp, (DetailTest)[i])
					}
				}
				return Temp, first
			}
		}
	}
	return DetailTest, first
}

func FilterDate(Detail []Detail, DetailTest []Detail, Temp []Detail, date int, first int) ([]Detail, int) { // filtre les artiste par date de creation

	if date > 1901 {
		if first == 0 {
			for i := 0; i < len(Detail); i++ {
				for y := date; y < date+10; y++ {
					if Detail[i].CreationDate == y {
						DetailTest = append(DetailTest, (Detail)[i])
						first = 1
					}
				}
			}
			return DetailTest, first
		} else {
			for i := 0; i < len(DetailTest); i++ {
				for y := date; y < date+10; y++ {
					if DetailTest[i].CreationDate == y {
						Temp = append(Temp, (DetailTest)[i])
						first = 1
					}
				}
			}
			return Temp, first
		}
	}
	return DetailTest, first
}

func FilterAlbum(Detail []Detail, DetailTest []Detail, Temp []Detail, album int, first int) ([]Detail, int) { // filtre les artiste par nombre d'album

	if album > 1901 {
		if first == 0 {
			for i := 0; i < len(Detail); i++ {
				for y := album; y < album+10; y++ {
					stringAlbumDate := Detail[i].FirstAlbum[6:]
					albumDate, _ := strconv.Atoi(stringAlbumDate)
					if albumDate == y {
						DetailTest = append(DetailTest, (Detail)[i])
					}
				}
			}
			return DetailTest, first
		}
	} else {
		if album > 1901 {
			for i := 0; i < len(DetailTest); i++ {
				for y := album; y < album+10; y++ {
					stringAlbumDate := DetailTest[i].FirstAlbum[6:]
					albumDate, _ := strconv.Atoi(stringAlbumDate)
					if albumDate == y {
						Temp = append(Temp, (DetailTest)[i])
					}
				}
			}
			return DetailTest, first
		}
	}
	return DetailTest, first
}

func FilterEurope(Detail []Detail, DetailTest []Detail, Temp []Detail, Vide *Vide, europe string, first int) ([]Detail, int) { // filtre les artiste par continent

	xd := 0

	if europe != "" {
		if first == 0 {
			for i := 0; i < len(Vide.Index); i++ {
				for y := 0; y < len(Vide.Index[i].Continent); y++ {
					if Vide.Index[i].Continent[y] == 1 {
						DetailTest = append(DetailTest, (Detail)[i])
						xd = 1
						break
					}
					if xd == 1 {
						break
					}
				}
			}

			return DetailTest, first

		} else {
			if europe != "" {
				for i := 0; i < len(Vide.Index); i++ {
					for y := 0; y < len(Vide.Index[i].Continent); y++ {
						if Vide.Index[i].Continent[y] == 1 {
							for z := 0; z < len(DetailTest); z++ {
								if Vide.Index[i].Id == DetailTest[z].Id {
									Temp = append(Temp, (Detail)[i])
									xd = 1
									break
								}
							}
							if xd == 1 {
								break
							}
						}
					}
				}

				return Temp, first
			}
		}
	}
	return DetailTest, first
}

func FilterAsia(Detail []Detail, DetailTest []Detail, Temp []Detail, Vide *Vide, asie string, first int) ([]Detail, int) { // filtre les artiste par continent

	xd := 0

	if asie != "" {
		if first == 0 {
			for i := 0; i < len(Vide.Index); i++ {
				for y := 0; y < len(Vide.Index[i].Continent); y++ {
					if Vide.Index[i].Continent[y] == 2 {
						DetailTest = append(DetailTest, (Detail)[i])
						xd = 1
						break
					}
					if xd == 1 {
						break
					}
				}
			}

			return DetailTest, first

		} else {
			if asie != "" {
				for i := 0; i < len(Vide.Index); i++ {
					for y := 0; y < len(Vide.Index[i].Continent); y++ {
						if Vide.Index[i].Continent[y] == 2 {
							for z := 0; z < len(DetailTest); z++ {
								if Vide.Index[i].Id == DetailTest[z].Id {
									Temp = append(Temp, (Detail)[i])
									xd = 1
									break
								}
							}
							if xd == 1 {
								break
							}
						}
					}
				}

				return Temp, first
			}
		}
	}
	return DetailTest, first
}

func FilterAmerica(Detail []Detail, DetailTest []Detail, Temp []Detail, Vide *Vide, amerique string, first int) ([]Detail, int) { // filtre les artiste par continent

	xd := 0

	if amerique != "" {
		if first == 0 {
			for i := 0; i < len(Vide.Index); i++ {
				for y := 0; y < len(Vide.Index[i].Continent); y++ {
					if Vide.Index[i].Continent[y] == 3 {
						DetailTest = append(DetailTest, (Detail)[i])
						xd = 1
						break
					}
					if xd == 1 {
						break
					}
				}
			}

			return DetailTest, first

		} else {
			if amerique != "" {
				for i := 0; i < len(Vide.Index); i++ {
					for y := 0; y < len(Vide.Index[i].Continent); y++ {
						if Vide.Index[i].Continent[y] == 3 {
							for z := 0; z < len(DetailTest); z++ {
								if Vide.Index[i].Id == DetailTest[z].Id {
									Temp = append(Temp, (Detail)[i])
									xd = 1
									break
								}
							}
							if xd == 1 {
								break
							}
						}
					}
				}

				return Temp, first
			}
		}
	}
	return DetailTest, first
}

func FilterAfrica(Detail []Detail, DetailTest []Detail, Temp []Detail, Vide *Vide, afrique string, first int) ([]Detail, int) { // filtre les artiste par continent

	xd := 0

	if afrique != "" {
		if first == 0 {
			for i := 0; i < len(Vide.Index); i++ {
				for y := 0; y < len(Vide.Index[i].Continent); y++ {
					if Vide.Index[i].Continent[y] == 4 {
						DetailTest = append(DetailTest, (Detail)[i])
						xd = 1
						break
					}
					if xd == 1 {
						break
					}
				}
			}

			return DetailTest, first

		} else {
			if afrique != "" {
				for i := 0; i < len(Vide.Index); i++ {
					for y := 0; y < len(Vide.Index[i].Continent); y++ {
						if Vide.Index[i].Continent[y] == 4 {
							for z := 0; z < len(DetailTest); z++ {
								if Vide.Index[i].Id == DetailTest[z].Id {
									Temp = append(Temp, (Detail)[i])
									xd = 1
									break
								}
							}
							if xd == 1 {
								break
							}
						}
					}
				}

				return Temp, first
			}
		}
	}
	return DetailTest, first
}

func FilterOceania(Detail []Detail, DetailTest []Detail, Temp []Detail, Vide *Vide, oceanie string, first int) ([]Detail, int) { // filtre les artiste par continent

	xd := 0

	if oceanie != "" {
		if first == 0 {
			for i := 0; i < len(Vide.Index); i++ {
				for y := 0; y < len(Vide.Index[i].Continent); y++ {
					if Vide.Index[i].Continent[y] == 2 {
						DetailTest = append(DetailTest, (Detail)[i])
						xd = 1
						break
					}
					if xd == 1 {
						break
					}
				}
			}

			return DetailTest, first

		} else {
			if oceanie != "" {
				for i := 0; i < len(Vide.Index); i++ {
					for y := 0; y < len(Vide.Index[i].Continent); y++ {
						if Vide.Index[i].Continent[y] == 2 {
							for z := 0; z < len(DetailTest); z++ {
								if Vide.Index[i].Id == DetailTest[z].Id {
									Temp = append(Temp, (Detail)[i])
									xd = 1
									break
								}
							}
							if xd == 1 {
								break
							}
						}
					}
				}

				return Temp, first
			}
		}
	}
	return DetailTest, first
}
