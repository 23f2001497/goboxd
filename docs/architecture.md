# GoBoxD Architecture

## Overview

GoBoxD is an HTTP service that executes user-submitted code and returns execution results.

## Architecture

HTTP Client
    |
    v
API Layer
    |
    v
Execution Service
    |
    v
Workspace Manager
    |
    v
Language Registry
    |
    v
Sandbox Runner
    |
    v
Result Formatter

## Component Responsibilities

API Layer
Receives HTTP requests.

Execution Service
Coordinates execution flow.

Workspace Manager
Creates and cleans temporary workspaces.

Language Registry
Maps language names to runtime commands.

Sandbox Runner
Executes submitted code.

Result Formatter
Converts execution results into JSON.

## Execution Flow

Client Request
    ↓
API Layer
    ↓
Execution Service
    ↓
Workspace Creation
    ↓
Source File Generation
    ↓
Program Execution
    ↓
Result Formatting
    ↓
JSON Response

## Current Stage 1 Status

Implemented:
- health endpoint
- run endpoint
- workspace manager
- python execution
- validation
- tests

Planned:
- nsjail
- cgroups
- resource limits
- multiple languages