package main

import (
	"context"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type MediaConvertOptions struct {
	TargetFormat  string `json:"target_format"`
	TargetBitrate struct {
		Video string `json:"video"`
		Audio string `json:"audio"`
	} `json:"target_bitrate"`
}

// MediaConvert converts the given media file to the given format
func (a *App) MediaConvert(options MediaConvertOptions) error {
	fileSource, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Videos",
				Pattern:     "*.mp4;*.mkv;*.avi;*.mov;*.wmv;*.flv;*.webm",
			},
			{
				DisplayName: "Audio",
				Pattern:     "*.mp3;*.wav;*.aac;*.flac;*.ogg;*.wma;*.m4a;*.opus",
			},
		},
	})
	if err != nil {
		runtime.EventsEmit(a.ctx, "failed", err.Error())
		return err
	}
	if fileSource == "" {
		runtime.EventsEmit(a.ctx, "failed", "No file selected")
		return nil
	}

	targetDir, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save file...",
		Filters: []runtime.FileFilter{
			{
				DisplayName: options.TargetFormat,
				Pattern:     "*." + options.TargetFormat,
			},
		},
		DefaultFilename: strings.Split(strings.Split(fileSource, "\\")[len(strings.Split(fileSource, "\\"))-1], ".")[0],
	})
	if err != nil {
		runtime.EventsEmit(a.ctx, "failed", err.Error())
		return err
	}
	if targetDir == "" {
		runtime.EventsEmit(a.ctx, "failed", "No target directory selected")
		return nil
	}

	targetDir = targetDir + "." + options.TargetFormat

	err = ffmpeg_go.Input(fileSource).
		Output(targetDir, ffmpeg_go.KwArgs{
			"b:v": options.TargetBitrate.Video,
			"b:a": options.TargetBitrate.Audio,
		}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()
	if err != nil {
		runtime.EventsEmit(a.ctx, "failed", err.Error())
		return err
	}

	log.Println("File converted to: " + targetDir)
	runtime.EventsEmit(a.ctx, "success", "File converted to: "+targetDir)

	return nil
}

type ImageConvertOptions struct {
	TargetFormat string `json:"target_format"`
}

func (a *App) ImageConvert(options ImageConvertOptions) error {
	fileSource, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images",
				Pattern:     "*.jpg;*.jpeg;*.png;*.gif;*.bmp;*.webp",
			},
			{
				DisplayName: "Videos",
				Pattern:     "*.mp4;*.mkv;*.avi;*.mov;*.wmv;*.flv;*.webm",
			},
		},
	})
	if err != nil {
		runtime.EventsEmit(a.ctx, "failed", err.Error())
		return err
	}
	if fileSource == "" {
		runtime.EventsEmit(a.ctx, "failed", "No file selected")
		return nil
	}

	targetDir, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save file...",
		Filters: []runtime.FileFilter{
			{
				DisplayName: options.TargetFormat,
				Pattern:     "*." + options.TargetFormat,
			},
		},
		DefaultFilename: strings.Split(strings.Split(fileSource, "\\")[len(strings.Split(fileSource, "\\"))-1], ".")[0],
	})
	if err != nil {
		runtime.EventsEmit(a.ctx, "failed", err.Error())
		return err
	}
	if targetDir == "" {
		runtime.EventsEmit(a.ctx, "failed", "No target directory selected")
		return nil
	}

	targetDir = targetDir + "." + options.TargetFormat

	err = ffmpeg_go.Input(fileSource).
		Output(targetDir).
		OverWriteOutput().
		ErrorToStdOut().
		Run()
	if err != nil {
		runtime.EventsEmit(a.ctx, "failed", err.Error())
		return err
	}

	log.Println("File converted to: " + targetDir)
	runtime.EventsEmit(a.ctx, "success", "File converted to: "+targetDir)

	return nil
}
