# 🧠 Logic Gate Simulation in Go

In digital logic design, the **NAND gate** is considered the most fundamental (primitive) gate. This is because **all other logic gates**—such as AND, OR, NOT, and XOR—can be constructed using only NAND gates.

However, in this simulation, to keep things simple and avoid dealing with low-level concepts like electrical voltage or transistor behavior, we leverage Go’s built-in **bitwise operators** to implement the logic gates. We start with a simulated NAND gate using bitwise operations and build the rest of the logic from there.

## ✅ Goals

- Simulate logic gates using Go
- Understand how to compose gates
- Learn how NAND can serve as the foundation of all logic

## 🛠️ Tech Stack

- Language: Go
- Concepts: Digital Logic, Bitwise Operations, Functional Composition

## 📂 Structure

```go
type Gate interface {
    Eval() uint8
}

type NAND struct { ... }
type AND  struct { ... }
type OR   struct { ... }
type NOT  struct { ... }
...