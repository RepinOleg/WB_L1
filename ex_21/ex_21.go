package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type MediaPlayer interface {
	Play(audioType, fileName string) error // Интерфейс MediaPlayer определяет метод Play для воспроизведения медиафайла определенного типа.
}

type AdvancedMediaPlayer interface {
	PlayVideo(filename string) error // Интерфейс AdvancedMediaPlayer определяет метод для воспроизведения видео.
	PlayAudio(filename string) error // Интерфейс AdvancedMediaPlayer определяет метод для воспроизведения аудио.
}

type VLCPlayer struct{} // Структура VLCPlayer представляет проигрыватель VLC.

func (v *VLCPlayer) PlayVideo(filename string) error {
	fmt.Printf("Playing video %s\n", filename)
	return nil
}

func (v *VLCPlayer) PlayAudio(filename string) error {
	fmt.Printf("Playing audio %s\n", filename)
	return nil
}

type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer // Структура MediaAdapter содержит экземпляр AdvancedMediaPlayer для адаптации к интерфейсу MediaPlayer.
}

func (m *MediaAdapter) Play(audioType, fileName string) error {
	if audioType == "vlc" {
		return m.advancedMusicPlayer.PlayVideo(fileName) // Адаптер MediaAdapter позволяет воспроизводить видео через AdvancedMediaPlayer.
	} else if audioType == "mp4" {
		return m.advancedMusicPlayer.PlayAudio(fileName) // Адаптер MediaAdapter позволяет воспроизводить аудио через AdvancedMediaPlayer.
	}
	return fmt.Errorf("invalid media type %s", audioType)
}

type AudioPlayer struct {
	mediaAdapter MediaPlayer // Структура AudioPlayer содержит адаптер для воспроизведения медиафайлов различных типов.
}

func (a *AudioPlayer) Play(audioType, filename string) error {
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file. Name: %s\n", filename)
		return nil
	} else if audioType == "vlc" || audioType == "mp4" {
		a.mediaAdapter = &MediaAdapter{&VLCPlayer{}}    // Создание экземпляра адаптера MediaAdapter для воспроизведения видео или аудио.
		return a.mediaAdapter.Play(audioType, filename) // Вызов метода Play через адаптер для воспроизведения соответствующего медиафайла.
	}
	return fmt.Errorf("invalid media type %s", audioType)
}

func main() {
	audioPlayer := &AudioPlayer{}

	err := audioPlayer.Play("mp3", "song.mp3")
	if err != nil {
		fmt.Println("Error playing mp3:", err)
	}

	err = audioPlayer.Play("vlc", "movie.vlc")
	if err != nil {
		fmt.Println("Error playing vlc:", err)
	}

	err = audioPlayer.Play("mp4", "song2.mp4")
	if err != nil {
		fmt.Println("Error playing mp4:", err)
	}
}
