package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableDensitySQL = `
	CREATE TABLE IF NOT EXISTS public."density"
	(
		"density_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"density_year" varchar(500) NOT NULL,

		"project_id" UUID,
		"platform_id" UUID,
		"station_id" UUID,
		"identification_id" UUID,
		"method_id" UUID,
		
		"major_group_id" UUID,
		"kingdom_id" UUID,
		"sub_kingdom_id" UUID,
		"infra_kingdom_id" UUID,
		"super_phylum_id" UUID,
		"phylum_id" UUID,
		"sub_phylum_id" UUID,
		"infra_phylum_id" UUID,
		"super_class_id" UUID,
		"class_id" UUID,
		"sub_class_id" UUID,
		"infra_class_id" UUID,
		"super_order_id" UUID,
		"order_id" UUID,
		"sub_order_id" UUID,
		"infra_order_id" UUID,
		"super_family_id" UUID,
		"family_id" UUID,
		"sub_family_id" UUID,
		"super_genus_id" UUID,
		"genus_id" UUID,
		"sub_genus_id" UUID,
		"species_id" UUID,
		"sub_species_id" UUID,
		"density" float,
		
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableDensitySQL = `DROP TABLE IF EXISTS public."density";`

const addFKDensityProjectSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_project
	FOREIGN KEY ("project_id") REFERENCES public."project" ("project_id")
	;`
const dropFKDensityProjectSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_project
	;`

const addFKDensityPlatformSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKDensityPlatformSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_platform
	;`

const addFKDensityStationSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_station
	FOREIGN KEY ("station_id") REFERENCES public."station" ("station_id")
	;`
const dropFKDensityStationSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_station
	;`

const addFKDensityIdentificationSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_identification
	FOREIGN KEY ("identification_id") REFERENCES public."identification" ("identification_id")
	;`
const dropFKDensityIdentificationSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_fk_density_identification
	;`

const addFKDensityMethodSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_method
	FOREIGN KEY ("method_id") REFERENCES public."method" ("method_id")
	;`
const dropFKDensityMethodSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_method
	;`

const addFKDensityMajorGroupSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_major_group
	FOREIGN KEY ("major_group_id") REFERENCES public."major_group" ("major_group_id")
	;`
const dropFKDensityMajorGroupSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_major_group
	;`

const addFKDensityKingdomSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_kingdom
	FOREIGN KEY ("kingdom_id") REFERENCES public."kingdom" ("kingdom_id")
	;`
const dropFKDensityKingdomSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_kingdom
	;`

const addFKDensitySubKingdomSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_sub_kingdom
	FOREIGN KEY ("sub_kingdom_id") REFERENCES public."sub_kingdom" ("sub_kingdom_id")
	;`
const dropFKDensitySubKingdomSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_sub_kingdom
	;`

const addFKDensityInfraKingdomSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_infra_kingdom
	FOREIGN KEY ("infra_kingdom_id") REFERENCES public."infra_kingdom" ("infra_kingdom_id")
	;`
const dropFKDensityInfraKingdomSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_infra_kingdom
	;`

const addFKDensitySuperPhylumSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_super_phylum
	FOREIGN KEY ("super_phylum_id") REFERENCES public."super_phylum" ("super_phylum_id")
	;`
const dropFKDensitySuperPhylumSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_super_phylum
	;`

const addFKDensityPhylumSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_phylum
	FOREIGN KEY ("phylum_id") REFERENCES public."phylum" ("phylum_id")
	;`
const dropFKDensityPhylumSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_phylum
	;`

const addFKDensitySubPhylumSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_sub_phylum
	FOREIGN KEY ("sub_phylum_id") REFERENCES public."sub_phylum" ("sub_phylum_id")
	;`
const dropFKDensitySubPhylumSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_sub_phylum
	;`

const addFKDensityInfraPhylumSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_infra_phylum
	FOREIGN KEY ("infra_phylum_id") REFERENCES public."infra_phylum" ("infra_phylum_id")
	;`
const dropFKDensityInfraPhylumSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_infra_phylum
	;`

const addFKDensitySuperClassSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_super_class
	FOREIGN KEY ("super_class_id") REFERENCES public."super_class" ("super_class_id")
	;`
const dropFKDensitySuperClassSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_super_class
	;`

const addFKDensityClassSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_class
	FOREIGN KEY ("class_id") REFERENCES public."class" ("class_id")
	;`
const dropFKDensityClassSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_class
	;`

const addFKDensitySubClassSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_sub_class
	FOREIGN KEY ("sub_class_id") REFERENCES public."sub_class" ("sub_class_id")
	;`
const dropFKDensitySubClassSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_sub_class
	;`

const addFKDensityInfraClassSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_infra_class
	FOREIGN KEY ("infra_class_id") REFERENCES public."infra_class" ("infra_class_id")
	;`
const dropFKDensityInfraClassSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_infra_class
	;`

const addFKDensitySuperOrderSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_super_order
	FOREIGN KEY ("super_order_id") REFERENCES public."super_order" ("super_order_id")
	;`
const dropFKDensitySuperOrderSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_super_order
	;`

const addFKDensityOrderSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_order
	FOREIGN KEY ("order_id") REFERENCES public."order" ("order_id")
	;`
const dropFKDensityOrderSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_order
	;`

const addFKDensitySubOrderSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_sub_order
	FOREIGN KEY ("sub_order_id") REFERENCES public."sub_order" ("sub_order_id")
	;`
const dropFKDensitySubOrderSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_sub_order
	;`

const addFKDensityInfraOrderSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_infra_order
	FOREIGN KEY ("infra_order_id") REFERENCES public."infra_order" ("infra_order_id")
	;`
const dropFKDensityInfraOrderSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_infra_order
	;`

const addFKDensitySuperFamilySQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_super_family
	FOREIGN KEY ("super_family_id") REFERENCES public."super_family" ("super_family_id")
	;`
const dropFKDensitySuperFamilySQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_super_family
	;`

const addFKDensityFamilySQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_family
	FOREIGN KEY ("family_id") REFERENCES public."family" ("family_id")
	;`
const dropFKDensityFamilySQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_family
	;`

const addFKDensitySubFamilySQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_sub_family
	FOREIGN KEY ("sub_family_id") REFERENCES public."sub_family" ("sub_family_id")
	;`
const dropFKDensitySubFamilySQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_sub_family
	;`

const addFKDensitySuperGenusSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_super_genus
	FOREIGN KEY ("super_genus_id") REFERENCES public."super_genus" ("super_genus_id")
	;`
const dropFKDensitySuperGenusSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_super_genus
	;`

const addFKDensityGenusSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_genus
	FOREIGN KEY ("genus_id") REFERENCES public."genus" ("genus_id")
	;`
const dropFKDensityGenusSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_genus
	;`

const addFKDensitySubGenusSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_sub_genus
	FOREIGN KEY ("sub_genus_id") REFERENCES public."sub_genus" ("sub_genus_id")
	;`
const dropFKDensitySubGenusSQL = `
	ALTER TABLE public."density"
	DROP CONSTRAINT fk_density_sub_genus
	;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table density...")
		var scripts = [28]string{
			createTableDensitySQL,
			addFKDensityProjectSQL,
			addFKDensityPlatformSQL,
			addFKDensityStationSQL,
			addFKDensityIdentificationSQL,
			addFKDensityMethodSQL,
			addFKDensityMajorGroupSQL,
			addFKDensityKingdomSQL,
			addFKDensitySubKingdomSQL,
			addFKDensityInfraKingdomSQL,
			addFKDensitySuperPhylumSQL,
			addFKDensityPhylumSQL,
			addFKDensitySubPhylumSQL,
			addFKDensityInfraPhylumSQL,
			addFKDensitySuperClassSQL,
			addFKDensityClassSQL,
			addFKDensitySubClassSQL,
			addFKDensityInfraClassSQL,
			addFKDensitySuperOrderSQL,
			addFKDensityOrderSQL,
			addFKDensitySubOrderSQL,
			addFKDensityInfraOrderSQL,
			addFKDensitySuperFamilySQL,
			addFKDensityFamilySQL,
			addFKDensitySubFamilySQL,
			addFKDensitySuperGenusSQL,
			addFKDensityGenusSQL,
			addFKDensitySubGenusSQL,
		}
		var firstError error
		for _, script := range scripts {
			_, err := db.Exec(script)
			if err != nil {
				firstError = err
				break
			}
		}
		if firstError != nil {
			return firstError
		}

		fmt.Println("[Migration] Seeding table density...")
		densitys, err := GetDensityData()
		if err != nil {
			return err
		}
		for _, s := range densitys {

			if s.Density == "NULL" {
				continue
			}

			insertDensitySQL := fmt.Sprintf(`
			INSERT INTO public."density" (
				"density_year",
				"project_id",
				"platform_id",
				"station_id",
				"identification_id",
				"method_id",
				"major_group_id",
				"kingdom_id",
				"sub_kingdom_id",
				"infra_kingdom_id",
				"super_phylum_id",
				"phylum_id",
				"sub_phylum_id",
				"infra_phylum_id",
				"super_class_id",
				"class_id",
				"sub_class_id",
				"infra_class_id",
				"super_order_id",
				"order_id",
				"sub_order_id",
				"infra_order_id",
				"super_family_id",
				"family_id",
				"sub_family_id",
				"super_genus_id",
				"genus_id",
				"sub_genus_id",
				"species_id",
				"density"
				) VALUES ('%s',
			(SELECT "project_id" FROM public."project" WHERE "project_name" = '%s' LIMIT 1),
			(SELECT "platform_id" FROM public."platform" WHERE "platform_name" = '%s' LIMIT 1),
			(SELECT "station_id" FROM public."station" WHERE "station_name" = '%s' LIMIT 1),
			(SELECT "identification_id" FROM public."identification" WHERE "identification_name" = '%s' LIMIT 1),
			(SELECT "method_id" FROM public."method" WHERE "method_name" = '%s' LIMIT 1),
			(SELECT "major_group_id" FROM public."major_group" WHERE "major_group_name" = '%s' LIMIT 1),
			(SELECT "kingdom_id" FROM public."kingdom" WHERE "kingdom_name" = '%s' LIMIT 1),
			(SELECT "sub_kingdom_id" FROM public."sub_kingdom" WHERE "sub_kingdom_name" = '%s' LIMIT 1),
			(SELECT "infra_kingdom_id" FROM public."infra_kingdom" WHERE "infra_kingdom_name" = '%s' LIMIT 1),
			(SELECT "super_phylum_id" FROM public."super_phylum" WHERE "super_phylum_name" = '%s' LIMIT 1),
			(SELECT "phylum_id" FROM public."phylum" WHERE "phylum_name" = '%s' LIMIT 1),
			(SELECT "sub_phylum_id" FROM public."sub_phylum" WHERE "sub_phylum_name" = '%s' LIMIT 1),
			(SELECT "infra_phylum_id" FROM public."infra_phylum" WHERE "infra_phylum_name" = '%s' LIMIT 1),
			(SELECT "super_class_id" FROM public."super_class" WHERE "super_class_name" = '%s' LIMIT 1),
			(SELECT "class_id" FROM public."class" WHERE "class_name" = '%s' LIMIT 1),
			(SELECT "sub_class_id" FROM public."sub_class" WHERE "sub_class_name" = '%s' LIMIT 1),
			(SELECT "infra_class_id" FROM public."infra_class" WHERE "infra_class_name" = '%s' LIMIT 1),
			(SELECT "super_order_id" FROM public."super_order" WHERE "super_order_name" = '%s' LIMIT 1),
			(SELECT "order_id" FROM public."order" WHERE "order_name" = '%s' LIMIT 1),
			(SELECT "sub_order_id" FROM public."sub_order" WHERE "sub_order_name" = '%s' LIMIT 1),
			(SELECT "infra_order_id" FROM public."infra_order" WHERE "infra_order_name" = '%s' LIMIT 1),
			(SELECT "super_family_id" FROM public."super_family" WHERE "super_family_name" = '%s' LIMIT 1),
			(SELECT "family_id" FROM public."family" WHERE "family_name" = '%s' LIMIT 1),
			(SELECT "sub_family_id" FROM public."sub_family" WHERE "sub_family_name" = '%s' LIMIT 1),
			(SELECT "super_genus_id" FROM public."super_genus" WHERE "super_genus_name" = '%s' LIMIT 1),
			(SELECT "genus_id" FROM public."genus" WHERE "genus_name" = '%s' LIMIT 1),
			(SELECT "sub_genus_id" FROM public."sub_genus" WHERE "sub_genus_name" = '%s' LIMIT 1),
			(SELECT "species_id" FROM public."species" WHERE "species_name" = '%s' LIMIT 1),
			'%s');
			`,
				s.Year,
				s.Project,
				s.Platform,
				s.Station,
				s.IdentificationTechnique,
				s.MethodOfCollection,
				s.MajorGroup,
				s.Kingdom,
				s.SubKingdom,
				s.InfraKingdom,
				s.SuperPhylum,
				s.Phylum,
				s.SubPhylum,
				s.InfraPhylum,
				s.SuperClass,
				s.Class,
				s.SubClass,
				s.InfraClass,
				s.SuperOrder,
				s.Order,
				s.SubOrder,
				s.InfraOrder,
				s.SuperFamily,
				s.Family,
				s.SubFamily,
				s.SuperGenus,
				s.Genus,
				s.SubGenus,
				s.SpeciesName,
				s.Density,
			)
			_, err := db.Exec(insertDensitySQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table density...")
		var scripts = [28]string{
			dropFKDensitySubGenusSQL,
			dropFKDensityGenusSQL,
			dropFKDensitySuperGenusSQL,
			dropFKDensitySubFamilySQL,
			dropFKDensityFamilySQL,
			dropFKDensitySuperFamilySQL,
			dropFKDensityInfraOrderSQL,
			dropFKDensitySubOrderSQL,
			dropFKDensityOrderSQL,
			dropFKDensitySuperOrderSQL,
			dropFKDensityInfraClassSQL,
			dropFKDensitySubClassSQL,
			dropFKDensityClassSQL,
			dropFKDensitySuperClassSQL,
			dropFKDensityInfraPhylumSQL,
			dropFKDensitySubPhylumSQL,
			dropFKDensityPhylumSQL,
			dropFKDensitySuperPhylumSQL,
			dropFKDensityInfraKingdomSQL,
			dropFKDensitySubKingdomSQL,
			dropFKDensityKingdomSQL,
			dropFKDensityMajorGroupSQL,
			dropFKDensityMethodSQL,
			dropFKDensityIdentificationSQL,
			dropFKDensityStationSQL,
			dropFKDensityPlatformSQL,
			dropFKDensityProjectSQL,
			dropTableDensitySQL,
		}
		var firstError error
		for _, script := range scripts {
			_, err := db.Exec(script)
			if err != nil {
				firstError = err
				break
			}
		}
		return firstError
	})
}
