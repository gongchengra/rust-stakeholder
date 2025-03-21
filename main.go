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
    Fullstack
    DataScience
    DevOps
    Blockchain
    MachineLearning
    SystemsProgramming
    GameDevelopment
    Security
)

// 删除重复的 DevelopmentType 常量声明
// const (
//     Backend DevelopmentType = iota
//     Frontend
//     FullStack
//     DevOps
//     Mobile
// )

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

// 添加术语生成器函数
func generateCodeJargon(devType DevelopmentType, level JargonLevel) string {
    basicTerms := map[DevelopmentType][]string{
        Backend: {
            "Optimized query execution paths for improved database throughput",
            "Reduced API latency via connection pooling and request batching",
            "Implemented stateless authentication with JWT token rotation",
            "Applied circuit breaker pattern to prevent cascading failures",
            "Utilized CQRS pattern for complex domain operations",
        },
        Frontend: {
            "Implemented virtual DOM diffing for optimal rendering performance",
            "Applied tree-shaking and code-splitting for bundle optimization",
            "Utilized CSS containment for layout performance improvement",
            "Implemented intersection observer for lazy-loading optimization",
            "Reduced reflow calculations with CSS will-change property",
        },
        // ... 其他开发类型的基础术语 ...
    }

    advancedTerms := map[DevelopmentType][]string{
        Backend: {
            "Implemented polyglot persistence with domain-specific data storage optimization",
            "Applied event-driven architecture with CQRS and event sourcing for eventual consistency",
            "Utilized domain-driven hexagonal architecture for maintainable business logic isolation",
            "Implemented reactive non-blocking I/O with backpressure handling for system resilience",
            "Applied saga pattern for distributed transaction management with compensating actions",
        },
        Frontend: {
            "Implemented compile-time static analysis for type-safe component composition",
            "Applied atomic CSS methodology with tree-shakable style injection",
            "Utilized custom rendering reconciliation with incremental DOM diffing",
            "Implemented time-sliced rendering with priority-based task scheduling",
            "Applied declarative animation system with hardware acceleration optimization",
        },
        // ... 其他开发类型的高级术语 ...
    }

    extremeTerms := []string{
        "Implemented isomorphic polymorphic runtime with transpiled metaprogramming for cross-paradigm interoperability",
        "Utilized quantum-resistant cryptographic primitives with homomorphic computation capabilities",
        "Applied non-euclidean topology optimization for multi-dimensional data representation",
        "Implemented stochastic gradient Langevin dynamics with cyclical annealing for robust convergence",
        "Utilized differentiable neural computers with external memory addressing for complex reasoning tasks",
    }

    switch level {
    case Low, Medium:
        if terms, ok := basicTerms[devType]; ok && len(terms) > 0 {
            return terms[rand.Intn(len(terms))]
        }
    case High:
        if terms, ok := advancedTerms[devType]; ok && len(terms) > 0 {
            return terms[rand.Intn(len(terms))]
        }
    case Expert:
        if rand.Float32() < 0.7 {
            return extremeTerms[rand.Intn(len(extremeTerms))]
        } else if terms, ok := advancedTerms[devType]; ok && len(terms) > 0 {
            return terms[rand.Intn(len(terms))]
        }
    }
    return "Optimizing system performance and resource utilization"
}

func generatePerformanceJargon(devType DevelopmentType, level JargonLevel) string {
    basicTerms := map[DevelopmentType][]string{
        Backend: {
            "Optimized request handling with connection pooling",
            "Implemented caching layer for frequently accessed data",
            "Applied query optimization for improved database performance",
            "Utilized async I/O for non-blocking request processing",
            "Implemented rate limiting to prevent resource contention",
        },
        Frontend: {
            "Optimized rendering pipeline with virtual DOM diffing",
            "Implemented code splitting for reduced initial load time",
            "Applied tree-shaking for reduced bundle size",
            "Utilized resource prioritization for critical path rendering",
            "Implemented request batching for reduced network overhead",
        },
    }

    // ... 其他性能相关术语 ...
    return getRandomTerm(devType, level, basicTerms)
}

func generateDataJargon(devType DevelopmentType, level JargonLevel) string {
    basicTerms := map[DevelopmentType][]string{
        DataScience: {
            "Applied feature normalization for improved model convergence",
            "Implemented data augmentation for enhanced training set diversity",
            "Utilized cross-validation for robust model evaluation",
            "Applied dimensionality reduction for feature space optimization",
            "Implemented ensemble methods for improved prediction accuracy",
        },
        MachineLearning: {
            "Applied feature normalization for improved model convergence",
            "Implemented data augmentation for enhanced training set diversity",
            "Utilized cross-validation for robust model evaluation",
            "Applied dimensionality reduction for feature space optimization",
            "Implemented ensemble methods for improved prediction accuracy",
        },
    }

    // ... 其他数据相关术语 ...
    return getRandomTerm(devType, level, basicTerms)
}

func generateNetworkJargon(devType DevelopmentType, level JargonLevel) string {
    basicTerms := map[DevelopmentType][]string{
        Backend: {
            "Optimized request batching for reduced network overhead",
            "Implemented connection pooling for improved throughput",
            "Applied response compression for bandwidth optimization",
            "Utilized HTTP/2 multiplexing for parallel requests",
            "Implemented retry strategies with exponential backoff",
        },
    }

    // ... 其他网络相关术语 ...
    return getRandomTerm(devType, level, basicTerms)
}

// 辅助函数，用于随机选择术语
func getRandomTerm(devType DevelopmentType, level JargonLevel, terms map[DevelopmentType][]string) string {
    if termList, ok := terms[devType]; ok && len(termList) > 0 {
        return termList[rand.Intn(len(termList))]
    }
    return "Optimizing system performance"
}

func generateJargon(devType DevelopmentType, level JargonLevel) string {
    basicTerms := map[DevelopmentType][]string{
        Backend: {
            "Optimized query execution paths for improved database throughput",
            "Reduced API latency via connection pooling and request batching",
            "Implemented stateless authentication with JWT token rotation",
            "Applied circuit breaker pattern to prevent cascading failures",
            "Utilized CQRS pattern for complex domain operations",
        },
        DataScience: {
            "Applied regularization techniques to prevent overfitting",
            "Implemented feature engineering pipeline with dimensionality reduction",
            "Utilized distributed computing for parallel data processing",
            "Optimized data transformations with vectorized operations",
            "Applied statistical significance testing to validate results",
        },
        Blockchain: {
            "Optimized transaction validation through merkle tree verification",
            "Implemented sharding for improved blockchain throughput",
            "Applied zero-knowledge proofs for privacy-preserving transactions",
            "Utilized state channels for off-chain scaling optimization",
            "Implemented consensus algorithm with Byzantine fault tolerance",
        },
        MachineLearning: {
            "Applied gradient boosting for improved model performance",
            "Implemented feature importance analysis for model interpretability",
            "Utilized transfer learning to optimize training efficiency",
            "Applied hyperparameter tuning with Bayesian optimization",
            "Implemented ensemble methods for model robustness",
        },
        SystemsProgramming: {
            "Optimized cache locality with data-oriented design patterns",
            "Implemented zero-copy memory management for I/O operations",
            "Applied lock-free algorithms for concurrent data structures",
            "Utilized SIMD instructions for vectorized processing",
            "Implemented memory pooling for reduced allocation overhead",
        },
        GameDevelopment: {
            "Optimized spatial partitioning for collision detection performance",
            "Implemented entity component system for flexible game architecture",
            "Applied level of detail techniques for rendering optimization",
            "Utilized GPU instancing for rendering large object counts",
            "Implemented deterministic physics for consistent simulation",
        },
        Security: {
            "Applied principle of least privilege across security boundaries",
            "Implemented defense-in-depth strategies for layered security",
            "Utilized cryptographic primitives for secure data exchange",
            "Applied security by design with threat modeling methodology",
            "Implemented zero-trust architecture for access control",
        },
    }

    advancedTerms := map[DevelopmentType][]string{
        Backend: {
            "Implemented polyglot persistence with domain-specific data storage optimization",
            "Applied event-driven architecture with CQRS and event sourcing for eventual consistency",
            "Utilized domain-driven hexagonal architecture for maintainable business logic isolation",
            "Implemented reactive non-blocking I/O with backpressure handling for system resilience",
            "Applied saga pattern for distributed transaction management with compensating actions",
        },
        MachineLearning: {
            "Implemented neural architecture search with reinforcement learning",
            "Applied differentiable programming for end-to-end trainable pipelines",
            "Utilized federated learning with secure aggregation protocols",
            "Implemented attention mechanisms with sparse transformers",
            "Applied meta-learning for few-shot adaptation capabilities",
        },
        Security: {
            "Implemented homomorphic encryption for secure multi-party computation",
            "Applied formal verification for cryptographic protocol security",
            "Utilized post-quantum cryptographic primitives for forward security",
            "Implemented secure multi-party computation with secret sharing",
            "Applied hardware-backed trusted execution environments for secure enclaves",
        },
    }

    extremeTerms := []string{
        "Implemented isomorphic polymorphic runtime with transpiled metaprogramming for cross-paradigm interoperability",
        "Utilized quantum-resistant cryptographic primitives with homomorphic computation capabilities",
        "Applied non-euclidean topology optimization for multi-dimensional data representation",
        "Implemented stochastic gradient Langevin dynamics with cyclical annealing for robust convergence",
        "Utilized differentiable neural computers with external memory addressing for complex reasoning tasks",
    }

    switch level {
    case Low, Medium:
        if terms, ok := basicTerms[devType]; ok && len(terms) > 0 {
            return terms[rand.Intn(len(terms))]
        }
    case High:
        if terms, ok := advancedTerms[devType]; ok && len(terms) > 0 {
            return terms[rand.Intn(len(terms))]
        }
    case Expert:
        if rand.Float32() < 0.7 {
            return extremeTerms[rand.Intn(len(extremeTerms))]
        } else if terms, ok := advancedTerms[devType]; ok && len(terms) > 0 {
            return terms[rand.Intn(len(terms))]
        }
    }

    return "Optimizing system performance and resource utilization"
}

func runDataProcessing(config *SessionConfig) {
    fmt.Println(blue("📊 Processing Data Streams"))
    
    dataPoints := rand.Intn(1000) + 500
    bar := progressbar.NewOptions(dataPoints,
        progressbar.OptionSetDescription("Processing data..."),
        progressbar.OptionShowCount(),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "▰",
            SaucerPadding: "▱",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    for i := 0; i < dataPoints; i++ {
        bar.Add(1)
        if i%50 == 0 {
            fmt.Printf("  🔄 %s\n", generateDataJargon(config.devType, config.jargonLevel))
        }
        time.Sleep(time.Duration(rand.Intn(50)+20) * time.Millisecond)
    }

    fmt.Printf("\n✅ Processed %d data points\n", dataPoints)
    fmt.Printf("💡 Insight: %s\n", generateDataJargon(config.devType, config.jargonLevel))
}

func runNetworkActivity(config *SessionConfig) {
    fmt.Println(yellow("🌐 Monitoring Network Activity"))
    
    packets := rand.Intn(200) + 100
    bar := progressbar.NewOptions(packets,
        progressbar.OptionSetDescription("Analyzing network..."),
        progressbar.OptionShowCount(),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "▰",
            SaucerPadding: "▱",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    for i := 0; i < packets; i++ {
        bar.Add(1)
        if i%20 == 0 {
            fmt.Printf("  📡 %s\n", generateNetworkJargon(config.devType, config.jargonLevel))
        }
        time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
    }

    fmt.Printf("\n📊 Network Analysis Complete\n")
    fmt.Printf("💡 Optimization: %s\n", generateNetworkJargon(config.devType, config.jargonLevel))
}

func displayRandomAlert(config *SessionConfig) {
    alerts := []string{
        "⚠️ High memory usage detected in worker process",
        "🔄 Auto-scaling triggered due to increased load",
        "📈 Performance threshold exceeded in API endpoint",
        "🔍 Unusual pattern detected in request flow",
        "⚡ Cache hit ratio below optimal threshold",
    }
    fmt.Printf("\n%s\n", alerts[rand.Intn(len(alerts))])
}

func displayTeamActivity(config *SessionConfig) {
    activities := []string{
        "👩‍💻 Team member pushing code updates",
        "👨‍💻 Code review in progress",
        "🤝 Merge request approved",
        "📝 Documentation update submitted",
        "🔧 Configuration changes deployed",
    }
    fmt.Printf("\n%s\n", activities[rand.Intn(len(activities))])
}

func getCodeAnalysisTitle(devType DevelopmentType, framework string) string {
    frameworkStr := ""
    if framework != "" {
        frameworkStr = fmt.Sprintf(" (%s)", framework)
    }
    return fmt.Sprintf("🔍 Running Code Analysis%s", frameworkStr)
}

func generateFileName(devType DevelopmentType) string {
    extensions := map[DevelopmentType][]string{
        Backend:     {".go", ".rs", ".java", ".py"},
        Frontend:    {".js", ".ts", ".vue", ".jsx"},
        Fullstack:   {".ts", ".go", ".py", ".jsx"},
    }

    prefixes := []string{"service", "controller", "model", "util", "helper"}
    names := []string{"user", "auth", "data", "config", "api"}

    ext := ".go"
    if exts, ok := extensions[devType]; ok {
        ext = exts[rand.Intn(len(exts))]
    }

    prefix := prefixes[rand.Intn(len(prefixes))]
    name := names[rand.Intn(len(names))]

    return fmt.Sprintf("%s_%s%s", prefix, name, ext)
}

func generateCodeIssue(devType DevelopmentType) string {
    issues := []string{
        "Potential memory leak",
        "Uncaught exception",
        "Resource not released",
        "Inefficient algorithm",
        "Security vulnerability",
    }
    return issues[rand.Intn(len(issues))]
}

func generateComplexityMetric() string {
    metrics := []string{
        "Cyclomatic complexity: 15",
        "Cognitive complexity: 8",
        "Maintainability index: 75",
        "Code coverage: 85%",
        "Technical debt ratio: 5%",
    }
    return metrics[rand.Intn(len(metrics))]
}

func getPerformanceTitle(devType DevelopmentType) string {
    return "⚡ Performance Analysis"
}

func generateBasePerformance(devType DevelopmentType) float64 {
    return 50.0 + rand.Float64()*30.0
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

func generatePerformanceMetric(devType DevelopmentType) string {
    metrics := map[DevelopmentType][]string{
        Backend: {"Response time", "Throughput", "Error rate", "Queue length", "Cache hit ratio"},
        Frontend: {"Time to interactive", "First paint", "Bundle size", "Memory usage", "Frame rate"},
        Fullstack: {"End-to-end latency", "API response time", "Database queries", "Cache efficiency", "Network latency"},
    }
    
    if metricList, ok := metrics[devType]; ok {
        return metricList[rand.Intn(len(metricList))]
    }
    return "Performance metric"
}

func generateMetricUnit(devType DevelopmentType) string {
    units := map[DevelopmentType][]string{
        Backend: {"req/s", "ms", "μs", "MB/s", "connections"},
        Frontend: {"ms", "KB", "fps", "MB", "req/s"},
        Fullstack: {"ms", "req/s", "MB/s", "ops/s", "connections"},
    }
    
    if unitList, ok := units[devType]; ok {
        return unitList[rand.Intn(len(unitList))]
    }
    return "units"
}

func generateOptimizationRecommendation(devType DevelopmentType) string {
    recommendations := map[DevelopmentType][]string{
        Backend: {
            "Consider implementing request caching to reduce database load",
            "Optimize database query patterns for improved throughput",
            "Implement connection pooling for better resource utilization",
            "Add request compression for reduced network overhead",
            "Consider implementing circuit breakers for external services",
        },
        Frontend: {
            "Implement lazy loading for improved initial load time",
            "Consider code splitting for optimized bundle size",
            "Add service worker for offline capabilities",
            "Optimize critical rendering path",
            "Implement resource prioritization",
        },
    }
    
    if recList, ok := recommendations[devType]; ok {
        return recList[rand.Intn(len(recList))]
    }
    return "Consider optimizing system performance"
}

func formatResourceValue(value, warningThreshold, criticalThreshold int) string {
    if value >= criticalThreshold {
        return red(fmt.Sprintf("%d%%", value))
    }
    if value >= warningThreshold {
        return yellow(fmt.Sprintf("%d%%", value))
    }
    return green(fmt.Sprintf("%d%%", value))
}

func generateSystemEvent() string {
    events := []string{
        "Container auto-scaling event triggered",
        "Cache invalidation completed",
        "Background job processing completed",
        "System health check passed",
        "Metrics collection cycle completed",
        "Log rotation executed",
        "Configuration refresh completed",
        "Resource cleanup task executed",
    }
    return events[rand.Intn(len(events))]
}

func generateSystemRecommendation() string {
    recommendations := []string{
        "Consider increasing cache size for improved performance",
        "Optimize background job scheduling for better resource utilization",
        "Review logging levels to reduce I/O overhead",
        "Consider implementing request rate limiting",
        "Optimize database connection pool settings",
        "Review auto-scaling thresholds for better resource efficiency",
    }
    return recommendations[rand.Intn(len(recommendations))]
}
