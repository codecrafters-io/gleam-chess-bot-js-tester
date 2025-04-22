package internal

import (
	"time"

	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{
		{
			Slug:     "anti-cheat-1",
			TestFunc: antiCheatExecute,
		},
		{
			Slug:     "anti-cheat-2",
			TestFunc: antiCheatDeps,
		},
	},
	ExecutableFileName: "your_program.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "qc6",
			TestFunc: test1,
		},
		{
			Slug:     "zo3",
			TestFunc: test2,
		},
		{
			Slug:     "xs2",
			TestFunc: test3,
			Timeout:  20 * time.Second,
		},
		{
			Slug:     "ix8",
			TestFunc: test4,
			Timeout:  20 * time.Second,
		},
		{
			Slug:     "dx9",
			TestFunc: test5,
			Timeout:  20 * time.Second,
		},
	},
}
