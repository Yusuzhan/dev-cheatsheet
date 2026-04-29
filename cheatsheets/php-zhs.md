---
title: PHP
icon: fa-php
primary: "#777BB4"
lang: php
locale: zhs
---

## fa-box 变量与类型

```php
$name = "Alice";
$age = 30;
$price = 9.99;
$active = true;
$nothing = null;

define('MAX_SIZE', 100);
const MIN_SIZE = 1;

$types = [
    gettype($name),
    is_string($name),
    is_int($age),
    is_array([]),
    is_null($nothing),
];

settype($val, "int");
(int)$val;
(string)$val;
```

## fa-font 字符串

```php
$s = "Hello, World!";
strlen($s);
strtoupper($s);
strtolower($s);
trim($s);
str_contains($s, "World");
str_starts_with($s, "Hello");
str_ends_with($s, "!");
str_replace("World", "PHP", $s);
explode(", ", $s);
implode(", ", $arr);
substr($s, 0, 5);
sprintf("Hello, %s! You are %d.", $name, $age);
"Hello, {$name}!"
```

## fa-list 数组

```php
$arr = [1, 2, 3, 4, 5];
$arr[] = 6;
array_push($arr, 7);
array_pop($arr);
array_shift($arr);
array_unshift($arr, 0);
count($arr);
in_array(3, $arr);
array_search(3, $arr);
array_slice($arr, 1, 3);
array_merge($a, $b);
array_unique($arr);
sort($arr);
array_map(fn($n) => $n * 2, $arr);
array_filter($arr, fn($n) => $n > 2);
array_keys($assoc);
array_values($assoc);
```

## fa-code-branch 控制流

```php
if ($x > 0) {
    echo "positive";
} elseif ($x === 0) {
    echo "zero";
} else {
    echo "negative";
}

$result = $x > 0 ? "pos" : "neg";
$result = $x ?? "default";

switch ($status) {
    case 'running':
        echo "ok";
        break;
    case 'stopped':
        echo "down";
        break;
    default:
        echo "unknown";
}

match($status) {
    'running' => 'ok',
    'stopped' => 'down',
    default => 'unknown',
};
```

## fa-pen-to-square 函数

```php
function greet(string $name): string {
    return "Hello, {$name}";
}

function add(int $a, int $b): int {
    return $a + $b;
}

function sum(int ...$nums): int {
    return array_sum($nums);
}

$double = fn($x) => $x * 2;

function process(array $items, callable $fn): array {
    return array_map($fn, $items);
}
```

## fa-cubes 类与 OOP

```php
class User
{
    public function __construct(
        public string $name,
        public int $age,
    ) {}

    public function greet(): string {
        return "I'm {$this->name}";
    }
}

class Admin extends User
{
    public function __construct(
        string $name,
        int $age,
        public string $role = 'admin',
    ) {
        parent::__construct($name, $age);
    }
}

$user = new User("Alice", 30);
$user->name;
```

## fa-puzzle-piece Trait 与接口

```php
interface Loggable
{
    public function log(string $message): void;
}

trait Timestamps
{
    public function createdAt(): string {
        return date('Y-m-d H:i:s');
    }
}

class Order implements Loggable
{
    use Timestamps;

    public function log(string $message): void {
        error_log($message);
    }
}
```

## fa-triangle-exclamation 错误处理

```php
try {
    $result = risky();
} catch (InvalidArgumentException $e) {
    echo $e->getMessage();
} catch (RuntimeException $e) {
    echo $e->getMessage();
} finally {
    cleanup();
}

throw new InvalidArgumentException("invalid input");

set_exception_handler(function ($e) {
    error_log($e->getMessage());
});

error_reporting(E_ALL);
ini_set('display_errors', '1');
```

## fa-file 文件 I/O

```php
file_put_contents("out.txt", "hello");
$content = file_get_contents("in.txt");
$lines = file("in.txt", FILE_IGNORE_NEW_LINES);

$handle = fopen("file.txt", "r");
while (($line = fgets($handle)) !== false) {
    echo $line;
}
fclose($handle);

file_exists("file.txt");
is_file("file.txt");
is_dir("dir");
mkdir("dir");
unlink("file.txt");
```

## fa-globe 超全局变量

```php
$_GET['key'];
$_POST['key'];
$_REQUEST['key'];
$_SERVER['HTTP_HOST'];
$_SERVER['REQUEST_METHOD'];
$_SESSION['user_id'];
$_COOKIE['token'];
$_FILES['upload']['tmp_name'];
$_ENV['API_KEY'];

htmlspecialchars($input, ENT_QUOTES, 'UTF-8');
filter_input(INPUT_GET, 'id', FILTER_VALIDATE_INT);
```

## fa-database PDO / MySQL

```php
$pdo = new PDO('mysql:host=localhost;dbname=test', 'user', 'pass');
$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

$stmt = $pdo->prepare('SELECT * FROM users WHERE id = ?');
$stmt->execute([$id]);
$user = $stmt->fetch(PDO::FETCH_ASSOC);

$stmt = $pdo->prepare('INSERT INTO users (name, email) VALUES (?, ?)');
$stmt->execute([$name, $email]);

$users = $pdo->query('SELECT * FROM users')->fetchAll(PDO::FETCH_OBJ);
$pdo->beginTransaction();
$pdo->commit();
$pdo->rollBack();
```

## fa-boxes-stacked Composer

```sh
composer init
composer install
composer require vendor/package
composer require --dev phpunit/phpunit
composer update
composer dump-autoload
composer dump-autoload -o
composer show
composer outdated
```

## fa-key Session 与 Cookie

```php
session_start();
$_SESSION['user'] = 'Alice';
$name = $_SESSION['user'] ?? '';
session_destroy();

setcookie('token', 'abc', [
    'expires' => time() + 3600,
    'path' => '/',
    'secure' => true,
    'httponly' => true,
    'samesite' => 'Strict',
]);

$token = $_COOKIE['token'] ?? '';
```

## fa-calendar DateTime

```php
$now = new DateTime();
$now->format('Y-m-d H:i:s');
$now->modify('+1 day');
$now->modify('+1 month');

$date = new DateTime('2024-01-15');
$diff = $now->diff($date);
$diff->days;

DateTime::createFromFormat('Y-m-d', '2024-01-15');
date('Y-m-d');
strtotime('+1 week');
$interval = new DateInterval('P1D');
$period = new DatePeriod($start, $interval, $end);
```
