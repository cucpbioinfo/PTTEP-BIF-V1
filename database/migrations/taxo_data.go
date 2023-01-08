package main

import (
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type TaxoData struct {
	IdentificationTechnique string `csv:"IdentificationTechnique"`
	MethodOfCollection      string `csv:"MethodOfCollection"`
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
}

type IdentificationData struct {
	IdentificationTechniqueName string
}

type MethodData struct {
	MethodName                  string
	IdentificationTechniqueName string
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
	IdentificationName string
	MethodName         string
	MajorGroupName     string
	KingdomName        string
	SubKingdomName     string
	InfraKingdomName   string
	SuperPhylumName    string
	PhylumName         string
	SubPhylumName      string
	InfraPhylumName    string
	SuperClassName     string
	ClassName          string
	SubClassName       string
	InfraClassName     string
	SuperOrderName     string
	OrderName          string
	SubOrderName       string
	InfraOrderName     string
	SuperFamilyName    string
	FamilyName         string
	SubFamilyName      string
	SuperGenusName     string
	GenusName          string
	SubGenusName       string
	SpeciesName        string
}

type SubSpeciesData struct {
	SpeciesName    string
	SubSpeciesName string
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
	filename := "./data/datainput_taxo_76.csv"

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
		}
	}

	return taxos, nil
}

func GetIdentificationData() ([]IdentificationData, error) {
	identifications := []IdentificationData{}

	identificationSet := make(map[IdentificationData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return identifications, err
	}

	for _, taxo := range taxos {
		identificationData := IdentificationData{
			IdentificationTechniqueName: taxo.IdentificationTechnique,
		}
		identificationSet[identificationData] = exists
	}

	for identification := range identificationSet {
		identifications = append(identifications, identification)
	}

	return identifications, nil
}

func GetMethodData() ([]MethodData, error) {
	methods := []MethodData{}

	methodSet := make(map[MethodData]struct{})
	var exists = struct{}{}

	taxos, err := GetTaxoData()
	if err != nil {
		return methods, err
	}

	for _, taxo := range taxos {
		methodData := MethodData{
			MethodName:                  taxo.MethodOfCollection,
			IdentificationTechniqueName: taxo.IdentificationTechnique,
		}
		methodSet[methodData] = exists
	}

	for method := range methodSet {
		methods = append(methods, method)
	}

	return methods, nil
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
			IdentificationName: taxo.IdentificationTechnique,
			MethodName:         taxo.MethodOfCollection,
			MajorGroupName:     taxo.MajorGroup,
			KingdomName:        taxo.Kingdom,
			SubKingdomName:     taxo.SubKingdom,
			InfraKingdomName:   taxo.InfraKingdom,
			SuperPhylumName:    taxo.SuperPhylum,
			PhylumName:         taxo.Phylum,
			SubPhylumName:      taxo.SubPhylum,
			InfraPhylumName:    taxo.InfraPhylum,
			SuperClassName:     taxo.SuperClass,
			ClassName:          taxo.Class,
			SubClassName:       taxo.SubClass,
			InfraClassName:     taxo.InfraClass,
			SuperOrderName:     taxo.SuperOrder,
			OrderName:          taxo.Order,
			SubOrderName:       taxo.SubOrder,
			InfraOrderName:     taxo.InfraOrder,
			SuperFamilyName:    taxo.SuperFamily,
			FamilyName:         taxo.Family,
			SubFamilyName:      taxo.SubFamily,
			SuperGenusName:     taxo.SuperGenus,
			GenusName:          taxo.Genus,
			SubGenusName:       taxo.SubGenus,
			SpeciesName:        taxo.SpeciesName,
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

// Location Import
type LocationData struct {
	Asset     string `csv:"Asset"`
	Platform  string `csv:"Platform"`
	Station   string `csv:"Station"`
	Latitude  string `csv:"Latitude"`
	Longitude string `csv:"Longitude"`
	Type      string `csv:"Type"`
}

type AssetData struct {
	Asset     string
	Latitude  string
	Longitude string
	Type      string
}

type PlatformData struct {
	Asset     string
	Platform  string
	Latitude  string
	Longitude string
	Type      string
}

type StationData struct {
	Asset     string
	Platform  string
	Station   string
	Latitude  string
	Longitude string
	Type      string
}

func GetLocationData() ([]LocationData, error) {
	locations := []LocationData{}
	filename := "./data/datax - Location.csv"

	in, err := os.Open(filename)
	if err != nil {
		return locations, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &locations); err != nil {
		return locations, err
	}

	for i, location := range locations {
		locations[i] = LocationData{
			Asset:     FormatData(location.Asset),
			Platform:  FormatData(location.Platform),
			Station:   FormatData(location.Station),
			Latitude:  FormatData(location.Latitude),
			Longitude: FormatData(location.Longitude),
			Type:      FormatData(location.Type),
		}
	}

	return locations, nil
}

func GetAssetData() ([]AssetData, error) {
	AssetList := []AssetData{}

	AssetSet := make(map[AssetData]struct{})
	var exists = struct{}{}

	assets, err := GetLocationData()
	if err != nil {
		return AssetList, err
	}

	for _, asset := range assets {
		AssetData := AssetData{
			Asset: asset.Asset,
		}
		AssetSet[AssetData] = exists
	}

	for Assets := range AssetSet {
		AssetList = append(AssetList, Assets)
	}

	return AssetList, nil
}

func GetPlatformData() ([]PlatformData, error) {
	PlatformList := []PlatformData{}

	PlatformSet := make(map[PlatformData]struct{})
	var exists = struct{}{}

	platforms, err := GetLocationData()
	if err != nil {
		return PlatformList, err
	}

	for _, platform := range platforms {
		PlatformData := PlatformData{
			Asset:     platform.Asset,
			Platform:  platform.Platform,
			Latitude:  platform.Latitude,
			Longitude: platform.Longitude,
			Type:      platform.Type,
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

	stations, err := GetLocationData()
	if err != nil {
		return StationList, err
	}

	for _, station := range stations {
		StationData := StationData{
			Asset:     station.Asset,
			Platform:  station.Platform,
			Station:   station.Station,
			Latitude:  station.Latitude,
			Longitude: station.Longitude,
		}
		StationSet[StationData] = exists
	}

	for Station := range StationSet {
		StationList = append(StationList, Station)
	}

	return StationList, nil
}

// //No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
// Density Import
type DensityData struct {
	Species       string `csv:"Species"`
	Year          string `csv:"Year"`
	Asset         string `csv:"Asset"`
	Platform      string `csv:"Platform"`
	Station       string `csv:"Station"`
	Surface       string `csv:"Surface"`
	Euphotic_zone string `csv:"Euphotic_zone"`
}

func GetDensityData() ([]DensityData, error) {
	densities := []DensityData{}
	filename := "./data/datax - Density.csv"

	in, err := os.Open(filename)
	if err != nil {
		return densities, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &densities); err != nil {
		return densities, err
	}

	for i, location := range densities {
		densities[i] = DensityData{
			Species:       FormatData(location.Species),
			Year:          FormatData(location.Year),
			Asset:         FormatData(location.Asset),
			Platform:      FormatData(location.Platform),
			Station:       FormatData(location.Station),
			Surface:       FormatData(location.Surface),
			Euphotic_zone: FormatData(location.Euphotic_zone),
		}
	}

	return densities, nil
}

func GetDensitiesData() ([]DensityData, error) {
	DensityList := []DensityData{}

	DensitySet := make(map[DensityData]struct{})
	var exists = struct{}{}

	densities, err := GetDensityData()
	if err != nil {
		return DensityList, err
	}

	for _, diversity := range densities {
		DensityData := DensityData{
			Species:       diversity.Species,
			Year:          diversity.Year,
			Asset:         diversity.Asset,
			Platform:      diversity.Platform,
			Station:       diversity.Station,
			Surface:       diversity.Surface,
			Euphotic_zone: diversity.Euphotic_zone,
		}
		DensitySet[DensityData] = exists
	}

	for densities := range DensitySet {
		DensityList = append(DensityList, densities)
	}

	return DensityList, nil
}

//Diversity Import

type DiversityData struct {
	Year          string `csv:"Year"`
	Asset         string `csv:"Asset"`
	Platform      string `csv:"Platform"`
	Station       string `csv:"Station"`
	Surface       string `csv:"Surface"`
	Euphotic_zone string `csv:"Euphotic_zone"`
}

func GetDiversityData() ([]DiversityData, error) {
	diversities := []DiversityData{}
	filename := "./data/datax - DI.csv"

	in, err := os.Open(filename)
	if err != nil {
		return diversities, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &diversities); err != nil {
		return diversities, err
	}

	for i, location := range diversities {
		diversities[i] = DiversityData{
			Year:          FormatData(location.Year),
			Asset:         FormatData(location.Asset),
			Platform:      FormatData(location.Platform),
			Station:       FormatData(location.Station),
			Surface:       FormatData(location.Surface),
			Euphotic_zone: FormatData(location.Euphotic_zone),
		}
	}

	return diversities, nil
}

func GetDiversitiesData() ([]DiversityData, error) {
	DiversityList := []DiversityData{}

	DiversitySet := make(map[DiversityData]struct{})
	var exists = struct{}{}

	diversities, err := GetDiversityData()
	if err != nil {
		return DiversityList, err
	}

	for _, diversity := range diversities {
		DiversityData := DiversityData{
			Year:          diversity.Year,
			Asset:         diversity.Asset,
			Platform:      diversity.Platform,
			Station:       diversity.Station,
			Surface:       diversity.Surface,
			Euphotic_zone: diversity.Euphotic_zone,
		}
		DiversitySet[DiversityData] = exists
	}

	for diversities := range DiversitySet {
		DiversityList = append(DiversityList, diversities)
	}

	return DiversityList, nil
}

// //Evenness Import
type EvennessData struct {
	Year          string `csv:"Year"`
	Asset         string `csv:"Asset"`
	Platform      string `csv:"Platform"`
	Station       string `csv:"Station"`
	Surface       string `csv:"Surface"`
	Euphotic_zone string `csv:"Euphotic_zone"`
}

func GetEvennessData() ([]EvennessData, error) {
	evennesses := []EvennessData{}
	filename := "./data/datax - EI.csv"

	in, err := os.Open(filename)
	if err != nil {
		return evennesses, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &evennesses); err != nil {
		return evennesses, err
	}

	for i, location := range evennesses {
		evennesses[i] = EvennessData{
			Year:          FormatData(location.Year),
			Asset:         FormatData(location.Asset),
			Platform:      FormatData(location.Platform),
			Station:       FormatData(location.Station),
			Surface:       FormatData(location.Surface),
			Euphotic_zone: FormatData(location.Euphotic_zone),
		}
	}

	return evennesses, nil
}

func GetEvennessesData() ([]EvennessData, error) {
	EvennessList := []EvennessData{}

	EvennessSet := make(map[EvennessData]struct{})
	var exists = struct{}{}

	evennesses, err := GetEvennessData()
	if err != nil {
		return EvennessList, err
	}

	for _, diversity := range evennesses {
		EvennessData := EvennessData{
			Year:          diversity.Year,
			Asset:         diversity.Asset,
			Platform:      diversity.Platform,
			Station:       diversity.Station,
			Surface:       diversity.Surface,
			Euphotic_zone: diversity.Euphotic_zone,
		}
		EvennessSet[EvennessData] = exists
	}

	for evennesses := range EvennessSet {
		EvennessList = append(EvennessList, evennesses)
	}

	return EvennessList, nil
}

// //Summary Import   No.,Year,Group,Identification,Asset,Platform,Station,SurfaceShannon,SurfaceNumber,SurfaceMax,SurfaceEvenness,
// Euphotic_zoneShannon,Euphotic_zoneNumber,Euphotic_zoneMax,Euphotic_zoneEvenness
// AverageShannon,AverageNumber,AverageMax,AverageEvenness
type SummaryData struct {
	Year                  string `csv:"Year"`
	Group                 string `csv:"Group"`
	Identification        string `csv:"Identification"`
	Asset                 string `csv:"Asset"`
	Platform              string `csv:"Platform"`
	Station               string `csv:"Station"`
	SurfaceShannon        string `csv:"SurfaceShannon"`
	SurfaceNumber         string `csv:"SurfaceNumber"`
	SurfaceMax            string `csv:"SurfaceMax"`
	SurfaceEvenness       string `csv:"SurfaceEvenness"`
	Euphotic_zoneShannon  string `csv:"Euphotic_zoneShannon"`
	Euphotic_zoneNumber   string `csv:"Euphotic_zoneNumber"`
	Euphotic_zoneMax      string `csv:"Euphotic_zoneMax"`
	Euphotic_zoneEvenness string `csv:"Euphotic_zoneEvenness"`
	AverageShannon        string `csv:"AverageShannon"`
	AverageNumber         string `csv:"AverageNumber"`
	AverageMax            string `csv:"AverageMax"`
	AverageEvenness       string `csv:"AverageEvenness"`
}

func GetSummaryData() ([]SummaryData, error) {
	summaries := []SummaryData{}
	filename := "./data/datax - Data_Summary.csv"

	in, err := os.Open(filename)
	if err != nil {
		return summaries, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &summaries); err != nil {
		return summaries, err
	}

	for i, sum := range summaries {
		summaries[i] = SummaryData{
			Year:                  FormatData(sum.Year),
			Group:                 FormatData(sum.Group),
			Identification:        FormatData(sum.Identification),
			Asset:                 FormatData(sum.Asset),
			Platform:              FormatData(sum.Platform),
			Station:               FormatData(sum.Station),
			SurfaceShannon:        FormatData(sum.SurfaceShannon),
			SurfaceNumber:         FormatData(sum.SurfaceNumber),
			SurfaceMax:            FormatData(sum.SurfaceMax),
			SurfaceEvenness:       FormatData(sum.SurfaceEvenness),
			Euphotic_zoneShannon:  FormatData(sum.Euphotic_zoneShannon),
			Euphotic_zoneNumber:   FormatData(sum.Euphotic_zoneNumber),
			Euphotic_zoneMax:      FormatData(sum.Euphotic_zoneMax),
			Euphotic_zoneEvenness: FormatData(sum.Euphotic_zoneEvenness),
			AverageShannon:        FormatData(sum.AverageShannon),
			AverageNumber:         FormatData(sum.AverageNumber),
			AverageMax:            FormatData(sum.AverageMax),
			AverageEvenness:       FormatData(sum.AverageEvenness),
		}
	}
	return summaries, nil
}

func GetSummariesData() ([]SummaryData, error) {
	SummaryList := []SummaryData{}

	SummarySet := make(map[SummaryData]struct{})
	var exists = struct{}{}

	summaries, err := GetSummaryData()
	if err != nil {
		return SummaryList, err
	}

	for _, summary := range summaries {
		SummaryData := SummaryData{
			Year:                  summary.Year,
			Group:                 summary.Group,
			Identification:        summary.Identification,
			Asset:                 summary.Asset,
			Platform:              summary.Platform,
			Station:               summary.Station,
			SurfaceShannon:        summary.SurfaceShannon,
			SurfaceNumber:         summary.SurfaceNumber,
			SurfaceMax:            summary.SurfaceMax,
			SurfaceEvenness:       summary.SurfaceEvenness,
			Euphotic_zoneShannon:  summary.Euphotic_zoneShannon,
			Euphotic_zoneNumber:   summary.Euphotic_zoneNumber,
			Euphotic_zoneMax:      summary.Euphotic_zoneMax,
			Euphotic_zoneEvenness: summary.Euphotic_zoneEvenness,
			AverageShannon:        summary.AverageShannon,
			AverageNumber:         summary.AverageNumber,
			AverageMax:            summary.AverageMax,
			AverageEvenness:       summary.AverageEvenness,
		}
		SummarySet[SummaryData] = exists
	}

	for summaries := range SummarySet {
		SummaryList = append(SummaryList, summaries)
	}

	return SummaryList, nil
}

// //No.,Species,Year,Asset,Platform,Surface,Euphotic_zone
// PlatformDensity Import
type PlatformDensityData struct {
	Species       string `csv:"Species"`
	Year          string `csv:"Year"`
	Asset         string `csv:"Asset"`
	Platform      string `csv:"Platform"`
	Surface       string `csv:"Surface"`
	Euphotic_zone string `csv:"Euphotic_zone"`
}

func GetPlatformDensityData() ([]PlatformDensityData, error) {
	densities := []PlatformDensityData{}
	filename := "./data/datax - PlatformDensity.csv"

	in, err := os.Open(filename)
	if err != nil {
		return densities, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &densities); err != nil {
		return densities, err
	}

	for i, location := range densities {
		densities[i] = PlatformDensityData{
			Species:       FormatData(location.Species),
			Year:          FormatData(location.Year),
			Asset:         FormatData(location.Asset),
			Platform:      FormatData(location.Platform),
			Surface:       FormatData(location.Surface),
			Euphotic_zone: FormatData(location.Euphotic_zone),
		}
	}

	return densities, nil
}

func GetPlatformDensitiesData() ([]PlatformDensityData, error) {
	DensityList := []PlatformDensityData{}

	DensitySet := make(map[PlatformDensityData]struct{})
	var exists = struct{}{}

	densities, err := GetPlatformDensityData()
	if err != nil {
		return DensityList, err
	}

	for _, diversity := range densities {
		PlatformDensityData := PlatformDensityData{
			Species:       diversity.Species,
			Year:          diversity.Year,
			Asset:         diversity.Asset,
			Platform:      diversity.Platform,
			Surface:       diversity.Surface,
			Euphotic_zone: diversity.Euphotic_zone,
		}
		DensitySet[PlatformDensityData] = exists
	}

	for densities := range DensitySet {
		DensityList = append(DensityList, densities)
	}

	return DensityList, nil
}

// //No.,Species,Year,Asset,Surface,Euphotic_zone
// AssetDensity Import
type AssetDensityData struct {
	Species       string `csv:"Species"`
	Year          string `csv:"Year"`
	Asset         string `csv:"Asset"`
	Surface       string `csv:"Surface"`
	Euphotic_zone string `csv:"Euphotic_zone"`
}

func GetAssetDensityData() ([]AssetDensityData, error) {
	densities := []AssetDensityData{}
	filename := "./data/datax - AssetDensity.csv"

	in, err := os.Open(filename)
	if err != nil {
		return densities, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &densities); err != nil {
		return densities, err
	}

	for i, location := range densities {
		densities[i] = AssetDensityData{
			Species:       FormatData(location.Species),
			Year:          FormatData(location.Year),
			Asset:         FormatData(location.Asset),
			Surface:       FormatData(location.Surface),
			Euphotic_zone: FormatData(location.Euphotic_zone),
		}
	}

	return densities, nil
}

func GetAssetDensitiesData() ([]AssetDensityData, error) {
	DensityList := []AssetDensityData{}

	DensitySet := make(map[AssetDensityData]struct{})
	var exists = struct{}{}

	densities, err := GetAssetDensityData()
	if err != nil {
		return DensityList, err
	}

	for _, diversity := range densities {
		AssetDensityData := AssetDensityData{
			Species:       diversity.Species,
			Year:          diversity.Year,
			Asset:         diversity.Asset,
			Surface:       diversity.Surface,
			Euphotic_zone: diversity.Euphotic_zone,
		}
		DensitySet[AssetDensityData] = exists
	}

	for densities := range DensitySet {
		DensityList = append(DensityList, densities)
	}

	return DensityList, nil
}

// //Platform Summary Import   No.,Year,Group,Identification,Asset,Platform,Station,SurfaceShannon,SurfaceNumber,SurfaceMax,SurfaceEvenness,
// Euphotic_zoneShannon,Euphotic_zoneNumber,Euphotic_zoneMax,Euphotic_zoneEvenness
// AverageShannon,AverageNumber,AverageMax,AverageEvenness
type PlatformSummaryData struct {
	Year                  string `csv:"Year"`
	Group                 string `csv:"Group"`
	Identification        string `csv:"Identification"`
	Asset                 string `csv:"Asset"`
	Platform              string `csv:"Platform"`
	SurfaceShannon        string `csv:"SurfaceShannon"`
	SurfaceNumber         string `csv:"SurfaceNumber"`
	SurfaceMax            string `csv:"SurfaceMax"`
	SurfaceEvenness       string `csv:"SurfaceEvenness"`
	Euphotic_zoneShannon  string `csv:"Euphotic_zoneShannon"`
	Euphotic_zoneNumber   string `csv:"Euphotic_zoneNumber"`
	Euphotic_zoneMax      string `csv:"Euphotic_zoneMax"`
	Euphotic_zoneEvenness string `csv:"Euphotic_zoneEvenness"`
	AverageShannon        string `csv:"AverageShannon"`
	AverageNumber         string `csv:"AverageNumber"`
	AverageMax            string `csv:"AverageMax"`
	AverageEvenness       string `csv:"AverageEvenness"`
}

func GetPlatformSummaryData() ([]PlatformSummaryData, error) {
	summaries := []PlatformSummaryData{}
	filename := "./data/datax - PlatformData_Summary.csv"

	in, err := os.Open(filename)
	if err != nil {
		return summaries, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &summaries); err != nil {
		return summaries, err
	}

	for i, sum := range summaries {
		summaries[i] = PlatformSummaryData{
			Year:                  FormatData(sum.Year),
			Group:                 FormatData(sum.Group),
			Identification:        FormatData(sum.Identification),
			Asset:                 FormatData(sum.Asset),
			Platform:              FormatData(sum.Platform),
			SurfaceShannon:        FormatData(sum.SurfaceShannon),
			SurfaceNumber:         FormatData(sum.SurfaceNumber),
			SurfaceMax:            FormatData(sum.SurfaceMax),
			SurfaceEvenness:       FormatData(sum.SurfaceEvenness),
			Euphotic_zoneShannon:  FormatData(sum.Euphotic_zoneShannon),
			Euphotic_zoneNumber:   FormatData(sum.Euphotic_zoneNumber),
			Euphotic_zoneMax:      FormatData(sum.Euphotic_zoneMax),
			Euphotic_zoneEvenness: FormatData(sum.Euphotic_zoneEvenness),
			AverageShannon:        FormatData(sum.AverageShannon),
			AverageNumber:         FormatData(sum.AverageNumber),
			AverageMax:            FormatData(sum.AverageMax),
			AverageEvenness:       FormatData(sum.AverageEvenness),
		}
	}
	return summaries, nil
}

func GetPlatformSummariesData() ([]PlatformSummaryData, error) {
	SummaryList := []PlatformSummaryData{}

	SummarySet := make(map[PlatformSummaryData]struct{})
	var exists = struct{}{}

	summaries, err := GetPlatformSummaryData()
	if err != nil {
		return SummaryList, err
	}

	for _, summary := range summaries {
		PlatformSummaryData := PlatformSummaryData{
			Year:                  summary.Year,
			Group:                 summary.Group,
			Identification:        summary.Identification,
			Asset:                 summary.Asset,
			Platform:              summary.Platform,
			SurfaceShannon:        summary.SurfaceShannon,
			SurfaceNumber:         summary.SurfaceNumber,
			SurfaceMax:            summary.SurfaceMax,
			SurfaceEvenness:       summary.SurfaceEvenness,
			Euphotic_zoneShannon:  summary.Euphotic_zoneShannon,
			Euphotic_zoneNumber:   summary.Euphotic_zoneNumber,
			Euphotic_zoneMax:      summary.Euphotic_zoneMax,
			Euphotic_zoneEvenness: summary.Euphotic_zoneEvenness,
			AverageShannon:        summary.AverageShannon,
			AverageNumber:         summary.AverageNumber,
			AverageMax:            summary.AverageMax,
			AverageEvenness:       summary.AverageEvenness,
		}
		SummarySet[PlatformSummaryData] = exists
	}

	for summaries := range SummarySet {
		SummaryList = append(SummaryList, summaries)
	}

	return SummaryList, nil
}

// //AssetSummary Import   No.,Year,Group,Identification,Asset,Platform,Station,SurfaceShannon,SurfaceNumber,SurfaceMax,SurfaceEvenness,
// Euphotic_zoneShannon,Euphotic_zoneNumber,Euphotic_zoneMax,Euphotic_zoneEvenness
// AverageShannon,AverageNumber,AverageMax,AverageEvenness
type AssetSummaryData struct {
	Year                  string `csv:"Year"`
	Group                 string `csv:"Group"`
	Identification        string `csv:"Identification"`
	Asset                 string `csv:"Asset"`
	SurfaceShannon        string `csv:"SurfaceShannon"`
	SurfaceNumber         string `csv:"SurfaceNumber"`
	SurfaceMax            string `csv:"SurfaceMax"`
	SurfaceEvenness       string `csv:"SurfaceEvenness"`
	Euphotic_zoneShannon  string `csv:"Euphotic_zoneShannon"`
	Euphotic_zoneNumber   string `csv:"Euphotic_zoneNumber"`
	Euphotic_zoneMax      string `csv:"Euphotic_zoneMax"`
	Euphotic_zoneEvenness string `csv:"Euphotic_zoneEvenness"`
	AverageShannon        string `csv:"AverageShannon"`
	AverageNumber         string `csv:"AverageNumber"`
	AverageMax            string `csv:"AverageMax"`
	AverageEvenness       string `csv:"AverageEvenness"`
}

func GetAssetSummaryData() ([]AssetSummaryData, error) {
	summaries := []AssetSummaryData{}
	filename := "./data/datax - AssetData_Summary.csv"

	in, err := os.Open(filename)
	if err != nil {
		return summaries, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &summaries); err != nil {
		return summaries, err
	}

	for i, sum := range summaries {
		summaries[i] = AssetSummaryData{
			Year:                  FormatData(sum.Year),
			Group:                 FormatData(sum.Group),
			Identification:        FormatData(sum.Identification),
			Asset:                 FormatData(sum.Asset),
			SurfaceShannon:        FormatData(sum.SurfaceShannon),
			SurfaceNumber:         FormatData(sum.SurfaceNumber),
			SurfaceMax:            FormatData(sum.SurfaceMax),
			SurfaceEvenness:       FormatData(sum.SurfaceEvenness),
			Euphotic_zoneShannon:  FormatData(sum.Euphotic_zoneShannon),
			Euphotic_zoneNumber:   FormatData(sum.Euphotic_zoneNumber),
			Euphotic_zoneMax:      FormatData(sum.Euphotic_zoneMax),
			Euphotic_zoneEvenness: FormatData(sum.Euphotic_zoneEvenness),
			AverageShannon:        FormatData(sum.AverageShannon),
			AverageNumber:         FormatData(sum.AverageNumber),
			AverageMax:            FormatData(sum.AverageMax),
			AverageEvenness:       FormatData(sum.AverageEvenness),
		}
	}
	return summaries, nil
}

func GetAssetSummariesData() ([]AssetSummaryData, error) {
	SummaryList := []AssetSummaryData{}

	SummarySet := make(map[AssetSummaryData]struct{})
	var exists = struct{}{}

	summaries, err := GetAssetSummaryData()
	if err != nil {
		return SummaryList, err
	}

	for _, summary := range summaries {
		AssetSummaryData := AssetSummaryData{
			Year:                  summary.Year,
			Group:                 summary.Group,
			Identification:        summary.Identification,
			Asset:                 summary.Asset,
			SurfaceShannon:        summary.SurfaceShannon,
			SurfaceNumber:         summary.SurfaceNumber,
			SurfaceMax:            summary.SurfaceMax,
			SurfaceEvenness:       summary.SurfaceEvenness,
			Euphotic_zoneShannon:  summary.Euphotic_zoneShannon,
			Euphotic_zoneNumber:   summary.Euphotic_zoneNumber,
			Euphotic_zoneMax:      summary.Euphotic_zoneMax,
			Euphotic_zoneEvenness: summary.Euphotic_zoneEvenness,
			AverageShannon:        summary.AverageShannon,
			AverageNumber:         summary.AverageNumber,
			AverageMax:            summary.AverageMax,
			AverageEvenness:       summary.AverageEvenness,
		}
		SummarySet[AssetSummaryData] = exists
	}

	for summaries := range SummarySet {
		SummaryList = append(SummaryList, summaries)
	}

	return SummaryList, nil
}

type UserData struct {
	Email    string `csv:"email"`
	Password string `csv:"password"`
}

func GetUserData() ([]UserData, error) {
	users := []UserData{}
	filename := "./data/datax - Username.csv"

	in, err := os.Open(filename)
	if err != nil {
		return users, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &users); err != nil {
		return users, err
	}

	for i, userdat := range users {
		users[i] = UserData{
			Email:    FormatData(userdat.Email),
			Password: userdat.Password,
		}
		//fmt.Println(FormatData(userdat.Password))
	}
	return users, nil
}

type AssetLocationData struct {
	Asset     string `csv:"Asset"`
	Latitude  string `csv:"Latitude"`
	Longitude string `csv:"Longitude"`
}

func GetAssetLocationData() ([]LocationData, error) {
	locations := []LocationData{}
	filename := "./data/datax - Location_Asset.csv"

	in, err := os.Open(filename)
	if err != nil {
		return locations, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &locations); err != nil {
		return locations, err
	}

	for i, location := range locations {
		locations[i] = LocationData{
			Asset:     FormatData(location.Asset),
			Latitude:  FormatData(location.Latitude),
			Longitude: FormatData(location.Longitude),
			Type:      FormatData(location.Type),
		}
	}

	return locations, nil
}

type PlatformLocationData struct {
	Asset     string `csv:"Asset"`
	Platform  string `csv:"Platform"`
	Latitude  string `csv:"Latitude"`
	Longitude string `csv:"Longitude"`
	Type      string `csv:"Type"`
}

func GetPlatformLocationData() ([]LocationData, error) {
	locations := []LocationData{}
	filename := "./data/datax - Location_Platform.csv"

	in, err := os.Open(filename)
	if err != nil {
		return locations, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &locations); err != nil {
		return locations, err
	}

	for i, location := range locations {
		locations[i] = LocationData{
			Asset:     FormatData(location.Asset),
			Platform:  FormatData(location.Platform),
			Latitude:  FormatData(location.Latitude),
			Longitude: FormatData(location.Longitude),
			Type:      FormatData(location.Type),
		}
	}

	return locations, nil
}
