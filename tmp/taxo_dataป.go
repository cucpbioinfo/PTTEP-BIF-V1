package main

import (
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type TaxoData struct {
	Year                    string `csv:"year"`
	Project                 string `csv:"project"`
	Platform                string `csv:"platform"`
	Station                 string `csv:"station"`
	Latitude                string `csv:"latitude"`
	Longtitude              string `csv:"longtitude,"`
	IdentificationTechnique string `csv:"identification technique"`
	MethodOfCollection      string `csv:"method of collection"`
	MajorGroup              string `csv:"major group"`
	Kingdom                 string `csv:"kingdom"`
	SubKingdom              string `csv:"subkingdom"`
	InfraKingdom            string `csv:"infrakingdom"`
	SuperPhylum             string `csv:"superphylum"`
	Phylum                  string `csv:"phylum"`
	SubPhylum               string `csv:"subphylum"`
	InfraPhylum             string `csv:"infraphylum"`
	SuperClass              string `csv:"superclass"`
	Class                   string `csv:"class"`
	SubClass                string `csv:"subclass"`
	InfraClass              string `csv:"infraclass"`
	SuperOrder              string `csv:"superorder"`
	Order                   string `csv:"order"`
	SubOrder                string `csv:"suborder"`
	InfraOrder              string `csv:"infraorder"`
	SuperFamily             string `csv:"superfamily"`
	Family                  string `csv:"family"`
	SubFamily               string `csv:"subfamily"`
	SuperGenus              string `csv:"supergenus"`
	Genus                   string `csv:"genus"`
	SubGenus                string `csv:"subgenus"`
	SpeciesName             string `csv:"species name"`
	SubSpecies              string `csv:"subspecies name"`
	Density                 string `csv:"density"`
}

type MajorGroupData struct {
	MajorGroupName string
}

type KingdomData struct {
	MajorGroupName string
	KingdomName    string
}

type SubKingdomData struct {
	KingdomName    string
	SubKingdomName string
}

type InfraKingdomData struct {
	KingdomName      string
	SubKingdomName   string
	InfraKingdomName string
}

type SuperPhylumData struct {
	KingdomName      string
	InfraKingdomName string
	SuperPhylumName  string
}

type PhylumData struct {
	KingdomName     string
	SuperPhylumName string
	PhylumName      string
}

type SubPhylumData struct {
	PhylumName    string
	SubPhylumName string
}

type InfraPhylumData struct {
	PhylumName      string
	SubPhylumName   string
	InfraPhylumName string
}

type SuperClassData struct {
	PhylumName      string
	InfraPhylumName string
	SuperClassName  string
}

type ClassData struct {
	PhylumName     string
	SuperClassName string
	ClassName      string
}

type SubClassData struct {
	ClassName    string
	SubClassName string
}

type InfraClassData struct {
	ClassName      string
	SubClassName   string
	InfraClassName string
}

type SuperOrderData struct {
	ClassName      string
	InfraClassName string
	SuperOrderName string
}

type OrderData struct {
	ClassName      string
	SuperOrderName string
	OrderName      string
}

type SubOrderData struct {
	OrderName    string
	SubOrderName string
}

type InfraOrderData struct {
	OrderName      string
	SubOrderName   string
	InfraOrderName string
}

type SuperFamilyData struct {
	OrderName       string
	InfraOrderName  string
	SuperFamilyName string
}

type FamilyData struct {
	OrderName       string
	SuperFamilyName string
	FamilyName      string
}

type SubFamilyData struct {
	FamilyName    string
	SubFamilyName string
}

type SuperGenusData struct {
	FamilyName     string
	SubFamilyName  string
	SuperGenusName string
}

type GenusData struct {
	FamilyName     string
	SuperGenusName string
	GenusName      string
}

type SubGenusData struct {
	GenusName    string
	SubGenusName string
}

type SpeciesData struct {
	MajorGroupName   string
	KingdomName      string
	SubKingdomName   string
	InfraKingdomName string
	SuperPhylumName  string
	PhylumName       string
	SubPhylumName    string
	InfraPhylumName  string
	SuperClassName   string
	ClassName        string
	SubClassName     string
	InfraClassName   string
	SuperOrderName   string
	OrderName        string
	SubOrderName     string
	InfraOrderName   string
	SuperFamilyName  string
	FamilyName       string
	SubFamilyName    string
	SuperGenusName   string
	GenusName        string
	SubGenusName     string
	SpeciesName      string
}

type SubSpeciesData struct {
	SpeciesName    string
	SubSpeciesName string
}

type ProjectData struct {
	Project string
}

type PlatformData struct {
	Platform   string
	Latitude   string
	Longtitude string
	Project    string
}

type StationData struct {
	Station  string
	Platform string
}

type IdentificationData struct {
	IdentificationTechnique string
}

type MethodData struct {
	MethodOfCollection      string
	IdentificationTechnique string
}

type DensityData struct {
	Year                    string
	Project                 string
	Platform                string
	Station                 string
	Latitude                string
	Longtitude              string
	IdentificationTechnique string
	MethodOfCollection      string
	MajorGroup              string
	Kingdom                 string
	SubKingdom              string
	InfraKingdom            string
	SuperPhylum             string
	Phylum                  string
	SubPhylum               string
	InfraPhylum             string
	SuperClass              string
	Class                   string
	SubClass                string
	InfraClass              string
	SuperOrder              string
	Order                   string
	SubOrder                string
	InfraOrder              string
	SuperFamily             string
	Family                  string
	SubFamily               string
	SuperGenus              string
	Genus                   string
	SubGenus                string
	SpeciesName             string
	SubSpecies              string
	Density                 string
}

func FormatData(s string) string {

	newString := strings.ToLower(
		strings.TrimSpace(s),
	)

	if newString == "-" || newString == "" {
		return "NULL"
	} else {
		return newString
	}
}

func GetTaxoData() ([]TaxoData, error) {
	taxos := []TaxoData{}
	filename := "./data/mockup.csv"

	in, err := os.Open(filename)
	if err != nil {
		return taxos, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &taxos); err != nil {
		return taxos, err
	}

	for i, taxo := range taxos {
		taxos[i] = TaxoData{
			Year:                    FormatData(taxo.Year),
			Project:                 FormatData(taxo.Project),
			Station:                 FormatData(taxo.Station),
			Platform:                FormatData(taxo.Platform),
			Latitude:                FormatData(taxo.Latitude),
			Longtitude:              FormatData(taxo.Longtitude),
			IdentificationTechnique: FormatData(taxo.IdentificationTechnique),
			MethodOfCollection:      FormatData(taxo.MethodOfCollection),
			MajorGroup:              FormatData(taxo.MajorGroup),
			Kingdom:                 FormatData(taxo.Kingdom),
			SubKingdom:              FormatData(taxo.SubKingdom),
			InfraKingdom:            FormatData(taxo.InfraKingdom),
			SuperPhylum:             FormatData(taxo.SuperPhylum),
			Phylum:                  FormatData(taxo.Phylum),
			SubPhylum:               FormatData(taxo.SubPhylum),
			InfraPhylum:             FormatData(taxo.InfraPhylum),
			SuperClass:              FormatData(taxo.SuperClass),
			Class:                   FormatData(taxo.Class),
			SubClass:                FormatData(taxo.SubClass),
			InfraClass:              FormatData(taxo.InfraClass),
			SuperOrder:              FormatData(taxo.SuperOrder),
			Order:                   FormatData(taxo.Order),
			SubOrder:                FormatData(taxo.SubOrder),
			InfraOrder:              FormatData(taxo.InfraOrder),
			SuperFamily:             FormatData(taxo.SuperFamily),
			Family:                  FormatData(taxo.Family),
			SubFamily:               FormatData(taxo.SubFamily),
			SuperGenus:              FormatData(taxo.SuperGenus),
			Genus:                   FormatData(taxo.Genus),
			SubGenus:                FormatData(taxo.SubGenus),
			SpeciesName:             FormatData(taxo.SpeciesName),
			SubSpecies:              FormatData(taxo.SubSpecies),
			Density:                 FormatData(taxo.Density),
		}
	}

	return taxos, nil
}

func GetMajorGroupData() ([]MajorGroupData, error) {
	majorGroups := []MajorGroupData{}

	majorGroupSet := make(map[MajorGroupData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return majorGroups, err
	}

	for _, taxo := range taxos {
		majorGroupData := MajorGroupData{
			MajorGroupName: taxo.MajorGroup,
		}
		majorGroupSet[majorGroupData] = exists
	}

	for majorGroup := range majorGroupSet {
		majorGroups = append(majorGroups, majorGroup)
	}

	return majorGroups, nil
}

func GetKingdomData() ([]KingdomData, error) {
	kingdoms := []KingdomData{}

	kingdomSet := make(map[KingdomData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return kingdoms, err
	}

	for _, taxo := range taxos {
		kingdomData := KingdomData{
			KingdomName: taxo.Kingdom,
		}
		kingdomSet[kingdomData] = exists
	}

	for kingdom := range kingdomSet {
		kingdoms = append(kingdoms, kingdom)
	}

	return kingdoms, nil
}

func GetSubKingdomData() ([]SubKingdomData, error) {
	subKingdoms := []SubKingdomData{}

	subKingdomSet := make(map[SubKingdomData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subKingdoms, err
	}

	for _, taxo := range taxos {
		subKingdomData := SubKingdomData{
			KingdomName:    taxo.Kingdom,
			SubKingdomName: taxo.SubKingdom,
		}
		subKingdomSet[subKingdomData] = exists
	}

	for subKingdom := range subKingdomSet {
		subKingdoms = append(subKingdoms, subKingdom)
	}

	return subKingdoms, nil
}

func GetInfraKingdomData() ([]InfraKingdomData, error) {
	infraKingdoms := []InfraKingdomData{}

	infraKingdomSet := make(map[InfraKingdomData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return infraKingdoms, err
	}

	for _, taxo := range taxos {
		infraKingdomData := InfraKingdomData{
			KingdomName:      taxo.Kingdom,
			SubKingdomName:   taxo.SubKingdom,
			InfraKingdomName: taxo.InfraKingdom,
		}
		infraKingdomSet[infraKingdomData] = exists
	}

	for infraKingdom := range infraKingdomSet {
		infraKingdoms = append(infraKingdoms, infraKingdom)
	}

	return infraKingdoms, nil
}

func GetSuperPhylumData() ([]SuperPhylumData, error) {
	superPhylums := []SuperPhylumData{}

	superPhylumSet := make(map[SuperPhylumData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return superPhylums, err
	}

	for _, taxo := range taxos {
		superPhylumData := SuperPhylumData{
			KingdomName:      taxo.Kingdom,
			InfraKingdomName: taxo.InfraKingdom,
			SuperPhylumName:  taxo.SuperPhylum,
		}
		superPhylumSet[superPhylumData] = exists
	}

	for superPhylum := range superPhylumSet {
		superPhylums = append(superPhylums, superPhylum)
	}

	return superPhylums, nil
}

func GetPhylumData() ([]PhylumData, error) {
	phylums := []PhylumData{}

	phylumSet := make(map[PhylumData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return phylums, err
	}

	for _, taxo := range taxos {
		phylumData := PhylumData{
			KingdomName:     taxo.Kingdom,
			SuperPhylumName: taxo.SuperPhylum,
			PhylumName:      taxo.Phylum,
		}
		phylumSet[phylumData] = exists
	}

	for phylum := range phylumSet {
		phylums = append(phylums, phylum)
	}

	return phylums, nil
}

func GetSubPhylumData() ([]SubPhylumData, error) {
	subPhylums := []SubPhylumData{}

	subPhylumSet := make(map[SubPhylumData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subPhylums, err
	}

	for _, taxo := range taxos {
		subPhylumData := SubPhylumData{
			PhylumName:    taxo.Phylum,
			SubPhylumName: taxo.SubPhylum,
		}
		subPhylumSet[subPhylumData] = exists
	}

	for subPhylum := range subPhylumSet {
		subPhylums = append(subPhylums, subPhylum)
	}

	return subPhylums, nil
}

func GetInfraPhylumData() ([]InfraPhylumData, error) {
	infraPhylums := []InfraPhylumData{}

	infraPhylumSet := make(map[InfraPhylumData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return infraPhylums, err
	}

	for _, taxo := range taxos {
		infraPhylumData := InfraPhylumData{
			PhylumName:      taxo.Phylum,
			SubPhylumName:   taxo.SubPhylum,
			InfraPhylumName: taxo.InfraPhylum,
		}
		infraPhylumSet[infraPhylumData] = exists
	}

	for infraPhylum := range infraPhylumSet {
		infraPhylums = append(infraPhylums, infraPhylum)
	}

	return infraPhylums, nil
}

func GetSuperClassData() ([]SuperClassData, error) {
	superClasses := []SuperClassData{}

	superClassSet := make(map[SuperClassData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return superClasses, err
	}

	for _, taxo := range taxos {
		superClassData := SuperClassData{
			PhylumName:      taxo.Phylum,
			InfraPhylumName: taxo.InfraPhylum,
			SuperClassName:  taxo.SuperClass,
		}
		superClassSet[superClassData] = exists
	}

	for superClass := range superClassSet {
		superClasses = append(superClasses, superClass)
	}

	return superClasses, nil
}

func GetClassData() ([]ClassData, error) {
	classes := []ClassData{}

	classSet := make(map[ClassData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return classes, err
	}

	for _, taxo := range taxos {
		classData := ClassData{
			PhylumName:     taxo.Phylum,
			SuperClassName: taxo.SuperClass,
			ClassName:      taxo.Class,
		}
		classSet[classData] = exists
	}

	for class := range classSet {
		classes = append(classes, class)
	}

	return classes, nil
}

func GetSubClassData() ([]SubClassData, error) {
	subClasses := []SubClassData{}

	subClassSet := make(map[SubClassData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subClasses, err
	}

	for _, taxo := range taxos {
		subClassData := SubClassData{
			ClassName:    taxo.Class,
			SubClassName: taxo.SubClass,
		}
		subClassSet[subClassData] = exists
	}

	for subClass := range subClassSet {
		subClasses = append(subClasses, subClass)
	}

	return subClasses, nil
}

func GetInfraClassData() ([]InfraClassData, error) {
	infraClasses := []InfraClassData{}

	infraClassSet := make(map[InfraClassData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return infraClasses, err
	}

	for _, taxo := range taxos {
		infraClassData := InfraClassData{
			ClassName:      taxo.Class,
			SubClassName:   taxo.SubClass,
			InfraClassName: taxo.InfraClass,
		}
		infraClassSet[infraClassData] = exists
	}

	for infraClass := range infraClassSet {
		infraClasses = append(infraClasses, infraClass)
	}

	return infraClasses, nil
}

func GetSuperOrderData() ([]SuperOrderData, error) {
	superOrders := []SuperOrderData{}

	superOrderSet := make(map[SuperOrderData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return superOrders, err
	}

	for _, taxo := range taxos {
		superOrderData := SuperOrderData{
			ClassName:      taxo.Class,
			InfraClassName: taxo.InfraClass,
			SuperOrderName: taxo.SuperOrder,
		}
		superOrderSet[superOrderData] = exists
	}

	for superOrder := range superOrderSet {
		superOrders = append(superOrders, superOrder)
	}

	return superOrders, nil
}

func GetOrderData() ([]OrderData, error) {
	orders := []OrderData{}

	orderSet := make(map[string]OrderData)
	// var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return orders, err
	}

	for _, taxo := range taxos {
		orderData := OrderData{
			ClassName:      taxo.Class,
			SuperOrderName: taxo.SuperOrder,
			OrderName:      taxo.Order,
		}
		orderSet[orderData.OrderName] = orderData
	}

	for order := range orderSet {
		orders = append(orders, orderSet[order])
	}

	return orders, nil
}

func GetSubOrderData() ([]SubOrderData, error) {
	subOrders := []SubOrderData{}

	subOrderSet := make(map[SubOrderData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subOrders, err
	}

	for _, taxo := range taxos {
		subOrderData := SubOrderData{
			OrderName:    taxo.Order,
			SubOrderName: taxo.SubOrder,
		}
		subOrderSet[subOrderData] = exists
	}

	for subOrder := range subOrderSet {
		subOrders = append(subOrders, subOrder)
	}

	return subOrders, nil
}

func GetInfraOrderData() ([]InfraOrderData, error) {
	infraOrders := []InfraOrderData{}

	infraOrderSet := make(map[InfraOrderData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return infraOrders, err
	}

	for _, taxo := range taxos {
		infraOrderData := InfraOrderData{
			OrderName:      taxo.Order,
			SubOrderName:   taxo.SubOrder,
			InfraOrderName: taxo.InfraOrder,
		}
		infraOrderSet[infraOrderData] = exists
	}

	for infraOrder := range infraOrderSet {
		infraOrders = append(infraOrders, infraOrder)
	}

	return infraOrders, nil
}

func GetSuperFamilyData() ([]SuperFamilyData, error) {
	superFamilies := []SuperFamilyData{}

	superFamilySet := make(map[SuperFamilyData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return superFamilies, err
	}

	for _, taxo := range taxos {
		superFamilyData := SuperFamilyData{
			OrderName:       taxo.Order,
			InfraOrderName:  taxo.InfraOrder,
			SuperFamilyName: taxo.SuperFamily,
		}
		superFamilySet[superFamilyData] = exists
	}

	for superFamily := range superFamilySet {
		superFamilies = append(superFamilies, superFamily)
	}

	return superFamilies, nil
}

func GetFamilyData() ([]FamilyData, error) {
	families := []FamilyData{}

	familySet := make(map[FamilyData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return families, err
	}

	for _, taxo := range taxos {
		familyData := FamilyData{
			OrderName:       taxo.Order,
			SuperFamilyName: taxo.SuperFamily,
			FamilyName:      taxo.Family,
		}
		familySet[familyData] = exists
	}

	for family := range familySet {
		families = append(families, family)
	}

	return families, nil
}

func GetSubFamilyData() ([]SubFamilyData, error) {
	subFamilies := []SubFamilyData{}

	subFamilySet := make(map[SubFamilyData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subFamilies, err
	}

	for _, taxo := range taxos {
		subFamilyData := SubFamilyData{
			FamilyName:    taxo.Family,
			SubFamilyName: taxo.SubFamily,
		}
		subFamilySet[subFamilyData] = exists
	}

	for subFamily := range subFamilySet {
		subFamilies = append(subFamilies, subFamily)
	}

	return subFamilies, nil
}

func GetSuperGenusData() ([]SuperGenusData, error) {
	superGenusses := []SuperGenusData{}

	superGenusSet := make(map[SuperGenusData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return superGenusses, err
	}

	for _, taxo := range taxos {
		superGenusData := SuperGenusData{
			FamilyName:     taxo.Family,
			SubFamilyName:  taxo.SubFamily,
			SuperGenusName: taxo.SuperGenus,
		}
		superGenusSet[superGenusData] = exists
	}

	for superGenus := range superGenusSet {
		superGenusses = append(superGenusses, superGenus)
	}

	return superGenusses, nil
}

func GetGenusData() ([]GenusData, error) {
	genusses := []GenusData{}

	genusSet := make(map[GenusData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return genusses, err
	}

	for _, taxo := range taxos {
		genusData := GenusData{
			FamilyName:     taxo.Family,
			SuperGenusName: taxo.SuperGenus,
			GenusName:      taxo.Genus,
		}
		genusSet[genusData] = exists
	}

	for genus := range genusSet {
		genusses = append(genusses, genus)
	}

	return genusses, nil
}

func GetSubGenusData() ([]SubGenusData, error) {
	subGenusses := []SubGenusData{}

	subGenusSet := make(map[SubGenusData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subGenusses, err
	}

	for _, taxo := range taxos {
		subGenusData := SubGenusData{
			GenusName:    taxo.Genus,
			SubGenusName: taxo.SubGenus,
		}
		subGenusSet[subGenusData] = exists
	}

	for subGenus := range subGenusSet {
		subGenusses = append(subGenusses, subGenus)
	}

	return subGenusses, nil
}

func GetSpeciesData() ([]SpeciesData, error) {
	speciesList := []SpeciesData{}

	speciesSet := make(map[SpeciesData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return speciesList, err
	}

	for _, taxo := range taxos {
		speciesData := SpeciesData{
			MajorGroupName:   taxo.MajorGroup,
			KingdomName:      taxo.Kingdom,
			SubKingdomName:   taxo.SubKingdom,
			InfraKingdomName: taxo.InfraKingdom,
			SuperPhylumName:  taxo.SuperPhylum,
			PhylumName:       taxo.Phylum,
			SubPhylumName:    taxo.SubPhylum,
			InfraPhylumName:  taxo.InfraPhylum,
			SuperClassName:   taxo.SuperClass,
			ClassName:        taxo.Class,
			SubClassName:     taxo.SubClass,
			InfraClassName:   taxo.InfraClass,
			SuperOrderName:   taxo.SuperOrder,
			OrderName:        taxo.Order,
			SubOrderName:     taxo.SubOrder,
			InfraOrderName:   taxo.InfraOrder,
			SuperFamilyName:  taxo.SuperFamily,
			FamilyName:       taxo.Family,
			SubFamilyName:    taxo.SubFamily,
			SuperGenusName:   taxo.SuperGenus,
			GenusName:        taxo.Genus,
			SubGenusName:     taxo.SubGenus,
			SpeciesName:      taxo.SpeciesName,
		}
		speciesSet[speciesData] = exists
	}

	for species := range speciesSet {
		speciesList = append(speciesList, species)
	}

	return speciesList, nil
}

func GetSubSpeciesData() ([]SubSpeciesData, error) {
	subSpeciesList := []SubSpeciesData{}

	subSpeciesSet := make(map[SubSpeciesData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return subSpeciesList, err
	}

	for _, taxo := range taxos {
		subSpeciesData := SubSpeciesData{
			SpeciesName:    taxo.SpeciesName,
			SubSpeciesName: taxo.SubSpecies,
		}
		subSpeciesSet[subSpeciesData] = exists
	}

	for subSpecies := range subSpeciesSet {
		subSpeciesList = append(subSpeciesList, subSpecies)
	}

	return subSpeciesList, nil
}

func GetProjectData() ([]ProjectData, error) {
	ProjectList := []ProjectData{}

	ProjectSet := make(map[ProjectData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return ProjectList, err
	}

	for _, taxo := range taxos {
		ProjectData := ProjectData{
			Project: taxo.Project,
		}
		ProjectSet[ProjectData] = exists
	}

	for Projects := range ProjectSet {
		ProjectList = append(ProjectList, Projects)
	}

	return ProjectList, nil
}

func GetPlatformData() ([]PlatformData, error) {
	PlatformList := []PlatformData{}

	PlatformSet := make(map[PlatformData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return PlatformList, err
	}

	for _, taxo := range taxos {
		PlatformData := PlatformData{
			Platform:   taxo.Platform,
			Latitude:   taxo.Latitude,
			Longtitude: taxo.Longtitude,
			Project:    taxo.Project,
		}
		PlatformSet[PlatformData] = exists
	}

	for Platform := range PlatformSet {
		PlatformList = append(PlatformList, Platform)
	}

	return PlatformList, nil
}

func GetStationData() ([]StationData, error) {
	StationList := []StationData{}

	StationSet := make(map[StationData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return StationList, err
	}

	for _, taxo := range taxos {
		StationData := StationData{
			Station:  taxo.Station,
			Platform: taxo.Platform,
		}
		StationSet[StationData] = exists
	}

	for Station := range StationSet {
		StationList = append(StationList, Station)
	}

	return StationList, nil
}

func GetIdentificationData() ([]IdentificationData, error) {
	IdentificationList := []IdentificationData{}

	IdentificationSet := make(map[IdentificationData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return IdentificationList, err
	}

	for _, taxo := range taxos {
		IdentificationData := IdentificationData{
			IdentificationTechnique: taxo.IdentificationTechnique,
		}
		IdentificationSet[IdentificationData] = exists
	}

	for Identification := range IdentificationSet {
		IdentificationList = append(IdentificationList, Identification)
	}

	return IdentificationList, nil
}

func GetMethodData() ([]MethodData, error) {
	MethodList := []MethodData{}

	MethodSet := make(map[MethodData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return MethodList, err
	}

	for _, taxo := range taxos {
		MethodData := MethodData{
			MethodOfCollection:      taxo.MethodOfCollection,
			IdentificationTechnique: taxo.IdentificationTechnique,
		}
		MethodSet[MethodData] = exists
	}

	for Method := range MethodSet {
		MethodList = append(MethodList, Method)
	}

	return MethodList, nil
}

func GetDensityData() ([]DensityData, error) {
	DensityList := []DensityData{}

	DensitySet := make(map[DensityData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return DensityList, err
	}

	for _, taxo := range taxos {
		densityData := DensityData{
			Year:                    taxo.Year,
			Project:                 taxo.Project,
			Platform:                taxo.Platform,
			Station:                 taxo.Station,
			IdentificationTechnique: taxo.IdentificationTechnique,
			MethodOfCollection:      taxo.MethodOfCollection,
			MajorGroup:              taxo.MajorGroup,
			Kingdom:                 taxo.Kingdom,
			SubKingdom:              taxo.SubKingdom,
			InfraKingdom:            taxo.InfraKingdom,
			SuperPhylum:             taxo.SuperPhylum,
			Phylum:                  taxo.Phylum,
			SubPhylum:               taxo.SubPhylum,
			InfraPhylum:             taxo.InfraPhylum,
			SuperClass:              taxo.SuperClass,
			Class:                   taxo.Class,
			SubClass:                taxo.SubClass,
			InfraClass:              taxo.InfraClass,
			SuperOrder:              taxo.SuperOrder,
			Order:                   taxo.Order,
			SubOrder:                taxo.SubOrder,
			InfraOrder:              taxo.InfraOrder,
			SuperFamily:             taxo.SuperFamily,
			Family:                  taxo.Family,
			SubFamily:               taxo.SubFamily,
			SuperGenus:              taxo.SuperGenus,
			Genus:                   taxo.Genus,
			SubGenus:                taxo.SubGenus,
			SpeciesName:             taxo.SpeciesName,
			Density:                 taxo.Density,
		}
		DensitySet[densityData] = exists
	}

	for densitys := range DensitySet {
		DensityList = append(DensityList, densitys)
	}

	return DensityList, nil
}
