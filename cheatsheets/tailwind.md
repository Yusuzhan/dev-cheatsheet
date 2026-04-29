---
title: Tailwind CSS
icon: fa-wind
primary: "#06B6D4"
lang: html
---

## fa-table-columns Layout (Flex/Grid)

```html
<div class="flex items-center justify-between gap-4">
  <div>Left</div>
  <div>Right</div>
</div>

<div class="flex flex-col md:flex-row">
  <div class="md:w-1/3">Sidebar</div>
  <div class="md:w-2/3">Main</div>
</div>

<div class="grid grid-cols-3 gap-4">
  <div>1</div><div>2</div><div>3</div>
</div>

<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
  <div class="col-span-2">Wide</div>
</div>
```

## fa-arrows-left-right Spacing

```html
<div class="p-4 px-6 py-2">
<div class="m-4 mx-auto my-2">

<div class="space-x-4">
<div class="space-y-2">

<div class="px-4 sm:px-6 lg:px-8">
```

## fa-ruler Sizing

```html
<div class="w-full w-1/2 w-64 h-screen h-8">
<div class="max-w-7xl mx-auto">
<div class="min-h-screen min-w-0">

<div class="w-12 h-12">
<img class="size-16 rounded-full" />
```

## fa-font Typography

```html
<p class="text-sm text-lg text-2xl text-4xl">
<p class="font-normal font-medium font-bold">
<p class="text-left text-center text-right">
<p class="leading-tight leading-relaxed tracking-wide">

<h1 class="text-3xl font-bold tracking-tight text-gray-900">
<p class="text-gray-600 line-clamp-3">
```

## fa-palette Colors

```html
<div class="bg-white bg-gray-100 bg-gray-900">
<div class="text-blue-600 text-red-500 text-green-700">
<button class="bg-indigo-600 hover:bg-indigo-700 text-white">

<div class="bg-gradient-to-r from-blue-500 to-purple-600">
<div class="bg-opacity-50 bg-black/50">
```

## fa-border-all Borders

```html
<div class="border border-gray-300 rounded-lg">
<div class="border-t border-b border-l-4 border-blue-500">
<div class="rounded-full rounded-xl rounded-t-lg">
<div class="ring-2 ring-blue-500 ring-offset-2">
<div class="divide-y divide-gray-200">
```

## fa-cloud Shadows

```html
<div class="shadow shadow-md shadow-lg shadow-xl shadow-2xl">
<div class="shadow-none">

<div class="drop-shadow-md drop-shadow-lg">
<div class="shadow-[0_4px_20px_rgba(0,0,0,0.1)]">
```

## fa-mobile-screen Responsive Design

```html
<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6">

<p class="text-sm md:text-base lg:text-lg">

<div class="hidden md:block">
<div class="block md:hidden">

<div class="px-4 sm:px-6 md:px-8 lg:px-12">

// sm: 640px, md: 768px, lg: 1024px, xl: 1280px, 2xl: 1536px
```

## fa-hand-pointer State Variants

```html
<button class="bg-blue-500 hover:bg-blue-700 active:bg-blue-900 text-white">
<input class="border focus:border-blue-500 focus:ring-2 focus:ring-blue-200">

<div class="group">
  <div class="group-hover:opacity-100 opacity-0">Tooltip</div>
</div>

<a class="text-blue-600 hover:text-blue-800 visited:text-purple-600">

<li class="first:mt-0 last:mb-0 odd:bg-gray-50">
```

## fa-moon Dark Mode

```html
<div class="bg-white dark:bg-gray-900">
<p class="text-gray-900 dark:text-gray-100">

<div class="border-gray-200 dark:border-gray-700">

<button class="bg-gray-100 dark:bg-gray-800 text-gray-800 dark:text-gray-200">
```

## fa-sliders Custom Configuration

```js
// tailwind.config.js
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        brand: { 500: "#3b82f6", 600: "#2563eb" },
      },
      fontFamily: {
        sans: ["Inter", "system-ui", "sans-serif"],
      },
      spacing: {
        128: "32rem",
      },
    },
  },
  plugins: [],
};
```

## fa-wand-magic-sparkles @apply Directive

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer components {
  .btn-primary {
    @apply px-4 py-2 bg-blue-600 text-white rounded-lg font-medium
           hover:bg-blue-700 transition-colors duration-200;
  }
  .card {
    @apply bg-white dark:bg-gray-800 rounded-xl shadow-md p-6 border
           border-gray-200 dark:border-gray-700;
  }
}
```

## fa-film Animation

```html
<div class="animate-spin">
<div class="animate-ping">
<div class="animate-pulse">
<div class="animate-bounce">

<button class="transition-all duration-300 ease-in-out
              hover:scale-105 hover:shadow-lg">

<div class="transition-opacity duration-500 opacity-0 hover:opacity-100">

// Custom keyframes in tailwind.config.js
keyframes: {
  "fade-in": {
    "0%": { opacity: "0", transform: "translateY(10px)" },
    "100%": { opacity: "1", transform: "translateY(0)" },
  },
},
animation: {
  "fade-in": "fade-in 0.3s ease-out",
},
```

## fa-shapes Common Patterns

```html
<nav class="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b">
  <div class="max-w-7xl mx-auto px-4 flex items-center h-16">
    <a class="font-bold text-xl">Brand</a>
    <div class="ml-auto flex gap-6">
      <a class="hover:text-blue-600">Link</a>
    </div>
  </div>
</nav>

<div class="min-h-screen flex items-center justify-center bg-gray-50">
  <div class="w-full max-w-md p-8 bg-white rounded-2xl shadow-xl">
    <h2 class="text-2xl font-bold text-center mb-6">Sign In</h2>
    <form class="space-y-4">
      <input class="w-full px-4 py-3 rounded-lg border focus:ring-2 focus:ring-blue-500" />
      <button class="w-full py-3 bg-blue-600 text-white rounded-lg font-medium hover:bg-blue-700 transition">
        Submit
      </button>
    </form>
  </div>
</div>
```
