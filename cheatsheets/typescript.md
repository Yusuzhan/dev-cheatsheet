---
title: TypeScript
icon: fa-code
primary: "#3178C6"
lang: typescript
---

## fa-tag Basic Types

```typescript
let str: string = "hello";
let num: number = 42;
let bool: boolean = true;
let n: null = null;
let u: undefined = undefined;
let anyVal: any = "anything";
let unknownVal: unknown = "safe";

let arr: number[] = [1, 2, 3];
let tuple: [string, number] = ["age", 25];
let ro: readonly string[] = ["a", "b"];
```

## fa-pen-to-square Interfaces & Types

```typescript
interface User {
  name: string;
  age: number;
  email?: string;
  readonly id: number;
}

type Status = "active" | "inactive";
type ID = string | number;
type Point = { x: number; y: number };

type Nullable<T> = T | null;
type Partial<T> = { [K in keyof T]?: T[K] };
type Required<T> = { [K in keyof T]-?: T[K] };
```

## fa-code Functions

```typescript
function add(a: number, b: number): number {
  return a + b;
}

const multiply = (a: number, b: number): number => a * b;

function greet(name: string, greeting = "Hello"): string {
  return `${greeting}, ${name}`;
}

function log(...args: unknown[]): void {
  console.log(...args);
}

type Callback = (data: string) => void;
```

## fa-layer-group Generics

```typescript
function identity<T>(value: T): T {
  return value;
}

function first<T>(arr: T[]): T | undefined {
  return arr[0];
}

class Box<T> {
  constructor(public value: T) {}
}

const numBox = new Box(42);

function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}
```

## fa-sitemap Classes

```typescript
abstract class Animal {
  constructor(public name: string) {}

  abstract speak(): string;

  move(distance: number): string {
    return `${this.name} moved ${distance}m`;
  }
}

class Dog extends Animal {
  speak(): string {
    return `${this.name} barks`;
  }

  override move(distance: number): string {
    return `${super.move(distance)} happily`;
  }
}
```

## fa-puzzle-piece Enums & Literal Types

```typescript
enum Direction {
  Up = "UP",
  Down = "DOWN",
  Left = "LEFT",
  Right = "RIGHT",
}

const enum Color {
  Red,
  Green,
  Blue,
}

type Method = "GET" | "POST" | "PUT" | "DELETE";
type NumericRange = 1 | 2 | 3 | 4 | 5;
```

## fa-shield Type Guards

```typescript
function isString(val: unknown): val is string {
  return typeof val === "string";
}

function isUser(val: any): val is User {
  return typeof val.name === "string";
}

function process(val: string | number) {
  if (typeof val === "string") {
    return val.toUpperCase();
  }
  return val.toFixed(2);
}
```

## fa-link Modules

```typescript
export interface Config {
  host: string;
  port: number;
}

export function createConnection(config: Config): void {}

export default class Database {
  constructor(private config: Config) {}
}

import { Config, createConnection } from "./db";
import type { User } from "./types";
import Database from "./db";
```

## fa-diagram-project Async & Promises

```typescript
async function fetchUser(id: number): Promise<User> {
  const res = await fetch(`/api/users/${id}`);
  return res.json();
}

async function getUsers(): Promise<User[]> {
  try {
    const res = await fetch("/api/users");
    return await res.json();
  } catch (err) {
    console.error(err);
    return [];
  }
}

Promise.all([fetchUser(1), fetchUser(2)]);
Promise.allSettled([fetchUser(1), fetchUser(2)]);
```

## fa-arrows-turn-to-dots Mapped & Conditional Types

```typescript
type Readonly<T> = { readonly [K in keyof T]: T[K] };
type Pick<T, K extends keyof T> = { [P in K]: T[P] };
type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
type Record<K extends string | number | symbol, V> = { [P in K]: V };

type IsString<T> = T extends string ? true : false;
type Unwrap<T> = T extends Promise<infer U> ? U : T;
type Flatten<T> = T extends Array<infer U> ? U : T;
```

## fa-wand-magic Utility Types

```typescript
interface Todo {
  title: string;
  done: boolean;
  description: string;
}

type TodoPreview = Pick<Todo, "title" | "done">;
type TodoInfo = Omit<Todo, "description">;
type OptionalTodo = Partial<Todo>;
type StrictTodo = Required<OptionalTodo>;
type ReadonlyTodo = Readonly<Todo>;
type TodoMap = Record<string, Todo>;
type ValueOf<T> = T[keyof T];

type NonNullable<T> = T extends null | undefined ? never : T;
type ReturnType<T> = T extends (...args: any) => infer R ? R : never;
type Parameters<T> = T extends (...args: infer P) => any ? P : never;
```

## fa-lightbulb Useful Patterns

```typescript
const assertion
const point = { x: 10, y: 20 } as const;
type Coord = typeof point;

satisfies operator
const config = { host: "localhost", port: 3000 } satisfies Config;

Template literal types
type EventName = `on${Capitalize<string>}`;
type CSSValue = `${number}${"px" | "rem" | "em"}`;
```
