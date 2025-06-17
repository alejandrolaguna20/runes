package cards

func GenerateMockCards() []Card {
	return []Card{
		{
			Front:          "What does 'Big O' notation describe?",
			Back:           "Big O notation describes the upper bound of an algorithm's time or space complexity in the worst-case scenario. It helps compare algorithm efficiency as input size grows.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is the difference between a stack and a queue?",
			Back:           "Stack: Last In, First Out (LIFO) - elements added/removed from the top. Queue: First In, First Out (FIFO) - elements added at rear, removed from front.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is a closure in programming?",
			Back:           "A closure is a function that captures and retains access to variables from its outer/enclosing scope, even after the outer function has returned.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What's the time complexity of binary search?",
			Back:           "O(log n) - Binary search eliminates half the search space with each comparison, resulting in logarithmic time complexity.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is polymorphism in OOP?",
			Back:           "Polymorphism allows objects of different types to be treated as instances of the same type through a common interface, enabling one interface to represent different underlying forms.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is the difference between '==' and '===' in JavaScript?",
			Back:           "'==' performs type coercion and compares values. '===' (strict equality) compares both value and type without coercion. Always prefer '===' for predictable behavior.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is a race condition?",
			Back:           "A race condition occurs when multiple threads/processes access shared data concurrently, and the outcome depends on the unpredictable timing of their execution.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is memoization?",
			Back:           "Memoization is an optimization technique that stores the results of expensive function calls and returns cached results for identical inputs to avoid redundant calculations.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What's the difference between HTTP and HTTPS?",
			Back:           "HTTP transmits data in plain text. HTTPS adds SSL/TLS encryption for secure data transmission, authentication, and data integrity protection.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
		{
			Front:          "What is dependency injection?",
			Back:           "Dependency injection is a design pattern where dependencies are provided to an object from external sources rather than the object creating them itself, improving testability and flexibility.",
			IsFlipped:      false,
			timesCorrect:   0,
			timesIncorrect: 0,
		},
	}
}
