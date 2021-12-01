package model

type Base struct {
	Url  string `json:"url"`
	Str2 string `json:"str_2"`
}

type App struct {
	Name string `json:"name"`
	Str1 string `json:"str_1"`
	Base `json:",inline"`
}
