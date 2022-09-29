package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSpeciesSQL = `
	CREATE TABLE IF NOT EXISTS public."species"
	(
		"species_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"species_name" varchar(500) UNIQUE NOT NULL,
		"scientific_name" varchar(500),
		"common_name_en" varchar(500),
		"common_name_th" varchar(500),
		"synonym" varchar(500),
		"dna" varchar(500),

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

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSpeciesSQL = `DROP TABLE IF EXISTS public."species";`

const addFKSpeciesIdentificationSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_identification
	FOREIGN KEY ("identification_id") REFERENCES public."identification" ("identification_id")
	;`
const dropFKSpeciesIdentificationSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_identification
	;`
const addFKSpeciesMethodSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_method
	FOREIGN KEY ("method_id") REFERENCES public."method" ("method_id")
	;`
const dropFKSpeciesMethodSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_method
	;`

const addFKSpeciesMajorGroupSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_major_group
	FOREIGN KEY ("major_group_id") REFERENCES public."major_group" ("major_group_id")
	;`
const dropFKSpeciesMajorGroupSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_major_group
	;`

const addFKSpeciesKingdomSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_kingdom
	FOREIGN KEY ("kingdom_id") REFERENCES public."kingdom" ("kingdom_id")
	;`
const dropFKSpeciesKingdomSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_kingdom
	;`

const addFKSpeciesSubKingdomSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_sub_kingdom
	FOREIGN KEY ("sub_kingdom_id") REFERENCES public."sub_kingdom" ("sub_kingdom_id")
	;`
const dropFKSpeciesSubKingdomSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_sub_kingdom
	;`

const addFKSpeciesInfraKingdomSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_infra_kingdom
	FOREIGN KEY ("infra_kingdom_id") REFERENCES public."infra_kingdom" ("infra_kingdom_id")
	;`
const dropFKSpeciesInfraKingdomSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_infra_kingdom
	;`

const addFKSpeciesSuperPhylumSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_super_phylum
	FOREIGN KEY ("super_phylum_id") REFERENCES public."super_phylum" ("super_phylum_id")
	;`
const dropFKSpeciesSuperPhylumSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_super_phylum
	;`

const addFKSpeciesPhylumSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_phylum
	FOREIGN KEY ("phylum_id") REFERENCES public."phylum" ("phylum_id")
	;`
const dropFKSpeciesPhylumSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_phylum
	;`

const addFKSpeciesSubPhylumSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_sub_phylum
	FOREIGN KEY ("sub_phylum_id") REFERENCES public."sub_phylum" ("sub_phylum_id")
	;`
const dropFKSpeciesSubPhylumSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_sub_phylum
	;`

const addFKSpeciesInfraPhylumSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_infra_phylum
	FOREIGN KEY ("infra_phylum_id") REFERENCES public."infra_phylum" ("infra_phylum_id")
	;`
const dropFKSpeciesInfraPhylumSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_infra_phylum
	;`

const addFKSpeciesSuperClassSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_super_class
	FOREIGN KEY ("super_class_id") REFERENCES public."super_class" ("super_class_id")
	;`
const dropFKSpeciesSuperClassSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_super_class
	;`

const addFKSpeciesClassSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_class
	FOREIGN KEY ("class_id") REFERENCES public."class" ("class_id")
	;`
const dropFKSpeciesClassSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_class
	;`

const addFKSpeciesSubClassSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_sub_class
	FOREIGN KEY ("sub_class_id") REFERENCES public."sub_class" ("sub_class_id")
	;`
const dropFKSpeciesSubClassSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_sub_class
	;`

const addFKSpeciesInfraClassSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_infra_class
	FOREIGN KEY ("infra_class_id") REFERENCES public."infra_class" ("infra_class_id")
	;`
const dropFKSpeciesInfraClassSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_infra_class
	;`

const addFKSpeciesSuperOrderSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_super_order
	FOREIGN KEY ("super_order_id") REFERENCES public."super_order" ("super_order_id")
	;`
const dropFKSpeciesSuperOrderSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_super_order
	;`

const addFKSpeciesOrderSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_order
	FOREIGN KEY ("order_id") REFERENCES public."order" ("order_id")
	;`
const dropFKSpeciesOrderSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_order
	;`

const addFKSpeciesSubOrderSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_sub_order
	FOREIGN KEY ("sub_order_id") REFERENCES public."sub_order" ("sub_order_id")
	;`
const dropFKSpeciesSubOrderSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_sub_order
	;`

const addFKSpeciesInfraOrderSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_infra_order
	FOREIGN KEY ("infra_order_id") REFERENCES public."infra_order" ("infra_order_id")
	;`
const dropFKSpeciesInfraOrderSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_infra_order
	;`

const addFKSpeciesSuperFamilySQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_super_family
	FOREIGN KEY ("super_family_id") REFERENCES public."super_family" ("super_family_id")
	;`
const dropFKSpeciesSuperFamilySQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_super_family
	;`

const addFKSpeciesFamilySQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_family
	FOREIGN KEY ("family_id") REFERENCES public."family" ("family_id")
	;`
const dropFKSpeciesFamilySQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_family
	;`

const addFKSpeciesSubFamilySQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_sub_family
	FOREIGN KEY ("sub_family_id") REFERENCES public."sub_family" ("sub_family_id")
	;`
const dropFKSpeciesSubFamilySQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_sub_family
	;`

const addFKSpeciesSuperGenusSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_super_genus
	FOREIGN KEY ("super_genus_id") REFERENCES public."super_genus" ("super_genus_id")
	;`
const dropFKSpeciesSuperGenusSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_super_genus
	;`

const addFKSpeciesGenusSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_genus
	FOREIGN KEY ("genus_id") REFERENCES public."genus" ("genus_id")
	;`
const dropFKSpeciesGenusSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_genus
	;`

const addFKSpeciesSubGenusSQL = `
	ALTER TABLE public."species"
	ADD CONSTRAINT fk_species_sub_genus
	FOREIGN KEY ("sub_genus_id") REFERENCES public."sub_genus" ("sub_genus_id")
	;`
const dropFKSpeciesSubGenusSQL = `
	ALTER TABLE public."species"
	DROP CONSTRAINT fk_species_sub_genus
	;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table species...")
		var scripts = [25]string{
			createTableSpeciesSQL,
			addFKSpeciesIdentificationSQL,
			addFKSpeciesMethodSQL,
			addFKSpeciesMajorGroupSQL,
			addFKSpeciesKingdomSQL,
			addFKSpeciesSubKingdomSQL,
			addFKSpeciesInfraKingdomSQL,
			addFKSpeciesSuperPhylumSQL,
			addFKSpeciesPhylumSQL,
			addFKSpeciesSubPhylumSQL,
			addFKSpeciesInfraPhylumSQL,
			addFKSpeciesSuperClassSQL,
			addFKSpeciesClassSQL,
			addFKSpeciesSubClassSQL,
			addFKSpeciesInfraClassSQL,
			addFKSpeciesSuperOrderSQL,
			addFKSpeciesOrderSQL,
			addFKSpeciesSubOrderSQL,
			addFKSpeciesInfraOrderSQL,
			addFKSpeciesSuperFamilySQL,
			addFKSpeciesFamilySQL,
			addFKSpeciesSubFamilySQL,
			addFKSpeciesSuperGenusSQL,
			addFKSpeciesGenusSQL,
			addFKSpeciesSubGenusSQL,
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

		fmt.Println("[Migration] Seeding table species...")
		species, err := GetSpeciesData()
		if err != nil {
			return err
		}
		for _, s := range species {

			if s.SpeciesName == "NULL" {
				continue
			}

			insertSpeciesSQL := fmt.Sprintf(`
			INSERT INTO public."species" (
				"species_name",
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
				"sub_genus_id"
				) VALUES ('%s',
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
			(SELECT "sub_genus_id" FROM public."sub_genus" WHERE "sub_genus_name" = '%s' LIMIT 1));
			`, s.SpeciesName, s.IdentificationName, s.MethodName, s.MajorGroupName, s.KingdomName, s.SubKingdomName,
				s.InfraKingdomName, s.SuperPhylumName, s.PhylumName, s.SubPhylumName,
				s.InfraPhylumName, s.SuperClassName, s.ClassName, s.SubClassName,
				s.InfraClassName, s.SuperOrderName, s.OrderName, s.SubOrderName,
				s.InfraOrderName, s.SuperFamilyName, s.FamilyName, s.SubFamilyName,
				s.SuperGenusName, s.GenusName, s.SubGenusName,
			)
			_, err := db.Exec(insertSpeciesSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table species...")
		var scripts = [25]string{
			dropFKSpeciesSubGenusSQL,
			dropFKSpeciesGenusSQL,
			dropFKSpeciesSuperGenusSQL,
			dropFKSpeciesSubFamilySQL,
			dropFKSpeciesFamilySQL,
			dropFKSpeciesSuperFamilySQL,
			dropFKSpeciesInfraOrderSQL,
			dropFKSpeciesSubOrderSQL,
			dropFKSpeciesOrderSQL,
			dropFKSpeciesSuperOrderSQL,
			dropFKSpeciesInfraClassSQL,
			dropFKSpeciesSubClassSQL,
			dropFKSpeciesClassSQL,
			dropFKSpeciesSuperClassSQL,
			dropFKSpeciesInfraPhylumSQL,
			dropFKSpeciesSubPhylumSQL,
			dropFKSpeciesPhylumSQL,
			dropFKSpeciesSuperPhylumSQL,
			dropFKSpeciesInfraKingdomSQL,
			dropFKSpeciesSubKingdomSQL,
			dropFKSpeciesKingdomSQL,
			dropFKSpeciesMajorGroupSQL,
			dropFKSpeciesMethodSQL,
			dropFKSpeciesIdentificationSQL,
			dropTableSpeciesSQL,
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
