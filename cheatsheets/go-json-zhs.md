---
title: Go encoding/json
icon: fa-brackets-curly
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-arrows-left-right Marshal / Unmarshal

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

user := User{Name: "Alice", Email: "alice@example.com", Age: 30}

data, err := json.Marshal(user)
data, err = json.MarshalIndent(user, "", "  ")

var result User
err = json.Unmarshal(data, &result)
```

## fa-tags Struct Tags

```go
type Product struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    SKU      string  `json:"sku,omitempty"`
    Internal string  `json:"-"`
    Count    int     `json:"item_count"`
}
```

## fa-pen-to-square Custom MarshalJSON

```go
type Date struct {
    time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
    return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}

type Person struct {
    Name string `json:"name"`
    DOB  Date   `json:"dob"`
}

func (p Person) MarshalJSON() ([]byte, error) {
    type Alias Person
    return json.Marshal(&struct {
        Alias
        FullName string `json:"full_name"`
    }{
        Alias:    Alias(p),
        FullName: p.Name,
    })
}
```

## fa-pen Custom UnmarshalJSON

```go
type FlexInt int

func (fi *FlexInt) UnmarshalJSON(data []byte) error {
    var v interface{}
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch t := v.(type) {
    case float64:
        *fi = FlexInt(t)
    case string:
        n, err := strconv.Atoi(t)
        if err != nil {
            return err
        }
        *fi = FlexInt(n)
    }
    return nil
}
```

## fa-stream JSON Encoder / Decoder (Streaming)

```go
var buf bytes.Buffer
enc := json.NewEncoder(&buf)
enc.SetIndent("", "  ")
enc.Encode(user)

dec := json.NewDecoder(strings.NewReader(`{"name":"Alice"}`))
var u User
dec.Decode(&u)

dec = json.NewDecoder(file)
for dec.More() {
    var item Item
    if err := dec.Decode(&item); err != nil {
        break
    }
    process(item)
}
```

## fa-shuffle Dynamic JSON (map/interface)

```go
var data map[string]interface{}
json.Unmarshal([]byte(`{"name":"Alice","age":30}`), &data)
name := data["name"].(string)
age := data["age"].(float64)

m := map[string]any{
    "name": "Alice",
    "tags": []string{"go", "dev"},
}
b, _ := json.Marshal(m)

var raw interface{}
json.Unmarshal(jsonStr, &raw)
```

## fa-layer-group Nested Structs

```go
type Address struct {
    City    string `json:"city"`
    Country string `json:"country"`
}

type Employee struct {
    Name    string  `json:"name"`
    Address Address `json:"address"`
}

emp := Employee{
    Name: "Alice",
    Address: Address{City: "Beijing", Country: "CN"},
}

data, _ := json.Marshal(emp)
```

## fa-eye-slash OmitEmpty & Ignore Fields

```go
type Request struct {
    ID     int    `json:"id"`
    Name   string `json:"name,omitempty"`
    Secret string `json:"-"`
    Data   []byte `json:"data,omitempty"`
}

req := Request{ID: 1, Secret: "hidden"}
b, _ := json.Marshal(req)
```

## fa-list JSON Arrays

```go
type Item struct {
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

items := []Item{
    {Name: "Apple", Price: 1.5},
    {Name: "Banana", Price: 0.8},
}

data, _ := json.Marshal(items)

var result []Item
json.Unmarshal(data, &result)

var names []string
json.Unmarshal([]byte(`["a","b","c"]`), &names)
```

## fa-clock Time Handling

```go
type Event struct {
    Time time.Time `json:"time"`
}

e := Event{Time: time.Now()}
data, _ := json.Marshal(e)

var parsed Event
json.Unmarshal(data, &parsed)

type EventCustom struct {
    Time string `json:"time"`
}

ec := EventCustom{Time: time.Now().Format(time.RFC3339)}
```

## fa-paintbrush Pretty Print

```go
data, _ := json.MarshalIndent(user, "", "    ")
fmt.Println(string(data))

var out bytes.Buffer
json.Indent(&out, rawJSON, "", "  ")
out.WriteTo(os.Stdout)

var obj interface{}
json.Unmarshal(rawJSON, &obj)
pretty, _ := json.MarshalIndent(obj, "", "  ")
```

## fa-box-open RawMessage

```go
type Envelope struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`
}

raw := []byte(`{"type":"user","data":{"name":"Alice"}}`)
var env Envelope
json.Unmarshal(raw, &env)

if env.Type == "user" {
    var user User
    json.Unmarshal(env.Data, &user)
}
```

## fa-shield DisallowUnknownFields

```go
dec := json.NewDecoder(strings.NewReader(`{"name":"Alice","unknown":"x"}`))
dec.DisallowUnknownFields()

var u User
err := dec.Decode(&u)
```

## fa-toolbox Common Patterns

```go
type APIResponse struct {
    Status  string      `json:"status"`
    Data    interface{} `json:"data,omitempty"`
    Message string      `json:"message,omitempty"`
}

resp := APIResponse{Status: "ok", Data: user}
b, _ := json.Marshal(resp)

type Pagination struct {
    Page  int `json:"page"`
    Total int `json:"total"`
}

type PagedResponse struct {
    Data       []User     `json:"data"`
    Pagination Pagination `json:"pagination"`
}
```
