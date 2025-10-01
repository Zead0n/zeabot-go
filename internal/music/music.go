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
			state:   PlayerStateEnded,
			queue:   make([]Track, 0),
			current: nil,
			mode:    LoopOff,
		}

		mm.musics[guildId] = music
	}

	return music
}

func (mm *MusicManager) Delete(guildId string) {
	delete(mm.musics, guildId)
}

type Music struct {
	state   PlayerState
	queue   []Track
	current *Track
	mode    LoopMode
}

func (m *Music) dequeue() (*Track, bool) {
	if len(m.queue) <= 0 {
		return &Track{}, false
	}

	next := m.queue[0]
	if len(m.queue) <= 1 {
		m.queue = m.queue[1:]
	} else {
		m.queue = []Track{}
	}

	return &next, true
}

func (m *Music) enqueue(tracks ...Track) {
	for _, track := range tracks[:min(len(tracks), 10)] {
		m.queue = append(m.queue, track)
	}
}
