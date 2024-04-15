package utils

import (
	"errors"
	"html"
	"strconv"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/microcosm-cc/bluemonday"
	"go.mongodb.org/mongo-driver/bson"
)

var p = bluemonday.UGCPolicy()

func GetIntQueryParam(param string) (int, error) {
	// prevent xss attack
	p.AllowComments()
	paramValue := p.Sanitize(param)
	if paramValue == "" {
		return 0, nil // No value provided, return default
	}
	intValue, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0, err // Invalid value
	}

	if intValue < 0 {
		return 0, errors.New("invalid value")
	}

	return intValue, nil
}

func ValidateStringQueryParam(param string) (string, error) {
	// prevent xss attack
	paramValue := p.Sanitize(param)
	if paramValue == "" {
		return "", nil // No value provided, return default
	}

	return paramValue, nil
}

func ValidateArrayQueryParam(param []string) ([]string, error) {
	// prevent xss attack
	var sanitizedParams []string
	sanitizedParams = append(sanitizedParams, param...)
	return sanitizedParams, nil
}

func SanitizeString(data string) string {
	s := p.Sanitize(data)
	return UnescapeString(s)
}

func PetSanitize(pet *models.Pet) {
	name := p.Sanitize(pet.Name)
	pet.Name = UnescapeString(name)
	desc := p.Sanitize(pet.Description)
	pet.Description = UnescapeString(desc)
	pet.Category = p.Sanitize(pet.Category)
	pet.Species = p.Sanitize(pet.Species)
	beh := p.Sanitize(pet.Behavior)
	pet.Behavior = UnescapeString(beh)
	// pet.Media = p.Sanitize(pet.Media)
	pet.Sex = p.Sanitize(pet.Sex)
	for i := range pet.Medical_records {
		medDesc := p.Sanitize(pet.Medical_records[i].Description)
		pet.Medical_records[i].Description = UnescapeString(medDesc)
		// pet.Medical_records[i].Description = p.Sanitize(pet.Medical_records[i].Description)
		pet.Medical_records[i].Medical_date = p.Sanitize(pet.Medical_records[i].Medical_date)
		pet.Medical_records[i].Medical_id = p.Sanitize(pet.Medical_records[i].Medical_id)
	}
}

func BsonSanitize(data *bson.M) {
	for key, value := range *data {
		switch value := value.(type) {
		case string:
			// (*data)[key] = p.Sanitize(value)
			s := p.Sanitize(value)
			(*data)[key] = UnescapeString(s)
		case bson.M:
			BsonSanitize(&value)
		}
	}
}

func UserSaniatize(user *models.User) {
	user.Username = p.Sanitize(user.Username)
	user.Email = p.Sanitize(user.Email)
	user.FirstName = p.Sanitize(user.FirstName)
	user.LastName = p.Sanitize(user.LastName)
	// user.Password = p.Sanitize(user.Password)
	user.Gender = p.Sanitize(user.Gender)
	user.PhoneNumber = p.Sanitize(user.PhoneNumber)
	// user.Image = p.Sanitize(user.Image)
	province := p.Sanitize(user.Address.Province)
	user.Address.Province = UnescapeString(province)
	district := p.Sanitize(user.Address.District)
	user.Address.District = UnescapeString(district)
	subdist := p.Sanitize(user.Address.SubDistrict)
	user.Address.SubDistrict = UnescapeString(subdist)
	user.Address.PostalCode = p.Sanitize(user.Address.PostalCode)
	street := p.Sanitize(user.Address.Street)
	user.Address.Street = UnescapeString(street)
	building := p.Sanitize(user.Address.Building)
	user.Address.Building = UnescapeString(building)
	houseNO := p.Sanitize(user.Address.HouseNumber)
	user.Address.HouseNumber = UnescapeString(houseNO)
}

func UnescapeString(data string) string {
	return html.UnescapeString(data)
}
