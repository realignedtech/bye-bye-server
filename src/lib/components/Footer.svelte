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
    <p class="mt-6 text-center text-sm text-slate-400">
      Bye Bye Server is maintained by Realigned Technologies Ltd, the team behind
      <a
        href="https://windshift.sh"
        target="_blank"
        rel="noopener noreferrer"
        class="text-slate-200 hover:text-white"
      >Windshift</a>. Listings are not curated by Realigned; anyone can submit an addition or update through
      <a
        href="https://github.com/realignedtech/bye-bye-server"
        target="_blank"
        rel="noopener noreferrer"
        class="text-slate-200 hover:text-white"
      >GitHub</a>.
    </p>
  </div>
</footer>
