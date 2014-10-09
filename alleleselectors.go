package main

func getHairColor(hc1 Allele, hc2 Allele) string {
	result := ""
	if hc1.Pos1 && hc1.Pos2 {
		result = "Black"
	} else {
		if hc1.Pos1 || hc1.Pos1 {
			result = "Brown"
		} else {
			if hc2.Pos1 || hc2.Pos2 {
				result = "Blonde"
			} else {
				result = "Red"
			}
		}
	}
	return result
}

func setDdAllele(allele Allele, dominant string, recessive string) string {
	if allele.Pos1 || allele.Pos2 {
		return dominant
	}
	return recessive
}

func getEyeColor(ec1 Allele, ec2 Allele) string {
	result := ""
	if ec1.Pos1 || ec1.Pos2 {
		result = "Brown"
	} else {
		if ec2.Pos1 || ec2.Pos2 {
			result = "Green"
		} else {
			result = "Blue"
		}
	}
	return result
}

func getColorBlind(cb1 Allele, cb2 Allele, gender string) string {
	//cb1 must be from father
	//cb2 must be from mother
	result := "Normal"
	rgblind := "R/G color blind"
	if gender == "X" {
		if !cb1.Pos1 && !cb1.Pos2 && !cb2.Pos1 && !cb2.Pos2 {
			result = rgblind
		}
	} else {
		if !cb2.Pos1 && !cb2.Pos2 {
			result = rgblind
		}
	}
	return result
}
