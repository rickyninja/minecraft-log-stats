package minelog

import (
	"encoding/json"
	"time"
)

type StatsContainer struct {
	Stats       Stats
	DataVersion int
}

type Stats struct {
	Custom   Custom   `json:"minecraft:custom"`
	Used     Used     `json:"minecraft:used"`
	Mined    Mined    `json:"minecraft:mined"`
	Broken   Broken   `json:"minecraft:broken"`
	Dropped  Dropped  `json:"minecraft:dropped"`
	KilledBy KilledBy `json:"minecraft:killed_by"`
	PickedUp PickedUp `json:"minecraft:picked_up"`
	Killed   Killed   `json:"minecraft:killed"`
	Crafted  Crafted  `json:"minecraft:crafted"`
}

// stat times are in game ticks: 1/20th of second.
const gameTicksPerSecond = 20

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(p []byte) error {
	var v int
	err := json.Unmarshal(p, &v)
	if err != nil {
		return err
	}
	d.Duration = time.Second * time.Duration(v) / gameTicksPerSecond
	return nil
}

type Used struct {
	NetherBricks               int `json:"minecraft:nether_bricks"`
	CartographyTable           int `json:"minecraft:cartography_table"`
	BoneMeal                   int `json:"minecraft:bone_meal"`
	DaylightDetector           int `json:"minecraft:daylight_detector"`
	BlackBed                   int `json:"minecraft:black_bed"`
	OakBoat                    int `json:"minecraft:oak_boat"`
	MusicDisc13                int `json:"minecraft:music_disc_13"`
	IronShovel                 int `json:"minecraft:iron_shovel"`
	Tnt                        int `json:"minecraft:tnt"`
	OakTrapdoor                int `json:"minecraft:oak_trapdoor"`
	TotemOfUndying             int `json:"minecraft:totem_of_undying"`
	IronHoe                    int `json:"minecraft:iron_hoe"`
	DarkOakButton              int `json:"minecraft:dark_oak_button"`
	RoseBush                   int `json:"minecraft:rose_bush"`
	SpruceLeaves               int `json:"minecraft:spruce_leaves"`
	SpruceTrapdoor             int `json:"minecraft:spruce_trapdoor"`
	FishingRod                 int `json:"minecraft:fishing_rod"`
	OakPressurePlate           int `json:"minecraft:oak_pressure_plate"`
	SpruceLog                  int `json:"minecraft:spruce_log"`
	Lantern                    int `json:"minecraft:lantern"`
	PumpkinSeeds               int `json:"minecraft:pumpkin_seeds"`
	SpruceSign                 int `json:"minecraft:spruce_sign"`
	MossyCobblestone           int `json:"minecraft:mossy_cobblestone"`
	IronSword                  int `json:"minecraft:iron_sword"`
	RedMushroom                int `json:"minecraft:red_mushroom"`
	MossyStoneBrickStairs      int `json:"minecraft:mossy_stone_brick_stairs"`
	OakDoor                    int `json:"minecraft:oak_door"`
	SmoothStone                int `json:"minecraft:smooth_stone"`
	BirchTrapdoor              int `json:"minecraft:birch_trapdoor"`
	Gravel                     int `json:"minecraft:gravel"`
	SpruceDoor                 int `json:"minecraft:spruce_door"`
	PolishedGraniteStairs      int `json:"minecraft:polished_granite_stairs"`
	TurtleSpawnEgg             int `json:"minecraft:turtle_spawn_egg"`
	OrangeTulip                int `json:"minecraft:orange_tulip"`
	NetherWart                 int `json:"minecraft:nether_wart"`
	Chest                      int `json:"minecraft:chest"`
	WhiteShulkerBox            int `json:"minecraft:white_shulker_box"`
	WoodenPickaxe              int `json:"minecraft:wooden_pickaxe"`
	MagmaBlock                 int `json:"minecraft:magma_block"`
	SnowBlock                  int `json:"minecraft:snow_block"`
	GrassBlock                 int `json:"minecraft:grass_block"`
	Kelp                       int `json:"minecraft:kelp"`
	NetherBrickStairs          int `json:"minecraft:nether_brick_stairs"`
	Diorite                    int `json:"minecraft:diorite"`
	GlassPane                  int `json:"minecraft:glass_pane"`
	OakSign                    int `json:"minecraft:oak_sign"`
	FlowerPot                  int `json:"minecraft:flower_pot"`
	OakPlanks                  int `json:"minecraft:oak_planks"`
	GrayBed                    int `json:"minecraft:gray_bed"`
	BirchSapling               int `json:"minecraft:birch_sapling"`
	BrownBed                   int `json:"minecraft:brown_bed"`
	DiamondBlock               int `json:"minecraft:diamond_block"`
	SweetBerries               int `json:"minecraft:sweet_berries"`
	JungleBoat                 int `json:"minecraft:jungle_boat"`
	Comparator                 int `json:"minecraft:comparator"`
	Pumpkin                    int `json:"minecraft:pumpkin"`
	OakSlab                    int `json:"minecraft:oak_slab"`
	SmoothStoneSlab            int `json:"minecraft:smooth_stone_slab"`
	ChiseledQuartzBlock        int `json:"minecraft:chiseled_quartz_block"`
	BirchSign                  int `json:"minecraft:birch_sign"`
	RabbitStew                 int `json:"minecraft:rabbit_stew"`
	PinkTulip                  int `json:"minecraft:pink_tulip"`
	TropicalFishBucket         int `json:"minecraft:tropical_fish_bucket"`
	CobblestoneStairs          int `json:"minecraft:cobblestone_stairs"`
	Snowball                   int `json:"minecraft:snowball"`
	Bow                        int `json:"minecraft:bow"`
	Lilac                      int `json:"minecraft:lilac"`
	StoneBrickWall             int `json:"minecraft:stone_brick_wall"`
	NetherWartBlock            int `json:"minecraft:nether_wart_block"`
	SpruceStairs               int `json:"minecraft:spruce_stairs"`
	Cake                       int `json:"minecraft:cake"`
	HoneycombBlock             int `json:"minecraft:honeycomb_block"`
	Glowstone                  int `json:"minecraft:glowstone"`
	Bricks                     int `json:"minecraft:bricks"`
	PurpurPillar               int `json:"minecraft:purpur_pillar"`
	Melon                      int `json:"minecraft:melon"`
	Dropper                    int `json:"minecraft:dropper"`
	Beehive                    int `json:"minecraft:beehive"`
	WheatSeeds                 int `json:"minecraft:wheat_seeds"`
	Sunflower                  int `json:"minecraft:sunflower"`
	BrownMushroom              int `json:"minecraft:brown_mushroom"`
	TurtleEgg                  int `json:"minecraft:turtle_egg"`
	DiamondPickaxe             int `json:"minecraft:diamond_pickaxe"`
	OakFence                   int `json:"minecraft:oak_fence"`
	NetherBrickFence           int `json:"minecraft:nether_brick_fence"`
	PackedIce                  int `json:"minecraft:packed_ice"`
	Granite                    int `json:"minecraft:granite"`
	Minecart                   int `json:"minecraft:minecart"`
	EnderPearl                 int `json:"minecraft:ender_pearl"`
	HopperMinecart             int `json:"minecraft:hopper_minecart"`
	PumpkinPie                 int `json:"minecraft:pumpkin_pie"`
	CookedSalmon               int `json:"minecraft:cooked_salmon"`
	CobblestoneWall            int `json:"minecraft:cobblestone_wall"`
	Egg                        int `json:"minecraft:egg"`
	Barrel                     int `json:"minecraft:barrel"`
	Anvil                      int `json:"minecraft:anvil"`
	PurpurBlock                int `json:"minecraft:purpur_block"`
	Scaffolding                int `json:"minecraft:scaffolding"`
	MagentaWool                int `json:"minecraft:magenta_wool"`
	Loom                       int `json:"minecraft:loom"`
	EndStone                   int `json:"minecraft:end_stone"`
	IronDoor                   int `json:"minecraft:iron_door"`
	CyanBanner                 int `json:"minecraft:cyan_banner"`
	MusicDiscCat               int `json:"minecraft:music_disc_cat"`
	Bamboo                     int `json:"minecraft:bamboo"`
	SoulSand                   int `json:"minecraft:soul_sand"`
	JungleFenceGate            int `json:"minecraft:jungle_fence_gate"`
	SprucePlanks               int `json:"minecraft:spruce_planks"`
	MossyStoneBrickWall        int `json:"minecraft:mossy_stone_brick_wall"`
	Lever                      int `json:"minecraft:lever"`
	StonePressurePlate         int `json:"minecraft:stone_pressure_plate"`
	Lectern                    int `json:"minecraft:lectern"`
	Carrot                     int `json:"minecraft:carrot"`
	JungleButton               int `json:"minecraft:jungle_button"`
	PinkBed                    int `json:"minecraft:pink_bed"`
	Glass                      int `json:"minecraft:glass"`
	LavaBucket                 int `json:"minecraft:lava_bucket"`
	DamagedAnvil               int `json:"minecraft:damaged_anvil"`
	EnderChest                 int `json:"minecraft:ender_chest"`
	QuartzBlock                int `json:"minecraft:quartz_block"`
	RedBed                     int `json:"minecraft:red_bed"`
	GreenConcrete              int `json:"minecraft:green_concrete"`
	StoneBricks                int `json:"minecraft:stone_bricks"`
	Smoker                     int `json:"minecraft:smoker"`
	WritableBook               int `json:"minecraft:writable_book"`
	Rail                       int `json:"minecraft:rail"`
	Dandelion                  int `json:"minecraft:dandelion"`
	Cobweb                     int `json:"minecraft:cobweb"`
	GreenConcretePowder        int `json:"minecraft:green_concrete_powder"`
	NoteBlock                  int `json:"minecraft:note_block"`
	Bookshelf                  int `json:"minecraft:bookshelf"`
	WhiteBanner                int `json:"minecraft:white_banner"`
	BirchPressurePlate         int `json:"minecraft:birch_pressure_plate"`
	StoneSword                 int `json:"minecraft:stone_sword"`
	CreeperSpawnEgg            int `json:"minecraft:creeper_spawn_egg"`
	Crossbow                   int `json:"minecraft:crossbow"`
	String                     int `json:"minecraft:string"`
	StoneShovel                int `json:"minecraft:stone_shovel"`
	BoneBlock                  int `json:"minecraft:bone_block"`
	Hopper                     int `json:"minecraft:hopper"`
	Stonecutter                int `json:"minecraft:stonecutter"`
	QuartzSlab                 int `json:"minecraft:quartz_slab"`
	FlintAndSteel              int `json:"minecraft:flint_and_steel"`
	OrangeCarpet               int `json:"minecraft:orange_carpet"`
	Sand                       int `json:"minecraft:sand"`
	BlackConcretePowder        int `json:"minecraft:black_concrete_powder"`
	DiamondSword               int `json:"minecraft:diamond_sword"`
	ShulkerBox                 int `json:"minecraft:shulker_box"`
	MilkBucket                 int `json:"minecraft:milk_bucket"`
	SmithingTable              int `json:"minecraft:smithing_table"`
	SpruceSlab                 int `json:"minecraft:spruce_slab"`
	Bell                       int `json:"minecraft:bell"`
	Lead                       int `json:"minecraft:lead"`
	BlackCarpet                int `json:"minecraft:black_carpet"`
	RottenFlesh                int `json:"minecraft:rotten_flesh"`
	IronBlock                  int `json:"minecraft:iron_block"`
	IronTrapdoor               int `json:"minecraft:iron_trapdoor"`
	NetherBrickWall            int `json:"minecraft:nether_brick_wall"`
	LightBlueWool              int `json:"minecraft:light_blue_wool"`
	Observer                   int `json:"minecraft:observer"`
	Cobblestone                int `json:"minecraft:cobblestone"`
	ChorusFlower               int `json:"minecraft:chorus_flower"`
	JungleSign                 int `json:"minecraft:jungle_sign"`
	PolishedAndesite           int `json:"minecraft:polished_andesite"`
	CoarseDirt                 int `json:"minecraft:coarse_dirt"`
	StoneStairs                int `json:"minecraft:stone_stairs"`
	BlueGlazedTerracotta       int `json:"minecraft:blue_glazed_terracotta"`
	Beacon                     int `json:"minecraft:beacon"`
	MossyCobblestoneWall       int `json:"minecraft:mossy_cobblestone_wall"`
	ChestMinecart              int `json:"minecraft:chest_minecart"`
	WhiteCarpet                int `json:"minecraft:white_carpet"`
	Allium                     int `json:"minecraft:allium"`
	SeaPickle                  int `json:"minecraft:sea_pickle"`
	Campfire                   int `json:"minecraft:campfire"`
	MossyStoneBricks           int `json:"minecraft:mossy_stone_bricks"`
	Trident                    int `json:"minecraft:trident"`
	YellowBed                  int `json:"minecraft:yellow_bed"`
	Obsidian                   int `json:"minecraft:obsidian"`
	HoneyBlock                 int `json:"minecraft:honey_block"`
	DetectorRail               int `json:"minecraft:detector_rail"`
	OakFenceGate               int `json:"minecraft:oak_fence_gate"`
	PrismarineBricks           int `json:"minecraft:prismarine_bricks"`
	CoalOre                    int `json:"minecraft:coal_ore"`
	CarvedPumpkin              int `json:"minecraft:carved_pumpkin"`
	CookedBeef                 int `json:"minecraft:cooked_beef"`
	DarkOakPressure_plate      int `json:"minecraft:dark_oak_pressure_plate"`
	StoneAxe                   int `json:"minecraft:stone_axe"`
	NetherQuartzOre            int `json:"minecraft:nether_quartz_ore"`
	Grindstone                 int `json:"minecraft:grindstone"`
	WitherSkeletonSkull        int `json:"minecraft:wither_skeleton_skull"`
	LimeBanner                 int `json:"minecraft:lime_banner"`
	PrismarineWall             int `json:"minecraft:prismarine_wall"`
	CookedChicken              int `json:"minecraft:cooked_chicken"`
	JungleFence                int `json:"minecraft:jungle_fence"`
	SeaLantern                 int `json:"minecraft:sea_lantern"`
	CodBucket                  int `json:"minecraft:cod_bucket"`
	RedstoneLamp               int `json:"minecraft:redstone_lamp"`
	RedstoneBlock              int `json:"minecraft:redstone_block"`
	JungleSapling              int `json:"minecraft:jungle_sapling"`
	BlastFurnace               int `json:"minecraft:blast_furnace"`
	Potion                     int `json:"minecraft:potion"`
	PolishedDiorite            int `json:"minecraft:polished_diorite"`
	StoneButton                int `json:"minecraft:stone_button"`
	SpruceFence                int `json:"minecraft:spruce_fence"`
	Sandstone                  int `json:"minecraft:sandstone"`
	OrangeTerracotta           int `json:"minecraft:orange_terracotta"`
	BirchFenceGate             int `json:"minecraft:birch_fence_gate"`
	DarkOakFenceGate           int `json:"minecraft:dark_oak_fence_gate"`
	Netherrack                 int `json:"minecraft:netherrack"`
	DiamondAxe                 int `json:"minecraft:diamond_axe"`
	Ladder                     int `json:"minecraft:ladder"`
	Terracotta                 int `json:"minecraft:terracotta"`
	BlueBanner                 int `json:"minecraft:blue_banner"`
	ArmorStand                 int `json:"minecraft:armor_stand"`
	BirchBoat                  int `json:"minecraft:birch_boat"`
	DiamondShovel              int `json:"minecraft:diamond_shovel"`
	Cornflower                 int `json:"minecraft:cornflower"`
	IronAxe                    int `json:"minecraft:iron_axe"`
	MushroomStew               int `json:"minecraft:mushroom_stew"`
	PurpleBanner               int `json:"minecraft:purple_banner"`
	CraftingTable              int `json:"minecraft:crafting_table"`
	SugarCane                  int `json:"minecraft:sugar_cane"`
	ChippedAnvil               int `json:"minecraft:chipped_anvil"`
	PufferfishBucket           int `json:"minecraft:pufferfish_bucket"`
	LimeCarpet                 int `json:"minecraft:lime_carpet"`
	FletchingTable             int `json:"minecraft:fletching_table"`
	Dispenser                  int `json:"minecraft:dispenser"`
	WhiteWool                  int `json:"minecraft:white_wool"`
	BirchLeaves                int `json:"minecraft:birch_leaves"`
	StickyPiston               int `json:"minecraft:sticky_piston"`
	Shears                     int `json:"minecraft:shears"`
	OakSapling                 int `json:"minecraft:oak_sapling"`
	WaterBucket                int `json:"minecraft:water_bucket"`
	OakButton                  int `json:"minecraft:oak_button"`
	StoneHoe                   int `json:"minecraft:stone_hoe"`
	IronPickaxe                int `json:"minecraft:iron_pickaxe"`
	BirchDoor                  int `json:"minecraft:birch_door"`
	StonePickaxe               int `json:"minecraft:stone_pickaxe"`
	Dirt                       int `json:"minecraft:dirt"`
	Poppy                      int `json:"minecraft:poppy"`
	Ice                        int `json:"minecraft:ice"`
	StoneSlab                  int `json:"minecraft:stone_slab"`
	DarkOakDoor                int `json:"minecraft:dark_oak_door"`
	HeavyWeightedPressurePlate int `json:"minecraft:heavy_weighted_pressure_plate"`
	DarkOakTrapdoor            int `json:"minecraft:dark_oak_trapdoor"`
	SplashPotion               int `json:"minecraft:splash_potion"`
	TrappedChest               int `json:"minecraft:trapped_chest"`
	Potato                     int `json:"minecraft:potato"`
	SalmonBucket               int `json:"minecraft:salmon_bucket"`
	Mycelium                   int `json:"minecraft:mycelium"`
	ItemFrame                  int `json:"minecraft:item_frame"`
	Torch                      int `json:"minecraft:torch"`
	JungleTrapdoor             int `json:"minecraft:jungle_trapdoor"`
	PurpleBed                  int `json:"minecraft:purple_bed"`
	PurpleWool                 int `json:"minecraft:purple_wool"`
	BlueBed                    int `json:"minecraft:blue_bed"`
	GlassBottle                int `json:"minecraft:glass_bottle"`
	NetherBrickSlab            int `json:"minecraft:nether_brick_slab"`
	OakLog                     int `json:"minecraft:oak_log"`
	Andesite                   int `json:"minecraft:andesite"`
	StoneBrickStairs           int `json:"minecraft:stone_brick_stairs"`
	IronBars                   int `json:"minecraft:iron_bars"`
	MossyCobblestoneStairs     int `json:"minecraft:mossy_cobblestone_stairs"`
	HayBlock                   int `json:"minecraft:hay_block"`
	MusicDiscStrad             int `json:"minecraft:music_disc_strad"`
	CutSandstone               int `json:"minecraft:cut_sandstone"`
	PolishedGranite            int `json:"minecraft:polished_granite"`
	MelonSeeds                 int `json:"minecraft:melon_seeds"`
	PoweredRail                int `json:"minecraft:powered_rail"`
	FireworkRocket             int `json:"minecraft:firework_rocket"`
	Bucket                     int `json:"minecraft:bucket"`
	Podzol                     int `json:"minecraft:podzol"`
	QuartzStairs               int `json:"minecraft:quartz_stairs"`
	BirchLog                   int `json:"minecraft:birch_log"`
	Bread                      int `json:"minecraft:bread"`
	RedstoneTorch              int `json:"minecraft:redstone_torch"`
	Cauldron                   int `json:"minecraft:cauldron"`
	RedShulkerBox              int `json:"minecraft:red_shulker_box"`
	CobblestoneSlab            int `json:"minecraft:cobblestone_slab"`
	JackOLantern               int `json:"minecraft:jack_o_lantern"`
	SuspiciousStew             int `json:"minecraft:suspicious_stew"`
	PrismarineBrickSlab        int `json:"minecraft:prismarine_brick_slab"`
	MagentaCarpet              int `json:"minecraft:magenta_carpet"`
	BlackConcrete              int `json:"minecraft:black_concrete"`
	Redstone                   int `json:"minecraft:redstone"`
	GreenCarpet                int `json:"minecraft:green_carpet"`
	VillagerSpawnEgg           int `json:"minecraft:villager_spawn_egg"`
	EnchantingTable            int `json:"minecraft:enchanting_table"`
	BeetrootSeeds              int `json:"minecraft:beetroot_seeds"`
	Piston                     int `json:"minecraft:piston"`
	Cactus                     int `json:"minecraft:cactus"`
	DarkOakSign                int `json:"minecraft:dark_oak_sign"`
	BrewingStand               int `json:"minecraft:brewing_stand"`
	BirchButton                int `json:"minecraft:birch_button"`
	RedstoneOre                int `json:"minecraft:redstone_ore"`
	MusicDiscWard              int `json:"minecraft:music_disc_ward"`
	Stone                      int `json:"minecraft:stone"`
	Composter                  int `json:"minecraft:composter"`
	BirchFence                 int `json:"minecraft:birch_fence"`
	ChorusFruit                int `json:"minecraft:chorus_fruit"`
	DarkOakFence               int `json:"minecraft:dark_oak_fence"`
	CookedCod                  int `json:"minecraft:cooked_cod"`
	EmeraldBlock               int `json:"minecraft:emerald_block"`
	TripwireHook               int `json:"minecraft:tripwire_hook"`
	EndStoneBricks             int `json:"minecraft:end_stone_bricks"`
	SlimeBlock                 int `json:"minecraft:slime_block"`
	Repeater                   int `json:"minecraft:repeater"`
	BirchPlanks                int `json:"minecraft:birch_planks"`
	StoneBrickSlab             int `json:"minecraft:stone_brick_slab"`
	PolishedGraniteSlab        int `json:"minecraft:polished_granite_slab"`
	GraniteStairs              int `json:"minecraft:granite_stairs"`
	BlueOrchid                 int `json:"minecraft:blue_orchid"`
	OrangeGlazedTerracotta     int `json:"minecraft:orange_glazed_terracotta"`
	WhiteBed                   int `json:"minecraft:white_bed"`
	Furnace                    int `json:"minecraft:furnace"`
}

type Custom struct {
	OpenChest                    int      `json:"minecraft:open_chest"`
	SneakTime                    Duration `json:"minecraft:sneak_time"`
	OpenEnderChest               int      `json:"minecraft:open_enderchest"`
	ClimbOneCM                   int      `json:"minecraft:climb_one_cm"`
	TuneNoteblock                int      `json:"minecraft:tune_noteblock"`
	MobKills                     int      `json:"minecraft:mob_kills"`
	PlayOneMinute                int      `json:"minecraft:play_one_minute"`
	Drop                         int      `json:"minecraft:drop"`
	InteractWithCraftingTable    int      `json:"minecraft:interact_with_crafting_table"`
	TalkedToVillager             int      `json:"minecraft:talked_to_villager"`
	DamageTaken                  int      `json:"minecraft:damage_taken"`
	TimeSinceRest                Duration `json:"minecraft:time_since_rest"`
	LeaveGame                    int      `json:"minecraft:leave_game"`
	FallOneCM                    int      `json:"minecraft:fall_one_cm"`
	DamageDealt                  int      `json:"minecraft:damage_dealt"`
	CrouchOneCM                  int      `json:"minecraft:crouch_one_cm"`
	InteractWithStonecutter      int      `json:"minecraft:interact_with_stonecutter"`
	FlyOneCM                     int      `json:"minecraft:fly_one_cm"`
	InteractWithFurnace          int      `json:"minecraft:interact_with_furnace"`
	TradedWithVillager           int      `json:"minecraft:traded_with_villager"`
	Jump                         int      `json:"minecraft:jump"`
	DamageBlockedByShield        int      `json:"minecraft:damage_blocked_by_shield"`
	WalkOneCM                    int      `json:"minecraft:walk_one_cm"`
	Deaths                       int      `json:"minecraft:deaths"`
	SwimOneCM                    int      `json:"minecraft:swim_one_cm"`
	TimeSinceDeath               Duration `json:"minecraft:time_since_death"`
	WalkUnderWaterOneCM          int      `json:"minecraft:walk_under_water_one_cm"`
	SprintOneCM                  int      `json:"minecraft:sprint_one_cm"`
	InteractWithLoom             int      `json:"minecraft:interact_with_loom"`
	SleepInBed                   int      `json:"minecraft:sleep_in_bed"`
	InteractWithAnvil            int      `json:"minecraft:interact_with_anvil"`
	WalkOnWaterOneCM             int      `json:"minecraft:walk_on_water_one_cm"`
	BoatOneCM                    int      `json:"minecraft:boat_one_cm"`
	AviateOneCM                  int      `json:"minecraft:aviate_one_cm"`
	PlayerKills                  int      `json:"minecraft:player_kills"`
	InteractWithBlastFurnace     int      `json:"minecraft:interact_with_blast_furnace"`
	OpenBarrel                   int      `json:"minecraft:open_barrel"`
	InteractWithGrindstone       int      `json:"minecraft:interact_with_grindstone"`
	AnimalsBred                  int      `json:"minecraft:animals_bred"`
	InspectDispenser             int      `json:"minecraft:inspect_dispenser"`
	RaidWin                      int      `json:"minecraft:raid_win"`
	EatCakeSlice                 int      `json:"minecraft:eat_cake_slice"`
	InteractWithCampfire         int      `json:"minecraft:interact_with_campfire"`
	PlayNoteblock                int      `json:"minecraft:play_noteblock"`
	InteractWithCartographyTable int      `json:"minecraft:interact_with_cartography_table"`
	FishCaught                   int      `json:"minecraft:fish_caught"`
	InteractWithBeacon           int      `json:"minecraft:interact_with_beacon"`
	TriggerTrappedChest          int      `json:"minecraft:trigger_trapped_chest"`
	InteractWithBrewingStand     int      `json:"minecraft:interact_with_brewingstand"`
	OpenShulkerBox               int      `json:"minecraft:open_shulker_box"`
	InteractWithSmoker           int      `json:"minecraft:interact_with_smoker"`
	InspectHopper                int      `json:"minecraft:inspect_hopper"`
	EnchantItem                  int      `json:"minecraft:enchant_item"`
	DamageDealtAbsorbed          int      `json:"minecraft:damage_dealt_absorbed"`
	RaidTrigger                  int      `json:"minecraft:raid_trigger"`
	FillCauldron                 int      `json:"minecraft:fill_cauldron"`
	PotFlower                    int      `json:"minecraft:pot_flower"`
	InspectDropper               int      `json:"minecraft:inspect_dropper"`
	MinecartOneCM                int      `json:"minecraft:minecart_one_cm"`
	PlayRecord                   int      `json:"minecraft:play_record"`
	HorseOneCM                   int      `json:"minecraft:horse_one_cm"`
	BellRing                     int      `json:"minecraft:bell_ring"`
}

type Mined struct {
	RedSand                    int `json:"minecraft:red_sand"`
	PinkTulip                  int `json:"minecraft:pink_tulip"`
	Pumpkin                    int `json:"minecraft:pumpkin"`
	OakSlab                    int `json:"minecraft:oak_slab"`
	SmoothStoneSlab            int `json:"minecraft:smooth_stone_slab"`
	BirchSign                  int `json:"minecraft:birch_sign"`
	ChiseledQuartzBlock        int `json:"minecraft:chiseled_quartz_block"`
	PurpurSlab                 int `json:"minecraft:purpur_slab"`
	CobblestoneStairs          int `json:"minecraft:cobblestone_stairs"`
	BirchSapling               int `json:"minecraft:birch_sapling"`
	JunglePlanks               int `json:"minecraft:jungle_planks"`
	Tripwire                   int `json:"minecraft:tripwire"`
	DarkOakLog                 int `json:"minecraft:dark_oak_log"`
	BrownBed                   int `json:"minecraft:brown_bed"`
	FlowerPot                  int `json:"minecraft:flower_pot"`
	OakPlanks                  int `json:"minecraft:oak_planks"`
	GrayBed                    int `json:"minecraft:gray_bed"`
	Comparator                 int `json:"minecraft:comparator"`
	DiamondBlock               int `json:"minecraft:diamond_block"`
	BrownMushroom              int `json:"minecraft:brown_mushroom"`
	DarkPrismarine             int `json:"minecraft:dark_prismarine"`
	OakFence                   int `json:"minecraft:oak_fence"`
	SandstoneStairs            int `json:"minecraft:sandstone_stairs"`
	NetherBrickFence           int `json:"minecraft:nether_brick_fence"`
	PackedIce                  int `json:"minecraft:packed_ice"`
	TurtleEgg                  int `json:"minecraft:turtle_egg"`
	InfestedStone              int `json:"minecraft:infested_stone"`
	ChorusPlant                int `json:"minecraft:chorus_plant"`
	StoneBrickWall             int `json:"minecraft:stone_brick_wall"`
	NetherWartBlock            int `json:"minecraft:nether_wart_block"`
	Grass                      int `json:"minecraft:grass"`
	SpruceStairs               int `json:"minecraft:spruce_stairs"`
	JungleStairs               int `json:"minecraft:jungle_stairs"`
	PurpurPillar               int `json:"minecraft:purpur_pillar"`
	Bricks                     int `json:"minecraft:bricks"`
	Melon                      int `json:"minecraft:melon"`
	Beehive                    int `json:"minecraft:beehive"`
	Dropper                    int `json:"minecraft:dropper"`
	HoneycombBlock             int `json:"minecraft:honeycomb_block"`
	Glowstone                  int `json:"minecraft:glowstone"`
	SpruceLeaves               int `json:"minecraft:spruce_leaves"`
	OakTrapdoor                int `json:"minecraft:oak_trapdoor"`
	SmoothSandstone            int `json:"minecraft:smooth_sandstone"`
	Lantern                    int `json:"minecraft:lantern"`
	SpruceTrapdoor             int `json:"minecraft:spruce_trapdoor"`
	OakPressurePlate           int `json:"minecraft:oak_pressure_plate"`
	SpruceLog                  int `json:"minecraft:spruce_log"`
	MushroomStem               int `json:"minecraft:mushroom_stem"`
	BlackBed                   int `json:"minecraft:black_bed"`
	DaylightDetector           int `json:"minecraft:daylight_detector"`
	WhiteWallBanner            int `json:"minecraft:white_wall_banner"`
	NetherBricks               int `json:"minecraft:nether_bricks"`
	CartographyTable           int `json:"minecraft:cartography_table"`
	EmeraldOre                 int `json:"minecraft:emerald_ore"`
	WhiteTerracotta            int `json:"minecraft:white_terracotta"`
	RedMushroomBlock           int `json:"minecraft:red_mushroom_block"`
	LightGrayTerracotta        int `json:"minecraft:light_gray_terracotta"`
	DarkOakSlab                int `json:"minecraft:dark_oak_slab"`
	GrassBlock                 int `json:"minecraft:grass_block"`
	Kelp                       int `json:"minecraft:kelp"`
	WhiteShulkerBox            int `json:"minecraft:white_shulker_box"`
	OrangeStainedGlass         int `json:"minecraft:orange_stained_glass"`
	MagmaBlock                 int `json:"minecraft:magma_block"`
	RedTerracotta              int `json:"minecraft:red_terracotta"`
	JungleLeaves               int `json:"minecraft:jungle_leaves"`
	Diorite                    int `json:"minecraft:diorite"`
	SandstoneSlab              int `json:"minecraft:sandstone_slab"`
	NetherBrickStairs          int `json:"minecraft:nether_brick_stairs"`
	GlassPane                  int `json:"minecraft:glass_pane"`
	LimeStainedGlass           int `json:"minecraft:lime_stained_glass"`
	OakSign                    int `json:"minecraft:oak_sign"`
	SmoothStone                int `json:"minecraft:smooth_stone"`
	BirchTrapdoor              int `json:"minecraft:birch_trapdoor"`
	Gravel                     int `json:"minecraft:gravel"`
	MossyCobblestone           int `json:"minecraft:mossy_cobblestone"`
	RedMushroom                int `json:"minecraft:red_mushroom"`
	MossyStoneBrickStairs      int `json:"minecraft:mossy_stone_brick_stairs"`
	Chest                      int `json:"minecraft:chest"`
	NetherWart                 int `json:"minecraft:nether_wart"`
	PurpleWallBanner           int `json:"minecraft:purple_wall_banner"`
	PolishedGraniteStairs      int `json:"minecraft:polished_granite_stairs"`
	OrangeTulip                int `json:"minecraft:orange_tulip"`
	Sand                       int `json:"minecraft:sand"`
	QuartzSlab                 int `json:"minecraft:quartz_slab"`
	Stonecutter                int `json:"minecraft:stonecutter"`
	OrangeCarpet               int `json:"minecraft:orange_carpet"`
	JungleLog                  int `json:"minecraft:jungle_log"`
	BlackConcretePowder        int `json:"minecraft:black_concrete_powder"`
	OakLeaves                  int `json:"minecraft:oak_leaves"`
	KelpPlant                  int `json:"minecraft:kelp_plant"`
	LilyPad                    int `json:"minecraft:lily_pad"`
	LilyOfTheValley            int `json:"minecraft:lily_of_the_valley"`
	Snow                       int `json:"minecraft:snow"`
	BoneBlock                  int `json:"minecraft:bone_block"`
	Hopper                     int `json:"minecraft:hopper"`
	Observer                   int `json:"minecraft:observer"`
	Cobblestone                int `json:"minecraft:cobblestone"`
	BlackCarpet                int `json:"minecraft:black_carpet"`
	LapisOre                   int `json:"minecraft:lapis_ore"`
	NetherBrickWall            int `json:"minecraft:nether_brick_wall"`
	IronTrapdoor               int `json:"minecraft:iron_trapdoor"`
	IronBlock                  int `json:"minecraft:iron_block"`
	LightBlueWool              int `json:"minecraft:light_blue_wool"`
	StoneStairs                int `json:"minecraft:stone_stairs"`
	CoarseDirt                 int `json:"minecraft:coarse_dirt"`
	BeeNest                    int `json:"minecraft:bee_nest"`
	BlueGlazedTerracotta       int `json:"minecraft:blue_glazed_terracotta"`
	Beacon                     int `json:"minecraft:beacon"`
	MossyCobblestoneWall       int `json:"minecraft:mossy_cobblestone_wall"`
	ChorusFlower               int `json:"minecraft:chorus_flower"`
	JungleSign                 int `json:"minecraft:jungle_sign"`
	SweetBerryBush             int `json:"minecraft:sweet_berry_bush"`
	PolishedAndesite           int `json:"minecraft:polished_andesite"`
	Beetroots                  int `json:"minecraft:beetroots"`
	MelonStem                  int `json:"minecraft:melon_stem"`
	ShulkerBox                 int `json:"minecraft:shulker_box"`
	SmithingTable              int `json:"minecraft:smithing_table"`
	SpruceSlab                 int `json:"minecraft:spruce_slab"`
	SprucePlanks               int `json:"minecraft:spruce_planks"`
	MossyStoneBrickWall        int `json:"minecraft:mossy_stone_brick_wall"`
	DarkOakLeaves              int `json:"minecraft:dark_oak_leaves"`
	Bamboo                     int `json:"minecraft:bamboo"`
	SoulSand                   int `json:"minecraft:soul_sand"`
	JungleFenceGate            int `json:"minecraft:jungle_fence_gate"`
	LightBlueTerracotta        int `json:"minecraft:light_blue_terracotta"`
	StonePressurePlate         int `json:"minecraft:stone_pressure_plate"`
	Lectern                    int `json:"minecraft:lectern"`
	JungleButton               int `json:"minecraft:jungle_button"`
	Lever                      int `json:"minecraft:lever"`
	CobblestoneWall            int `json:"minecraft:cobblestone_wall"`
	Granite                    int `json:"minecraft:granite"`
	EndStone                   int `json:"minecraft:end_stone"`
	Loom                       int `json:"minecraft:loom"`
	Barrel                     int `json:"minecraft:barrel"`
	PurpurBlock                int `json:"minecraft:purpur_block"`
	Anvil                      int `json:"minecraft:anvil"`
	Scaffolding                int `json:"minecraft:scaffolding"`
	MagentaWool                int `json:"minecraft:magenta_wool"`
	GreenConcretePowder        int `json:"minecraft:green_concrete_powder"`
	NoteBlock                  int `json:"minecraft:note_block"`
	Bookshelf                  int `json:"minecraft:bookshelf"`
	Cobweb                     int `json:"minecraft:cobweb"`
	BirchPressurePlate         int `json:"minecraft:birch_pressure_plate"`
	PurpleGlazedTerracotta     int `json:"minecraft:purple_glazed_terracotta"`
	PolishedDioriteStairs      int `json:"minecraft:polished_diorite_stairs"`
	DamagedAnvil               int `json:"minecraft:damaged_anvil"`
	EnderChest                 int `json:"minecraft:ender_chest"`
	RedBed                     int `json:"minecraft:red_bed"`
	QuartzBlock                int `json:"minecraft:quartz_block"`
	PinkBed                    int `json:"minecraft:pink_bed"`
	Glass                      int `json:"minecraft:glass"`
	Farmland                   int `json:"minecraft:farmland"`
	Rail                       int `json:"minecraft:rail"`
	Dandelion                  int `json:"minecraft:dandelion"`
	Cocoa                      int `json:"minecraft:cocoa"`
	GreenConcrete              int `json:"minecraft:green_concrete"`
	StoneBricks                int `json:"minecraft:stone_bricks"`
	Clay                       int `json:"minecraft:clay"`
	DarkOakFenceGate           int `json:"minecraft:dark_oak_fence_gate"`
	Netherrack                 int `json:"minecraft:netherrack"`
	OrangeTerracotta           int `json:"minecraft:orange_terracotta"`
	BirchFenceGate             int `json:"minecraft:birch_fence_gate"`
	InfestedStoneBricks        int `json:"minecraft:infested_stone_bricks"`
	Ladder                     int `json:"minecraft:ladder"`
	Terracotta                 int `json:"minecraft:terracotta"`
	BlueBanner                 int `json:"minecraft:blue_banner"`
	AzureBluet                 int `json:"minecraft:azure_bluet"`
	StoneButton                int `json:"minecraft:stone_button"`
	RedTulip                   int `json:"minecraft:red_tulip"`
	PolishedDiorite            int `json:"minecraft:polished_diorite"`
	SpruceFence                int `json:"minecraft:spruce_fence"`
	Sandstone                  int `json:"minecraft:sandstone"`
	Prismarine                 int `json:"minecraft:prismarine"`
	BirchLeaves                int `json:"minecraft:birch_leaves"`
	DarkOakWallSign            int `json:"minecraft:dark_oak_wall_sign"`
	LimeCarpet                 int `json:"minecraft:lime_carpet"`
	Dispenser                  int `json:"minecraft:dispenser"`
	WhiteWool                  int `json:"minecraft:white_wool"`
	OakSapling                 int `json:"minecraft:oak_sapling"`
	OakButton                  int `json:"minecraft:oak_button"`
	StickyPiston               int `json:"minecraft:sticky_piston"`
	BlueWallBanner             int `json:"minecraft:blue_wall_banner"`
	Cornflower                 int `json:"minecraft:cornflower"`
	Potatoes                   int `json:"minecraft:potatoes"`
	SugarCane                  int `json:"minecraft:sugar_cane"`
	ChippedAnvil               int `json:"minecraft:chipped_anvil"`
	RedstoneWallTorch          int `json:"minecraft:redstone_wall_torch"`
	RedstoneWire               int `json:"minecraft:redstone_wire"`
	PurpleBanner               int `json:"minecraft:purple_banner"`
	Seagrass                   int `json:"minecraft:seagrass"`
	DarkOakPlanks              int `json:"minecraft:dark_oak_planks"`
	CraftingTable              int `json:"minecraft:crafting_table"`
	CoalOre                    int `json:"minecraft:coal_ore"`
	CarvedPumpkin              int `json:"minecraft:carved_pumpkin"`
	GrassPath                  int `json:"minecraft:grass_path"`
	Grindstone                 int `json:"minecraft:grindstone"`
	WitherSkeletonSkull        int `json:"minecraft:wither_skeleton_skull"`
	LimeBanner                 int `json:"minecraft:lime_banner"`
	NetherQuartzOre            int `json:"minecraft:nether_quartz_ore"`
	Campfire                   int `json:"minecraft:campfire"`
	MossyStoneBricks           int `json:"minecraft:mossy_stone_bricks"`
	WhiteCarpet                int `json:"minecraft:white_carpet"`
	FrostedIce                 int `json:"minecraft:frosted_ice"`
	Allium                     int `json:"minecraft:allium"`
	SeaPickle                  int `json:"minecraft:sea_pickle"`
	PrismarineBricks           int `json:"minecraft:prismarine_bricks"`
	ChiseledSandstone          int `json:"minecraft:chiseled_sandstone"`
	YellowBed                  int `json:"minecraft:yellow_bed"`
	BirchWallSign              int `json:"minecraft:birch_wall_sign"`
	Obsidian                   int `json:"minecraft:obsidian"`
	OakFenceGate               int `json:"minecraft:oak_fence_gate"`
	HoneyBlock                 int `json:"minecraft:honey_block"`
	DetectorRail               int `json:"minecraft:detector_rail"`
	SeaLantern                 int `json:"minecraft:sea_lantern"`
	CyanWallBanner             int `json:"minecraft:cyan_wall_banner"`
	JungleFence                int `json:"minecraft:jungle_fence"`
	Fern                       int `json:"minecraft:fern"`
	YellowTerracotta           int `json:"minecraft:yellow_terracotta"`
	StrippedDarkOakLog         int `json:"minecraft:stripped_dark_oak_log"`
	RedstoneBlock              int `json:"minecraft:redstone_block"`
	BlastFurnace               int `json:"minecraft:blast_furnace"`
	RedstoneLamp               int `json:"minecraft:redstone_lamp"`
	JungleWallSign             int `json:"minecraft:jungle_wall_sign"`
	InfestedCobblestone        int `json:"minecraft:infested_cobblestone"`
	PrismarineWall             int `json:"minecraft:prismarine_wall"`
	WallTorch                  int `json:"minecraft:wall_torch"`
	BlackConcrete              int `json:"minecraft:black_concrete"`
	OakWallSign                int `json:"minecraft:oak_wall_sign"`
	PrismarineBrickSlab        int `json:"minecraft:prismarine_brick_slab"`
	Piston                     int `json:"minecraft:piston"`
	GreenCarpet                int `json:"minecraft:green_carpet"`
	BrownMushroomBlock         int `json:"minecraft:brown_mushroom_block"`
	JackOLantern               int `json:"minecraft:jack_o_lantern"`
	GoldOre                    int `json:"minecraft:gold_ore"`
	SlimeBlock                 int `json:"minecraft:slime_block"`
	Repeater                   int `json:"minecraft:repeater"`
	BirchPlanks                int `json:"minecraft:birch_planks"`
	DarkOakStairs              int `json:"minecraft:dark_oak_stairs"`
	TripwireHook               int `json:"minecraft:tripwire_hook"`
	EndStoneBricks             int `json:"minecraft:end_stone_bricks"`
	StoneBrickSlab             int `json:"minecraft:stone_brick_slab"`
	PolishedGraniteSlab        int `json:"minecraft:polished_granite_slab"`
	GraniteStairs              int `json:"minecraft:granite_stairs"`
	Vine                       int `json:"minecraft:vine"`
	BlueOrchid                 int `json:"minecraft:blue_orchid"`
	OrangeGlazedTerracotta     int `json:"minecraft:orange_glazed_terracotta"`
	WhiteBed                   int `json:"minecraft:white_bed"`
	JungleSlab                 int `json:"minecraft:jungle_slab"`
	Furnace                    int `json:"minecraft:furnace"`
	OxeyeDaisy                 int `json:"minecraft:oxeye_daisy"`
	WhiteTulip                 int `json:"minecraft:white_tulip"`
	PistonHead                 int `json:"minecraft:piston_head"`
	BrewingStand               int `json:"minecraft:brewing_stand"`
	Carrots                    int `json:"minecraft:carrots"`
	RedstoneOre                int `json:"minecraft:redstone_ore"`
	BirchButton                int `json:"minecraft:birch_button"`
	Stone                      int `json:"minecraft:stone"`
	Cactus                     int `json:"minecraft:cactus"`
	PurpurStairs               int `json:"minecraft:purpur_stairs"`
	EmeraldBlock               int `json:"minecraft:emerald_block"`
	Composter                  int `json:"minecraft:composter"`
	BirchFence                 int `json:"minecraft:birch_fence"`
	Wheat                      int `json:"minecraft:wheat"`
	DarkOakFence               int `json:"minecraft:dark_oak_fence"`
	TrappedChest               int `json:"minecraft:trapped_chest"`
	EndRod                     int `json:"minecraft:end_rod"`
	PurpleBed                  int `json:"minecraft:purple_bed"`
	JungleTrapdoor             int `json:"minecraft:jungle_trapdoor"`
	PurpleWool                 int `json:"minecraft:purple_wool"`
	Mycelium                   int `json:"minecraft:mycelium"`
	Torch                      int `json:"minecraft:torch"`
	MagentaWallBanner          int `json:"minecraft:magenta_wall_banner"`
	Dirt                       int `json:"minecraft:dirt"`
	LargeFern                  int `json:"minecraft:large_fern"`
	HeavyWeightedPressurePlate int `json:"minecraft:heavy_weighted_pressure_plate"`
	IronOre                    int `json:"minecraft:iron_ore"`
	DarkOakTrapdoor            int `json:"minecraft:dark_oak_trapdoor"`
	Poppy                      int `json:"minecraft:poppy"`
	Ice                        int `json:"minecraft:ice"`
	PolishedGranite            int `json:"minecraft:polished_granite"`
	PoweredRail                int `json:"minecraft:powered_rail"`
	CutSandstone               int `json:"minecraft:cut_sandstone"`
	RedShulkerBox              int `json:"minecraft:red_shulker_box"`
	Cauldron                   int `json:"minecraft:cauldron"`
	CobblestoneSlab            int `json:"minecraft:cobblestone_slab"`
	CrackedStoneBricks         int `json:"minecraft:cracked_stone_bricks"`
	Podzol                     int `json:"minecraft:podzol"`
	QuartzStairs               int `json:"minecraft:quartz_stairs"`
	BirchLog                   int `json:"minecraft:birch_log"`
	RedstoneTorch              int `json:"minecraft:redstone_torch"`
	DiamondOre                 int `json:"minecraft:diamond_ore"`
	OakLog                     int `json:"minecraft:oak_log"`
	BlueBed                    int `json:"minecraft:blue_bed"`
	NetherBrickSlab            int `json:"minecraft:nether_brick_slab"`
	IronBars                   int `json:"minecraft:iron_bars"`
	MossyCobblestoneStairs     int `json:"minecraft:mossy_cobblestone_stairs"`
	HayBlock                   int `json:"minecraft:hay_block"`
	Andesite                   int `json:"minecraft:andesite"`
	StoneBrickStairs           int `json:"minecraft:stone_brick_stairs"`
}

type Broken struct {
	FlintAndSteel     int `json:"minecraft:flint_and_steel" `
	GoldenHelmet      int `json:"minecraft:golden_helmet" `
	StoneAxe          int `json:"minecraft:stone_axe" `
	FishingRod        int `json:"minecraft:fishing_rod" `
	GoldenBoots       int `json:"minecraft:golden_boots" `
	Shears            int `json:"minecraft:shears" `
	ChainmailHelmet   int `json:"minecraft:chainmail_helmet" `
	StoneSword        int `json:"minecraft:stone_sword" `
	StoneHoe          int `json:"minecraft:stone_hoe" `
	IronPickaxe       int `json:"minecraft:iron_pickaxe" `
	IronChestplate    int `json:"minecraft:iron_chestplate" `
	StonePickaxe      int `json:"minecraft:stone_pickaxe" `
	LeatherBoots      int `json:"minecraft:leather_boots" `
	LeatherChestplate int `json:"minecraft:leather_chestplate" `
	Shield            int `json:"minecraft:shield" `
	GoldenLeggings    int `json:"minecraft:golden_leggings" `
	GoldenChestplate  int `json:"minecraft:golden_chestplate" `
	StoneShovel       int `json:"minecraft:stone_shovel" `
	IronShovel        int `json:"minecraft:iron_shovel" `
}
type Dropped struct {
	Stick                      int `json:"minecraft:stick" `
	IronAxe                    int `json:"minecraft:iron_axe" `
	Cornflower                 int `json:"minecraft:cornflower" `
	DiamondShovel              int `json:"minecraft:diamond_shovel" `
	BirchBoat                  int `json:"minecraft:birch_boat" `
	SugarCane                  int `json:"minecraft:sugar_cane" `
	CraftingTable              int `json:"minecraft:crafting_table" `
	Seagrass                   int `json:"minecraft:seagrass" `
	DarkOakPlanks              int `json:"minecraft:dark_oak_planks" `
	Gunpowder                  int `json:"minecraft:gunpowder" `
	WhiteWool                  int `json:"minecraft:white_wool" `
	PufferfishBucket           int `json:"minecraft:pufferfish_bucket" `
	StoneHoe                   int `json:"minecraft:stone_hoe" `
	OakButton                  int `json:"minecraft:oak_button" `
	WaterBucket                int `json:"minecraft:water_bucket" `
	OakSapling                 int `json:"minecraft:oak_sapling" `
	Shears                     int `json:"minecraft:shears" `
	WoodenShovel               int `json:"minecraft:wooden_shovel" `
	LapisBlock                 int `json:"minecraft:lapis_block" `
	StickyPiston               int `json:"minecraft:sticky_piston" `
	NameTag                    int `json:"minecraft:name_tag" `
	StoneButton                int `json:"minecraft:stone_button" `
	Bowl                       int `json:"minecraft:bowl" `
	AzureBluet                 int `json:"minecraft:azure_bluet" `
	Potion                     int `json:"minecraft:potion" `
	OrangeWool                 int `json:"minecraft:orange_wool" `
	DiamondAxe                 int `json:"minecraft:diamond_axe" `
	Netherrack                 int `json:"minecraft:netherrack" `
	SpruceSapling              int `json:"minecraft:spruce_sapling" `
	Ladder                     int `json:"minecraft:ladder" `
	LeatherBoots               int `json:"minecraft:leather_boots" `
	BubbleCoralFan             int `json:"minecraft:bubble_coral_fan" `
	TubeCoral                  int `json:"minecraft:tube_coral" `
	Feather                    int `json:"minecraft:feather" `
	CookedChicken              int `json:"minecraft:cooked_chicken" `
	CyanWool                   int `json:"minecraft:cyan_wool" `
	IronHelmet                 int `json:"minecraft:iron_helmet" `
	BlazeRod                   int `json:"minecraft:blaze_rod" `
	Fern                       int `json:"minecraft:fern" `
	GoldIngot                  int `json:"minecraft:gold_ingot" `
	DiamondBoots               int `json:"minecraft:diamond_boots" `
	RedstoneBlock              int `json:"minecraft:redstone_block" `
	FireCoralBlock             int `json:"minecraft:fire_coral_block" `
	RedstoneLamp               int `json:"minecraft:redstone_lamp" `
	Sugar                      int `json:"minecraft:sugar" `
	Trident                    int `json:"minecraft:trident" `
	BrainCoralFan              int `json:"minecraft:brain_coral_fan" `
	PinkWool                   int `json:"minecraft:pink_wool" `
	Campfire                   int `json:"minecraft:campfire" `
	SeaPickle                  int `json:"minecraft:sea_pickle" `
	Obsidian                   int `json:"minecraft:obsidian" `
	RabbitHide                 int `json:"minecraft:rabbit_hide" `
	Diamond                    int `json:"minecraft:diamond" `
	Grindstone                 int `json:"minecraft:grindstone" `
	StoneAxe                   int `json:"minecraft:stone_axe" `
	Stone                      int `json:"minecraft:stone" `
	PurpurStairs               int `json:"minecraft:purpur_stairs" `
	Cactus                     int `json:"minecraft:cactus" `
	CookedCod                  int `json:"minecraft:cooked_cod" `
	HornCoral                  int `json:"minecraft:horn_coral" `
	DarkOakFence               int `json:"minecraft:dark_oak_fence" `
	ChorusFruit                int `json:"minecraft:chorus_fruit" `
	BirchFence                 int `json:"minecraft:birch_fence" `
	Wheat                      int `json:"minecraft:wheat" `
	BirchPlanks                int `json:"minecraft:birch_planks" `
	BlueWool                   int `json:"minecraft:blue_wool" `
	SlimeBlock                 int `json:"minecraft:slime_block" `
	EndStoneBricks             int `json:"minecraft:end_stone_bricks" `
	TripwireHook               int `json:"minecraft:tripwire_hook" `
	OxeyeDaisy                 int `json:"minecraft:oxeye_daisy" `
	Furnace                    int `json:"minecraft:furnace" `
	StoneBrickSlab             int `json:"minecraft:stone_brick_slab" `
	Emerald                    int `json:"minecraft:emerald" `
	DiamondLeggings            int `json:"minecraft:diamond_leggings" `
	BrainCoralBlock            int `json:"minecraft:brain_coral_block" `
	JackOLantern               int `json:"minecraft:jack_o_lantern" `
	Arrow                      int `json:"minecraft:arrow" `
	BrownMushroomBlock         int `json:"minecraft:brown_mushroom_block" `
	SuspiciousStew             int `json:"minecraft:suspicious_stew" `
	Honeycomb                  int `json:"minecraft:honeycomb" `
	GoldOre                    int `json:"minecraft:gold_ore" `
	Redstone                   int `json:"minecraft:redstone" `
	Beetroot                   int `json:"minecraft:beetroot" `
	IronIngot                  int `json:"minecraft:iron_ingot" `
	EnchantedBook              int `json:"minecraft:enchanted_book" `
	Piston                     int `json:"minecraft:piston" `
	BeetrootSeeds              int `json:"minecraft:beetroot_seeds" `
	ChainmailHelmet            int `json:"minecraft:chainmail_helmet" `
	FilledMap                  int `json:"minecraft:filled_map" `
	Saddle                     int `json:"minecraft:saddle" `
	SpruceBoat                 int `json:"minecraft:spruce_boat" `
	OakLog                     int `json:"minecraft:oak_log" `
	IronBoots                  int `json:"minecraft:iron_boots" `
	NetherBrickSlab            int `json:"minecraft:nether_brick_slab" `
	RedDye                     int `json:"minecraft:red_dye" `
	IronNugget                 int `json:"minecraft:iron_nugget" `
	PoisonousPotato            int `json:"minecraft:poisonous_potato" `
	GlassBottle                int `json:"minecraft:glass_bottle" `
	IronBars                   int `json:"minecraft:iron_bars" `
	StoneBrickStairs           int `json:"minecraft:stone_brick_stairs" `
	RedWool                    int `json:"minecraft:red_wool" `
	Andesite                   int `json:"minecraft:andesite" `
	DiamondHelmet              int `json:"minecraft:diamond_helmet" `
	Bucket                     int `json:"minecraft:bucket" `
	LeatherHelmet              int `json:"minecraft:leather_helmet" `
	FireworkRocket             int `json:"minecraft:firework_rocket" `
	MelonSeeds                 int `json:"minecraft:melon_seeds" `
	PurpleStainedGlass         int `json:"minecraft:purple_stained_glass" `
	Coal                       int `json:"minecraft:coal" `
	IronLeggings               int `json:"minecraft:iron_leggings" `
	Flint                      int `json:"minecraft:flint" `
	CrackedStoneBricks         int `json:"minecraft:cracked_stone_bricks" `
	CobblestoneSlab            int `json:"minecraft:cobblestone_slab" `
	RedstoneTorch              int `json:"minecraft:redstone_torch" `
	AcaciaPlanks               int `json:"minecraft:acacia_planks" `
	InkSac                     int `json:"minecraft:ink_sac" `
	GoldenBoots                int `json:"minecraft:golden_boots" `
	BirchLog                   int `json:"minecraft:birch_log" `
	YellowWool                 int `json:"minecraft:yellow_wool" `
	TubeCoralFan               int `json:"minecraft:tube_coral_fan" `
	Dirt                       int `json:"minecraft:dirt" `
	BirchDoor                  int `json:"minecraft:birch_door" `
	StonePickaxe               int `json:"minecraft:stone_pickaxe" `
	IronPickaxe                int `json:"minecraft:iron_pickaxe" `
	IronOre                    int `json:"minecraft:iron_ore" `
	HeavyWeightedPressurePlate int `json:"minecraft:heavy_weighted_pressure_plate" `
	Ice                        int `json:"minecraft:ice" `
	Poppy                      int `json:"minecraft:poppy" `
	TubeCoralBlock             int `json:"minecraft:tube_coral_block" `
	SplashPotion               int `json:"minecraft:splash_potion" `
	PurpleWool                 int `json:"minecraft:purple_wool" `
	PurpleBed                  int `json:"minecraft:purple_bed" `
	Torch                      int `json:"minecraft:torch" `
	Potato                     int `json:"minecraft:potato" `
	SpruceStairs               int `json:"minecraft:spruce_stairs" `
	Grass                      int `json:"minecraft:grass" `
	Lilac                      int `json:"minecraft:lilac" `
	NetherWartBlock            int `json:"minecraft:nether_wart_block" `
	ChainmailBoots             int `json:"minecraft:chainmail_boots" `
	Bow                        int `json:"minecraft:bow" `
	Snowball                   int `json:"minecraft:snowball" `
	WheatSeeds                 int `json:"minecraft:wheat_seeds" `
	Melon                      int `json:"minecraft:melon" `
	Bricks                     int `json:"minecraft:bricks" `
	GoldenLeggings             int `json:"minecraft:golden_leggings" `
	PurpurPillar               int `json:"minecraft:purpur_pillar" `
	FireCoralFan               int `json:"minecraft:fire_coral_fan" `
	BrownMushroom              int `json:"minecraft:brown_mushroom" `
	Sunflower                  int `json:"minecraft:sunflower" `
	BlazePowder                int `json:"minecraft:blaze_powder" `
	PackedIce                  int `json:"minecraft:packed_ice" `
	WoodenSword                int `json:"minecraft:wooden_sword" `
	BrownWool                  int `json:"minecraft:brown_wool" `
	DiamondPickaxe             int `json:"minecraft:diamond_pickaxe" `
	OakFence                   int `json:"minecraft:oak_fence" `
	GrayWool                   int `json:"minecraft:gray_wool" `
	DarkOakLog                 int `json:"minecraft:dark_oak_log" `
	HornCoralBlock             int `json:"minecraft:horn_coral_block" `
	MelonSlice                 int `json:"minecraft:melon_slice" `
	JunglePlanks               int `json:"minecraft:jungle_planks" `
	BirchSapling               int `json:"minecraft:birch_sapling" `
	OakPlanks                  int `json:"minecraft:oak_planks" `
	Chicken                    int `json:"minecraft:chicken" `
	Shield                     int `json:"minecraft:shield" `
	SlimeBall                  int `json:"minecraft:slime_ball" `
	Cod                        int `json:"minecraft:cod" `
	RabbitStew                 int `json:"minecraft:rabbit_stew" `
	BirchSign                  int `json:"minecraft:birch_sign" `
	SmoothStoneSlab            int `json:"minecraft:smooth_stone_slab" `
	Pumpkin                    int `json:"minecraft:pumpkin" `
	CobblestoneStairs          int `json:"minecraft:cobblestone_stairs" `
	BrainCoral                 int `json:"minecraft:brain_coral" `
	Gravel                     int `json:"minecraft:gravel" `
	SmoothStone                int `json:"minecraft:smooth_stone" `
	GreenBanner                int `json:"minecraft:green_banner" `
	RedMushroom                int `json:"minecraft:red_mushroom" `
	IronSword                  int `json:"minecraft:iron_sword" `
	Chest                      int `json:"minecraft:chest" `
	TurtleSpawnEgg             int `json:"minecraft:turtle_spawn_egg" `
	Kelp                       int `json:"minecraft:kelp" `
	GrassBlock                 int `json:"minecraft:grass_block" `
	SnowBlock                  int `json:"minecraft:snow_block" `
	WoodenPickaxe              int `json:"minecraft:wooden_pickaxe" `
	LimeWool                   int `json:"minecraft:lime_wool" `
	OakSign                    int `json:"minecraft:oak_sign" `
	GlassPane                  int `json:"minecraft:glass_pane" `
	Diorite                    int `json:"minecraft:diorite" `
	BlackBed                   int `json:"minecraft:black_bed" `
	DaylightDetector           int `json:"minecraft:daylight_detector" `
	SpiderEye                  int `json:"minecraft:spider_eye" `
	GoldenApple                int `json:"minecraft:golden_apple" `
	BoneMeal                   int `json:"minecraft:bone_meal" `
	IronChestplate             int `json:"minecraft:iron_chestplate" `
	GreenDye                   int `json:"minecraft:green_dye" `
	NetherBricks               int `json:"minecraft:nether_bricks" `
	IronShovel                 int `json:"minecraft:iron_shovel" `
	OakBoat                    int `json:"minecraft:oak_boat" `
	IronHoe                    int `json:"minecraft:iron_hoe" `
	BubbleCoral                int `json:"minecraft:bubble_coral" `
	OakTrapdoor                int `json:"minecraft:oak_trapdoor" `
	PumpkinSeeds               int `json:"minecraft:pumpkin_seeds" `
	SpruceLog                  int `json:"minecraft:spruce_log" `
	FishingRod                 int `json:"minecraft:fishing_rod" `
	GreenWool                  int `json:"minecraft:green_wool" `
	DiamondSword               int `json:"minecraft:diamond_sword" `
	Pufferfish                 int `json:"minecraft:pufferfish" `
	SpruceSlab                 int `json:"minecraft:spruce_slab" `
	Beef                       int `json:"minecraft:beef" `
	GoldenSword                int `json:"minecraft:golden_sword" `
	LeatherChestplate          int `json:"minecraft:leather_chestplate" `
	Cobblestone                int `json:"minecraft:cobblestone" `
	Observer                   int `json:"minecraft:observer" `
	LightBlueWool              int `json:"minecraft:light_blue_wool" `
	IronBlock                  int `json:"minecraft:iron_block" `
	RottenFlesh                int `json:"minecraft:rotten_flesh" `
	Salmon                     int `json:"minecraft:salmon" `
	BlueGlazedTerracotta       int `json:"minecraft:blue_glazed_terracotta" `
	BlackWool                  int `json:"minecraft:black_wool" `
	ChorusFlower               int `json:"minecraft:chorus_flower" `
	Crossbow                   int `json:"minecraft:crossbow" `
	Bone                       int `json:"minecraft:bone" `
	LightGrayWool              int `json:"minecraft:light_gray_wool" `
	StoneShovel                int `json:"minecraft:stone_shovel" `
	PhantomMembrane            int `json:"minecraft:phantom_membrane" `
	ChainmailChestplate        int `json:"minecraft:chainmail_chestplate" `
	String                     int `json:"minecraft:string" `
	Sand                       int `json:"minecraft:sand" `
	WoodenAxe                  int `json:"minecraft:wooden_axe" `
	ShulkerShell               int `json:"minecraft:shulker_shell" `
	FlintAndSteel              int `json:"minecraft:flint_and_steel" `
	JungleLog                  int `json:"minecraft:jungle_log" `
	EnderChest                 int `json:"minecraft:ender_chest" `
	LavaBucket                 int `json:"minecraft:lava_bucket" `
	Glass                      int `json:"minecraft:glass" `
	Rail                       int `json:"minecraft:rail" `
	WritableBook               int `json:"minecraft:writable_book" `
	StoneBricks                int `json:"minecraft:stone_bricks" `
	GreenConcrete              int `json:"minecraft:green_concrete" `
	WhiteBanner                int `json:"minecraft:white_banner" `
	LapisLazuli                int `json:"minecraft:lapis_lazuli" `
	LeatherLeggings            int `json:"minecraft:leather_leggings" `
	GoldenHelmet               int `json:"minecraft:golden_helmet" `
	StoneSword                 int `json:"minecraft:stone_sword" `
	BirchPressurePlate         int `json:"minecraft:birch_pressure_plate" `
	Book                       int `json:"minecraft:book" `
	Egg                        int `json:"minecraft:egg" `
	CookedSalmon               int `json:"minecraft:cooked_salmon" `
	PumpkinPie                 int `json:"minecraft:pumpkin_pie" `
	EnderPearl                 int `json:"minecraft:ender_pearl" `
	Granite                    int `json:"minecraft:granite" `
	Minecart                   int `json:"minecraft:minecart" `
	DiamondChestplate          int `json:"minecraft:diamond_chestplate" `
	GoldenChestplate           int `json:"minecraft:golden_chestplate" `
	EndStone                   int `json:"minecraft:end_stone" `
	Scaffolding                int `json:"minecraft:scaffolding" `
	MagentaWool                int `json:"minecraft:magenta_wool" `
	GoldNugget                 int `json:"minecraft:gold_nugget" `
	PurpurBlock                int `json:"minecraft:purpur_block" `
	SprucePlanks               int `json:"minecraft:spruce_planks" `
	HornCoralFan               int `json:"minecraft:horn_coral_fan" `
	Bamboo                     int `json:"minecraft:bamboo" `
	Carrot                     int `json:"minecraft:carrot" `
	StonePressurePlate         int `json:"minecraft:stone_pressure_plate" `
	Lever                      int `json:"minecraft:lever" `
}

type KilledBy struct {
	Witch          int `json:"minecraft:witch" `
	ZombieVillager int `json:"minecraft:zombie_villager" `
	Ghast          int `json:"minecraft:ghast" `
	Vindicator     int `json:"minecraft:vindicator" `
	ZombiePigman   int `json:"minecraft:zombie_pigman" `
	Blaze          int `json:"minecraft:blaze" `
	IronGolem      int `json:"minecraft:iron_golem" `
	Wolf           int `json:"minecraft:wolf" `
	Vex            int `json:"minecraft:vex" `
	Ravager        int `json:"minecraft:ravager" `
	Skeleton       int `json:"minecraft:skeleton" `
	CaveSpider     int `json:"minecraft:cave_spider" `
	Zombie         int `json:"minecraft:zombie" `
	Guardian       int `json:"minecraft:guardian" `
	Creeper        int `json:"minecraft:creeper" `
	EnderDragon    int `json:"minecraft:ender_dragon" `
}

type PickedUp struct {
	LilyOfTheValley            int `json:"minecraft:lily_of_the_valley" `
	MagentaStainedGlass        int `json:"minecraft:magenta_stained_glass" `
	Bone                       int `json:"minecraft:bone" `
	LightGrayWool              int `json:"minecraft:light_gray_wool" `
	Crossbow                   int `json:"minecraft:crossbow" `
	LilyPad                    int `json:"minecraft:lily_pad" `
	String                     int `json:"minecraft:string" `
	BoneBlock                  int `json:"minecraft:bone_block" `
	Hopper                     int `json:"minecraft:hopper" `
	StoneShovel                int `json:"minecraft:stone_shovel" `
	PhantomMembrane            int `json:"minecraft:phantom_membrane" `
	ChainmailChestplate        int `json:"minecraft:chainmail_chestplate" `
	WoodenAxe                  int `json:"minecraft:wooden_axe" `
	OrangeCarpet               int `json:"minecraft:orange_carpet" `
	ShulkerShell               int `json:"minecraft:shulker_shell" `
	NetherStar                 int `json:"minecraft:nether_star" `
	Porkchop                   int `json:"minecraft:porkchop" `
	FlintAndSteel              int `json:"minecraft:flint_and_steel" `
	Stonecutter                int `json:"minecraft:stonecutter" `
	QuartzSlab                 int `json:"minecraft:quartz_slab" `
	Sand                       int `json:"minecraft:sand" `
	OakLeaves                  int `json:"minecraft:oak_leaves" `
	GreenStainedGlass          int `json:"minecraft:green_stained_glass" `
	BlackConcretePowder        int `json:"minecraft:black_concrete_powder" `
	WetSponge                  int `json:"minecraft:wet_sponge" `
	JungleLog                  int `json:"minecraft:jungle_log" `
	TropicalFish               int `json:"minecraft:tropical_fish" `
	DiamondSword               int `json:"minecraft:diamond_sword" `
	ShulkerBox                 int `json:"minecraft:shulker_box" `
	Pufferfish                 int `json:"minecraft:pufferfish" `
	MilkBucket                 int `json:"minecraft:milk_bucket" `
	GoldenSword                int `json:"minecraft:golden_sword" `
	LeatherChestplate          int `json:"minecraft:leather_chestplate" `
	SpectralArrow              int `json:"minecraft:spectral_arrow" `
	SmithingTable              int `json:"minecraft:smithing_table" `
	SpruceSlab                 int `json:"minecraft:spruce_slab" `
	BirchSlab                  int `json:"minecraft:birch_slab" `
	Beef                       int `json:"minecraft:beef" `
	LightBlueWool              int `json:"minecraft:light_blue_wool" `
	IronBlock                  int `json:"minecraft:iron_block" `
	IronTrapdoor               int `json:"minecraft:iron_trapdoor" `
	NetherBrickWall            int `json:"minecraft:nether_brick_wall" `
	RottenFlesh                int `json:"minecraft:rotten_flesh" `
	BlackCarpet                int `json:"minecraft:black_carpet" `
	Lead                       int `json:"minecraft:lead" `
	Cobblestone                int `json:"minecraft:cobblestone" `
	Observer                   int `json:"minecraft:observer" `
	PolishedAndesite           int `json:"minecraft:polished_andesite" `
	JungleSign                 int `json:"minecraft:jungle_sign" `
	BlackWool                  int `json:"minecraft:black_wool" `
	ChorusFlower               int `json:"minecraft:chorus_flower" `
	MossyCobblestoneWall       int `json:"minecraft:mossy_cobblestone_wall" `
	Salmon                     int `json:"minecraft:salmon" `
	Beacon                     int `json:"minecraft:beacon" `
	BlueGlazedTerracotta       int `json:"minecraft:blue_glazed_terracotta" `
	Compass                    int `json:"minecraft:compass" `
	CoarseDirt                 int `json:"minecraft:coarse_dirt" `
	StoneStairs                int `json:"minecraft:stone_stairs" `
	EnderPearl                 int `json:"minecraft:ender_pearl" `
	Granite                    int `json:"minecraft:granite" `
	Minecart                   int `json:"minecraft:minecart" `
	CobblestoneWall            int `json:"minecraft:cobblestone_wall" `
	Egg                        int `json:"minecraft:egg" `
	CookedSalmon               int `json:"minecraft:cooked_salmon" `
	PumpkinPie                 int `json:"minecraft:pumpkin_pie" `
	MagentaWool                int `json:"minecraft:magenta_wool" `
	Scaffolding                int `json:"minecraft:scaffolding" `
	Anvil                      int `json:"minecraft:anvil" `
	PurpurBlock                int `json:"minecraft:purpur_block" `
	GoldNugget                 int `json:"minecraft:gold_nugget" `
	Barrel                     int `json:"minecraft:barrel" `
	MusicDiscCat               int `json:"minecraft:music_disc_cat" `
	CyanBanner                 int `json:"minecraft:cyan_banner" `
	DiamondChestplate          int `json:"minecraft:diamond_chestplate" `
	GoldenChestplate           int `json:"minecraft:golden_chestplate" `
	CookedMutton               int `json:"minecraft:cooked_mutton" `
	IronDoor                   int `json:"minecraft:iron_door" `
	Loom                       int `json:"minecraft:loom" `
	EndStone                   int `json:"minecraft:end_stone" `
	JungleFenceGate            int `json:"minecraft:jungle_fence_gate" `
	SoulSand                   int `json:"minecraft:soul_sand" `
	Bamboo                     int `json:"minecraft:bamboo" `
	FermentedSpiderEye         int `json:"minecraft:fermented_spider_eye" `
	Elytra                     int `json:"minecraft:elytra" `
	MossyStoneBrickWall        int `json:"minecraft:mossy_stone_brick_wall" `
	ChainmailLeggings          int `json:"minecraft:chainmail_leggings" `
	SprucePlanks               int `json:"minecraft:spruce_planks" `
	HornCoralFan               int `json:"minecraft:horn_coral_fan" `
	Lever                      int `json:"minecraft:lever" `
	JungleButton               int `json:"minecraft:jungle_button" `
	Carrot                     int `json:"minecraft:carrot" `
	Lectern                    int `json:"minecraft:lectern" `
	LightBlueTerracotta        int `json:"minecraft:light_blue_terracotta" `
	HoneyBottle                int `json:"minecraft:honey_bottle" `
	StonePressurePlate         int `json:"minecraft:stone_pressure_plate" `
	Glass                      int `json:"minecraft:glass" `
	LavaBucket                 int `json:"minecraft:lava_bucket" `
	PinkBed                    int `json:"minecraft:pink_bed" `
	RedBed                     int `json:"minecraft:red_bed" `
	QuartzBlock                int `json:"minecraft:quartz_block" `
	EnderChest                 int `json:"minecraft:ender_chest" `
	DamagedAnvil               int `json:"minecraft:damaged_anvil" `
	StoneBricks                int `json:"minecraft:stone_bricks" `
	Rabbit                     int `json:"minecraft:rabbit" `
	GreenConcrete              int `json:"minecraft:green_concrete" `
	Quartz                     int `json:"minecraft:quartz" `
	WritableBook               int `json:"minecraft:writable_book" `
	Dandelion                  int `json:"minecraft:dandelion" `
	Rail                       int `json:"minecraft:rail" `
	GoldenHelmet               int `json:"minecraft:golden_helmet" `
	Cobweb                     int `json:"minecraft:cobweb" `
	WhiteBanner                int `json:"minecraft:white_banner" `
	NoteBlock                  int `json:"minecraft:note_block" `
	LapisLazuli                int `json:"minecraft:lapis_lazuli" `
	GreenConcretePowder        int `json:"minecraft:green_concrete_powder" `
	LeatherLeggings            int `json:"minecraft:leather_leggings" `
	Paper                      int `json:"minecraft:paper" `
	PolishedDioriteStairs      int `json:"minecraft:polished_diorite_stairs" `
	StoneSword                 int `json:"minecraft:stone_sword" `
	Charcoal                   int `json:"minecraft:charcoal" `
	BirchPressurePlate         int `json:"minecraft:birch_pressure_plate" `
	PurpleGlazedTerracotta     int `json:"minecraft:purple_glazed_terracotta" `
	Book                       int `json:"minecraft:book" `
	GrayBed                    int `json:"minecraft:gray_bed" `
	OakPlanks                  int `json:"minecraft:oak_planks" `
	Chicken                    int `json:"minecraft:chicken" `
	FlowerPot                  int `json:"minecraft:flower_pot" `
	BrownBed                   int `json:"minecraft:brown_bed" `
	DarkOakLog                 int `json:"minecraft:dark_oak_log" `
	HornCoralBlock             int `json:"minecraft:horn_coral_block" `
	JunglePlanks               int `json:"minecraft:jungle_planks" `
	BirchSapling               int `json:"minecraft:birch_sapling" `
	MelonSlice                 int `json:"minecraft:melon_slice" `
	SlimeBall                  int `json:"minecraft:slime_ball" `
	SweetBerries               int `json:"minecraft:sweet_berries" `
	Cod                        int `json:"minecraft:cod" `
	DiamondBlock               int `json:"minecraft:diamond_block" `
	Comparator                 int `json:"minecraft:comparator" `
	JungleBoat                 int `json:"minecraft:jungle_boat" `
	Shield                     int `json:"minecraft:shield" `
	ChiseledQuartzBlock        int `json:"minecraft:chiseled_quartz_block" `
	BirchSign                  int `json:"minecraft:birch_sign" `
	SmoothStoneSlab            int `json:"minecraft:smooth_stone_slab" `
	OakSlab                    int `json:"minecraft:oak_slab" `
	Pumpkin                    int `json:"minecraft:pumpkin" `
	PinkTulip                  int `json:"minecraft:pink_tulip" `
	RedSand                    int `json:"minecraft:red_sand" `
	RabbitStew                 int `json:"minecraft:rabbit_stew" `
	BlueDye                    int `json:"minecraft:blue_dye" `
	PrismarineCrystals         int `json:"minecraft:prismarine_crystals" `
	BrainCoral                 int `json:"minecraft:brain_coral" `
	CobblestoneStairs          int `json:"minecraft:cobblestone_stairs" `
	PurpurSlab                 int `json:"minecraft:purpur_slab" `
	Bow                        int `json:"minecraft:bow" `
	Snowball                   int `json:"minecraft:snowball" `
	SpruceStairs               int `json:"minecraft:spruce_stairs" `
	Grass                      int `json:"minecraft:grass" `
	Lilac                      int `json:"minecraft:lilac" `
	NetherWartBlock            int `json:"minecraft:nether_wart_block" `
	StoneBrickWall             int `json:"minecraft:stone_brick_wall" `
	ChainmailBoots             int `json:"minecraft:chainmail_boots" `
	Glowstone                  int `json:"minecraft:glowstone" `
	HoneycombBlock             int `json:"minecraft:honeycomb_block" `
	FireCoralFan               int `json:"minecraft:fire_coral_fan" `
	Dropper                    int `json:"minecraft:dropper" `
	WheatSeeds                 int `json:"minecraft:wheat_seeds" `
	Beehive                    int `json:"minecraft:beehive" `
	Melon                      int `json:"minecraft:melon" `
	Bricks                     int `json:"minecraft:bricks" `
	PurpurPillar               int `json:"minecraft:purpur_pillar" `
	GoldenLeggings             int `json:"minecraft:golden_leggings" `
	DarkPrismarine             int `json:"minecraft:dark_prismarine" `
	Sunflower                  int `json:"minecraft:sunflower" `
	BrownMushroom              int `json:"minecraft:brown_mushroom" `
	TurtleEgg                  int `json:"minecraft:turtle_egg" `
	PackedIce                  int `json:"minecraft:packed_ice" `
	WoodenSword                int `json:"minecraft:wooden_sword" `
	NetherBrickFence           int `json:"minecraft:nether_brick_fence" `
	BrownWool                  int `json:"minecraft:brown_wool" `
	OakFence                   int `json:"minecraft:oak_fence" `
	SandstoneStairs            int `json:"minecraft:sandstone_stairs" `
	DiamondPickaxe             int `json:"minecraft:diamond_pickaxe" `
	GrayWool                   int `json:"minecraft:gray_wool" `
	CartographyTable           int `json:"minecraft:cartography_table" `
	IronChestplate             int `json:"minecraft:iron_chestplate" `
	GreenDye                   int `json:"minecraft:green_dye" `
	NetherBricks               int `json:"minecraft:nether_bricks" `
	BlackBed                   int `json:"minecraft:black_bed" `
	DaylightDetector           int `json:"minecraft:daylight_detector" `
	SpiderEye                  int `json:"minecraft:spider_eye" `
	BoneMeal                   int `json:"minecraft:bone_meal" `
	GoldenApple                int `json:"minecraft:golden_apple" `
	DarkOakSlab                int `json:"minecraft:dark_oak_slab" `
	LightGrayTerracotta        int `json:"minecraft:light_gray_terracotta" `
	OakBoat                    int `json:"minecraft:oak_boat" `
	MusicDisc13                int `json:"minecraft:music_disc_13" `
	IronShovel                 int `json:"minecraft:iron_shovel" `
	WhiteTerracotta            int `json:"minecraft:white_terracotta" `
	DarkOakButton              int `json:"minecraft:dark_oak_button" `
	SmoothSandstone            int `json:"minecraft:smooth_sandstone" `
	IronHoe                    int `json:"minecraft:iron_hoe" `
	TotemOfUndying             int `json:"minecraft:totem_of_undying" `
	BubbleCoral                int `json:"minecraft:bubble_coral" `
	OakTrapdoor                int `json:"minecraft:oak_trapdoor" `
	SpruceLeaves               int `json:"minecraft:spruce_leaves" `
	RoseBush                   int `json:"minecraft:rose_bush" `
	SpruceLog                  int `json:"minecraft:spruce_log" `
	OakPressurePlate           int `json:"minecraft:oak_pressure_plate" `
	FishingRod                 int `json:"minecraft:fishing_rod" `
	SpruceTrapdoor             int `json:"minecraft:spruce_trapdoor" `
	GreenWool                  int `json:"minecraft:green_wool" `
	CreeperHead                int `json:"minecraft:creeper_head" `
	PumpkinSeeds               int `json:"minecraft:pumpkin_seeds" `
	Lantern                    int `json:"minecraft:lantern" `
	MossyStoneBrickStairs      int `json:"minecraft:mossy_stone_brick_stairs" `
	RedMushroom                int `json:"minecraft:red_mushroom" `
	IronSword                  int `json:"minecraft:iron_sword" `
	MossyCobblestone           int `json:"minecraft:mossy_cobblestone" `
	Gravel                     int `json:"minecraft:gravel" `
	SmoothStone                int `json:"minecraft:smooth_stone" `
	BirchTrapdoor              int `json:"minecraft:birch_trapdoor" `
	OakDoor                    int `json:"minecraft:oak_door" `
	TurtleSpawnEgg             int `json:"minecraft:turtle_spawn_egg" `
	OrangeTulip                int `json:"minecraft:orange_tulip" `
	Apple                      int `json:"minecraft:apple" `
	PolishedGraniteStairs      int `json:"minecraft:polished_granite_stairs" `
	ClayBall                   int `json:"minecraft:clay_ball" `
	Chest                      int `json:"minecraft:chest" `
	NetherWart                 int `json:"minecraft:nether_wart" `
	SnowBlock                  int `json:"minecraft:snow_block" `
	MagmaBlock                 int `json:"minecraft:magma_block" `
	WoodenPickaxe              int `json:"minecraft:wooden_pickaxe" `
	WhiteShulkerBox            int `json:"minecraft:white_shulker_box" `
	GrassBlock                 int `json:"minecraft:grass_block" `
	Kelp                       int `json:"minecraft:kelp" `
	OakSign                    int `json:"minecraft:oak_sign" `
	GlassPane                  int `json:"minecraft:glass_pane" `
	SandstoneSlab              int `json:"minecraft:sandstone_slab" `
	NetherBrickStairs          int `json:"minecraft:nether_brick_stairs" `
	Diorite                    int `json:"minecraft:diorite" `
	LimeWool                   int `json:"minecraft:lime_wool" `
	RedTerracotta              int `json:"minecraft:red_terracotta" `
	DiamondLeggings            int `json:"minecraft:diamond_leggings" `
	BrainCoralBlock            int `json:"minecraft:brain_coral_block" `
	JackOLantern               int `json:"minecraft:jack_o_lantern" `
	BrownMushroomBlock         int `json:"minecraft:brown_mushroom_block" `
	Arrow                      int `json:"minecraft:arrow" `
	CoalBlock                  int `json:"minecraft:coal_block" `
	DarkOakSapling             int `json:"minecraft:dark_oak_sapling" `
	GoldOre                    int `json:"minecraft:gold_ore" `
	MagmaCream                 int `json:"minecraft:magma_cream" `
	SuspiciousStew             int `json:"minecraft:suspicious_stew" `
	PinkBanner                 int `json:"minecraft:pink_banner" `
	Honeycomb                  int `json:"minecraft:honeycomb" `
	EnchantedBook              int `json:"minecraft:enchanted_book" `
	MagentaCarpet              int `json:"minecraft:magenta_carpet" `
	PrismarineBrickSlab        int `json:"minecraft:prismarine_brick_slab" `
	Redstone                   int `json:"minecraft:redstone" `
	BlackConcrete              int `json:"minecraft:black_concrete" `
	Beetroot                   int `json:"minecraft:beetroot" `
	IronIngot                  int `json:"minecraft:iron_ingot" `
	Saddle                     int `json:"minecraft:saddle" `
	FilledMap                  int `json:"minecraft:filled_map" `
	GreenCarpet                int `json:"minecraft:green_carpet" `
	SpruceBoat                 int `json:"minecraft:spruce_boat" `
	Piston                     int `json:"minecraft:piston" `
	ChainmailHelmet            int `json:"minecraft:chainmail_helmet" `
	BeetrootSeeds              int `json:"minecraft:beetroot_seeds" `
	DarkOakSign                int `json:"minecraft:dark_oak_sign" `
	PurpurStairs               int `json:"minecraft:purpur_stairs" `
	Cactus                     int `json:"minecraft:cactus" `
	Stone                      int `json:"minecraft:stone" `
	RedstoneOre                int `json:"minecraft:redstone_ore" `
	MusicDiscWard              int `json:"minecraft:music_disc_ward" `
	BirchButton                int `json:"minecraft:birch_button" `
	BrewingStand               int `json:"minecraft:brewing_stand" `
	DarkOakFence               int `json:"minecraft:dark_oak_fence" `
	Wheat                      int `json:"minecraft:wheat" `
	BirchFence                 int `json:"minecraft:birch_fence" `
	ChorusFruit                int `json:"minecraft:chorus_fruit" `
	Composter                  int `json:"minecraft:composter" `
	EmeraldBlock               int `json:"minecraft:emerald_block" `
	CookedCod                  int `json:"minecraft:cooked_cod" `
	HornCoral                  int `json:"minecraft:horn_coral" `
	EndStoneBricks             int `json:"minecraft:end_stone_bricks" `
	DarkOakStairs              int `json:"minecraft:dark_oak_stairs" `
	TripwireHook               int `json:"minecraft:tripwire_hook" `
	BirchPlanks                int `json:"minecraft:birch_planks" `
	Repeater                   int `json:"minecraft:repeater" `
	BlueWool                   int `json:"minecraft:blue_wool" `
	SlimeBlock                 int `json:"minecraft:slime_block" `
	WhiteTulip                 int `json:"minecraft:white_tulip" `
	Emerald                    int `json:"minecraft:emerald" `
	OxeyeDaisy                 int `json:"minecraft:oxeye_daisy" `
	JungleSlab                 int `json:"minecraft:jungle_slab" `
	Furnace                    int `json:"minecraft:furnace" `
	WhiteBed                   int `json:"minecraft:white_bed" `
	BlueOrchid                 int `json:"minecraft:blue_orchid" `
	OrangeGlazedTerracotta     int `json:"minecraft:orange_glazed_terracotta" `
	PolishedGraniteSlab        int `json:"minecraft:polished_granite_slab" `
	GraniteStairs              int `json:"minecraft:granite_stairs" `
	StoneBrickSlab             int `json:"minecraft:stone_brick_slab" `
	Vine                       int `json:"minecraft:vine" `
	TubeCoralFan               int `json:"minecraft:tube_coral_fan" `
	Dirt                       int `json:"minecraft:dirt" `
	BirchDoor                  int `json:"minecraft:birch_door" `
	StonePickaxe               int `json:"minecraft:stone_pickaxe" `
	IronPickaxe                int `json:"minecraft:iron_pickaxe" `
	YellowWool                 int `json:"minecraft:yellow_wool" `
	Ice                        int `json:"minecraft:ice" `
	StoneSlab                  int `json:"minecraft:stone_slab" `
	ZombieHead                 int `json:"minecraft:zombie_head" `
	Leather                    int `json:"minecraft:leather" `
	Poppy                      int `json:"minecraft:poppy" `
	DarkOakTrapdoor            int `json:"minecraft:dark_oak_trapdoor" `
	IronOre                    int `json:"minecraft:iron_ore" `
	GlowstoneDust              int `json:"minecraft:glowstone_dust" `
	HeavyWeightedPressurePlate int `json:"minecraft:heavy_weighted_pressure_plate" `
	DarkOakDoor                int `json:"minecraft:dark_oak_door" `
	TubeCoralBlock             int `json:"minecraft:tube_coral_block" `
	TrappedChest               int `json:"minecraft:trapped_chest" `
	ItemFrame                  int `json:"minecraft:item_frame" `
	Torch                      int `json:"minecraft:torch" `
	Potato                     int `json:"minecraft:potato" `
	PurpleWool                 int `json:"minecraft:purple_wool" `
	PurpleBed                  int `json:"minecraft:purple_bed" `
	JungleTrapdoor             int `json:"minecraft:jungle_trapdoor" `
	EndRod                     int `json:"minecraft:end_rod" `
	NetherBrickSlab            int `json:"minecraft:nether_brick_slab" `
	PoisonousPotato            int `json:"minecraft:poisonous_potato" `
	GlassBottle                int `json:"minecraft:glass_bottle" `
	IronNugget                 int `json:"minecraft:iron_nugget" `
	BlueBed                    int `json:"minecraft:blue_bed" `
	OakLog                     int `json:"minecraft:oak_log" `
	IronBoots                  int `json:"minecraft:iron_boots" `
	CocoaBeans                 int `json:"minecraft:cocoa_beans" `
	StoneBrickStairs           int `json:"minecraft:stone_brick_stairs" `
	RedWool                    int `json:"minecraft:red_wool" `
	Andesite                   int `json:"minecraft:andesite" `
	DiamondHelmet              int `json:"minecraft:diamond_helmet" `
	HayBlock                   int `json:"minecraft:hay_block" `
	IronBars                   int `json:"minecraft:iron_bars" `
	MossyCobblestoneStairs     int `json:"minecraft:mossy_cobblestone_stairs" `
	ActivatorRail              int `json:"minecraft:activator_rail" `
	CutSandstone               int `json:"minecraft:cut_sandstone" `
	Coal                       int `json:"minecraft:coal" `
	MusicDiscStrad             int `json:"minecraft:music_disc_strad" `
	IronLeggings               int `json:"minecraft:iron_leggings" `
	Bucket                     int `json:"minecraft:bucket" `
	LeatherHelmet              int `json:"minecraft:leather_helmet" `
	PoweredRail                int `json:"minecraft:powered_rail" `
	MelonSeeds                 int `json:"minecraft:melon_seeds" `
	PurpleStainedGlass         int `json:"minecraft:purple_stained_glass" `
	PolishedGranite            int `json:"minecraft:polished_granite" `
	AcaciaPlanks               int `json:"minecraft:acacia_planks" `
	RedstoneTorch              int `json:"minecraft:redstone_torch" `
	InkSac                     int `json:"minecraft:ink_sac" `
	Bread                      int `json:"minecraft:bread" `
	GoldenBoots                int `json:"minecraft:golden_boots" `
	BirchLog                   int `json:"minecraft:birch_log" `
	QuartzStairs               int `json:"minecraft:quartz_stairs" `
	Flint                      int `json:"minecraft:flint" `
	CrackedStoneBricks         int `json:"minecraft:cracked_stone_bricks" `
	CobblestoneSlab            int `json:"minecraft:cobblestone_slab" `
	RedShulkerBox              int `json:"minecraft:red_shulker_box" `
	Cauldron                   int `json:"minecraft:cauldron" `
	PolishedDiorite            int `json:"minecraft:polished_diorite" `
	RedTulip                   int `json:"minecraft:red_tulip" `
	Potion                     int `json:"minecraft:potion" `
	NameTag                    int `json:"minecraft:name_tag" `
	StoneButton                int `json:"minecraft:stone_button" `
	Mutton                     int `json:"minecraft:mutton" `
	Bowl                       int `json:"minecraft:bowl" `
	AzureBluet                 int `json:"minecraft:azure_bluet" `
	Prismarine                 int `json:"minecraft:prismarine" `
	OrangeWool                 int `json:"minecraft:orange_wool" `
	Sandstone                  int `json:"minecraft:sandstone" `
	SpruceFence                int `json:"minecraft:spruce_fence" `
	BirchFenceGate             int `json:"minecraft:birch_fence_gate" `
	WitherRose                 int `json:"minecraft:wither_rose" `
	OrangeTerracotta           int `json:"minecraft:orange_terracotta" `
	DiamondAxe                 int `json:"minecraft:diamond_axe" `
	Netherrack                 int `json:"minecraft:netherrack" `
	DarkOakFenceGate           int `json:"minecraft:dark_oak_fence_gate" `
	SpruceSapling              int `json:"minecraft:spruce_sapling" `
	BlueBanner                 int `json:"minecraft:blue_banner" `
	Ladder                     int `json:"minecraft:ladder" `
	Terracotta                 int `json:"minecraft:terracotta" `
	DiamondShovel              int `json:"minecraft:diamond_shovel" `
	Clock                      int `json:"minecraft:clock" `
	BirchBoat                  int `json:"minecraft:birch_boat" `
	Stick                      int `json:"minecraft:stick" `
	MagentaBanner              int `json:"minecraft:magenta_banner" `
	Cornflower                 int `json:"minecraft:cornflower" `
	IronAxe                    int `json:"minecraft:iron_axe" `
	CraftingTable              int `json:"minecraft:crafting_table" `
	Seagrass                   int `json:"minecraft:seagrass" `
	DarkOakPlanks              int `json:"minecraft:dark_oak_planks" `
	PurpleBanner               int `json:"minecraft:purple_banner" `
	SkeletonSkull              int `json:"minecraft:skeleton_skull" `
	SugarCane                  int `json:"minecraft:sugar_cane" `
	ChippedAnvil               int `json:"minecraft:chipped_anvil" `
	Gunpowder                  int `json:"minecraft:gunpowder" `
	WhiteWool                  int `json:"minecraft:white_wool" `
	Dispenser                  int `json:"minecraft:dispenser" `
	GhastTear                  int `json:"minecraft:ghast_tear" `
	LimeCarpet                 int `json:"minecraft:lime_carpet" `
	PufferfishBucket           int `json:"minecraft:pufferfish_bucket" `
	ExperienceBottle           int `json:"minecraft:experience_bottle" `
	BirchLeaves                int `json:"minecraft:birch_leaves" `
	Shears                     int `json:"minecraft:shears" `
	LapisBlock                 int `json:"minecraft:lapis_block" `
	WoodenShovel               int `json:"minecraft:wooden_shovel" `
	StickyPiston               int `json:"minecraft:sticky_piston" `
	StoneHoe                   int `json:"minecraft:stone_hoe" `
	OakButton                  int `json:"minecraft:oak_button" `
	WaterBucket                int `json:"minecraft:water_bucket" `
	OakSapling                 int `json:"minecraft:oak_sapling" `
	SeaPickle                  int `json:"minecraft:sea_pickle" `
	Allium                     int `json:"minecraft:allium" `
	WhiteCarpet                int `json:"minecraft:white_carpet" `
	Trident                    int `json:"minecraft:trident" `
	MossyStoneBricks           int `json:"minecraft:mossy_stone_bricks" `
	BrainCoralFan              int `json:"minecraft:brain_coral_fan" `
	PinkWool                   int `json:"minecraft:pink_wool" `
	Campfire                   int `json:"minecraft:campfire" `
	HoneyBlock                 int `json:"minecraft:honey_block" `
	OakFenceGate               int `json:"minecraft:oak_fence_gate" `
	DetectorRail               int `json:"minecraft:detector_rail" `
	Obsidian                   int `json:"minecraft:obsidian" `
	YellowBed                  int `json:"minecraft:yellow_bed" `
	ChiseledSandstone          int `json:"minecraft:chiseled_sandstone" `
	PrismarineBricks           int `json:"minecraft:prismarine_bricks" `
	RabbitHide                 int `json:"minecraft:rabbit_hide" `
	Scute                      int `json:"minecraft:scute" `
	CookedBeef                 int `json:"minecraft:cooked_beef" `
	NautilusShell              int `json:"minecraft:nautilus_shell" `
	CarvedPumpkin              int `json:"minecraft:carved_pumpkin" `
	CoalOre                    int `json:"minecraft:coal_ore" `
	NetherQuartzOre            int `json:"minecraft:nether_quartz_ore" `
	StoneAxe                   int `json:"minecraft:stone_axe" `
	LimeBanner                 int `json:"minecraft:lime_banner" `
	WitherSkeletonSkull        int `json:"minecraft:wither_skeleton_skull" `
	NetherBrick                int `json:"minecraft:nether_brick" `
	Diamond                    int `json:"minecraft:diamond" `
	Grindstone                 int `json:"minecraft:grindstone" `
	Feather                    int `json:"minecraft:feather" `
	LeatherBoots               int `json:"minecraft:leather_boots" `
	PrismarineShard            int `json:"minecraft:prismarine_shard" `
	BubbleCoralFan             int `json:"minecraft:bubble_coral_fan" `
	TubeCoral                  int `json:"minecraft:tube_coral" `
	PrismarineWall             int `json:"minecraft:prismarine_wall" `
	CyanWool                   int `json:"minecraft:cyan_wool" `
	CookedChicken              int `json:"minecraft:cooked_chicken" `
	StrippedDarkOakLog         int `json:"minecraft:stripped_dark_oak_log" `
	YellowTerracotta           int `json:"minecraft:yellow_terracotta" `
	Fern                       int `json:"minecraft:fern" `
	JungleFence                int `json:"minecraft:jungle_fence" `
	BlazeRod                   int `json:"minecraft:blaze_rod" `
	GoldIngot                  int `json:"minecraft:gold_ingot" `
	SeaLantern                 int `json:"minecraft:sea_lantern" `
	IronHelmet                 int `json:"minecraft:iron_helmet" `
	Sugar                      int `json:"minecraft:sugar" `
	RedstoneLamp               int `json:"minecraft:redstone_lamp" `
	DiamondBoots               int `json:"minecraft:diamond_boots" `
	BlastFurnace               int `json:"minecraft:blast_furnace" `
	RedstoneBlock              int `json:"minecraft:redstone_block" `
	GoldenHorseArmor           int `json:"minecraft:golden_horse_armor" `
	JungleSapling              int `json:"minecraft:jungle_sapling" `
	FireCoralBlock             int `json:"minecraft:fire_coral_block" `
}

type Killed struct {
	Turtle          int `json:"minecraft:turtle" `
	Pillager        int `json:"minecraft:pillager" `
	WanderingTrader int `json:"minecraft:wandering_trader" `
	Chicken         int `json:"minecraft:chicken" `
	ElderGuardian   int `json:"minecraft:elder_guardian" `
	Enderman        int `json:"minecraft:enderman" `
	Creeper         int `json:"minecraft:creeper" `
	Witch           int `json:"minecraft:witch" `
	Bat             int `json:"minecraft:bat" `
	PolarBear       int `json:"minecraft:polar_bear" `
	IronGolem       int `json:"minecraft:iron_golem" `
	Spider          int `json:"minecraft:spider" `
	Vindicator      int `json:"minecraft:vindicator" `
	CaveSpider      int `json:"minecraft:cave_spider" `
	Ravager         int `json:"minecraft:ravager" `
	Skeleton        int `json:"minecraft:skeleton" `
	Vex             int `json:"minecraft:vex" `
	Sheep           int `json:"minecraft:sheep" `
	SnowGolem       int `json:"minecraft:snow_golem" `
	Squid           int `json:"minecraft:squid" `
	MagmaCube       int `json:"minecraft:magma_cube" `
	Pufferfish      int `json:"minecraft:pufferfish" `
	Drowned         int `json:"minecraft:drowned" `
	Evoker          int `json:"minecraft:evoker" `
	Phantom         int `json:"minecraft:phantom" `
	Guardian        int `json:"minecraft:guardian" `
	Zombie          int `json:"minecraft:zombie" `
	Cow             int `json:"minecraft:cow" `
	Wither          int `json:"minecraft:wither" `
	Slime           int `json:"minecraft:slime" `
	ZombieVillager  int `json:"minecraft:zombie_villager" `
	SkeletonHorse   int `json:"minecraft:skeleton_horse" `
	WitherSkeleton  int `json:"minecraft:wither_skeleton" `
	Villager        int `json:"minecraft:villager" `
	Salmon          int `json:"minecraft:salmon" `
	Husk            int `json:"minecraft:husk" `
	Blaze           int `json:"minecraft:blaze" `
	Shulker         int `json:"minecraft:shulker" `
	ZombiePigman    int `json:"minecraft:zombie_pigman" `
	Silverfish      int `json:"minecraft:silverfish" `
	Ghast           int `json:"minecraft:ghast" `
}

type Crafted struct {
	StoneShovel                int `json:"minecraft:stone_shovel" `
	BoneBlock                  int `json:"minecraft:bone_block" `
	Hopper                     int `json:"minecraft:hopper" `
	GlisteringMelonSlice       int `json:"minecraft:glistering_melon_slice" `
	BlackConcretePowder        int `json:"minecraft:black_concrete_powder" `
	OrangeCarpet               int `json:"minecraft:orange_carpet" `
	Stonecutter                int `json:"minecraft:stonecutter" `
	FlintAndSteel              int `json:"minecraft:flint_and_steel" `
	SmithingTable              int `json:"minecraft:smithing_table" `
	SpruceSlab                 int `json:"minecraft:spruce_slab" `
	Bell                       int `json:"minecraft:bell" `
	LeatherChestplate          int `json:"minecraft:leather_chestplate" `
	ShulkerBox                 int `json:"minecraft:shulker_box" `
	DiamondSword               int `json:"minecraft:diamond_sword" `
	Beacon                     int `json:"minecraft:beacon" `
	CoarseDirt                 int `json:"minecraft:coarse_dirt" `
	StoneStairs                int `json:"minecraft:stone_stairs" `
	Compass                    int `json:"minecraft:compass" `
	BlackWool                  int `json:"minecraft:black_wool" `
	JungleSign                 int `json:"minecraft:jungle_sign" `
	Observer                   int `json:"minecraft:observer" `
	IronBlock                  int `json:"minecraft:iron_block" `
	IronTrapdoor               int `json:"minecraft:iron_trapdoor" `
	Lead                       int `json:"minecraft:lead" `
	BlackCarpet                int `json:"minecraft:black_carpet" `
	CyanBanner                 int `json:"minecraft:cyan_banner" `
	DiamondChestplate          int `json:"minecraft:diamond_chestplate" `
	CookedMutton               int `json:"minecraft:cooked_mutton" `
	Loom                       int `json:"minecraft:loom" `
	IronDoor                   int `json:"minecraft:iron_door" `
	Scaffolding                int `json:"minecraft:scaffolding" `
	Barrel                     int `json:"minecraft:barrel" `
	GoldNugget                 int `json:"minecraft:gold_nugget" `
	Anvil                      int `json:"minecraft:anvil" `
	CookedSalmon               int `json:"minecraft:cooked_salmon" `
	PumpkinPie                 int `json:"minecraft:pumpkin_pie" `
	HopperMinecart             int `json:"minecraft:hopper_minecart" `
	Minecart                   int `json:"minecraft:minecart" `
	JungleButton               int `json:"minecraft:jungle_button" `
	StonePressurePlate         int `json:"minecraft:stone_pressure_plate" `
	Lectern                    int `json:"minecraft:lectern" `
	Lever                      int `json:"minecraft:lever" `
	ChainmailLeggings          int `json:"minecraft:chainmail_leggings" `
	SprucePlanks               int `json:"minecraft:spruce_planks" `
	FermentedSpiderEye         int `json:"minecraft:fermented_spider_eye" `
	JungleFenceGate            int `json:"minecraft:jungle_fence_gate" `
	WritableBook               int `json:"minecraft:writable_book" `
	Rail                       int `json:"minecraft:rail" `
	Smoker                     int `json:"minecraft:smoker" `
	StoneBricks                int `json:"minecraft:stone_bricks" `
	YellowDye                  int `json:"minecraft:yellow_dye" `
	EnderChest                 int `json:"minecraft:ender_chest" `
	QuartzBlock                int `json:"minecraft:quartz_block" `
	RedBed                     int `json:"minecraft:red_bed" `
	Glass                      int `json:"minecraft:glass" `
	PinkBed                    int `json:"minecraft:pink_bed" `
	BirchPressurePlate         int `json:"minecraft:birch_pressure_plate" `
	Charcoal                   int `json:"minecraft:charcoal" `
	StoneSword                 int `json:"minecraft:stone_sword" `
	Book                       int `json:"minecraft:book" `
	Paper                      int `json:"minecraft:paper" `
	WhiteBanner                int `json:"minecraft:white_banner" `
	Bookshelf                  int `json:"minecraft:bookshelf" `
	LapisLazuli                int `json:"minecraft:lapis_lazuli" `
	LeatherLeggings            int `json:"minecraft:leather_leggings" `
	NoteBlock                  int `json:"minecraft:note_block" `
	JungleBoat                 int `json:"minecraft:jungle_boat" `
	Shield                     int `json:"minecraft:shield" `
	Comparator                 int `json:"minecraft:comparator" `
	Brick                      int `json:"minecraft:brick" `
	SlimeBall                  int `json:"minecraft:slime_ball" `
	DiamondBlock               int `json:"minecraft:diamond_block" `
	BrownBed                   int `json:"minecraft:brown_bed" `
	JunglePlanks               int `json:"minecraft:jungle_planks" `
	OakPlanks                  int `json:"minecraft:oak_planks" `
	GrayBed                    int `json:"minecraft:gray_bed" `
	FlowerPot                  int `json:"minecraft:flower_pot" `
	CobblestoneStairs          int `json:"minecraft:cobblestone_stairs" `
	TropicalFishBucket         int `json:"minecraft:tropical_fish_bucket" `
	BlueDye                    int `json:"minecraft:blue_dye" `
	Air                        int `json:"minecraft:air" `
	RabbitStew                 int `json:"minecraft:rabbit_stew" `
	BirchSign                  int `json:"minecraft:birch_sign" `
	OakSlab                    int `json:"minecraft:oak_slab" `
	Melon                      int `json:"minecraft:melon" `
	Beehive                    int `json:"minecraft:beehive" `
	Dropper                    int `json:"minecraft:dropper" `
	Glowstone                  int `json:"minecraft:glowstone" `
	Cake                       int `json:"minecraft:cake" `
	HoneycombBlock             int `json:"minecraft:honeycomb_block" `
	GoldenCarrot               int `json:"minecraft:golden_carrot" `
	SpruceStairs               int `json:"minecraft:spruce_stairs" `
	NetherWartBlock            int `json:"minecraft:nether_wart_block" `
	Bow                        int `json:"minecraft:bow" `
	NetherBrickFence           int `json:"minecraft:nether_brick_fence" `
	OakFence                   int `json:"minecraft:oak_fence" `
	DiamondPickaxe             int `json:"minecraft:diamond_pickaxe" `
	BlazePowder                int `json:"minecraft:blaze_powder" `
	IronShovel                 int `json:"minecraft:iron_shovel" `
	OakBoat                    int `json:"minecraft:oak_boat" `
	BlackBed                   int `json:"minecraft:black_bed" `
	DaylightDetector           int `json:"minecraft:daylight_detector" `
	WhiteGlazedTerracotta      int `json:"minecraft:white_glazed_terracotta" `
	BoneMeal                   int `json:"minecraft:bone_meal" `
	GoldenApple                int `json:"minecraft:golden_apple" `
	IronChestplate             int `json:"minecraft:iron_chestplate" `
	CartographyTable           int `json:"minecraft:cartography_table" `
	NetherBricks               int `json:"minecraft:nether_bricks" `
	GreenDye                   int `json:"minecraft:green_dye" `
	SpruceSign                 int `json:"minecraft:spruce_sign" `
	Lantern                    int `json:"minecraft:lantern" `
	PumpkinSeeds               int `json:"minecraft:pumpkin_seeds" `
	OakPressurePlate           int `json:"minecraft:oak_pressure_plate" `
	SpruceTrapdoor             int `json:"minecraft:spruce_trapdoor" `
	GreenWool                  int `json:"minecraft:green_wool" `
	FishingRod                 int `json:"minecraft:fishing_rod" `
	IronHoe                    int `json:"minecraft:iron_hoe" `
	DarkOakButton              int `json:"minecraft:dark_oak_button" `
	OakTrapdoor                int `json:"minecraft:oak_trapdoor" `
	OrangeBanner               int `json:"minecraft:orange_banner" `
	Chest                      int `json:"minecraft:chest" `
	GoldBlock                  int `json:"minecraft:gold_block" `
	Map                        int `json:"minecraft:map" `
	SmoothStone                int `json:"minecraft:smooth_stone" `
	BirchTrapdoor              int `json:"minecraft:birch_trapdoor" `
	SpruceDoor                 int `json:"minecraft:spruce_door" `
	OakDoor                    int `json:"minecraft:oak_door" `
	GreenBanner                int `json:"minecraft:green_banner" `
	MossyCobblestone           int `json:"minecraft:mossy_cobblestone" `
	IronSword                  int `json:"minecraft:iron_sword" `
	OakSign                    int `json:"minecraft:oak_sign" `
	GlassPane                  int `json:"minecraft:glass_pane" `
	WoodenPickaxe              int `json:"minecraft:wooden_pickaxe" `
	SnowBlock                  int `json:"minecraft:snow_block" `
	WhiteShulkerBox            int `json:"minecraft:white_shulker_box" `
	SuspiciousStew             int `json:"minecraft:suspicious_stew" `
	CoalBlock                  int `json:"minecraft:coal_block" `
	JackOLantern               int `json:"minecraft:jack_o_lantern" `
	DiamondLeggings            int `json:"minecraft:diamond_leggings" `
	Arrow                      int `json:"minecraft:arrow" `
	Piston                     int `json:"minecraft:piston" `
	ChainmailHelmet            int `json:"minecraft:chainmail_helmet" `
	EnchantingTable            int `json:"minecraft:enchanting_table" `
	GreenCarpet                int `json:"minecraft:green_carpet" `
	Redstone                   int `json:"minecraft:redstone" `
	IronIngot                  int `json:"minecraft:iron_ingot" `
	EnchantedBook              int `json:"minecraft:enchanted_book" `
	EmeraldBlock               int `json:"minecraft:emerald_block" `
	CookedCod                  int `json:"minecraft:cooked_cod" `
	Composter                  int `json:"minecraft:composter" `
	Wheat                      int `json:"minecraft:wheat" `
	BirchFence                 int `json:"minecraft:birch_fence" `
	Stone                      int `json:"minecraft:stone" `
	BrewingStand               int `json:"minecraft:brewing_stand" `
	BirchButton                int `json:"minecraft:birch_button" `
	DarkOakSign                int `json:"minecraft:dark_oak_sign" `
	Furnace                    int `json:"minecraft:furnace" `
	JungleSlab                 int `json:"minecraft:jungle_slab" `
	WhiteBed                   int `json:"minecraft:white_bed" `
	Emerald                    int `json:"minecraft:emerald" `
	Repeater                   int `json:"minecraft:repeater" `
	BirchPlanks                int `json:"minecraft:birch_planks" `
	SlimeBlock                 int `json:"minecraft:slime_block" `
	BlueWool                   int `json:"minecraft:blue_wool" `
	HeavyWeightedPressurePlate int `json:"minecraft:heavy_weighted_pressure_plate" `
	DarkOakTrapdoor            int `json:"minecraft:dark_oak_trapdoor" `
	BrownDye                   int `json:"minecraft:brown_dye" `
	DarkOakDoor                int `json:"minecraft:dark_oak_door" `
	StoneSlab                  int `json:"minecraft:stone_slab" `
	CookedPorkchop             int `json:"minecraft:cooked_porkchop" `
	YellowWool                 int `json:"minecraft:yellow_wool" `
	StonePickaxe               int `json:"minecraft:stone_pickaxe" `
	BirchDoor                  int `json:"minecraft:birch_door" `
	IronPickaxe                int `json:"minecraft:iron_pickaxe" `
	JungleTrapdoor             int `json:"minecraft:jungle_trapdoor" `
	PurpleBed                  int `json:"minecraft:purple_bed" `
	Torch                      int `json:"minecraft:torch" `
	ItemFrame                  int `json:"minecraft:item_frame" `
	TubeCoralBlock             int `json:"minecraft:tube_coral_block" `
	TrappedChest               int `json:"minecraft:trapped_chest" `
	IronBars                   int `json:"minecraft:iron_bars" `
	HayBlock                   int `json:"minecraft:hay_block" `
	RedWool                    int `json:"minecraft:red_wool" `
	DiamondHelmet              int `json:"minecraft:diamond_helmet" `
	IronBoots                  int `json:"minecraft:iron_boots" `
	RedDye                     int `json:"minecraft:red_dye" `
	BlueBed                    int `json:"minecraft:blue_bed" `
	IronNugget                 int `json:"minecraft:iron_nugget" `
	GlassBottle                int `json:"minecraft:glass_bottle" `
	Cauldron                   int `json:"minecraft:cauldron" `
	RedShulkerBox              int `json:"minecraft:red_shulker_box" `
	CobblestoneSlab            int `json:"minecraft:cobblestone_slab" `
	RedstoneTorch              int `json:"minecraft:redstone_torch" `
	Podzol                     int `json:"minecraft:podzol" `
	PoweredRail                int `json:"minecraft:powered_rail" `
	Bucket                     int `json:"minecraft:bucket" `
	FireworkRocket             int `json:"minecraft:firework_rocket" `
	PolishedGranite            int `json:"minecraft:polished_granite" `
	MelonSeeds                 int `json:"minecraft:melon_seeds" `
	Coal                       int `json:"minecraft:coal" `
	IronLeggings               int `json:"minecraft:iron_leggings" `
	SpruceFence                int `json:"minecraft:spruce_fence" `
	NameTag                    int `json:"minecraft:name_tag" `
	EnderEye                   int `json:"minecraft:ender_eye" `
	StoneButton                int `json:"minecraft:stone_button" `
	PolishedDiorite            int `json:"minecraft:polished_diorite" `
	ArmorStand                 int `json:"minecraft:armor_stand" `
	BlueBanner                 int `json:"minecraft:blue_banner" `
	SpruceSapling              int `json:"minecraft:spruce_sapling" `
	Ladder                     int `json:"minecraft:ladder" `
	DiamondAxe                 int `json:"minecraft:diamond_axe" `
	DarkOakFenceGate           int `json:"minecraft:dark_oak_fence_gate" `
	Clay                       int `json:"minecraft:clay" `
	BirchFenceGate             int `json:"minecraft:birch_fence_gate" `
	CraftingTable              int `json:"minecraft:crafting_table" `
	PurpleBanner               int `json:"minecraft:purple_banner" `
	DarkOakPlanks              int `json:"minecraft:dark_oak_planks" `
	Stick                      int `json:"minecraft:stick" `
	IronAxe                    int `json:"minecraft:iron_axe" `
	Clock                      int `json:"minecraft:clock" `
	DiamondShovel              int `json:"minecraft:diamond_shovel" `
	BirchBoat                  int `json:"minecraft:birch_boat" `
	OakButton                  int `json:"minecraft:oak_button" `
	StoneHoe                   int `json:"minecraft:stone_hoe" `
	Shears                     int `json:"minecraft:shears" `
	StickyPiston               int `json:"minecraft:sticky_piston" `
	LapisBlock                 int `json:"minecraft:lapis_block" `
	Dispenser                  int `json:"minecraft:dispenser" `
	Gunpowder                  int `json:"minecraft:gunpowder" `
	WhiteWool                  int `json:"minecraft:white_wool" `
	PufferfishBucket           int `json:"minecraft:pufferfish_bucket" `
	FletchingTable             int `json:"minecraft:fletching_table" `
	LimeCarpet                 int `json:"minecraft:lime_carpet" `
	OakFenceGate               int `json:"minecraft:oak_fence_gate" `
	HoneyBlock                 int `json:"minecraft:honey_block" `
	AcaciaSapling              int `json:"minecraft:acacia_sapling" `
	YellowBed                  int `json:"minecraft:yellow_bed" `
	MossyStoneBricks           int `json:"minecraft:mossy_stone_bricks" `
	Campfire                   int `json:"minecraft:campfire" `
	SeaPickle                  int `json:"minecraft:sea_pickle" `
	ChestMinecart              int `json:"minecraft:chest_minecart" `
	WhiteCarpet                int `json:"minecraft:white_carpet" `
	BlueIce                    int `json:"minecraft:blue_ice" `
	LimeBanner                 int `json:"minecraft:lime_banner" `
	Diamond                    int `json:"minecraft:diamond" `
	BlackDye                   int `json:"minecraft:black_dye" `
	Grindstone                 int `json:"minecraft:grindstone" `
	NetherBrick                int `json:"minecraft:nether_brick" `
	StoneAxe                   int `json:"minecraft:stone_axe" `
	DarkOakPressurePlate       int `json:"minecraft:dark_oak_pressure_plate" `
	CookedBeef                 int `json:"minecraft:cooked_beef" `
	CookedChicken              int `json:"minecraft:cooked_chicken" `
	WhiteDye                   int `json:"minecraft:white_dye" `
	LeatherBoots               int `json:"minecraft:leather_boots" `
	RedstoneBlock              int `json:"minecraft:redstone_block" `
	DiamondBoots               int `json:"minecraft:diamond_boots" `
	BlastFurnace               int `json:"minecraft:blast_furnace" `
	TurtleHelmet               int `json:"minecraft:turtle_helmet" `
	RedstoneLamp               int `json:"minecraft:redstone_lamp" `
	Sugar                      int `json:"minecraft:sugar" `
	CodBucket                  int `json:"minecraft:cod_bucket" `
	IronHelmet                 int `json:"minecraft:iron_helmet" `
	GoldIngot                  int `json:"minecraft:gold_ingot" `
	Fern                       int `json:"minecraft:fern" `
	JungleFence                int `json:"minecraft:jungle_fence" `
}
