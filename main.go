package main

import "fmt"

func main() {
	// Function call.
	test("a()", "a()")
	test("a(b)", "a(b)")
	test("a(b, c)", "a(b, c)")
	test("a(b)(c)", "a(b)(c)")
	test("a(b) + c(d)", "(a(b) + c(d))")
	test("a(b ? c : d, e + f)", "a((b ? c : d), (e + f))")

	// Unary precedence.
	test("~!-+a", "(~(!(-(+a))))")
	test("a!!!", "(((a!)!)!)")

	// Unary and binary predecence.
	test("-a * b", "((-a) * b)")
	test("!a + b", "((!a) + b)")
	test("~a ^ b", "((~a) ^ b)")
	test("-a!", "(-(a!))")
	test("!a!", "(!(a!))")

	// Binary precedence.
	test("a = b + c * d ^ e - f / g", "(a = ((b + (c * (d ^ e))) - (f / g)))")

	// Binary associativity.
	test("a = b = c", "(a = (b = c))")
	test("a + b - c", "((a + b) - c)")
	test("a * b / c", "((a * b) / c)")
	test("a ^ b ^ c", "(a ^ (b ^ c))")

	// Conditional operator.
	test("a ? b : c ? d : e", "(a ? b : (c ? d : e))")
	test("a ? b ? c : d : e", "(a ? (b ? c : d) : e)")
	test("a + b ? c * d : e / f", "((a + b) ? (c * d) : (e / f))")

	// Grouping.
	test("a + (b + c) + d", "((a + (b + c)) + d)")
	test("a ^ (b + c)", "(a ^ (b + c))")
	test("(!a)!", "((!a)!)")
}

func test(source, expected string) {
	fmt.Println(source, expected)
}
