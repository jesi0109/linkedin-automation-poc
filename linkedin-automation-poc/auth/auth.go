package auth

import "linkedin-automation-poc/logger"

// Mock authentication for demonstration
func Login() bool {
	logger.Info("Simulating login process (POC)")
	return true
}
