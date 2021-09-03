✨✨根据配置自动选择和生成任务，治好你的强迫症!

```
已从 7 项任务配置中为您生成了 5 个任务!
1. 【日常】: dart 收藏文章 学习
2. 【日常】: 笔记复习、迁移、整理
3. 【支线】:
      * 英语(单词 10, 书籍一小节)
      * 数学(公式或公式记忆) * 1
4. 【实战任务】: 进行实战任务
5. 【探索任务】: 深入学习go携程
请按Enter退出...
```

<br>

包含以下可选配置:
* 休息日，概率出现，当前不进行任何任务
* 动态分类，从子分类中随机选择
* 每个任务都可以配置出现概率
* 周任务

<br>

使用

1. 新建目录，创建配置文件 task.json （请手动去掉注释）
```js
{
  // 休息日出现的概率
  "restDayProbability": 3,
  // 任务列表
  "tasks": [
    // 描述一项任务
    {
      // 任务名
      "label": "【日常】: 笔记复习、迁移、整理",
      // 在提供的类型中随机选择一项, label 中通过$category1、$category2来注入对应项的名称
      "randCategory": [],
      // 任务的出现概率
      "probability": 50,
      // 是否为周任务，周任务固定在周末出现
      "isWeekTask": false
    },
    {
      // 这里的 $category* 是通过下方的randCategory随机获取的
      "label": "【日常】: $category1 $category2 学习",
      "randCategory": [
        [
          {
            // 选项名
            "label": "js",
            // 出现概率
            "ratio": 40
          },
          {
            "label": "go",
            "ratio": 40
          },
          {
            "label": "dart",
            "ratio": 20
          }
        ],
        [
          {
            "label": "收藏文章",
            "ratio": 50
          },
          {
            "label": "标准库 > TODO > 生态内容",
            "ratio": 50
          }
        ]
      ],
      "probability": 85
    },
    {
      "label": "【探索任务】: go学习",
      "probability": 85
    },
    {
      "label": "【实战任务】: 进行实战任务             ",
      "probability": 90
    },
    {
      "label": "【支线】:\n      * 政治(循环学习一个章节)\n      * 英语(单词 10, 书籍一小节)\n      * 数学(公式或公式记忆) * 1  ",
      "probability": 85
    },
    {
      "label": "【周末特供】: 整理当周新的技术文章并记录到对应todo * 1",
      "isWeekTask": true
    }
  ]
}
```

2. 根据系统, 下载仓库内的 task_generate (linux/mac) 或 task_generate.exe (windows) 到配置目录并执行
