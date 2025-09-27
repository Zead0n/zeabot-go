package music

type MusicManager struct {
	musics map[string]*Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{
		musics: make(map[string]*Music),
	}
}

func (mm *MusicManager) Get(guildId string) *Music {
	music, ok := mm.musics[guildId]
	if !ok {
		music = &Music{
			player: Player{
				state: PlayerStateEnded,
			},
			queue: Queue{
				Tracks:  make([]Track, 0),
				current: &Track{},
				mode:    LoopOff,
			},
		}

		mm.musics[guildId] = music
	}

	return music
}

func (mm *MusicManager) Delete(guildId string) {
	delete(mm.musics, guildId)
}

type Music struct {
	player Player
	queue  Queue
}
