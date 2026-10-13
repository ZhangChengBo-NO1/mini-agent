package tools

import "time"

func GetTime(_ string) (string, error) {
    now := time.Now()

    return now.Format("2006-01-02 15:04:05"),nil
}