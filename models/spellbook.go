package models

var m map[string]SpellBook

const BARD = "bard"
const CLERIC = "cleric"
const DRUID = "druid"
const WIZARD = "wizard"
const RANGER = "ranger"
const PALADIN = "paladin"

type Spell struct {
	Title string
	School string
	Save string
	Range string
	Description string
}

type LevelSection struct {
	Spells []Spell
}

type SpellBook struct {
	Sections []LevelSection
}

func GetSpellBookMap() map[string]SpellBook{
	m = make(map[string]SpellBook)

	//Bard Spells
	book := SpellBook{Sections: []LevelSection{
		//0th lvl spells
		LevelSection{Spells: []Spell{
			Spell{
				Title: "DancingLights",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "Creates torches or other lights",
			},
			Spell{
				Title: "Daze",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "Humanoid creature of 4 HD or less loses next action",
			},
		}},
		//1st lvl spells
		LevelSection{Spells: []Spell{
			Spell{
				Title: "Alarm",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "Wards an area for 2 hours/level",
			},
			Spell{
				Title: "Animate Rope",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "Makes a rope move at your command",
			},
		}},
	}}
	m[BARD] = book

	//Cleric Spells
	book = SpellBook{Sections: []LevelSection{
		//0th lvl spells
		LevelSection{Spells: []Spell{
			Spell{
				Title: "CureMinorWounds",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "You heal 1hp to yourself or a willing subject within touch range",
			},
			Spell{
				Title: "CreateWater",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "Creates 2 gallons/level of pure water",
			},
		}},
		//1st lvl spells
		LevelSection{Spells: []Spell{
			Spell{
				Title: "CureLightWounds",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "You heal 1d8+CasterLvl to yourself or a willing subject within touch range",
			},
			Spell{
				Title: "BlessWater",
				School: "School",
				Save: "Instant",
				Range: "Touch",
				Description: "Makes holy water",
			},
		}},
	}}
	m[CLERIC] = book

	return m
}


