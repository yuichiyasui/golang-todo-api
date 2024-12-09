/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'

// Create Virtual Routes

const SignInLazyImport = createFileRoute('/sign-in')()
const NewLazyImport = createFileRoute('/new')()
const TaskIdLazyImport = createFileRoute('/$taskId')()
const IndexLazyImport = createFileRoute('/')()
const SignUpEmailLazyImport = createFileRoute('/sign-up/email')()

// Create/Update Routes

const SignInLazyRoute = SignInLazyImport.update({
  path: '/sign-in',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/sign-in.lazy').then((d) => d.Route))

const NewLazyRoute = NewLazyImport.update({
  path: '/new',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/new.lazy').then((d) => d.Route))

const TaskIdLazyRoute = TaskIdLazyImport.update({
  path: '/$taskId',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/$taskId.lazy').then((d) => d.Route))

const IndexLazyRoute = IndexLazyImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/index.lazy').then((d) => d.Route))

const SignUpEmailLazyRoute = SignUpEmailLazyImport.update({
  path: '/sign-up/email',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/sign-up/email.lazy').then((d) => d.Route))

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexLazyImport
      parentRoute: typeof rootRoute
    }
    '/$taskId': {
      id: '/$taskId'
      path: '/$taskId'
      fullPath: '/$taskId'
      preLoaderRoute: typeof TaskIdLazyImport
      parentRoute: typeof rootRoute
    }
    '/new': {
      id: '/new'
      path: '/new'
      fullPath: '/new'
      preLoaderRoute: typeof NewLazyImport
      parentRoute: typeof rootRoute
    }
    '/sign-in': {
      id: '/sign-in'
      path: '/sign-in'
      fullPath: '/sign-in'
      preLoaderRoute: typeof SignInLazyImport
      parentRoute: typeof rootRoute
    }
    '/sign-up/email': {
      id: '/sign-up/email'
      path: '/sign-up/email'
      fullPath: '/sign-up/email'
      preLoaderRoute: typeof SignUpEmailLazyImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren({
  IndexLazyRoute,
  TaskIdLazyRoute,
  NewLazyRoute,
  SignInLazyRoute,
  SignUpEmailLazyRoute,
})

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/$taskId",
        "/new",
        "/sign-in",
        "/sign-up/email"
      ]
    },
    "/": {
      "filePath": "index.lazy.tsx"
    },
    "/$taskId": {
      "filePath": "$taskId.lazy.tsx"
    },
    "/new": {
      "filePath": "new.lazy.tsx"
    },
    "/sign-in": {
      "filePath": "sign-in.lazy.tsx"
    },
    "/sign-up/email": {
      "filePath": "sign-up/email.lazy.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
