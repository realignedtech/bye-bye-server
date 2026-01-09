<script>
  import CategorySection from '$lib/components/CategorySection.svelte';
  import { products } from '$lib/data/products.js';
  import { ArrowRight, GitBranch } from 'lucide-svelte';
</script>

<!-- Hero Section -->
<div class="hero-gradient relative overflow-hidden">
  <div class="hero-pattern absolute inset-0"></div>
  <div class="relative max-w-6xl mx-auto px-4 py-20 text-center">
    <h1 class="text-5xl md:text-6xl font-bold text-white mb-4">
      Atlassian Alternatives, Compared
    </h1>
    <p class="text-xl md:text-2xl text-white/90 max-w-2xl mx-auto mb-4">
      Find the right tools as Server and Data Center wind down.
    </p>
    <div class="mb-8">
      <a
        href="https://www.atlassian.com/blog/announcements/atlassian-ascend"
        target="_blank"
        rel="noopener noreferrer"
        class="text-sm font-medium text-white hover:opacity-80 transition-opacity"
      >Read the official announcement</a>
    </div>
    <a
      href="https://github.com/realignedtechnologies/bye-bye-server"
      target="_blank"
      rel="noopener noreferrer"
      class="inline-flex items-center gap-2 px-6 py-3 text-base font-semibold rounded-lg bg-white text-slate-900 hover:bg-slate-100 hover:-translate-y-0.5 transition-all cursor-pointer"
    >
      <GitBranch class="w-5 h-5" />
      Contribute on GitHub
      <ArrowRight class="w-4 h-4" />
    </a>
  </div>
</div>

<!-- Stats Bar -->
<div class="border-b border-gray-200 bg-white">
  <div class="max-w-6xl mx-auto px-4 py-6">
    <div class="grid grid-cols-3 gap-8 text-center">
      <div>
        <div class="text-3xl font-bold text-[#013373]">{products.jirasoftware.length}</div>
        <div class="text-sm text-slate-500 mt-1">Jira Alternatives</div>
      </div>
      <div>
        <div class="text-3xl font-bold text-[#013373]">{products.jiraservicedesk.length}</div>
        <div class="text-sm text-slate-500 mt-1">Service Desk Options</div>
      </div>
      <div>
        <div class="text-3xl font-bold text-[#013373]">{products.confluence.length}</div>
        <div class="text-sm text-slate-500 mt-1">Wiki Solutions</div>
      </div>
    </div>
  </div>
</div>

<!-- Product Categories -->
<div class="max-w-6xl mx-auto px-4 py-16">
  <CategorySection title="Jira Software Alternatives" products={products.jirasoftware} />
  <CategorySection title="Jira Service Desk Alternatives" products={products.jiraservicedesk} />
  <CategorySection title="Confluence Alternatives" products={products.confluence} />
</div>
