package medicineapi

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	pdfToText "github.com/heussd/pdftotext-go"
	"github.com/higorjsilva/goapi/types"
)

type MedicineAPI struct {
	baseUrl string
}

func NewApi() *MedicineAPI {
	return &MedicineAPI{ baseUrl: "https://consultas.anvisa.gov.br/api/consulta",}
}

func (m *MedicineAPI) GetMedicines(query types.SearchMedicinePayload) (types.MedicinesAPIResponse, error)  {

	urlParams, err := setQueryParams(query)

	if err != nil {
		fmt.Println("Error:", err)
		return types.MedicinesAPIResponse{}, err
	}

	baseURL := m.baseUrl + "/bulario"
	url := fmt.Sprintf("%s?%s", baseURL, urlParams)
	method := "GET"

	body, err := MakeRequest(method, url)

	if err != nil {
		log.Fatalln("Error making request:", err)
		return types.MedicinesAPIResponse{}, err
	}

	var apiResponse types.MedicinesAPIResponse

	err1 := json.Unmarshal([]byte(body), &apiResponse)

	if err1 != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return apiResponse, nil
}

func (m *MedicineAPI) GetSideEffects(query types.SearchMedicinePayload) (string, error)  {

	medicines, err := m.GetMedicines(query)

	if err != nil  || len(medicines.Content) == 0{
		return "Error searching medicine",  errors.New("Error searching medicine")
	}

	baseURL := m.baseUrl + "/medicamentos/arquivo/bula/parecer/%s/?Authorization="
	url := fmt.Sprintf(baseURL, medicines.Content[0].IDBulaPacienteProtegido)
	method := "GET"

	body, err := MakeRequest(method, url)

	if err != nil {
		log.Print("Error making request:", err)
		return "Error making request", err
	}

	pages, err := readFile(body)

	if err != nil {
		return "error converting pdf", errors.New("error converting pdf")
	}

	var builder strings.Builder

	for _, page := range pages {
		builder.WriteString(page.Content)
		builder.WriteString("\n")
	}

	allContent := builder.String()
	pattern := `8\. Q([\s\S]*?)9\.`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(allContent)

	if len(match) > 0 {
		return strings.Replace(match[0], "8. ", "", -1), err
	}

	return "Not Found", errors.New("Not Found")
}

func MakeRequest(method, url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Authorization", "Guest")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func setQueryParams(payload types.SearchMedicinePayload) (string, error){
	v := reflect.ValueOf(payload)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected a struct, but got %T", payload)
	}

	params := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i).Interface()

		if value != "" && value != 0 {
			params.Add(fmt.Sprintf("filter[%s]", field.Tag.Get("json")), fmt.Sprintf("%v", value))
		}
	}

	return params.Encode(), nil
}

func readFile(pdf []byte) ([]pdfToText.PdfPage, error) {
	pages, err := pdfToText.Extract(pdf)

	if err != nil {
		log.Println("error converting pdf", err)
	}

	return pages, err
}


