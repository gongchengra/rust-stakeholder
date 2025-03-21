package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	// 删除未使用的 strings 包
	"sync/atomic"
	"syscall"
	"time"

	"github.com/fatih/color"
)

// DevelopmentType 开发活动类型
type DevelopmentType int

const (
	Backend DevelopmentType = iota
	Frontend
	FullStack
	DevOps
	Mobile
)

// JargonLevel 技术术语级别
type JargonLevel int

const (
	Low JargonLevel = iota
	Medium
	High
	Expert
)

// Complexity 复杂度级别
type Complexity int

const (
	ComplexityLow Complexity = iota
	ComplexityMedium
	ComplexityHigh
	ComplexityExtreme
)

// SessionConfig 会话配置
type SessionConfig struct {
	devType       DevelopmentType
	jargonLevel   JargonLevel
	complexity    Complexity
	alertsEnabled bool
	projectName   string
	minimalOutput bool
	teamActivity  bool
	framework     string
}

// 全局变量
var (
	running atomic.Bool
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
)

func main() {
	config := parseArgs()
	running.Store(true)

	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		running.Store(false)
	}()

	// 清屏
	fmt.Print("\033[H\033[2J")

	// 显示启动序列
	displayBootSequence(config)

	startTime := time.Now()
	var targetDuration *time.Duration
	if duration := getDuration(); duration > 0 {
		d := time.Duration(duration) * time.Second
		targetDuration = &d
	}

	for running.Load() {
		if targetDuration != nil && time.Since(startTime) >= *targetDuration {
			break
		}

		// 根据复杂度确定同时显示的活动数量
		activitiesCount := getActivitiesCount(config.complexity)

		// 随机选择并运行活动
		activities := []func(*SessionConfig){
			runCodeAnalysis,
			runPerformanceMetrics,
			runSystemMonitoring,
			runDataProcessing,
			runNetworkActivity,
		}
		rand.Shuffle(len(activities), func(i, j int) {
			activities[i], activities[j] = activities[j], activities[i]
		})

		for i := 0; i < activitiesCount && i < len(activities); i++ {
			activities[i](config)

			// 随机暂停
			time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)

			if !running.Load() || (targetDuration != nil && time.Since(startTime) >= *targetDuration) {
				break
			}
		}

		if config.alertsEnabled && rand.Float32() < 0.1 {
			displayRandomAlert(config)
		}

		if config.teamActivity && rand.Float32() < 0.2 {
			displayTeamActivity(config)
		}
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println(green("Session terminated."))
}

func parseArgs() *SessionConfig {
	// 简化起见，这里使用默认配置
	return &SessionConfig{
		devType:       Backend,
		jargonLevel:   Medium,
		complexity:    ComplexityMedium,
		alertsEnabled: false,
		projectName:   "distributed-cluster",
		minimalOutput: false,
		teamActivity:  false,
		framework:     "",
	}
}

func getDuration() int64 {
	return 0 // 默认运行直到中断
}

func getActivitiesCount(complexity Complexity) int {
	switch complexity {
	case ComplexityLow:
		return 1
	case ComplexityMedium:
		return 2
	case ComplexityHigh:
		return 3
	case ComplexityExtreme:
		return 4
	default:
		return 2
	}
}

func displayBootSequence(config *SessionConfig) {
	fmt.Println(green("Initializing system..."))
	time.Sleep(500 * time.Millisecond)
	fmt.Println(blue("Loading configuration..."))
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Project: %s\n", yellow(config.projectName))
	if config.framework != "" {
		fmt.Printf("Framework: %s\n", yellow(config.framework))
	}
	time.Sleep(500 * time.Millisecond)
}

func runCodeAnalysis(config *SessionConfig) {
	fmt.Println(blue("Analyzing code structure..."))
	fmt.Println(yellow("Scanning dependencies..."))
}

func runPerformanceMetrics(config *SessionConfig) {
	fmt.Println(green("Collecting performance metrics..."))
	fmt.Printf("CPU Usage: %d%%\n", rand.Intn(60)+20)
}

func runSystemMonitoring(config *SessionConfig) {
	fmt.Println(blue("Monitoring system resources..."))
	fmt.Printf("Memory utilization: %d%%\n", rand.Intn(40)+30)
}

func runDataProcessing(config *SessionConfig) {
	fmt.Println(yellow("Processing data streams..."))
	fmt.Printf("Throughput: %d MB/s\n", rand.Intn(100)+50)
}

func runNetworkActivity(config *SessionConfig) {
	fmt.Println(green("Monitoring network activity..."))
	fmt.Printf("Active connections: %d\n", rand.Intn(1000)+100)
}

func displayRandomAlert(config *SessionConfig) {
	alerts := []string{
		"Warning: High memory usage detected",
		"Notice: Network latency spike observed",
		"Alert: CPU utilization above threshold",
		"Warning: Database connection pool near capacity",
	}
	fmt.Println(red(alerts[rand.Intn(len(alerts))]))
}

func displayTeamActivity(config *SessionConfig) {
	activities := []string{
		"Team member pushing changes to repository",
		"Code review requested for feature branch",
		"CI/CD pipeline triggered by recent commit",
		"Team chat message received",
	}
	fmt.Println(blue(activities[rand.Intn(len(activities))]))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}