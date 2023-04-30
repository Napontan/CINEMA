package movie // ตั้งชื่อ package ไม่ตรงกับ ไฟล์ไม่ได้

func FindMovieName(imdbId string) string {
	switch imdbId {
	case "tt4154796":
		return "Avengers: Endgame"
	case "tt1825683":
		return "Black Panther"
	}

	return "not found"
}
