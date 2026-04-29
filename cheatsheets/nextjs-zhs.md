---
title: Next.js
icon: fa-n
primary: "#000000"
lang: jsx
locale: zhs
---

## fa-folder-tree App Router Setup

```jsx
// app/layout.tsx
export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}

// app/page.tsx
export default function Home() {
  return <h1>Hello World</h1>;
}
```

## fa-layer-group Pages & Layouts

```jsx
// app/dashboard/layout.tsx
export default function DashboardLayout({ children }) {
  return (
    <div className="flex">
      <nav>Sidebar</nav>
      <main>{children}</main>
    </div>
  );
}

// app/dashboard/page.tsx
export default function DashboardPage() {
  return <h1>Dashboard</h1>;
}
```

## fa-server Server Components

```jsx
// app/posts/page.tsx (默认为 Server Component)
import { db } from "@/lib/db";

export default async function PostsPage() {
  const posts = await db.post.findMany();
  return (
    <ul>
      {posts.map(post => <li key={post.id}>{post.title}</li>)}
    </ul>
  );
}
```

## fa-display Client Components

```jsx
"use client";

import { useState } from "react";

export default function Counter() {
  const [count, setCount] = useState(0);
  return (
    <button onClick={() => setCount(c => c + 1)}>
      Count: {count}
    </button>
  );
}
```

## fa-route Routing

```jsx
// 动态路由: app/blog/[slug]/page.tsx
export default function BlogPost({ params }) {
  return <h1>Post: {params.slug}</h1>;
}

// Catch-all 路由: app/shop/[...slug]/page.tsx
export default function ShopPage({ params }) {
  return <p>Path: {params.slug.join("/")}</p>;
}

// 路由分组 (不影响 URL): app/(marketing)/about/page.tsx
```

## fa-spinner Loading & Error

```jsx
// app/dashboard/loading.tsx
export default function Loading() {
  return <div>Loading...</div>;
}

// app/dashboard/error.tsx
"use client";
export default function Error({ error, reset }) {
  return (
    <div>
      <p>{error.message}</p>
      <button onClick={reset}>Retry</button>
    </div>
  );
}

// app/not-found.tsx
export default function NotFound() {
  return <h1>404 - Page Not Found</h1>;
}
```

## fa-download Data Fetching

```jsx
// 服务端数据获取，不缓存
async function getUsers() {
  const res = await fetch("https://api.example.com/users", {
    cache: "no-store",
  });
  return res.json();
}

// 每 60 秒重新验证
async function getProducts() {
  const res = await fetch("https://api.example.com/products", {
    next: { revalidate: 60 },
  });
  return res.json();
}

// 并行数据获取
export default async function Page() {
  const [users, products] = await Promise.all([getUsers(), getProducts()]);
  return <Dashboard users={users} products={products} />;
}
```

## fa-paper-plane Server Actions

```jsx
// app/actions.ts
"use server";

export async function createPost(formData) {
  const title = formData.get("title");
  await db.post.create({ data: { title } });
  revalidatePath("/posts");
}

export async function deletePost(id) {
  await db.post.delete({ where: { id } });
  revalidatePath("/posts");
}

// app/posts/page.tsx
import { createPost } from "./actions";

export default function NewPost() {
  return (
    <form action={createPost}>
      <input name="title" />
      <button type="submit">Create</button>
    </form>
  );
}
```

## fa-plug API Routes

```jsx
// app/api/users/route.ts
import { NextResponse } from "next/server";

export async function GET(request) {
  const users = await db.user.findMany();
  return NextResponse.json(users);
}

export async function POST(request) {
  const body = await request.json();
  const user = await db.user.create({ data: body });
  return NextResponse.json(user, { status: 201 });
}

// app/api/users/[id]/route.ts
export async function DELETE(request, { params }) {
  await db.user.delete({ where: { id: params.id } });
  return new Response(null, { status: 204 });
}
```

## fa-shield Middleware

```jsx
// middleware.ts
import { NextResponse } from "next/server";
import { getToken } from "next-auth/jwt";

export async function middleware(request) {
  const token = await getToken({ req: request });
  if (!token) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  return NextResponse.next();
}

export const config = {
  matcher: ["/dashboard/:path*", "/admin/:path*"],
};
```

## fa-image Image Optimization

```jsx
import Image from "next/image";

export default function Avatar() {
  return (
    <Image
      src="/profile.jpg"
      alt="Profile"
      width={200}
      height={200}
      priority
    />
  );
}

export function HeroImage() {
  return (
    <Image
      src="https://cdn.example.com/hero.jpg"
      alt="Hero"
      fill
      className="object-cover"
      sizes="100vw"
    />
  );
}
```

## fa-tags Metadata & SEO

```jsx
// app/layout.tsx
export const metadata = {
  title: "My App",
  description: "A Next.js application",
  openGraph: {
    title: "My App",
    description: "A Next.js application",
    url: "https://example.com",
    siteName: "My App",
  },
};

// 页面级动态 metadata
export async function generateMetadata({ params }) {
  const post = await getPost(params.slug);
  return {
    title: post.title,
    description: post.excerpt,
  };
}
```

## fa-gear Environment Variables

```jsx
// 仅服务端 (无前缀)
const dbUrl = process.env.DATABASE_URL;

// 暴露给客户端 (NEXT_PUBLIC_ 前缀)
const apiKey = process.env.NEXT_PUBLIC_API_KEY;

// .env.local (git 忽略)
// DATABASE_URL=postgresql://...
// NEXT_PUBLIC_API_URL=https://api.example.com
```

## fa-rocket Deployment

```bash
# 构建并启动生产环境
next build && next start

# 静态导出
next build  # next.config.js 中设置 output: 'export'

# Docker
FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build
EXPOSE 3000
CMD ["npm", "start"]
```
