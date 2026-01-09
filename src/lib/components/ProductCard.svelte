<script>
  import { ExternalLink, Cloud, Server, Globe } from 'lucide-svelte';

  let { product } = $props();

  function getLicenseBadge(license) {
    const l = license.toLowerCase();
    if (l.includes('mit') || l.includes('bsd')) return 'bg-green-100 text-green-800';
    if (l.includes('gpl') || l.includes('agpl') || l.includes('lgpl')) return 'bg-blue-100 text-blue-800';
    if (l.includes('proprietary')) return 'bg-amber-100 text-amber-700';
    return 'bg-purple-100 text-purple-700';
  }

  function getHostingIcon(hosting) {
    const h = hosting.toLowerCase();
    if (h.includes('cloud') && h.includes('premises')) return Globe;
    if (h.includes('cloud')) return Cloud;
    return Server;
  }

  const HostingIcon = getHostingIcon(product.hosting);
</script>

<div class="bg-white border border-slate-200 rounded-xl p-6 transition-all duration-200 hover:-translate-y-1 hover:shadow-lg flex flex-col h-full">
  <div class="flex items-start justify-between mb-4">
    <h3 class="text-lg font-semibold text-gray-900">{product.title}</h3>
    <div class="p-2 bg-slate-50 rounded-md text-slate-400">
      <HostingIcon class="w-5 h-5" />
    </div>
  </div>

  <div class="mb-4">
    <span class="inline-flex items-center px-2.5 py-1 text-xs font-medium rounded-full uppercase tracking-wide {getLicenseBadge(product.license)}">{product.license}</span>
  </div>

  <dl class="space-y-3 text-sm flex-1">
    <div class="flex justify-between">
      <dt class="text-gray-500">Hosting</dt>
      <dd class="text-gray-700 font-medium text-right">{product.hosting}</dd>
    </div>
    <div class="flex justify-between">
      <dt class="text-gray-500">Subscription</dt>
      <dd class="text-gray-700 font-medium text-right">{product.subscription}</dd>
    </div>
    <div class="flex justify-between">
      <dt class="text-gray-500">Marketplace</dt>
      <dd class="text-gray-700 font-medium text-right">{product.marketplace}</dd>
    </div>
  </dl>

  <div class="mt-6 pt-4 border-t border-gray-100">
    <a href={product.link} target="_blank" rel="noopener noreferrer" class="inline-flex items-center gap-2 text-[#013373] font-medium transition-opacity hover:opacity-80">
      <span>Visit Website</span>
      <ExternalLink class="w-4 h-4" />
    </a>
  </div>
</div>
