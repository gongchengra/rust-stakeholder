package main

import (
    "fmt"
    "math"
    "math/rand"
    "os"
    "os/signal"
    "sort"
    "sync/atomic"
    "syscall"
    "time"

    "github.com/fatih/color"
    "github.com/schollz/progressbar/v3"
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

// 首先添加必要的依赖
func runCodeAnalysis(config *SessionConfig) {
    filesToAnalyze := rand.Intn(20) + 5
    totalLines := rand.Intn(9000) + 1000

    title := getCodeAnalysisTitle(config.devType, config.framework)
    fmt.Println(blue(title))

    // 创建进度条
    bar := progressbar.NewOptions(filesToAnalyze,
        progressbar.OptionSetDescription("Analyzing files..."),
        progressbar.OptionShowCount(),
        progressbar.OptionShowIts(),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "▰",
            SaucerPadding: "▱",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    for i := 0; i < filesToAnalyze; i++ {
        bar.Add(1)
        if rand.Float32() < 0.3 {
            fileName := generateFileName(config.devType)
            issueType := generateCodeIssue(config.devType)
            complexity := generateComplexityMetric()

            if rand.Float32() < 0.25 {
                fmt.Printf("  ⚠️ %s - %s: %s\n", fileName, issueType, complexity)
            } else {
                fmt.Printf("  ✓ %s - %s\n", fileName, complexity)
            }
        }
        time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
    }

    // 分析总结
    fmt.Printf("\n📊 Analysis Complete: %d files, %d lines of code\n", filesToAnalyze, totalLines)
    fmt.Printf("  - Issues found: %d\n", rand.Intn(5))
    fmt.Printf("  - Code quality score: %d%%\n", rand.Intn(14)+85)
    fmt.Printf("  - Technical debt: %d%%\n", rand.Intn(14)+1)
}

// 扩充性能指标功能
func runPerformanceMetrics(config *SessionConfig) {
    title := getPerformanceTitle(config.devType)
    fmt.Println(yellow(title))

    iterations := rand.Intn(150) + 50
    bar := progressbar.NewOptions(iterations,
        progressbar.OptionSetDescription("Collecting metrics..."),
        progressbar.OptionShowCount(),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "▰",
            SaucerPadding: "▱",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    var performanceData []float64

    for i := 0; i < iterations; i++ {
        bar.Add(1)
        basePerf := generateBasePerformance(config.devType)
        jitter := (rand.Float64() * 10) - 5
        perfValue := math.Max(basePerf+jitter, 1.0)
        performanceData = append(performanceData, perfValue)

        if i%10 == 0 && rand.Float32() < 0.3 {
            metricName := generatePerformanceMetric(config.devType)
            metricValue := rand.Intn(989) + 10
            metricUnit := generateMetricUnit(config.devType)
            fmt.Printf("  📊 %s: %d %s\n", metricName, metricValue, metricUnit)
        }

        time.Sleep(time.Duration(rand.Intn(50)+50) * time.Millisecond)
    }

    // 计算并显示指标
    sort.Float64s(performanceData)
    avg := calculateAverage(performanceData)
    median := performanceData[len(performanceData)/2]
    p95 := performanceData[int(float64(len(performanceData))*0.95)]
    p99 := performanceData[int(float64(len(performanceData))*0.99)]

    fmt.Println("\n📈 Performance Results:")
    fmt.Printf("  - Average: %.2f ms\n", avg)
    fmt.Printf("  - Median: %.2f ms\n", median)
    fmt.Printf("  - P95: %.2f ms\n", p95)
    fmt.Printf("  - P99: %.2f ms\n", p99)

    // 添加优化建议
    fmt.Printf("💡 Recommendation: %s\n", generateOptimizationRecommendation(config.devType))
}

// 扩充系统监控功能
func runSystemMonitoring(config *SessionConfig) {
    fmt.Println(green("🖥️ System Resource Monitoring"))

    duration := rand.Intn(10) + 5
    bar := progressbar.NewOptions(duration,
        progressbar.OptionSetDescription("Monitoring..."),
        progressbar.OptionShowCount(),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "▰",
            SaucerPadding: "▱",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    cpuBase := rand.Intn(50) + 10
    memoryBase := rand.Intn(40) + 30
    networkBase := rand.Intn(19) + 1
    diskBase := rand.Intn(35) + 5

    for i := 0; i < duration; i++ {
        bar.Add(1)

        cpu := cpuBase + rand.Intn(15) - 5
        memory := memoryBase + rand.Intn(8) - 3
        network := networkBase + rand.Intn(4) - 1
        disk := diskBase + rand.Intn(6) - 2
        processes := rand.Intn(120) + 80

        cpuStr := formatResourceValue(cpu, 80, 60)
        memStr := formatResourceValue(memory, 85, 70)

        fmt.Printf("  CPU: %s  |  RAM: %s  |  Network: %d MB/s  |  Disk I/O: %d MB/s  |  Processes: %d\n",
            cpuStr, memStr, network, disk, processes)

        if i%3 == 0 && rand.Float32() < 0.3 {
            fmt.Printf("  🔄 %s\n", generateSystemEvent())
        }

        time.Sleep(time.Duration(rand.Intn(300)+200) * time.Millisecond)
    }

    // 显示总结
    fmt.Println("\n📊 Resource Utilization Summary:")
    fmt.Printf("  - Peak CPU: %d%%\n", cpuBase+rand.Intn(10)+5)
    fmt.Printf("  - Peak Memory: %d%%\n", memoryBase+rand.Intn(10)+5)
    fmt.Printf("  - Network Throughput: %d MB/s\n", networkBase+rand.Intn(5)+5)
    fmt.Printf("  - Disk Throughput: %d MB/s\n", diskBase+rand.Intn(6)+2)
    fmt.Printf("  - %s\n", generateSystemRecommendation())
}

// 添加新的生成器函数
func generateMetricUnit(devType DevelopmentType) string {
    units := map[DevelopmentType][]string{
        Backend: {"req/s", "ms", "μs", "MB/s", "connections", "sessions", "%", "threads", "MB", "ops/s"},
        Frontend: {"ms", "fps", "KB", "MB", "elements", "nodes", "req/s", "s", "μs", "%"},
        // ... 可以添加其他类型的单位 ...
    }
    if metricUnits, ok := units[devType]; ok {
        return metricUnits[rand.Intn(len(metricUnits))]
    }
    defaultUnits := []string{"ms", "s", "MB/s", "GB/s", "ops/s", "%", "MB", "KB", "count", "ratio"}
    return defaultUnits[rand.Intn(len(defaultUnits))]
}

func generateOptimizationRecommendation(devType DevelopmentType) string {
    recommendations := map[DevelopmentType][]string{
        Backend: {
            "Consider implementing request batching for high-volume endpoints",
            "Database query optimization could improve response times by 15-20%",
            "Adding a distributed cache layer would reduce database load",
            // ... 添加更多建议 ...
        },
        Frontend: {
            "Implement code splitting to reduce initial bundle size",
            "Consider lazy loading for off-screen components",
            "Optimize critical rendering path for faster first paint",
            // ... 添加更多建议 ...
        },
        // ... 可以添加其他类型的建议 ...
    }
    if typeRecommendations, ok := recommendations[devType]; ok {
        return typeRecommendations[rand.Intn(len(typeRecommendations))]
    }
    return "Consider optimizing resource utilization"
}

func generatePerformanceMetric(devType DevelopmentType) string {
    metrics := map[DevelopmentType][]string{
        Backend: {
            "API Response Time",
            "Database Query Latency",
            "Request Throughput",
            "Cache Hit Ratio",
            "Connection Pool Utilization",
        },
        Frontend: {
            "Render Time",
            "First Contentful Paint",
            "Time to Interactive",
            "Bundle Size",
            "DOM Node Count",
        },
    }
    if typeMetrics, ok := metrics[devType]; ok {
        return typeMetrics[rand.Intn(len(typeMetrics))]
    }
    defaultMetrics := []string{"Processing Time", "Resource Usage", "Operation Latency"}
    return defaultMetrics[rand.Intn(len(defaultMetrics))]
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

func getCodeAnalysisTitle(devType DevelopmentType, framework string) string {
    frameworkSpecific := ""
    if framework != "" {
        frameworkSpecific = fmt.Sprintf(" (%s specific)", framework)
    }

    titles := map[DevelopmentType]string{
        Backend:  fmt.Sprintf("🔍 Running Code Analysis on API Components%s", frameworkSpecific),
        Frontend: fmt.Sprintf("🔍 Analyzing UI Components%s", frameworkSpecific),
        FullStack: "🔍 Analyzing Full-Stack Integration Points",
        DevOps:    "🔍 Analyzing Infrastructure Configuration",
        Mobile:    "🔍 Analyzing Mobile App Components",
    }

    if title, ok := titles[devType]; ok {
        return title
    }
    return "🔍 Running Code Analysis"
}

func generateFileName(devType DevelopmentType) string {
    backendFiles := []string{"api.go", "service.go", "repository.go", "middleware.go", "handler.go"}
    frontendFiles := []string{"app.js", "component.tsx", "styles.css", "utils.js", "router.js"}

    switch devType {
    case Backend:
        return backendFiles[rand.Intn(len(backendFiles))]
    case Frontend:
        return frontendFiles[rand.Intn(len(frontendFiles))]
    default:
        return fmt.Sprintf("file_%d.txt", rand.Intn(100))
    }
}

func generateCodeIssue(devType DevelopmentType) string {
    issues := []string{
        "Potential memory leak",
        "Unused variable",
        "Complex function",
        "Missing error handling",
        "Code duplication",
    }
    return issues[rand.Intn(len(issues))]
}

func generateComplexityMetric() string {
    metrics := []string{
        "Cyclomatic complexity: 15",
        "Cognitive complexity: 8",
        "Maintainability index: 75",
        "Code coverage: 85%",
    }
    return metrics[rand.Intn(len(metrics))]
}

func getPerformanceTitle(devType DevelopmentType) string {
    titles := map[DevelopmentType]string{
        Backend:   "⚡ Analyzing API Response Time",
        Frontend:  "⚡ Measuring UI Rendering Performance",
        FullStack: "⚡ Evaluating End-to-End Performance",
        DevOps:    "⚡ Evaluating Infrastructure Performance",
        Mobile:    "⚡ Analyzing Mobile App Performance",
    }

    if title, ok := titles[devType]; ok {
        return title
    }
    return "⚡ Analyzing Performance"
}

func generateBasePerformance(devType DevelopmentType) float64 {
    switch devType {
    case Backend:
        return float64(rand.Intn(60) + 20)
    case Frontend:
        return float64(rand.Intn(25) + 5)
    default:
        return float64(rand.Intn(90) + 10)
    }
}

func calculateAverage(data []float64) float64 {
    if len(data) == 0 {
        return 0
    }
    sum := 0.0
    for _, v := range data {
        sum += v
    }
    return sum / float64(len(data))
}

func formatResourceValue(value, highThreshold, mediumThreshold int) string {
    str := fmt.Sprintf("%d%%", value)
    if value > highThreshold {
        return red(str)
    } else if value > mediumThreshold {
        return yellow(str)
    }
    return green(str)
}

func generateSystemEvent() string {
    events := []string{
        "Service auto-scaling triggered",
        "Cache invalidation completed",
        "Background job completed",
        "Config reload successful",
        "Backup process initiated",
    }
    return events[rand.Intn(len(events))]
}

func generateSystemRecommendation() string {
    recommendations := []string{
        "Consider increasing cache size",
        "Optimize background job frequency",
        "Review auto-scaling thresholds",
        "Implement resource usage alerts",
        "Schedule routine maintenance",
    }
    return recommendations[rand.Intn(len(recommendations))]
}
