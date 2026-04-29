---
title: Zig
icon: fa-bolt
primary: "#F7A41D"
lang: zig
---

## fa-box Variables & Constants

```zig
const x: i32 = 42;
var y: f64 = 3.14;
const z = @as(u8, 100);
const pi: comptime_float = 3.14159;

const yes: bool = true;
const letter: u8 = 'A';
const byte: u8 = 0xFF;

var n: usize = 0;
n += 1;
```

## fa-code Primitive Types

```zig
i8  i16  i32  i64  i128  isize
u8  u16  u32  u64  u128  usize
f16  f32  f64  f80  f128
bool  void  noreturn  type
*const T  ?T  anyerror!T
[c]u8  []const u8  [*]const u8
```

## fa-code-branch Control Flow

```zig
if (x > 0) {
    print("positive\n", .{});
} else if (x == 0) {
    print("zero\n", .{});
} else {
    print("negative\n", .{});
}

const abs = if (x < 0) -x else x;

for (0..10) |i| {
    print("{}\n", .{i});
}

while (i < 10) : (i += 1) {
    print("{}\n", .{i});
}

while (node) |n| : (node = n.next) {
    print("{}\n", .{n.data});
}

loop: {
    break :loop;
}
```

## fa-pen-to-square Functions

```zig
fn add(a: i32, b: i32) i32 {
    return a + b;
}

fn greet(name: []const u8) void {
    print("Hello, {s}!\n", .{name});
}

fn maybeValue() ?i32 {
    return 42;
}

fn fallible() !i32 {
    return error.SomethingWrong;
}

const add_fn = fn (i32, i32) i32;
```

## fa-cubes Structs

```zig
const Point = struct {
    x: f64,
    y: f64,

    fn init(x: f64, y: f64) Point {
        return .{ .x = x, .y = y };
    }

    fn distance(self: Point, other: Point) f64 {
        const dx = self.x - other.x;
        const dy = self.y - other.y;
        return @sqrt(dx * dx + dy * dy);
    }
};

const p = Point{ .x = 1.0, .y = 2.0 };
const p2 = Point.init(3.0, 4.0);
```

## fa-layer-group Enums & Tagged Unions

```zig
const Color = enum { red, green, blue };
const c: Color = .red;

const Shape = union(enum) {
    circle: f64,
    rectangle: struct { w: f64, h: f64 },
    triangle,

    fn area(self: Shape) f64 {
        return switch (self) {
            .circle => |r| std.math.pi * r * r,
            .rectangle => |r| r.w * r.h,
            .triangle => 0.0,
        };
    }
};

const s = Shape{ .circle = 5.0 };
```

## fa-triangle-exclamation Error Handling

```zig
const FileOpenError = error{
    NotFound,
    PermissionDenied,
    TooBig,
};

fn readFile(path: []const u8) ![]const u8 {
    const data = std.fs.cwd().readFileAlloc(allocator, path, max_size) catch |err| {
        return err;
    };
    return data;
}

const result = fallible() catch |err| switch (err) {
    error.NotFound => "default",
    else => return err,
};

const val = try fallible();
const opt = maybeValue() orelse 0;
```

## fa-list Arrays & Slices

```zig
const arr = [5]i32{ 1, 2, 3, 4, 5 };
const slice: []const i32 = arr[1..3];

var buf: [1024]u8 = undefined;
var dynamic = std.ArrayList(i32).init(allocator);
try dynamic.append(42);

for (arr, 0..) |val, i| {
    print("[{}] = {}\n", .{ i, val });
}

const len = arr.len;
const first = arr[0];
const last = arr[arr.len - 1];
```

## fa-arrow-pointer Pointers

```zig
var x: i32 = 42;
const ptr: *i32 = &x;
ptr.* += 1;

const cptr: *const i32 = &x;

var buf: [10]u8 = undefined;
const many: [*]u8 = &buf;
const slice: []u8 = many[0..10];

const optional_ptr: ?*i32 = null;
if (optional_ptr) |p| {
    p.* += 1;
}
```

## fa-memory-stick Allocators

```zig
var gpa = std.heap.GeneralPurposeAllocator(.{}){};
const allocator = gpa.allocator();

var list = std.ArrayList(i32).init(allocator);
defer list.deinit();
try list.append(42);

var map = std.StringHashMap(i32).init(allocator);
defer map.deinit();
try map.put("key", 1);

var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
defer arena.deinit();
const arena_alloc = arena.allocator();
```

## fa-font Strings

```zig
const msg: []const u8 = "hello world";
const len = msg.len;

const buf = try std.fmt.allocPrint(allocator, "{s} is {d}", .{ "answer", 42 });

std.debug.print("{s}\n", .{msg});

const eql = std.mem.eql(u8, "abc", "abc");
const starts = std.mem.startsWith(u8, msg, "hello");
const parts = std.mem.split(u8, "a,b,c", ",");

const upper = try std.ascii.allocUpperString(allocator, "hello");
```

## fa-brain Comptime

```zig
fn Matrix(comptime T: type, comptime width: usize) type {
    return [width][width]T;
}

const Mat3f = Matrix(f32, 3);
var m: Mat3f = undefined;

comptime {
    assert(@sizeOf(u32) == 4);
}

fn max(comptime T: type, a: T, b: T) T {
    return if (a > b) a else b;
}

const type_info = @typeInfo(Point);
const fields = @typeInfo(Point).Struct.fields;
```

## fa-gears Build System

```zig
const std = @import("std");

pub fn build(b: *std.Build) void {
    const exe = b.addExecutable(.{
        .name = "app",
        .root_source_file = b.path("src/main.zig"),
        .target = b.standardTargetOptions(.{}),
        .optimize = b.standardOptimizeOption(.{}),
    });

    exe.linkLibC();
    b.installArtifact(exe);

    const run_cmd = b.addRunArtifact(exe);
    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);

    const unit_tests = b.addTest(.{
        .root_source_file = b.path("src/main.zig"),
    });
    const test_step = b.step("test", "Run unit tests");
    test_step.dependOn(&unit_tests.step);
}
```

```sh
zig build
zig build run
zig build test
zig build -Dtarget=x86_64-linux
zig build -Doptimize=ReleaseFast
```

## fa-link Interop with C

```zig
const c = @cImport({
    @cInclude("stdio.h");
    @cInclude("stdlib.h");
});

pub fn main() void {
    c.printf("hello from C: %d\n", 42);
}

const MyStruct = extern struct {
    x: c_int,
    y: c_int,
};

extern fn malloc(size: usize) ?*anyopaque;
extern fn free(ptr: ?*anyopaque) void;

export fn add(a: i32, b: i32) i32 {
    return a + b;
}
```
