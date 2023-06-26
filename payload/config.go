package payload

type Config struct {
	ThirdPartyAPIEndpoint     string `json:"thirdPartyAPIEndpoint"`
	OutputFileCreationPath    string `json:"outputFileCreationPath"`
	DataFetchDurationInMinute int    `json:"dataFetchDurationInMinute"`
	FileNameFormat            string `json:"fileNameFormat"`
}
