package export

import (
	"encoding/csv"
	"gin/pkg/setting"
	"os"
)

func GetExcelFullUrl(name string) string  {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string  {
	return setting.AppSetting.ExportSavePath;
}

func GetExcelFullPath() string  {
	return setting.AppSetting.RuntimeRootPath+GetExcelPath()
}

func ExportCsv(fileName string,data [][]string) (filePath string,err error) {

	filePath = GetExcelFullPath()+fileName+".csv"
	f, err := os.Create(filePath)

	if err !=nil {
		return filePath,err
	}

	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	w :=csv.NewWriter(f)
	w.WriteAll(data)

	return filePath,nil
}