---
title: Dart
icon: fa-dart
primary: "#0175C2"
lang: dart
---

## fa-box Variables & Types

```dart
final String name = 'Alice';
const int max = 100;
var count = 0;
double score = 95.5;
bool active = true;
dynamic anything = 42;
anything = 'hello';

Object obj = 'string';
String? nullable;

int x = 5;
double d = x.toDouble();
String s = 42.toString();
int.parse('42');
```

## fa-shield Null Safety

```dart
String? name;
String nonNull = name ?? 'default';
String result = name!.toUpperCase();

if (name != null) {
  print(name.length);
}

String? maybeNull;
String val = maybeNull ?? 'fallback';
int? len = maybeNull?.length;
```

## fa-font Strings

```dart
var s = 'Hello, World!';
s.length;
s.isEmpty;
s.toUpperCase();
s.toLowerCase();
s.contains('World');
s.startsWith('Hello');
s.endsWith('!');
s.replaceAll('o', '0');
s.substring(0, 5);
s.split(', ');
var parts = ['a', 'b'].join(', ');
var msg = 'Hello, $name!';
var expr = '2 + 2 = ${2 + 2}';
var multi = '''
  line one
  line two
  ''';
```

## fa-code-branch Control Flow

```dart
if (x > 0) {
  print('positive');
} else if (x == 0) {
  print('zero');
} else {
  print('negative');
}

var result = x > 0 ? 'pos' : 'neg';

switch (status) {
  case 'running':
    print('ok');
  case 'stopped':
    print('down');
  default:
    print('unknown');
}

for (var i = 0; i < 10; i++) { print(i); }
for (var item in list) { print(item); }
while (condition) { }
do { } while (condition);
```

## fa-pen-to-square Functions & Closures

```dart
String greet(String name) => 'Hello, $name';

int add(int a, int b) {
  return a + b;
}

void log(String msg, [String level = 'INFO']) {
  print('[$level] $msg');
}

int sum(int a, int b, {int c = 0}) => a + b + c;

var numbers = [1, 2, 3];
var doubled = numbers.map((n) => n * 2).toList();

var multiplier = (int n) => n * 3;
```

## fa-cubes Classes & Mixins

```dart
class User {
  final String name;
  int _age;

  User(this.name, this._age);

  String greet() => 'I am $name';
}

class Admin extends User {
  final String role;
  Admin(String name, int age, this.role) : super(name, age);
}

mixin Loggable {
  void log(String msg) => print('[LOG] $msg');
}

class Service with Loggable {
  void run() => log('running');
}
```

## fa-list Collections

```dart
var list = [1, 2, 3];
list.add(4);
list.addAll([5, 6]);
list.remove(3);
list.indexOf(2);
list.sort();
list.reversed;

var set = <int>{1, 2, 3};
set.add(4);
set.contains(2);
set.intersection({2, 3, 4});

var map = {'a': 1, 'b': 2};
map['c'] = 3;
map.containsKey('a');
map.remove('a');
map.keys;
map.values;
```

## fa-layer-group Generics

```dart
T first<T>(List<T> items) => items.first;

class Box<T> {
  final T value;
  Box(this.value);
  R map<R>(R Function(T) fn) => fn(value);
}

var intBox = Box(42);
var strBox = Box('hello');

List<T> filter<T>(List<T> list, bool Function(T) test) {
  return list.where(test).toList();
}
```

## fa-arrows-spin Async/Await & Future

```dart
Future<String> fetchData() async {
  await Future.delayed(Duration(seconds: 1));
  return 'data';
}

Future<void> main() async {
  var data = await fetchData();
  print(data);
}

Future<int> compute() {
  return Future.value(42);
}

Future.wait([fetchA(), fetchB(), fetchC()]);
Future.any([fetchFast1(), fetchFast2()]);
```

## fa-water Streams

```dart
Stream<int> countStream(int max) async* {
  for (var i = 0; i < max; i++) {
    await Future.delayed(Duration(seconds: 1));
    yield i;
  }
}

await for (var value in countStream(5)) {
  print(value);
}

var controller = StreamController<int>();
controller.add(1);
controller.add(2);
await controller.close();
```

## fa-triangle-exclamation Error Handling

```dart
try {
  var result = await riskyOperation();
} on FormatException catch (e) {
  print('Format error: $e');
} catch (e, stackTrace) {
  print('Error: $e');
  print(stackTrace);
} finally {
  cleanup();
}

throw Exception('something went wrong');
throw FormatException('invalid format');
```

## fa-puzzle-piece Extension Methods

```dart
extension StringX on String {
  String get capitalized =>
      isEmpty ? '' : '${this[0].toUpperCase()}${substring(1)}';

  int get wordCount => split(RegExp(r'\s+')).length;
}

'hello world'.capitalized;
'hello world'.wordCount;

extension ListX<T> on List<T> {
  T? get second => length > 1 ? this[1] : null;
}
```

## fa-layer-group Enums & Sealed Classes

```dart
enum Color { red, green, blue }
Color.values;
Color.red.name;

enum Status with { ok, error, loading }
switch (status) {
  case Status.ok: print('ok');
  case Status.error: print('err');
  case Status.loading: print('loading');
}

sealed class Result<T> {}
class Success<T> extends Result<T> {
  final T value;
  Success(this.value);
}
class Failure<T> extends Result<T> {
  final String message;
  Failure(this.message);
}
```

## fa-lightbulb Common Patterns

```dart
var json = {'name': 'Alice', 'age': 30};
var user = User.fromJson(json);

var users = list.where((u) => u.age > 18).toList();
var names = list.map((u) => u.name).toList();

final cache = <String, dynamic>{};
T memoize<T>(String key, T Function() compute) {
  return cache.putIfAbsent(key, compute) as T;
}

list.sort((a, b) => a.name.compareTo(b.name));
list.fold<int>(0, (sum, item) => sum + item.value);
```
