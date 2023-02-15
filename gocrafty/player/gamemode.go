package player

type GameMode uint8

const (
	GameModeSurvival GameMode = iota
	GameModeCreative
	GameModeAdventure
	GameModeSpectator
)

func (g GameMode) String() string {
	switch g {
	case GameModeSurvival:
		return "Survival"
	case GameModeCreative:
		return "Creative"
	case GameModeAdventure:
		return "Adventure"
	case GameModeSpectator:
		return "Spectator"
	default:
		return "Unknown"
	}
}
