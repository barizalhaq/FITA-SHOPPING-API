package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/barizalhaq/fita_shopping_api/domain"
)

func (r *promoRepository) GetPromos() ([]domain.Promo, error) {
	var promo []domain.Promo

	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	promoFile, err := os.Open(filepath.Join(rootDir, "misc", "promo.json"))
	if err != nil {
		return nil, err
	}
	defer promoFile.Close()

	byteVal, err := ioutil.ReadAll(promoFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteVal, &promo)
	if err != nil {
		return nil, err
	}

	return promo, nil
}
