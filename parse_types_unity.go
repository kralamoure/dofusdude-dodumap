package dodumap

func (i JSONGameSetUnityRaw) GetID() int {
	return i.Id
}

type JSONGameSetUnityRaw struct {
	Id int `json:"id"`
	//BonusIsSecret int                        `json:"bonusIsSecret"` // always 0, maybe used in the future
	ItemIds JSONGameUnityAnkamaIdArray `json:"items"`
	NameId  int                        `json:"nameId"`
	Effects JSONGameUnityArray[struct {
		Values JSONGameUnityRefArray `json:"values"`
	}] `json:"effects"`
}

func (i *JSONGameSetUnityRaw) Merge(other [][]JSONGameItemPossibleEffectUnity) JSONGameSetUnity {
	return JSONGameSetUnity{
		Id:      i.Id,
		NameId:  i.NameId,
		Effects: other,
		ItemIds: i.ItemIds.Array,
	}
}

func (i JSONGameSetUnity) GetID() int {
	return i.Id
}

type JSONGameSetUnity struct {
	Id int `json:"id"`
	//BonusIsSecret int                        `json:"bonusIsSecret"` // always 0, maybe used in the future
	ItemIds []int                               `json:"items"`
	NameId  int                                 `json:"nameId"`
	Effects [][]JSONGameItemPossibleEffectUnity `json:"effects"`
}

type JsonGameUnityRefLookup[T any, A any] struct {
	Ref      map[int64]A
	AnkamaId map[int]T
}

type JsonGameUnityRef struct {
	Ref int64 `json:"rid"`
}

type JSONGameUnityArray[T any] struct {
	Array []T `json:"Array"`
}

type JSONGameUnityAnkamaIdArray = JSONGameUnityArray[int]
type JSONGameUnityRefArray = JSONGameUnityArray[JsonGameUnityRef]

func (i JSONGameItemUnity) GetID() int {
	return i.Id
}

type JSONGameItemUnityRaw struct {
	Id            int `json:"id"`
	TypeId        int `json:"typeId"`
	DescriptionId int `json:"descriptionId"`
	IconId        int `json:"iconId"`
	NameId        int `json:"nameId"`
	Level         int `json:"level"`

	PossibleEffects        JSONGameUnityRefArray      `json:"possibleEffects"`
	RecipeIds              JSONGameUnityAnkamaIdArray `json:"recipeIds"`
	Pods                   int                        `json:"realWeight"`
	EvolutiveEffectIds     JSONGameUnityAnkamaIdArray `json:"evolutiveEffectIds"`
	DropMonsterIds         JSONGameUnityAnkamaIdArray `json:"dropMonsterIds"`
	ItemSetId              int                        `json:"itemSetId"`
	Criteria               string                     `json:"criteria"`
	CriticalHitBonus       int                        `json:"criticalHitBonus"`
	MaxCastPerTurn         int                        `json:"maxCastPerTurn"`
	ApCost                 int                        `json:"apCost"`
	Range                  int                        `json:"range"`
	MinRange               int                        `json:"minRange"`
	CriticalHitProbability int                        `json:"criticalHitProbability"`
}

func (i JSONGameItemUnityRaw) GetID() int {
	return i.Id
}

type JSONGameItemUnity struct {
	Id            int `json:"id"`
	TypeId        int `json:"typeId"`
	DescriptionId int `json:"descriptionId"`
	IconId        int `json:"iconId"`
	NameId        int `json:"nameId"`
	Level         int `json:"level"`

	PossibleEffects        []JSONGameItemPossibleEffectUnity `json:"possibleEffects"`
	RecipeIds              JSONGameUnityAnkamaIdArray        `json:"recipeIds"`
	Pods                   int                               `json:"realWeight"`
	EvolutiveEffectIds     JSONGameUnityAnkamaIdArray        `json:"evolutiveEffectIds"`
	DropMonsterIds         JSONGameUnityAnkamaIdArray        `json:"dropMonsterIds"`
	ItemSetId              int                               `json:"itemSetId"`
	Criteria               string                            `json:"criteria"`
	CriticalHitBonus       int                               `json:"criticalHitBonus"`
	MaxCastPerTurn         int                               `json:"maxCastPerTurn"`
	ApCost                 int                               `json:"apCost"`
	Range                  int                               `json:"range"`
	MinRange               int                               `json:"minRange"`
	CriticalHitProbability int                               `json:"criticalHitProbability"`
}

func (i *JSONGameItemUnityRaw) Merge(other []JSONGameItemPossibleEffectUnity) JSONGameItemUnity {
	return JSONGameItemUnity{
		Id:                     i.Id,
		TypeId:                 i.TypeId,
		DescriptionId:          i.DescriptionId,
		IconId:                 i.IconId,
		NameId:                 i.NameId,
		Level:                  i.Level,
		PossibleEffects:        other,
		RecipeIds:              i.RecipeIds,
		Pods:                   i.Pods,
		EvolutiveEffectIds:     i.EvolutiveEffectIds,
		DropMonsterIds:         i.DropMonsterIds,
		ItemSetId:              i.ItemSetId,
		Criteria:               i.Criteria,
		CriticalHitBonus:       i.CriticalHitBonus,
		CriticalHitProbability: i.CriticalHitProbability,
	}
}

type JSONGameItemTypeUnity struct {
	Id          int `json:"id"`
	NameId      int `json:"nameId"`
	SuperTypeId int `json:"superTypeId"`
	CategoryId  int `json:"categoryId"`
	Plural      int `json:"plural"` // 0, 1
	Gender      int `json:"gender"` // 0, 1, 2
}

func (i JSONGameItemTypeUnity) GetID() int {
	return i.Id
}

type JSONGameEffectUnity struct {
	Id                       int    `json:"id"`
	DescriptionId            int    `json:"descriptionId"`
	IconId                   int    `json:"iconId"`
	Characteristic           int    `json:"characteristic"`
	Category                 int    `json:"category"`
	UseDice                  int    `json:"useDice"`                  // bool
	Active                   int    `json:"active"`                   // bool
	TheoreticalDescriptionId string `json:"theoreticalDescriptionId"` // and int inside a string lol
	BonusType                int    `json:"bonusType"`                // -1,0,+1
	ElementId                int    `json:"elementId"`
	UseInFight               int    `json:"useInFight"`  // bool
	IsInPercent              int    `json:"isInPercent"` // bool
}

func (i JSONGameEffectUnity) GetID() int {
	return i.Id
}

type JSONGameBonusUnity struct {
	Amount        int                        `json:"amount"`
	Id            int                        `json:"id"`
	CriterionsIds JSONGameUnityAnkamaIdArray `json:"criterionsIds"`
	Type          int                        `json:"type"`
}

func (i JSONGameBonusUnity) GetID() int {
	return i.Id
}

type JSONGameRecipeUnity struct {
	Id            int                        `json:"resultId"`
	NameId        string                     `json:"resultNameId"` // int in a string
	TypeId        int                        `json:"resultTypeId"`
	Level         int                        `json:"resultLevel"`
	IngredientIds JSONGameUnityAnkamaIdArray `json:"ingredientIds"`
	Quantities    JSONGameUnityArray[int]    `json:"quantities"`
	JobId         int                        `json:"jobId"`
	SkillId       int                        `json:"skillId"`
}

func (i JSONGameRecipeUnity) GetID() int {
	return i.Id
}

type JSONGameMountUnityRaw struct {
	Id            int                   `json:"id"`
	FamilyId      int                   `json:"familyId"`
	NameId        int                   `json:"nameId"`
	Effects       JSONGameUnityRefArray `json:"effects"`
	CertificateId int                   `json:"certificateId"`
}

func (i JSONGameMountUnityRaw) GetID() int {
	return i.Id
}

func (i *JSONGameMountUnityRaw) Merge(other []JSONGameItemPossibleEffectUnity) JSONGameMountUnity {
	return JSONGameMountUnity{
		Id:            i.Id,
		FamilyId:      i.FamilyId,
		NameId:        i.NameId,
		Effects:       other,
		CertificateId: i.CertificateId,
	}
}

type JSONGameMountUnity struct {
	Id            int                               `json:"id"`
	FamilyId      int                               `json:"familyId"`
	NameId        int                               `json:"nameId"`
	Effects       []JSONGameItemPossibleEffectUnity `json:"effects"`
	CertificateId int                               `json:"certificateId"`
}

func (i JSONGameMountUnity) GetID() int {
	return i.Id
}

// actually a playable class
type JSONGameBreedUnity struct {
	Id                         int                        `json:"id"`
	ShortNameId                string                     `json:"shortNameId"` // int
	LongNameId                 string                     `json:"longNameId"`  // int
	DescriptionId              int                        `json:"descriptionId"`
	GameplayClassDescriptionId string                     `json:"gameplayClassDescriptionId"` // int
	GuideItemId                int                        `json:"guideItemId"`
	MaleArtwork                int                        `json:"maleArtwork"`
	FemaleArtwork              int                        `json:"femaleArtwork"`
	BreedSpellsId              JSONGameUnityAnkamaIdArray `json:"breedSpellsId"`
	BreedRolesId               JSONGameUnityArray[struct {
		BreedId       int `json:"breedId"`
		RoleId        int `json:"roleId"`
		DescriptionId int `json:"descriptionId"`
		Value         int `json:"value"`
		Order         int `json:"order"`
	}] `json:"breedRolesId"`
	Complexity int `json:"complexity"`
	SortIndex  int `json:"sortIndex"`
}

func (i JSONGameBreedUnity) GetID() int {
	return i.Id
}

type JSONGameMountFamilyUnity struct {
	Id      int    `json:"id"`
	NameId  int    `json:"nameId"`
	HeadUri string `json:"headUri"`
}

func (i JSONGameMountFamilyUnity) GetID() int {
	return i.Id
}

type JSONGameNPCUnity struct {
	Id             int     `json:"id"`
	NameId         int     `json:"nameId"`
	DialogMessages [][]int `json:"dialogMessages"`
	DialogReplies  [][]int `json:"dialogReplies"`
	Actions        []int   `json:"actions"`
}

func (i JSONGameNPCUnity) GetID() int {
	return i.Id
}

type JSONGameTitleUnity struct {
	Id           int    `json:"id"`
	NameMaleId   string `json:"nameMaleId"`   // int in a string
	NameFemaleId string `json:"nameFemaleId"` // int in a string
	Visible      int    `json:"visible"`      // bool
	CategoryId   int    `json:"categoryId"`
}

func (i JSONGameTitleUnity) GetID() int {
	return i.Id
}

type JSONGameQuestUnity struct {
	Id             int                        `json:"id"`
	NameId         int                        `json:"nameId"`
	StepIds        JSONGameUnityAnkamaIdArray `json:"stepIds"`
	CategoryId     int                        `json:"categoryId"`
	RepeatType     int                        `json:"repeatType"`
	RepeatLimit    int                        `json:"repeatLimit"`
	IsDungeonQuest int                        `json:"isDungeonQuest"` // bool
	LevelMin       int                        `json:"levelMin"`
	LevelMax       int                        `json:"levelMax"`
	Followable     int                        `json:"followable"`   // bool
	IsPartyQuest   int                        `json:"isPartyQuest"` // bool
	StartCriterion string                     `json:"startCriterion"`
}

func (i JSONGameQuestUnity) GetID() int {
	return i.Id
}

type JSONGameQuestStepUnity struct {
	Id            int                        `json:"id"`
	DescriptionId int                        `json:"descriptionId"`
	DialogId      int                        `json:"dialogId"`
	NameId        int                        `json:"nameId"`
	OptimalLevel  int                        `json:"optimalLevel"`
	Duration      float64                    `json:"duration"`
	ObjectiveIds  JSONGameUnityAnkamaIdArray `json:"objectiveIds"`
	RewardsIds    JSONGameUnityAnkamaIdArray `json:"rewardsIds"`
	QuestId       int                        `json:"questId"`
}

func (i JSONGameQuestStepUnity) GetID() int {
	return i.Id
}

type JSONGameQuestObjectiveUnity struct {
	Id         int                    `json:"id"`
	Coords     JSONGameCoordinate     `json:"coords"`
	MapId      int                    `json:"mapId"`
	Parameters JSONGameQuestParameter `json:"parameters"`
	StepId     int                    `json:"stepId"`
	TypeId     int                    `json:"typeId"`
}

func (i JSONGameQuestObjectiveUnity) GetID() int {
	return i.Id
}

type JSONGameQuestStepRewardsUnity struct {
	Id              int     `json:"id"`
	ExperienceRatio float64 `json:"experienceRatio"`
	KamasRatio      float64 `json:"kamasRatio"`
	ItemsReward     JSONGameUnityArray[struct {
		Values JSONGameUnityArray[int] `json:"values"`
	}] `json:"itemsReward"`
	KamasScaleWithPlayerLevel int `json:"kamasScaleWithPlayerLevel"` // bool
	LevelMax                  int `json:"levelMax"`
	LevelMin                  int `json:"levelMin"`
	//SpellsReward []int `json:"spellsReward"`
	//EmotesReward []int `json:"emotesReward"`
	//TitlesReward []int `json:"titlesReward"`
	StepId int `json:"stepId"`
}

func (i JSONGameQuestStepRewardsUnity) GetID() int {
	return i.Id
}

type JSONGameQuestCategoryUnity struct {
	Id       int                        `json:"id"`
	NameId   int                        `json:"nameId"`
	Order    int                        `json:"order"`
	QuestIds JSONGameUnityAnkamaIdArray `json:"questIds"`
}

func (i JSONGameQuestCategoryUnity) GetID() int {
	return i.Id
}

type JSONGameAlamanaxCalendarUnity struct {
	Id         int                        `json:"id"`
	DescId     string                     `json:"descId"` // int
	NameId     int                        `json:"nameId"`
	NpcId      int                        `json:"npcId"`
	BonusesIds JSONGameUnityAnkamaIdArray `json:"bonusesIds"`
}

func (i JSONGameAlamanaxCalendarUnity) GetID() int {
	return i.Id
}

type JSONGameItemPossibleEffectUnity struct {
	EffectId     int `json:"effectId"`
	MinimumValue int `json:"diceNum"`
	MaximumValue int `json:"diceSide"`
	Value        int `json:"value"`

	BaseEffectId  int `json:"baseEffectId"`
	EffectElement int `json:"effectElement"`
	Dispellable   int `json:"dispellable"`
	SpellId       int `json:"spellId"`
	Duration      int `json:"duration"`
}

func (i JSONGameItemPossibleEffectUnity) GetID() int {
	return i.EffectId
}

type JSONGameDataUnity struct {
	Items     map[int]JSONGameItemUnity
	Sets      map[int]JSONGameSetUnity
	ItemTypes map[int]JSONGameItemTypeUnity
	effects   map[int]JSONGameEffectUnity
	bonuses   map[int]JSONGameBonusUnity
	Recipes   map[int]JSONGameRecipeUnity
	/*spells           map[int]JSONGameSpell
	spellTypes       map[int]JSONGameSpellType
	areas            map[int]JSONGameArea*/
	Mounts           map[int]JSONGameMountUnity
	classes          map[int]JSONGameBreedUnity
	MountFamilys     map[int]JSONGameMountFamilyUnity
	npcs             map[int]JSONGameNPCUnity
	titles           map[int]JSONGameTitleUnity
	quests           map[int]JSONGameQuestUnity
	questSteps       map[int]JSONGameQuestStepUnity
	questObjectives  map[int]JSONGameQuestObjectiveUnity
	questStepRewards map[int]JSONGameQuestStepRewardsUnity
	questCategories  map[int]JSONGameQuestCategoryUnity
	almanaxCalendars map[int]JSONGameAlamanaxCalendarUnity
}