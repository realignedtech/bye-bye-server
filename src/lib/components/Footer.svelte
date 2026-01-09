<script>
  import { Heart } from 'lucide-svelte';
  const year = new Date().getFullYear();
</script>

<footer class="bg-slate-900 text-slate-400 py-12">
  <div class="max-w-6xl mx-auto px-4">
    <div class="flex flex-col md:flex-row justify-between items-center gap-4">
      <p class="text-sm">{year} Realigned Technologies Ltd</p>
      <p class="text-sm flex items-center gap-1">
        Made with <Heart class="w-4 h-4 text-red-500" /> for the Atlassian community
      </p>
    </div>
  </div>
</footer>
