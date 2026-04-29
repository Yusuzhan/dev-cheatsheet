---
title: React Hooks
icon: fa-atom
primary: "#61DAFB"
lang: jsx
---

## fa-database useState

```jsx
const [count, setCount] = useState(0);
const [user, setUser] = useState({ name: "", age: 0 });

setCount(prev => prev + 1);
setUser(prev => ({ ...prev, name: "Alice" }));
```

## fa-rotate useEffect

```jsx
useEffect(() => {
  document.title = `Count: ${count}`;
  return () => {
    document.title = "App";
  };
}, [count]);

useEffect(() => {
  const controller = new AbortController();
  fetch("/api/data", { signal: controller.signal })
    .then(res => res.json())
    .then(setData);
  return () => controller.abort();
}, []);
```

## fa-thumbtack useRef

```jsx
const inputRef = useRef(null);
const renderCount = useRef(0);

const focus = () => inputRef.current.focus();

return <input ref={inputRef} />;
```

## fa-layer-group useContext

```jsx
const ThemeContext = createContext("light");

function App() {
  return (
    <ThemeContext.Provider value="dark">
      <Child />
    </ThemeContext.Provider>
  );
}

function Child() {
  const theme = useContext(ThemeContext);
  return <div className={theme}>Current: {theme}</div>;
}
```

## fa-bolt useMemo

```jsx
const sorted = useMemo(
  () => items.filter(i => i.active).sort((a, b) => a.name.localeCompare(b.name)),
  [items]
);

const expensiveValue = useMemo(() => compute(a, b), [a, b]);
```

## fa-reply useCallback

```jsx
const handleClick = useCallback(() => {
  setCount(prev => prev + 1);
}, []);

const handleSubmit = useCallback((e) => {
  e.preventDefault();
  saveForm(formData);
}, [formData]);

return <Child onClick={handleClick} onSubmit={handleSubmit} />;
```

## fa-code-branch useReducer

```jsx
const reducer = (state, action) => {
  switch (action.type) {
    case "increment": return { count: state.count + 1 };
    case "decrement": return { count: state.count - 1 };
    case "reset": return { count: 0 };
    default: return state;
  }
};

const [state, dispatch] = useReducer(reducer, { count: 0 });

dispatch({ type: "increment" });
```

## fa-puzzle-piece Custom Hooks

```jsx
function useDebounce(value, delay) {
  const [debounced, setDebounced] = useState(value);
  useEffect(() => {
    const timer = setTimeout(() => setDebounced(value), delay);
    return () => clearTimeout(timer);
  }, [value, delay]);
  return debounced;
}

function useLocalStorage(key, initial) {
  const [value, setValue] = useState(() => {
    const stored = localStorage.getItem(key);
    return stored ? JSON.parse(stored) : initial;
  });
  useEffect(() => {
    localStorage.setItem(key, JSON.stringify(value));
  }, [key, value]);
  return [value, setValue];
}
```

## fa-ruler-combined useLayoutEffect

```jsx
useLayoutEffect(() => {
  const height = ref.current.getBoundingClientRect().height;
  setTooltipPos(height + 8);
}, [deps]);
```

## fa-hand-pointer useImperativeHandle

```jsx
const FancyInput = forwardRef((props, ref) => {
  const inputRef = useRef();
  useImperativeHandle(ref, () => ({
    focus: () => inputRef.current.focus(),
    clear: () => { inputRef.current.value = ""; },
  }));
  return <input ref={inputRef} />;
});

const inputRef = useRef();
<inputRef.current?.focus() />;
```

## fa-fingerprint useId

```jsx
function Form() {
  const id = useId();
  return (
    <>
      <label htmlFor={id + "-name"}>Name</label>
      <input id={id + "-name"} />
      <label htmlFor={id + "-email"}>Email</label>
      <input id={id + "-email"} />
    </>
  );
}
```

## fa-bug useDebugValue

```jsx
function useOnlineStatus() {
  const [online, setOnline] = useState(navigator.onLine);
  useEffect(() => {
    const goOnline = () => setOnline(true);
    const goOffline = () => setOnline(false);
    window.addEventListener("online", goOnline);
    window.addEventListener("offline", goOffline);
    return () => {
      window.removeEventListener("online", goOnline);
      window.removeEventListener("offline", goOffline);
    };
  }, []);
  useDebugValue(online ? "Online" : "Offline");
  return online;
}
```

## fa-list-check Rules of Hooks

```jsx
// Only call Hooks at the top level
function Bad() {
  if (condition) {
    useState(0);
  }
}

function Good() {
  const [val, setVal] = useState(0);
  if (condition) {
    // conditional logic using val
  }
}

// Only call Hooks from React functions
// Good: function components or custom hooks
// Bad: regular JS functions, class components
```

## fa-shapes Common Patterns

```jsx
function useFetch(url) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    const controller = new AbortController();
    setLoading(true);
    fetch(url, { signal: controller.signal })
      .then(res => { if (!res.ok) throw new Error(res.statusText); return res.json(); })
      .then(setData)
      .catch(err => { if (err.name !== "AbortError") setError(err.message); })
      .finally(() => setLoading(false));
    return () => controller.abort();
  }, [url]);
  return { data, loading, error };
}

function useToggle(initial = false) {
  const [value, setValue] = useState(initial);
  const toggle = useCallback(() => setValue(v => !v), []);
  return [value, toggle];
}

function usePrevious(value) {
  const ref = useRef();
  useEffect(() => { ref.current = value; });
  return ref.current;
}
```
