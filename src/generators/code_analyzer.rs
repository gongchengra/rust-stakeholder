use crate::DevelopmentType;
use rand::{prelude::*, rng};

pub fn generate_filename(dev_type: DevelopmentType) -> String {
    let extensions = match dev_type {
        DevelopmentType::Backend => [
            "rs", "go", "java", "py", "js", "ts", "rb", "php", "cs", "scala",
        ],
        DevelopmentType::Frontend => [
            "js", "ts", "jsx", "tsx", "vue", "scss", "css", "html", "svelte", "elm",
        ],
        DevelopmentType::Fullstack => [
            "js", "ts", "rs", "go", "py", "jsx", "tsx", "vue", "rb", "php",
        ],
        DevelopmentType::DataScience => [
            "py", "ipynb", "R", "jl", "scala", "sql", "m", "stan", "cpp", "h",
        ],
        DevelopmentType::DevOps => [
            "yaml",
            "yml",
            "tf",
            "hcl",
            "sh",
            "Dockerfile",
            "json",
            "toml",
            "ini",
            "conf",
        ],
        DevelopmentType::Blockchain => [
            "sol", "rs", "go", "js", "ts", "wasm", "move", "cairo", "vy", "cpp",
        ],
        DevelopmentType::MachineLearning => [
            "py", "ipynb", "pth", "h5", "pb", "tflite", "onnx", "pt", "cpp", "cu",
        ],
        DevelopmentType::SystemsProgramming => {
            ["rs", "c", "cpp", "h", "hpp", "asm", "s", "go", "zig", "d"]
        }
        DevelopmentType::GameDevelopment => [
            "cpp", "h", "cs", "js", "ts", "glsl", "hlsl", "shader", "unity", "prefab",
        ],
        DevelopmentType::Security => [
            "rs", "go", "c", "cpp", "py", "java", "js", "ts", "rb", "php",
        ],
    };

    let components = match dev_type {
        DevelopmentType::Backend => [
            "Service",
            "Controller",
            "Repository",
            "DAO",
            "Manager",
            "Factory",
            "Provider",
            "Client",
            "Handler",
            "Middleware",
            "Interceptor",
            "Connector",
            "Processor",
            "Worker",
            "Queue",
            "Cache",
            "Store",
            "Adapter",
            "Wrapper",
            "Mapper",
        ],
        DevelopmentType::Frontend => [
            "Component",
            "Container",
            "Page",
            "View",
            "Screen",
            "Element",
            "Layout",
            "Widget",
            "Hook",
            "Context",
            "Provider",
            "Reducer",
            "Action",
            "State",
            "Form",
            "Modal",
            "Card",
            "Button",
            "Input",
            "Selector",
        ],
        DevelopmentType::Fullstack => [
            "Service",
            "Controller",
            "Component",
            "Container",
            "Connector",
            "Integration",
            "Provider",
            "Client",
            "Api",
            "Interface",
            "Bridge",
            "Adapter",
            "Manager",
            "Handler",
            "Processor",
            "Orchestrator",
            "Facade",
            "Proxy",
            "Wrapper",
            "Mapper",
        ],
        DevelopmentType::DataScience => [
            "Analysis",
            "Processor",
            "Transformer",
            "Pipeline",
            "Extractor",
            "Loader",
            "Model",
            "Predictor",
            "Classifier",
            "Regressor",
            "Clusterer",
            "Encoder",
            "Trainer",
            "Evaluator",
            "Feature",
            "Dataset",
            "Optimizer",
            "Validator",
            "Sampler",
            "Splitter",
        ],
        DevelopmentType::DevOps => [
            "Config",
            "Setup",
            "Deployment",
            "Pipeline",
            "Builder",
            "Runner",
            "Provisioner",
            "Monitor",
            "Logger",
            "Alerter",
            "Scanner",
            "Tester",
            "Backup",
            "Security",
            "Network",
            "Cluster",
            "Container",
            "Orchestrator",
            "Manager",
            "Scheduler",
        ],
        DevelopmentType::Blockchain => [
            "Contract",
            "Wallet",
            "Token",
            "Chain",
            "Block",
            "Transaction",
            "Validator",
            "Miner",
            "Node",
            "Consensus",
            "Ledger",
            "Network",
            "Pool",
            "Oracle",
            "Signer",
            "Verifier",
            "Bridge",
            "Protocol",
            "Exchange",
            "Market",
        ],
        DevelopmentType::MachineLearning => [
            "Model",
            "Trainer",
            "Predictor",
            "Pipeline",
            "Transformer",
            "Encoder",
            "Embedder",
            "Classifier",
            "Regressor",
            "Optimizer",
            "Layer",
            "Network",
            "DataLoader",
            "Preprocessor",
            "Evaluator",
            "Validator",
            "Callback",
            "Metric",
            "Loss",
            "Sampler",
        ],
        DevelopmentType::SystemsProgramming => [
            "Allocator",
            "Memory",
            "Thread",
            "Process",
            "Scheduler",
            "Dispatcher",
            "Device",
            "Driver",
            "Buffer",
            "Stream",
            "Channel",
            "IO",
            "FS",
            "Network",
            "Synchronizer",
            "Lock",
            "Atomic",
            "Signal",
            "Interrupt",
            "Handler",
        ],
        DevelopmentType::GameDevelopment => [
            "Engine",
            "Renderer",
            "Physics",
            "Audio",
            "Input",
            "Entity",
            "Component",
            "System",
            "Scene",
            "Level",
            "Player",
            "Character",
            "Animation",
            "Sprite",
            "Camera",
            "Light",
            "Particle",
            "Collision",
            "AI",
            "Pathfinding",
        ],
        DevelopmentType::Security => [
            "Auth",
            "Identity",
            "Credential",
            "Token",
            "Certificate",
            "Encryption",
            "Hasher",
            "Signer",
            "Verifier",
            "Scanner",
            "Detector",
            "Analyzer",
            "Filter",
            "Firewall",
            "Proxy",
            "Inspector",
            "Monitor",
            "Logger",
            "Policy",
            "Permission",
        ],
    };

    let domain_prefixes = match dev_type {
        DevelopmentType::Backend => [
            "User",
            "Account",
            "Order",
            "Payment",
            "Product",
            "Inventory",
            "Customer",
            "Shipment",
            "Transaction",
            "Notification",
            "Message",
            "Event",
            "Task",
            "Job",
            "Schedule",
            "Catalog",
            "Cart",
            "Recommendation",
            "Analytics",
            "Report",
        ],
        DevelopmentType::Frontend => [
            "User",
            "Auth",
            "Product",
            "Cart",
            "Checkout",
            "Profile",
            "Dashboard",
            "Settings",
            "Notification",
            "Message",
            "Search",
            "List",
            "Detail",
            "Home",
            "Landing",
            "Admin",
            "Modal",
            "Navigation",
            "Theme",
            "Responsive",
        ],
        _ => [
            "Core", "Main", "Base", "Shared", "Util", "Helper", "Abstract", "Default", "Custom",
            "Advanced", "Simple", "Complex", "Dynamic", "Static", "Global", "Local", "Internal",
            "External", "Public", "Private",
        ],
    };

    let prefix = if rng().random_ratio(2, 3) {
        domain_prefixes.choose(&mut rng()).unwrap()
    } else {
        components.choose(&mut rng()).unwrap()
    };

    let component = components.choose(&mut rng()).unwrap();
    let extension = extensions.choose(&mut rng()).unwrap();

    // Only use prefix if it's different from component
    if prefix == component {
        format!("{}.{}", component, extension)
    } else {
        format!("{}{}.{}", prefix, component, extension)
    }
}

pub fn generate_code_issue(dev_type: DevelopmentType) -> String {
    let common_issues = [
        "Unused variable",
        "Unreachable code",
        "Redundant calculation",
        "Missing error handling",
        "Inefficient algorithm",
        "Potential null reference",
        "Code duplication",
        "Overly complex method",
        "Deprecated API usage",
        "Resource leak",
    ];

    let specific_issues = match dev_type {
        DevelopmentType::Backend => [
            "Unoptimized database query",
            "Missing transaction boundary",
            "Potential SQL injection",
            "Inefficient connection management",
            "Improper error propagation",
            "Race condition in concurrent request handling",
            "Inadequate request validation",
            "Excessive logging",
            "Missing authentication check",
            "Insufficient rate limiting",
        ],
        DevelopmentType::Frontend => [
            "Unnecessary component re-rendering",
            "Unhandled promise rejection",
            "Excessive DOM manipulation",
            "Memory leak in event listener",
            "Non-accessible UI element",
            "Inconsistent styling approach",
            "Unoptimized asset loading",
            "Browser compatibility issue",
            "Inefficient state management",
            "Poor mobile responsiveness",
        ],
        DevelopmentType::Fullstack => [
            "Inconsistent data validation",
            "Redundant data transformation",
            "Inefficient client-server communication",
            "Mismatched data types",
            "Inconsistent error handling",
            "Overly coupled client-server logic",
            "Duplicated business logic",
            "Inconsistent state management",
            "Security vulnerability in API integration",
            "Race condition in state synchronization",
        ],
        DevelopmentType::DataScience => [
            "Potential data leakage",
            "Inadequate data normalization",
            "Inefficient data transformation",
            "Missing null value handling",
            "Improper train-test split",
            "Unoptimized feature selection",
            "Insufficient data validation",
            "Model overfitting risk",
            "Numerical instability in calculation",
            "Memory inefficient data processing",
        ],
        DevelopmentType::DevOps => [
            "Insecure configuration default",
            "Missing resource constraint",
            "Inadequate error recovery mechanism",
            "Inefficient resource allocation",
            "Hardcoded credential",
            "Insufficient monitoring setup",
            "Non-idempotent operation",
            "Missing backup strategy",
            "Inadequate security policy",
            "Inefficient deployment process",
        ],
        DevelopmentType::Blockchain => [
            "Gas inefficient operation",
            "Potential reentrancy vulnerability",
            "Improper access control",
            "Integer overflow/underflow risk",
            "Unchecked external call result",
            "Inadequate transaction validation",
            "Front-running vulnerability",
            "Improper randomness source",
            "Inefficient storage pattern",
            "Missing event emission",
        ],
        DevelopmentType::MachineLearning => [
            "Potential data leakage",
            "Inefficient model architecture",
            "Improper learning rate scheduling",
            "Unhandled gradient explosion risk",
            "Inefficient batch processing",
            "Inadequate model evaluation metric",
            "Memory inefficient tensor operation",
            "Missing early stopping criteria",
            "Unoptimized hyperparameter",
            "Inefficient feature engineering",
        ],
        DevelopmentType::SystemsProgramming => [
            "Potential memory leak",
            "Uninitialized memory access",
            "Thread synchronization issue",
            "Inefficient memory allocation",
            "Resource cleanup failure",
            "Buffer overflow risk",
            "Race condition in concurrent access",
            "Inefficient cache usage pattern",
            "Blocking I/O in critical path",
            "Undefined behavior risk",
        ],
        DevelopmentType::GameDevelopment => [
            "Inefficient rendering call",
            "Physics calculation in rendering thread",
            "Unoptimized asset loading",
            "Missing frame rate cap",
            "Memory fragmentation risk",
            "Inefficient collision detection",
            "Unoptimized shader complexity",
            "Animation state machine complexity",
            "Inefficient particle system update",
            "Missing object pooling",
        ],
        DevelopmentType::Security => [
            "Potential privilege escalation",
            "Insecure cryptographic algorithm",
            "Missing input validation",
            "Hardcoded credential",
            "Insufficient authentication check",
            "Security misconfiguration",
            "Inadequate error handling exposing details",
            "Missing rate limiting",
            "Insecure direct object reference",
            "Improper certificate validation",
        ],
    };

    if rng().random_ratio(1, 3) {
        common_issues.choose(&mut rng()).unwrap().to_string()
    } else {
        specific_issues.choose(&mut rng()).unwrap().to_string()
    }
}

pub fn generate_complexity_metric() -> String {
    let complexity_metrics = [
        "Cyclomatic complexity: 5 (good)",
        "Cyclomatic complexity: 8 (acceptable)",
        "Cyclomatic complexity: 12 (moderate)",
        "Cyclomatic complexity: 18 (high)",
        "Cyclomatic complexity: 25 (very high)",
        "Cognitive complexity: 4 (good)",
        "Cognitive complexity: 7 (acceptable)",
        "Cognitive complexity: 15 (moderate)",
        "Cognitive complexity: 22 (high)",
        "Cognitive complexity: 30 (very high)",
        "Maintainability index: 85 (highly maintainable)",
        "Maintainability index: 75 (maintainable)",
        "Maintainability index: 65 (moderately maintainable)",
        "Maintainability index: 55 (difficult to maintain)",
        "Maintainability index: 45 (very difficult to maintain)",
        "Lines of code: 25 (compact)",
        "Lines of code: 75 (moderate)",
        "Lines of code: 150 (large)",
        "Lines of code: 300 (very large)",
        "Lines of code: 500+ (extremely large)",
        "Nesting depth: 2 (good)",
        "Nesting depth: 3 (acceptable)",
        "Nesting depth: 4 (moderate)",
        "Nesting depth: 5 (high)",
        "Nesting depth: 6+ (very high)",
    ];

    complexity_metrics.choose(&mut rng()).unwrap().to_string()
}
