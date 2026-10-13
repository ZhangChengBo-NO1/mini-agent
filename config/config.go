package config

import "os"

var APIKey = os.Getenv("SILICONFLOW_API_KEY")

const BaseURL = "https://api.siliconflow.cn/v1"

const Model = "Qwen/Qwen2.5-7B-Instruct"
