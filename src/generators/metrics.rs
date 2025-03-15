use crate::DevelopmentType;
use rand::{prelude::*, rng};

pub fn generate_metric_unit(dev_type: DevelopmentType) -> String {
    let units = match dev_type {
        DevelopmentType::DataScience | DevelopmentType::MachineLearning => [
            "MB/s",
            "GB/s",
            "records/s",
            "samples/s",
            "iterations/s",
            "ms/batch",
            "s/epoch",
            "%",
            "MB",
            "GB",
        ],
        DevelopmentType::Backend | DevelopmentType::Fullstack => [
            "req/s",
            "ms",
            "μs",
            "MB/s",
            "connections",
            "sessions",
            "%",
            "threads",
            "MB",
            "ops/s",
        ],
        DevelopmentType::Frontend => [
            "ms", "fps", "KB", "MB", "elements", "nodes", "req/s", "s", "μs", "%",
        ],
        _ => [
            "ms", "s", "MB/s", "GB/s", "ops/s", "%", "MB", "KB", "count", "ratio",
        ],
    };

    units.choose(&mut rng()).unwrap().to_string()
}

pub fn generate_optimization_recommendation(dev_type: DevelopmentType) -> String {
    let recommendations = match dev_type {
        DevelopmentType::Backend => [
            "Consider implementing request batching for high-volume endpoints",
            "Database query optimization could improve response times by 15-20%",
            "Adding a distributed cache layer would reduce database load",
            "Implement connection pooling to reduce connection overhead",
            "Consider async processing for non-critical operations",
            "Implement circuit breakers for external service dependencies",
            "Database index optimization could improve query performance",
            "Consider implementing a read replica for heavy read workloads",
            "API response compression could reduce bandwidth consumption",
            "Implement rate limiting to protect against traffic spikes",
        ],
        DevelopmentType::Frontend => [
            "Implement code splitting to reduce initial bundle size",
            "Consider lazy loading for off-screen components",
            "Optimize critical rendering path for faster first paint",
            "Use memoization for expensive component calculations",
            "Implement virtualization for long scrollable lists",
            "Consider using web workers for CPU-intensive tasks",
            "Optimize asset loading with preload/prefetch strategies",
            "Implement request batching for multiple API calls",
            "Reduce JavaScript execution time with debouncing/throttling",
            "Optimize animation performance with CSS GPU acceleration",
        ],
        DevelopmentType::Fullstack => [
            "Implement more efficient data serialization between client and server",
            "Consider GraphQL for more efficient data fetching",
            "Optimize state management to reduce unnecessary renders",
            "Implement server-side rendering for improved initial load time",
            "Consider BFF pattern for optimized client-specific endpoints",
            "Reduce client-server round trips with data denormalization",
            "Implement WebSocket for real-time updates instead of polling",
            "Consider implementing a service worker for offline capabilities",
            "Optimize API contract for reduced payload sizes",
            "Implement shared validation logic between client and server",
        ],
        DevelopmentType::DataScience => [
            "Optimize feature engineering pipeline for parallel processing",
            "Consider incremental processing for large datasets",
            "Implement vectorized operations for numerical computations",
            "Consider dimensionality reduction to improve model efficiency",
            "Optimize data loading with memory-mapped files",
            "Implement distributed processing for large-scale computations",
            "Consider feature selection to reduce model complexity",
            "Optimize hyperparameter search strategy",
            "Implement early stopping criteria for training efficiency",
            "Consider model quantization for inference optimization",
        ],
        DevelopmentType::DevOps => [
            "Implement horizontal scaling for improved throughput",
            "Consider containerization for consistent deployment",
            "Optimize CI/CD pipeline for faster build times",
            "Implement infrastructure as code for reproducible environments",
            "Consider implementing a service mesh for observability",
            "Optimize resource allocation based on usage patterns",
            "Implement automated scaling policies based on demand",
            "Consider implementing blue-green deployments for zero downtime",
            "Optimize container image size for faster deployments",
            "Implement distributed tracing for performance bottleneck identification",
        ],
        DevelopmentType::Blockchain => [
            "Optimize smart contract gas usage with storage pattern refinement",
            "Consider implementing a layer 2 solution for improved throughput",
            "Optimize transaction validation with batched signature verification",
            "Implement more efficient consensus algorithm for reduced latency",
            "Consider sharding for improved scalability",
            "Optimize state storage with pruning strategies",
            "Implement efficient merkle tree computation",
            "Consider optimistic execution for improved transaction throughput",
            "Optimize P2P network propagation with better peer selection",
            "Implement efficient cryptographic primitives for reduced overhead",
        ],
        DevelopmentType::MachineLearning => [
            "Implement model quantization for faster inference",
            "Consider knowledge distillation for smaller model footprint",
            "Optimize batch size for improved training throughput",
            "Implement mixed-precision training for better GPU utilization",
            "Consider implementing gradient accumulation for larger effective batch sizes",
            "Optimize data loading pipeline with prefetching",
            "Implement model pruning for reduced parameter count",
            "Consider feature selection for improved model efficiency",
            "Optimize distributed training communication patterns",
            "Implement efficient checkpoint strategies for reduced storage requirements",
        ],
        DevelopmentType::SystemsProgramming => [
            "Optimize memory access patterns for improved cache utilization",
            "Consider implementing custom memory allocators for specific workloads",
            "Implement lock-free data structures for concurrent access",
            "Optimize instruction pipelining with code layout restructuring",
            "Consider SIMD instructions for vectorized processing",
            "Implement efficient thread pooling for reduced creation overhead",
            "Optimize I/O operations with asynchronous processing",
            "Consider memory-mapped I/O for large file operations",
            "Implement efficient serialization for data interchange",
            "Consider zero-copy strategies for data processing pipelines",
        ],
        DevelopmentType::GameDevelopment => [
            "Implement object pooling for frequently created entities",
            "Consider frustum culling optimization for rendering performance",
            "Optimize draw call batching for reduced GPU overhead",
            "Implement level of detail (LOD) for distant objects",
            "Consider async loading for game assets",
            "Optimize physics simulation with spatial partitioning",
            "Implement efficient animation blending techniques",
            "Consider GPU instancing for similar objects",
            "Optimize shader complexity for better performance",
            "Implement efficient collision detection with broad-phase algorithms",
        ],
        DevelopmentType::Security => [
            "Implement cryptographic acceleration for improved performance",
            "Consider session caching for reduced authentication overhead",
            "Optimize security scanning with incremental analysis",
            "Implement efficient key management for reduced overhead",
            "Consider least-privilege optimization for security checks",
            "Optimize certificate validation with efficient revocation checking",
            "Implement efficient secure channel negotiation",
            "Consider security policy caching for improved evaluation performance",
            "Optimize encryption algorithm selection based on data sensitivity",
            "Implement efficient log analysis with streaming processing",
        ],
    };

    recommendations.choose(&mut rng()).unwrap().to_string()
}

pub fn generate_performance_metric(dev_type: DevelopmentType) -> String {
    let metrics = match dev_type {
        DevelopmentType::Backend => [
            "API Response Time",
            "Database Query Latency",
            "Request Throughput",
            "Cache Hit Ratio",
            "Connection Pool Utilization",
            "Thread Pool Saturation",
            "Queue Depth",
            "Active Sessions",
            "Error Rate",
            "GC Pause Time",
        ],
        DevelopmentType::Frontend => [
            "Render Time",
            "First Contentful Paint",
            "Time to Interactive",
            "Bundle Size",
            "DOM Node Count",
            "Frame Rate",
            "Memory Usage",
            "Network Request Count",
            "Asset Load Time",
            "Input Latency",
        ],
        DevelopmentType::Fullstack => [
            "End-to-End Response Time",
            "API Integration Latency",
            "Data Serialization Time",
            "Client-Server Round Trip",
            "Authentication Time",
            "State Synchronization Time",
            "Cache Coherency Ratio",
            "Concurrent User Sessions",
            "Bandwidth Utilization",
            "Resource Contention Index",
        ],
        DevelopmentType::DataScience => [
            "Data Processing Time",
            "Model Training Iteration",
            "Feature Extraction Time",
            "Data Transformation Throughput",
            "Prediction Latency",
            "Dataset Load Time",
            "Memory Utilization",
            "Parallel Worker Efficiency",
            "I/O Throughput",
            "Query Execution Time",
        ],
        DevelopmentType::DevOps => [
            "Deployment Time",
            "Build Duration",
            "Resource Provisioning Time",
            "Autoscaling Response Time",
            "Container Startup Time",
            "Service Discovery Latency",
            "Configuration Update Time",
            "Health Check Response Time",
            "Log Processing Rate",
            "Alert Processing Time",
        ],
        DevelopmentType::Blockchain => [
            "Transaction Validation Time",
            "Block Creation Time",
            "Consensus Round Duration",
            "Smart Contract Execution Time",
            "Network Propagation Delay",
            "Cryptographic Verification Time",
            "Merkle Tree Computation",
            "State Transition Latency",
            "Chain Sync Rate",
            "Gas Utilization Efficiency",
        ],
        DevelopmentType::MachineLearning => [
            "Model Inference Time",
            "Training Epoch Duration",
            "Feature Engineering Throughput",
            "Gradient Computation Time",
            "Batch Processing Rate",
            "Model Serialization Time",
            "Memory Utilization",
            "GPU Utilization",
            "Data Loading Throughput",
            "Hyperparameter Evaluation Time",
        ],
        DevelopmentType::SystemsProgramming => [
            "Memory Allocation Time",
            "Context Switch Overhead",
            "Lock Contention Ratio",
            "Cache Miss Rate",
            "Syscall Latency",
            "I/O Operation Throughput",
            "Thread Synchronization Time",
            "Memory Bandwidth Utilization",
            "Instruction Throughput",
            "Branch Prediction Accuracy",
        ],
        DevelopmentType::GameDevelopment => [
            "Frame Render Time",
            "Physics Simulation Time",
            "Asset Loading Duration",
            "Particle System Update Time",
            "Animation Blending Time",
            "AI Pathfinding Computation",
            "Collision Detection Time",
            "Memory Fragmentation Ratio",
            "Draw Call Count",
            "Audio Processing Latency",
        ],
        DevelopmentType::Security => [
            "Encryption/Decryption Time",
            "Authentication Latency",
            "Signature Verification Time",
            "Security Scan Duration",
            "Threat Detection Latency",
            "Policy Evaluation Time",
            "Access Control Check Latency",
            "Certificate Validation Time",
            "Secure Channel Establishment",
            "Log Analysis Throughput",
        ],
    };

    metrics.choose(&mut rng()).unwrap().to_string()
}
