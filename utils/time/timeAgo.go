package time

import (
	"errors"
	"fmt"
	"time"
)

// TimeAgo 计算时间差并返回类似“XX秒前/分钟前/小时前/天前”或具体日期的字符串。
func TimeAgo(timestamp time.Time) (string, error) {
	// 获取当前时间
	now := time.Now()

	// 检查时间戳是否有效
	if timestamp.IsZero() {
		return "", errors.New("invalid timestamp")
	}

	// 计算时间差
	duration := now.Sub(timestamp)

	// 基于时间差返回不同的描述
	switch {
	case duration < 10*time.Second:
		return "刚刚", nil
	case duration < time.Minute:
		return fmt.Sprintf("%d秒前", int(duration.Seconds())), nil
	case duration < time.Hour:
		return fmt.Sprintf("%d分钟前", int(duration.Minutes())), nil
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d小时前", int(duration.Hours())), nil
	case duration < 7*24*time.Hour:
		return fmt.Sprintf("%d天前", int(duration.Hours()/24)), nil
	case now.Year() == timestamp.Year():
		// 如果是同一年，只显示月/日/时/分
		return timestamp.Format("1月2日 15:04"), nil
	default:
		// 不同年，显示完整日期
		return timestamp.Format("2006-01-02"), nil
	}
}
