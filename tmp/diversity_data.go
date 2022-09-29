package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

type DiverData struct {
	Year      string `csv:"year"`
	Project   string `csv:"project"`
	Platform  string `csv:"platform"`
	Station   string `csv:"station"`
	Diversity string `csv:"diversity"`
}

type DiversityData struct {
	Year      string
	Project   string
	Platform  string
	Station   string
	Diversity string
}

func GetDiverData() ([]DiverData, error) {
	divers := []DiverData{}
	filename := "./data/diversitymockup.csv"

	in, err := os.Open(filename)
	if err != nil {
		return divers, err
	}

	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &divers); err != nil {
		return divers, err
	}

	for i, diver := range divers {
		divers[i] = DiverData{
			Year:      FormatData(diver.Year),
			Project:   FormatData(diver.Project),
			Station:   FormatData(diver.Station),
			Platform:  FormatData(diver.Platform),
			Diversity: FormatData(diver.Diversity),
		}
	}

	return divers, nil
}

func GetDiverProjectData() ([]ProjectData, error) {
	ProjectList := []ProjectData{}

	ProjectSet := make(map[ProjectData]struct{})
	var exists = struct{}{}

	divers, err := GetDiverData()
	if err != nil {
		return ProjectList, err
	}

	for _, diver := range divers {
		ProjectData := ProjectData{
			Project: diver.Project,
		}
		ProjectSet[ProjectData] = exists
	}

	for Projects := range ProjectSet {
		ProjectList = append(ProjectList, Projects)
	}

	return ProjectList, nil
}

func GetDiverPlatformData() ([]PlatformData, error) {
	PlatformList := []PlatformData{}

	PlatformSet := make(map[PlatformData]struct{})
	var exists = struct{}{}

	divers, err := GetDiverData()
	if err != nil {
		return PlatformList, err
	}

	for _, diver := range divers {
		PlatformData := PlatformData{
			Platform: diver.Platform,
			Project:  diver.Project,
		}
		PlatformSet[PlatformData] = exists
	}

	for Platform := range PlatformSet {
		PlatformList = append(PlatformList, Platform)
	}

	return PlatformList, nil
}

func GetDiverStationData() ([]StationData, error) {
	StationList := []StationData{}

	StationSet := make(map[StationData]struct{})
	var exists = struct{}{}

	divers, err := GetDiverData()
	if err != nil {
		return StationList, err
	}

	for _, diver := range divers {
		StationData := StationData{
			Station:  diver.Station,
			Platform: diver.Platform,
		}
		StationSet[StationData] = exists
	}

	for Station := range StationSet {
		StationList = append(StationList, Station)
	}

	return StationList, nil
}

func GetDiversityData() ([]DiversityData, error) {
	DiversityList := []DiversityData{}

	DiversitySet := make(map[DiversityData]struct{})
	var exists = struct{}{}

	divers, err := GetDiverData()
	if err != nil {
		return DiversityList, err
	}

	for _, diver := range divers {
		diversityData := DiversityData{
			Year:      diver.Year,
			Project:   diver.Project,
			Platform:  diver.Platform,
			Station:   diver.Station,
			Diversity: diver.Diversity,
		}
		DiversitySet[diversityData] = exists
	}

	for diversities := range DiversitySet {
		DiversityList = append(DiversityList, diversities)
	}

	return DiversityList, nil
}
