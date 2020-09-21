package delete

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func Delete(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, v := range files {
		if matchFile(v.Name()) {
			p := filepath.Join(dir, v.Name())
			if err = os.Remove(p); err != nil {
				return err
			}
		}
	}
	return nil
}

///_(\d{4})-(\d{2})-(\d{2})-\d{2}_\d{2}.mp3$/u
func matchFile(str string) bool {
	r, _ := regexp.Compile(`(\d{4})-(\d{2})-(\d{2})-\d{2}_\d{2}.mp3$`)
	m := r.FindStringSubmatch(str)
	if len(m) == 0 {
		return false
	}

	year, _ := strconv.Atoi(m[1])
	month, _ := strconv.Atoi(m[2])
	day, _ := strconv.Atoi(m[3])

	fileDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if fileDate.Before(time.Now().AddDate(0, -1, 0)) {
		return true
	}
	return false

}
