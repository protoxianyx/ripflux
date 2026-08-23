package paths

type logFiles struct {
	INPUT_LOG      string
	CONSOLE_OUTPUT string
	BACKEND_LOGS   string
	TEST_LOGS      string
	COMBINED_LOGS  string
}

var LogFiles = logFiles{
	INPUT_LOG:      "inputs.log",
	CONSOLE_OUTPUT: "console_output.log",
	BACKEND_LOGS:   "backend_output.log",
	TEST_LOGS:      "test_output.log",
	COMBINED_LOGS:  "combined_backend.log",
}
