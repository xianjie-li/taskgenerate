package g

import (
	"encoding/json"
	"fmt"
	"github.com/shogo82148/go-shuffle"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RandCategory 随机分类，对应 Task.randCategory
type RandCategory struct {
	Label string  `json:"label"` // 选项名
	Ratio float64 `json:"ratio"` // 出现概率
}

// Task 描述一项任务
type Task struct {
	Label        string           `json:"label"`                  // 任务名
	RandCategory [][]RandCategory `json:"randCategory,omitempty"` // 在提供的类型中随机选择一项, label 中通过$category1、$category2来注入对应项的名称
	Probability  float64          `json:"probability,omitempty"`  // 任务的出现概率
	IsWeekTask   bool             `json:"isWeekTask,omitempty"`   // 是否为周任务，周任务固定在周末出现
}

type TaskConfig struct {
	RestDayProbability float64 `json:"restDayProbability"` // 休息日出现的概率
	Tasks              []Task  `json:"tasks"`              // 任务列表
}

func G() {
	rand.Seed(time.Now().Unix())
	var input string
	config, err := readConfig()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		if getProportion() <= config.RestDayProbability {
			fmt.Println("✨ 今天是休息日，好好休息吧 ✨")
		} else {
			tasks := generate(config.Tasks)

			if len(tasks) == 0 {
				fmt.Println("✨ 今天没有要进行的任务 ✨")
			} else {
				fmt.Printf("已从 %d 项任务配置中为您生成了 %d 个任务!\r\n", len(config.Tasks), len(tasks))

				for i, task := range tasks {
					fmt.Println(strconv.Itoa(i+1)+".", task.Label)
				}
			}
		}
	}

	fmt.Println("请按Enter退出...")

	fmt.Scanln(&input)
}

// readConfig 从配置文件中读取配置
func readConfig() (TaskConfig, error) {
	file, err := ioutil.ReadFile("./task.json")
	if err != nil {
		return TaskConfig{}, err
	}

	var tasks TaskConfig

	if err = json.Unmarshal(file, &tasks); err != nil {
		return TaskConfig{}, err
	}

	return tasks, nil
}

// getProportion 返回0 - 100的值
func getProportion() float64 {
	return float64(rand.Intn(101))
}

// generate 根据任务列表配置生成任务
func generate(tasks []Task) []Task {
	newList := make([]Task, 0)
	isWeekend := time.Now().Weekday() == time.Sunday

	shuffleTask := make([]Task, len(tasks))
	copy(shuffleTask, tasks)
	shuffle.Slice(shuffleTask)

	for _, task := range shuffleTask {
		probabilityRand := getProportion()

		// 跳过
		if task.Probability != 0 && probabilityRand > task.Probability {
			continue
		}

		// 处理随机选项，并更改label
		if len(task.RandCategory) != 0 {
			var labelValues []string

			for _, categories := range task.RandCategory {
				categoryRand := getProportion()

				// 取值为0-99
				if categoryRand == 100 {
					categoryRand--
				}

				lastRatio := 0.0

				for _, category := range categories {
					if categoryRand < category.Ratio+lastRatio {
						labelValues = append(labelValues, category.Label)
						break
					}

					lastRatio += category.Ratio
				}

			}

			for i, value := range labelValues {
				task.Label = strings.ReplaceAll(task.Label, "$category"+strconv.Itoa(i+1), value)
			}
		}

		// 周任务直接添加
		if task.IsWeekTask {
			if isWeekend {
				newList = append(newList, task)
			}
		} else {
			newList = append(newList, task)
		}
	}

	return newList
}
