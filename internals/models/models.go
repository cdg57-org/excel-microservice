package models

type RGPD_COLL struct {
	RgpdColCode      string  `json:"RGPD_COL_CODE"`
	ColIdentite      string  `json:"COL_IDENTITE"`
	ColEmail         string  `json:"COL_EMAIL"`
	ColTel           string  `json:"COL_TEL"`
	Cnracl           float64 `json:"CNRACL"`
	Rg               float64 `json:"RG"`
	Autre            float64 `json:"AUTRE"`
	Total            float64 `json:"TOTAL"`
	Facture1ErAnnee  int64   `json:"FACTURE_1ER_ANNEE"`
	Facture2EmeAnnee int64   `json:"FACTURE_2EME_ANNEE"`
	Facture3EmeAnnee int64   `json:"FACTURE_3EME_ANNEE"`
	Facture4EmeAnnee int64   `json:"FACTURE_4EME_ANNEE"`
	Facture5EmeAnnee int64   `json:"FACTURE_5EME_ANNEE"`
}

type RGPD_COLL_COMPLET struct {
	RgpdColCode string `json:"RGPD_COL_CODE"`
	ColIdentite string `json:"COL_IDENTITE"`
	ColEmail    string `json:"COL_EMAIL"`
	ColTel      string `json:"COL_TEL"`
}
