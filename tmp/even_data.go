package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

type EvenData struct {
	Year     string `csv:"year"`
	Project  string `csv:"project"`
	Platform string `csv:"platform"`
	Station  string `csv:"station"`
	Surface  string `csv:"surface"`
	Bottom   string `csv:"bottom"`
	Evenness string `csv:"evenness"`
}

type EvennessData struct {
	Year     string
	Project  string
	Platform string
	Station  string
	Surface  string
	Bottom   string
	Evenness string
}

func GetEvenData() ([]EvenData, error) {
	evens := []EvenData{}
	filename := "./data/evenmockup.csv"

	in, err := os.Open(filename)
	if err != nil {
		return evens, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &evens); err != nil {
		return evens, err
	}

	for i, even := range evens {
		evens[i] = EvenData{
			Year:     FormatData(even.Year),
			Project:  FormatData(even.Project),
			Station:  FormatData(even.Station),
			Platform: FormatData(even.Platform),
			Surface:  FormatData(even.Surface),
			Bottom:   FormatData(even.Bottom),
			Evenness: FormatData(even.Evenness),
		}
	}

	return evens, nil
}

func GetEvenProjectData() ([]ProjectData, error) {
	ProjectList := []ProjectData{}

	ProjectSet := make(map[ProjectData]struct{})
	var exists = struct{}{}

	evens, err := GetEvenData()
	if err != nil {
		return ProjectList, err
	}

	for _, even := range evens {
		ProjectData := ProjectData{
			Project: even.Project,
		}
		ProjectSet[ProjectData] = exists
	}

	for Projects := range ProjectSet {
		ProjectList = append(ProjectList, Projects)
	}

	return ProjectList, nil
}

func GetEvenPlatformData() ([]PlatformData, error) {
	PlatformList := []PlatformData{}

	PlatformSet := make(map[PlatformData]struct{})
	var exists = struct{}{}

	evens, err := GetEvenData()
	if err != nil {
		return PlatformList, err
	}

	for _, even := range evens {
		PlatformData := PlatformData{
			Platform: even.Platform,
			Project:  even.Project,
		}
		PlatformSet[PlatformData] = exists
	}

	for Platform := range PlatformSet {
		PlatformList = append(PlatformList, Platform)
	}

	return PlatformList, nil
}

func GetEvenStationData() ([]StationData, error) {
	StationList := []StationData{}

	StationSet := make(map[StationData]struct{})
	var exists = struct{}{}

	evens, err := GetEvenData()
	if err != nil {
		return StationList, err
	}

	for _, even := range evens {
		StationData := StationData{
			Station:  even.Station,
			Platform: even.Platform,
		}
		StationSet[StationData] = exists
	}

	for Station := range StationSet {
		StationList = append(StationList, Station)
	}

	return StationList, nil
}

func GetEvennessData() ([]EvennessData, error) {
	EvennessList := []EvennessData{}

	EvennessSet := make(map[EvennessData]struct{})
	var exists = struct{}{}

	evens, err := GetEvenData()
	if err != nil {
		return EvennessList, err
	}

	for _, even := range evens {
		evennessData := EvennessData{
			Year:     even.Year,
			Project:  even.Project,
			Platform: even.Platform,
			Station:  even.Station,
			Surface:  even.Surface,
			Bottom:   even.Bottom,
			Evenness: even.Evenness,
		}
		EvennessSet[evennessData] = exists
	}

	for evenesses := range EvennessSet {
		EvennessList = append(EvennessList, evenesses)
	}

	return EvennessList, nil
}
