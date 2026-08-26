# Bye Bye Server

Atlassian is retiring their Server and Data Center products. This site helps you find alternatives.

Website: https://bye-bye-server.com

## Contributing

PRs welcome. To add or update a product, edit `src/lib/data/products.js`.

Requirements for listed products:
- Must support on-premises deployment (alternatives for Atlassian Server customers)
- Include accurate license information
- Provide a working website URL

## Development

### Prerequisites
- Node.js 22+
- Go 1.22+

### Run locally
```
npm install
npm run dev
```

### Build
```
npm run build
```

### Format/Lint
```
npm run format
npm run lint
```

## Deployment

Uses Docker for deployment:

```
docker compose up -d
```

## Tech Stack
- Frontend: SvelteKit/Svelte 5, statically prerendered with Vite and Tailwind CSS
- Backend: Go static file server with explicit 404 handling
