package main

type AffixType string

const (
	Prefix AffixType = "Prefix"
	Suffix AffixType = "Suffix"
)

func (a AffixType) String() string {
	return string(a)
}

func (a AffixType) IsPrefix() bool {
	return a == Prefix
}

func (a AffixType) IsSuffix() bool {
	return a == Suffix
}

type Stats struct {
	modifier string
	currency uint64
	maps     uint64
	scarab   uint64
	packsize uint64
	rarity   uint64
	quantity uint64
}

type MapModifier struct {
	Name           string
	Stats          Stats
	SpawnWeighting uint64
	AffixType      AffixType
}

var mapMods []MapModifier = []MapModifier{
	{
		Name: "Afflicting",
		Stats: Stats{
			modifier: "Monsters Ignite, Freeze and Shock on Hit\nAll Monster Damage can Ignite, Freeze and Shock",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			scarab:   60,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Antagonist's",
		Stats: Stats{
			modifier: "Rare Monsters each have 1 additional Modifier\n(35-45)% increased number of Rare Monsters",
			packsize: 5,
			rarity:   8,
			quantity: 13,
		},
		SpawnWeighting: 30, AffixType: Prefix},
	{
		Name: "Buffered",
		Stats: Stats{
			modifier: "Monsters gain (70-80)% of Maximum Life as Extra Maximum Energy Shield",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     35,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Chaining",
		Stats: Stats{
			modifier: "Monsters' skills Chain 3 additional times\nMonsters' Projectiles can Chain when colliding with Terrain",
			packsize: 6,
			rarity:   56,
			quantity: 16,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Diluted",
		Stats: Stats{
			modifier: "Players have 40% less effect of Flasks applied to them",
			packsize: 5,
			rarity:   8,
			quantity: 16,
			currency: 47,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "Enthralled",
		Stats: Stats{
			modifier: "Unique Bosses are Possessed",
			packsize: 6,
			rarity:   56,
			quantity: 16,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Fecund",
		Stats: Stats{
			modifier: "(90-100)% more Monster Life",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			currency: 47,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Fleet",
		Stats: Stats{
			modifier: "(35-45)% increased Monster Cast Speed\n(35-45)% increased Monster Attack Speed\n(25-30)% increased Monster Movement Speed",
			packsize: 7,
			rarity:   56,
			quantity: 19,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Grasping",
		Stats: Stats{
			modifier: "Monsters inflict 2 Grasping Vines on Hit",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     40,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Hungering",
		Stats: Stats{
			modifier: "Area contains Drowning Orbs",
			packsize: 20,
			rarity:   8,
			quantity: 16,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Labyrinth's",
		Stats: Stats{
			modifier: "Area contains Labyrinth Hazards",
			packsize: 5,
			rarity:   54,
			quantity: 19,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "Magnifying",
		Stats: Stats{
			modifier: "Monsters fire 2 additional Projectiles\nMonsters have 100% increased Area of Effect",
			packsize: 7,
			rarity:   56,
			quantity: 19,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Oppressive",
		Stats: Stats{
			modifier: "Monsters have +100% chance to Suppress Spell Damage",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     35,
		},
		SpawnWeighting: 60, AffixType: Prefix},
	{
		Name: "Prismatic",
		Stats: Stats{
			modifier: "Monsters gain (180-200)% of their Physical Damage as Extra Damage of a random Element",
			packsize: 6,
			rarity:   71,
			quantity: 16,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Profane",
		Stats: Stats{
			modifier: "Monsters gain (80-100)% of their Physical Damage as Extra Chaos Damage",
			packsize: 6,
			rarity:   46,
			quantity: 16,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Protected",
		Stats: Stats{
			modifier: "+50% Monster Physical Damage Reduction\n+55% Monster Elemental Resistances\n+35% Monster Chaos Resistance",
			packsize: 5,
			rarity:   8,
			quantity: 13,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Punishing",
		Stats: Stats{
			modifier: "Monsters reflect 20% of Physical Damage\nMonsters reflect 20% of Elemental Damage",
			packsize: 5,
			rarity:   8,
			quantity: 10,
			currency: 47,
		},
		SpawnWeighting: 40, AffixType: Prefix},
	{
		Name: "Savage",
		Stats: Stats{
			modifier: "(30-40)% increased Monster Damage",
			packsize: 7,
			rarity:   56,
			quantity: 19,
		},
		SpawnWeighting: 100, AffixType: Prefix},
	{
		Name: "Searing",
		Stats: Stats{
			modifier: "Area contains Runes of the Searing Exarch",
			packsize: 20,
			rarity:   8,
			quantity: 13,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Stalwart",
		Stats: Stats{
			modifier: "Monsters have +50% Chance to Block Attack Damage",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			currency: 47,
		},
		SpawnWeighting: 60, AffixType: Prefix},
	{
		Name: "Synthetic",
		Stats: Stats{
			modifier: "Map Boss is accompanied by a Synthesis Boss",
			packsize: 20,
			rarity:   8,
			quantity: 13,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Ultimate",
		Stats: Stats{
			modifier: "Players are assaulted by Bloodstained Sawblades",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			currency: 47,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Valdo's",
		Stats: Stats{
			modifier: "Rare monsters in area are Shaper-Touched",
			packsize: 20,
			rarity:   11,
			quantity: 13,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "Volatile",
		Stats: Stats{
			modifier: "Rare Monsters have Volatile Cores",
			packsize: 5,
			rarity:   11,
			quantity: 13,
			scarab:   53,
		},
		SpawnWeighting: 50, AffixType: Prefix},
	{
		Name: "of Collection",
		Stats: Stats{
			modifier: "The Maven interferes with Players",
			packsize: 20,
			rarity:   8,
			quantity: 13,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Congealment",
		Stats: Stats{
			modifier: "Players have (60-50)% reduced Maximum total Life, Mana and Energy Shield Recovery per second from Leech",
			packsize: 6,
			rarity:   49,
			quantity: 16,
		},
		SpawnWeighting: 30,
	},
	{
		Name: "of Curses",
		Stats: Stats{
			modifier: "Players are Cursed with Vulnerability\nPlayers are Cursed with Temporal Chains\nPlayers are Cursed with Elemental Weakness",
			packsize: 6,
			rarity:   11,
			quantity: 13,
		},
		SpawnWeighting: 60,
	},
	{
		Name: "of Deadliness",
		Stats: Stats{
			modifier: "+(70-75)% to Monster Critical Strike Multiplier\nMonsters have (650-700)% increased Critical Strike Chance",
			packsize: 6,
			rarity:   56,
			quantity: 16,
		},
		SpawnWeighting: 100,
	},
	{
		Name: "of Decaying",
		Stats: Stats{
			modifier: "Area contains Unstable Tentacle Fiends",
			packsize: 20,
			rarity:   11,
			quantity: 16,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Defiance",
		Stats: Stats{
			modifier: "Debuffs on Monsters expire 100% faster",
			packsize: 6,
			rarity:   12,
			quantity: 16,
			currency: 64,
		},
		SpawnWeighting: 40,
	},
	{
		Name: "of Domination",
		Stats: Stats{
			modifier: "Unique Monsters have a random Shrine Buff",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     35,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Endurance",
		Stats: Stats{
			modifier: "Monsters have +1 to Maximum Endurance Charges\nMonsters gain an Endurance Charge when hit",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     35,
		},
		SpawnWeighting: 100,
	},
	{
		Name: "of Exposure",
		Stats: Stats{
			modifier: "Players have -20% to all maximum Resistances",
			packsize: 20,
			rarity:   11,
			quantity: 19,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Frenzy",
		Stats: Stats{
			modifier: "Monsters gain a Frenzy Charge on Hit\nMonsters have +1 to Maximum Frenzy Charges",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     35,
		},
		SpawnWeighting: 100,
	},
	{
		Name: "of Imbibing",
		Stats: Stats{
			modifier: "Players are targeted by a Meteor when they use a Flask",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			currency: 45,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Impotence",
		Stats: Stats{
			modifier: "Players have (30-25)% less Area of Effect",
			packsize: 5,
			rarity:   69,
			quantity: 13,
		},
		SpawnWeighting: 100,
	},
	{
		Name: "of Marking",
		Stats: Stats{
			modifier: "Area contains patches of moving Marked Ground, inflicting random Marks",
			packsize: 5,
			rarity:   8,
			quantity: 16,
			scarab:   36,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Miring",
		Stats: Stats{
			modifier: "Players have (30-25)% less Defences",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			scarab:   35,
		},
		SpawnWeighting: 60,
	},
	{
		Name: "of Penetration",
		Stats: Stats{
			modifier: "Monster Damage Penetrates 15% Elemental Resistances",
			packsize: 6,
			rarity:   72,
			quantity: 19,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Persistence",
		Stats: Stats{
			modifier: "Rare monsters in area Temporarily Revive on death",
			packsize: 17,
			rarity:   8,
			quantity: 16,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Power",
		Stats: Stats{
			modifier: "Monsters gain a Power Charge on Hit\nMonsters have +1 to Maximum Power Charges",
			packsize: 5,
			rarity:   8,
			quantity: 13,
			maps:     35,
		},
		SpawnWeighting: 100,
	},
	{
		Name: "of Splinters",
		Stats: Stats{
			modifier: "25% chance for Rare Monsters to Fracture on death",
			packsize: 5,
			rarity:   8,
			quantity: 19,
		},
		SpawnWeighting: 30,
	},
	{
		Name: "of the Juggernaut",
		Stats: Stats{
			modifier: "Monsters cannot be Stunned\nMonsters' Action Speed cannot be modified to below Base Value\nMonsters' Movement Speed cannot be modified to below Base Value",
			packsize: 5,
			rarity:   8,
			quantity: 13,
		},
		SpawnWeighting: 40,
	},
	{
		Name: "of Toughness",
		Stats: Stats{
			modifier: "Monsters take (35-45)% reduced Extra Damage from Critical Strikes",
			packsize: 5,
			rarity:   56,
			quantity: 10,
		},
		SpawnWeighting: 30,
	},
	{
		Name: "of Transience",
		Stats: Stats{
			modifier: "Buffs on Players expire 100% faster",
			packsize: 4,
			rarity:   6,
			quantity: 10,
			scarab:   35,
		},
		SpawnWeighting: 60,
	},
	{
		Name: "of Treachery",
		Stats: Stats{
			modifier: "Auras from Player Skills which affect Allies also affect Enemies",
			packsize: 5,
			rarity:   54,
			quantity: 16,
		},
		SpawnWeighting: 50,
	},
	{
		Name: "of Venom",
		Stats: Stats{
			modifier: "Monsters Poison on Hit\nMonsters have 100% increased Poison Duration\nAll Damage from Monsters' Hits can Poison",
			packsize: 6,
			rarity:   49,
			quantity: 16,
		},
		SpawnWeighting: 50,
	},
}

func main() {
	_ = mapMods
}
