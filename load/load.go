//package load helps the mars load redcode warriors into the core based on the
// configuration. this package must:
// 1) lex all supplied warriors
// 2) parse and simplify all lables and expressions
// 3) read config and inspect existing core
// 4) return byte arrays and locations that they should be placed in the core
//
package load
