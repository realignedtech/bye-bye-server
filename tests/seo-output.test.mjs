import assert from 'node:assert/strict';
import { existsSync, readFileSync } from 'node:fs';
import { test } from 'node:test';

function output(name) {
  return readFileSync(new URL(`../dist/${name}`, import.meta.url), 'utf8');
}

test('homepage ships meaningful content in the initial HTML response', () => {
  const html = output('index.html');

  assert.match(html, /<title>Bye Bye Server - Alternatives to Atlassian Products<\/title>/);
  assert.match(html, /<link rel="canonical" href="https:\/\/bye-bye-server\.com\/"/);
  assert.match(html, /<h1[^>]*>\s*Atlassian Alternatives, Compared\s*<\/h1>/);
  assert.match(html, /<h2[^>]*>Jira Software Alternatives<\/h2>/);
  assert.match(html, /<h2[^>]*>Jira Service Desk Alternatives<\/h2>/);
  assert.match(html, /<h2[^>]*>Confluence Alternatives<\/h2>/);
  assert.match(html, /href="https:\/\/github\.com\/realignedtech\/bye-bye-server"/);
  assert.match(html, /src="https:\/\/data\.windshift\.dev\/script\.js"/);
  assert.match(html, /data-website-id="be0c9ff8-cc28-4f6b-a046-a9902e6755ff"/);

  const crawlableLinks = html.match(/<a\b[^>]*href=/g) ?? [];
  assert.ok(
    crawlableLinks.length > 20,
    `expected more than 20 crawlable links, got ${crawlableLinks.length}`
  );
});

test('static output includes machine-readable files and a noindex 404', () => {
  const robots = output('robots.txt');
  const sitemap = output('sitemap.xml');
  const notFound = output('404.html');

  assert.match(robots, /^User-agent: \*$/m);
  assert.match(robots, /^Sitemap: https:\/\/bye-bye-server\.com\/sitemap\.xml$/m);
  assert.match(sitemap, /<loc>https:\/\/bye-bye-server\.com\/<\/loc>/);
  assert.match(notFound, /<meta name="robots" content="noindex"/);
  assert.match(notFound, /<h1>This page said goodbye\.<\/h1>/);
  assert.equal(existsSync(new URL('../dist/not-a-real-page', import.meta.url)), false);
});
