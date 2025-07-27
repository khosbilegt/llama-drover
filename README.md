# Llama Drover (WIP)

A distributed load balancer and coordinator for Large Language Model (LLM) inference nodes. Llama Drover intelligently routes LLM requests to the fastest available nodes in your cluster.

## Overview

Llama Drover consists of a central coordinator server that accepts HTTP API calls for LLM questions and intelligently delegates them to connected worker nodes via gRPC. The coordinator maintains real-time statistics about each node's performance and health, ensuring requests are always routed to the most optimal available node.

## Architecture

```
┌─────────────────┐    HTTP API    ┌─────────────────┐
│   Client Apps   │ ────────────►  │   Coordinator   │
└─────────────────┘                │     Server      │
                                   └─────────┬───────┘
                                             │ gRPC
                              ┌──────────────┼──────────────┐
                              │              │              │
                              ▼              ▼              ▼
                    ┌─────────────┐ ┌─────────────┐ ┌─────────────┐
                    │ LLM Node 1  │ │ LLM Node 2  │ │ LLM Node N  │
                    └─────────────┘ └─────────────┘ └─────────────┘
```

## Features

### Coordinator Server

- **HTTP API**: RESTful endpoints for LLM inference requests
- **Intelligent Load Balancing**: Routes requests to the fastest available nodes
- **Real-time Monitoring**: Tracks node health status and performance metrics
- **Statistics Collection**: Maintains average response times and success rates
- **Health Checks**: Continuous monitoring of node availability
- **Failover Support**: Automatic routing around unhealthy nodes

### Node Management

- **gRPC Communication**: High-performance communication with worker nodes
- **Performance Tracking**: Real-time latency and throughput monitoring
- **Dynamic Node Discovery**: Automatic registration and deregistration of nodes
- **Load Distribution**: Smart request distribution based on current node capacity
