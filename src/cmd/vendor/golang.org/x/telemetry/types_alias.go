package telemetry

import "golang.org/x/telemetry/internal/telemetry"

// Common types and directories used by multiple packages.

// An UploadConfig controls what data is uploaded.
type UploadConfig = telemetry.UploadConfig

type ProgramConfig = telemetry.ProgramConfig

type CounterConfig = telemetry.CounterConfig

// A Report is what's uploaded (or saved locally)
type Report = telemetry.Report

type ProgramReport = telemetry.ProgramReport
