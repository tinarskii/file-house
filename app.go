package main

import (
	"context"
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type MediaConvertOptions struct {
	Source           string `json:"source"`
	TargetName       string `json:"target_name"`
	TargetFormat     string `json:"target_format"`
	TargetDir        string `json:"target_dir"`
	TargetVideoCodec string `json:"target_video_codec"`
	TargetAudioCodec string `json:"target_audio_codec"`
}

// MediaConvert converts the given media file to the given format
func (a *App) MediaConvert(options MediaConvertOptions) []byte {

	outputPath := os.ExpandEnv(options.TargetDir) + options.TargetName + "." + options.TargetFormat

	err := ffmpeg_go.Input(os.ExpandEnv(options.Source)).
		Output(outputPath, ffmpeg_go.KwArgs{
			"vcodec":   options.TargetVideoCodec,
			"acodec":   options.TargetAudioCodec,
			"movflags": "faststart",
		}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

	if err != nil {
		fmt.Println("Error converting file: " + err.Error())
		return nil
	}

	fmt.Println("File converted to: " + outputPath)

	fileContent, err := os.ReadFile(outputPath)
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		return nil
	}

	return fileContent
}

func (a *App) WriteFile(fileContent string, filePath string) error {
	err := os.WriteFile(os.ExpandEnv(filePath), []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Error writing file: " + err.Error())
		return err
	}

	fmt.Println("File written to: " + filePath)
	return nil
}

func (a *App) RemoveFile(filePath string) error {
	err := os.Remove(os.ExpandEnv(filePath))
	if err != nil {
		fmt.Println("Error removing file: " + err.Error())
		return err
	}

	fmt.Println("File removed: " + filePath)
	return nil
}
