package models

type TextToReplace struct {
	Old *string `csv:"old"`
	New *string `csv:"new"`
}
